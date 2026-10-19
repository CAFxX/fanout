package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 uint64; F2 int8 }
func genFn0(p0 uint8, p1 int, p2 int64) int8 {
    v5 := GenS1{F0: (p0 >> ((2) & 63)), F1: "hello", F2: uint64(p0)}
    var v6 uint16 = 875
    var v7 float32 = (0.0 + -2.25)
    var v8 uint8 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -99
}
func genFn1(p0 int8, p1 int32, p2 int32) uint32 {
    if true {
        v9 := []int{-1, 0, -705, 32767}
        _ = v9
    } else {
        var v10 int32 = 2147483647
        if (127 <= 42) {
            if ("foo bar" != "\n\t") {
                v11 := GenS2{F0: 65535, F1: p1, F2: 995}
                v12 := GenS1{F0: 2, F1: "", F2: 1}
                _ = v11
                _ = v12
            } else {
                var v13 uint32 = 4294967295
                var v14 uint32 = 9
                for fi15 := 0; fi15 < 9; fi15++ {
                    v10 += p1
                    var v16 float32 = 0.0
                    _ = fi15
                    _ = v16
                }
                _ = v13
                _ = v14
            }
            v17 := GenS2{F0: 2147483648, F1: 1, F2: 570}
            _ = v17
        } else {
            p1 += v10
            var v18 uint32 = 0
            _ = v18
        }
        _ = v10
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    return 93
}

func main() {
    var v1 int64 = -597
    v2 := []string{}
    v3 := []int{32767, 42, 739}
    v4 := make(map[int]string)
    var v19 uint16 = 1
    var v20 uint16 = v19
    if len(v3) > 0 {
        v3[0] = -9223372036854775808
    }
    v21 := GenS3{F0: 127}
    var v22 int8 = genFn0(128, 42, v1)
    fmt.Println(v20)
    switch ("a" + "foo bar") {
    case "a":
        var v23 int8 = (v22 / ((v22) | 1))
        _ = v23
    case "foo bar":
        v24 := GenS1{F0: 86, F1: "a", F2: 4294967295}
        _ = v24
    }
    v19 -= (v20 * 501)
    v19 = (v20 + 32767)
    for fi25 := 0; fi25 < 20; fi25++ {
        var v26 bool = ((!false) && (529 == -32768))
        _ = fi25
        _ = v26
    }
    var v27 int16 = 1000
    var v28 int = 0
    var v29 float64 = (-2.25 / (((3.14159 + 0.1) / 1.0) + 1e10))
    switch 792 {
    case 2147483647:
        var v30 int = (v28 & (v28 << ((0) & 63)))
        _ = v30
    default:
        switch v28 {
        case -59:
            v31 := GenS2{F0: 42, F1: -627, F2: v20}
            _ = v31
        case 1:
            switch "\n\t" {
            case "":
                v1 -= (v1 >> ((18446744073709551615) & 63))
            case "hello":
                v32 := GenS1{F0: 110, F1: ("" + "hello"), F2: 782}
                _ = v32
            default:
                if true {
                    var v33 uint = 65535
                    var v34 uint16 = 32767
                    _ = v33
                    _ = v34
                } else {
                    var v35 uint8 = uint8(v19)
                    for fi36 := 0; fi36 < 9; fi36++ {
                        v37 := GenS4{F0: "foo bar", F1: 238, F2: (v22 >> ((78) & 63))}
                        switch "hello" {
                        case "a":
                            v3 = append(v3, (^(fi36)))
                        }
                        for fi38 := 0; fi38 < 5; fi38++ {
                            switch "\n\t" {
                            case "\n\t":
                                var v39 string = "\n\t"
                                _ = v39
                            }
                            switch 2147483647 {
                            case 9223372036854775807:
                                fmt.Println(fi38, v29)
                            }
                            v40 := GenS1{F0: 127, F1: ("hello" + "foo bar"), F2: 2}
                            _ = fi38
                            _ = v40
                        }
                        _ = fi36
                        _ = v37
                    }
                    _ = v35
                }
            }
        case 610:
            for fi41 := 0; fi41 < 7; fi41++ {
                for fi42 := 0; fi42 < 12; fi42++ {
                    v29 += (1e10 + v29)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi42
                }
                v1 = int64(v20)
                _ = fi41
            }
        }
    }
    var v43 float32 = (1.5 + -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v43
}
