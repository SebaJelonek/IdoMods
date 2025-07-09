curl --request POST \
 --url https://zooart6.yourtechnicaldomain.com/api/admin/v5/orders/orders/search \
 --header 'X-API-KEY: IDOSELL_API_KEY' \
 --header 'accept: application/json' \
 --header 'content-type: application/json' \
 --data '
{"params": {"ordersStatuses": ["new","finished","false","lost","on_order","packed","ready","canceled","payment_waiting","delivery_waiting","suspended","joined","finished_ext"]}}'
