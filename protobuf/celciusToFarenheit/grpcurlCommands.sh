# To list all the services hosted on "localhost:9092"
grpcurl --plaintext localhost:9092 list
# Response : 
# TempratureConversion
# grpc.reflection.v1alpha.ServerReflection

# To listy all the methods on service TempratureConversion
grpcurl --plaintext localhost:9092 list TempratureConversion
# Response : 
# TempratureConversion.ConvertFromCelsiusToFarenheit
# TempratureConversion.ConvertFromFarenheitToCelsius

# To desctibe the i/p and o/p on method  "TempratureConversion.ConvertFromCelsiusToFarenheit"
grpcurl --plaintext localhost:9092 describe TempratureConversion.ConvertFromCelsiusToFarenheit
# Response : 
# TempratureConversion.ConvertFromCelsiusToFarenheit is a method:
# rpc ConvertFromCelsiusToFarenheit ( .Temprature ) returns ( .Temprature );

# To describe the details of "Temprature"
grpcurl --plaintext localhost:9092 describe .Temprature
# Response : 
# Temprature is a message:
# message Temprature {
#   double Temprature = 1;
# }

# To request the API and get the data
grpcurl --plaintext -d '{"Temprature":0}' localhost:9092 TempratureConversion.ConvertFromCelsiusToFarenheit
# Response : 
# {
#   "Temprature": 32
# }