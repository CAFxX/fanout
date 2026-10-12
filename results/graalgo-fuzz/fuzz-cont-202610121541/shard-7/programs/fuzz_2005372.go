package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint8, p1 uint16) int32 {
    p0 -= (^(p0))
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := GenS1{F0: p0, F1: (^(p1)), F2: 2147483647}
        p0 -= p0
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    return 242
}

func main() {
    var v1 float32 = 0.0
    v2 := []string{"foo bar", "", ("\n\t" + ("\n\t" + ""))}
    var v3 uint32 = 1
    var v4 int = -2147483648
    var v5 uint8 = 198
    fmt.Println(v3)
    if (42 == -1) {
        fmt.Println(v1, v1, v5)
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        if (v1 < float32(0.0)) {
            v3 += v3
        } else {
            v1 = (3.14159 - (v1 - (v1 / (v1 - 0.1))))
            v8 := GenS1{F0: (v5 * v5), F1: 32767, F2: 65535}
            v9 := GenS1{F0: v5, F1: 255, F2: 2}
            _ = v8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        for fi11 := 0; fi11 < 6; fi11++ {
            if (!(float32(0.1) <= (v1 - 3.14159))) {
                v2 = append(v2, "")
                v2 = append(v2, "\n\t")
            }
            if true {
                v12 := GenS2{F0: -9223372036854775808}
                _ = v12
            } else {
                v13 := GenS2{F0: -584}
                v3 -= v3
                var v14 int64 = 1
                _ = v13
                _ = v14
            }
            _ = fi11
        }
        _ = fi10
    }
    _ = len(v2)
    switch "hello" {
    case "\n\t":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "a":
        v5 += v5
    case "foo bar":
        v5 = v5
    default:
        v4 = 815
    }
    fmt.Println(v4)
    if (-862 == -1) {
        switch 128 {
        case -2147483648:
            var v15 uint8 = (v5 ^ (v5 | (^(v5))))
            _ = v15
        case 561:
            var v16 string = "a"
            _ = v16
        default:
            v17 := GenS2{F0: -479}
            _ = v17
        }
        fmt.Println(v4, v3)
        v3 += (v3 & (^(v3)))
    }
    var v18 int8 = -128
    if len(v2) > 0 {
        v2[0] = (("" + "\n\t") + "\n\t")
    }
    v19 := GenS2{F0: int64(v3)}
    var v20 uint64 = 735
    var v21 int8 = (v18 | (v18 >> ((12) & 63)))
    v1 = float32(1.5)
    v21 -= v21
    v1 += v1
    for fi22 := 0; fi22 < 10; fi22++ {
        fmt.Println(v5, v20)
        _ = fi22
    }
    var v23 int = -9223372036854775808
    v1 = ((v1 / 1.5) / (float32((float64((((0.1 / 1.0) - ((3.14159 / 0.1) + 1.5)) - v1)) + ((-2.25 + 1.5) - 0.0))) * (float32(0.0) - v1)))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
