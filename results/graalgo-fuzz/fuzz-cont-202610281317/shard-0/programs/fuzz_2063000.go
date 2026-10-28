package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int64; F1 int8; F2 uint }
type GenS4 struct { F0 int8; F1 uint16; F2 int }
func genFn0() int64 {
    v5 := GenS1{F0: 32768}
    var v6 uint64 = 0
    v6 = 18446744073709551615
    v6 -= v6
    fmt.Println(v6)
    _ = v5
    _ = v6
    return -9223372036854775808
}
func genFn1() uint {
    var v7 int32 = -1
    v7 += (v7 + -2147483648)
    var v8 uint32 = 802
    v9 := []int{1, 253, -1, 9223372036854775807}
    switch 1 {
    case 42:
        switch ("" + "a") {
        case "a":
            var v10 float32 = 0.0
            _ = v10
        case "\n\t":
            v11 := GenS2{F0: 2147483647, F1: 9223372036854775807}
            _ = v11
        case "":
            switch (((("\n\t" + "foo bar") + "hello") + "\n\t") + ((("foo bar" + ("a" + "")) + "a") + "foo bar")) {
            case "a":
                v12 := GenS3{F0: -1, F1: 1, F2: 220}
                _ = v12
            }
        }
    case 128:
        switch "hello" {
        case "\n\t":
            var v13 float64 = float64((1.5 + 0.1))
            _ = v13
        case "hello":
            var v14 float32 = 0.1
            _ = v14
        default:
            if len(v9) > 0 {
                v9[0] = -9223372036854775808
            }
        }
    case -2147483648:
        var v15 uint64 = 269
        _ = v15
    default:
        if (!(128 < 0)) {
            switch "foo bar" {
            case "foo bar":
                v7 = (v7 * v7)
            case "hello":
                if (-2.25 != -2.25) {
                    for fi16 := 0; fi16 < 17; fi16++ {
                        fmt.Println(v8, v7)
                        v7 = 101
                        _ = fi16
                    }
                    _ = len(v9)
                } else {
                    var v17 uint16 = 701
                    _ = v17
                }
            default:
                var v18 uint8 = 2
                _ = v18
            }
            var v19 uint16 = 42
            v7 = (v7 - (v7 / ((v7) | 1)))
            _ = v19
        }
    }
    _ = v7
    _ = v8
    _ = v9
    return 828
}

func main() {
    v1 := make(map[int]string)
    var v2 int32 = -1
    v3 := make(map[int]int)
    v4 := []string{"", (("" + "") + ("" + ((("" + "hello") + "\n\t") + ""))), "\n\t", "\n\t"}
    v20 := GenS4{F0: -73, F1: 32767, F2: -1}
    var v21 uint32 = 2147483648
    if (0 == 18446744073709551615) {
        for fi22 := 0; fi22 < 20; fi22++ {
            for fi23 := 0; fi23 < 5; fi23++ {
                var v24 uint = 2
                var v25 float64 = 0.0
                var v26 float64 = (float64(0.0) + float64(1e10))
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
            _ = fi22
        }
        fmt.Println(v2, v2)
        if true {
            if true {
                var v27 int64 = 9223372036854775807
                for fi28 := 0; fi28 < 9; fi28++ {
                    if false {
                        v4 = append(v4, "hello")
                    } else {
                        for fi29 := 0; fi29 < 5; fi29++ {
                            var v30 uint = 255
                            fmt.Println(v27)
                            _ = fi29
                            _ = v30
                        }
                        var v31 int = (fi28 | fi28)
                        _ = v31
                    }
                    fmt.Println(v2, fi28)
                    _ = fi28
                }
                v32 := GenS3{F0: 539, F1: 0, F2: 1}
                _ = v27
                _ = v32
            }
        }
    }
    if (255 <= 2) {
        for fi33 := 0; fi33 < 18; fi33++ {
            switch ("hello" + "foo bar") {
            case "foo bar":
                v2 = v2
            default:
                var v34 uint = 0
                _ = v34
            }
            fmt.Println(fi33)
            _ = fi33
        }
        var v35 uint8 = 127
        _ = v35
    }
    v36 := v3[-541]
    for fi37 := 0; fi37 < 5; fi37++ {
        if ((18446744073709551615 == 1) || true) {
            _ = len(v3)
            v1[v36] = "foo bar"
            _ = len(v1)
        } else {
            if true {
                for fi38 := 0; fi38 < 9; fi38++ {
                    v3[fi37] = -299
                    _ = fi38
                }
                v2 -= -198
                for fi39 := 0; fi39 < 7; fi39++ {
                    var v40 int16 = 0
                    var v41 uint64 = 4294967295
                    switch (((("" + "a") + "a") + ("a" + "")) + "hello") {
                    case "":
                        var v42 float64 = 3.14159
                        _ = v42
                    case "hello":
                        v36 = 4294967295
                    }
                    _ = fi39
                    _ = v40
                    _ = v41
                }
            } else {
                var v43 int32 = int32(v36)
                fmt.Println(v21)
                v21 -= (v21 & (v21 % (((v21 >> ((255) & 63))) | 1)))
                _ = v43
            }
            v21 = 326
        }
        v2 += v2
        _ = fi37
    }
    fmt.Println(v21, v36, v21)
    fmt.Println(v21)
    switch v36 {
    case 41:
        var v44 uint32 = 4294967295
        _ = v44
    case -594:
        if false {
            for fi45 := 0; fi45 < 15; fi45++ {
                var v46 uint64 = 18446744073709551615
                _ = fi45
                _ = v46
            }
            if (true && true) {
                var v47 float32 = (3.14159 / ((-2.25 - (1e10 * -2.25)) + (0.1 + 1.5)))
                v36 = v36
                _ = v47
            } else {
                var v48 uint16 = 420
                v4 = append(v4, ("foo bar" + (("hello" + ((("a" + "") + "foo bar") + "\n\t")) + "")))
                switch (^(v36)) {
                case -2147483648:
                    v2 -= (-(v2))
                case 32767:
                    var v49 int8 = -65
                    _ = v49
                case 354:
                    var v50 uint8 = 255
                    _ = v50
                default:
                    v48 -= uint16(v21)
                }
                _ = v48
            }
        }
    }
    switch ("\n\t" + (("\n\t" + "hello") + "")) {
    case "\n\t":
        switch 42 {
        case -9223372036854775808:
            var v51 uint32 = (v21 | 65535)
            _ = v51
        case 4294967295:
            v36 += v36
        }
    case "a":
        v2 -= (v2 * v2)
    case "hello":
        var v52 int16 = -32768
        _ = v52
    default:
        v53 := GenS2{F0: v2, F1: genFn0()}
        _ = v53
    }
    fmt.Println(v21, v21, v36)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v36 = int(v2)
    if (v2 >= v2) {
        var v54 uint32 = v21
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v54
    }
    var v55 float32 = (-2.25 / 1.0)
    if (!true) {
        var v56 bool = (-159 > -9223372036854775808)
        for fi57 := 0; fi57 < 6; fi57++ {
            for fi58 := 0; fi58 < 20; fi58++ {
                var v59 int = fi57
                _ = fi58
                _ = v59
            }
            v56 = true
            v2 = (v2 ^ v2)
            _ = fi57
        }
        _ = v56
    } else {
        for fi60 := 0; fi60 < 19; fi60++ {
            if len(v4) > 0 {
                v4[0] = ""
            }
            _ = fi60
        }
        _ = len(v1)
    }
    var v61 int16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v55:", v55)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v36
    _ = v4
    _ = v55
    _ = v61
}
