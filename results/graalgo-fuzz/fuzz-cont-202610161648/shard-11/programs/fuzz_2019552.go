package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 string }
func genFn0() int64 {
    if false {
        for fi5 := 0; fi5 < 7; fi5++ {
            var v6 int64 = 1
            for fi7 := 0; fi7 < 10; fi7++ {
                v8 := []int{4294967295, 887, 32767}
                if len(v8) > 0 {
                    v8[0] = (fi7 | (fi7 / ((2147483647) | 1)))
                }
                _ = fi7
                _ = v8
            }
            v9 := GenS1{F0: 255}
            _ = fi5
            _ = v6
            _ = v9
        }
    } else {
        switch "hello" {
        case "hello":
            switch 4294967295 {
            case 2147483647:
                var v10 uint32 = 42
                _ = v10
            }
        case "a":
            if (!(42 >= -281)) {
                v11 := []int{}
                _ = v11
            }
        default:
            v12 := GenS2{F0: 32767, F1: -2147483648}
            _ = v12
        }
        var v13 float32 = 3.14159
        v13 -= float32((1.5 * 0.1))
        _ = v13
    }
    var v14 int16 = 32767
    _ = v14
    return -1
}

func main() {
    var v1 float64 = 1.5
    var v2 uint8 = 255
    var v3 int16 = -1
    v4 := make(map[string]bool)
    var v15 int64 = 42
    if ("\n\t" < ("hello" + "\n\t")) {
        v15 += 9223372036854775807
        var v16 uint8 = (v2 / ((79) | 1))
        _ = v16
    } else {
        v17 := GenS1{F0: 65535}
        v3 -= v3
        var v18 string = "\n\t"
        _ = v17
        _ = v18
    }
    v1 -= v1
    v19 := GenS3{F0: ("a" + "")}
    v4[((((((("a" + ("foo bar" + ("hello" + "hello"))) + "foo bar") + "foo bar") + "") + "foo bar") + ("foo bar" + "hello")) + "a")] = (65535 < 1)
    if false {
        if true {
            if true {
                switch ("" + ("" + "foo bar")) {
                case "hello":
                    if ((3.14159 + (((((1e10 * 0.0) - (float32(v1) * (-2.25 / float32(((3.14159 + (((v1 - v1) + 1e10) / 3.14159)) * (-2.25 * float64(0.1))))))) + 0.0) / float32(float64(0.1))) - 1.5)) > -2.25) {
                        for fi20 := 0; fi20 < 17; fi20++ {
                            var v21 int = (fi20 >> ((1) & 63))
                            _ = fi20
                            _ = v21
                        }
                        var v22 uint = 42
                        _ = v22
                    }
                case "":
                    if true {
                        var v23 uint32 = 988
                        v15 -= (-(v15))
                        var v24 int64 = genFn0()
                        _ = v23
                        _ = v24
                    } else {
                        if (-705 < -2147483648) {
                            for fi25 := 0; fi25 < 17; fi25++ {
                                v2 = (v2 << ((4294967295) & 63))
                                _ = fi25
                            }
                            v15 -= (v15 % ((v15) | 1))
                        }
                        if (!(!(44 != -74))) {
                            switch ("" + "\n\t") {
                            case "a":
                                v1 += 0.1
                            case "\n\t":
                                v2 += v2
                            case "":
                                fmt.Println(v1, v3, v3)
                            default:
                                fmt.Println(v1, v2)
                            }
                            fmt.Println(v15, v15, v1)
                        } else {
                            var v26 int32 = -2147483648
                            _ = v26
                        }
                    }
                case "foo bar":
                    for fi27 := 0; fi27 < 11; fi27++ {
                        v1 += v1
                        _ = fi27
                    }
                }
                for fi28 := 0; fi28 < 1; fi28++ {
                    v2 = v2
                    _ = fi28
                }
            }
        }
        var v29 bool = false
        _ = v29
    } else {
        v2 += (v2 | v2)
        for fi30 := 0; fi30 < 10; fi30++ {
            v31 := GenS3{F0: "\n\t"}
            fmt.Println(fi30, fi30, v1)
            _ = fi30
            _ = v31
        }
        v15 = -1
    }
    v32 := GenS1{F0: 4294967295}
    v2 = (v2 + (v2 % ((1) | 1)))
    v33 := []int{693, 42}
    fmt.Println(v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
