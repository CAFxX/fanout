package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 uint8; F2 int32 }
type GenS5 struct { F0 int; F1 uint; F2 int16 }
func genFn0(p0 uint64, p1 int16) bool {
    if false {
        var v7 uint16 = 861
        switch 32767 {
        case 0:
            if (false || (!(!false))) {
                var v8 float32 = 1e10
                _ = v8
            }
        case 2147483647:
            fmt.Println(p1, p0)
        }
        var v9 int32 = -1
        _ = v7
        _ = v9
    } else {
        for fi10 := 0; fi10 < 6; fi10++ {
            p0 -= p0
            p0 = 0
            _ = fi10
        }
        p0 = (p0 & 1)
    }
    switch 32768 {
    case 339:
        var v11 string = "\n\t"
        _ = v11
    }
    switch ("" + "a") {
    case "foo bar":
        if ("" > "a") {
            var v12 string = ((("" + ("a" + "\n\t")) + "hello") + ("" + ("hello" + (("a" + ("a" + "hello")) + "\n\t"))))
            _ = v12
        } else {
            v13 := GenS1{F0: p0, F1: ((p1 == p1) || (((((1e10 - float64(0.1)) / 3.14159) != (((((float64((-2.25 - 0.1)) * 1.5) + 1e10) - float64((((1.5 - float32(-2.25)) * 3.14159) + -2.25))) + (1e10 / (3.14159 - 0.1))) + 0.1)) && (65535 == 80)) && false))}
            _ = v13
        }
    default:
        v14 := []int{}
        _ = v14
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        if (1 < 127) {
            p1 = p1
            v16 := GenS2{F0: (fi15 | fi15)}
            _ = v16
        }
        if false {
            if ((p0 >> ((65535) & 63)) < 1) {
                var v17 float64 = (0.0 * float64(0.0))
                v18 := GenS3{F0: (fi15 << ((186) & 63))}
                p0 = (p0 | 42)
                _ = v17
                _ = v18
            }
        } else {
            p1 += (p1 << ((65535) & 63))
            switch fi15 {
            case -2147483648:
                var v19 uint16 = 787
                _ = v19
            default:
                p1 = (-(p1))
            }
        }
        _ = fi15
    }
    _ = p0
    _ = p1
    return true
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{}
    var v3 int8 = 1
    var v4 uint = 363
    var v5 int32 = 72
    var v6 int = 2147483647
    var v20 int8 = 42
    var v21 uint32 = 2147483648
    v5 += v5
    fmt.Println(v21, v6)
    switch 4294967295 {
    case -2147483648:
        var v22 string = "a"
        _ = v22
    default:
        v4 += 2
    }
    switch (^(v6)) {
    case 2147483647:
        v21 += v21
    case 4294967295:
        var v23 string = ("hello" + "a")
        _ = v23
    case 160:
        v20 -= -1
    }
    fmt.Println(v3, v21)
    if (!(65535 >= v4)) {
        fmt.Println(v5)
        switch ("a" + ("hello" + ("foo bar" + ("" + "a")))) {
        case "foo bar":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            var v24 float32 = float32(0.0)
            _ = v24
        }
        v4 -= (v4 + v4)
    } else {
        var v25 float32 = 1e10
        if false {
            var v26 int = v6
            _ = v26
        } else {
            v3 += v3
            for fi27 := 0; fi27 < 1; fi27++ {
                var v28 uint8 = 127
                v25 += float32(-2.25)
                fmt.Println(v6, v6)
                _ = fi27
                _ = v28
            }
            switch "\n\t" {
            case "foo bar":
                var v29 uint32 = v21
                _ = v29
            case "":
                fmt.Println(v6, v21)
            case "hello":
                v30 := GenS2{F0: 9223372036854775807}
                _ = v30
            default:
                if true {
                    v6 = 4294967295
                }
            }
        }
        var v31 float64 = (1.5 - 0.0)
        _ = v25
        _ = v31
    }
    v32 := GenS3{F0: (v6 >> ((v4) & 63))}
    var v33 bool = (false && false)
    v33 = (18446744073709551615 > 907)
    v34 := GenS4{F0: 32767, F1: 232, F2: -827}
    var v35 uint32 = 597
    var v36 uint8 = 159
    for fi37 := 0; fi37 < 4; fi37++ {
        v38 := GenS5{F0: int(v35), F1: (v4 ^ 18446744073709551615), F2: 593}
        var v39 uint8 = (^(v36))
        _ = fi37
        _ = v38
        _ = v39
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v36)
    var v40 uint16 = 752
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
