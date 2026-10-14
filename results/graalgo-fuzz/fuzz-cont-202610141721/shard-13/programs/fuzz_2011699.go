package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64 }
func genFn0() int64 {
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    if false {
        if (false && (0.0 >= 0.1)) {
            if (-1 >= -27) {
                var v9 uint8 = 255
                v9 += 192
                v9 -= 227
                _ = v9
            } else {
                v10 := GenS1{F0: 2}
                var v11 int8 = -25
                _ = v10
                _ = v11
            }
            var v12 int16 = -766
            fmt.Println(v12)
            _ = v12
        } else {
            var v13 int8 = 42
            fmt.Println(v13)
            _ = v13
        }
    }
    return -9223372036854775808
}

func main() {
    var v1 uint8 = 128
    var v2 int = 926
    var v3 int32 = 337
    var v4 uint32 = 2147483647
    var v5 uint16 = 0
    var v6 uint32 = 250
    v14 := []int{(v2 | -2147483648), 32767, 9223372036854775807}
    if false {
        for fi15 := 0; fi15 < 15; fi15++ {
            var v16 int16 = -1
            var v17 uint16 = (^(v5))
            _ = fi15
            _ = v16
            _ = v17
        }
        if ((v5 & 32767) >= 2) {
            var v18 float32 = 1e10
            _ = v18
        } else {
            if true {
                v1 += uint8(v4)
            } else {
                var v19 uint = uint(v1)
                _ = v19
            }
        }
        switch "\n\t" {
        case "foo bar":
            if true {
                fmt.Println(v5, v5)
            }
        case "hello":
            var v20 uint = 255
            _ = v20
        case "a":
            if len(v14) > 0 {
                v14[0] = 32768
            }
        default:
            v21 := GenS2{F0: 4294967295}
            _ = v21
        }
    } else {
        v5 = 1
    }
    var v22 int = v2
    fmt.Println(v5, v6, v6)
    if false {
        v2 = 42
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 int64 = -9223372036854775808
            v24 -= v24
            _ = fi23
            _ = v24
        }
        v6 -= v4
    } else {
        v25 := GenS1{F0: 0}
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = len(v14)
        _ = v25
    }
    for fi26 := 0; fi26 < 14; fi26++ {
        var v27 int16 = 32767
        var v28 bool = false
        _ = fi26
        _ = v27
        _ = v28
    }
    if ((!false) || false) {
        fmt.Println(v6, v6, v2)
        fmt.Println(v2, v3)
    }
    v1 = (^(v1))
    switch 9223372036854775807 {
    case 127:
        var v29 uint8 = (v1 ^ 1)
        _ = v29
    case -53:
        var v30 int32 = (v3 * (v3 & (v3 / ((1) | 1))))
        _ = v30
    }
    v4 -= v6
    if len(v14) > 0 {
        v14[0] = (v2 & v2)
    }
    v2 += (v22 / ((v22) | 1))
    fmt.Println(v3, v1, v2)
    fmt.Println(v22)
    var v31 string = (("\n\t" + "hello") + (("hello" + "") + "foo bar"))
    var v32 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
