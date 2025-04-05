// 路線作成1
CREATE (tokyo:Station {name: '東京'})
CREATE (shinagawa:Station {name: '品川'})
CREATE (ebisu:Station {name: '恵比寿'})
CREATE (shibuya:Station {name: '渋谷'})
CREATE (shinjuku:Station {name: '新宿'})
CREATE (ikebukuro:Station {name: '池袋'})
CREATE (sugamo:Station {name: '巣鴨'})
CREATE (komagome:Station {name: '駒込'})
CREATE (otsuka:Station {name: '大塚'})
CREATE (mejiro:Station {name: '目白'})
CREATE (takadanobaba:Station {name: '高田馬場'})
CREATE (waseda:Station {name: '早稲田'})

// 路線経路作成1
CREATE (tokyo)-[:CONNECTED_TO {travel_time: 5}]->(shinagawa)
CREATE (shinagawa)-[:CONNECTED_TO {travel_time: 5}]->(ebisu)
CREATE (ebisu)-[:CONNECTED_TO {travel_time: 5}]->(shibuya)
CREATE (shibuya)-[:CONNECTED_TO {travel_time: 5}]->(shinjuku)
CREATE (shinjuku)-[:CONNECTED_TO {travel_time: 5}]->(ikebukuro)
CREATE (ikebukuro)-[:CONNECTED_TO {travel_time: 5}]->(sugamo)
CREATE (sugamo)-[:CONNECTED_TO {travel_time: 5}]->(komagome)
CREATE (komagome)-[:CONNECTED_TO {travel_time: 5}]->(otsuka)
CREATE (otsuka)-[:CONNECTED_TO {travel_time: 5}]->(mejiro)
CREATE (mejiro)-[:CONNECTED_TO {travel_time: 5}]->(takadanobaba)
CREATE (takadanobaba)-[:CONNECTED_TO {travel_time: 5}]->(waseda)

// 路線作成2
CREATE (ocha:Station {name: '御茶ノ水'})
CREATE (kanda:Station {name: '神田'})
CREATE (hatchobori:Station {name: '八丁堀'})
CREATE (kayabacho:Station {name: '茅場町'})
CREATE (kyobashi:Station {name: '京橋'})
CREATE (ginza:Station {name: '銀座'})

// 路線経路作成2
CREATE (shinjuku)-[:CONNECTED_TO {travel_time: 3}]->(ocha)
CREATE (ocha)-[:CONNECTED_TO {travel_time: 3}]->(kanda)
CREATE (kanda)-[:CONNECTED_TO {travel_time: 2}]->(tokyo)
CREATE (tokyo)-[:CONNECTED_TO {travel_time: 3}]->(hatchobori)
CREATE (hatchobori)-[:CONNECTED_TO {travel_time: 3}]->(kayabacho)
CREATE (kayabacho)-[:CONNECTED_TO {travel_time: 3}]->(kyobashi)
CREATE (kyobashi)-[:CONNECTED_TO {travel_time: 3}]->(ginza)

// 2路線の交差点作成
CREATE (shinjuku)-[:CONNECTED_TO {travel_time: 2}]->(tokyo)
