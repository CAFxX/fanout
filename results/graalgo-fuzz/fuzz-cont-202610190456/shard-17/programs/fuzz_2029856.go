package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 bool; F2 int32 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 int64; F1 uint; F2 uint64 }
type GenS5 struct { F0 uint; F1 bool }
func genFn0() int {
    switch -9223372036854775808 {
    case 32767:
        var v4 uint8 = 231
        _ = v4
    case 128:
        v5 := []int{}
        _ = v5
    }
    if (!false) {
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 uint32 = 2147483647
            v8 := GenS1{F0: -19}
            _ = fi6
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 7; fi9++ {
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 uint16 = 0
                _ = fi10
                _ = v11
            }
            switch "foo bar" {
            case "foo bar":
                fmt.Println(fi9)
            case "\n\t":
                for fi12 := 0; fi12 < 20; fi12++ {
                    var v13 uint32 = 1
                    var v14 string = "hello"
                    v15 := GenS2{F0: -128, F1: (!true), F2: 42}
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
            case "hello":
                v16 := GenS1{F0: -127}
                _ = v16
            }
            fmt.Println(fi9)
            _ = fi9
        }
    } else {
        for fi17 := 0; fi17 < 7; fi17++ {
            v18 := GenS1{F0: -1}
            _ = fi17
            _ = v18
        }
    }
    switch "" {
    case "\n\t":
        if (!(((float64(0.0) * (0.1 / 1.5)) != 0.1) && true)) {
            switch -1 {
            case 9223372036854775807:
                var v19 int16 = 926
                _ = v19
            case -350:
                if (942 <= 127) {
                    v20 := GenS3{F0: true, F1: 18446744073709551615}
                    _ = v20
                }
            }
            var v21 uint64 = 18446744073709551615
            var v22 int64 = 883
            _ = v21
            _ = v22
        } else {
            var v23 int8 = -1
            for fi24 := 0; fi24 < 4; fi24++ {
                var v25 int16 = 32767
                v23 = (v23 % (((v23 % ((v23) | 1))) | 1))
                var v26 int8 = v23
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = v23
        }
    default:
        var v27 uint8 = 15
        _ = v27
    }
    return -1
}

func main() {
    var v1 int8 = 42
    v2 := []uint32{1, 4294967295, 2147483648, 2147483647}
    v3 := []string{"\n\t", (("hello" + ("a" + "")) + "hello"), "", ("" + ("foo bar" + ("\n\t" + "a")))}
    v1 += v1
    for fi28 := 0; fi28 < 3; fi28++ {
        switch fi28 {
        case 2147483647:
            for fi29 := 0; fi29 < 14; fi29++ {
                v1 = (v1 + -1)
                _ = fi29
            }
        case -976:
            _ = len(v3)
        case 4294967295:
            for fi30 := 0; fi30 < 14; fi30++ {
                var v31 uint8 = 2
                _ = fi30
                _ = v31
            }
        }
        _ = fi28
    }
    if false {
        _ = len(v2)
        var v32 uint8 = 52
        _ = v32
    } else {
        var v33 uint16 = 384
        for fi34 := 0; fi34 < 5; fi34++ {
            var v35 bool = true
            var v36 int64 = 0
            v37 := GenS4{F0: (v36 | (v36 + (v36 >> ((18446744073709551615) & 63)))), F1: 891, F2: 609}
            _ = fi34
            _ = v35
            _ = v36
            _ = v37
        }
        if true {
            fmt.Println(v33, v33)
        } else {
            v1 = v1
            if true {
                v2 = append(v2, 2)
                _ = len(v2)
            } else {
                var v38 int32 = 533
                _ = v38
            }
        }
        _ = v33
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi39 := 0; fi39 < 15; fi39++ {
        v1 = v1
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi39
    }
    if len(v2) > 0 {
        v2[0] = 997
    }
    for fi40 := 0; fi40 < 13; fi40++ {
        var v41 uint32 = 655
        v41 += v41
        _ = fi40
        _ = v41
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        _ = len(v2)
        for fi42 := 0; fi42 < 16; fi42++ {
            var v43 uint16 = 65535
            var v44 int = genFn0()
            _ = fi42
            _ = v43
            _ = v44
        }
    }
    var v45 bool = true
    var v46 int32 = 42
    fmt.Println(v1)
    v47 := GenS4{F0: 42, F1: 42, F2: 1}
    for fi48 := 0; fi48 < 9; fi48++ {
        v49 := GenS5{F0: 1, F1: true}
        v46 = (v46 | -1)
        _ = fi48
        _ = v49
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (813 != -1) {
        v1 -= v1
    } else {
        var v50 bool = (0.0 <= 3.14159)
        v51 := GenS5{F0: 4294967295, F1: v50}
        _ = v50
        _ = v51
    }
    var v52 bool = (245 < 1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v47
    _ = v52
}
