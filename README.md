# Desafio DevOps - Korp

Este repositório contém a solução para o desafio técnico de estágio em DevOps. A infraestrutura foi totalmente projetada em containers utilizando Docker e Docker Compose, com proxy reverso NGINX e monitoramento via Prometheus.

## 🚀 Como Executar o Projeto

Toda a infraestrutura (configuração de pastas, build da imagem, subida dos containers e proxy) foi automatizada para ser provisionada com **um único comando Ansible**.

Na raiz do projeto dentro do ambiente Linux/WSL, execute:

```bash
ansible-playbook playbook.yml
