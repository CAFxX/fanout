package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint }
type GenS2 struct { F0 string; F1 uint }
func genFn0(p0 int8, p1 int64, p2 uint16) uint32 {
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 uint64 = 18446744073709551615
        v5 = (v5 + 242)
        _ = fi4
        _ = v5
    }
    p1 -= (p1 % ((9223372036854775807) | 1))
    _ = p0
    _ = p1
    _ = p2
    return 65535
}
func genFn1() uint32 {
    v6 := []int{0, 954, 561, 127}
    var v7 int = 127
    var v8 uint = 4294967295
    if (false && (true || true)) {
        v9 := GenS1{F0: v7, F1: 42, F2: 4294967295}
        _ = v9
    }
    var v10 uint = v8
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int = 0
    v2 := make(map[string]string)
    v3 := make(map[uint32]bool)
    v1 = v1
    switch 0 {
    case 128:
        v11 := []int{2147483647, (v1 << ((18446744073709551615) & 63)), (v1 * (v1 - (v1 + -874)))}
        _ = v11
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        switch ("foo bar" + "a") {
        case "":
            fmt.Println(v1, fi12)
        case "\n\t":
            _ = len(v2)
        default:
            if (!(1 > 1)) {
                fmt.Println(v1, v1)
                var v13 uint64 = 1
                _ = v13
            } else {
                var v14 float64 = ((-2.25 + 3.14159) + 3.14159)
                if (65535 >= 744) {
                    v1 = v1
                    v15 := GenS1{F0: fi12, F1: 2147483647, F2: 2}
                    _ = v15
                }
                _ = v14
            }
        }
        var v16 uint32 = 2147483648
        _ = fi12
        _ = v16
    }
    if true {
        for fi17 := 0; fi17 < 14; fi17++ {
            v18 := v3[4294967295]
            fmt.Println(fi17)
            fmt.Println(v18)
            _ = fi17
            _ = v18
        }
    } else {
        v19 := v2["\n\t"]
        v1 += (-(v1))
        _ = v19
    }
    fmt.Println(v1)
    v20 := v3[2]
    for fi21 := 0; fi21 < 8; fi21++ {
        var v22 float64 = 1e10
        _ = fi21
        _ = v22
    }
    v1 -= (v1 % ((v1) | 1))
    v20 = v20
    v23 := []int{150, (-(v1))}
    for fi24 := 0; fi24 < 18; fi24++ {
        v25 := GenS2{F0: ("\n\t" + ""), F1: 255}
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
}
