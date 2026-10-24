package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 uint32; F2 int }
type GenS5 struct { F0 string; F1 int8; F2 uint32 }
func genFn0(p0 uint16) string {
    p0 = (p0 ^ p0)
    switch 548 {
    case -889:
        v7 := GenS1{F0: 42}
        _ = v7
    default:
        fmt.Println(p0)
    }
    _ = p0
    return "hello"
}
func genFn1(p0 int16, p1 int16, p2 int) uint64 {
    fmt.Println(p2)
    p1 = p0
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint = 0
    var v2 int64 = 9223372036854775807
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]int)
    v5 := make(map[int]bool)
    var v6 float64 = (float64(0.1) * (1.5 - ((1e10 - (1.5 * 3.14159)) - -2.25)))
    v1 += 2
    v2 = (v2 - (v2 << ((825) & 63)))
    if (65535 == 2147483648) {
        switch 2147483647 {
        case -2147483648:
            v8 := GenS1{F0: genFn1(-32768, 1000, -202)}
            _ = v8
        default:
            v9 := GenS1{F0: 1}
            _ = v9
        }
        var v10 uint16 = 571
        var v11 uint16 = v10
        _ = v10
        _ = v11
    }
    switch ("foo bar" + "\n\t") {
    case "":
        v1 = v1
    case "foo bar":
        delete(v5, -2147483648)
    default:
        _ = len(v4)
    }
    var v12 float64 = 0.1
    v13 := GenS2{F0: "foo bar"}
    for fi14 := 0; fi14 < 5; fi14++ {
        if false {
            v15 := GenS3{F0: 2147483647}
            v16 := []int{(fi14 ^ (fi14 >> (((v1 + 1)) & 63))), fi14}
            _ = len(v4)
            _ = v15
            _ = v16
        } else {
            var v17 int = fi14
            v18 := GenS3{F0: -1}
            v2 = -1
            _ = v17
            _ = v18
        }
        var v19 int8 = 67
        v2 = v2
        _ = fi14
        _ = v19
    }
    v20 := GenS3{F0: 179}
    if (127 > 127) {
        if true {
            var v21 bool = false
            _ = v21
        } else {
            switch "a" {
            case "":
                fmt.Println(v2, v6)
            case "a":
                var v22 float64 = 1e10
                _ = v22
            case "\n\t":
                var v23 uint8 = uint8(v1)
                _ = v23
            default:
                v12 = v6
            }
            for fi24 := 0; fi24 < 7; fi24++ {
                var v25 int = int(v2)
                var v26 uint64 = 853
                _ = fi24
                _ = v25
                _ = v26
            }
        }
        fmt.Println(v6, v6)
    }
    var v27 uint = v1
    if (v2 != 1) {
        if true {
            if true {
                v2 -= v2
                v28 := GenS1{F0: 0}
                _ = v28
            } else {
                fmt.Println(v6)
                var v29 int8 = 92
                _ = v29
            }
            switch "\n\t" {
            case "\n\t":
                switch 32767 {
                case 4294967295:
                    var v30 string = "hello"
                    _ = v30
                }
            case "foo bar":
                v31 := GenS4{F0: "hello", F1: 4294967295, F2: -2147483648}
                _ = v31
            case "hello":
                v6 = v6
            }
        }
        v1 += (v1 / ((v1) | 1))
        for fi32 := 0; fi32 < 12; fi32++ {
            var v33 uint16 = 658
            v6 = v6
            for fi34 := 0; fi34 < 14; fi34++ {
                var v35 uint64 = 1
                switch ("foo bar" + ("a" + "a")) {
                case "a":
                    v33 = 54
                default:
                    if true {
                        v36 := []int{332, fi34, (fi32 + (fi34 + fi34))}
                        if len(v36) >= 2 {
                            v36 = v36[1:len(v36)]
                        }
                        _ = v36
                    } else {
                        v37 := []int{fi34, 127, fi34, 127}
                        if (0 >= 9223372036854775807) {
                            v6 -= v6
                            v12 += ((0.0 * 1e10) / 3.14159)
                        } else {
                            var v38 uint32 = 581
                            var v39 uint8 = 128
                            _ = v38
                            _ = v39
                        }
                        var v40 float64 = float64(1.5)
                        _ = v37
                        _ = v40
                    }
                }
                var v41 float64 = (((float64(((-2.25 * (-2.25 * (0.0 / 1.0))) / 1.0)) * v12) * (v6 - (0.0 * float64((1.5 - (0.0 + 1e10)))))) + float64(3.14159))
                _ = fi34
                _ = v35
                _ = v41
            }
            _ = fi32
            _ = v33
        }
    }
    v42 := GenS2{F0: ("foo bar" + "foo bar")}
    v43 := GenS4{F0: ("hello" + ""), F1: 921, F2: -792}
    v6 -= v12
    v3[206] = 210
    v6 -= v12
    switch "foo bar" {
    case "\n\t":
        v44 := []int{127, 128, 32767}
        _ = v44
    case "foo bar":
        var v45 int8 = 1
        _ = v45
    default:
        v46 := GenS5{F0: "a", F1: 127, F2: 4294967295}
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
