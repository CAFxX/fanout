package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint16) uint64 {
    var v7 int8 = 42
    fmt.Println(p0, v7)
    switch -535 {
    case -1:
        if (!(128 == 180)) {
            v8 := []int{-257}
            _ = v8
        }
    case 32768:
        v9 := GenS1{F0: (v7 ^ v7)}
        _ = v9
    default:
        var v10 int64 = -1
        _ = v10
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        for fi12 := 0; fi12 < 5; fi12++ {
            p0 = (p0 + p0)
            p0 = p0
            switch "hello" {
            case "hello":
                var v13 uint = 18446744073709551615
                _ = v13
            }
            _ = fi12
        }
        p0 = p0
        _ = fi11
    }
    _ = p0
    _ = v7
    return 190
}

func main() {
    v1 := []string{}
    var v2 uint = 0
    var v3 uint32 = 42
    v4 := make(map[uint32]int)
    var v5 uint16 = 65535
    v6 := make(map[uint32]int)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v4)
    if false {
        v5 = 360
    } else {
        var v14 int16 = -598
        for fi15 := 0; fi15 < 15; fi15++ {
            v2 -= (v2 * (^(v2)))
            fmt.Println(v3, v2, v2)
            switch fi15 {
            case 42:
                v4[v3] = int(v5)
            case 32767:
                if ((true && false) || false) {
                    var v16 float64 = 0.0
                    _ = v16
                } else {
                    v2 -= 788
                }
            default:
                switch "" {
                case "hello":
                    if false {
                        v14 += v14
                        v17 := GenS1{F0: 127}
                        _ = v17
                    }
                case "":
                    v18 := GenS1{F0: 100}
                    _ = v18
                case "foo bar":
                    var v19 int32 = -666
                    _ = v19
                default:
                    for fi20 := 0; fi20 < 2; fi20++ {
                        if len(v1) > 0 {
                            v1[0] = "\n\t"
                        }
                        _ = fi20
                    }
                }
            }
            _ = fi15
        }
        _ = v14
    }
    var v21 int64 = 1
    v5 = 65535
    fmt.Println(v2)
    if ((v5 % ((42) | 1)) == v5) {
        var v22 float32 = 3.14159
        if len(v1) > 0 {
            v1[0] = "hello"
        }
        _ = v22
    } else {
        delete(v6, v3)
    }
    v21 += v21
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
