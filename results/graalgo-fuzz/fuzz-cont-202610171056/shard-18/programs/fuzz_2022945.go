package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 int8; F1 bool; F2 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 bool }

func main() {
    var v1 float64 = 1e10
    v2 := []int{-945}
    var v3 int32 = 2147483647
    v4 := []int{64}
    var v5 bool = (v3 < v3)
    var v6 int8 = 127
    var v7 string = ((((("\n\t" + "") + ("" + "\n\t")) + "\n\t") + "a") + "")
    v5 = v5
    switch ("a" + v7) {
    case "hello":
        var v8 uint32 = 757
        _ = v8
    case "":
        for fi9 := 0; fi9 < 19; fi9++ {
            for fi10 := 0; fi10 < 17; fi10++ {
                switch "a" {
                case "\n\t":
                    switch fi10 {
                    case 9223372036854775807:
                        switch (fi9 >> ((916) & 63)) {
                        case 4294967295:
                            v6 -= v6
                        default:
                            var v11 int64 = 816
                            _ = v11
                        }
                        var v12 uint32 = 1
                        v6 = (v6 & (^(v6)))
                        _ = v12
                    case 32768:
                        v7 = v7
                    }
                case "hello":
                    v13 := GenS1{F0: 0, F1: 1, F2: 22}
                    _ = v13
                case "foo bar":
                    var v14 uint32 = 65535
                    _ = v14
                default:
                    v15 := GenS1{F0: 2, F1: 2, F2: 0}
                    _ = v15
                }
                v16 := GenS1{F0: 693, F1: 128, F2: -884}
                if v5 {
                    v4 = append(v4, fi9)
                } else {
                    var v17 int32 = -272
                    _ = v17
                }
                v18 := GenS1{F0: 65535, F1: 109, F2: -227}
                _ = fi10
                _ = v16
                _ = v18
            }
            _ = fi9
        }
    case "a":
        for fi19 := 0; fi19 < 11; fi19++ {
            var v20 string = v7
            _ = fi19
            _ = v20
        }
    }
    v7 = v7
    v21 := GenS2{F0: (!true), F1: 289}
    v6 -= v6
    for fi22 := 0; fi22 < 9; fi22++ {
        for fi23 := 0; fi23 < 3; fi23++ {
            fmt.Println(v7, fi22, v5)
            v4 = append(v4, (fi22 / ((9223372036854775807) | 1)))
            if (((v3 | 2147483647) > (v3 << ((0) & 63))) && true) {
                if (!true) {
                    fmt.Println(v7, v5)
                    v24 := GenS2{F0: v5, F1: -472}
                    var v25 uint16 = 42
                    _ = v24
                    _ = v25
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = fi23
        }
        var v26 uint32 = 2
        v27 := GenS2{F0: (v1 > (v1 + v1)), F1: fi22}
        _ = fi22
        _ = v26
        _ = v27
    }
    var v28 uint8 = 128
    if len(v2) > 0 {
        v2[0] = -501
    }
    for fi29 := 0; fi29 < 20; fi29++ {
        var v30 int8 = v6
        _ = fi29
        _ = v30
    }
    v7 = (v7 + (v7 + (v7 + v7)))
    if (0.0 == 1e10) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if (-899 == -331) {
            fmt.Println(v5, v3)
        } else {
            var v31 float64 = v1
            v4 = append(v4, -1)
            var v32 int16 = 598
            _ = v31
            _ = v32
        }
        v33 := GenS3{F0: (v6 - -128), F1: (737 != 18446744073709551615), F2: v28}
        _ = v33
    } else {
        if v5 {
            v34 := GenS4{F0: false}
            for fi35 := 0; fi35 < 9; fi35++ {
                v6 -= 1
                _ = fi35
            }
            v6 = -128
            _ = v34
        } else {
            v36 := GenS5{F0: (v28 | 42), F1: v5}
            v5 = v5
            _ = v36
        }
        switch -2147483648 {
        case 42:
            v37 := GenS3{F0: (v6 | (v6 % ((v6) | 1))), F1: (!false), F2: (v28 | (v28 | 1))}
            _ = v37
        }
        for fi38 := 0; fi38 < 19; fi38++ {
            v4 = append(v4, (fi38 << ((1) & 63)))
            v39 := GenS3{F0: (v6 ^ (v6 | (-(v6)))), F1: v5, F2: (v28 % ((0) | 1))}
            _ = fi38
            _ = v39
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
