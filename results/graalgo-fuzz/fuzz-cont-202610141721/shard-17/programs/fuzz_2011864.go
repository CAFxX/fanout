package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 int; F1 int32 }
func genFn0(p0 uint8, p1 int64, p2 uint64) int32 {
    var v5 int16 = -32768
    switch "hello" {
    case "\n\t":
        if ((p0 + p0) == 178) {
            var v6 bool = (!(255 != 992))
            fmt.Println(v6)
            _ = v6
        }
    case "foo bar":
        v5 = (v5 ^ v5)
    case "hello":
        var v7 int32 = -2147483648
        _ = v7
    default:
        v8 := []int{-9223372036854775808, 42}
        _ = v8
    }
    fmt.Println(p2)
    var v9 int = int(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v9
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{}
    var v3 string = "hello"
    var v4 string = v3
    switch 1 {
    case 127:
        fmt.Println(v3, v4)
    default:
        if false {
            if true {
                for fi10 := 0; fi10 < 12; fi10++ {
                    v3 = v3
                    _ = fi10
                }
            }
        }
    }
    fmt.Println(v3)
    var v11 uint = 4294967295
    var v12 uint = v11
    if ((v11 / (((v11 - (v11 >> ((v11) & 63)))) | 1)) < (v12 | v11)) {
        if len(v2) > 0 {
            v2[0] = 287
        }
        var v13 uint32 = 723
        _ = v13
    }
    switch 2147483647 {
    case 954:
        var v14 int32 = 2147483647
        _ = v14
    }
    v15 := GenS1{F0: 2, F1: 0, F2: 2}
    var v16 int16 = -32768
    if true {
        if (4294967295 == v11) {
            v17 := GenS1{F0: v11, F1: 42, F2: 2}
            _ = v17
        } else {
            var v18 bool = (true && (false && true))
            var v19 float64 = float64(float32(0.1))
            v12 -= 930
            _ = v18
            _ = v19
        }
        v1[558] = "hello"
        v20 := GenS1{F0: (v11 << ((294) & 63)), F1: -1, F2: 11}
        _ = v20
    }
    switch v4 {
    case "\n\t":
        for fi21 := 0; fi21 < 2; fi21++ {
            v22 := GenS2{F0: (^(v11)), F1: (v11 | (v11 | 0))}
            _ = fi21
            _ = v22
        }
    case "":
        v23 := GenS2{F0: (v11 - v11), F1: 149}
        _ = v23
    default:
        v11 = v12
    }
    for fi24 := 0; fi24 < 14; fi24++ {
        if ((384 < 46) || false) {
            switch (fi24 >> ((2) & 63)) {
            case 9223372036854775807:
                v3 = v3
            case -2147483648:
                var v25 int32 = 1
                _ = v25
            default:
                v26 := GenS2{F0: v12, F1: v11}
                _ = v26
            }
        }
        switch fi24 {
        case 127:
            v12 = 18446744073709551615
        case -9223372036854775808:
            for fi27 := 0; fi27 < 11; fi27++ {
                var v28 uint64 = 1
                fmt.Println(v28, v28, fi27)
                fmt.Println(v4, fi27, v28)
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
                fmt.Println(v3, fi24, v11)
                var v29 string = v3
                _ = fi27
                _ = v28
                _ = v29
            }
        case 111:
            if true {
                v16 -= (v16 % (((v16 * (v16 + v16))) | 1))
                var v30 int16 = 42
                _ = v30
            } else {
                for fi31 := 0; fi31 < 11; fi31++ {
                    var v32 string = "hello"
                    _ = fi31
                    _ = v32
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if (genFn0(170, -9223372036854775808, 0) >= 2147483647) {
                    switch ("a" + ((v4 + v4) + (v3 + v4))) {
                    case "":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "a":
                        delete(v1, 0)
                    }
                    v33 := GenS3{F0: -9223372036854775808, F1: ((!true) && false)}
                    _ = v33
                } else {
                    for fi34 := 0; fi34 < 12; fi34++ {
                        var v35 int32 = 0
                        var v36 float32 = (((-2.25 * 1e10) * (0.1 * (-2.25 + 0.0))) / ((0.1 - 1e10) * 3.14159))
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    for fi37 := 0; fi37 < 3; fi37++ {
                        var v38 int = 128
                        _ = fi37
                        _ = v38
                    }
                    for fi39 := 0; fi39 < 11; fi39++ {
                        _ = len(v2)
                        _ = fi39
                    }
                }
            }
        }
        for fi40 := 0; fi40 < 1; fi40++ {
            switch (fi40 % ((fi40) | 1)) {
            case 32767:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 32768:
                fmt.Println(v3)
            default:
                v41 := GenS3{F0: 0, F1: (3.14159 > float64((float32((1e10 + 1.5)) - float32(1e10))))}
                _ = v41
            }
            _ = fi40
        }
        _ = fi24
    }
    for fi42 := 0; fi42 < 9; fi42++ {
        v4 = "hello"
        _ = fi42
    }
    var v43 uint = v12
    if true {
        var v44 int32 = 0
        var v45 int = 127
        var v46 uint32 = 0
        _ = v44
        _ = v45
        _ = v46
    }
    switch "a" {
    case "hello":
        var v47 string = "foo bar"
        _ = v47
    }
    if (0.0 < 0.1) {
        var v48 int64 = 42
        v49 := GenS4{F0: 0, F1: 42}
        var v50 int8 = 0
        _ = v48
        _ = v49
        _ = v50
    }
    var v51 int32 = 1
    var v52 int = -331
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v51
    _ = v52
}
