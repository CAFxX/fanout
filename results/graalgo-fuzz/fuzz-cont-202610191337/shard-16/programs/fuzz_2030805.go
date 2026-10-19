package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 string }
type GenS2 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 bool; F1 uint8 }
func genFn0() uint64 {
    switch "" {
    case "a":
        var v7 int8 = 100
        _ = v7
    case "hello":
        var v8 uint = 640
        _ = v8
    case "\n\t":
        var v9 float32 = 0.0
        _ = v9
    }
    v10 := GenS1{F0: 9223372036854775807, F1: 4294967295, F2: ("a" + "\n\t")}
    _ = v10
    return 369
}
func genFn1() bool {
    var v11 int32 = 2147483647
    var v12 float32 = float32(0.0)
    v12 = ((float32(1e10) - float32(0.0)) * v12)
    _ = v11
    _ = v12
    return (true || (false && (42 >= 1)))
}

func main() {
    var v1 float64 = (3.14159 - -2.25)
    var v2 uint = 24
    v3 := []string{(("hello" + "foo bar") + "foo bar"), "foo bar", "\n\t", ("\n\t" + "hello")}
    var v4 int = 4294967295
    v5 := []uint32{833, 2147483647, 65535, 65535}
    v6 := make(map[int]bool)
    if len(v5) > 0 {
        v5[0] = 2147483648
    }
    v2 = v2
    var v13 int = (v4 | (v4 ^ -9223372036854775808))
    fmt.Println(v4, v13, v4)
    for fi14 := 0; fi14 < 15; fi14++ {
        switch (^(v4)) {
        case 9223372036854775807:
            if true {
                v2 = (v2 & 0)
                if (1 != 1) {
                    if (-1 < -931) {
                        for fi15 := 0; fi15 < 3; fi15++ {
                            v16 := GenS1{F0: 125, F1: 2, F2: "hello"}
                            v1 += 1.5
                            var v17 int32 = -2147483648
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                        var v18 int64 = 9223372036854775807
                        _ = v18
                    } else {
                        fmt.Println(fi14, fi14, v4)
                        v1 -= 3.14159
                        if true {
                            v2 -= v2
                            fmt.Println(v13, v1)
                        } else {
                            if ((v1 < (3.14159 / (-2.25 + (-2.25 / v1)))) && true) {
                                v19 := GenS1{F0: -794, F1: 584, F2: "a"}
                                var v20 string = "\n\t"
                                _ = v19
                                _ = v20
                            }
                        }
                    }
                    if (529 >= 2147483647) {
                        v21 := GenS2{F0: 255, F1: 2147483647, F2: 42}
                        v1 -= 1.5
                        _ = v21
                    }
                }
                switch (fi14 + v4) {
                case 758:
                    _ = len(v3)
                case 81:
                    v1 = v1
                case 9223372036854775807:
                    v22 := GenS3{F0: 789, F1: 18446744073709551615, F2: 911}
                    _ = v22
                default:
                    var v23 uint = v2
                    _ = v23
                }
            }
        default:
            var v24 string = "hello"
            _ = v24
        }
        v25 := GenS4{F0: 18446744073709551615, F1: -234}
        _ = fi14
        _ = v25
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v6)
    for fi26 := 0; fi26 < 14; fi26++ {
        for fi27 := 0; fi27 < 5; fi27++ {
            var v28 uint = v2
            if ((!false) && true) {
                var v29 int8 = 100
                _ = v29
            } else {
                v28 -= v28
            }
            var v30 uint64 = uint64(v28)
            _ = fi27
            _ = v28
            _ = v30
        }
        var v31 uint8 = uint8(v2)
        _ = fi26
        _ = v31
    }
    v6[(v13 - (v13 + (v4 + v13)))] = false
    var v32 float32 = 0.1
    v33 := v6[(v13 + v4)]
    for fi34 := 0; fi34 < 3; fi34++ {
        var v35 uint = v2
        for fi36 := 0; fi36 < 17; fi36++ {
            var v37 int32 = -2147483648
            _ = fi36
            _ = v37
        }
        var v38 uint64 = 18446744073709551615
        _ = fi34
        _ = v35
        _ = v38
    }
    switch ("" + "a") {
    case "foo bar":
        var v39 uint = v2
        _ = v39
    case "\n\t":
        switch 128 {
        case -1:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case 0:
            var v40 int64 = -9223372036854775808
            _ = v40
        case -9223372036854775808:
            v4 += v4
        }
    case "hello":
        fmt.Println(v33, v2)
    default:
        var v41 int16 = 0
        _ = v41
    }
    if len(v5) > 0 {
        v5[0] = 42
    }
    switch "a" {
    case "foo bar":
        var v42 string = ("hello" + "hello")
        _ = v42
    case "hello":
        _ = len(v5)
    default:
        var v43 uint8 = 234
        _ = v43
    }
    switch "hello" {
    case "a":
        v13 += v13
    case "":
        if (!((2147483647 == 1) || ((v33 && v33) || v33))) {
            if (0 < 2) {
                v4 += (v13 % ((v13) | 1))
                v13 = v4
            } else {
                v44 := GenS5{F0: (v33 && (0 >= 128)), F1: 59}
                _ = v44
            }
            var v45 uint8 = 105
            var v46 int = (v13 - (v4 + v13))
            _ = v45
            _ = v46
        }
    case "foo bar":
        for fi47 := 0; fi47 < 11; fi47++ {
            var v48 uint = (v2 >> ((750) & 63))
            fmt.Println(v2, v33)
            var v49 float64 = (float64(v32) + float64(v32))
            _ = fi47
            _ = v48
            _ = v49
        }
    }
    var v50 float64 = float64(v32)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v50
    _ = v6
}
