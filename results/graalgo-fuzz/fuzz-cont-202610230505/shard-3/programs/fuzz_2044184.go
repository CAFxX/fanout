package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint32; F1 uint8; F2 int64 }
func genFn0(p0 uint64) uint16 {
    fmt.Println(p0)
    v5 := []int{128, -2147483648, 9223372036854775807, 4294967295}
    p0 -= 82
    _ = p0
    _ = v5
    return 65535
}
func genFn1() string {
    v6 := []int{2147483647, 9223372036854775807, -354, 32768}
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 bool = false
        v9 := GenS1{F0: 255, F1: ""}
        v10 := GenS1{F0: 1, F1: "\n\t"}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 uint32 = 0
    v11 = (v11 | (v11 & 42))
    var v12 int64 = 130
    _ = v11
    _ = v12
    _ = v6
    return "hello"
}

func main() {
    var v1 int32 = 35
    var v2 uint32 = 954
    var v3 uint32 = 0
    var v4 int = -1
    v13 := []int{128}
    v4 = v4
    v2 -= 42
    for fi14 := 0; fi14 < 1; fi14++ {
        _ = len(v13)
        _ = fi14
    }
    var v15 bool = (false || (42 != (v1 >> ((0) & 63))))
    switch "hello" {
    case "foo bar":
        v4 = -1
    case "a":
        if v15 {
            switch v4 {
            case -2147483648:
                if ((("" + "hello") + "") == "hello") {
                    v16 := GenS1{F0: 128, F1: (("\n\t" + (("\n\t" + "a") + "a")) + "")}
                    v4 -= (v4 >> ((916) & 63))
                    var v17 float64 = float64((-2.25 * (3.14159 + float32(((-2.25 * 1.5) - (-2.25 * (1e10 * 1.5)))))))
                    _ = v16
                    _ = v17
                }
            case -1:
                v18 := GenS1{F0: 128, F1: "a"}
                _ = v18
            }
            switch 9223372036854775807 {
            case -2147483648:
                var v19 int16 = 1
                _ = v19
            }
        }
    case "hello":
        var v20 int64 = 760
        _ = v20
    default:
        switch v4 {
        case -831:
            for fi21 := 0; fi21 < 17; fi21++ {
                for fi22 := 0; fi22 < 7; fi22++ {
                    v23 := GenS2{F0: 2147483648, F1: 42, F2: 9223372036854775807}
                    var v24 bool = (18446744073709551615 == 0)
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                _ = fi21
            }
        default:
            if true {
                if v15 {
                    switch "a" {
                    case "hello":
                        v25 := GenS1{F0: 2, F1: "hello"}
                        _ = v25
                    case "foo bar":
                        fmt.Println(v3, v15)
                    }
                }
            } else {
                for fi26 := 0; fi26 < 4; fi26++ {
                    fmt.Println(v3, v2)
                    var v27 uint = 0
                    _ = fi26
                    _ = v27
                }
                var v28 uint16 = 65535
                v2 = v3
                _ = v28
            }
        }
    }
    var v29 uint32 = v3
    v13 = append(v13, (v4 >> ((4294967295) & 63)))
    switch -117 {
    case 4294967295:
        for fi30 := 0; fi30 < 7; fi30++ {
            var v31 uint16 = 586
            _ = fi30
            _ = v31
        }
    case 128:
        var v32 int32 = 27
        _ = v32
    }
    fmt.Println(v15)
    v1 += 1
    v33 := GenS1{F0: 2, F1: ("hello" + (("foo bar" + "foo bar") + ("foo bar" + "\n\t")))}
    v34 := GenS1{F0: 128, F1: ""}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
