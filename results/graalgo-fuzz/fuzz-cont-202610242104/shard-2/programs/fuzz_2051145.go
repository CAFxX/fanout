package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int8 }

func main() {
    v1 := make(map[string]int)
    var v2 string = "\n\t"
    var v3 string = ((v2 + v2) + ("" + "hello"))
    var v4 uint = 42
    _ = len(v1)
    v4 = (v4 % (((v4 >> (((^(v4))) & 63))) | 1))
    var v5 uint = (v4 / (((v4 >> ((86) & 63))) | 1))
    fmt.Println(v2, v2, v2)
    v4 = v5
    v6 := []int{67, 127}
    var v7 float64 = 3.14159
    switch (v2 + v3) {
    case "":
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 uint8 = 255
            for fi10 := 0; fi10 < 8; fi10++ {
                v4 += 65535
                var v11 bool = true
                v12 := GenS1{F0: 136, F1: -1, F2: 0}
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi8
            _ = v9
        }
    case "\n\t":
        delete(v1, v3)
    }
    if true {
        for fi13 := 0; fi13 < 7; fi13++ {
            v3 = ""
            _ = fi13
        }
        v6 = append(v6, 4294967295)
        fmt.Println(v2, v5)
    } else {
        v2 = v3
        var v14 float64 = float64(-2.25)
        _ = v14
    }
    v15 := GenS1{F0: 42, F1: 94, F2: -1}
    if true {
        var v16 uint32 = 2
        _ = len(v1)
        v1["hello"] = 1
        _ = v16
    }
    var v17 int8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
