package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 uint8; F1 int16 }

func main() {
    v1 := make(map[int]int)
    var v2 string = "hello"
    v3 := make(map[int]int)
    var v4 int32 = 1
    var v5 uint64 = 2
    v3[4294967295] = -303
    for fi6 := 0; fi6 < 8; fi6++ {
        v5 += 195
        _ = fi6
    }
    var v7 int32 = 42
    delete(v3, 32768)
    _ = len(v1)
    _ = len(v1)
    v8 := []int{-1, 163, 2147483647}
    v7 += (v4 / (((v7 * v7)) | 1))
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 uint64 = (v5 + (v5 ^ v5))
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: ((((("foo bar" + (((v2 + v2) + v2) + (("" + "a") + v2))) + (v2 + (v2 + ""))) + (v2 + v2)) + v2) + v2)}
    for fi12 := 0; fi12 < 17; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            v14 := GenS1{F0: v2}
            var v15 int32 = (v4 >> ((408) & 63))
            if false {
                var v16 uint16 = 439
                _ = v16
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 int8 = -128
        for fi19 := 0; fi19 < 10; fi19++ {
            switch fi19 {
            case 899:
                var v20 string = "foo bar"
                _ = v20
            case 4294967295:
                v5 += 96
            default:
                v2 = v2
            }
            _ = fi19
        }
        fmt.Println(v5)
        _ = fi17
        _ = v18
    }
    if true {
        if false {
            v3[-97] = 639
            switch v2 {
            case "a":
                _ = len(v3)
            case "":
                _ = len(v8)
            case "foo bar":
                fmt.Println(v7, v2)
            }
            var v21 uint64 = (^(v5))
            _ = v21
        } else {
            if true {
                v8 = append(v8, 9223372036854775807)
                for fi22 := 0; fi22 < 15; fi22++ {
                    fmt.Println(v7)
                    v23 := GenS1{F0: v2}
                    v7 -= 219
                    _ = fi22
                    _ = v23
                }
                if (!true) {
                    for fi24 := 0; fi24 < 6; fi24++ {
                        for fi25 := 0; fi25 < 5; fi25++ {
                            switch -52 {
                            case -1:
                                v3[(fi25 + (fi24 & (fi25 | 986)))] = (^(fi24))
                                v2 = v2
                            case 31:
                                var v26 uint = 670
                                _ = v26
                            default:
                                var v27 uint = 230
                                _ = v27
                            }
                            fmt.Println(v4, fi24, v2)
                            _ = fi25
                        }
                        _ = fi24
                    }
                    switch (v2 + ("a" + (v2 + v2))) {
                    case "a":
                        if true {
                            for fi28 := 0; fi28 < 13; fi28++ {
                                switch (^(fi28)) {
                                case 8:
                                    if (42 != -1) {
                                        var v29 int = (fi28 % ((-2147483648) | 1))
                                        var v30 int = v29
                                        for fi31 := 0; fi31 < 10; fi31++ {
                                            var v32 int16 = 1
                                            switch ((("foo bar" + v2) + (v2 + "\n\t")) + v2) {
                                            case "hello":
                                                v8 = append(v8, 42)
                                            case "\n\t":
                                                v33 := GenS2{F0: -135}
                                                _ = v33
                                            default:
                                                if (!true) {
                                                    var v34 int8 = 100
                                                    var v35 uint = 0
                                                    _ = v34
                                                    _ = v35
                                                }
                                            }
                                            if (!((v7 % ((724) | 1)) < v7)) {
                                                v36 := GenS1{F0: ("hello" + (v2 + "a"))}
                                                var v37 bool = false
                                                var v38 int16 = (v32 >> ((0) & 63))
                                                v29 = (fi28 << ((62) & 63))
                                                _ = v36
                                                _ = v37
                                                _ = v38
                                            }
                                            _ = fi31
                                            _ = v32
                                        }
                                        _ = v29
                                        _ = v30
                                    } else {
                                        v39 := GenS1{F0: ("\n\t" + (v2 + (v2 + "\n\t")))}
                                        _ = v39
                                    }
                                default:
                                    switch fi28 {
                                    case 9223372036854775807:
                                        fmt.Println(v2, v2, fi28)
                                    case 2147483647:
                                        var v40 uint16 = 652
                                        _ = v40
                                    case -2147483648:
                                        for fi41 := 0; fi41 < 7; fi41++ {
                                            v1[fi28] = fi28
                                            v2 = ("\n\t" + "a")
                                            _ = fi41
                                        }
                                    default:
                                        _ = len(v8)
                                    }
                                }
                                _ = fi28
                            }
                            v4 += 0
                            var v42 int = 4294967295
                            _ = v42
                        } else {
                            v43 := GenS3{F0: 3, F1: 2}
                            switch 1 {
                            case -9223372036854775808:
                                v5 += v5
                            case 1:
                                var v44 uint64 = v5
                                _ = v44
                            default:
                                v4 = (v7 + v7)
                            }
                            v45 := GenS4{F0: uint8(v5), F1: 32768}
                            _ = v43
                            _ = v45
                        }
                    default:
                        v4 += v4
                    }
                }
            }
            var v46 uint64 = v5
            v5 -= (v5 % ((v46) | 1))
            _ = v46
        }
    }
    if (true || (2 >= 196)) {
        var v47 uint8 = 42
        var v48 int64 = 0
        _ = v47
        _ = v48
    } else {
        if false {
            v5 += (v5 << ((0) & 63))
        } else {
            v49 := GenS5{F0: 128, F1: 1000}
            fmt.Println(v5)
            _ = v49
        }
        if true {
            v50 := GenS5{F0: 0, F1: 32767}
            _ = v50
        } else {
            switch v2 {
            case "hello":
                v4 = (-(v7))
            case "\n\t":
                var v51 int64 = 9223372036854775807
                _ = v51
            }
            v5 = (v5 / ((0) | 1))
        }
        if ("foo bar" != (v2 + "a")) {
            var v52 uint = 1
            _ = v52
        }
    }
    var v53 uint = 65535
    if (((1.5 + 0.1) / 3.14159) > (0.0 - 3.14159)) {
        var v54 bool = (!false)
        _ = len(v8)
        var v55 uint16 = 163
        _ = v54
        _ = v55
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v7
    _ = v8
}
