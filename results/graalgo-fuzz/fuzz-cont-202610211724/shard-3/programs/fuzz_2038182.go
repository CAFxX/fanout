package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int16; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint16; F1 int; F2 string }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 uint16; F1 int16 }
func genFn0(p0 uint32, p1 uint16, p2 uint8) int32 {
    if ("" > "foo bar") {
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 float64 = 1e10
            _ = fi6
            _ = v7
        }
    } else {
        if (false || (p0 != (p0 & (^(p0))))) {
            var v8 uint = 255
            _ = v8
        }
    }
    v9 := []int{32767, 95, -2147483648, 127}
    p1 += p1
    p2 = 233
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -2147483648
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 2
    var v3 uint16 = 42
    var v4 uint16 = (v3 >> ((255) & 63))
    var v5 float64 = (-2.25 + float64(float32(((1e10 + 1e10) + 0.1))))
    var v10 float64 = 1.5
    switch ("a" + "\n\t") {
    case "a":
        delete(v1, 2147483647)
    default:
        var v11 uint8 = 103
        _ = v11
    }
    v3 = 2
    switch 42 {
    case 4294967295:
        v1[0] = ("a" + "hello")
    case 2147483647:
        v12 := GenS1{F0: ("a" + ("a" + (((("\n\t" + ("hello" + "hello")) + "\n\t") + "\n\t") + ("a" + "a")))), F1: 492}
        _ = v12
    }
    v2 = (v3 % ((v2) | 1))
    if false {
        v13 := v1[297]
        _ = v13
    }
    v14 := []int{42, 42, 9223372036854775807}
    _ = len(v14)
    switch -785 {
    case 0:
        var v15 uint = 65535
        _ = v15
    case -9223372036854775808:
        for fi16 := 0; fi16 < 3; fi16++ {
            v5 += 0.0
            delete(v1, 65535)
            _ = fi16
        }
    case -2147483648:
        var v17 int = 127
        _ = v17
    }
    if false {
        switch 1 {
        case 32768:
            v18 := GenS1{F0: "hello", F1: (v2 | (v3 / (((v4 / ((v2) | 1))) | 1)))}
            _ = v18
        case 137:
            v1[42] = "\n\t"
        }
    } else {
        if true {
            var v19 uint8 = 157
            if false {
                var v20 uint32 = 463
                v21 := GenS2{F0: 42, F1: 127, F2: (v3 % ((uint16(v19)) | 1))}
                if false {
                    fmt.Println(v19, v10, v2)
                    if false {
                        switch "" {
                        case "":
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                        default:
                            v3 = v3
                        }
                        v2 -= v2
                        v22 := GenS3{F0: (v4 % (((v4 >> ((2) & 63))) | 1)), F1: 860, F2: "\n\t"}
                        _ = v22
                    } else {
                        if (v3 > v4) {
                            var v23 string = ("\n\t" + "foo bar")
                            _ = len(v1)
                            _ = v23
                        }
                        var v24 bool = false
                        _ = v24
                    }
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                }
                _ = v20
                _ = v21
            }
            _ = v19
        } else {
            v4 -= v3
            v25 := GenS4{F0: 0, F1: -976}
            _ = v25
        }
    }
    var v26 uint8 = 25
    fmt.Println(v3)
    var v27 uint16 = v3
    var v28 uint = 739
    switch 2147483647 {
    case 42:
        var v29 uint16 = v2
        _ = v29
    default:
        v30 := GenS3{F0: (v4 % (((v27 >> ((65535) & 63))) | 1)), F1: 0, F2: "a"}
        _ = v30
    }
    for fi31 := 0; fi31 < 16; fi31++ {
        v32 := GenS5{F0: (v4 << (((^(v28))) & 63)), F1: -1}
        switch -2147483648 {
        case 4294967295:
            for fi33 := 0; fi33 < 10; fi33++ {
                var v34 int16 = 1
                v35 := GenS5{F0: v27, F1: (v34 ^ 296)}
                v4 -= (v4 >> ((v28) & 63))
                _ = fi33
                _ = v34
                _ = v35
            }
        }
        for fi36 := 0; fi36 < 1; fi36++ {
            v37 := GenS4{F0: -510, F1: 0}
            _ = fi36
            _ = v37
        }
        _ = fi31
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
