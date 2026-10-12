package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int16; F2 int }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 string; F1 int; F2 uint8 }
func genFn0(p0 int16, p1 int) int32 {
    v5 := []int{0, p1, p1}
    v6 := GenS1{F0: (^(p0))}
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 int64 = 42
        var v9 int16 = (p0 / (((p0 >> ((65535) & 63))) | 1))
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -830
}
func genFn1() uint8 {
    var v10 float64 = 1e10
    if false {
        var v11 int32 = -2147483648
        var v12 int = -2147483648
        if false {
            v10 += (v10 - 0.0)
        }
        _ = v11
        _ = v12
    } else {
        v13 := []int{9223372036854775807, 32768, -64}
        fmt.Println(v10)
        _ = v13
    }
    v14 := GenS1{F0: 592}
    _ = v10
    _ = v14
    return 255
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]int)
    var v3 string = "hello"
    var v4 int16 = 704
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 uint = 896
        var v17 bool = true
        switch v3 {
        case "\n\t":
            fmt.Println(v16, v4)
        case "":
            if (0 <= 1) {
                v3 = "foo bar"
                v17 = ((-2.25 * (((-2.25 / 1.5) / 3.14159) * 1e10)) >= -2.25)
                v18 := GenS1{F0: v4}
                _ = v18
            }
        case "foo bar":
            v17 = (65535 <= 32767)
        default:
            v17 = (v17 && (255 > 128))
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    v4 = 0
    v19 := GenS1{F0: (v4 | v4)}
    if (105 <= -109) {
        delete(v2, 1)
        fmt.Println(v3, v4)
    } else {
        fmt.Println(v4)
        fmt.Println(v3, v3)
    }
    if true {
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 uint = 65535
            if true {
                v3 = (v3 + ("" + (v3 + ((v3 + v3) + "hello"))))
            } else {
                var v22 uint64 = 4294967295
                if (true && false) {
                    v4 += (v4 >> (((^(v21))) & 63))
                }
                fmt.Println(v3, v4, fi20)
                _ = v22
            }
            _ = fi20
            _ = v21
        }
        var v23 int = -2147483648
        var v24 uint64 = 18446744073709551615
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 19; fi25++ {
        v4 = (v4 + (^(v4)))
        for fi26 := 0; fi26 < 10; fi26++ {
            if (v4 > 594) {
                v1[(v3 + v3)] = v3
            }
            fmt.Println(fi25, v3, v4)
            v27 := GenS2{F0: fi26, F1: (-(v4)), F2: fi25}
            _ = fi26
            _ = v27
        }
        for fi28 := 0; fi28 < 6; fi28++ {
            if false {
                v3 = v3
                v29 := GenS3{F0: 32767, F1: ""}
                _ = v29
            } else {
                v30 := GenS3{F0: 2, F1: v3}
                fmt.Println(fi25, fi25, v4)
                v31 := GenS1{F0: v4}
                v32 := GenS2{F0: (fi25 ^ (^(fi25))), F1: (v4 ^ 691), F2: (-(fi25))}
                _ = v30
                _ = v31
                _ = v32
            }
            fmt.Println(fi25, fi25)
            var v33 uint = 42
            _ = fi28
            _ = v33
        }
        _ = fi25
    }
    for fi34 := 0; fi34 < 13; fi34++ {
        v3 = "\n\t"
        if ((fi34 * 2147483647) < (fi34 * (fi34 % (((fi34 - fi34)) | 1)))) {
            _ = len(v1)
            v35 := []int{fi34, (fi34 << ((895) & 63)), (fi34 >> ((42) & 63)), (fi34 / (((fi34 % ((4294967295) | 1))) | 1))}
            _ = v35
        } else {
            switch v3 {
            case "foo bar":
                v4 -= 32767
            default:
                var v36 uint8 = 255
                _ = v36
            }
        }
        _ = fi34
    }
    v37 := GenS2{F0: 32767, F1: (v4 >> ((217) & 63)), F2: 32768}
    var v38 int8 = 0
    v39 := GenS2{F0: -9223372036854775808, F1: (v4 << ((0) & 63)), F2: 127}
    for fi40 := 0; fi40 < 15; fi40++ {
        switch (v3 + (v3 + "hello")) {
        case "foo bar":
            for fi41 := 0; fi41 < 2; fi41++ {
                fmt.Println(v3, v38, v38)
                v38 = (^(v38))
                v38 -= (v38 & int8(fi40))
                _ = fi41
            }
        case "\n\t":
            var v42 float32 = (-2.25 * (3.14159 - ((((((3.14159 / 1e10) / 3.14159) - -2.25) + ((((0.1 / 1.5) / 1.0) + -2.25) / 3.14159)) + -2.25) / 1.0)))
            _ = v42
        }
        v43 := GenS3{F0: 2, F1: "foo bar"}
        var v44 uint32 = 2
        _ = fi40
        _ = v43
        _ = v44
    }
    for fi45 := 0; fi45 < 14; fi45++ {
        v46 := GenS4{F0: "\n\t", F1: (fi45 * fi45), F2: 2}
        var v47 int16 = v4
        _ = fi45
        _ = v46
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
