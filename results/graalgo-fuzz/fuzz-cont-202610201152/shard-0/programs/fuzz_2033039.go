package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 string; F1 int64 }
func genFn0(p0 int32) uint64 {
    var v4 bool = false
    for fi5 := 0; fi5 < 9; fi5++ {
        fmt.Println(fi5, v4, p0)
        v6 := []int{}
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 bool = (2 != 2)
        v9 := GenS1{F0: "hello", F1: -1, F2: (p0 ^ p0)}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v4
    return 83
}
func genFn1(p0 uint, p1 int) int {
    if (147 < 0) {
        v10 := GenS2{F0: -738}
        var v11 uint = p0
        p1 += 2147483647
        _ = v10
        _ = v11
    } else {
        v12 := []int{}
        _ = v12
    }
    if (!false) {
        v13 := GenS3{F0: 471, F1: -54}
        for fi14 := 0; fi14 < 18; fi14++ {
            switch "hello" {
            case "a":
                p0 = p0
            case "foo bar":
                var v15 string = "hello"
                _ = v15
            case "hello":
                p0 = (p0 % (((p0 + 65535)) | 1))
            }
            _ = fi14
        }
        _ = v13
    }
    var v16 int16 = -1
    _ = p0
    _ = p1
    _ = v16
    return (p1 << (((p0 * p0)) & 63))
}

func main() {
    var v1 bool = (!(65535 == 958))
    var v2 uint64 = 18446744073709551615
    v3 := make(map[uint32]string)
    v2 -= (v2 | (v2 << ((530) & 63)))
    for fi17 := 0; fi17 < 11; fi17++ {
        v2 -= (v2 | (v2 / (((v2 | v2)) | 1)))
        _ = fi17
    }
    v2 += v2
    v18 := GenS4{F0: 0, F1: (v2 << ((255) & 63))}
    v19 := GenS4{F0: 2147483647, F1: (v2 % ((0) | 1))}
    var v20 string = "hello"
    var v21 bool = v1
    if ((v20 + v20) < v20) {
        var v22 uint64 = 42
        v23 := []int{32768, 42, 352, -665}
        switch v20 {
        case "":
            var v24 uint64 = v22
            _ = v24
        default:
            for fi25 := 0; fi25 < 2; fi25++ {
                for fi26 := 0; fi26 < 11; fi26++ {
                    fmt.Println(fi26)
                    delete(v3, 65535)
                    _ = fi26
                }
                if v1 {
                    var v27 uint8 = 255
                    _ = v27
                } else {
                    switch genFn1(43, 1) {
                    case -9223372036854775808:
                        var v28 int64 = -9223372036854775808
                        _ = v28
                    case 42:
                        fmt.Println(v20, v21, fi25)
                    default:
                        var v29 uint = 4294967295
                        _ = v29
                    }
                    v30 := v3[2]
                    var v31 int32 = 383
                    _ = v30
                    _ = v31
                }
                _ = fi25
            }
        }
        _ = v22
        _ = v23
    }
    var v32 int32 = -1
    v33 := []int{977, genFn1(2, 0), -2147483648, 32767}
    v20 = "foo bar"
    v20 = (v20 + v20)
    v20 = (("\n\t" + v20) + "\n\t")
    switch v20 {
    case "foo bar":
        for fi34 := 0; fi34 < 12; fi34++ {
            var v35 bool = v1
            v36 := GenS5{F0: "a", F1: 9223372036854775807}
            _ = fi34
            _ = v35
            _ = v36
        }
    case "hello":
        var v37 uint64 = 820
        _ = v37
    case "\n\t":
        if len(v33) >= 2 {
            v33 = v33[1:len(v33)]
        }
    default:
        if v21 {
            v20 = v20
            if false {
                fmt.Println(v1, v32)
                for fi38 := 0; fi38 < 5; fi38++ {
                    var v39 int64 = -1
                    var v40 string = (v20 + "foo bar")
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                v41 := GenS2{F0: v32}
                _ = v41
            } else {
                fmt.Println(v1, v20, v21)
            }
            v20 = ((v20 + (("\n\t" + v20) + "a")) + "hello")
        } else {
            if false {
                v21 = (((!(!(1 != v32))) && (!false)) || false)
                v1 = v1
            }
            var v42 bool = ((0.1 * float32(-2.25)) > 3.14159)
            _ = v42
        }
    }
    fmt.Println(v1, v2)
    v20 = v20
    var v43 float32 = ((0.1 * 0.0) - float32(3.14159))
    var v44 uint64 = (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v32
    _ = v33
    _ = v43
    _ = v44
}
