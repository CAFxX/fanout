package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() uint8 {
    v7 := GenS1{F0: 2}
    var v8 int8 = 42
    v9 := GenS1{F0: 0}
    if (32767 >= -785) {
        if (42 == 0) {
            var v10 uint8 = 125
            switch 2147483647 {
            case -398:
                v11 := []int{32768, -431}
                _ = v11
            case 32768:
                v12 := []int{128}
                _ = v12
            case -208:
                for fi13 := 0; fi13 < 18; fi13++ {
                    var v14 int64 = 9223372036854775807
                    var v15 uint64 = 589
                    var v16 int8 = 127
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
            }
            switch ("\n\t" + "\n\t") {
            case "hello":
                var v17 float64 = ((0.0 / 1.5) * 3.14159)
                _ = v17
            case "a":
                var v18 string = "foo bar"
                _ = v18
            case "\n\t":
                v19 := GenS1{F0: 18446744073709551615}
                _ = v19
            }
            _ = v10
        } else {
            var v20 uint = 1
            for fi21 := 0; fi21 < 13; fi21++ {
                for fi22 := 0; fi22 < 20; fi22++ {
                    v23 := GenS1{F0: (v20 * v20)}
                    v20 -= v20
                    _ = fi22
                    _ = v23
                }
                _ = fi21
            }
            v8 -= v8
            _ = v20
        }
        var v24 uint = 510
        _ = v24
    } else {
        v25 := []int{-2147483648, 128, -472, 2147483647}
        fmt.Println(v8)
        var v26 int64 = 55
        _ = v25
        _ = v26
    }
    fmt.Println(v8)
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint16 = 2
    var v2 uint32 = 407
    v3 := []uint32{(v2 + v2), v2}
    var v4 int8 = -67
    var v5 float32 = -2.25
    var v6 int64 = 0
    var v27 uint = 2
    if false {
        var v28 int32 = 1
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
        var v29 uint = (v27 * v27)
        _ = v28
        _ = v29
    } else {
        var v30 float32 = 1e10
        v27 -= v27
        _ = v30
    }
    var v31 uint32 = (v2 * v2)
    if len(v3) > 0 {
        v3[0] = 707
    }
    var v32 uint32 = (^(v31))
    var v33 float64 = 0.1
    v3 = append(v3, 632)
    _ = len(v3)
    for fi34 := 0; fi34 < 18; fi34++ {
        var v35 float32 = v5
        for fi36 := 0; fi36 < 8; fi36++ {
            var v37 uint32 = v2
            _ = fi36
            _ = v37
        }
        _ = fi34
        _ = v35
    }
    v33 += 1e10
    v38 := GenS1{F0: v27}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
