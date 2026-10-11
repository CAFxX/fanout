package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint; F1 int; F2 int }
type GenS3 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint64; F2 int16 }

func main() {
    var v1 bool = false
    v2 := []string{"hello", "a", "foo bar"}
    v3 := make(map[int]uint32)
    var v4 int = 42
    v5 := make(map[int]uint32)
    var v6 int8 = 100
    for fi7 := 0; fi7 < 5; fi7++ {
        if ("hello" <= "a") {
            v8 := GenS1{F0: 32767, F1: -32768}
            var v9 int = (v4 >> ((65535) & 63))
            _ = v8
            _ = v9
        }
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 uint16 = 99
            var v12 uint32 = 439
            if (v6 > v6) {
                v12 = 4294967295
                v12 = (v12 | (v12 / ((v12) | 1)))
                v3[4294967295] = (v12 << ((1) & 63))
            } else {
                v13 := GenS2{F0: 65535, F1: (v4 + fi10), F2: fi7}
                var v14 string = "hello"
                if ((^(v11)) <= (v11 << ((619) & 63))) {
                    switch 4294967295 {
                    case 128:
                        if len(v2) > 0 {
                            v2[0] = v14
                        }
                    case 127:
                        if (false && true) {
                            var v15 int = v4
                            var v16 int64 = 0
                            var v17 uint16 = (v11 % ((42) | 1))
                            switch (v4 % ((4294967295) | 1)) {
                            case 1:
                                var v18 int32 = 1
                                _ = v18
                            }
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                    case 42:
                        var v19 int16 = 88
                        _ = v19
                    }
                    if v1 {
                        _ = len(v2)
                        var v20 int64 = -9223372036854775808
                        v2 = append(v2, v14)
                        var v21 int16 = 0
                        _ = v20
                        _ = v21
                    }
                } else {
                    fmt.Println(fi10)
                }
                switch (^(v4)) {
                case 128:
                    var v22 int32 = int32(v11)
                    _ = v22
                case 32768:
                    var v23 string = v14
                    _ = v23
                case -2147483648:
                    v24 := GenS3{F0: 65535, F1: uint8(v11), F2: (v12 % (((v12 >> ((2) & 63))) | 1))}
                    _ = v24
                default:
                    var v25 int8 = v6
                    _ = v25
                }
                fmt.Println(v14, v12)
                _ = v13
                _ = v14
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        v26 := GenS1{F0: 612, F1: 1}
        _ = fi7
        _ = v26
    }
    v4 += (v4 * v4)
    if (0 <= 1) {
        fmt.Println(v4)
        v1 = ((18446744073709551615 == 65535) || (!true))
    }
    switch -290 {
    case 529:
        if (18446744073709551615 <= 860) {
            switch "a" {
            case "hello":
                v27 := GenS1{F0: 483, F1: 1}
                _ = v27
            case "a":
                v4 += 0
            case "":
                if true {
                    var v28 uint64 = 18446744073709551615
                    _ = len(v2)
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 2; fi29++ {
                        var v30 uint64 = 4294967295
                        v30 = v30
                        _ = fi29
                        _ = v30
                    }
                    for fi31 := 0; fi31 < 13; fi31++ {
                        if v1 {
                            v32 := GenS1{F0: 967, F1: 42}
                            _ = v32
                        }
                        if ((!(!(255 <= 788))) && v1) {
                            fmt.Println(fi31)
                            switch "a" {
                            case "hello":
                                v1 = (true || (!false))
                            case "foo bar":
                                switch (-(fi31)) {
                                case 936:
                                    var v33 int32 = int32(v6)
                                    _ = v33
                                }
                            case "a":
                                v1 = v1
                            }
                        }
                        _ = fi31
                    }
                }
            }
            switch "hello" {
            case "a":
                v34 := GenS3{F0: 65535, F1: 8, F2: 4294967295}
                _ = v34
            case "foo bar":
                v35 := GenS2{F0: 2, F1: (v4 + v4), F2: v4}
                _ = v35
            case "":
                var v36 uint16 = 1
                _ = v36
            }
        }
    case -1:
        v1 = true
    case 1:
        v37 := v5[(v4 % ((v4) | 1))]
        _ = v37
    default:
        v1 = v1
    }
    var v38 uint8 = 0
    v39 := GenS3{F0: 1, F1: (v38 >> ((uint(v38)) & 63)), F2: 1}
    if v1 {
        for fi40 := 0; fi40 < 13; fi40++ {
            if (-1 != 42) {
                v41 := GenS4{F0: -150, F1: 0, F2: 32767}
                switch 0 {
                case -60:
                    v38 = 128
                case 911:
                    var v42 int16 = -799
                    _ = v42
                case 938:
                    for fi43 := 0; fi43 < 16; fi43++ {
                        if (42 > 42) {
                            v1 = (0 < 32767)
                            v1 = (805 <= 503)
                            var v44 uint32 = 4294967295
                            _ = v44
                        } else {
                            if (true && v1) {
                                fmt.Println(fi43, v38, v1)
                            } else {
                                var v45 uint16 = 32767
                                v45 += v45
                                _ = v45
                            }
                        }
                        _ = len(v5)
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                        _ = fi43
                    }
                default:
                    v6 = (v6 >> ((18446744073709551615) & 63))
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v41
            }
            _ = fi40
        }
        for fi46 := 0; fi46 < 18; fi46++ {
            var v47 uint = 18446744073709551615
            v47 = 614
            switch fi46 {
            case 725:
                for fi48 := 0; fi48 < 17; fi48++ {
                    switch -1 {
                    case -2147483648:
                        fmt.Println(v47, v1)
                    case 9223372036854775807:
                        var v49 int16 = -32768
                        _ = v49
                    case 1:
                        v47 = v47
                    }
                    _ = fi48
                }
            case -9223372036854775808:
                fmt.Println(v1, v47)
            default:
                if ((1.5 * (1e10 + 1e10)) > float64(3.14159)) {
                    for fi50 := 0; fi50 < 16; fi50++ {
                        var v51 uint8 = v38
                        v4 = (fi50 % (((v4 | -223)) | 1))
                        v4 += (v4 / (((v4 - fi50)) | 1))
                        _ = fi50
                        _ = v51
                    }
                    fmt.Println(v4, v38, v6)
                }
            }
            _ = fi46
            _ = v47
        }
        v52 := GenS1{F0: 255, F1: -749}
        _ = v52
    } else {
        fmt.Println(v4, v6, v6)
        var v53 uint = 255
        _ = v53
    }
    switch (("hello" + (("hello" + "a") + "\n\t")) + "hello") {
    case "hello":
        fmt.Println(v38, v38)
    default:
        var v54 int16 = 235
        _ = v54
    }
    var v55 uint16 = 0
    if (!(v1 && false)) {
        fmt.Println(v38, v55)
        v3[(v4 << ((uint(v55)) & 63))] = 42
        var v56 uint8 = 67
        _ = v56
    } else {
        switch ("" + "foo bar") {
        case "":
            v38 -= v38
        case "\n\t":
            v4 = (^(v4))
        default:
            fmt.Println(v4)
        }
        var v57 int32 = 582
        _ = v57
    }
    if (-1 != 42) {
        var v58 int8 = (v6 + v6)
        _ = len(v2)
        v55 = (^(v55))
        _ = v58
    } else {
        for fi59 := 0; fi59 < 15; fi59++ {
            var v60 int = (fi59 ^ (fi59 / (((^(fi59))) | 1)))
            _ = fi59
            _ = v60
        }
        fmt.Println(v38, v55, v55)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v55
    _ = v6
}
