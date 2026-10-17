package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS4 struct { F0 int }
func genFn0(p0 bool, p1 int8) int16 {
    for fi7 := 0; fi7 < 7; fi7++ {
        p0 = p0
        _ = fi7
    }
    switch 32768 {
    case 4294967295:
        v8 := GenS1{F0: 42}
        _ = v8
    case -2147483648:
        for fi9 := 0; fi9 < 7; fi9++ {
            fmt.Println(p1)
            switch "hello" {
            case "\n\t":
                var v10 float64 = float64((float32(3.14159) - (((((-2.25 * 3.14159) * (-2.25 + float32(3.14159))) + 1e10) - 1e10) + 3.14159)))
                _ = v10
            default:
                var v11 uint8 = 0
                _ = v11
            }
            _ = fi9
        }
    default:
        v12 := []int{1, -9223372036854775808}
        _ = v12
    }
    var v13 int16 = -1
    v14 := []int{}
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return (v13 | 32767)
}

func main() {
    v1 := []string{"hello", "hello", (("\n\t" + (("a" + (("" + "foo bar") + ("a" + "foo bar"))) + "foo bar")) + ("a" + ""))}
    var v2 int16 = 128
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    v5 := []uint32{}
    v6 := []string{"hello"}
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 uint = 0
        var v17 string = (("a" + "foo bar") + "")
        _ = fi15
        _ = v16
        _ = v17
    }
    v2 += v2
    if false {
        var v18 int32 = 785
        fmt.Println(v18, v18)
        for fi19 := 0; fi19 < 19; fi19++ {
            for fi20 := 0; fi20 < 4; fi20++ {
                switch "a" {
                case "a":
                    v21 := GenS1{F0: (^(v18))}
                    _ = v21
                case "":
                    if (((("\n\t" + ("hello" + "")) + "a") + (("" + ("a" + ("\n\t" + ("a" + ("foo bar" + "foo bar"))))) + ("" + "hello"))) == "a") {
                        var v22 float32 = 1e10
                        fmt.Println(v18, v2)
                        _ = v22
                    }
                case "hello":
                    v23 := GenS2{F0: 0}
                    _ = v23
                }
                var v24 uint8 = 1
                _ = fi20
                _ = v24
            }
            switch ("" + "hello") {
            case "\n\t":
                switch 9223372036854775807 {
                case 712:
                    var v25 uint64 = 2
                    _ = v25
                default:
                    v5 = append(v5, 2)
                }
            }
            var v26 int = fi19
            _ = fi19
            _ = v26
        }
        _ = v18
    } else {
        v27 := GenS3{F0: 100, F1: 2147483647, F2: 255}
        _ = v27
    }
    v28 := GenS3{F0: -44, F1: 829, F2: 255}
    var v29 bool = ((!(!false)) || false)
    if v29 {
        v29 = (((!v29) && false) && v29)
        v2 = v2
    }
    if (973 <= 2) {
        for fi30 := 0; fi30 < 5; fi30++ {
            var v31 uint16 = 1
            _ = fi30
            _ = v31
        }
    }
    var v32 uint32 = 624
    delete(v4, 4294967295)
    if (236 > 171) {
        fmt.Println(v32, v2)
        var v33 string = "hello"
        _ = v33
    }
    v2 += (v2 & genFn0(v29, -1))
    if v29 {
        v2 += v2
        v34 := v4[2147483647]
        _ = v34
    } else {
        _ = len(v4)
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        if (42 != 466) {
            var v36 bool = true
            if true {
                var v37 uint = 844
                if v36 {
                    v38 := GenS3{F0: -47, F1: 65535, F2: 735}
                    _ = v38
                }
                switch "a" {
                case "a":
                    switch "\n\t" {
                    case "":
                        _ = len(v4)
                    case "hello":
                        if (911 > 2) {
                            v39 := GenS4{F0: (fi35 ^ fi35)}
                            _ = v39
                        } else {
                            fmt.Println(v36)
                            var v40 uint8 = 128
                            v41 := GenS3{F0: int8(v32), F1: v32, F2: 32767}
                            _ = v40
                            _ = v41
                        }
                    case "a":
                        v29 = (v32 == (v32 % ((v32) | 1)))
                    default:
                        switch fi35 {
                        case -995:
                            v32 -= (v32 | (^(v32)))
                        default:
                            v42 := GenS3{F0: 127, F1: v32, F2: 0}
                            _ = v42
                        }
                    }
                case "\n\t":
                    fmt.Println(v36)
                case "":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = v37
            } else {
                var v43 int8 = 100
                v29 = v29
                var v44 int = (fi35 - (fi35 ^ (fi35 - 678)))
                _ = v43
                _ = v44
            }
            var v45 bool = (1e10 == 1e10)
            _ = v36
            _ = v45
        }
        fmt.Println(v29)
        _ = fi35
    }
    v29 = (1 > 18446744073709551615)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
