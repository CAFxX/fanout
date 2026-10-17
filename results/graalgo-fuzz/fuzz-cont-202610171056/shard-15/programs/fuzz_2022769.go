package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 string; F2 uint32 }
type GenS6 struct { F0 int64; F1 uint8; F2 string }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint; F1 bool; F2 int }

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 761
    var v3 uint = 2
    var v4 bool = ((v2 + 65535) > v2)
    _ = len(v1)
    v2 = v2
    v1[0] = ("a" + "")
    v3 += (v3 & (v3 / ((962) | 1)))
    if true {
        for fi5 := 0; fi5 < 1; fi5++ {
            var v6 int32 = 0
            if v4 {
                v7 := GenS1{F0: (v2 >> (((v3 * v3)) & 63))}
                fmt.Println(v3, v6, v6)
                v8 := []int{}
                _ = v7
                _ = v8
            }
            for fi9 := 0; fi9 < 5; fi9++ {
                var v10 float32 = 1e10
                if (v4 && v4) {
                    v11 := GenS2{F0: 42, F1: v6}
                    v6 -= v6
                    v6 = v6
                    _ = v11
                } else {
                    v12 := []int{1, 2147483647, (^(fi9))}
                    var v13 int = fi5
                    fmt.Println(fi9, v3)
                    _ = v12
                    _ = v13
                }
                fmt.Println(v4)
                _ = fi9
                _ = v10
            }
            _ = fi5
            _ = v6
        }
        fmt.Println(v3)
        switch "\n\t" {
        case "hello":
            switch 527 {
            case -9223372036854775808:
                fmt.Println(v3, v2)
            case 0:
                v14 := []int{0, 9223372036854775807}
                _ = v14
            default:
                delete(v1, 4294967295)
            }
        case "\n\t":
            _ = len(v1)
        }
    } else {
        switch "" {
        case "hello":
            var v15 uint16 = v2
            _ = v15
        case "\n\t":
            v4 = v4
        case "foo bar":
            v16 := GenS3{F0: 0, F1: 2}
            _ = v16
        }
    }
    v17 := GenS4{F0: "a"}
    if (0 > 977) {
        v1[-573] = ("a" + "a")
    } else {
        if (v4 && v4) {
            var v18 uint8 = 1
            _ = v18
        }
    }
    v19 := []int{431, 32768, 495}
    fmt.Println(v3, v4)
    for fi20 := 0; fi20 < 9; fi20++ {
        v19 = append(v19, (fi20 ^ fi20))
        var v21 float64 = (((float64(-2.25) + 3.14159) + 3.14159) * (0.0 / -2.25))
        switch fi20 {
        case 127:
            if v4 {
                var v22 bool = (((v3 / ((42) | 1)) != (^(v3))) && false)
                v22 = v22
                var v23 uint = (v3 >> (((v3 ^ (^(v3)))) & 63))
                _ = v22
                _ = v23
            } else {
                for fi24 := 0; fi24 < 16; fi24++ {
                    var v25 uint16 = (v2 & 263)
                    _ = fi24
                    _ = v25
                }
                if (1 < 2) {
                    _ = len(v1)
                    switch "" {
                    case "":
                        fmt.Println(fi20)
                    case "\n\t":
                        v26 := GenS5{F0: 0, F1: "a", F2: 4294967295}
                        _ = v26
                    case "hello":
                        v27 := GenS6{F0: 826, F1: 127, F2: ((("" + ("foo bar" + "\n\t")) + "hello") + "")}
                        _ = v27
                    default:
                        var v28 uint16 = 728
                        _ = v28
                    }
                } else {
                    v29 := GenS7{F0: 1}
                    v30 := GenS8{F0: 192}
                    var v31 string = "foo bar"
                    _ = v29
                    _ = v30
                    _ = v31
                }
                for fi32 := 0; fi32 < 3; fi32++ {
                    fmt.Println(fi32, fi20)
                    var v33 int = 2147483647
                    _ = fi32
                    _ = v33
                }
            }
        case 859:
            if (fi20 != 9223372036854775807) {
                v3 += (^(v3))
            }
        case 2147483647:
            v34 := GenS2{F0: 2, F1: -1}
            _ = v34
        }
        _ = fi20
        _ = v21
    }
    v35 := GenS9{F0: (v3 | (v3 / ((v3) | 1))), F1: false, F2: 0}
    v1[-702] = ("" + "foo bar")
    for fi36 := 0; fi36 < 12; fi36++ {
        var v37 bool = (132 >= 42)
        _ = fi36
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
