package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 string }

func main() {
    var v1 float32 = 3.14159
    var v2 int8 = -1
    var v3 int16 = -547
    v4 := []uint32{65535, 4294967295}
    var v5 int16 = 1000
    var v6 uint16 = 603
    for fi7 := 0; fi7 < 18; fi7++ {
        if (1e10 < 1.5) {
            v8 := GenS1{F0: 891, F1: (v6 % ((v6) | 1)), F2: "a"}
            _ = v8
        } else {
            fmt.Println(v1)
            var v9 int = -220
            v10 := GenS1{F0: (v5 - v5), F1: v6, F2: ("a" + "")}
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    var v11 int8 = v2
    var v12 bool = (4294967295 > 407)
    switch -56 {
    case 42:
        v6 = v6
    case -2147483648:
        v2 += (v2 & (-(v11)))
    default:
        var v13 float32 = ((1.5 / v1) / 1.0)
        _ = v13
    }
    var v14 uint64 = 4294967295
    v6 -= v6
    var v15 string = ("foo bar" + ("hello" + "foo bar"))
    fmt.Println(v3, v1)
    v11 += v11
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v16 int16 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
