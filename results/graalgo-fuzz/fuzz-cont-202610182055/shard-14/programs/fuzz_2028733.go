package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint16, p1 uint32, p2 uint16) uint64 {
    p2 -= (p0 + p0)
    fmt.Println(p0, p2)
    switch -568 {
    case 128:
        p1 += (^(p1))
    case 127:
        if (18446744073709551615 <= 710) {
            for fi6 := 0; fi6 < 3; fi6++ {
                if true {
                    for fi7 := 0; fi7 < 7; fi7++ {
                        var v8 bool = ((fi7 == (fi7 * (fi7 - (fi6 << ((uint(p1)) & 63))))) || (1000 < 32767))
                        _ = fi7
                        _ = v8
                    }
                    fmt.Println(p2)
                }
                _ = fi6
            }
        } else {
            p2 = (^(p0))
        }
    case 564:
        p2 += (^(p2))
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 float64 = 1e10
    var v2 uint = 4294967295
    v3 := make(map[uint32]bool)
    v4 := make(map[string]string)
    var v5 uint32 = 65535
    v9 := GenS1{F0: "a", F1: (v2 << ((65535) & 63))}
    if false {
        var v10 float32 = 1.5
        v1 -= v1
        _ = v10
    }
    if true {
        v11 := []int{}
        _ = v11
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        if false {
            switch ("" + "\n\t") {
            case "\n\t":
                var v13 uint = (v2 ^ v2)
                _ = v13
            default:
                for fi14 := 0; fi14 < 8; fi14++ {
                    var v15 int16 = 1000
                    _ = fi14
                    _ = v15
                }
            }
            v16 := GenS1{F0: "", F1: (v2 << (((v2 & v2)) & 63))}
            _ = v16
        } else {
            fmt.Println(v2, fi12)
            v17 := GenS1{F0: "", F1: (v2 - uint(v5))}
            _ = v17
        }
        var v18 int32 = 0
        _ = fi12
        _ = v18
    }
    fmt.Println(v5, v5, v5)
    v2 -= (v2 >> ((v2) & 63))
    if false {
        for fi19 := 0; fi19 < 3; fi19++ {
            v20 := GenS2{F0: int64(fi19), F1: 127}
            fmt.Println(v2)
            _ = fi19
            _ = v20
        }
        v21 := []int{9223372036854775807, 881, 152, 1}
        fmt.Println(v2)
        _ = v21
    }
    fmt.Println(v5)
    v1 -= v1
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 int64 = -1
        _ = fi22
        _ = v23
    }
    var v24 uint = uint(v5)
    v5 += v5
    v25 := GenS3{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
