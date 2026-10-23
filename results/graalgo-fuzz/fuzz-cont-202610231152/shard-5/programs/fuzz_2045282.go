package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int }
func genFn0(p0 uint8) uint16 {
    p0 = 236
    v5 := []int{}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, 9223372036854775807)
    switch 32767 {
    case 0:
        if len(v5) > 0 {
            v5[0] = 32768
        }
    case 9223372036854775807:
        if (0 < 238) {
            var v6 uint8 = 128
            var v7 int32 = -123
            _ = v6
            _ = v7
        }
    default:
        if true {
            if len(v5) > 0 {
                v5[0] = 1
            }
            var v8 int32 = 0
            _ = v8
        } else {
            if (1 != 576) {
                var v9 uint16 = 255
                _ = v9
            }
        }
    }
    _ = p0
    _ = v5
    return 1
}
func genFn1(p0 string, p1 bool, p2 int64) int8 {
    var v10 int64 = (p2 + -9223372036854775808)
    var v11 bool = p1
    v10 += v10
    p2 = (v10 / ((1) | 1))
    if true {
        if (1e10 > -2.25) {
            v11 = (!v11)
        } else {
            fmt.Println(v11, p2, v11)
            if p1 {
                v12 := GenS1{F0: 5, F1: 0}
                _ = v12
            }
        }
        v10 -= v10
        if (2 >= 7) {
            var v13 int = 0
            var v14 int = (v13 + v13)
            _ = v13
            _ = v14
        }
    } else {
        v15 := []int{-9223372036854775808}
        var v16 int8 = -92
        switch (p0 + "foo bar") {
        case "":
            var v17 int8 = v16
            _ = v17
        case "a":
            var v18 uint = 678
            _ = v18
        default:
            for fi19 := 0; fi19 < 20; fi19++ {
                _ = len(v15)
                var v20 int = (-(fi19))
                _ = fi19
                _ = v20
            }
        }
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 0
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int8 = -127
    v3 := make(map[uint32]string)
    var v4 float64 = 1e10
    fmt.Println(v1)
    if (-2147483648 <= -2147483648) {
        if false {
            v1 -= -705
            switch 4294967295 {
            case 32768:
                var v21 uint16 = 960
                _ = v21
            case 2147483647:
                v3[65535] = "a"
            }
        }
        var v22 int8 = genFn1("", (true || false), (v1 >> ((1) & 63)))
        var v23 uint = 687
        _ = v22
        _ = v23
    } else {
        v1 = (v1 >> ((18446744073709551615) & 63))
        var v24 string = ("\n\t" + ("hello" + ""))
        var v25 int32 = 42
        _ = v24
        _ = v25
    }
    v4 += (v4 / v4)
    if (!((true || true) && (-117 >= v2))) {
        if false {
            fmt.Println(v1)
            v4 += v4
        }
    } else {
        v26 := GenS1{F0: 2147483647, F1: 2147483647}
        v27 := GenS1{F0: -2147483648, F1: 2147483648}
        if false {
            switch "hello" {
            case "a":
                fmt.Println(v1, v2, v4)
            }
            v28 := GenS2{F0: 42, F1: (^(v2))}
            v29 := GenS1{F0: 42, F1: 2147483647}
            _ = v28
            _ = v29
        }
        _ = v26
        _ = v27
    }
    var v30 string = "a"
    switch (v30 + "a") {
    case "hello":
        var v31 int16 = 1
        _ = v31
    default:
        for fi32 := 0; fi32 < 13; fi32++ {
            v4 = ((v4 * float64((3.14159 / 1e10))) / v4)
            _ = fi32
        }
    }
    for fi33 := 0; fi33 < 11; fi33++ {
        v4 = (float64(3.14159) - float64((0.0 / (0.0 - 1.5))))
        v34 := GenS2{F0: 934, F1: v2}
        _ = fi33
        _ = v34
    }
    v35 := GenS3{F0: -108}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v4
}
