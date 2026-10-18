package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 bool; F2 string }
type GenS3 struct { F0 bool; F1 int; F2 uint }
type GenS4 struct { F0 uint8 }

func main() {
    v1 := []uint32{653, 467, 962, 574}
    var v2 float32 = 0.1
    var v3 uint16 = 206
    for fi4 := 0; fi4 < 13; fi4++ {
        for fi5 := 0; fi5 < 2; fi5++ {
            v2 = (((v2 + v2) * (v2 + v2)) / (v2 * (v2 + -2.25)))
            if (42 > 127) {
                var v6 int8 = -1
                fmt.Println(fi5)
                fmt.Println(v2, fi4)
                _ = v6
            } else {
                if false {
                    v3 += (v3 & v3)
                    var v7 float64 = -2.25
                    var v8 int = fi4
                    if true {
                        var v9 uint8 = 2
                        var v10 string = (((("hello" + ("" + "\n\t")) + "a") + "foo bar") + "foo bar")
                        v10 = v10
                        v1 = append(v1, 28)
                        _ = v10
                        _ = v9
                    }
                    _ = v7
                    _ = v8
                } else {
                    var v11 string = "a"
                    fmt.Println(fi5, fi4)
                    _ = v11
                }
                fmt.Println(fi4, v2)
            }
            v2 = 0.1
            switch ("foo bar" + "hello") {
            case "\n\t":
                v12 := GenS1{F0: (fi5 & (fi5 - (fi4 & 42)))}
                _ = v12
            default:
                v3 = v3
            }
            var v13 uint64 = 18446744073709551615
            var v14 uint64 = v13
            _ = fi5
            _ = v13
            _ = v14
        }
        _ = fi4
    }
    v3 = v3
    v3 = (v3 >> ((245) & 63))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v15 := GenS2{F0: 42, F1: true, F2: "\n\t"}
    if false {
        v16 := GenS2{F0: 2147483647, F1: (((float64(v2) + 1e10) / 1.5) > (((-2.25 * (-2.25 * (1.5 * 0.0))) * (3.14159 * (((((1e10 / float64(float32((0.1 * 1.5)))) + 1.5) / 1e10) - (1e10 * (0.0 / 1.0))) + 0.1))) / float64(((v2 - v2) / v2)))), F2: ""}
        v17 := GenS2{F0: 1, F1: false, F2: ("foo bar" + "")}
        _ = v16
        _ = v17
    }
    v18 := GenS2{F0: 2147483647, F1: true, F2: ("a" + "\n\t")}
    for fi19 := 0; fi19 < 13; fi19++ {
        v2 = 0.1
        if false {
            if (v3 != (v3 + v3)) {
                for fi20 := 0; fi20 < 11; fi20++ {
                    switch (^(fi20)) {
                    case -487:
                        _ = len(v1)
                    }
                    if true {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v3 = v3
                    }
                    _ = fi20
                }
                for fi21 := 0; fi21 < 9; fi21++ {
                    var v22 uint64 = 2
                    switch fi21 {
                    case -1:
                        v23 := GenS3{F0: (false || false), F1: (fi19 * fi21), F2: 65535}
                        _ = v23
                    case 32768:
                        v24 := GenS2{F0: -934, F1: (4294967295 <= 65535), F2: "foo bar"}
                        var v25 bool = false
                        v3 -= (^(v3))
                        _ = v24
                        _ = v25
                    default:
                        var v26 bool = false
                        _ = v26
                    }
                    _ = fi21
                    _ = v22
                }
                if (18446744073709551615 <= 4294967295) {
                    v2 += v2
                }
            } else {
                if len(v1) > 0 {
                    v1[0] = 1
                }
                fmt.Println(fi19)
                var v27 bool = true
                _ = v27
            }
            if (!("a" != "a")) {
                switch fi19 {
                case -9223372036854775808:
                    v3 -= (v3 % (((v3 & (v3 & v3))) | 1))
                case -2147483648:
                    for fi28 := 0; fi28 < 8; fi28++ {
                        v1 = append(v1, 1)
                        fmt.Println(v3)
                        _ = fi28
                    }
                default:
                    if false {
                        v2 = v2
                        v29 := GenS1{F0: -598}
                        fmt.Println(v3)
                        _ = v29
                    }
                }
            } else {
                var v30 uint64 = 408
                v31 := GenS1{F0: (^(fi19))}
                v32 := GenS4{F0: 255}
                _ = v30
                _ = v31
                _ = v32
            }
        }
        _ = fi19
    }
    var v33 uint32 = uint32(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v33
}
