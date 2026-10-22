package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool; F1 uint32; F2 int }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 255
    v3 := []int{42, -2147483648, -1}
    var v4 uint8 = 255
    var v5 int16 = -79
    var v6 int8 = 24
    var v7 float32 = 0.0
    switch 0 {
    case 4294967295:
        if (true || (true && true)) {
            if (true || ("a" == "foo bar")) {
                fmt.Println(v5, v7, v2)
            } else {
                _ = len(v1)
            }
            _ = len(v1)
        }
    default:
        switch 43 {
        case 2147483647:
            v8 := GenS1{F0: 42, F1: v2}
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        fmt.Println(fi9, fi9)
        v3 = append(v3, fi9)
        _ = fi9
    }
    delete(v1, "hello")
    var v10 bool = (!false)
    v2 = (v2 << ((v2) & 63))
    v6 -= v6
    var v11 uint = 0
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 6; fi13++ {
            _ = len(v3)
            v5 = (v5 + v5)
            v3 = append(v3, -1)
            _ = fi13
        }
        if false {
            var v14 int16 = (v5 << (((v11 % (((^(v11))) | 1))) & 63))
            v15 := GenS1{F0: -1, F1: 976}
            v16 := GenS2{F0: -9223372036854775808, F1: 0}
            _ = v14
            _ = v15
            _ = v16
        }
        _ = fi12
    }
    v17 := GenS2{F0: -9223372036854775808, F1: -2147483648}
    var v18 string = "foo bar"
    v10 = v10
    v10 = v10
    if (4294967295 == 4294967295) {
        var v19 int8 = 36
        for fi20 := 0; fi20 < 14; fi20++ {
            v21 := GenS3{F0: (fi20 * fi20)}
            var v22 uint32 = 880
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = v19
    } else {
        v4 += (v4 ^ (v4 & (v4 << ((v2) & 63))))
    }
    if v10 {
        v3 = append(v3, 32768)
        if true {
            if v10 {
                fmt.Println(v2, v10)
                if v10 {
                    var v23 uint8 = (^(v4))
                    switch -766 {
                    case 2147483647:
                        for fi24 := 0; fi24 < 2; fi24++ {
                            if false {
                                for fi25 := 0; fi25 < 10; fi25++ {
                                    v7 = 0.1
                                    v7 -= v7
                                    _ = fi25
                                }
                                v26 := GenS4{F0: 65535}
                                _ = v26
                            } else {
                                _ = len(v1)
                                _ = len(v1)
                                v3 = append(v3, fi24)
                            }
                            if v10 {
                                v27 := GenS4{F0: 331}
                                _ = v27
                            } else {
                                v5 = int16(fi24)
                            }
                            _ = fi24
                        }
                    case -780:
                        if v10 {
                            if (!(986 == -2147483648)) {
                                if (1 > v2) {
                                    var v28 int = 42
                                    _ = v28
                                }
                                for fi29 := 0; fi29 < 20; fi29++ {
                                    var v30 int32 = 44
                                    v31 := GenS1{F0: -1, F1: (v11 | (v2 - (v11 - (^(v11)))))}
                                    for fi32 := 0; fi32 < 9; fi32++ {
                                        fmt.Println(v4, v11, v23)
                                        v33 := GenS3{F0: 2147483647}
                                        switch v18 {
                                        case "\n\t":
                                            v34 := GenS4{F0: 104}
                                            _ = v34
                                        }
                                        _ = fi32
                                        _ = v33
                                    }
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                }
                                v11 += v2
                            }
                        } else {
                            v35 := GenS3{F0: -912}
                            var v36 uint = 4294967295
                            _ = v35
                            _ = v36
                        }
                    case 9223372036854775807:
                        var v37 uint8 = v23
                        _ = v37
                    }
                    _ = v23
                } else {
                    for fi38 := 0; fi38 < 7; fi38++ {
                        v2 = (v2 % (((v2 / (((v11 / ((18446744073709551615) | 1))) | 1))) | 1))
                        if true {
                            fmt.Println(v6)
                            var v39 int32 = 42
                            v7 = v7
                            _ = v39
                        } else {
                            for fi40 := 0; fi40 < 17; fi40++ {
                                v4 = (^(v4))
                                var v41 int = 285
                                fmt.Println(v10)
                                _ = fi40
                                _ = v41
                            }
                            for fi42 := 0; fi42 < 3; fi42++ {
                                v6 = (v6 >> (((v2 % ((v11) | 1))) & 63))
                                v6 -= v6
                                fmt.Println(fi42, v6, v11)
                                _ = fi42
                            }
                            for fi43 := 0; fi43 < 12; fi43++ {
                                v5 = -576
                                var v44 int16 = v5
                                _ = fi43
                                _ = v44
                            }
                        }
                        _ = fi38
                    }
                    switch (("hello" + (v18 + ("a" + ""))) + "foo bar") {
                    case "":
                        v10 = false
                    case "hello":
                        if (((0 > 65535) || (!true)) || (444 != 124)) {
                            for fi45 := 0; fi45 < 18; fi45++ {
                                fmt.Println(v6)
                                if v10 {
                                    var v46 uint64 = 694
                                    _ = v46
                                }
                                _ = fi45
                            }
                        }
                    case "\n\t":
                        v11 = 582
                    }
                }
                _ = len(v3)
            }
            v7 = (-2.25 - v7)
        }
        var v47 float64 = float64(((1e10 / float32(((0.1 - 1.5) * float64(v7)))) + 0.0))
        _ = v47
    } else {
        var v48 float64 = 1e10
        _ = v48
    }
    v49 := GenS5{F0: (-1 <= int32(v6)), F1: 4294967295, F2: -581}
    fmt.Println(v2, v5)
    switch v18 {
    case "":
        var v50 uint64 = 2
        _ = v50
    default:
        _ = len(v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v7
}
