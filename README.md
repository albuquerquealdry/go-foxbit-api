
# 📄 Documentação da API de Cálculo FoxBit 🦊🧡

## 1. Descrição

API de cálculo que fornece operações básicas (soma, subtração, multiplicação e divisão) e endpoints de health para probes do Kubernetes.

Sobre o código:

A aplicação segue o **padrão SOLID**, que são princípios de design de software que ajudam a manter o código **organizado, modular e fácil de estender ou modificar**.
No contexto desta API:

* Cada operação está separada em **service**, mantendo responsabilidades isoladas (**Single Responsibility**).
* Os controllers usam interfaces mínimas e retornos padronizados, facilitando a substituição ou adição de novas operações (**Open/Closed**).
* O código é modular, promovendo baixo acoplamento e alta coesão, permitindo que novos endpoints ou serviços matemáticos sejam adicionados sem impactar funcionalidades existentes.

Para executar localmente a aplicação é necessário:

1. Golang em versão **1.24.0+**
2. Instalar as dependências do código (`go get github.com/gin-gonic/gin`)
3. Executar o `main.go` que está dentro de `cmd`

Localmente, podemos testar todos os endpoints documentados abaixo.

### 1.1 Informações gerais sobre a app

* **Base URL:** `http://<host>:8080`
* **Formato de resposta:** JSON
* **Probes de health:** `/health/liveness` e `/health/readiness`

---

## 2. Endpoints

### 2.1 Operações Matemáticas

| Método | Endpoint   | Parâmetros             | Descrição       | Exemplo                                                |
| ------ | ---------- | ---------------------- | --------------- | ------------------------------------------------------ |
| GET    | `/api/sum` | `term_one`, `term_two` | Soma os valores | `/api/sum?term_one=10&term_two=5` → `{"result":15}` |
| GET    | `/api/sub` | `term_one`, `term_two` | Subtração       | `/api/sub?term_one=10&term_two=5` → `{"result":5}`  |
| GET    | `/api/mul` | `term_one`, `term_two` | Multiplicação   | `/api/mul?term_one=10&term_two=5` → `{"result":50}` |
| GET    | `/api/div` | `term_one`, `term_two` | Divisão         | `/api/div?term_one=10&term_two=5` → `{"result":2}`  |

> ⚠️ Divisão por zero retorna erro:
> `{"error":"divisão por zero não permitida"}`

### 2.2 Health Checks

| Método | Endpoint            | Descrição                                                | Exemplo              |
| ------ | ------------------- | -------------------------------------------------------- | -------------------- |
| GET    | `/health/liveness`  | Verifica se a aplicação está viva                        | `{"status":"alive"}` |
| GET    | `/health/readiness` | Verifica se a aplicação está pronta para receber tráfego | `{"status":"ready"}` |

---

## 3. Kubernetes

Para fazer o deploy no Kubernetes, basta aplicar os manifestos que estão no diretório `k8s`.

Os manifestos contêm:

* **Deployment**: já apontando para a imagem do registry público.
* **Service**: tipo ClusterIP, expondo a porta da aplicação.

### Comandos úteis:

Para aplicar e visualizar os recursos:

```bash
kubectl apply -f ./k8s 
kubectl get pods
kubectl get service
kubectl get deployment
kubectl port-forward svc/foxbit-api-calc-service 8000:8000
```
 Para remover:

```bash
kubectl delete -f ./k8s 
```


Após o port-forward, é possível consumir os endpoints localmente via browser ou `curl`.

> ⚠️ A aplicação responde apenas por **ClusterIP**, por isso o port-forward é necessário para testes locais.

---

## 4. Testes de unidade

Para rodar os testes:

```bash
go test ./... -cover
```

Exemplo de saída:

```
ok      foxbit-calc-api/src/service     (cached)        coverage: 100.0% of statements
```

---
