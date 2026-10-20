package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int16; F2 uint }
type GenS3 struct { F0 int64; F1 uint64; F2 int64 }
type GenS4 struct { F0 uint64; F1 int32 }
func genFn0(p0 int16) int16 {
    var v4 float32 = 1.5
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := []int{fi5}
        fmt.Println(p0, p0, fi5)
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 496}
    fmt.Println(p0, v4)
    v4 = v4
    _ = p0
    _ = v4
    _ = v7
    return (p0 / ((p0) | 1))
}
func genFn1(p0 int32, p1 uint) string {
    v8 := []int{1, 42, -1}
    p1 -= p1
    switch 260 {
    case 32767:
        if (100 != 42) {
            v8 = append(v8, -1)
            for fi9 := 0; fi9 < 14; fi9++ {
                switch "\n\t" {
                case "hello":
                    for fi10 := 0; fi10 < 10; fi10++ {
                        var v11 string = "hello"
                        _ = fi10
                        _ = v11
                    }
                }
                _ = len(v8)
                _ = fi9
            }
            p0 = p0
        }
    case -1:
        var v12 int64 = 9223372036854775807
        _ = v12
    case -2147483648:
        if true {
            v13 := GenS2{F0: 42, F1: genFn0(-638), F2: 65535}
            _ = v13
        }
    default:
        var v14 float32 = (0.0 * ((1e10 * 0.0) * (float32(float64((0.1 + 1e10))) + float32(1.5))))
        _ = v14
    }
    v15 := GenS2{F0: 42, F1: -32768, F2: (p1 + p1)}
    p0 = p0
    _ = p0
    _ = p1
    _ = v15
    _ = v8
    return ("foo bar" + "")
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 166
    var v3 uint8 = v2
    if (553 > 2) {
        v16 := GenS2{F0: 1, F1: 856, F2: 1}
        fmt.Println(v2, v3)
        v17 := GenS1{F0: 776}
        _ = v16
        _ = v17
    } else {
        var v18 int16 = 1000
        _ = v18
    }
    var v19 uint16 = 1
    var v20 int16 = 1000
    v21 := GenS2{F0: -110, F1: (v20 - -706), F2: 2}
    if (0 <= 0) {
        switch 9223372036854775807 {
        case -831:
            fmt.Println(v2, v3, v19)
        case 32767:
            v19 = (v19 << ((65535) & 63))
        default:
            for fi22 := 0; fi22 < 8; fi22++ {
                v23 := GenS1{F0: 32768}
                _ = fi22
                _ = v23
            }
        }
        switch ("hello" + ("\n\t" + "hello")) {
        case "":
            for fi24 := 0; fi24 < 14; fi24++ {
                fmt.Println(v20)
                _ = fi24
            }
        case "a":
            if true {
                for fi25 := 0; fi25 < 16; fi25++ {
                    v1[2] = "a"
                    for fi26 := 0; fi26 < 7; fi26++ {
                        v2 = 133
                        v20 -= v20
                        _ = fi26
                    }
                    _ = fi25
                }
                v19 -= (v19 % ((v19) | 1))
            } else {
                var v27 float32 = (-2.25 + 0.1)
                _ = v27
            }
        case "\n\t":
            v3 = 127
        }
        v2 -= (v3 | (^(v2)))
    } else {
        delete(v1, 0)
    }
    v28 := GenS3{F0: 221, F1: uint64(v2), F2: -1}
    for fi29 := 0; fi29 < 10; fi29++ {
        var v30 int32 = 1
        for fi31 := 0; fi31 < 1; fi31++ {
            v19 = (v19 & (v19 * (v19 & 266)))
            var v32 uint8 = v3
            var v33 float32 = 0.1
            v34 := []int{fi31, (fi31 / ((-9223372036854775808) | 1))}
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
        }
        v3 += (^(v2))
        _ = fi29
        _ = v30
    }
    v35 := GenS4{F0: 0, F1: 749}
    v20 = v20
    if false {
        switch "a" {
        case "\n\t":
            delete(v1, 0)
        case "foo bar":
            for fi36 := 0; fi36 < 7; fi36++ {
                var v37 int16 = -495
                _ = fi36
                _ = v37
            }
        }
        switch 42 {
        case 4294967295:
            v1[2] = ((("hello" + ("a" + "foo bar")) + "hello") + (("a" + "foo bar") + "\n\t"))
        case 224:
            delete(v1, 42)
        case 0:
            for fi38 := 0; fi38 < 19; fi38++ {
                var v39 bool = (-1 > 100)
                _ = fi38
                _ = v39
            }
        }
    } else {
        v1[2] = (((("\n\t" + ("foo bar" + "hello")) + (("a" + "a") + "a")) + "foo bar") + ("" + ""))
    }
    var v40 uint = 1
    v41 := GenS1{F0: -9223372036854775808}
    v42 := []int{}
    var v43 uint8 = 128
    fmt.Println(v19, v20, v43)
    v43 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v3
    _ = v35
    _ = v40
    _ = v41
    _ = v42
    _ = v43
}
