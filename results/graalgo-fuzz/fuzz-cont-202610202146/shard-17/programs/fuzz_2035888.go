package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int16 }
func genFn0() int32 {
    var v7 int64 = 42
    v8 := GenS1{F0: 1, F1: false, F2: 32767}
    fmt.Println(v7)
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint, p1 uint32, p2 bool) int8 {
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 bool = false
        fmt.Println(p1)
        _ = fi9
        _ = v10
    }
    if (0 < -39) {
        v11 := []int{128}
        p0 = p0
        var v12 uint32 = (p1 / ((2147483647) | 1))
        _ = v11
        _ = v12
    }
    var v13 int8 = -128
    p1 -= 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return v13
}

func main() {
    v1 := []uint32{567, 4294967295, 2147483647}
    v2 := make(map[int]string)
    v3 := make(map[string]bool)
    v4 := make(map[int]string)
    v5 := make(map[int]string)
    var v6 int8 = -128
    if (v6 >= v6) {
        v1 = append(v1, 871)
        var v14 uint = 42
        _ = v14
    } else {
        v15 := GenS1{F0: v6, F1: false, F2: 0}
        v6 += 0
        _ = v15
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v16 string = ("hello" + (((("foo bar" + "foo bar") + "") + ("a" + "foo bar")) + "a"))
    if true {
        fmt.Println(v16)
        delete(v3, (v16 + v16))
        var v17 int16 = -988
        _ = v17
    }
    var v18 int32 = int32(v6)
    fmt.Println(v16, v18)
    for fi19 := 0; fi19 < 15; fi19++ {
        var v20 int64 = 1
        if (!true) {
            var v21 int32 = (v18 >> ((0) & 63))
            if true {
                v16 = v16
                if (0.1 != (-2.25 - 1.5)) {
                    switch ("" + v16) {
                    case "hello":
                        fmt.Println(v20, v6, v20)
                    case "foo bar":
                        for fi22 := 0; fi22 < 13; fi22++ {
                            if true {
                                var v23 float64 = 1e10
                                _ = v23
                            }
                            var v24 int64 = v20
                            var v25 uint16 = 1
                            _ = fi22
                            _ = v24
                            _ = v25
                        }
                    default:
                        switch (fi19 ^ 32767) {
                        case 0:
                            var v26 float32 = (-2.25 - -2.25)
                            _ = v26
                        default:
                            fmt.Println(v21, fi19)
                        }
                    }
                } else {
                    v21 = v21
                }
                fmt.Println(v18)
            }
            fmt.Println(v16, fi19)
            _ = v21
        }
        v6 = v6
        _ = fi19
        _ = v20
    }
    v27 := GenS1{F0: (v6 + v6), F1: (1 == 94), F2: 653}
    v16 = (v16 + "foo bar")
    switch -639 {
    case 845:
        v28 := v4[9223372036854775807]
        _ = v28
    case -2147483648:
        var v29 uint8 = 2
        _ = v29
    case 2147483647:
        var v30 uint32 = 33
        _ = v30
    default:
        fmt.Println(v18, v6)
    }
    var v31 uint16 = 374
    var v32 uint = 405
    v31 = (^(v31))
    for fi33 := 0; fi33 < 9; fi33++ {
        v31 += 42
        switch fi33 {
        case 128:
            v31 += 0
        case 0:
            for fi34 := 0; fi34 < 19; fi34++ {
                var v35 uint8 = 74
                if true {
                    v31 += (^(v31))
                }
                _ = fi34
                _ = v35
            }
        case 32768:
            v16 = ("a" + ("a" + ("foo bar" + (v16 + v16))))
        }
        _ = fi33
    }
    for fi36 := 0; fi36 < 9; fi36++ {
        delete(v2, fi36)
        _ = len(v1)
        v6 = v6
        _ = fi36
    }
    if false {
        fmt.Println(v18, v31)
        if false {
            var v37 bool = true
            if ((v37 && (-2.25 > 0.0)) || ((1.5 - (3.14159 * float64(1e10))) != (((1.5 + 1.5) * 1.5) - -2.25))) {
                v38 := GenS1{F0: v6, F1: (float64(1.5) >= ((float64(0.1) * 1e10) - 1e10)), F2: 1}
                var v39 string = ("hello" + (("hello" + v16) + (((("" + "a") + (((v16 + "hello") + v16) + v16)) + ("" + "a")) + (("\n\t" + (v16 + "foo bar")) + v16))))
                switch v16 {
                case "":
                    fmt.Println(v39, v18)
                case "\n\t":
                    var v40 int = -2147483648
                    _ = v40
                case "a":
                    var v41 uint32 = 4294967295
                    _ = v41
                default:
                    var v42 int16 = 42
                    _ = v42
                }
                _ = v38
                _ = v39
            } else {
                _ = len(v1)
                var v43 uint64 = 56
                _ = v43
            }
            v44 := GenS1{F0: (v6 * (v6 >> (((v32 ^ uint(v31))) & 63))), F1: (9223372036854775807 != -1), F2: 0}
            _ = v37
            _ = v44
        } else {
            var v45 int = -9223372036854775808
            var v46 uint64 = 2
            v6 = (v6 + genFn1((v32 - (v32 << (((v32 * v32)) & 63))), 366, true))
            _ = v45
            _ = v46
        }
    } else {
        v18 = (^(v18))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
