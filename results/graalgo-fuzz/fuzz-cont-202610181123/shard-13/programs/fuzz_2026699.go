package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 string; F1 uint32 }
func genFn0(p0 uint64, p1 int16, p2 int) uint {
    p0 = p0
    fmt.Println(p1)
    var v5 float32 = (float32(3.14159) / 3.14159)
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 4294967295
}

func main() {
    var v1 uint64 = 805
    var v2 bool = (32767 != 42)
    v3 := []string{}
    var v4 int16 = 1
    if (true && true) {
        fmt.Println(v1)
        for fi6 := 0; fi6 < 5; fi6++ {
            fmt.Println(v4, v1, fi6)
            v7 := GenS1{F0: ((v2 || (false || v2)) && v2), F1: 0}
            _ = fi6
            _ = v7
        }
        var v8 int16 = 291
        _ = v8
    }
    var v9 int32 = 42
    switch 127 {
    case -183:
        for fi10 := 0; fi10 < 15; fi10++ {
            fmt.Println(v2)
            v2 = v2
            var v11 int32 = -1
            _ = fi10
            _ = v11
        }
    case -505:
        v12 := GenS2{F0: "\n\t", F1: 2147483648}
        _ = v12
    case -2147483648:
        if v2 {
            v13 := GenS2{F0: ("" + ((((("foo bar" + ("hello" + (("foo bar" + "foo bar") + "foo bar"))) + "a") + "\n\t") + "foo bar") + "")), F1: 2147483648}
            v9 -= v9
            _ = v13
        }
    }
    var v14 int16 = (v4 - int16(v1))
    v3 = append(v3, (("a" + "hello") + "foo bar"))
    var v15 int8 = 42
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = ((((0 > -9223372036854775808) || v2) && (("" > "hello") && v2)) || (1 != 715))
    v16 := GenS1{F0: (604 < -814), F1: (-(v15))}
    var v17 uint32 = 216
    var v18 int = 42
    if (1 == 18446744073709551615) {
        v19 := GenS1{F0: (v9 < -1), F1: v15}
        _ = v19
    }
    v20 := GenS2{F0: "a", F1: v17}
    for fi21 := 0; fi21 < 19; fi21++ {
        for fi22 := 0; fi22 < 16; fi22++ {
            var v23 float32 = ((1e10 / float32((0.1 * float64(float32(-2.25))))) + 3.14159)
            if (v2 && true) {
                fmt.Println(v2, fi22, v1)
            }
            v24 := GenS2{F0: ("a" + ("a" + "hello")), F1: (v17 ^ 420)}
            if ((v17 + v17) < v17) {
                var v25 int8 = 127
                _ = v25
            }
            if (fi22 > (fi21 ^ (v18 * 9223372036854775807))) {
                fmt.Println(fi22, v14)
            } else {
                var v26 uint16 = 130
                v27 := GenS1{F0: v2, F1: v15}
                _ = v26
                _ = v27
            }
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v9
}
