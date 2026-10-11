package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
func genFn0() int64 {
    var v5 uint = 2
    v5 = v5
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 uint = 1
        if false {
            if (-344 > 1) {
                v7 = v5
                var v8 uint32 = 0
                var v9 int8 = 40
                _ = v8
                _ = v9
            }
            v5 = v5
        }
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 int8 = 42
            _ = fi10
            _ = v11
        }
        _ = fi6
        _ = v7
    }
    _ = v5
    return 42
}
func genFn1() string {
    var v12 uint32 = 199
    var v13 int = 66
    fmt.Println(v12)
    _ = v12
    _ = v13
    return "\n\t"
}

func main() {
    var v1 int = 2147483647
    var v2 int32 = 624
    var v3 int8 = -128
    v4 := []uint32{65535, 741, 682}
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    v3 = v3
    v3 = v3
    _ = len(v4)
    for fi14 := 0; fi14 < 7; fi14++ {
        v4 = append(v4, 2147483648)
        _ = fi14
    }
    if true {
        switch "hello" {
        case "a":
            var v15 uint16 = 255
            _ = v15
        case "hello":
            v2 = (-(v2))
        case "\n\t":
            v3 -= (v3 ^ (v3 ^ 0))
        default:
            v3 += (-(v3))
        }
        v3 -= 8
    } else {
        var v16 float64 = 3.14159
        switch "\n\t" {
        case "\n\t":
            var v17 uint8 = 8
            _ = v17
        case "hello":
            for fi18 := 0; fi18 < 20; fi18++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                fmt.Println(v3, v2)
                v3 = v3
                _ = fi18
            }
        default:
            for fi19 := 0; fi19 < 10; fi19++ {
                v20 := GenS1{F0: v3, F1: -32768}
                _ = fi19
                _ = v20
            }
        }
        _ = v16
    }
    if true {
        var v21 float64 = -2.25
        _ = v21
    } else {
        v4 = append(v4, 1)
        var v22 uint32 = 2147483648
        v2 += v2
        _ = v22
    }
    var v23 uint = 65535
    v3 += (v3 * v3)
    for fi24 := 0; fi24 < 4; fi24++ {
        var v25 uint16 = 42
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
        _ = fi24
        _ = v25
    }
    var v26 int8 = v3
    v23 += 248
    fmt.Println(v3, v2, v23)
    if true {
        v27 := GenS1{F0: v3, F1: 698}
        _ = v27
    } else {
        var v28 int64 = -1
        var v29 uint16 = 1
        var v30 int = (-(v1))
        _ = v28
        _ = v29
        _ = v30
    }
    var v31 string = "foo bar"
    _ = len(v4)
    v32 := GenS1{F0: v3, F1: 0}
    var v33 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
