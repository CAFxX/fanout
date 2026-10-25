package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 string; F1 uint8; F2 int8 }
func genFn0(p0 string) uint64 {
    switch 42 {
    case 887:
        var v6 uint16 = 427
        _ = v6
    case -2147483648:
        var v7 bool = (0 >= 855)
        _ = v7
    case 127:
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := []int{4294967295, fi8, (fi8 ^ (-(fi8)))}
            if len(v9) > 0 {
                v9[0] = (fi8 + 32768)
            }
            var v10 uint8 = 0
            _ = fi8
            _ = v10
            _ = v9
        }
    default:
        for fi11 := 0; fi11 < 10; fi11++ {
            for fi12 := 0; fi12 < 20; fi12++ {
                p0 = ""
                var v13 uint16 = 1
                _ = fi12
                _ = v13
            }
            if false {
                var v14 float32 = ((((float32(1.5) - float32(-2.25)) - float32(-2.25)) * -2.25) - (0.0 - (1e10 / (float32((3.14159 / 3.14159)) / 1.0))))
                _ = v14
            } else {
                p0 = (p0 + p0)
            }
            var v15 float32 = 1e10
            _ = fi11
            _ = v15
        }
    }
    fmt.Println(p0)
    if true {
        var v16 int64 = -9223372036854775808
        p0 = p0
        var v17 uint = 65535
        _ = v16
        _ = v17
    } else {
        switch 32768 {
        case 42:
            fmt.Println(p0)
        default:
            p0 = "foo bar"
        }
    }
    p0 = p0
    _ = p0
    return 42
}

func main() {
    var v1 bool = (false && (0 > 2))
    var v2 int64 = -1
    var v3 int8 = 127
    var v4 int = 9223372036854775807
    var v5 uint16 = 255
    if true {
        v4 = (v4 * v4)
        v4 = (v4 ^ (v4 << ((65535) & 63)))
    }
    v18 := GenS1{F0: v4, F1: 42}
    v3 = int8(v2)
    for fi19 := 0; fi19 < 6; fi19++ {
        for fi20 := 0; fi20 < 16; fi20++ {
            v3 += int8(v5)
            v1 = (18 == 1)
            v4 = (fi19 + (v4 % ((fi20) | 1)))
            _ = fi20
        }
        v21 := []int{(v4 % (((fi19 + fi19)) | 1)), (fi19 ^ 4294967295)}
        _ = fi19
        _ = v21
    }
    if (9223372036854775807 == (v2 ^ (v2 ^ (v2 ^ v2)))) {
        fmt.Println(v1, v2)
        var v22 float64 = 3.14159
        v5 += 2
        _ = v22
    } else {
        if (1 <= 255) {
            v1 = v1
            switch v4 {
            case 32768:
                v23 := []int{-358, v4, v4, -9223372036854775808}
                _ = v23
            }
            fmt.Println(v4)
        }
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        for fi25 := 0; fi25 < 1; fi25++ {
            v26 := []int{(v4 >> ((18446744073709551615) & 63))}
            v1 = true
            var v27 bool = true
            _ = fi25
            _ = v26
            _ = v27
        }
        v28 := GenS2{F0: "a", F1: 42, F2: v3}
        v4 += (fi24 - 127)
        _ = fi24
        _ = v28
    }
    v29 := []int{(v4 - v4), (v4 / (((v4 * v4)) | 1))}
    v30 := GenS2{F0: ("hello" + "a"), F1: 0, F2: (v3 << ((1) & 63))}
    v31 := GenS2{F0: "", F1: 1, F2: v3}
    v5 -= v5
    v5 += 262
    for fi32 := 0; fi32 < 12; fi32++ {
        if len(v29) > 0 {
            v29[0] = (fi32 | v4)
        }
        v33 := GenS1{F0: v4, F1: 0}
        v29 = append(v29, (fi32 | 4294967295))
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 17; fi34++ {
        v2 = v2
        fmt.Println(v5, fi34)
        var v35 int8 = 42
        _ = fi34
        _ = v35
    }
    v3 += (v3 / (((v3 ^ (v3 ^ v3))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
