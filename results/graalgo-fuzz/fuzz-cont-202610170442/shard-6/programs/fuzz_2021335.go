package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 uint32; F1 bool; F2 uint }
type GenS6 struct { F0 bool; F1 int16 }
type GenS7 struct { F0 int64; F1 bool }

func main() {
    var v1 uint32 = 2147483648
    var v2 float32 = float32((1e10 - -2.25))
    var v3 uint = 50
    if false {
        v4 := GenS1{F0: -3, F1: 128}
        switch "foo bar" {
        case "\n\t":
            var v5 int32 = 1
            _ = v5
        }
        var v6 uint = 20
        _ = v4
        _ = v6
    } else {
        switch 506 {
        case -887:
            for fi7 := 0; fi7 < 8; fi7++ {
                var v8 int16 = int16(v1)
                var v9 uint16 = 42
                if (!((^(v9)) > (v9 | v9))) {
                    v10 := []int{fi7, 618}
                    var v11 float32 = v2
                    if (-9223372036854775808 > 650) {
                        for fi12 := 0; fi12 < 4; fi12++ {
                            var v13 int64 = 0
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            _ = fi12
                            _ = v13
                        }
                    }
                    _ = v10
                    _ = v11
                } else {
                    var v14 uint16 = 387
                    var v15 float32 = (3.14159 * (v2 + v2))
                    _ = v14
                    _ = v15
                }
                _ = fi7
                _ = v8
                _ = v9
            }
        case 128:
            v16 := []int{-312}
            _ = v16
        }
        var v17 int8 = 1
        _ = v17
    }
    if true {
        fmt.Println(v3)
        if ((!((true || (v3 <= v3)) || true)) || false) {
            v18 := GenS2{F0: (v3 << (((^(v3))) & 63)), F1: -9223372036854775808, F2: 42}
            var v19 uint32 = (v1 >> (((^(v3))) & 63))
            _ = v18
            _ = v19
        } else {
            v20 := GenS2{F0: v3, F1: -998, F2: uint16(v3)}
            v21 := GenS2{F0: 704, F1: 448, F2: 1}
            v22 := GenS3{F0: 2, F1: 42}
            _ = v20
            _ = v21
            _ = v22
        }
    } else {
        v2 += 0.0
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        v2 = v2
        _ = fi23
    }
    v2 -= (v2 / v2)
    fmt.Println(v1)
    v1 = v1
    switch -764 {
    case 128:
        v24 := GenS3{F0: 1, F1: 0}
        _ = v24
    case 127:
        switch -29 {
        case -9223372036854775808:
            switch ("hello" + "foo bar") {
            case "\n\t":
                v25 := GenS4{F0: 255, F1: 1}
                _ = v25
            case "hello":
                var v26 uint = v3
                _ = v26
            case "":
                for fi27 := 0; fi27 < 19; fi27++ {
                    var v28 uint64 = 121
                    _ = fi27
                    _ = v28
                }
            default:
                v29 := []int{1, -290}
                _ = v29
            }
        }
    }
    var v30 int16 = 1000
    if (!(1e10 >= 1.5)) {
        for fi31 := 0; fi31 < 20; fi31++ {
            v2 -= v2
            v32 := []int{}
            switch "a" {
            case "hello":
                v30 += v30
            case "foo bar":
                var v33 int8 = 1
                _ = v33
            case "\n\t":
                v34 := GenS5{F0: v1, F1: true, F2: 18446744073709551615}
                _ = v34
            default:
                for fi35 := 0; fi35 < 14; fi35++ {
                    var v36 uint32 = 4294967295
                    var v37 int = (^(fi35))
                    switch ("hello" + "\n\t") {
                    case "a":
                        fmt.Println(v36, v36)
                    case "\n\t":
                        _ = len(v32)
                    case "":
                        var v38 int = 48
                        _ = v38
                    }
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            }
            _ = fi31
            _ = v32
        }
    } else {
        var v39 string = "hello"
        if false {
            v1 = v1
            fmt.Println(v39, v2, v1)
        }
        var v40 uint = (v3 | (v3 << ((v3) & 63)))
        _ = v39
        _ = v40
    }
    v41 := []int{}
    for fi42 := 0; fi42 < 13; fi42++ {
        var v43 uint16 = 65535
        for fi44 := 0; fi44 < 3; fi44++ {
            v45 := GenS6{F0: ((v30 + v30) > (v30 + v30)), F1: (v30 - -32768)}
            v1 += (v1 & v1)
            _ = fi44
            _ = v45
        }
        _ = fi42
        _ = v43
    }
    v46 := GenS6{F0: (65535 != (v3 >> (((v3 >> ((1) & 63))) & 63))), F1: v30}
    if (0.0 <= 3.14159) {
        var v47 int = -2147483648
        _ = v47
    }
    v48 := GenS7{F0: -1, F1: (96 >= 1)}
    fmt.Println(v30)
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v41:", len(v41), v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v41
    _ = v46
    _ = v48
}
