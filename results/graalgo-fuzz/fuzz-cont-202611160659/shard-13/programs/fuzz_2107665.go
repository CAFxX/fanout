package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
func genFn0(p0 int64, p1 uint16, p2 int) uint16 {
    v5 := []int{(p2 / ((-1) | 1))}
    switch "\n\t" {
    case "hello":
        fmt.Println(p1, p2)
    case "":
        if true {
            for fi6 := 0; fi6 < 11; fi6++ {
                fmt.Println(fi6, p1)
                switch "foo bar" {
                case "a":
                    for fi7 := 0; fi7 < 11; fi7++ {
                        p2 -= fi7
                        p2 = (fi7 & 127)
                        p1 += p1
                        _ = fi7
                    }
                }
                _ = fi6
            }
            p2 = 0
            var v8 uint64 = 2
            _ = v8
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    case "foo bar":
        var v9 uint = 18446744073709551615
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p1
}
func genFn1() int64 {
    switch ("a" + "a") {
    case "":
        var v10 int64 = -1
        _ = v10
    }
    switch 32767 {
    case 9223372036854775807:
        if false {
            switch -9223372036854775808 {
            case -862:
                var v11 float32 = 1.5
                _ = v11
            default:
                v12 := []int{}
                _ = v12
            }
            var v13 int64 = -361
            _ = v13
        } else {
            for fi14 := 0; fi14 < 15; fi14++ {
                for fi15 := 0; fi15 < 5; fi15++ {
                    v16 := []int{fi14, 2147483647, (fi14 + (fi15 ^ -2147483648))}
                    _ = fi15
                    _ = v16
                }
                var v17 float64 = 0.1
                _ = fi14
                _ = v17
            }
            v18 := GenS1{F0: ((!((("" + "a") + "foo bar") == (("" + "hello") + "foo bar"))) && (223 < -1)), F1: 0}
            _ = v18
        }
    default:
        v19 := []int{32767, 737}
        _ = v19
    }
    return 349
}

func main() {
    var v1 bool = (3.14159 >= float32(0.0))
    var v2 int = -9223372036854775808
    v3 := []int{(v2 | (-(v2))), -441, v2}
    v4 := make(map[string]string)
    v2 = (v2 + (v2 * v2))
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 float32 = (-2.25 + -2.25)
        _ = fi20
        _ = v21
    }
    fmt.Println(v1, v1)
    var v22 uint8 = 127
    var v23 uint32 = 954
    v2 += (^(v2))
    fmt.Println(v2, v23)
    var v24 string = ""
    if (v23 < v23) {
        for fi25 := 0; fi25 < 6; fi25++ {
            v4[((v24 + "hello") + v24)] = (v24 + (("foo bar" + v24) + ((("hello" + "") + (((v24 + (v24 + (v24 + "hello"))) + ("\n\t" + ("foo bar" + ("\n\t" + v24)))) + "")) + v24)))
            v2 = v2
            _ = fi25
        }
    } else {
        if ((1e10 / 1.5) != float32((float64(1e10) + 3.14159))) {
            v23 -= v23
            var v26 uint = 0
            _ = v26
        }
    }
    var v27 string = ""
    v28 := GenS1{F0: v1, F1: 100}
    var v29 bool = (v23 != 343)
    v2 -= -2147483648
    v30 := GenS1{F0: (v2 < v2), F1: -1}
    var v31 uint8 = (v22 / ((1) | 1))
    var v32 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
