package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string }
func genFn0(p0 uint32, p1 int16) uint64 {
    var v7 float64 = 1e10
    p1 += p1
    if (!(!(2 != 151))) {
        var v8 string = ((("\n\t" + "") + (("" + "foo bar") + "hello")) + ("\n\t" + "\n\t"))
        _ = v8
    } else {
        if false {
            var v9 float64 = v7
            var v10 int64 = -1
            _ = v10
            _ = v9
        }
        v11 := GenS1{F0: -1, F1: 0, F2: 42}
        _ = v11
    }
    v12 := []int{9223372036854775807, 127, 4294967295}
    switch 4294967295 {
    case 959:
        fmt.Println(p0)
    case 988:
        p0 = p0
    default:
        fmt.Println(v7, p1)
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v7
    return 42
}
func genFn1(p0 int8, p1 uint64) int {
    v13 := GenS2{F0: 415, F1: 32767}
    v14 := GenS1{F0: -1, F1: 0, F2: (p0 % ((p0) | 1))}
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := GenS2{F0: 42, F1: 255}
        v17 := []int{(fi15 * 127), fi15, fi15}
        p0 = (p0 - p0)
        _ = fi15
        _ = v16
        _ = v17
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 4294967295
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{}
    v3 := make(map[uint32]string)
    var v4 bool = (false || false)
    v5 := []uint32{65535}
    var v6 bool = v4
    v6 = v4
    v18 := GenS2{F0: 42, F1: 692}
    var v19 int16 = -94
    v20 := GenS3{F0: 1}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v21 := GenS4{F0: 128, F1: 2, F2: -1}
    v1[1] = (v4 && v6)
    if v4 {
        if v6 {
            delete(v3, 754)
            if (true || (!(308 <= 434))) {
                var v22 bool = v4
                v23 := GenS5{F0: 4294967295}
                _ = v22
                _ = v23
            }
            var v24 uint8 = 0
            _ = v24
        }
    } else {
        fmt.Println(v4, v19)
        for fi25 := 0; fi25 < 19; fi25++ {
            for fi26 := 0; fi26 < 15; fi26++ {
                var v27 int32 = -1
                switch "hello" {
                case "":
                    if (!v6) {
                        v19 = v19
                    } else {
                        var v28 float64 = float64((float32(1.5) * 1e10))
                        _ = v28
                    }
                case "hello":
                    switch ("foo bar" + ("a" + "\n\t")) {
                    case "a":
                        v29 := GenS2{F0: 2147483648, F1: 928}
                        _ = v29
                    case "\n\t":
                        if (1 >= 0) {
                            v4 = true
                        } else {
                            if (!v6) {
                                _ = len(v2)
                                switch ("hello" + "foo bar") {
                                case "foo bar":
                                    v6 = (false && ((3.14159 / (1.5 + (float32((1.5 / 3.14159)) - -2.25))) != (0.1 / 0.1)))
                                default:
                                    var v30 uint32 = 523
                                    _ = v30
                                }
                            }
                        }
                    default:
                        var v31 bool = (!(v6 && v6))
                        _ = v31
                    }
                default:
                    switch fi25 {
                    case 0:
                        v27 -= 0
                    case 404:
                        v19 = 0
                    default:
                        var v32 int32 = (v27 + (v27 % ((v27) | 1)))
                        fmt.Println(fi25, v19, fi26)
                        _ = v32
                    }
                }
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
        var v33 float32 = 3.14159
        _ = v33
    }
    var v34 int8 = 1
    v19 = v19
    for fi35 := 0; fi35 < 11; fi35++ {
        v6 = (v34 > (^(v34)))
        _ = fi35
    }
    for fi36 := 0; fi36 < 14; fi36++ {
        switch "a" {
        case "foo bar":
            var v37 float64 = 0.0
            _ = v37
        }
        if (!false) {
            fmt.Println(v19, v34, fi36)
            if (1000 >= 42) {
                v4 = ((fi36 & 32767) != fi36)
            }
        }
        var v38 int32 = 2147483647
        _ = fi36
        _ = v38
    }
    if (!(673 == 186)) {
        v39 := GenS6{F0: ("\n\t" + ("" + "\n\t"))}
        v19 = v19
        _ = v39
    }
    switch 290 {
    case -497:
        v4 = v4
    case 32768:
        var v40 uint8 = 2
        _ = v40
    }
    if len(v2) > 0 {
        v2[0] = -1
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
