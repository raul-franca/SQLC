# SQLC
SQLC é um compilador de SQL para Go que permite aos desenvolvedores trabalhar com bancos de dados de forma mais segura e eficiente. Ele gera código Go a partir de consultas SQL, garantindo que essas consultas estejam corretas em tempo de compilação, o que reduz a chance de erros em tempo de execução.

Principais Características do SQLC:
Segurança de Tipos: SQLC gera código Go fortemente tipado com base nas consultas SQL, facilitando o trabalho com registros de banco de dados de maneira segura.
Suporte a Vários Bancos de Dados: SQLC suporta PostgreSQL, MySQL, SQLite, entre outros, permitindo sua utilização em diversos tipos de projetos.
Geração de Código: SQLC converte consultas SQL escritas em arquivos .sql em código Go que pode ser usado diretamente na aplicação.
ORM Mínimo: Diferente dos ORMs tradicionais, SQLC não abstrai as consultas SQL, mas se concentra em tornar as consultas SQL mais seguras e fáceis de usar em Go.
Integração Simples: SQLC pode ser integrado em projetos Go com pouca configuração, funcionando bem com módulos Go.
Resumo do Fluxo de Trabalho com SQLC:
Escreva as Consultas SQL: Defina suas consultas em arquivos .sql.
Gere o Código Go: Execute o SQLC para gerar o código Go correspondente.
Use o Código Gerado: Importe o pacote Go gerado e utilize as funções e tipos fornecidos no seu código.
SQLC é uma ferramenta útil para desenvolvedores Go que preferem escrever SQL puro, garantindo interações mais seguras e eficientes com o banco de dados.


# SQLC
https://github.com/golang-migrate/migrate

Ver tambem SQLX