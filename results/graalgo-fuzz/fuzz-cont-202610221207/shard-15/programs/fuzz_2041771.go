package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int32 }
type GenS2 struct { F0 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 0
    v3 := []uint32{65535, 42}
    switch "" {
    case "a":
        if false {
            v2 += (v2 & (v2 & v2))
            if (!true) {
                v2 = (v2 * 544)
                for fi4 := 0; fi4 < 13; fi4++ {
                    v5 := GenS1{F0: "", F1: -1, F2: -2147483648}
                    _ = fi4
                    _ = v5
                }
                switch 0 {
                case 32768:
                    if false {
                        var v6 uint32 = 105
                        for fi7 := 0; fi7 < 11; fi7++ {
                            v8 := GenS2{F0: false}
                            var v9 int32 = 42
                            var v10 uint64 = 436
                            _ = fi7
                            _ = v10
                            _ = v8
                            _ = v9
                        }
                        v2 = -9223372036854775808
                        _ = v6
                    }
                default:
                    for fi11 := 0; fi11 < 11; fi11++ {
                        _ = len(v1)
                        _ = fi11
                    }
                }
            }
        }
    case "foo bar":
        var v12 int8 = 0
        _ = v12
    case "\n\t":
        var v13 uint32 = 2147483647
        _ = v13
    }
    _ = len(v3)
    var v14 int = -783
    fmt.Println(v2)
    v2 -= 42
    if len(v3) > 0 {
        v3[0] = 100
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 11; fi16++ {
            switch ("a" + "") {
            case "a":
                var v17 uint64 = 42
                v14 = v14
                _ = v17
            case "hello":
                var v18 string = ""
                _ = v18
            default:
                switch "\n\t" {
                case "":
                    v19 := GenS2{F0: true}
                    _ = v19
                default:
                    fmt.Println(fi16)
                }
            }
            _ = fi16
        }
        var v20 string = "foo bar"
        _ = fi15
        _ = v20
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        _ = len(v3)
        v3 = append(v3, 268)
        fmt.Println(v2)
        _ = fi21
    }
    v22 := GenS2{F0: true}
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 float32 = ((1.5 + 0.1) * 0.0)
        _ = fi23
        _ = v24
    }
    var v25 string = ("\n\t" + ("a" + "\n\t"))
    v26 := GenS2{F0: (-1 < 100)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
}
