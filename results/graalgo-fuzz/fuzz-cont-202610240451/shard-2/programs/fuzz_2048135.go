package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int32; F2 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 string; F2 bool }

func main() {
    var v1 int16 = -32768
    var v2 float64 = 3.14159
    var v3 int16 = 1000
    var v4 string = "\n\t"
    for fi5 := 0; fi5 < 2; fi5++ {
        for fi6 := 0; fi6 < 5; fi6++ {
            fmt.Println(v4, fi5, fi5)
            v4 = v4
            _ = fi6
        }
        fmt.Println(v4, v2)
        for fi7 := 0; fi7 < 20; fi7++ {
            v3 = int16(fi5)
            v4 = "hello"
            var v8 uint64 = 1
            switch "hello" {
            case "foo bar":
                switch v4 {
                case "\n\t":
                    v9 := GenS1{F0: 2147483647}
                    _ = v9
                case "foo bar":
                    v2 += 0.1
                case "a":
                    var v10 int16 = (v3 | -487)
                    _ = v10
                }
            case "hello":
                switch 42 {
                case 273:
                    v11 := []int{fi5, (fi5 & 32768), 32768}
                    _ = v11
                case -103:
                    var v12 float64 = v2
                    _ = v12
                }
            case "":
                switch (((v4 + (("" + "foo bar") + v4)) + (v4 + v4)) + "hello") {
                case "foo bar":
                    if (!(!true)) {
                        if true {
                            var v13 uint16 = 22
                            var v14 int32 = -381
                            _ = v13
                            _ = v14
                        } else {
                            fmt.Println(v4, v3, v4)
                        }
                    } else {
                        switch int(v3) {
                        case 0:
                            v15 := GenS2{F0: -1, F1: 0, F2: 528}
                            _ = v15
                        case 9223372036854775807:
                            v16 := GenS1{F0: 2147483648}
                            _ = v16
                        default:
                            if (v2 > v2) {
                                v17 := GenS2{F0: 42, F1: -1, F2: (v1 << ((18446744073709551615) & 63))}
                                _ = v17
                            }
                            v18 := []int{(fi7 + 128), (fi7 ^ fi7), 2147483647, 2147483647}
                            if true {
                                v2 = (v2 / float64(3.14159))
                                v4 = v4
                            } else {
                                var v19 uint64 = 93
                                fmt.Println(fi5)
                                _ = v19
                            }
                            _ = v18
                        }
                        v20 := []int{fi5, (fi7 / ((2147483647) | 1))}
                        _ = v20
                    }
                case "":
                    v4 = (v4 + v4)
                case "a":
                    v3 -= v3
                default:
                    if ((-671 == 42) || ((127 < 2) && (true && (false && ((160 < 128) && true))))) {
                        switch "foo bar" {
                        case "hello":
                            if true {
                                if true {
                                    v1 -= v1
                                    var v21 uint8 = 127
                                    var v22 int64 = 1
                                    _ = v21
                                    _ = v22
                                }
                                if ((true && (v4 <= v4)) && false) {
                                    v23 := []int{(-(fi7)), fi7}
                                    v24 := GenS3{F0: 469}
                                    var v25 int16 = 42
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                            }
                        default:
                            var v26 uint = 532
                            _ = v26
                        }
                        v27 := GenS2{F0: 2147483647, F1: 2147483647, F2: (v1 >> ((18446744073709551615) & 63))}
                        v28 := []int{}
                        _ = v27
                        _ = v28
                    }
                }
            }
            if (((float64((-2.25 + 1.5)) * -2.25) + v2) != v2) {
                fmt.Println(v4, v4, v8)
                v8 = (v8 ^ 378)
                v29 := []int{(-(fi7))}
                _ = v29
            }
            _ = fi7
            _ = v8
        }
        _ = fi5
    }
    v30 := GenS3{F0: 127}
    v4 = v4
    switch "\n\t" {
    case "hello":
        if true {
            var v31 bool = false
            v31 = (-2147483648 != -900)
            v32 := []int{775, 0, -2147483648, 4294967295}
            _ = v31
            _ = v32
        } else {
            v33 := []int{}
            var v34 uint8 = 1
            _ = v33
            _ = v34
        }
    case "":
        var v35 float32 = 0.1
        _ = v35
    case "\n\t":
        switch v4 {
        case "foo bar":
            v36 := GenS4{F0: 97, F1: ((v4 + v4) + "foo bar"), F2: (false || false)}
            _ = v36
        }
    }
    v4 = (v4 + "foo bar")
    if true {
        var v37 int = -423
        switch (v4 + v4) {
        case "foo bar":
            v4 = ("hello" + ("" + v4))
        case "a":
            v38 := []int{(v37 % ((-1) | 1)), 42, v37}
            _ = v38
        default:
            var v39 uint32 = 779
            _ = v39
        }
        if true {
            for fi40 := 0; fi40 < 19; fi40++ {
                if ((0.1 + (1.5 + 1e10)) >= v2) {
                    switch "" {
                    case "foo bar":
                        var v41 int64 = 1
                        _ = v41
                    case "hello":
                        var v42 bool = (!(!(v4 > v4)))
                        _ = v42
                    case "\n\t":
                        var v43 float64 = v2
                        _ = v43
                    }
                    fmt.Println(v2)
                }
                _ = fi40
            }
        }
        _ = v37
    }
    v44 := []int{2147483647}
    if len(v44) >= 2 {
        v44 = v44[1:len(v44)]
    }
    var v45 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v44
    _ = v45
}
