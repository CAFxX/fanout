package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int; F2 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint16) bool {
    var v7 int64 = 1
    p0 = 0
    var v8 bool = (true && true)
    v8 = (128 <= 214)
    var v9 int32 = 42
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return false
}
func genFn1(p0 int, p1 uint64) uint8 {
    p1 = p1
    var v10 int16 = -519
    v10 -= v10
    var v11 float64 = 0.1
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{4294967295, 42, 65535}
    var v3 uint = 65535
    var v4 float32 = float32(1e10)
    v5 := make(map[string]uint32)
    var v6 int16 = 1
    if (!(1 > genFn1(4294967295, 42))) {
        v4 = 0.1
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch 4294967295 {
        case 128:
            switch "" {
            case "a":
                v12 := GenS1{F0: "a"}
                _ = v12
            }
        case 4294967295:
            var v13 uint8 = 128
            _ = v13
        }
    }
    if false {
        switch ("hello" + "\n\t") {
        case "foo bar":
            delete(v1, 1)
        case "":
            var v14 uint16 = 498
            _ = v14
        }
        v15 := GenS1{F0: "a"}
        var v16 uint = v3
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 8; fi17++ {
            switch (fi17 ^ (fi17 % ((fi17) | 1))) {
            case -9223372036854775808:
                var v18 float64 = (float64(v4) + ((((((float64((((v4 - ((v4 + 3.14159) - 0.0)) + (((v4 + v4) / 0.1) / (((v4 / (v4 + v4)) - float32(((((3.14159 / 1.5) - ((-2.25 * 0.1) / -2.25)) - 3.14159) + 3.14159))) - v4))) * v4)) * -2.25) - 3.14159) / (((0.1 - 0.1) - float64((v4 - v4))) + 3.14159)) - (((float64(float32((1.5 + -2.25))) * 0.1) * ((3.14159 * float64(v4)) + (0.1 + 0.1))) / (-2.25 * 3.14159))) - 1.5) / (3.14159 / 1.5)))
                _ = v18
            case 127:
                if ((2 != 42) && ((^(v6)) != 32767)) {
                    v3 = (v3 - v3)
                    v19 := GenS2{F0: 127, F1: (^(fi17)), F2: uint16(v3)}
                    _ = v19
                } else {
                    delete(v1, 65535)
                    if false {
                        var v20 uint8 = genFn1(32767, 1)
                        fmt.Println(v3, v3)
                        _ = v20
                    } else {
                        var v21 float32 = 1e10
                        v22 := GenS1{F0: ("hello" + "foo bar")}
                        var v23 int = fi17
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    if true {
                        var v24 uint = (v3 | (v3 % ((65535) | 1)))
                        v24 -= 42
                        for fi25 := 0; fi25 < 9; fi25++ {
                            v26 := GenS2{F0: int8(v24), F1: 0, F2: 2}
                            var v27 uint16 = 799
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        _ = v24
                    } else {
                        v6 += v6
                    }
                }
            }
            _ = fi17
        }
    }
    if (0 != -652) {
        _ = len(v2)
    }
    fmt.Println(v4, v4)
    v4 = float32((0.0 - 1.5))
    delete(v1, 42)
    for fi28 := 0; fi28 < 8; fi28++ {
        v2 = append(v2, 2)
        switch "foo bar" {
        case "\n\t":
            v29 := GenS2{F0: 10, F1: (fi28 + fi28), F2: 32767}
            _ = v29
        case "foo bar":
            v30 := GenS3{F0: 628}
            _ = v30
        case "a":
            switch fi28 {
            case 42:
                fmt.Println(v3, v6, v4)
            default:
                if false {
                    for fi31 := 0; fi31 < 13; fi31++ {
                        _ = len(v1)
                        var v32 int = fi28
                        var v33 uint64 = 0
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                } else {
                    switch fi28 {
                    case -920:
                        var v34 uint16 = 533
                        _ = v34
                    }
                }
            }
        default:
            v3 += 1
        }
        switch fi28 {
        case -2147483648:
            var v35 uint32 = 2147483647
            _ = v35
        case -1:
            v3 = 4294967295
        case -9223372036854775808:
            switch "hello" {
            case "":
                var v36 string = ("\n\t" + "foo bar")
                _ = v36
            case "\n\t":
                v37 := v1[1]
                _ = v37
            case "hello":
                v2 = append(v2, 2147483648)
            }
        }
        _ = fi28
    }
    switch "" {
    case "foo bar":
        if false {
            fmt.Println(v3)
            for fi38 := 0; fi38 < 12; fi38++ {
                var v39 float64 = (-2.25 / 3.14159)
                v40 := GenS4{F0: 6}
                v41 := GenS5{F0: -32768, F1: genFn1(fi38, uint64(v3))}
                _ = fi38
                _ = v39
                _ = v40
                _ = v41
            }
        } else {
            delete(v5, "\n\t")
        }
    case "hello":
        fmt.Println(v3)
    case "a":
        if (42 != 1) {
            v42 := GenS5{F0: 32767, F1: 12}
            var v43 int8 = 1
            _ = v42
            _ = v43
        }
    default:
        var v44 int32 = -1
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
