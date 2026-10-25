package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
func genFn0(p0 int8, p1 int8, p2 uint8) uint8 {
    fmt.Println(p2, p1)
    if (32767 >= -1) {
        var v7 int32 = 550
        _ = v7
    } else {
        v8 := GenS1{F0: 571, F1: p2}
        var v9 uint = 18446744073709551615
        _ = v8
        _ = v9
    }
    v10 := []int{1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 1
}

func main() {
    v1 := []string{}
    v2 := make(map[int]int)
    v3 := make(map[uint32]int)
    var v4 int64 = -1
    var v5 uint8 = 81
    var v6 float32 = 0.0
    v11 := GenS1{F0: 18446744073709551615, F1: v5}
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := GenS2{F0: -2147483648}
        _ = fi12
        _ = v13
    }
    v4 += (v4 ^ v4)
    _ = len(v1)
    delete(v3, 2147483648)
    var v14 int32 = -516
    switch 127 {
    case 741:
        v14 += 284
    case 882:
        v14 -= v14
    case -1:
        var v15 uint8 = v5
        _ = v15
    }
    if true {
        v6 -= ((v6 - v6) + v6)
        fmt.Println(v5, v5)
    }
    v16 := GenS1{F0: 42, F1: 31}
    for fi17 := 0; fi17 < 8; fi17++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi18 := 0; fi18 < 19; fi18++ {
            fmt.Println(v14, v6, v14)
            _ = len(v1)
            switch "\n\t" {
            case "a":
                switch ("\n\t" + "\n\t") {
                case "a":
                    v19 := GenS1{F0: 136, F1: 249}
                    _ = v19
                case "":
                    if false {
                        var v20 uint64 = 978
                        _ = v20
                    }
                    var v21 int16 = 32767
                    v6 = ((v6 - v6) / (v6 / v6))
                    _ = v21
                default:
                    v14 = v14
                }
            case "\n\t":
                v22 := v3[4294967295]
                _ = v22
            }
            _ = fi18
        }
        _ = fi17
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        v24 := GenS2{F0: v14}
        switch fi23 {
        case 435:
            var v25 int8 = 0
            _ = v25
        default:
            if true {
                var v26 int64 = v4
                fmt.Println(fi23, v4, v4)
                _ = v26
            } else {
                delete(v3, 65535)
            }
        }
        var v27 uint = 42
        _ = fi23
        _ = v24
        _ = v27
    }
    v28 := GenS3{F0: (v4 * v4)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
