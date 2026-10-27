package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int32 }
type GenS2 struct { F0 string; F1 int64 }
func genFn0() uint32 {
    v5 := GenS1{F0: 1, F1: ("hello" + (("foo bar" + ("foo bar" + "hello")) + "hello")), F2: 1}
    var v6 int16 = -475
    for fi7 := 0; fi7 < 5; fi7++ {
        switch (fi7 >> ((4294967295) & 63)) {
        case -2147483648:
            var v8 string = "a"
            _ = v8
        }
        _ = fi7
    }
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 float64 = 3.14159
    var v2 string = "\n\t"
    var v3 bool = (-576 > 882)
    var v4 uint32 = 2
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 float64 = v1
        v11 := GenS2{F0: "", F1: -9223372036854775808}
        _ = fi9
        _ = v10
        _ = v11
    }
    v1 = -2.25
    v1 -= ((1.5 - 0.1) - 1.5)
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 int32 = 2147483647
        v4 = (v4 << ((18446744073709551615) & 63))
        _ = fi12
        _ = v13
    }
    var v14 string = "hello"
    if v3 {
        v4 -= (^(v4))
    } else {
        fmt.Println(v14)
        var v15 int64 = 591
        _ = v15
    }
    if v3 {
        switch v14 {
        case "a":
            for fi16 := 0; fi16 < 11; fi16++ {
                v14 = (v14 + "a")
                v3 = ((!(604 > 544)) || ((!true) || false))
                _ = fi16
            }
        case "\n\t":
            v17 := GenS1{F0: 240, F1: (v2 + (v14 + "")), F2: 2147483647}
            _ = v17
        case "foo bar":
            v18 := []int{-976, -2147483648}
            _ = v18
        default:
            v14 = "hello"
        }
        if (4294967295 > 0) {
            fmt.Println(v14, v2)
            var v19 int8 = 100
            v20 := []int{}
            _ = v19
            _ = v20
        }
        v21 := []int{4294967295, 127, 42}
        _ = v21
    }
    var v22 int8 = 127
    v23 := []int{0, -2147483648}
    if (42 != uint(v4)) {
        v1 -= (0.1 * (float64(((float32(v1) / ((0.0 - (-2.25 + (0.0 + ((-2.25 * ((1.5 - 0.1) / -2.25)) - (3.14159 - (1e10 - (0.0 / -2.25))))))) / -2.25)) + 0.0)) + (v1 * 3.14159)))
        for fi24 := 0; fi24 < 5; fi24++ {
            v2 = "foo bar"
            _ = fi24
        }
        fmt.Println(v3, v2)
    } else {
        for fi25 := 0; fi25 < 4; fi25++ {
            var v26 uint32 = v4
            if v3 {
                var v27 int8 = v22
                _ = v27
            } else {
                var v28 uint32 = (^(v26))
                _ = v28
            }
            _ = fi25
            _ = v26
        }
    }
    var v29 uint16 = 186
    v3 = v3
    var v30 uint32 = (v4 >> ((4294967295) & 63))
    fmt.Println(v29, v2, v29)
    v29 += v29
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
