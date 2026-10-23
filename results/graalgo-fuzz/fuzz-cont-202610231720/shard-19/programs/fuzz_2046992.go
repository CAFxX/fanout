package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 int8; F1 uint64; F2 int }
type GenS4 struct { F0 int64 }
func genFn0() uint8 {
    switch (("" + "\n\t") + ("hello" + ("hello" + ""))) {
    case "foo bar":
        var v5 int8 = 100
        _ = v5
    default:
        v6 := []int{692, 664}
        _ = v6
    }
    if true {
        v7 := []int{98}
        _ = v7
    } else {
        var v8 bool = true
        v9 := GenS1{F0: 0, F1: -9223372036854775808, F2: 0}
        _ = v8
        _ = v9
    }
    var v10 uint32 = 65535
    fmt.Println(v10)
    _ = v10
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 18446744073709551615
    var v3 int8 = 0
    var v4 int = -9223372036854775808
    v11 := GenS2{F0: 2147483647, F1: 42}
    var v12 string = "hello"
    switch "foo bar" {
    case "a":
        v2 += 42
    case "foo bar":
        for fi13 := 0; fi13 < 1; fi13++ {
            for fi14 := 0; fi14 < 7; fi14++ {
                var v15 int32 = 42
                v16 := []int{}
                if false {
                    v1[148] = true
                    var v17 int64 = 42
                    _ = v17
                } else {
                    v2 = (v2 | (v2 % (((^(v2))) | 1)))
                    v12 = v12
                }
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi13
        }
    default:
        var v18 uint16 = 1
        _ = v18
    }
    switch (v4 * 128) {
    case 127:
        var v19 int8 = v3
        _ = v19
    case 4294967295:
        var v20 int16 = 1
        _ = v20
    }
    var v21 uint8 = 128
    if false {
        if false {
            for fi22 := 0; fi22 < 19; fi22++ {
                for fi23 := 0; fi23 < 1; fi23++ {
                    if true {
                        var v24 int16 = 32767
                        _ = v24
                    }
                    _ = fi23
                }
                _ = fi22
            }
            v1[2147483647] = ((v4 << ((2) & 63)) >= (-(v4)))
            v25 := GenS3{F0: 100, F1: 491, F2: -660}
            _ = v25
        }
    }
    v1[1] = true
    if false {
        var v26 int = -815
        var v27 float32 = (1.5 + 3.14159)
        for fi28 := 0; fi28 < 8; fi28++ {
            if true {
                switch (v26 ^ (^(v4))) {
                case 230:
                    switch "\n\t" {
                    case "\n\t":
                        if true {
                            fmt.Println(fi28, v2)
                        }
                    case "hello":
                        for fi29 := 0; fi29 < 8; fi29++ {
                            if ((!false) && (32767 == 2)) {
                                v27 = ((float32(-2.25) / 0.1) - (v27 + (v27 * (0.1 + ((v27 - v27) - v27)))))
                            }
                            delete(v1, uint32(v2))
                            _ = fi29
                        }
                    case "foo bar":
                        switch (v12 + v12) {
                        case "foo bar":
                            if false {
                                fmt.Println(v21)
                                fmt.Println(v3, v4)
                                var v30 float32 = v27
                                _ = v30
                            }
                        case "\n\t":
                            _ = len(v1)
                        case "":
                            fmt.Println(v12, v27, v4)
                        default:
                            v31 := GenS2{F0: 42, F1: 81}
                            _ = v31
                        }
                    default:
                        var v32 int8 = (v3 | v3)
                        _ = v32
                    }
                case 4294967295:
                    var v33 int64 = -681
                    _ = v33
                default:
                    var v34 bool = (((127 < (^(v3))) && false) || false)
                    _ = v34
                }
                v35 := GenS2{F0: 42, F1: 255}
                var v36 int64 = -1
                _ = v35
                _ = v36
            }
            _ = fi28
        }
        _ = v26
        _ = v27
    }
    if (0.1 >= 0.0) {
        var v37 int = v4
        if (42 != -9223372036854775808) {
            v38 := GenS2{F0: uint32(v21), F1: 122}
            _ = v38
        }
        var v39 string = v12
        _ = v37
        _ = v39
    } else {
        v40 := GenS4{F0: -1}
        var v41 uint = 599
        for fi42 := 0; fi42 < 15; fi42++ {
            v43 := GenS2{F0: 876, F1: v41}
            _ = fi42
            _ = v43
        }
        _ = v40
        _ = v41
    }
    switch int(v2) {
    case 32767:
        switch v4 {
        case 1:
            switch v4 {
            case 127:
                var v44 float64 = float64((0.0 + (1e10 + float32(3.14159))))
                _ = v44
            }
        case -2147483648:
            var v45 uint32 = 365
            _ = v45
        }
    default:
        var v46 int = (^(v4))
        _ = v46
    }
    v4 -= (v4 & v4)
    if false {
        v4 = (-(v4))
        v4 = v4
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
