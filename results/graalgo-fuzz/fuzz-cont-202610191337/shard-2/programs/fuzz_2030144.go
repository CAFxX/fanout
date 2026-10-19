package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int16 }
func genFn0() int64 {
    var v4 int8 = 100
    var v5 float32 = 0.1
    var v6 uint32 = 2147483647
    _ = v4
    _ = v5
    _ = v6
    return 9223372036854775807
}
func genFn1() int64 {
    var v7 uint8 = 195
    v8 := []int{32768}
    switch "hello" {
    case "hello":
        v7 = (v7 % ((0) | 1))
    }
    fmt.Println(v7)
    _ = v7
    _ = v8
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = 0
    var v3 int16 = -23
    var v9 float32 = float32(0.1)
    v10 := GenS1{F0: "\n\t", F1: -667}
    delete(v1, 987)
    if false {
        for fi11 := 0; fi11 < 18; fi11++ {
            v2 -= v2
            v2 -= (v2 / ((-2147483648) | 1))
            if (448 < 18446744073709551615) {
                for fi12 := 0; fi12 < 4; fi12++ {
                    fmt.Println(v2, fi11)
                    _ = fi12
                }
                var v13 int16 = v3
                var v14 string = ("" + ("a" + ("" + "\n\t")))
                _ = v13
                _ = v14
            }
            _ = fi11
        }
        v15 := []int{}
        var v16 uint32 = 2147483648
        _ = v15
        _ = v16
    }
    var v17 bool = (!(206 <= 879))
    var v18 int8 = 100
    var v19 int8 = -1
    if (43 < 127) {
        v20 := GenS1{F0: "hello", F1: 327}
        _ = v20
    }
    v21 := []int{}
    switch -9223372036854775808 {
    case -2147483648:
        v17 = v17
    case 32768:
        for fi22 := 0; fi22 < 3; fi22++ {
            var v23 uint16 = 14
            v24 := GenS1{F0: ("a" + ("\n\t" + ("" + "hello"))), F1: (^(fi22))}
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    _ = len(v21)
    v3 -= v3
    var v25 uint8 = 127
    if v17 {
        for fi26 := 0; fi26 < 11; fi26++ {
            v2 += v2
            var v27 float64 = (0.0 - 3.14159)
            _ = fi26
            _ = v27
        }
        if (!((float64((float32(3.14159) / (v9 * v9))) >= float64(v9)) || true)) {
            switch 1 {
            case 32767:
                switch "\n\t" {
                case "\n\t":
                    v17 = v17
                case "":
                    if (0 < 870) {
                        v28 := GenS1{F0: "hello", F1: 4294967295}
                        var v29 uint32 = 0
                        var v30 float64 = 1.5
                        _ = v28
                        _ = v29
                        _ = v30
                    } else {
                        _ = len(v21)
                    }
                }
            case 0:
                switch "a" {
                case "hello":
                    v31 := GenS2{F0: (v3 - (v3 / (((v3 | v3)) | 1))), F1: "\n\t"}
                    _ = v31
                case "\n\t":
                    switch "" {
                    case "\n\t":
                        switch 32768 {
                        case 32768:
                            v19 += v18
                        }
                    case "hello":
                        if (false && v17) {
                            var v32 int8 = v18
                            _ = v32
                        }
                    case "a":
                        var v33 uint = 1
                        _ = v33
                    }
                case "a":
                    for fi34 := 0; fi34 < 9; fi34++ {
                        for fi35 := 0; fi35 < 2; fi35++ {
                            var v36 int64 = 740
                            _ = fi35
                            _ = v36
                        }
                        v2 = 231
                        var v37 float64 = -2.25
                        _ = fi34
                        _ = v37
                    }
                }
            case 49:
                if (true || true) {
                    var v38 uint16 = 0
                    if (!true) {
                        v21 = append(v21, 1)
                    } else {
                        if v17 {
                            for fi39 := 0; fi39 < 11; fi39++ {
                                v19 += (v18 | int8(v3))
                                for fi40 := 0; fi40 < 1; fi40++ {
                                    var v41 uint16 = v38
                                    if (v3 > (^(v3))) {
                                        var v42 uint8 = v25
                                        _ = v42
                                    } else {
                                        var v43 uint64 = 2
                                        switch "" {
                                        case "\n\t":
                                            v19 = (v19 >> ((255) & 63))
                                        default:
                                            v19 -= int8(v2)
                                        }
                                        _ = v43
                                    }
                                    v18 = v18
                                    _ = fi40
                                    _ = v41
                                }
                                _ = fi39
                            }
                            delete(v1, 0)
                            fmt.Println(v9, v19, v9)
                        } else {
                            v9 += v9
                        }
                        fmt.Println(v3)
                        v19 += v18
                    }
                    fmt.Println(v25, v3)
                    _ = v38
                } else {
                    v44 := v1[127]
                    v45 := GenS3{F0: 42}
                    _ = v44
                    _ = v45
                }
            }
            v9 = 0.1
        }
    } else {
        var v46 int8 = v19
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v9
}
