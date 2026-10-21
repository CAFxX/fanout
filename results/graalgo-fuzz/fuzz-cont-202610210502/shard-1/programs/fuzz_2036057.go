package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 string }
func genFn0(p0 uint32, p1 uint64, p2 int) uint8 {
    v4 := []int{-316, 387, 891}
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 int = (^(fi5))
        _ = fi5
        _ = v6
    }
    if len(v4) > 0 {
        v4[0] = p2
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        if len(v4) > 0 {
            v4[0] = p2
        }
        if true {
            v8 := GenS1{F0: (p1 & (p1 - 140))}
            _ = v8
        } else {
            fmt.Println(fi7, fi7, p1)
            var v9 uint8 = 42
            _ = v9
        }
        if len(v4) > 0 {
            v4[0] = (p2 * fi7)
        }
        _ = fi7
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        switch -2147483648 {
        case 735:
            v11 := GenS2{F0: 879, F1: "\n\t"}
            _ = v11
        case -1:
            switch "foo bar" {
            case "hello":
                for fi12 := 0; fi12 < 12; fi12++ {
                    if ((247 <= -169) && true) {
                        var v13 uint8 = 255
                        p1 = p1
                        _ = v13
                    } else {
                        fmt.Println(fi10, p1, fi10)
                    }
                    var v14 int = (fi12 + fi10)
                    var v15 uint32 = p0
                    _ = fi12
                    _ = v14
                    _ = v15
                }
            default:
                if (0 >= uint16(p0)) {
                    var v16 int = p2
                    _ = v16
                }
            }
        case 9223372036854775807:
            v17 := GenS3{F0: true, F1: 519}
            _ = v17
        default:
            for fi18 := 0; fi18 < 4; fi18++ {
                v19 := GenS2{F0: 255, F1: "\n\t"}
                _ = fi18
                _ = v19
            }
        }
        if true {
            p1 -= (p1 - (p1 ^ 0))
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2
}
func genFn1() uint32 {
    v20 := GenS1{F0: 42}
    var v21 int8 = 1
    _ = v20
    _ = v21
    return 2
}

func main() {
    var v1 int64 = 1
    var v2 uint16 = 42
    var v3 uint = 860
    switch "" {
    case "foo bar":
        v22 := []int{}
        _ = v22
    }
    v23 := []int{-1, -9223372036854775808, 253, -2147483648}
    v3 = (v3 % ((v3) | 1))
    fmt.Println(v1)
    if true {
        v2 = (v2 & 255)
        v24 := GenS3{F0: false, F1: 740}
        v1 -= (v1 + (v1 % (((v1 % (((v1 - int64(v2))) | 1))) | 1)))
        _ = v24
    } else {
        v1 = v1
        switch 42 {
        case 128:
            v25 := GenS2{F0: 358, F1: "\n\t"}
            _ = v25
        case 2147483647:
            if (0 < v2) {
                v1 = 1
                v1 += 1
            }
        case 1:
            v1 += (v1 ^ v1)
        }
    }
    v26 := GenS4{F0: ""}
    var v27 float32 = 1e10
    fmt.Println(v27, v3, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v3
}
