package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 uint32; F2 int }
type GenS6 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS7 struct { F0 string; F1 int16 }
func genFn0(p0 int64) bool {
    var v4 int16 = -804
    switch "a" {
    case "":
        p0 += 869
    case "foo bar":
        v5 := GenS1{F0: (p0 >> ((18446744073709551615) & 63))}
        _ = v5
    case "hello":
        for fi6 := 0; fi6 < 3; fi6++ {
            v4 -= (v4 << ((2) & 63))
            for fi7 := 0; fi7 < 4; fi7++ {
                v4 += v4
                fmt.Println(fi7, fi6)
                _ = fi7
            }
            _ = fi6
        }
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        if (1 != 1) {
            switch fi8 {
            case -9223372036854775808:
                v9 := GenS2{F0: 1000}
                _ = v9
            case 42:
                switch "\n\t" {
                case "":
                    v10 := []int{}
                    _ = v10
                case "a":
                    var v11 int16 = v4
                    _ = v11
                case "\n\t":
                    for fi12 := 0; fi12 < 15; fi12++ {
                        v4 -= -1
                        var v13 bool = ("\n\t" < ("foo bar" + "foo bar"))
                        fmt.Println(fi8)
                        _ = fi12
                        _ = v13
                    }
                }
            case 828:
                var v14 float32 = 3.14159
                _ = v14
            }
            v15 := GenS2{F0: (v4 + v4)}
            v16 := GenS3{F0: 4294967295, F1: 2}
            _ = v15
            _ = v16
        } else {
            var v17 bool = (v4 == 1)
            _ = v17
        }
        _ = fi8
    }
    v4 += -32768
    _ = p0
    _ = v4
    return (0.0 <= 1e10)
}

func main() {
    v1 := []int{-1, 1, 0}
    var v2 float64 = 0.1
    v3 := []int{1, 381, 991}
    v18 := GenS4{F0: 2}
    v1 = append(v1, -2147483648)
    v3 = append(v3, 2147483647)
    switch 654 {
    case 134:
        for fi19 := 0; fi19 < 7; fi19++ {
            v20 := GenS5{F0: 783, F1: 1, F2: fi19}
            if (!false) {
                fmt.Println(v2, fi19)
                v21 := GenS6{F0: 128, F1: 1, F2: 2}
                _ = v21
            }
            if len(v3) > 0 {
                v3[0] = (fi19 % ((fi19) | 1))
            }
            _ = fi19
            _ = v20
        }
    case 2147483647:
        switch "" {
        case "hello":
            fmt.Println(v2)
        case "a":
            v3 = append(v3, 368)
        case "\n\t":
            var v22 uint64 = 2
            _ = v22
        default:
            v1 = append(v1, 4294967295)
        }
    }
    switch 1 {
    case 9223372036854775807:
        var v23 string = "hello"
        _ = v23
    case -2147483648:
        v3 = append(v3, 779)
    default:
        var v24 uint16 = 2
        _ = v24
    }
    v2 = (v2 - v2)
    v25 := GenS2{F0: 1000}
    fmt.Println(v2)
    var v26 int64 = 1
    if (!(838 <= 1)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        v2 = 1.5
        _ = fi27
    }
    switch "" {
    case "":
        var v28 float64 = v2
        _ = v28
    case "hello":
        v26 += -613
    case "foo bar":
        if false {
            v2 = v2
            if true {
                var v29 uint32 = 2147483648
                v26 = (v26 * (-(v26)))
                v30 := GenS1{F0: (v26 << ((535) & 63))}
                _ = v29
                _ = v30
            } else {
                v31 := GenS7{F0: ("a" + ("hello" + "hello")), F1: 42}
                var v32 int16 = -791
                _ = v31
                _ = v32
            }
            v26 += 9223372036854775807
        } else {
            switch ("a" + "a") {
            case "":
                var v33 string = "a"
                _ = v33
            default:
                for fi34 := 0; fi34 < 8; fi34++ {
                    var v35 int16 = 392
                    _ = fi34
                    _ = v35
                }
            }
        }
    default:
        fmt.Println(v2, v26)
    }
    var v36 int64 = (v26 / ((-966) | 1))
    var v37 uint32 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v36
    _ = v37
}
