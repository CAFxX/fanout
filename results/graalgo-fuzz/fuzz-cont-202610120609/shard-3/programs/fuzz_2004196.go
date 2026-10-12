package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint }
type GenS3 struct { F0 int8; F1 int; F2 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 string }
func genFn0(p0 uint32) uint8 {
    fmt.Println(p0)
    p0 = 2
    _ = p0
    return 48
}
func genFn1() uint8 {
    v7 := []int{1, 127}
    var v8 int64 = -379
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := GenS1{F0: -32768, F1: ("" + "\n\t"), F2: -933}
        for fi11 := 0; fi11 < 7; fi11++ {
            if true {
                v8 = (v8 % ((-434) | 1))
                var v12 int32 = 1
                switch (("hello" + "foo bar") + "") {
                case "a":
                    v12 -= (v12 / (((v12 << ((2) & 63))) | 1))
                case "foo bar":
                    switch "hello" {
                    case "\n\t":
                        var v13 uint16 = 1
                        _ = v13
                    case "hello":
                        if len(v7) > 0 {
                            v7[0] = -2147483648
                        }
                    default:
                        var v14 float64 = ((1.5 + 1e10) - ((1e10 / 1.0) - -2.25))
                        _ = v14
                    }
                }
                _ = v12
            }
            var v15 uint16 = 425
            _ = fi11
            _ = v15
        }
        var v16 int64 = (v8 + (-(v8)))
        _ = fi9
        _ = v10
        _ = v16
    }
    var v17 int64 = v8
    fmt.Println(v17)
    _ = v17
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint32 = 42
    v2 := []int{1}
    var v3 uint = 0
    var v4 uint64 = 1
    var v5 uint = v3
    var v6 int16 = -32768
    if (!false) {
        v5 = (v5 << ((2) & 63))
        var v18 int16 = v6
        switch -9223372036854775808 {
        case -97:
            switch ((("hello" + "hello") + "hello") + "hello") {
            case "":
                v5 = (v3 * v5)
            case "\n\t":
                var v19 uint16 = 1
                _ = v19
            case "a":
                switch "a" {
                case "foo bar":
                    v2 = append(v2, -2147483648)
                default:
                    v20 := GenS1{F0: v6, F1: "hello", F2: 883}
                    _ = v20
                }
            }
        }
        _ = v18
    } else {
        var v21 uint64 = 396
        _ = v21
    }
    v22 := GenS2{F0: 127, F1: (v1 % ((v1) | 1)), F2: (v3 << ((255) & 63))}
    if true {
        v4 = 4294967295
    } else {
        v3 = v5
        switch "foo bar" {
        case "foo bar":
            _ = len(v2)
        case "\n\t":
            var v23 int64 = 405
            _ = v23
        case "hello":
            if false {
                var v24 uint32 = v1
                v25 := GenS1{F0: v6, F1: "a", F2: 568}
                _ = v24
                _ = v25
            }
        default:
            var v26 int8 = 0
            _ = v26
        }
    }
    switch ("a" + "a") {
    case "hello":
        fmt.Println(v3, v6, v1)
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "":
        if len(v2) > 0 {
            v2[0] = 537
        }
    }
    if false {
        fmt.Println(v3, v6, v4)
    } else {
        fmt.Println(v1, v4, v4)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi27 := 0; fi27 < 16; fi27++ {
            switch "hello" {
            case "a":
                _ = len(v2)
            default:
                _ = len(v2)
            }
            _ = fi27
        }
    }
    v5 += v5
    switch 0 {
    case -833:
        switch "\n\t" {
        case "hello":
            v5 = 446
        }
    case 127:
        var v28 int16 = v6
        _ = v28
    case -9223372036854775808:
        _ = len(v2)
    default:
        v29 := GenS3{F0: -1, F1: 0, F2: v4}
        _ = v29
    }
    v5 += (v5 / (((v3 | uint(v1))) | 1))
    switch "\n\t" {
    case "\n\t":
        for fi30 := 0; fi30 < 16; fi30++ {
            v1 -= v1
            v6 -= (v6 & v6)
            _ = fi30
        }
    case "":
        for fi31 := 0; fi31 < 1; fi31++ {
            v32 := GenS4{F0: 0}
            var v33 bool = (false && (!false))
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    var v34 uint32 = 2147483647
    fmt.Println(v1, v4, v1)
    v4 -= uint64(v1)
    v35 := GenS4{F0: 788}
    if (4294967295 >= 145) {
        v6 = v6
        var v36 int8 = 100
        v1 -= (v34 - (v1 | uint32(v5)))
        _ = v36
    }
    var v37 uint8 = genFn0(v34)
    switch ("\n\t" + "") {
    case "\n\t":
        switch -295 {
        case 2147483647:
            var v38 float32 = 0.0
            _ = v38
        default:
            v2 = append(v2, 32768)
        }
    default:
        if false {
            v39 := GenS5{F0: v6, F1: "hello"}
            _ = v39
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
