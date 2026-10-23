package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int32 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int8; F1 uint8 }
func genFn0() int16 {
    var v6 int = 32768
    switch -2147483648 {
    case -217:
        var v7 int8 = 24
        _ = v7
    default:
        switch 32768 {
        case 1:
            var v8 int8 = -12
            _ = v8
        case 127:
            var v9 int32 = 42
            _ = v9
        case 4294967295:
            var v10 string = ("hello" + (("hello" + "foo bar") + ((((("\n\t" + "a") + "\n\t") + "foo bar") + (((("foo bar" + "a") + ("hello" + "hello")) + "\n\t") + ("foo bar" + "hello"))) + "foo bar")))
            _ = v10
        }
    }
    var v11 uint64 = 2
    var v12 uint = 65535
    var v13 int16 = 1000
    _ = v11
    _ = v12
    _ = v13
    _ = v6
    return (v13 >> ((453) & 63))
}
func genFn1(p0 int32) uint64 {
    p0 = 0
    v14 := []int{42, 972, 4294967295, 0}
    for fi15 := 0; fi15 < 13; fi15++ {
        p0 -= 2147483647
        _ = fi15
    }
    v16 := GenS1{F0: true, F1: 2, F2: 391}
    _ = p0
    _ = v14
    _ = v16
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = -173
    var v3 int8 = 1
    var v4 int64 = 9223372036854775807
    var v5 int8 = 42
    v5 -= -1
    if false {
        v17 := GenS1{F0: false, F1: 28, F2: 223}
        _ = v17
    } else {
        var v18 float64 = (float64(((-2.25 - 3.14159) - -2.25)) - -2.25)
        if (!(!(-2147483648 > -496))) {
            v5 = (^(v5))
            fmt.Println(v3, v3)
        }
        var v19 uint64 = 4294967295
        _ = v18
        _ = v19
    }
    v20 := GenS1{F0: false, F1: 0, F2: 340}
    fmt.Println(v4, v3, v3)
    v21 := []int{0, 1, 1}
    v22 := GenS2{F0: 286, F1: v5}
    fmt.Println(v3, v3)
    for fi23 := 0; fi23 < 15; fi23++ {
        var v24 float64 = 3.14159
        if ((false && true) && (float64(3.14159) <= 1e10)) {
            delete(v1, -1)
            for fi25 := 0; fi25 < 1; fi25++ {
                fmt.Println(v24, v2, v4)
                var v26 int = -475
                v1[128] = ((100 == (v5 & v5)) || false)
                _ = fi25
                _ = v26
            }
            var v27 string = "\n\t"
            _ = v27
        } else {
            v28 := GenS1{F0: true, F1: 41, F2: 2147483647}
            var v29 int8 = v3
            _ = v28
            _ = v29
        }
        _ = fi23
        _ = v24
    }
    var v30 int64 = v4
    fmt.Println(v3, v30)
    fmt.Println(v4, v2, v30)
    for fi31 := 0; fi31 < 6; fi31++ {
        v32 := GenS3{F0: -26, F1: 127}
        switch "foo bar" {
        case "foo bar":
            v4 = (-(v30))
        default:
            v33 := v1[fi31]
            _ = v33
        }
        _ = len(v1)
        _ = fi31
        _ = v32
    }
    var v34 bool = (((v5 - v5) < v5) && ((1 > 407) && true))
    v35 := v1[-9223372036854775808]
    if v35 {
        var v36 int32 = 42
        _ = v36
    }
    var v37 int = 42
    switch -509 {
    case 42:
        var v38 uint8 = 127
        _ = v38
    case 2147483647:
        for fi39 := 0; fi39 < 5; fi39++ {
            fmt.Println(v37, v2, v4)
            v2 = genFn0()
            fmt.Println(v5, v35)
            _ = fi39
        }
    default:
        var v40 int32 = -1
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v37
    _ = v4
    _ = v5
}
