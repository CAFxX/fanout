package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 uint }

func main() {
    var v1 uint16 = 527
    var v2 int32 = 961
    var v3 uint8 = 128
    var v4 uint32 = 87
    var v5 int8 = 44
    var v6 uint = 7
    var v7 float64 = ((1.5 - (1.5 - 1e10)) * 0.0)
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: "\n\t", F1: 376, F2: (^(v3))}
        var v10 int = fi8
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 float32 = -2.25
        _ = fi11
        _ = v12
    }
    switch -630 {
    case 1:
        v13 := []int{128}
        _ = v13
    case 192:
        var v14 uint32 = (v4 - v4)
        _ = v14
    }
    var v15 uint8 = v3
    switch "\n\t" {
    case "foo bar":
        for fi16 := 0; fi16 < 7; fi16++ {
            v5 += (v5 ^ 0)
            for fi17 := 0; fi17 < 7; fi17++ {
                v6 = (^(v6))
                if false {
                    v7 = (v7 * 3.14159)
                    v7 = (-2.25 * v7)
                    v18 := GenS1{F0: "foo bar", F1: 1, F2: v3}
                    _ = v18
                }
                _ = fi17
            }
            _ = fi16
        }
    case "\n\t":
        v19 := GenS2{F0: v3}
        _ = v19
    default:
        v20 := []int{2147483647, -392}
        _ = v20
    }
    var v21 float32 = ((-2.25 - 3.14159) - float32((1e10 + v7)))
    if true {
        switch ("foo bar" + "a") {
        case "\n\t":
            var v22 int8 = v5
            _ = v22
        case "a":
            switch 354 {
            case 1:
                v23 := []int{-1, 0, int(v3)}
                _ = v23
            }
        case "foo bar":
            fmt.Println(v7)
        }
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        switch "a" {
        case "foo bar":
            var v25 uint16 = v1
            _ = v25
        case "":
            var v26 uint32 = 595
            _ = v26
        case "\n\t":
            v27 := GenS3{F0: 2}
            _ = v27
        default:
            var v28 int64 = -9223372036854775808
            _ = v28
        }
        for fi29 := 0; fi29 < 18; fi29++ {
            switch "foo bar" {
            case "a":
                if true {
                    v30 := []int{}
                    v31 := GenS4{F0: 9223372036854775807, F1: 4294967295}
                    v30 = append(v30, fi24)
                    _ = v30
                    _ = v31
                } else {
                    v7 = (-2.25 * -2.25)
                }
                switch (fi24 >> ((uint(v4)) & 63)) {
                case 360:
                    var v32 float64 = v7
                    _ = v32
                }
            case "\n\t":
                if ((true && (v4 != 4294967295)) && ((false || false) && false)) {
                    fmt.Println(v1)
                } else {
                    v33 := []int{-84}
                    _ = v33
                }
            case "hello":
                var v34 uint64 = 598
                v2 = v2
                _ = v34
            default:
                v2 += v2
            }
            var v35 int8 = (v5 * (v5 % (((-(v5))) | 1)))
            v36 := []int{fi29}
            _ = fi29
            _ = v35
            _ = v36
        }
        _ = fi24
    }
    v37 := []int{-2147483648, -1, 32768, 975}
    if (1 == 853) {
        var v38 int32 = (-(v2))
        _ = v38
    } else {
        var v39 int32 = v2
        for fi40 := 0; fi40 < 8; fi40++ {
            fmt.Println(v5)
            _ = fi40
        }
        v2 = -251
        _ = v39
    }
    v3 = (v3 | (v3 + 234))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
