package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 bool; F1 int; F2 int64 }
type GenS7 struct { F0 int8 }
func genFn0(p0 uint8, p1 int32) int32 {
    var v4 uint16 = 0
    for fi5 := 0; fi5 < 5; fi5++ {
        for fi6 := 0; fi6 < 13; fi6++ {
            v4 = v4
            v7 := []int{fi5, 4294967295, fi6, 128}
            v7 = append(v7, fi5)
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    p1 -= p1
    _ = p0
    _ = p1
    _ = v4
    return p1
}

func main() {
    var v1 int = 128
    var v2 float64 = 1.5
    v3 := make(map[string]int)
    var v8 float64 = (v2 - (float64(float32(v2)) / (((v2 - v2) * 1.5) / v2)))
    v9 := []int{}
    fmt.Println(v8, v8, v2)
    v2 = v8
    v8 -= -2.25
    if true {
        switch (("" + "foo bar") + "hello") {
        case "\n\t":
            v10 := GenS1{F0: 42}
            _ = v10
        }
        for fi11 := 0; fi11 < 1; fi11++ {
            var v12 uint64 = 476
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 16; fi13++ {
            for fi14 := 0; fi14 < 10; fi14++ {
                if true {
                    var v15 int32 = 2147483647
                    switch (fi13 - (^(fi13))) {
                    case 2147483647:
                        var v16 uint = 460
                        _ = v16
                    }
                    v15 = genFn0(101, 507)
                    _ = v15
                }
                _ = fi14
            }
            switch fi13 {
            case 32768:
                if ("" >= "hello") {
                    switch fi13 {
                    case 2147483647:
                        var v17 uint16 = 0
                        _ = v17
                    case 855:
                        v18 := GenS2{F0: 18446744073709551615}
                        _ = v18
                    case 0:
                        var v19 uint = 42
                        _ = v19
                    default:
                        for fi20 := 0; fi20 < 17; fi20++ {
                            if (0 > 255) {
                                v21 := GenS2{F0: 2}
                                v22 := GenS3{F0: 1, F1: 2}
                                v23 := GenS4{F0: (v1 ^ fi20)}
                                fmt.Println(v1, fi20)
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                            _ = fi20
                        }
                    }
                } else {
                    for fi24 := 0; fi24 < 15; fi24++ {
                        var v25 string = (("hello" + "a") + "")
                        _ = fi24
                        _ = v25
                    }
                }
            default:
                switch "\n\t" {
                case "foo bar":
                    v8 = -2.25
                case "hello":
                    var v26 float32 = 0.1
                    _ = v26
                default:
                    v27 := GenS4{F0: 1}
                    _ = v27
                }
            }
            v28 := GenS5{F0: -880, F1: (fi13 * (-(fi13)))}
            _ = fi13
            _ = v28
        }
    }
    v29 := GenS6{F0: false, F1: v1, F2: -261}
    var v30 uint8 = 0
    for fi31 := 0; fi31 < 17; fi31++ {
        v32 := GenS7{F0: 42}
        var v33 int16 = 745
        _ = fi31
        _ = v32
        _ = v33
    }
    v8 -= (-2.25 * float64(1.5))
    v30 = v30
    for fi34 := 0; fi34 < 3; fi34++ {
        v35 := GenS6{F0: false, F1: v1, F2: 9223372036854775807}
        _ = fi34
        _ = v35
    }
    v2 = ((3.14159 / (1e10 * float64(0.1))) - ((v2 / float64(1.5)) * (((-2.25 * float64(0.0)) / v8) / 1e10)))
    switch ("" + ("\n\t" + "foo bar")) {
    case "a":
        var v36 string = "a"
        _ = v36
    case "":
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    case "foo bar":
        v37 := GenS3{F0: 42, F1: 18446744073709551615}
        _ = v37
    }
    if (false || (!false)) {
        if (1 > -1) {
            v38 := GenS2{F0: 18446744073709551615}
            v39 := GenS1{F0: 529}
            var v40 string = ("" + "foo bar")
            _ = v38
            _ = v39
            _ = v40
        }
    }
    v41 := GenS2{F0: 18446744073709551615}
    v30 -= v30
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v41
    _ = v8
    _ = v9
}
