package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int64 = v1
    v3 := make(map[uint32]int)
    v4 := []uint32{65535, 4294967295}
    v5 := make(map[int]string)
    var v6 bool = (1e10 < (1e10 - float64((0.1 - 1e10))))
    if (!v6) {
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 int64 = (v1 << ((255) & 63))
            _ = fi7
            _ = v8
        }
        v2 = v1
        switch (("a" + ("a" + "\n\t")) + "\n\t") {
        case "hello":
            for fi9 := 0; fi9 < 13; fi9++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v10 int = (fi9 + fi9)
                _ = fi9
                _ = v10
            }
        case "\n\t":
            v1 -= 9223372036854775807
        }
    }
    v11 := GenS1{F0: 0}
    v2 -= -289
    _ = len(v5)
    v2 += v1
    var v12 bool = (v6 || (2147483647 != 42))
    var v13 uint16 = 32767
    v2 = v1
    fmt.Println(v12, v6, v6)
    var v14 int8 = 7
    v15 := v5[-2147483648]
    _ = len(v5)
    for fi16 := 0; fi16 < 7; fi16++ {
        v13 -= (v13 % (((v13 + (v13 % ((v13) | 1)))) | 1))
        _ = fi16
    }
    v3[1] = 0
    switch v15 {
    case "foo bar":
        if v12 {
            v17 := GenS1{F0: 42}
            var v18 bool = (v15 == v15)
            var v19 uint8 = 1
            _ = v17
            _ = v18
            _ = v19
        }
    case "":
        v20 := GenS1{F0: 127}
        _ = v20
    default:
        v13 = (v13 & v13)
    }
    v21 := GenS1{F0: 128}
    v14 -= (-(v14))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
