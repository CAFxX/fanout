package main

import (
    "fmt"
)

func genFn0(p0 int, p1 int64) int8 {
    var v5 uint = 554
    p1 = 42
    _ = p0
    _ = p1
    _ = v5
    return -128
}

func main() {
    var v1 float64 = ((3.14159 * (-2.25 / 1.0)) - 1e10)
    var v2 float64 = 1e10
    var v3 string = ("a" + "a")
    var v4 uint16 = 0
    var v6 bool = (!((v4 & v4) >= v4))
    v7 := []int{2147483647, 575, 2147483647, -2147483648}
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int16 = 322
        v1 -= ((0.0 * 1e10) * 3.14159)
        _ = fi8
        _ = v9
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v10 uint32 = 65535
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 int64 = 42
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        if ((^(v4)) <= (v4 << ((18446744073709551615) & 63))) {
            var v14 int = fi13
            if (!true) {
                if (v6 || v6) {
                    if (0 == 741) {
                        fmt.Println(v3)
                        v7 = append(v7, -9223372036854775808)
                        var v15 uint = 42
                        _ = v15
                    }
                }
                v10 = v10
            } else {
                v1 = float64(3.14159)
                var v16 uint16 = (v4 * v4)
                _ = len(v7)
                _ = v16
            }
            if v6 {
                fmt.Println(v6, v6)
                fmt.Println(v3, v3)
            }
            _ = v14
        } else {
            var v17 float64 = (v2 - -2.25)
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = v17
        }
        _ = fi13
    }
    v3 = ("" + ((("" + v3) + "\n\t") + (((v3 + v3) + "\n\t") + ((v3 + "foo bar") + "a"))))
    v6 = false
    var v18 uint16 = v4
    var v19 string = (v3 + v3)
    fmt.Println(v4, v3)
    switch "a" {
    case "hello":
        switch 2147483647 {
        case 9223372036854775807:
            var v20 int = -9223372036854775808
            _ = v20
        case 4294967295:
            if ((((float64(0.1) - v1) - 0.1) + float64(3.14159)) == v2) {
                var v21 uint = 519
                var v22 uint = (v21 ^ 0)
                _ = v21
                _ = v22
            }
        }
    case "foo bar":
        v19 = v3
    default:
        var v23 float32 = (float32(v2) / 1.0)
        _ = v23
    }
    v19 = v3
    var v24 bool = v6
    var v25 int8 = 122
    if (1 > 186) {
        if (!((1e10 + v1) == float64(1.5))) {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
