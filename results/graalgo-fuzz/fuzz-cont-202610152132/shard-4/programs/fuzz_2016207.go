package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int8, p1 int8) uint64 {
    var v4 int = 9223372036854775807
    var v5 uint16 = 32767
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}

func main() {
    v1 := []uint32{}
    var v2 uint = 42
    var v3 float32 = 1e10
    v1 = append(v1, 4294967295)
    if false {
        v2 += v2
    }
    if false {
        v3 = 0.0
        fmt.Println(v2, v3)
        switch ("foo bar" + "") {
        case "foo bar":
            fmt.Println(v2, v2)
        default:
            fmt.Println(v3)
        }
    }
    var v6 uint32 = 2147483648
    switch "\n\t" {
    case "hello":
        if true {
            var v7 int16 = 1000
            _ = v7
        } else {
            v3 -= 1e10
            fmt.Println(v3, v3)
        }
    case "":
        var v8 int8 = -105
        _ = v8
    }
    if false {
        v3 -= 0.1
    } else {
        var v9 uint16 = 0
        var v10 int16 = -1
        _ = v10
        _ = v9
    }
    var v11 float64 = (float64((((3.14159 - (v3 / v3)) * ((v3 * 1.5) / v3)) / v3)) + ((((0.0 / float64((v3 - 0.0))) - 0.1) / (1e10 / (1.5 - 0.1))) / 3.14159))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch -453 {
    case -2147483648:
        fmt.Println(v2)
    case 42:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        v12 := GenS1{F0: 23}
        _ = v12
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := GenS1{F0: 127}
        _ = fi13
        _ = v14
    }
    if (true || true) {
        if (true || (((v11 * 1.5) * (1.5 * float64(((1e10 / (v3 - 1.5)) - v3)))) == v11)) {
            for fi15 := 0; fi15 < 4; fi15++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi15
            }
        }
        fmt.Println(v2, v2, v11)
    } else {
        var v16 int64 = -1
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 uint16 = 255
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    v2 = (v2 ^ (^(v2)))
    if false {
        if (int32(v2) > -2147483648) {
            v19 := GenS1{F0: 127}
            _ = v19
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v1 = append(v1, (^(v6)))
    fmt.Println(v6)
    switch "foo bar" {
    case "\n\t":
        v6 = v6
    case "a":
        if true {
            v20 := GenS1{F0: 230}
            _ = v20
        }
    default:
        if ((^(v2)) <= (v2 % (((v2 | (v2 + 4294967295))) | 1))) {
            switch "foo bar" {
            case "hello":
                v1 = append(v1, (v6 - v6))
            case "\n\t":
                for fi21 := 0; fi21 < 1; fi21++ {
                    var v22 uint64 = 802
                    var v23 int = fi21
                    var v24 string = ("" + ("hello" + "a"))
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                }
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v11 -= (v11 / float64((1e10 * 0.1)))
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v6
}
