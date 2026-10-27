package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 bool }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 979
    var v3 uint16 = 42
    var v4 uint64 = 0
    v5 := []uint32{929, 4294967295, 1}
    v5 = append(v5, uint32(v4))
    v2 += (^(v2))
    fmt.Println(v3, v2)
    if true {
        switch 34 {
        case -419:
            if true {
                v2 = v2
                for fi6 := 0; fi6 < 7; fi6++ {
                    _ = len(v1)
                    v4 = (v4 ^ v4)
                    v3 -= v3
                    _ = fi6
                }
                for fi7 := 0; fi7 < 12; fi7++ {
                    _ = len(v5)
                    var v8 uint32 = 0
                    for fi9 := 0; fi9 < 19; fi9++ {
                        fmt.Println(v4)
                        _ = fi9
                    }
                    _ = fi7
                    _ = v8
                }
            }
        }
    } else {
        v3 += 32767
    }
    v10 := GenS1{F0: v4, F1: uint(v3), F2: (true && ((("" + "") + "foo bar") >= ("foo bar" + "")))}
    v11 := GenS2{F0: 4294967295, F1: false}
    v3 -= v2
    if true {
        v12 := GenS2{F0: int(v3), F1: (((!true) || false) && (3.14159 < 1e10))}
        _ = v12
    }
    var v13 uint = 4294967295
    var v14 float32 = ((0.0 + (1e10 / (3.14159 / 0.1))) + 1e10)
    var v15 string = "a"
    switch (("a" + "hello") + v15) {
    case "\n\t":
        v16 := GenS2{F0: -9223372036854775808, F1: ((^(v4)) <= (v4 & v4))}
        _ = v16
    case "":
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 int8 = -1
            _ = fi17
            _ = v18
        }
    case "hello":
        if (uint32(v2) >= 703) {
            v13 = uint(v4)
            var v19 bool = true
            if v19 {
                if (1 != uint8(v13)) {
                    if len(v5) > 0 {
                        v5[0] = 1
                    }
                }
                switch -2147483648 {
                case 149:
                    var v20 uint64 = 25
                    _ = v20
                case 479:
                    if ((!("hello" == "")) || (true && v19)) {
                        v15 = v15
                        v21 := GenS3{F0: 223}
                        var v22 bool = (0.1 <= 3.14159)
                        _ = v21
                        _ = v22
                    }
                default:
                    v5 = append(v5, 700)
                }
                var v23 uint16 = v2
                _ = v23
            }
            _ = v19
        }
    default:
        var v24 uint64 = (v4 * v4)
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
