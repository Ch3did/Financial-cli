#!/bin/bash

function add_category() {
    NAME="$1"
    DESCRIPTION="$2"
    EXPECTED="$3"

    echo -e "${NAME}\n${DESCRIPTION}\n${EXPECTED}" | cx add-category
}

add_category "Casa" "Gastos relacionados à manutenção da casa, como aluguel, contas fixas (água, luz, internet), e serviços domésticos." -2000.00
add_category "Alimentos" "Despesas com alimentação em geral, incluindo supermercado, delivery, refeições fora de casa e lanches ocasionais." -600.00
add_category "Eventos" "Gastos com lazer e entretenimento como festas, eventos, ingressos, bebidas alcoólicas e saídas sociais." -500.00
add_category "Transporte" "Custos com deslocamento, como corridas de Uber e 99, transporte público, metrô e eventuais manutenções ou pedágios." -300.00
add_category "Saúde" "Investimentos em saúde e bem-estar, como consultas médicas, remédios, exames, tratamentos e uso de produtos como minoxidil." -1000.00
add_category "Pagamentos" "Valores recebidos como salário, pagamentos por serviços, reembolsos ou outras fontes de renda recorrente ou pontual." 5000.00
add_category "Outros" "Categoria genérica para transações não classificadas ou ajustes manuais feitos posteriormente." -100.00
