package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 0
    var v3 float32 = 0.1
    v4 := make(map[string]int)
    var v5 uint32 = 42
    v6 := GenS1{F0: -9223372036854775808, F1: 382}
    if false {
        switch ("\n\t" + "a") {
        case "\n\t":
            v7 := []int{}
            _ = v7
        case "":
            v8 := v4["a"]
            _ = v8
        default:
            if ((false || (!(2 != 4294967295))) || (255 == 19)) {
                var v9 uint = 515
                _ = v9
            }
        }
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        delete(v4, (("a" + "") + "\n\t"))
        v11 := v1["a"]
        v2 = (v2 & (v2 + 42))
        _ = fi10
        _ = v11
    }
    fmt.Println(v2)
    v5 = v5
    var v12 string = "hello"
    var v13 int64 = v2
    var v14 int = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
