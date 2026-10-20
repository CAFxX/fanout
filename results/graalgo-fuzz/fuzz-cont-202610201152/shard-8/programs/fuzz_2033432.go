package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 uint16 }

func main() {
    v1 := make(map[int]bool)
    var v2 bool = (!(1 < 32767))
    var v3 uint16 = 220
    var v4 int16 = 42
    if false {
        var v5 float64 = 0.0
        _ = v5
    }
    v4 = (v4 << ((42) & 63))
    fmt.Println(v4, v2, v4)
    for fi6 := 0; fi6 < 18; fi6++ {
        v3 = v3
        _ = fi6
    }
    v3 += v3
    for fi7 := 0; fi7 < 11; fi7++ {
        v1[fi7] = false
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 uint8 = 1
            v1[128] = (-2147483648 <= 1)
            _ = fi8
            _ = v9
        }
        var v10 int32 = -1
        _ = fi7
        _ = v10
    }
    var v11 uint64 = 131
    var v12 uint64 = (v11 / ((v11) | 1))
    if v2 {
        if v2 {
            for fi13 := 0; fi13 < 15; fi13++ {
                _ = len(v1)
                v1[int(v12)] = (!(-122 <= -69))
                v14 := GenS1{F0: (v12 - v11)}
                _ = fi13
                _ = v14
            }
            fmt.Println(v11, v12, v12)
        }
        var v15 uint16 = 1
        var v16 uint = 0
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        for fi18 := 0; fi18 < 7; fi18++ {
            var v19 uint32 = 0
            _ = fi18
            _ = v19
        }
        switch "foo bar" {
        case "foo bar":
            _ = len(v1)
        default:
            for fi20 := 0; fi20 < 4; fi20++ {
                switch fi20 {
                case 128:
                    v4 -= -710
                case 127:
                    switch "hello" {
                    case "":
                        var v21 uint16 = v3
                        _ = v21
                    case "a":
                        fmt.Println(v4, fi20)
                    case "hello":
                        if v2 {
                            v22 := GenS1{F0: v12}
                            if ("" < "a") {
                                if (!v2) {
                                    var v23 string = ("foo bar" + "a")
                                    if v2 {
                                        var v24 uint = uint(v3)
                                        v12 = 18446744073709551615
                                        if (!v2) {
                                            v25 := []int{(fi20 % (((fi17 >> (((v24 / ((65535) | 1))) & 63))) | 1)), fi20, fi17}
                                            fmt.Println(v3, v24)
                                            _ = v25
                                        } else {
                                            v4 = v4
                                        }
                                        _ = v24
                                    }
                                    v1[-2147483648] = (("" + "\n\t") > ((v23 + (v23 + ("\n\t" + "foo bar"))) + v23))
                                    _ = v23
                                }
                                var v26 int = int(v4)
                                _ = v26
                            }
                            var v27 float64 = 3.14159
                            _ = v22
                            _ = v27
                        }
                        if v2 {
                            v2 = (!false)
                        }
                        v2 = v2
                    }
                }
                v28 := []int{(fi20 % ((-9223372036854775808) | 1)), fi20, fi17, (fi17 | fi20)}
                _ = fi20
                _ = v28
            }
        }
        if v2 {
            for fi29 := 0; fi29 < 17; fi29++ {
                if (false && true) {
                    delete(v1, (fi17 >> ((18446744073709551615) & 63)))
                    if v2 {
                        v12 = (v12 % ((944) | 1))
                        switch (fi29 % (((^(fi29))) | 1)) {
                        case 936:
                            var v30 int64 = 1
                            _ = v30
                        default:
                            if v2 {
                                switch (-(fi17)) {
                                case 1:
                                    v31 := []int{(fi17 + fi17), (fi17 % (((fi17 % ((fi29) | 1))) | 1)), (fi17 - 128), fi17}
                                    _ = v31
                                case 4294967295:
                                    v32 := GenS2{F0: 31, F1: 255}
                                    _ = v32
                                default:
                                    fmt.Println(v3, v2, fi29)
                                }
                                v12 = v11
                                fmt.Println(v4, fi17, fi29)
                            } else {
                                switch (fi29 << ((1) & 63)) {
                                case 1:
                                    v33 := []int{}
                                    _ = v33
                                case 128:
                                    v34 := GenS3{F0: v4}
                                    _ = v34
                                }
                                switch fi17 {
                                case -636:
                                    var v35 int32 = 2147483647
                                    _ = v35
                                default:
                                    var v36 float32 = 1e10
                                    _ = v36
                                }
                                if v2 {
                                    v37 := GenS4{F0: 375}
                                    _ = v37
                                }
                                var v38 int64 = 170
                                v4 -= int16(v38)
                                _ = v38
                            }
                        }
                        var v39 int = 32768
                        _ = v39
                    }
                } else {
                    switch "\n\t" {
                    case "a":
                        var v40 uint = 1
                        _ = v40
                    case "\n\t":
                        v41 := GenS5{F0: 1, F1: 32767}
                        var v42 int32 = 1
                        _ = v41
                        _ = v42
                    }
                }
                if true {
                    var v43 bool = v2
                    _ = v43
                }
                v44 := []int{fi17, int(v11), (fi29 % ((723) | 1)), (-(fi17))}
                _ = fi29
                _ = v44
            }
            fmt.Println(v12)
        }
        _ = fi17
    }
    if ((1e10 < 1.5) && v2) {
        v12 = (v11 / ((42) | 1))
        for fi45 := 0; fi45 < 1; fi45++ {
            fmt.Println(v12)
            if v2 {
                v46 := GenS4{F0: 739}
                _ = v46
            } else {
                if v2 {
                    var v47 int = (^(fi45))
                    _ = v47
                } else {
                    for fi48 := 0; fi48 < 1; fi48++ {
                        var v49 float64 = ((float64(1e10) * 1.5) * (1e10 - 1.5))
                        _ = fi48
                        _ = v49
                    }
                    var v50 float32 = 3.14159
                    _ = v50
                }
            }
            var v51 int64 = 682
            _ = fi45
            _ = v51
        }
    }
    var v52 float32 = (-2.25 * -2.25)
    var v53 int = 996
    var v54 uint64 = (v12 ^ 1)
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v53
    _ = v54
}
