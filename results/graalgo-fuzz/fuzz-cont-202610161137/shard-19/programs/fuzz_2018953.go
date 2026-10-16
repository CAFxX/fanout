package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 int }
func genFn0() uint16 {
    for fi6 := 0; fi6 < 18; fi6++ {
        if ("" >= "hello") {
            if ((fi6 / (((fi6 >> ((2) & 63))) | 1)) == 884) {
                if (!true) {
                    switch "\n\t" {
                    case "a":
                        var v7 uint64 = 18446744073709551615
                        _ = v7
                    default:
                        fmt.Println(fi6)
                    }
                    if ((fi6 + fi6) < fi6) {
                        v8 := []int{}
                        switch "hello" {
                        case "a":
                            switch "\n\t" {
                            case "foo bar":
                                var v9 uint8 = 85
                                _ = v9
                            case "\n\t":
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                            case "a":
                                if false {
                                    var v10 uint16 = 1
                                    _ = v10
                                }
                            default:
                                fmt.Println(fi6)
                            }
                        }
                        _ = v8
                    }
                    v11 := []int{}
                    _ = v11
                } else {
                    fmt.Println(fi6)
                    var v12 bool = (!false)
                    v12 = (!(((!v12) && (0.0 == 1e10)) || (!(207 != 58))))
                    _ = v12
                }
                var v13 int64 = -1
                _ = v13
            }
            var v14 uint8 = 127
            if false {
                for fi15 := 0; fi15 < 12; fi15++ {
                    fmt.Println(fi6, v14)
                    switch 9223372036854775807 {
                    case 127:
                        v16 := GenS1{F0: 0, F1: 18446744073709551615, F2: v14}
                        _ = v16
                    case 128:
                        v14 = (v14 % ((v14) | 1))
                    case 0:
                        var v17 uint64 = 2
                        _ = v17
                    default:
                        v14 -= v14
                    }
                    _ = fi15
                }
            }
            _ = v14
        }
        var v18 int = (fi6 << ((2) & 63))
        _ = fi6
        _ = v18
    }
    switch -9223372036854775808 {
    case -1:
        switch "a" {
        case "":
            var v19 uint = 881
            _ = v19
        }
    case 0:
        v20 := []int{-1, 9223372036854775807, 1}
        _ = v20
    default:
        var v21 float32 = -2.25
        _ = v21
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 bool = true
        switch -9223372036854775808 {
        case 32767:
            v23 = (!v23)
        case -879:
            v24 := GenS1{F0: 1, F1: 560, F2: 55}
            _ = v24
        case 42:
            v25 := GenS1{F0: 1, F1: 690, F2: 255}
            _ = v25
        }
        v26 := []int{2147483647, -9223372036854775808, fi22, 9223372036854775807}
        _ = fi22
        _ = v23
        _ = v26
    }
    var v27 uint = 2
    _ = v27
    return 666
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]bool)
    v3 := []uint32{4294967295, 0}
    var v4 uint = 4294967295
    var v5 uint = v4
    if false {
        var v28 float64 = 0.1
        fmt.Println(v4, v4)
        _ = v28
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        switch "hello" {
        case "foo bar":
            var v30 uint = v4
            _ = v30
        }
        v4 += v5
        _ = fi29
    }
    switch "foo bar" {
    case "\n\t":
        _ = len(v3)
    case "a":
        if true {
            _ = len(v3)
            fmt.Println(v5, v5)
        } else {
            for fi31 := 0; fi31 < 20; fi31++ {
                v5 += 65535
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if false {
                    v32 := GenS1{F0: 0, F1: 733, F2: 255}
                    _ = v32
                }
                _ = fi31
            }
            v33 := GenS2{F0: ((1 < -128) || (224 < 0)), F1: 217}
            fmt.Println(v5)
            _ = v33
        }
    }
    switch -9223372036854775808 {
    case 0:
        var v34 int64 = -1
        _ = v34
    }
    for fi35 := 0; fi35 < 1; fi35++ {
        for fi36 := 0; fi36 < 8; fi36++ {
            v37 := GenS3{F0: 1, F1: 2}
            v4 = v4
            _ = fi36
            _ = v37
        }
        v38 := GenS4{F0: fi35}
        _ = len(v3)
        _ = fi35
        _ = v38
    }
    fmt.Println(v5, v5)
    var v39 string = "a"
    v40 := v2[42]
    var v41 uint32 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
