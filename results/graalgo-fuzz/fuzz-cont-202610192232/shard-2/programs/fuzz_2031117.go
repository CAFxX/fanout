package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int }
func genFn0(p0 int64, p1 string, p2 bool) bool {
    p1 = p1
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 int64 = p0
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 0, F1: 2, F2: 18446744073709551615}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return false
}

func main() {
    var v1 uint16 = 0
    var v2 bool = (2147483647 >= -975)
    v3 := make(map[int]uint32)
    var v4 float64 = -2.25
    var v5 bool = true
    var v6 uint64 = 42
    for fi10 := 0; fi10 < 20; fi10++ {
        v4 = (1.5 - (((v4 * (v4 * ((3.14159 + v4) + 0.0))) * 3.14159) * (v4 - -2.25)))
        v11 := []int{int(v1)}
        _ = fi10
        _ = v11
    }
    switch 128 {
    case 4294967295:
        switch "" {
        case "\n\t":
            for fi12 := 0; fi12 < 2; fi12++ {
                v4 -= v4
                var v13 float32 = 3.14159
                _ = fi12
                _ = v13
            }
        case "hello":
            v2 = false
        case "a":
            var v14 uint8 = 117
            _ = v14
        }
    case 32768:
        if (v2 && ((!v5) && v2)) {
            switch 128 {
            case 914:
                var v15 uint8 = 255
                _ = v15
            default:
                v1 += 1
            }
        } else {
            if true {
                switch 493 {
                case 2147483647:
                    fmt.Println(v6, v1)
                case -9223372036854775808:
                    v4 -= 0.0
                }
                fmt.Println(v2, v4)
            }
        }
    case -1:
        v5 = true
    default:
        v16 := GenS2{F0: 0, F1: 128}
        _ = v16
    }
    fmt.Println(v4, v4, v1)
    for fi17 := 0; fi17 < 17; fi17++ {
        for fi18 := 0; fi18 < 18; fi18++ {
            if true {
                v19 := v3[fi17]
                var v20 float32 = -2.25
                _ = v19
                _ = v20
            }
            if v2 {
                v21 := []int{(fi17 - (fi17 * 9223372036854775807)), 9223372036854775807}
                var v22 int64 = -9223372036854775808
                _ = v21
                _ = v22
            }
            var v23 uint = 981
            _ = fi18
            _ = v23
        }
        fmt.Println(v4, v2)
        _ = fi17
    }
    switch "a" {
    case "a":
        for fi24 := 0; fi24 < 16; fi24++ {
            fmt.Println(fi24, v6)
            _ = fi24
        }
    case "\n\t":
        v25 := GenS2{F0: 65535, F1: 2147483647}
        _ = v25
    case "hello":
        if v2 {
            v26 := []int{}
            if false {
                v2 = (!(25 == 0))
                fmt.Println(v1)
            }
            _ = v26
        } else {
            v1 = (v1 / (((v1 / ((v1) | 1))) | 1))
            v6 -= 4294967295
            fmt.Println(v2)
        }
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        v6 -= (^(v6))
        var v28 int64 = 479
        _ = fi27
        _ = v28
    }
    v29 := []int{4294967295, 236, -1, 4294967295}
    var v30 uint16 = (v1 ^ (v1 + (v1 + (v1 - (v1 << ((471) & 63))))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
