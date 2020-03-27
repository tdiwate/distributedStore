# distributedStore

The objective is to store a key-value pair in a distributed store based on the key.

For simplicity, we can assume the following:

The stores are held in-mem
There are 10 stores
Key is a string type
Value can be of any type.
You will write 3 functions

Write(key, value)
Read (key) //returns value
Dump() //Dumps all the keys and values by store acrosss all stores
Extra points for accounting for the following aspects: - Concurrency (Any means neccesary) - Performance - Reliability (What would happen on a store failure)
