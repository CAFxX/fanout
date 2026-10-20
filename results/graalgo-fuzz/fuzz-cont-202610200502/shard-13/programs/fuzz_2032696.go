package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() int {
    var v6 int8 = 0
    v7 := []int{-2147483648, 42, 989}
    fmt.Println(v6)
    _ = v6
    _ = v7
    return 761
}
func genFn1(p0 string) int64 {
    fmt.Println(p0)
    switch (p0 + "\n\t") {
    case "hello":
        fmt.Println(p0)
    case "foo bar":
        if (42 < 65535) {
            if true {
                if false {
                    for fi8 := 0; fi8 < 17; fi8++ {
                        p0 = p0
                        for fi9 := 0; fi9 < 20; fi9++ {
                            v10 := []int{}
                            var v11 uint8 = 255
                            if len(v10) > 0 {
                                v10[0] = 42
                            }
                            _ = fi9
                            _ = v10
                            _ = v11
                        }
                        _ = fi8
                    }
                } else {
                    p0 = p0
                    fmt.Println(p0)
                }
            }
            var v12 uint = 2
            _ = v12
        } else {
            for fi13 := 0; fi13 < 14; fi13++ {
                fmt.Println(p0)
                _ = fi13
            }
            var v14 float32 = 0.0
            v15 := GenS1{F0: 0}
            _ = v14
            _ = v15
        }
    case "a":
        for fi16 := 0; fi16 < 13; fi16++ {
            v17 := GenS1{F0: -966}
            var v18 int16 = -355
            _ = fi16
            _ = v17
            _ = v18
        }
    default:
        if false {
            v19 := []int{}
            v20 := GenS1{F0: 1}
            _ = v19
            _ = v20
        }
    }
    _ = p0
    return -1
}

func main() {
    var v1 float32 = 1e10
    var v2 uint64 = 1
    var v3 uint = 0
    var v4 int16 = 0
    var v5 int16 = 32767
    v21 := []int{680, 82, 0}
    switch 32768 {
    case 9223372036854775807:
        var v22 int16 = (v5 & (v5 << ((255) & 63)))
        _ = v22
    case 32768:
        v21 = append(v21, 9223372036854775807)
    case 32767:
        fmt.Println(v4)
    }
    var v23 uint = 255
    v1 -= v1
    var v24 int = 32768
    var v25 int8 = -50
    for fi26 := 0; fi26 < 8; fi26++ {
        v24 -= fi26
        switch (("foo bar" + ("" + (("" + "a") + ("\n\t" + "foo bar")))) + "") {
        case "hello":
            var v27 int32 = 2147483647
            _ = v27
        case "\n\t":
            v21 = append(v21, (fi26 & (v24 * fi26)))
        case "a":
            for fi28 := 0; fi28 < 14; fi28++ {
                v29 := GenS1{F0: 759}
                var v30 float32 = 1.5
                _ = fi28
                _ = v29
                _ = v30
            }
        default:
            for fi31 := 0; fi31 < 4; fi31++ {
                var v32 uint = v23
                fmt.Println(v3, v5)
                _ = fi31
                _ = v32
            }
        }
        v3 -= v23
        _ = fi26
    }
    for fi33 := 0; fi33 < 17; fi33++ {
        v25 = (-(v25))
        fmt.Println(v5)
        var v34 int64 = 0
        _ = fi33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
