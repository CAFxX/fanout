package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = (3.14159 / (3.14159 + float64(-2.25)))
    v3 := []uint32{2147483647, 42, 2147483647, 2147483648}
    v4 := make(map[int]bool)
    var v5 uint8 = 255
    var v6 int16 = 39
    switch "hello" {
    case "hello":
        fmt.Println(v6, v5)
    case "\n\t":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "a":
        if len(v3) > 0 {
            v3[0] = 953
        }
    default:
        var v7 int32 = 0
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 float32 = 0.0
        _ = fi8
        _ = v9
    }
    v5 += (^(v5))
    v1["a"] = (888 != 922)
    v10 := GenS1{F0: "hello"}
    if false {
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        v6 += 0
        var v11 int16 = (v6 / (((v6 | (v6 | v6))) | 1))
        _ = v11
    }
    fmt.Println(v5, v6)
    fmt.Println(v5, v5, v2)
    if ((v5 + (^(v5))) > (v5 << ((42) & 63))) {
        fmt.Println(v5, v5, v2)
    } else {
        for fi12 := 0; fi12 < 12; fi12++ {
            for fi13 := 0; fi13 < 18; fi13++ {
                v6 = v6
                v6 = 13
                _ = fi13
            }
            _ = fi12
        }
        _ = len(v4)
        for fi14 := 0; fi14 < 4; fi14++ {
            fmt.Println(v5)
            var v15 bool = (true && ("a" > "hello"))
            _ = fi14
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        v17 := v1["hello"]
        _ = fi16
        _ = v17
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
