package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS4 struct { F0 int16; F1 int8; F2 uint8 }
func genFn0(p0 uint16, p1 uint64) string {
    if (-2.25 <= 0.1) {
        var v7 int = 4294967295
        v7 -= -2147483648
        _ = v7
    } else {
        v8 := GenS1{F0: 1, F1: ""}
        var v9 uint = 4294967295
        _ = v8
        _ = v9
    }
    var v10 int = 4294967295
    _ = p0
    _ = p1
    _ = v10
    return ""
}
func genFn1(p0 uint16, p1 int64) uint64 {
    var v11 float32 = float32(0.1)
    p1 += p1
    v12 := []int{42, 127, -2147483648, 127}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 1
}

func main() {
    var v1 int64 = 588
    v2 := []int{42, -9223372036854775808, 206}
    var v3 uint64 = 4294967295
    var v4 uint = 18446744073709551615
    var v5 int32 = 132
    v6 := make(map[int]int)
    var v13 int16 = -1
    fmt.Println(v4, v13, v1)
    switch (((("" + "hello") + "") + ("foo bar" + (("foo bar" + "foo bar") + ("foo bar" + "foo bar")))) + "\n\t") {
    case "a":
        v13 = (v13 << ((0) & 63))
    }
    v4 = v4
    var v14 int32 = v5
    _ = len(v2)
    v14 += (^(v14))
    var v15 float64 = (0.0 - 1.5)
    if ((((0.1 * (3.14159 + 1e10)) - 0.1) + float32(((float64((float32(0.0) / 1.5)) + -2.25) / v15))) != 1.5) {
        v6[32767] = 128
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        for fi17 := 0; fi17 < 3; fi17++ {
            var v18 uint32 = 911
            var v19 string = ("" + "\n\t")
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = len(v6)
        _ = fi16
    }
    v4 = 1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        for fi21 := 0; fi21 < 7; fi21++ {
            if (v13 >= v13) {
                fmt.Println(v13)
                var v22 uint8 = 2
                _ = v22
            }
            _ = fi21
        }
        switch fi20 {
        case 42:
            fmt.Println(v14, v14)
        case 0:
            switch (fi20 + (fi20 / ((fi20) | 1))) {
            case 42:
                for fi23 := 0; fi23 < 10; fi23++ {
                    var v24 uint = 65535
                    fmt.Println(v14)
                    _ = fi23
                    _ = v24
                }
            }
        default:
            for fi25 := 0; fi25 < 15; fi25++ {
                if true {
                    fmt.Println(v14, v15, v1)
                    var v26 float64 = v15
                    v15 = -2.25
                    _ = v26
                }
                v27 := GenS1{F0: (v13 * (v13 + (^(v13)))), F1: ""}
                var v28 float32 = 1e10
                _ = fi25
                _ = v27
                _ = v28
            }
        }
        _ = fi20
    }
    switch ("" + "") {
    case "\n\t":
        for fi29 := 0; fi29 < 9; fi29++ {
            var v30 uint32 = 0
            _ = fi29
            _ = v30
        }
    default:
        v31 := GenS2{F0: v1, F1: 1}
        _ = v31
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        var v33 int16 = 42
        switch "foo bar" {
        case "foo bar":
            if (((fi32 << (((^(v4))) & 63)) <= 128) || (2147483647 > 0)) {
                for fi34 := 0; fi34 < 10; fi34++ {
                    var v35 string = "hello"
                    v14 = v14
                    var v36 string = "\n\t"
                    fmt.Println(v4, v5)
                    v37 := GenS3{F0: 2, F1: 1, F2: v14}
                    v3 = (v3 * genFn1(32767, (v1 >> ((v4) & 63))))
                    _ = fi34
                    _ = v35
                    _ = v36
                    _ = v37
                }
                var v38 uint8 = 7
                _ = v38
            } else {
                var v39 int = int(v1)
                _ = v39
            }
        case "a":
            fmt.Println(v15)
        default:
            for fi40 := 0; fi40 < 10; fi40++ {
                _ = len(v6)
                v4 -= (v4 % ((2) | 1))
                v41 := GenS4{F0: int16(v4), F1: 127, F2: 128}
                _ = fi40
                _ = v41
            }
        }
        var v42 uint = (v4 + 18446744073709551615)
        _ = fi32
        _ = v33
        _ = v42
    }
    var v43 uint8 = 42
    var v44 string = ""
    v15 = (3.14159 * float64(1e10))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
