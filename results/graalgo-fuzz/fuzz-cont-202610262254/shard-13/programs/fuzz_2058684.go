package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 uint64; F2 string }
type GenS4 struct { F0 int32; F1 bool; F2 string }

func main() {
    var v1 float64 = 1e10
    var v2 uint64 = 478
    v3 := make(map[int]uint32)
    var v4 uint64 = (^(v2))
    var v5 int8 = 42
    v2 = v2
    v6 := []int{int(v5), -2147483648, -330, 4294967295}
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint8 = 223
        var v9 int = fi7
        if false {
            v1 += 0.1
            var v10 uint16 = 215
            if len(v6) > 0 {
                v6[0] = (^(v9))
            }
            _ = v10
        } else {
            v2 = 2
            var v11 bool = false
            v9 += 127
            _ = v11
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 string = (("hello" + "hello") + "a")
        switch (^(fi12)) {
        case 793:
            v14 := GenS1{F0: v4, F1: 527}
            _ = v14
        case 9223372036854775807:
            v4 = (v4 >> ((255) & 63))
        }
        _ = fi12
        _ = v13
    }
    v5 += v5
    if false {
        v1 = ((-2.25 / -2.25) + (-2.25 / float64((1e10 - 3.14159))))
    } else {
        v3[4294967295] = 1
        if false {
            fmt.Println(v5)
        } else {
            v15 := GenS2{F0: 704}
            var v16 int32 = 42
            if false {
                switch "foo bar" {
                case "":
                    v1 += 0.1
                case "a":
                    var v17 string = ""
                    _ = v17
                }
                fmt.Println(v4, v4)
            }
            _ = v15
            _ = v16
        }
        delete(v3, 932)
    }
    switch "foo bar" {
    case "foo bar":
        for fi18 := 0; fi18 < 11; fi18++ {
            v5 = v5
            _ = fi18
        }
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        v4 = v2
        _ = fi19
    }
    v2 -= 2
    if ((^(v5)) < 127) {
        for fi20 := 0; fi20 < 16; fi20++ {
            v5 -= (v5 % ((v5) | 1))
            for fi21 := 0; fi21 < 4; fi21++ {
                v4 += v2
                switch fi21 {
                case -1:
                    var v22 uint = 42
                    _ = v22
                default:
                    v23 := v3[(fi21 + fi20)]
                    _ = v23
                }
                v24 := GenS2{F0: 117}
                _ = fi21
                _ = v24
            }
            v25 := GenS3{F0: "hello", F1: v2, F2: "hello"}
            _ = fi20
            _ = v25
        }
        fmt.Println(v2, v2, v4)
        fmt.Println(v5, v1, v1)
    } else {
        for fi26 := 0; fi26 < 8; fi26++ {
            v2 -= 319
            _ = fi26
        }
    }
    var v27 int = -844
    fmt.Println(v5, v2, v1)
    v28 := GenS2{F0: 42}
    switch v27 {
    case 32768:
        if true {
            v29 := v3[v27]
            fmt.Println(v4, v2)
            for fi30 := 0; fi30 < 9; fi30++ {
                for fi31 := 0; fi31 < 11; fi31++ {
                    switch "foo bar" {
                    case "a":
                        v27 = (^(fi30))
                    case "hello":
                        fmt.Println(fi31)
                    }
                    fmt.Println(v1)
                    var v32 uint = 255
                    _ = fi31
                    _ = v32
                }
                if (!false) {
                    if (42 == 339) {
                        for fi33 := 0; fi33 < 13; fi33++ {
                            v27 = v27
                            v29 += 2147483648
                            var v34 int16 = 42
                            if false {
                                fmt.Println(v27, v34)
                                var v35 int64 = -612
                                if true {
                                    if len(v6) > 0 {
                                        v6[0] = (^(fi33))
                                    }
                                    if (fi30 <= (fi33 >> ((443) & 63))) {
                                        fmt.Println(v2)
                                        v35 -= (v35 | v35)
                                        switch -9223372036854775808 {
                                        case 4294967295:
                                            fmt.Println(v2, v29)
                                        case 137:
                                            _ = len(v6)
                                        }
                                        if false {
                                            fmt.Println(v2)
                                            var v36 bool = false
                                            fmt.Println(v34, v27, v1)
                                            _ = v36
                                        }
                                        fmt.Println(v4)
                                        v37 := GenS1{F0: v4, F1: 0}
                                        if (42 <= 237) {
                                            v38 := GenS4{F0: 1, F1: (false && false), F2: "\n\t"}
                                            _ = v38
                                        }
                                        _ = v37
                                    }
                                    fmt.Println(v1, v34, v5)
                                    var v39 uint16 = uint16(v2)
                                    switch (fi30 + int(v4)) {
                                    case 4294967295:
                                        var v40 float32 = 3.14159
                                        _ = v40
                                    case -198:
                                        if (((^(fi33)) == 42) && false) {
                                            v41 := GenS1{F0: 4294967295, F1: 0}
                                            fmt.Println(v35, v35, v39)
                                            v42 := GenS1{F0: (v2 * (v2 >> ((4294967295) & 63))), F1: (v29 ^ (v29 / ((uint32(v39)) | 1)))}
                                            _ = v41
                                            _ = v42
                                        }
                                    case 32767:
                                        var v43 float64 = 3.14159
                                        _ = v43
                                    default:
                                        v3[fi30] = 65535
                                    }
                                    _ = v39
                                }
                                _ = v35
                            }
                            _ = fi33
                            _ = v34
                        }
                        v44 := GenS2{F0: 371}
                        _ = v44
                    }
                    v6 = append(v6, v27)
                } else {
                    var v45 uint16 = uint16(v4)
                    _ = v45
                }
                fmt.Println(fi30, v5)
                _ = fi30
            }
            _ = v29
        } else {
            switch v27 {
            case 818:
                var v46 uint64 = v4
                _ = v46
            default:
                v47 := GenS4{F0: 42, F1: false, F2: "a"}
                _ = v47
            }
        }
    case 128:
        fmt.Println(v4)
    case 2147483647:
        var v48 int16 = 439
        _ = v48
    }
    fmt.Println(v5, v1, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
