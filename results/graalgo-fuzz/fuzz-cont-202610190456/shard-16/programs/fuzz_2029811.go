package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 uint32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{42, 2147483648, 0, 1}
    var v3 float32 = 0.0
    v4 := make(map[int]bool)
    v5 := make(map[int]string)
    var v6 uint = 42
    var v7 uint = 4294967295
    fmt.Println(v7)
    v7 -= (v7 ^ v7)
    for fi8 := 0; fi8 < 2; fi8++ {
        _ = len(v2)
        _ = fi8
    }
    v9 := GenS1{F0: (!(43 < 145)), F1: 1, F2: 913}
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 int16 = -32768
        _ = fi10
        _ = v11
    }
    if true {
        for fi12 := 0; fi12 < 11; fi12++ {
            v3 = float32(0.0)
            fmt.Println(fi12, v6)
            for fi13 := 0; fi13 < 1; fi13++ {
                switch (fi13 << (((v7 & (v6 << ((v6) & 63)))) & 63)) {
                case 0:
                    switch "" {
                    case "":
                        v3 -= v3
                    case "a":
                        var v14 uint64 = 2
                        _ = v14
                    case "hello":
                        if false {
                            switch "" {
                            case "":
                                switch fi13 {
                                case -9223372036854775808:
                                    v7 = v7
                                case 420:
                                    switch "\n\t" {
                                    case "\n\t":
                                        switch ("a" + "hello") {
                                        case "\n\t":
                                            var v15 int8 = 100
                                            _ = v15
                                        default:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                    default:
                                        var v16 int8 = 0
                                        _ = v16
                                    }
                                default:
                                    _ = len(v5)
                                }
                            }
                            v5[-452] = "\n\t"
                            fmt.Println(fi13)
                        }
                    default:
                        v3 = v3
                    }
                case -1:
                    var v17 float64 = -2.25
                    _ = v17
                }
                v3 -= 3.14159
                if false {
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(v7, v7, v6)
                    case "hello":
                        v3 = (0.0 / (float32((((0.1 / float64((v3 * -2.25))) * -2.25) - (0.1 - 0.1))) + v3))
                    }
                }
                _ = fi13
            }
            _ = fi12
        }
        if ("\n\t" < "foo bar") {
            if true {
                var v18 int16 = 187
                v18 = 1
                _ = v18
            } else {
                var v19 int64 = 0
                v20 := GenS2{F0: v6}
                var v21 int64 = (v19 & (v19 ^ 0))
                _ = v19
                _ = v20
                _ = v21
            }
            var v22 uint16 = 136
            if (42 <= 0) {
                var v23 uint16 = v22
                _ = v23
            } else {
                v24 := GenS3{F0: 2147483647, F1: 42}
                if true {
                    var v25 int8 = 42
                    switch -9223372036854775808 {
                    case 2147483647:
                        if (42 >= (v7 ^ v7)) {
                            v22 -= (v22 | v22)
                            v3 -= v3
                        } else {
                            var v26 string = "a"
                            v25 = (v25 * v25)
                            v22 -= (v22 + v22)
                            _ = v26
                        }
                    case 42:
                        v27 := GenS3{F0: 2147483648, F1: 2}
                        _ = v27
                    case -359:
                        for fi28 := 0; fi28 < 9; fi28++ {
                            var v29 int32 = -901
                            for fi30 := 0; fi30 < 17; fi30++ {
                                v22 = 224
                                var v31 float64 = (1.5 - ((0.1 - 0.0) * -2.25))
                                _ = len(v2)
                                _ = fi30
                                _ = v31
                            }
                            for fi32 := 0; fi32 < 12; fi32++ {
                                v3 -= v3
                                v6 += (^(v6))
                                v4[(^(fi32))] = ("hello" >= "a")
                                _ = fi32
                            }
                            _ = fi28
                            _ = v29
                        }
                    }
                    _ = v25
                } else {
                    for fi33 := 0; fi33 < 15; fi33++ {
                        if false {
                            var v34 float32 = (v3 - (0.1 + v3))
                            for fi35 := 0; fi35 < 3; fi35++ {
                                var v36 string = "foo bar"
                                var v37 uint = v6
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                            _ = v34
                        }
                        var v38 uint = v7
                        if ((v3 != (float32(0.1) + (1.5 - (v3 * 0.1)))) && (668 != 65535)) {
                            for fi39 := 0; fi39 < 17; fi39++ {
                                v40 := GenS3{F0: 2147483647, F1: 2147483647}
                                var v41 float32 = v3
                                _ = fi39
                                _ = v40
                                _ = v41
                            }
                        }
                        _ = fi33
                        _ = v38
                    }
                    fmt.Println(v6, v3)
                    _ = len(v2)
                }
                _ = v24
            }
            _ = v22
        }
        var v42 float32 = v3
        _ = v42
    }
    var v43 int64 = -1
    var v44 uint32 = 2147483647
    var v45 string = ("" + "foo bar")
    if (!true) {
        v45 = "\n\t"
    }
    fmt.Println(v45, v3)
    var v46 uint = 957
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
