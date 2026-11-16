package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 string; F1 uint16 }
func genFn0(p0 uint) int {
    for fi6 := 0; fi6 < 12; fi6++ {
        p0 = (p0 & 4294967295)
        _ = fi6
    }
    if true {
        for fi7 := 0; fi7 < 16; fi7++ {
            p0 -= (p0 >> (((^(p0))) & 63))
            _ = fi7
        }
        fmt.Println(p0)
    } else {
        switch 42 {
        case -159:
            switch "" {
            case "foo bar":
                for fi8 := 0; fi8 < 14; fi8++ {
                    p0 -= (^(p0))
                    p0 += 4294967295
                    _ = fi8
                }
            case "a":
                if true {
                    for fi9 := 0; fi9 < 18; fi9++ {
                        v10 := []int{fi9, 42}
                        v10 = append(v10, fi9)
                        _ = fi9
                        _ = v10
                    }
                    if (32767 != 447) {
                        if (18446744073709551615 < 4294967295) {
                            p0 = (p0 - p0)
                        } else {
                            p0 = p0
                            p0 = (^(p0))
                        }
                        var v11 int16 = 444
                        switch 127 {
                        case 128:
                            var v12 float32 = 1.5
                            _ = v12
                        case 0:
                            fmt.Println(v11, v11)
                        }
                        _ = v11
                    }
                    if ((2147483647 >= -1) && true) {
                        var v13 int = -859
                        switch ((("" + "\n\t") + "a") + "hello") {
                        case "\n\t":
                            v14 := []int{}
                            _ = v14
                        case "hello":
                            v13 -= v13
                        case "":
                            fmt.Println(p0, v13)
                        }
                        if true {
                            p0 += 352
                            v15 := GenS1{F0: 231, F1: 1000}
                            if true {
                                var v16 uint = 4294967295
                                v16 = p0
                                _ = v16
                            }
                            _ = v15
                        }
                        _ = v13
                    }
                }
            }
        case -9223372036854775808:
            p0 -= 65535
        default:
            p0 += (p0 / ((p0) | 1))
        }
        v17 := []int{-509, 1, 1}
        _ = v17
    }
    if true {
        p0 += (^(p0))
        fmt.Println(p0)
        var v18 float32 = 3.14159
        _ = v18
    }
    p0 = (p0 | (p0 << ((p0) & 63)))
    _ = p0
    return -1
}
func genFn1(p0 uint, p1 uint16, p2 int16) int16 {
    switch 647 {
    case 557:
        for fi19 := 0; fi19 < 5; fi19++ {
            var v20 uint64 = 18446744073709551615
            _ = fi19
            _ = v20
        }
    case -221:
        v21 := []int{-9223372036854775808, 690}
        _ = v21
    }
    v22 := []int{42}
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return (p2 - (p2 % (((p2 - (p2 >> ((p0) & 63)))) | 1)))
}

func main() {
    var v1 float64 = 1.5
    var v2 uint32 = 83
    var v3 int16 = 70
    v4 := []string{"hello", "a", ("hello" + ("a" + "hello"))}
    var v5 uint = 65535
    fmt.Println(v5, v2)
    var v23 uint8 = uint8(v2)
    if true {
        var v24 float32 = -2.25
        _ = v24
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        fmt.Println(v5, v1, v2)
        var v26 string = "a"
        _ = fi25
        _ = v26
    }
    v3 += (v3 & (v3 * -1))
    var v27 uint64 = 164
    for fi28 := 0; fi28 < 17; fi28++ {
        var v29 int8 = -81
        v27 += (v27 * v27)
        fmt.Println(v1)
        _ = fi28
        _ = v29
    }
    var v30 bool = false
    if len(v4) > 0 {
        v4[0] = ("" + "")
    }
    v31 := GenS2{F0: -240, F1: 0}
    v32 := GenS3{F0: "\n\t", F1: 2}
    v2 = (v2 >> ((v5) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
