package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint8 }
type GenS2 struct { F0 bool }
func genFn0(p0 int8, p1 uint16) string {
    p0 = -35
    v5 := []int{127, -1, -65}
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            switch -11 {
            case 161:
                if len(v5) > 0 {
                    v5[0] = fi6
                }
            }
            var v8 uint = 0
            var v9 int32 = 42
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    var v10 int = 127
    _ = len(v5)
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    return "\n\t"
}
func genFn1(p0 string, p1 uint32, p2 int16) uint8 {
    v11 := GenS1{F0: -1, F1: 4294967295, F2: 61}
    v12 := GenS1{F0: 2147483647, F1: 776, F2: 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 2
}

func main() {
    v1 := []string{((("hello" + "foo bar") + "\n\t") + "hello"), "\n\t", ((("" + "") + "foo bar") + ""), (("hello" + (("a" + "foo bar") + "hello")) + ("foo bar" + ("\n\t" + "a")))}
    var v2 uint8 = 166
    var v3 bool = false
    var v4 uint32 = 722
    v13 := GenS1{F0: -2147483648, F1: 4294967295, F2: v2}
    for fi14 := 0; fi14 < 10; fi14++ {
        for fi15 := 0; fi15 < 8; fi15++ {
            switch "" {
            case "hello":
                fmt.Println(v3, fi15)
            case "\n\t":
                switch fi15 {
                case 192:
                    v4 -= 239
                default:
                    switch "hello" {
                    case "hello":
                        v2 = (v2 | v2)
                    default:
                        v16 := GenS2{F0: (!false)}
                        _ = v16
                    }
                }
            }
            _ = fi15
        }
        _ = fi14
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        v3 = (42 > 255)
        var v18 int64 = 42
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        var v20 uint32 = 1
        _ = fi19
        _ = v20
    }
    v21 := GenS1{F0: -124, F1: 1, F2: (^(v2))}
    if len(v1) > 0 {
        v1[0] = "a"
    }
    v3 = (128 <= -9223372036854775808)
    v22 := GenS1{F0: 661, F1: 2, F2: (v2 | (^(v2)))}
    v23 := GenS2{F0: (v2 == 0)}
    v1 = append(v1, "a")
    var v24 uint = 42
    v1 = append(v1, ("foo bar" + ""))
    v24 = v24
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
