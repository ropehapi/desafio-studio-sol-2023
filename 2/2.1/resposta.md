## Resposta
Levando em consideração o modelo relacional apresentado e que as tabelas não possuem índices, posso citar como necessários para os cenários de busca mais comuns desse schema os seguintes índices:

- Tabela authors
    - id: Primary key
- Tabela genres
    - id: Primary key
- Tabela publishers
    - id: Primary key
- Tabela books
    - id: Primary Key
    - author_id: Key
    - genre_id: Key
    - publisher_id: Key
    - barcode: Unique
- Tabela users
    - id: Primary key
- Tabela books_ratings
    - id: Primary Key
    - user_id: Key
    - book_id: Key
- Tabela books_comments
    - id: Primary key
    - user_id: Key
    - book_id: Key


Aqui, acredito que o único índice que precise ser justificado seja o barcode. Sua utilidade se dá justamente pelo fato de um código de barras ser algo único como um identificador. Tornando útil a consulta por essa coluna em algumas ocasiões.

Sobre a redução de conjuntos de dados e a cardinalidade, acredito que todos os índices trabalharão muito bem pois quanto maior a cardinalidade de um valor, melhor o desempenho das consultas utilizando índices, e todos os índices foram colocados estratégicamente nos identificadores dos registros, valores esses que possuem uma cardinalidade gigantesca, tendo em vista que não se repetem.

 ## Buscas necessárias
 [Indexação de consultas SQL](https://www.devmedia.com.br/otimizacao-de-consultas-no-mysql/6178)