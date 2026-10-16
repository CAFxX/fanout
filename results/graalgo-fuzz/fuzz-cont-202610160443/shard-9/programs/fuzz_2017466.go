package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
func genFn0(p0 uint) uint16 {
    var v7 float64 = (1e10 - float64(3.14159))
    switch "\n\t" {
    case "foo bar":
        switch ("" + "hello") {
        case "foo bar":
            var v8 int8 = 118
            _ = v8
        case "hello":
            v7 = (v7 - v7)
        default:
            var v9 uint32 = 2147483647
            _ = v9
        }
    case "\n\t":
        v10 := []int{}
        _ = v10
    case "":
        var v11 float64 = v7
        _ = v11
    }
    _ = p0
    _ = v7
    return 2
}
func genFn1(p0 int64, p1 int32) int16 {
    fmt.Println(p0)
    v12 := []int{618, 0, 484}
    p1 = 1
    _ = p0
    _ = p1
    _ = v12
    return 32767
}

func main() {
    var v1 int32 = 566
    var v2 float32 = (1.5 + float32(-2.25))
    var v3 string = ("\n\t" + ("hello" + "foo bar"))
    var v4 uint = 18446744073709551615
    var v5 bool = (-32768 < 0)
    var v6 uint = v4
    v13 := GenS1{F0: 42, F1: -2147483648}
    switch -1 {
    case 127:
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 uint64 = 2
            fmt.Println(v2, v1)
            _ = fi14
            _ = v15
        }
    case -119:
        v2 += 3.14159
    case 32768:
        v16 := []int{2147483647}
        _ = v16
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        fmt.Println(v5, v4)
        fmt.Println(v1, v1, v6)
        _ = fi17
    }
    if v5 {
        var v18 string = v3
        for fi19 := 0; fi19 < 15; fi19++ {
            for fi20 := 0; fi20 < 18; fi20++ {
                var v21 float64 = (1.5 * 3.14159)
                var v22 float32 = float32(((1e10 * float64(3.14159)) * -2.25))
                if (491 <= 2) {
                    if ((v18 != ("foo bar" + "hello")) && v5) {
                        fmt.Println(v6, v22)
                        if v5 {
                            v23 := GenS1{F0: 2147483648, F1: fi19}
                            var v24 string = (((v3 + v3) + (("hello" + (("foo bar" + ("hello" + (v3 + v3))) + ((v3 + (v3 + (v3 + v18))) + "hello"))) + ("hello" + v18))) + ("foo bar" + v3))
                            _ = v23
                            _ = v24
                        }
                        fmt.Println(v22, v18, v2)
                    } else {
                        var v25 uint32 = 628
                        _ = v25
                    }
                    switch 1 {
                    case -1:
                        var v26 string = ((v18 + "\n\t") + (((("a" + ("" + "hello")) + ("a" + (v18 + "\n\t"))) + v3) + v3))
                        v6 -= (^(v6))
                        _ = v26
                    case 793:
                        v2 = v22
                    default:
                        v1 = v1
                    }
                    v27 := []int{fi19, (fi20 >> ((v4) & 63)), fi19}
                    _ = v27
                }
                var v28 int32 = v1
                _ = fi20
                _ = v21
                _ = v22
                _ = v28
            }
            v4 = v4
            _ = fi19
        }
        _ = v18
    } else {
        if (255 > 42) {
            v1 = (v1 << (((v4 << (((v4 + v6)) & 63))) & 63))
            v29 := GenS1{F0: 326, F1: 0}
            _ = v29
        }
        fmt.Println(v5)
    }
    v5 = true
    v30 := GenS1{F0: 1, F1: 32767}
    if false {
        var v31 string = ""
        if v5 {
            for fi32 := 0; fi32 < 10; fi32++ {
                v1 = (v1 * v1)
                var v33 uint32 = 82
                _ = fi32
                _ = v33
            }
        }
        if v5 {
            v2 = float32(-2.25)
            var v34 int32 = (v1 << (((v4 ^ 65535)) & 63))
            _ = v34
        }
        _ = v31
    } else {
        var v35 int16 = 371
        _ = v35
    }
    var v36 int = 9223372036854775807
    fmt.Println(v1)
    fmt.Println(v5, v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
