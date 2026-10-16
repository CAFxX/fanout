package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint }
func genFn0(p0 int16, p1 int16, p2 uint16) string {
    fmt.Println(p1, p1, p2)
    p0 = p0
    p0 -= p0
    if true {
        v5 := []int{-2147483648, 391, -1}
        _ = v5
    }
    if false {
        v6 := GenS1{F0: true, F1: 0}
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return ((("a" + "foo bar") + "a") + "")
}

func main() {
    var v1 int8 = 37
    var v2 int8 = (v1 >> ((597) & 63))
    v3 := []int{}
    v4 := make(map[int]uint32)
    for fi7 := 0; fi7 < 13; fi7++ {
        v1 = v1
        v1 -= v2
        v8 := v4[fi7]
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 uint32 = 1
        _ = fi9
        _ = v10
    }
    fmt.Println(v1, v1)
    fmt.Println(v2)
    for fi11 := 0; fi11 < 5; fi11++ {
        if (32767 != 42) {
            v1 = v1
        }
        switch ("a" + "") {
        case "\n\t":
            fmt.Println(fi11, fi11, v1)
        case "":
            for fi12 := 0; fi12 < 10; fi12++ {
                switch 213 {
                case 9223372036854775807:
                    if ((v2 == 100) || false) {
                        if false {
                            if true {
                                var v13 uint8 = 0
                                _ = v13
                            } else {
                                if true {
                                    if (132 == 4294967295) {
                                        var v14 uint16 = 255
                                        v3 = append(v3, fi12)
                                        _ = v14
                                    } else {
                                        v1 = (v1 - (v2 << ((0) & 63)))
                                        fmt.Println(v1, fi11)
                                    }
                                } else {
                                    v1 = v1
                                }
                                if true {
                                    v2 += (v2 >> ((807) & 63))
                                    var v15 int8 = v2
                                    _ = v15
                                }
                            }
                        }
                        if len(v3) > 0 {
                            v3[0] = 2147483647
                        }
                    }
                }
                v3 = append(v3, (fi12 | (^(fi11))))
                _ = fi12
            }
        default:
            v2 = (v1 / (((-(v1))) | 1))
        }
        switch "foo bar" {
        case "":
            switch fi11 {
            case 0:
                fmt.Println(v1)
            case 128:
                v2 = 14
            case 413:
                v2 = (^(v1))
            }
        default:
            if true {
                var v16 string = "hello"
                switch (fi11 % (((fi11 - fi11)) | 1)) {
                case -2147483648:
                    var v17 bool = (2 >= 1)
                    _ = v17
                }
                v1 -= 127
                _ = v16
            } else {
                v1 -= -1
                v2 = (v2 % (((v2 / (((-(v2))) | 1))) | 1))
                var v18 float32 = 0.1
                _ = v18
            }
        }
        _ = fi11
    }
    var v19 bool = (1000 < 622)
    fmt.Println(v19, v1, v2)
    v20 := GenS1{F0: v19, F1: 128}
    if false {
        fmt.Println(v19, v19)
        switch "" {
        case "a":
            if (42 < -1) {
                var v21 uint16 = 2
                v1 += (v2 >> ((255) & 63))
                _ = v21
            }
        case "":
            for fi22 := 0; fi22 < 18; fi22++ {
                v19 = (!v19)
                var v23 int = fi22
                v23 = (fi22 << ((4294967295) & 63))
                _ = fi22
                _ = v23
            }
        }
    }
    if v19 {
        v24 := v4[9223372036854775807]
        var v25 float32 = 1e10
        _ = v24
        _ = v25
    }
    if v19 {
        switch -231 {
        case 9223372036854775807:
            v19 = (1000 > 0)
        case 4294967295:
            delete(v4, -433)
        }
    }
    if (!((!false) && (0 != 128))) {
        v1 = v2
        for fi26 := 0; fi26 < 8; fi26++ {
            delete(v4, fi26)
            if (1 <= 2) {
                var v27 float32 = 1.5
                v19 = v19
                v28 := GenS1{F0: false, F1: 42}
                _ = v27
                _ = v28
            }
            var v29 int = (fi26 | (^(fi26)))
            _ = fi26
            _ = v29
        }
        v2 = 42
    } else {
        v19 = true
        var v30 uint32 = 0
        switch "foo bar" {
        case "":
            fmt.Println(v2, v19, v30)
        }
        _ = v30
    }
    switch "" {
    case "a":
        var v31 bool = (363 != -679)
        _ = v31
    case "\n\t":
        var v32 float64 = ((-2.25 * 3.14159) / ((-2.25 - 1.5) / 1.5))
        _ = v32
    case "":
        switch ("" + "a") {
        case "":
            var v33 uint64 = 294
            _ = v33
        case "hello":
            v34 := GenS2{F0: 18446744073709551615, F1: 1000}
            _ = v34
        case "a":
            var v35 uint = 65535
            _ = v35
        }
    default:
        v36 := GenS3{F0: 0, F1: 59, F2: 4294967295}
        _ = v36
    }
    fmt.Println(v2, v1, v19)
    v19 = v19
    v37 := GenS3{F0: 128, F1: 42, F2: 2}
    v2 = (-(v1))
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v4
}
