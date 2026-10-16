package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int16; F2 uint8 }
type GenS3 struct { F0 string; F1 string; F2 int16 }
func genFn0() int {
    var v6 int = 4294967295
    var v7 int = (v6 | (v6 + (^(v6))))
    v6 += v6
    var v8 int16 = 1000
    var v9 int64 = 1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return int(v8)
}
func genFn1(p0 uint32, p1 uint8, p2 int8) int32 {
    v10 := []int{32767, 1, 1}
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := GenS1{F0: 2147483647}
        fmt.Println(p2, fi11, p1)
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        p2 = (p2 / ((0) | 1))
        _ = fi13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 int32 = 79
        _ = fi14
        _ = v15
    }
    var v16 int = 105
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v16
    return -2147483648
}

func main() {
    var v1 int8 = 42
    var v2 int8 = (v1 % (((v1 / ((v1) | 1))) | 1))
    v3 := []string{(((("a" + "a") + (("" + (("\n\t" + ("foo bar" + ((("a" + "foo bar") + (("a" + "hello") + "hello")) + ("a" + ("hello" + ("hello" + ("foo bar" + "a"))))))) + ("a" + ("foo bar" + ("\n\t" + ""))))) + ("hello" + ("" + "hello")))) + "hello") + "a"), "a", ("hello" + "foo bar")}
    var v4 int32 = 2
    v5 := make(map[int]uint32)
    if (!true) {
        fmt.Println(v2, v2)
        switch -2 {
        case 42:
            fmt.Println(v2, v1)
        }
    }
    if ("foo bar" <= "a") {
        if (!true) {
            v2 = (v1 * 100)
            fmt.Println(v2, v4, v1)
            fmt.Println(v4, v4, v1)
        } else {
            v17 := GenS2{F0: -9223372036854775808, F1: 1, F2: 2}
            _ = v17
        }
        v4 = (v4 * (v4 ^ v4))
        _ = len(v5)
    } else {
        var v18 int16 = 415
        _ = v18
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 int = -644
        _ = fi19
        _ = v20
    }
    v2 -= v1
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 string = "foo bar"
        if true {
            _ = len(v3)
        } else {
            v4 = (v4 % (((v4 / (((v4 % ((351) | 1))) | 1))) | 1))
            switch fi21 {
            case 0:
                var v23 int8 = 0
                _ = v23
            case 1:
                switch v22 {
                case "foo bar":
                    v24 := GenS3{F0: "hello", F1: "a", F2: 32767}
                    _ = v24
                default:
                    v2 = v1
                }
            case 32768:
                switch ("a" + v22) {
                case "hello":
                    if true {
                        v25 := v5[fi21]
                        _ = v25
                    }
                case "\n\t":
                    v26 := GenS2{F0: -9223372036854775808, F1: 959, F2: 201}
                    _ = v26
                default:
                    v22 = "\n\t"
                }
            default:
                fmt.Println(v1, v1, v2)
            }
            var v27 uint16 = 425
            _ = v27
        }
        v22 = ("\n\t" + v22)
        _ = fi21
        _ = v22
    }
    var v28 bool = (((!false) || (255 >= 165)) && (true && false))
    var v29 string = ""
    delete(v5, 2147483647)
    v28 = v28
    var v30 uint16 = 939
    var v31 float32 = (3.14159 / 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
