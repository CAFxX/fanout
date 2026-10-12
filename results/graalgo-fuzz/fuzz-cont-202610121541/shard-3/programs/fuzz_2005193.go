package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint16; F2 int64 }
type GenS3 struct { F0 string; F1 int }

func main() {
    var v1 float32 = 1e10
    var v2 int32 = -2147483648
    var v3 uint32 = 2147483647
    v4 := make(map[uint32]bool)
    var v5 uint64 = 2
    var v6 uint64 = v5
    if true {
        for fi7 := 0; fi7 < 11; fi7++ {
            fmt.Println(v6)
            _ = fi7
        }
    }
    switch 2147483647 {
    case -919:
        v5 -= 864
    case 32767:
        v1 = v1
    default:
        if ((v2 - v2) == (v2 & -200)) {
            if (!true) {
                v3 = v3
                switch "\n\t" {
                case "a":
                    switch ("foo bar" + "hello") {
                    case "foo bar":
                        v8 := []int{-1}
                        _ = v8
                    case "":
                        v6 += 18446744073709551615
                    }
                case "hello":
                    var v9 int64 = -1
                    _ = v9
                default:
                    if false {
                        v10 := GenS1{F0: 81}
                        switch "hello" {
                        case "a":
                            v11 := []int{-1, 402}
                            _ = v11
                        }
                        var v12 float64 = ((((1.5 * ((0.0 - 3.14159) / (1e10 + -2.25))) / (3.14159 / 3.14159)) * float64((0.0 / 0.1))) * 3.14159)
                        _ = v10
                        _ = v12
                    } else {
                        v13 := GenS1{F0: 1}
                        v1 += v1
                        _ = v13
                    }
                }
            }
        } else {
            v3 = (v3 + (v3 / ((99) | 1)))
        }
    }
    v4[4294967295] = true
    if (!true) {
        for fi14 := 0; fi14 < 6; fi14++ {
            v5 += (v6 >> ((65535) & 63))
            v2 = v2
            _ = fi14
        }
    } else {
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 bool = true
            v6 = (v5 << ((42) & 63))
            v2 = v2
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 5; fi17++ {
            v18 := v4[(v3 >> ((18446744073709551615) & 63))]
            v4[(^(v3))] = (false || (!(255 <= 341)))
            var v19 uint32 = v3
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    if (994 <= (v2 ^ (-(v2)))) {
        v6 = 0
    } else {
        v20 := []int{-837, 9223372036854775807, 32768, 919}
        var v21 float64 = ((0.0 * 1.5) + (-2.25 * 0.1))
        _ = v20
        _ = v21
    }
    fmt.Println(v6, v1)
    for fi22 := 0; fi22 < 11; fi22++ {
        for fi23 := 0; fi23 < 8; fi23++ {
            v2 -= v2
            _ = fi23
        }
        if (32767 < 1) {
            var v24 uint16 = 571
            var v25 float64 = float64((((1e10 / v1) * v1) + v1))
            v26 := GenS2{F0: "\n\t", F1: 42, F2: -9223372036854775808}
            _ = v24
            _ = v25
            _ = v26
        } else {
            v27 := GenS2{F0: "", F1: 255, F2: -9223372036854775808}
            v28 := []int{(fi22 + (fi22 << ((42) & 63))), 240, (^(fi22))}
            _ = v27
            _ = v28
        }
        _ = fi22
    }
    if true {
        v29 := GenS3{F0: ("foo bar" + ("a" + "a")), F1: 2147483647}
        v30 := GenS2{F0: "hello", F1: 65535, F2: -362}
        var v31 int8 = 127
        _ = v29
        _ = v30
        _ = v31
    }
    var v32 uint8 = 127
    var v33 bool = (2 >= 1)
    switch 1 {
    case 428:
        var v34 int8 = int8(v32)
        _ = v34
    case 32768:
        fmt.Println(v32, v33)
    case 127:
        var v35 float32 = 3.14159
        _ = v35
    default:
        for fi36 := 0; fi36 < 15; fi36++ {
            for fi37 := 0; fi37 < 15; fi37++ {
                v5 = 889
                _ = fi37
            }
            _ = fi36
        }
    }
    for fi38 := 0; fi38 < 15; fi38++ {
        switch 671 {
        case 32767:
            fmt.Println(v3, v6, v5)
        case 698:
            fmt.Println(v3, v32)
        case 47:
            v2 += (v2 | v2)
        }
        _ = fi38
    }
    switch "foo bar" {
    case "hello":
        v33 = (!(!((^(v6)) >= v6)))
    default:
        v1 = v1
    }
    if true {
        v39 := []int{1}
        var v40 int32 = v2
        _ = v39
        _ = v40
    } else {
        v1 = float32(0.0)
    }
    v41 := GenS2{F0: "", F1: 623, F2: -834}
    if ((("a" + ("" + "hello")) > "hello") || ((7 < 4294967295) || (-9223372036854775808 <= -237))) {
        v4[v3] = v33
    }
    switch -930 {
    case -2147483648:
        fmt.Println(v6)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
