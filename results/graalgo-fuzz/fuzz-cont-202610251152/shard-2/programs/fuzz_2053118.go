package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 int8; F2 int }
type GenS6 struct { F0 int64 }
func genFn0(p0 int16, p1 int64) uint64 {
    v4 := GenS1{F0: p1, F1: (p0 << ((586) & 63)), F2: 4294967295}
    p1 = p1
    for fi5 := 0; fi5 < 8; fi5++ {
        p0 -= -663
        v6 := GenS2{F0: -1, F1: 2}
        for fi7 := 0; fi7 < 3; fi7++ {
            fmt.Println(fi5)
            if (false || true) {
                var v8 int = 1
                var v9 bool = (true || (p0 == (p0 + -60)))
                _ = v8
                _ = v9
            } else {
                var v10 int64 = (p1 | p1)
                v11 := GenS1{F0: 1, F1: p0, F2: 581}
                var v12 uint = 2
                var v13 int32 = 1
                v14 := GenS1{F0: 9223372036854775807, F1: (p0 & (p0 | p0)), F2: 228}
                _ = v10
                _ = v11
                _ = v12
                _ = v13
                _ = v14
            }
            if false {
                var v15 float32 = (0.0 * ((((float32(((((-2.25 * 1.5) + 3.14159) + 3.14159) * 1e10)) + 3.14159) * ((((1e10 / 1e10) / (0.1 * (1e10 + (1e10 * -2.25)))) * -2.25) + 1e10)) / 1.0) - 0.0))
                if false {
                    var v16 uint16 = 1
                    _ = v16
                }
                var v17 int64 = (p1 & 72)
                v18 := GenS3{F0: 42}
                _ = v15
                _ = v17
                _ = v18
            }
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    var v19 uint = 65535
    _ = p0
    _ = p1
    _ = v19
    _ = v4
    return 42
}

func main() {
    v1 := []string{}
    var v2 float64 = float64((1.5 / 0.1))
    var v3 int8 = -1
    for fi20 := 0; fi20 < 2; fi20++ {
        fmt.Println(v2)
        _ = fi20
    }
    fmt.Println(v3, v2)
    switch ("foo bar" + "") {
    case "hello":
        v1 = append(v1, "a")
    }
    v2 -= v2
    if (1 > 32767) {
        if (!(0 == 4294967295)) {
            var v21 uint16 = 1
            var v22 int8 = 42
            _ = v21
            _ = v22
        }
        var v23 int = 9223372036854775807
        v24 := GenS1{F0: 0, F1: 42, F2: 532}
        _ = v23
        _ = v24
    } else {
        v2 -= float64(0.0)
    }
    v25 := GenS4{F0: "hello"}
    v26 := GenS5{F0: 42, F1: (v3 >> ((65535) & 63)), F2: -2147483648}
    var v27 int32 = -2147483648
    if len(v1) > 0 {
        v1[0] = "a"
    }
    v28 := GenS2{F0: (v27 >> ((255) & 63)), F1: 2147483647}
    v29 := GenS6{F0: 9223372036854775807}
    var v30 int16 = -32768
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    switch ("a" + ("" + "a")) {
    case "":
        var v31 uint16 = 255
        _ = v31
    case "a":
        if (("" + "hello") != "hello") {
            v27 += -1
            if len(v1) > 0 {
                v1[0] = ("\n\t" + ("hello" + "hello"))
            }
            var v32 int64 = 9223372036854775807
            _ = v32
        }
    }
    if ((1 != 18446744073709551615) && (!true)) {
        var v33 bool = (4294967295 == 2147483647)
        fmt.Println(v3, v2)
        _ = v33
    } else {
        switch 32768 {
        case 2147483647:
            var v34 uint8 = 127
            _ = v34
        case -9223372036854775808:
            var v35 int16 = v30
            _ = v35
        default:
            var v36 int64 = 1
            _ = v36
        }
        v30 += v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
