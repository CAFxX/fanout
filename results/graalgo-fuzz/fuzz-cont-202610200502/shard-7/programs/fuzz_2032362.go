package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint16; F2 uint16 }
func genFn0(p0 bool, p1 uint16, p2 bool) int16 {
    v7 := GenS1{F0: -2147483648, F1: 1, F2: uint64(p1)}
    var v8 int = 32768
    for fi9 := 0; fi9 < 3; fi9++ {
        fmt.Println(fi9, p1)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 int64 = 93
    v3 := []int{-1, 0}
    var v4 uint32 = 65535
    var v5 int32 = -2147483648
    var v6 float32 = (1.5 / 1.5)
    var v10 uint32 = v4
    for fi11 := 0; fi11 < 12; fi11++ {
        fmt.Println(v6)
        v3 = append(v3, -9223372036854775808)
        _ = fi11
    }
    v12 := GenS2{F0: ("a" + "\n\t"), F1: uint(v4)}
    switch ("hello" + ("" + "\n\t")) {
    case "a":
        fmt.Println(v2)
    default:
        for fi13 := 0; fi13 < 18; fi13++ {
            if (42 <= -1) {
                if true {
                    var v14 uint8 = 255
                    v2 = (v2 << ((uint(v4)) & 63))
                    _ = v14
                }
            }
            v6 = (v6 + float32(3.14159))
            _ = fi13
        }
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        v16 := GenS1{F0: (v5 >> ((42) & 63)), F1: 42, F2: 18446744073709551615}
        v2 -= (v2 / ((1) | 1))
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        for fi18 := 0; fi18 < 8; fi18++ {
            v19 := GenS3{F0: 55}
            v20 := GenS4{F0: 1, F1: 65535, F2: 32767}
            v5 += (v5 - v5)
            _ = fi18
            _ = v19
            _ = v20
        }
        v21 := GenS3{F0: 0}
        _ = fi17
        _ = v21
    }
    switch 622 {
    case 127:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case -9223372036854775808:
        switch "" {
        case "a":
            fmt.Println(v5, v5, v5)
        case "\n\t":
            v22 := GenS1{F0: -1, F1: 912, F2: 918}
            _ = v22
        case "foo bar":
            var v23 int = 0
            _ = v23
        default:
            var v24 string = ("" + "\n\t")
            _ = v24
        }
    case 2147483647:
        var v25 uint64 = 18446744073709551615
        _ = v25
    default:
        v4 = v4
    }
    v5 = v5
    switch ("foo bar" + "hello") {
    case "a":
        var v26 uint64 = 213
        _ = v26
    case "foo bar":
        var v27 float32 = 3.14159
        _ = v27
    case "\n\t":
        v5 -= -1
    }
    var v28 float64 = 1.5
    for fi29 := 0; fi29 < 10; fi29++ {
        var v30 int32 = (v5 | v5)
        var v31 float64 = (v28 * ((3.14159 / float64(float32((v28 * v28)))) * (v28 + v28)))
        _ = fi29
        _ = v30
        _ = v31
    }
    switch ("" + "hello") {
    case "\n\t":
        for fi32 := 0; fi32 < 5; fi32++ {
            var v33 string = ""
            if (v6 < (1.5 / (3.14159 + (((float32(float64(float32(((v28 + float64(-2.25)) / v28)))) / (1.5 / 1.5)) * 1e10) - v6)))) {
                for fi34 := 0; fi34 < 7; fi34++ {
                    _ = len(v1)
                    var v35 float64 = v28
                    _ = fi34
                    _ = v35
                }
            }
            if (!false) {
                v1[42] = 772
                fmt.Println(v6, fi32)
            }
            _ = fi32
            _ = v33
        }
    case "hello":
        for fi36 := 0; fi36 < 15; fi36++ {
            for fi37 := 0; fi37 < 13; fi37++ {
                v10 -= (v10 % ((v10) | 1))
                v6 -= v6
                v5 = v5
                _ = fi37
            }
            v10 -= v10
            for fi38 := 0; fi38 < 4; fi38++ {
                v2 -= 934
                if false {
                    var v39 uint8 = 1
                    _ = v39
                } else {
                    var v40 uint64 = 2
                    fmt.Println(v4)
                    v41 := GenS2{F0: "foo bar", F1: 0}
                    _ = v40
                    _ = v41
                }
                _ = fi38
            }
            _ = fi36
        }
    case "a":
        delete(v1, 40)
    }
    for fi42 := 0; fi42 < 9; fi42++ {
        fmt.Println(v4, v2, v2)
        var v43 float64 = v28
        switch (("a" + "hello") + "") {
        case "hello":
            v10 = (v4 * 2147483647)
        case "\n\t":
            if len(v3) > 0 {
                v3[0] = 142
            }
        case "foo bar":
            var v44 uint64 = 2
            _ = v44
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi42
        _ = v43
    }
    v45 := GenS1{F0: (v5 / ((v5) | 1)), F1: 18446744073709551615, F2: 0}
    var v46 int64 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
