package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint8 }
func genFn0(p0 bool, p1 int16, p2 int16) bool {
    if (p0 && p0) {
        var v7 float32 = 1e10
        v8 := []int{127}
        _ = v7
        _ = v8
    } else {
        var v9 int32 = 1
        _ = v9
    }
    v10 := []int{42}
    fmt.Println(p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return p0
}
func genFn1() uint32 {
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := []int{fi11, (fi11 - 554), (fi11 - fi11), (fi11 + (-(fi11)))}
        _ = fi11
        _ = v12
    }
    var v13 uint8 = 2
    _ = v13
    return 65535
}

func main() {
    var v1 int64 = 0
    var v2 int8 = 42
    var v3 uint32 = 33
    var v4 uint32 = (v3 % (((^(v3))) | 1))
    var v5 float64 = ((((0.1 - 1.5) * float64(1e10)) - 3.14159) + (0.0 * (-2.25 * (3.14159 / ((((1e10 - (1e10 + 3.14159)) * float64(float32((0.1 * 1.5)))) * (0.0 + 1e10)) - 3.14159)))))
    v6 := make(map[string]string)
    var v14 int32 = 1
    var v15 float32 = (-2.25 - -2.25)
    if (false || false) {
        if (!false) {
            fmt.Println(v4)
        } else {
            if (!true) {
                fmt.Println(v3, v5)
                v6[("" + "hello")] = "a"
                delete(v6, "\n\t")
            }
            fmt.Println(v2, v4, v1)
            var v16 float32 = v15
            _ = v16
        }
        switch "foo bar" {
        case "\n\t":
            v17 := []int{-2147483648, 193}
            _ = v17
        default:
            v14 = v14
        }
        var v18 string = "foo bar"
        _ = v18
    } else {
        for fi19 := 0; fi19 < 15; fi19++ {
            delete(v6, ("a" + ("" + ("foo bar" + "foo bar"))))
            _ = fi19
        }
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        var v21 uint64 = 514
        switch (fi20 * (fi20 & fi20)) {
        case 2147483647:
            var v22 float32 = 1.5
            _ = v22
        default:
            v2 -= v2
        }
        _ = fi20
        _ = v21
    }
    if false {
        v23 := GenS1{F0: v14, F1: (v14 % ((0) | 1)), F2: 189}
        _ = v23
    }
    fmt.Println(v14, v1, v2)
    v5 -= v5
    if false {
        for fi24 := 0; fi24 < 15; fi24++ {
            v14 -= (^(v14))
            _ = fi24
        }
        v5 = (v5 / (-2.25 + v5))
        var v25 int8 = -1
        _ = v25
    } else {
        var v26 int64 = v1
        _ = v26
    }
    var v27 int8 = -1
    v28 := []int{-9223372036854775808, 32767, 821, 0}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
