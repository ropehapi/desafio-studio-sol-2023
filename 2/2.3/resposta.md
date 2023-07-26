## Resposta
Levando em consideração a query apresentada, consigo listar três medidas de otimização do banco de dados que podem ser tomadas.

### 1 - Índices adequados
- Prós: Criar índices nas colunas usadas nas cláusulas WHERE e ORDER BY pode acelerar a pesquisa e a filtragem dos dados necessários para a agregação.
- Contras: Índices podem ocupar espaço em disco e podem levar a uma leve degradação no desempenho das operações de escrita (INSERT, UPDATE, DELETE). Além disso, índices devem ser escolhidos com cuidado, pois índices em excesso podem diminuir o desempenho geral do banco de dados.

### 2 - Índices Compostos
- Prós: Criar índices compostos que abrangem várias colunas usadas na consulta pode melhorar o desempenho, especialmente quando várias colunas são usadas em conjunto nas cláusulas WHERE.
- Contras: Como com qualquer índice, é importante encontrar um equilíbrio, pois índices em excesso podem levar a uma sobrecarga de indexação e aumento do espaço em disco.

### 3 - Materialized views
- Prós: Uma visão materializada pode pré-calcular a agregação e armazená-la em uma tabela, o que acelera a consulta, especialmente se os dados mudam lentamente. Isso reduz a necessidade de recalcular a agregação repetidamente.
- Contras: As visões materializadas ocupam espaço em disco e requerem esforço extra para manter os dados atualizados à medida que as classificações dos livros mudam.

## Buscas necessárias
[Vantagens e desvantagens dos indíces](https://pt.stackoverflow.com/questions/35088/quais-as-vantagens-e-desvantagens-do-uso-de-índices-em-base-de-dados)

[Índices compostos](https://codigofonte.org/indice-composto-mysql/#:~:text=Um%20índice%20composto%20é%20pelo,GROUP%20BY%20e%20MIN%2FMAX%20.)

[Materialized views](https://learn.microsoft.com/en-us/sql/t-sql/statements/create-materialized-view-as-select-transact-sql?view=azure-sqldw-latest)