package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int; F1 int64; F2 uint32 }
type GenS4 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS5 struct { F0 int64; F1 int64; F2 int64 }
type GenS6 struct { F0 int16; F1 uint16 }

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 2
    var v3 int16 = -32768
    var v4 uint16 = 2
    if false {
        fmt.Println(v3, v2, v4)
        switch -9223372036854775808 {
        case 0:
            for fi5 := 0; fi5 < 15; fi5++ {
                v6 := GenS1{F0: 237, F1: (fi5 | fi5)}
                _ = fi5
                _ = v6
            }
        case -866:
            v4 = v4
        default:
            var v7 bool = true
            _ = v7
        }
    }
    if (false || false) {
        switch 9223372036854775807 {
        case 0:
            if false {
                v4 = 255
            }
        case 9223372036854775807:
            v8 := GenS1{F0: uint8(v4), F1: -1}
            _ = v8
        case 42:
            switch "hello" {
            case "":
                v9 := GenS1{F0: 2, F1: 32767}
                _ = v9
            case "a":
                v10 := GenS1{F0: 0, F1: -2147483648}
                _ = v10
            }
        }
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 int8 = 42
            switch (fi11 << ((371) & 63)) {
            case -427:
                for fi13 := 0; fi13 < 13; fi13++ {
                    v12 += (v12 & (v12 & (v12 << ((1) & 63))))
                    _ = fi13
                }
            }
            switch "foo bar" {
            case "hello":
                if true {
                    for fi14 := 0; fi14 < 3; fi14++ {
                        var v15 int32 = 847
                        v1[2147483647] = -771
                        _ = len(v1)
                        var v16 int = int(v15)
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    var v17 float32 = ((0.0 / 1.0) / (3.14159 * 1.5))
                    for fi18 := 0; fi18 < 20; fi18++ {
                        var v19 uint = 466
                        fmt.Println(v3)
                        _ = fi18
                        _ = v19
                    }
                    _ = v17
                }
            case "a":
                for fi20 := 0; fi20 < 8; fi20++ {
                    fmt.Println(v3)
                    _ = fi20
                }
            case "foo bar":
                fmt.Println(v4, fi11, v3)
            }
            _ = fi11
            _ = v12
        }
        var v21 float32 = 3.14159
        _ = v21
    }
    var v22 int64 = 585
    if (517 != (v22 ^ v22)) {
        switch (("\n\t" + "") + (((("\n\t" + "") + ("a" + ("foo bar" + "hello"))) + "") + "foo bar")) {
        case "\n\t":
            _ = len(v1)
        default:
            v1[-1] = 318
        }
        var v23 float32 = (1.5 - float32(-2.25))
        v24 := []int{}
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        for fi26 := 0; fi26 < 4; fi26++ {
            v2 = (v2 - 0)
            var v27 int8 = 123
            fmt.Println(v4, v3)
            v22 -= (v22 + 170)
            if true {
                if false {
                    v28 := []int{(fi26 << ((1) & 63)), (fi25 ^ (-(fi25)))}
                    v27 = (v27 | v27)
                    v29 := GenS1{F0: 1, F1: (fi25 | 2147483647)}
                    _ = v28
                    _ = v29
                } else {
                    switch (((("a" + "foo bar") + ("" + "\n\t")) + "foo bar") + "a") {
                    case "":
                        var v30 int32 = 869
                        _ = v30
                    case "a":
                        v31 := []int{(-(fi26)), -9223372036854775808, 32768}
                        _ = v31
                    case "hello":
                        v4 = (v4 * (^(v4)))
                    default:
                        v32 := GenS1{F0: 120, F1: fi25}
                        _ = v32
                    }
                    if true {
                        v27 = v27
                    }
                    if true {
                        v33 := GenS2{F0: 255, F1: 42}
                        v34 := GenS2{F0: 873, F1: 1}
                        v35 := GenS1{F0: 42, F1: fi25}
                        _ = v33
                        _ = v34
                        _ = v35
                    } else {
                        switch -402 {
                        case 32767:
                            v36 := []int{(fi26 + 42), 2147483647}
                            _ = v36
                        case 853:
                            var v37 uint16 = 339
                            _ = v37
                        default:
                            fmt.Println(fi25)
                        }
                        v38 := GenS3{F0: (fi25 - (fi26 >> ((753) & 63))), F1: (-(v22)), F2: 885}
                        var v39 uint = 4294967295
                        _ = v38
                        _ = v39
                    }
                }
                var v40 int16 = v3
                v41 := GenS4{F0: (-(v22)), F1: 2, F2: v2}
                _ = v40
                _ = v41
            } else {
                if (42 >= -576) {
                    v22 -= (v22 / ((v22) | 1))
                }
                v42 := GenS2{F0: 255, F1: 42}
                _ = v42
            }
            var v43 uint16 = (v4 / ((uint16(v2)) | 1))
            _ = fi26
            _ = v27
            _ = v43
        }
        fmt.Println(v3)
        _ = fi25
    }
    fmt.Println(v4, v4, v4)
    for fi44 := 0; fi44 < 11; fi44++ {
        for fi45 := 0; fi45 < 2; fi45++ {
            v46 := GenS5{F0: v22, F1: v22, F2: -1}
            v3 = (v3 % (((-(v3))) | 1))
            _ = fi45
            _ = v46
        }
        v22 = 0
        v47 := GenS6{F0: 32767, F1: (v4 >> ((0) & 63))}
        _ = fi44
        _ = v47
    }
    var v48 float32 = ((((3.14159 - 1.5) / 1e10) - 0.0) / 0.1)
    v48 += float32(0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v48
}
