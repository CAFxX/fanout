package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int16 }

func main() {
    var v1 bool = (!(true && true))
    var v2 uint8 = 127
    v3 := []uint32{2147483648, 1, 4294967295, 237}
    var v4 float64 = 3.14159
    var v5 uint = 42
    v6 := GenS1{F0: 1, F1: v2, F2: 0}
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            fmt.Println(v1)
            _ = fi8
        }
        _ = fi7
    }
    _ = len(v3)
    fmt.Println(v1, v5, v1)
    for fi9 := 0; fi9 < 15; fi9++ {
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := GenS1{F0: 0, F1: v2, F2: -1}
            var v12 uint64 = 4294967295
            _ = fi10
            _ = v11
            _ = v12
        }
        v1 = v1
        _ = fi9
    }
    if v1 {
        var v13 bool = false
        _ = v13
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v14 uint16 = 42
        for fi15 := 0; fi15 < 12; fi15++ {
            v4 -= v4
            for fi16 := 0; fi16 < 19; fi16++ {
                var v17 bool = (42 < 1)
                var v18 string = ("\n\t" + "")
                var v19 bool = v1
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            _ = fi15
        }
        _ = v14
    }
    var v20 int = 128
    if len(v3) > 0 {
        v3[0] = 24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
