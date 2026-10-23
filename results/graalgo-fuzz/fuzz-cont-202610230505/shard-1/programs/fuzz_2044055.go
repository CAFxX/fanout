package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int32; F1 int8; F2 bool }
type GenS3 struct { F0 int32; F1 int32 }
func genFn0(p0 bool) uint8 {
    var v5 float32 = 3.14159
    p0 = ((!p0) || (-32768 == 517))
    v6 := []int{42, 1, -2147483648}
    _ = p0
    _ = v5
    _ = v6
    return 128
}
func genFn1() uint32 {
    var v7 uint = 597
    var v8 string = (((("foo bar" + "foo bar") + "a") + ("" + "\n\t")) + (("\n\t" + ("\n\t" + ("a" + "\n\t"))) + "a"))
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 uint16 = uint16(v7)
        _ = fi9
        _ = v10
    }
    if (-1 > -128) {
        v11 := []int{}
        var v12 uint16 = 877
        _ = v11
        _ = v12
    }
    var v13 float64 = ((1e10 * 1.5) - ((3.14159 * ((-2.25 - 1.5) - 1.5)) - 0.1))
    _ = v13
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 float64 = 1.5
    var v2 uint64 = 0
    var v3 int8 = 42
    v4 := []string{"a", "hello", ((("hello" + "hello") + "") + "hello")}
    v2 += (v2 - v2)
    for fi14 := 0; fi14 < 9; fi14++ {
        v15 := GenS1{F0: 1, F1: -149}
        v16 := GenS1{F0: 681, F1: 513}
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 int32 = 298
    var v18 int32 = (^(v17))
    switch -9223372036854775808 {
    case 9223372036854775807:
        var v19 int16 = 0
        _ = v19
    case -2147483648:
        v20 := GenS2{F0: 0, F1: (v3 & v3), F2: ((-2.25 + -2.25) != ((((0.1 / 3.14159) * (3.14159 * 0.1)) * 0.1) * 1e10))}
        _ = v20
    case -240:
        var v21 int16 = -1
        _ = v21
    }
    v2 += (v2 | v2)
    switch 128 {
    case 32768:
        if true {
            for fi22 := 0; fi22 < 6; fi22++ {
                _ = len(v4)
                var v23 uint8 = 147
                v24 := GenS3{F0: (v17 | 1), F1: (v17 * v17)}
                _ = fi22
                _ = v23
                _ = v24
            }
        } else {
            if true {
                switch ("hello" + ((("a" + "\n\t") + ("foo bar" + "")) + "foo bar")) {
                case "a":
                    for fi25 := 0; fi25 < 12; fi25++ {
                        if false {
                            var v26 int = 42
                            _ = v26
                        } else {
                            v3 += v3
                            var v27 uint32 = 468
                            _ = v27
                        }
                        _ = fi25
                    }
                case "hello":
                    if (!true) {
                        fmt.Println(v2, v18)
                        v2 -= (v2 & 4294967295)
                        var v28 int32 = v18
                        _ = v28
                    }
                case "\n\t":
                    var v29 uint32 = 2147483647
                    _ = v29
                default:
                    var v30 bool = false
                    _ = v30
                }
            }
            v31 := GenS2{F0: (v18 % ((690) | 1)), F1: v3, F2: false}
            _ = v31
        }
    case 1:
        v4 = append(v4, "hello")
    case -962:
        for fi32 := 0; fi32 < 2; fi32++ {
            _ = len(v4)
            fmt.Println(v17, v1)
            _ = fi32
        }
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
