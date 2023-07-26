## Resposta
Levando em consideração o cenário apresentado, as modificações necessárias para atingir os requisitos não funcionais são:

### Disponibilidade de histórico offline
- **Banco de dados offline com cache em memória:** Para que pudéssemos dispor de um histórico de notícias mesmo com o banco de dados offline, o certo seria implementar um sistema de cache em memória como o Redis. Dessa forma, mesmo que o banco de dados relacional esteja offline, as informações mais recentes ainda poderão ser acessadas.

### Latência inferior a 20ms em 95% dos casos
Não possuo certeza quanto a alcançarmos essa latência de 20ms, mas aqui vão algumas medidas que poderiam ser adotadas para reduzir essa latência:

- **Redução de registros e otimização da tabela:** Seria necessária uma análise na tabela "history" para que identificássemos campos desnecessários ou que podem ser compactados para reduzir o tamanho dos registros. Além disso, devemos nos certificar de que todos os índices relevantes estejam devidamente criados e otimizados.

- **Compressão de dados:** Devemos também implementar técnicas de compressão de dados nos registros da tabela "history" para reduzir o tamanho físico do banco de dados e, consequentemente, melhorar a latência das operações de leitura.

## Buscas necessárias
[Compactação de dados](https://www.devmedia.com.br/artigo-sql-magazine-68-compactacao-de-dados-com-o-sql-server-2008/14300)