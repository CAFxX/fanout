package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint; F1 bool; F2 uint32 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    var v3 int = 32768
    var v4 float32 = 3.14159
    if true {
        for fi5 := 0; fi5 < 16; fi5++ {
            if (84 <= 561) {
                if (18446744073709551615 <= 42) {
                    fmt.Println(v3, fi5)
                    if true {
                        var v6 int = (-(fi5))
                        _ = v6
                    } else {
                        v7 := []int{}
                        fmt.Println(v3, fi5, v3)
                        _ = v7
                    }
                } else {
                    switch (fi5 + fi5) {
                    case -1:
                        switch "\n\t" {
                        case "":
                            var v8 int64 = 622
                            _ = v8
                        case "foo bar":
                            _ = len(v2)
                        case "hello":
                            _ = len(v2)
                        }
                    case 127:
                        v9 := GenS1{F0: -9223372036854775808}
                        _ = v9
                    case 1:
                        if (!false) {
                            v3 = -544
                        }
                    default:
                        if (!false) {
                            fmt.Println(fi5)
                            v4 += v4
                            var v10 float64 = 0.1
                            _ = v10
                        }
                    }
                    var v11 int8 = -51
                    if true {
                        switch (((("" + "") + (("hello" + "hello") + "\n\t")) + "\n\t") + "\n\t") {
                        case "foo bar":
                            switch "foo bar" {
                            case "hello":
                                for fi12 := 0; fi12 < 14; fi12++ {
                                    v13 := []int{}
                                    _ = len(v1)
                                    _ = fi12
                                    _ = v13
                                }
                            case "foo bar":
                                var v14 int32 = 1
                                _ = v14
                            case "\n\t":
                                var v15 uint32 = 665
                                _ = v15
                            default:
                                var v16 uint64 = 1
                                _ = v16
                            }
                        case "\n\t":
                            var v17 uint16 = 0
                            _ = v17
                        }
                        v18 := []int{}
                        _ = v18
                    }
                    _ = v11
                }
                var v19 uint16 = 2
                if false {
                    v4 = ((1.5 - float32(0.1)) * 1.5)
                    switch (fi5 << ((18446744073709551615) & 63)) {
                    case 127:
                        v20 := GenS2{F0: 18446744073709551615, F1: -1}
                        _ = v20
                    case 2147483647:
                        if (3.14159 >= 1e10) {
                            v19 += 1
                            v19 = 1
                            for fi21 := 0; fi21 < 6; fi21++ {
                                if (32767 < 1000) {
                                    _ = len(v1)
                                    var v22 int = fi21
                                    _ = v22
                                }
                                _ = fi21
                            }
                        }
                    default:
                        v3 = fi5
                    }
                }
                _ = v19
            } else {
                if false {
                    var v23 int16 = 275
                    var v24 string = ("hello" + "hello")
                    _ = v23
                    _ = v24
                } else {
                    var v25 bool = (v4 == v4)
                    _ = v25
                }
            }
            _ = fi5
        }
    }
    delete(v2, 0)
    var v26 int32 = 93
    switch v3 {
    case 224:
        v27 := v2[2]
        _ = v27
    case 128:
        var v28 string = ((("foo bar" + "a") + "a") + "foo bar")
        _ = v28
    }
    var v29 bool = true
    switch ("hello" + "\n\t") {
    case "foo bar":
        var v30 uint16 = 671
        _ = v30
    default:
        fmt.Println(v4)
    }
    if v29 {
        v29 = v29
        v4 = 1e10
    } else {
        delete(v1, (v3 % (((v3 >> ((733) & 63))) | 1)))
        for fi31 := 0; fi31 < 12; fi31++ {
            if v29 {
                switch v3 {
                case -826:
                    var v32 uint32 = 798
                    _ = v32
                case 545:
                    for fi33 := 0; fi33 < 2; fi33++ {
                        var v34 int32 = v26
                        if (42 < 65535) {
                            v4 += (1e10 / 1.0)
                        }
                        var v35 int32 = v34
                        var v36 int64 = 9223372036854775807
                        _ = fi33
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                case 4294967295:
                    v37 := GenS3{F0: 744, F1: false, F2: 594}
                    _ = v37
                }
            } else {
                var v38 uint = 964
                v26 += (v26 * 2147483647)
                _ = v38
            }
            _ = fi31
        }
        v4 += v4
    }
    v39 := GenS1{F0: 0}
    v40 := []int{}
    var v41 string = "foo bar"
    switch v41 {
    case "foo bar":
        if true {
            v26 -= v26
            fmt.Println(v41, v26, v3)
            v42 := GenS3{F0: 109, F1: false, F2: 446}
            _ = v42
        } else {
            var v43 uint64 = 4294967295
            if (!v29) {
                v43 = (v43 / ((906) | 1))
                v44 := GenS1{F0: 358}
                v3 = 460
                _ = v44
            }
            v41 = "foo bar"
            _ = v43
        }
    case "hello":
        fmt.Println(v41, v29, v4)
    }
    fmt.Println(v4)
    delete(v2, 4294967295)
    v41 = v41
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
