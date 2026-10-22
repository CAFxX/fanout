package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 int; F2 int8 }
type GenS6 struct { F0 int32 }

func main() {
    var v1 uint = 738
    var v2 int8 = 65
    v3 := make(map[int]int)
    var v4 uint8 = 1
    var v5 int = 32767
    fmt.Println(v5, v2, v4)
    var v6 int16 = 1
    var v7 int32 = -2147483648
    v1 = 0
    var v8 int = (v5 + -170)
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 string = ""
        var v11 int8 = v2
        v12 := GenS1{F0: ((^(v1)) == v1)}
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    v4 += v4
    for fi13 := 0; fi13 < 14; fi13++ {
        v8 -= 79
        if ((((1.5 / 1.0) - (3.14159 - -2.25)) * 0.1) != ((0.1 + -2.25) + (0.0 * (0.1 * 3.14159)))) {
            fmt.Println(v1, v4)
        } else {
            switch ("a" + "\n\t") {
            case "":
                if ((0.0 + 0.1) <= 0.0) {
                    v14 := GenS2{F0: 4294967295, F1: -1}
                    var v15 uint = (v1 ^ v1)
                    for fi16 := 0; fi16 < 18; fi16++ {
                        var v17 uint32 = 2
                        _ = fi16
                        _ = v17
                    }
                    _ = v14
                    _ = v15
                }
            case "\n\t":
                if (0 >= 2147483648) {
                    v1 += v1
                    v18 := []int{-291, fi13, (v5 << ((255) & 63)), (v5 | v8)}
                    v2 += v2
                    _ = v18
                }
            default:
                v19 := GenS3{F0: v4}
                _ = v19
            }
            v20 := v3[(v8 / (((-(v5))) | 1))]
            _ = v20
        }
        var v21 uint8 = (v4 - 235)
        _ = fi13
        _ = v21
    }
    v1 = (v1 % ((v1) | 1))
    switch v8 {
    case -59:
        v6 = v6
    default:
        var v22 int8 = v2
        _ = v22
    }
    if (2 >= 366) {
        if (v8 < (-(v8))) {
            if false {
                fmt.Println(v8, v8)
            }
            switch 698 {
            case 32767:
                v1 -= 0
            case -52:
                v23 := GenS1{F0: false}
                _ = v23
            default:
                for fi24 := 0; fi24 < 13; fi24++ {
                    var v25 bool = (1 >= v1)
                    _ = fi24
                    _ = v25
                }
            }
            fmt.Println(v2, v8, v4)
        } else {
            v26 := GenS2{F0: 42, F1: 0}
            v5 = v5
            switch (^(v5)) {
            case 107:
                v27 := []int{}
                _ = v27
            case -2147483648:
                var v28 int64 = 0
                _ = v28
            case 2147483647:
                v2 -= (v2 >> (((^(v1))) & 63))
            }
            _ = v26
        }
        for fi29 := 0; fi29 < 8; fi29++ {
            v4 = v4
            if (!false) {
                var v30 string = "hello"
                if (false && (v2 >= v2)) {
                    if ((fi29 <= v8) && true) {
                        var v31 float32 = 1.5
                        var v32 int32 = v7
                        var v33 int = (v5 / ((v5) | 1))
                        _ = v31
                        _ = v32
                        _ = v33
                    } else {
                        if (1 >= 42) {
                            v7 = v7
                            v7 -= 42
                            v30 = v30
                        } else {
                            v34 := GenS4{F0: (fi29 << ((v1) & 63))}
                            v35 := []int{}
                            _ = v34
                            _ = v35
                        }
                        v6 = (v6 & 865)
                        for fi36 := 0; fi36 < 19; fi36++ {
                            var v37 int = (fi36 ^ v5)
                            if true {
                                switch ((v30 + v30) + v30) {
                                case "foo bar":
                                    var v38 int8 = v2
                                    _ = v38
                                case "a":
                                    _ = len(v3)
                                case "hello":
                                    var v39 int16 = (-(v6))
                                    _ = v39
                                }
                                var v40 float32 = 1.5
                                v7 += int32(v6)
                                v1 -= 964
                                v41 := GenS4{F0: fi29}
                                _ = v40
                                _ = v41
                            } else {
                                var v42 float32 = 1e10
                                v2 = (v2 + v2)
                                _ = v42
                            }
                            _ = fi36
                            _ = v37
                        }
                    }
                    switch (((((v30 + ((v30 + ("\n\t" + ("\n\t" + v30))) + v30)) + "hello") + v30) + "") + v30) {
                    case "":
                        var v43 int16 = v6
                        _ = v43
                    case "foo bar":
                        var v44 bool = ((v1 % ((uint(v4)) | 1)) > 1)
                        _ = v44
                    }
                } else {
                    var v45 bool = (-2.25 > 1.5)
                    v46 := GenS4{F0: v8}
                    _ = v45
                    _ = v46
                }
                var v47 uint = v1
                _ = v30
                _ = v47
            } else {
                var v48 uint16 = 65535
                _ = v48
            }
            if true {
                v49 := GenS5{F0: false, F1: (-(v8)), F2: 0}
                if true {
                    switch "hello" {
                    case "a":
                        if (v4 >= 255) {
                            v50 := GenS6{F0: (v7 + 0)}
                            v51 := GenS6{F0: (^(v7))}
                            _ = v50
                            _ = v51
                        } else {
                            v52 := []int{}
                            _ = v52
                        }
                    case "\n\t":
                        fmt.Println(v7, v5, v4)
                    case "foo bar":
                        v4 += 42
                    }
                } else {
                    var v53 float32 = (-2.25 / -2.25)
                    v53 += 0.1
                    _ = v53
                }
                var v54 uint8 = (v4 / (((v4 / (((v4 ^ 0)) | 1))) | 1))
                _ = v49
                _ = v54
            }
            _ = fi29
        }
        var v55 int64 = 512
        _ = v55
    } else {
        var v56 int32 = (v7 << ((255) & 63))
        for fi57 := 0; fi57 < 16; fi57++ {
            var v58 uint32 = 304
            var v59 int16 = int16(v7)
            switch (-(v5)) {
            case -957:
                fmt.Println(v6, v56)
            }
            _ = fi57
            _ = v58
            _ = v59
        }
        v60 := GenS6{F0: (-(v56))}
        _ = v56
        _ = v60
    }
    var v61 float32 = (((0.0 / ((1e10 * ((1e10 - 0.1) * (0.1 * (3.14159 * (float32(1.5) * 3.14159))))) * (((-2.25 / -2.25) - 1e10) - 1.5))) * (-2.25 * float32(1e10))) * (1e10 + -2.25))
    fmt.Println(v1, v5, v5)
    if (-2.25 < (float32(float64(v61)) + (0.1 - 0.0))) {
        v8 = v5
    }
    var v62 int32 = (v7 | (v7 % ((v7) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v62
    _ = v7
    _ = v8
}
