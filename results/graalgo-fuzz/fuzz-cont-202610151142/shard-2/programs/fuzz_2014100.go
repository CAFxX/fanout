package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint, p1 int32) uint64 {
    v4 := []int{626, 128, 2147483647, -9223372036854775808}
    var v5 bool = true
    if v5 {
        var v6 float32 = 3.14159
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 int = fi7
            v9 := GenS1{F0: 850, F1: int8(fi7), F2: -1}
            v10 := GenS1{F0: 1, F1: -117, F2: 74}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        _ = v6
    } else {
        v5 = v5
    }
    fmt.Println(p0, p1)
    fmt.Println(p1, v5, v5)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 42
    v3 := make(map[uint32]bool)
    v2 -= v2
    fmt.Println(v2)
    if false {
        if false {
            var v11 string = ""
            var v12 float32 = 1e10
            v11 = "a"
            _ = v11
            _ = v12
        }
        v2 = 65535
        var v13 int64 = -1
        _ = v13
    }
    var v14 int = -9223372036854775808
    v2 = v2
    v14 = v14
    v2 = (v2 & (^(v2)))
    for fi15 := 0; fi15 < 7; fi15++ {
        if (true && true) {
            var v16 int16 = 559
            _ = v16
        } else {
            for fi17 := 0; fi17 < 13; fi17++ {
                if true {
                    fmt.Println(v14, fi15, v14)
                    fmt.Println(fi15, v14)
                    v18 := GenS2{F0: v2}
                    if true {
                        var v19 int64 = 42
                        fmt.Println(v2, v2)
                        switch fi17 {
                        case 2147483647:
                            var v20 float32 = 1e10
                            _ = v20
                        default:
                            _ = len(v3)
                        }
                        _ = v19
                    } else {
                        v2 -= (v2 ^ (v2 | (v2 ^ 277)))
                    }
                    _ = v18
                }
                var v21 uint32 = 2
                switch "\n\t" {
                case "hello":
                    var v22 uint16 = 0
                    _ = v22
                case "\n\t":
                    var v23 uint8 = 42
                    _ = v23
                case "foo bar":
                    delete(v1, v21)
                default:
                    if false {
                        v2 = v2
                        fmt.Println(fi15, fi15)
                    }
                }
                _ = fi17
                _ = v21
            }
        }
        var v24 int64 = 1
        _ = fi15
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
}
