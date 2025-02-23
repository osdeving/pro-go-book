# Workflow envolvendo **`git worktree`**

## **📌 Cenário: Você não tem nada, nem local, nem remoto**

### **1️⃣ Criar o repositório remoto**

Você pode criar o repositório diretamente no [GitHub](https://github.com), mas aqui usei o GitHub CLI:

```sh
gh repo create meu-repo --private
```

### **2️⃣ Clonar o repositório vazio para trabalhar localmente**

```sh
git clone git@github.com:<seu-usuario>meu-repo.git
cd meu-repo
```

🚨 O repositório está vazio, **faça pelo menos um commit inicial.**

```sh
echo "# Meu Projeto" > README.md
git add README.md
git commit -m "chore: projeto inicial"
git push -u origin main
```

---

### **3️⃣ Criar uma pasta no repositório que será também uma branch **

```sh
git worktree add -b feature/experimentos experimentos main
```

---

### **4️⃣ Criando outra pasta para **

```sh
git worktree add -b feature/loucuras loucuras main
```

Agora a estrutura está assim:

```
.
├── README.md
├── experimentos
│   └── README.md
└── loucuras
    └── README.md
```

Cada pasta representa um branch diferente.
Se fizer cd experimentos, você estará na branch feature/experimentos.
Se fizer cd loucuras, você estará na branch feature/loucuras.

---

### **5️⃣ Trabalhando em `feature/experimentos`**

```sh
cd experimentos
echo "## Experimentos " >> README.md
git add .
git commit -m "feat: adicionou readme"
git push -u origin feature/experimentos
cd ..
```

---

### **6️⃣ Trabalhando em `feature/loucuras`**

```sh
cd loucuras
echo "# Loucuras" > README.md
git add .
git commit -m "feat: adicionou readme"
git push -u origin feature/loucuras
cd ..
```

---

### **7️⃣ Trabalhando na `main`**

Aqui, caso você faça um git add . você estará adicionando os arquivos em todas as pastas.
Você pode colocar as subpastas no .gitignore ou fazer git add apenas dentro de cada pasta.

```sh
echo "## Estrutura de Pastas" >> README.md
git add README.md
git commit -m "docs: editar README"
git push
```

---

## **📌 Resumo**

- ✅ **Agora, ao entrar em cada pasta (`cd minha-pasta`), o Git automaticamente entra na branch correspondente.**
- ✅ Cada commit dentro de uma pasta **vai para a branch correta automaticamente**.

🚀 **Agora você pode ler aquele livro maroto e separar por capítulos. Ou você pode estruturar seu curso por pastas, etc..** 🚀🔥

