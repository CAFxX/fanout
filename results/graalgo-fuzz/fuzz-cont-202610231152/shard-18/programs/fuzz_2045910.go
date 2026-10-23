package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int; F1 uint16 }
func genFn0(p0 int8) string {
    v6 := []int{32768, 0, 42, -2147483648}
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 int64 = -1
        fmt.Println(p0)
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v6
    return ("hello" + "foo bar")
}
func genFn1() uint {
    for fi9 := 0; fi9 < 3; fi9++ {
        fmt.Println(fi9)
        v10 := []int{fi9}
        switch "a" {
        case "a":
            if false {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                if true {
                    var v11 int8 = -31
                    var v12 string = "foo bar"
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    _ = v11
                    _ = v12
                }
                if len(v10) > 0 {
                    v10[0] = fi9
                }
            } else {
                var v13 string = "a"
                _ = v13
            }
        }
        _ = fi9
        _ = v10
    }
    switch 4294967295 {
    case 127:
        switch 127 {
        case 71:
            v14 := GenS1{F0: 172, F1: -1}
            _ = v14
        case 42:
            var v15 int32 = 0
            _ = v15
        case 569:
            if (-761 < -9223372036854775808) {
                var v16 int16 = -100
                _ = v16
            }
        }
    case 128:
        var v17 uint = 0
        _ = v17
    }
    var v18 uint16 = 65535
    v18 += (^(v18))
    v19 := GenS1{F0: -633, F1: 108}
    _ = v18
    _ = v19
    return 18446744073709551615
}

func main() {
    v1 := []string{("\n\t" + "a")}
    var v2 uint32 = 65535
    v3 := make(map[int]uint32)
    var v4 float32 = (-2.25 - (float32(((-2.25 - (((0.0 / -2.25) / (float64((3.14159 - ((float32(1e10) * (((0.1 - 0.0) / 0.1) / -2.25)) / (float32(0.1) + -2.25)))) + (float64(0.0) / 0.1))) / float64(float32(3.14159)))) + 1.5)) + -2.25))
    v5 := make(map[string]int)
    v20 := GenS2{F0: 42}
    if true {
        v21 := GenS3{F0: (v2 % ((v2) | 1)), F1: v2}
        v2 = 0
        var v22 uint = 44
        _ = v21
        _ = v22
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v4 -= ((v4 * 0.1) + -2.25)
    }
    v1 = append(v1, "")
    fmt.Println(v4, v4)
    var v23 string = ""
    switch ("foo bar" + v23) {
    case "":
        fmt.Println(v4)
    case "hello":
        if true {
            if (((v2 / ((321) | 1)) <= v2) && true) {
                var v24 bool = false
                var v25 uint8 = 0
                _ = v24
                _ = v25
            } else {
                var v26 uint16 = 32767
                _ = v26
            }
        } else {
            for fi27 := 0; fi27 < 18; fi27++ {
                switch v23 {
                case "a":
                    for fi28 := 0; fi28 < 8; fi28++ {
                        v2 = v2
                        var v29 uint8 = 42
                        v30 := GenS3{F0: 2147483647, F1: v2}
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                case "\n\t":
                    var v31 int = 545
                    _ = v31
                default:
                    v32 := GenS3{F0: (v2 & 831), F1: 234}
                    _ = v32
                }
                v2 -= 2147483647
                fmt.Println(fi27, fi27)
                _ = fi27
            }
            v23 = (v23 + v23)
        }
    }
    var v33 float32 = (((v4 / (3.14159 / -2.25)) * v4) + v4)
    v1 = append(v1, v23)
    v1 = append(v1, v23)
    v34 := GenS4{F0: -2147483648, F1: 32767}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
