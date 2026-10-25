package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int16; F1 int16; F2 int }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 int8 }

func main() {
    v1 := []int{}
    var v2 float64 = -2.25
    var v3 uint16 = 2
    var v4 float64 = (v2 + v2)
    fmt.Println(v4, v3, v2)
    if true {
        v3 = v3
        if false {
            v5 := GenS1{F0: (2147483648 > 65535), F1: 1, F2: 42}
            fmt.Println(v2)
            _ = v5
        } else {
            v4 -= (0.1 + v4)
            v2 -= (v4 + float64(0.1))
            switch "hello" {
            case "":
                if false {
                    var v6 int8 = -76
                    _ = v6
                }
            case "foo bar":
                var v7 float32 = (float32(1e10) * ((((3.14159 - 1e10) + 0.0) + 0.1) + 1e10))
                _ = v7
            case "hello":
                v3 -= v3
            }
        }
        for fi8 := 0; fi8 < 13; fi8++ {
            switch 32767 {
            case 127:
                for fi9 := 0; fi9 < 7; fi9++ {
                    var v10 uint32 = 2
                    if len(v1) > 0 {
                        v1[0] = fi8
                    }
                    _ = fi9
                    _ = v10
                }
            default:
                if (42 < 554) {
                    for fi11 := 0; fi11 < 8; fi11++ {
                        if (-2.25 == (3.14159 + 0.0)) {
                            v12 := GenS1{F0: true, F1: 9223372036854775807, F2: 113}
                            v4 += ((v2 - float64(float32(1.5))) * v4)
                            v1 = append(v1, (fi11 & 32767))
                            _ = v12
                        }
                        _ = fi11
                    }
                    var v13 int8 = 75
                    _ = v13
                }
            }
            var v14 uint8 = 255
            _ = fi8
            _ = v14
        }
    } else {
        v15 := GenS2{F0: 18446744073709551615}
        v2 = v2
        v16 := GenS3{F0: 42, F1: 42}
        _ = v15
        _ = v16
    }
    switch 32767 {
    case -9223372036854775808:
        for fi17 := 0; fi17 < 9; fi17++ {
            for fi18 := 0; fi18 < 14; fi18++ {
                if (2147483647 != 0) {
                    v3 = (v3 / ((v3) | 1))
                    if ((v3 != (v3 * (v3 >> ((255) & 63)))) && true) {
                        v3 = v3
                    } else {
                        if len(v1) > 0 {
                            v1[0] = (-(fi18))
                        }
                        fmt.Println(v4, v3, fi17)
                    }
                } else {
                    fmt.Println(v2, v3, v2)
                }
                if (true || true) {
                    var v19 uint32 = 433
                    v2 += 0.1
                    var v20 bool = false
                    _ = v19
                    _ = v20
                }
                _ = fi18
            }
            _ = len(v1)
            _ = fi17
        }
    default:
        switch 9223372036854775807 {
        case -678:
            v21 := GenS3{F0: 100, F1: 32767}
            _ = v21
        case 1:
            var v22 int32 = 0
            _ = v22
        }
    }
    if true {
        if (v4 != (3.14159 + v4)) {
            v23 := GenS3{F0: -101, F1: 32767}
            _ = len(v1)
            _ = v23
        } else {
            var v24 uint64 = 18446744073709551615
            for fi25 := 0; fi25 < 5; fi25++ {
                v24 = 18446744073709551615
                _ = fi25
            }
            _ = v24
        }
        v3 -= (v3 + v3)
        v26 := GenS2{F0: 2}
        _ = v26
    }
    var v27 float64 = float64((3.14159 * 1.5))
    for fi28 := 0; fi28 < 7; fi28++ {
        var v29 uint8 = 255
        fmt.Println(v27, v27, v27)
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        v31 := GenS1{F0: false, F1: 42, F2: 92}
        switch (fi30 ^ fi30) {
        case 1:
            var v32 int32 = 42
            _ = v32
        case 42:
            for fi33 := 0; fi33 < 14; fi33++ {
                v3 += v3
                var v34 int64 = -721
                var v35 int32 = -131
                v36 := GenS3{F0: 109, F1: 1}
                _ = fi33
                _ = v34
                _ = v35
                _ = v36
            }
        case 9223372036854775807:
            v3 -= (v3 / ((65535) | 1))
        default:
            fmt.Println(v4, v2)
        }
        _ = fi30
        _ = v31
    }
    var v37 float32 = (0.1 * (3.14159 * 3.14159))
    var v38 float32 = v37
    if (1000 <= 42) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = len(v1)
        for fi39 := 0; fi39 < 20; fi39++ {
            v37 -= v37
            _ = fi39
        }
    } else {
        if (!true) {
            v40 := GenS4{F0: -32768, F1: -32768, F2: 32768}
            fmt.Println(v3)
            if (4294967295 != 870) {
                v41 := GenS5{F0: 65535}
                switch 1 {
                case -439:
                    for fi42 := 0; fi42 < 6; fi42++ {
                        var v43 uint = 2
                        var v44 uint8 = 2
                        switch ("a" + ((("a" + "hello") + "a") + "")) {
                        case "hello":
                            for fi45 := 0; fi45 < 19; fi45++ {
                                v4 += v2
                                fmt.Println(fi45, v43)
                                _ = fi45
                            }
                        case "foo bar":
                            if (261 != 1) {
                                v37 += 3.14159
                                v46 := GenS6{F0: v43, F1: -100}
                                v43 = v43
                                _ = v46
                            }
                        }
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                case 127:
                    var v47 string = "a"
                    _ = v47
                case 1:
                    var v48 uint16 = 65535
                    _ = v48
                }
                _ = v41
            }
            _ = v40
        } else {
            var v49 uint8 = 42
            switch "hello" {
            case "hello":
                var v50 float32 = v38
                _ = v50
            case "foo bar":
                var v51 uint8 = v49
                _ = v51
            default:
                switch ("foo bar" + "") {
                case "hello":
                    if (!true) {
                        var v52 int16 = 0
                        v4 += v2
                        for fi53 := 0; fi53 < 2; fi53++ {
                            var v54 string = ("foo bar" + "\n\t")
                            if len(v1) > 0 {
                                v1[0] = 127
                            }
                            _ = fi53
                            _ = v54
                        }
                        _ = v52
                    } else {
                        fmt.Println(v2)
                        if false {
                            v55 := GenS7{F0: 17}
                            _ = v55
                        }
                    }
                case "foo bar":
                    v56 := GenS1{F0: true, F1: -9223372036854775808, F2: 0}
                    _ = v56
                case "a":
                    v1 = append(v1, -2147483648)
                }
            }
            switch "foo bar" {
            case "foo bar":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = v49
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v57 float64 = float64(0.1)
        _ = v57
    }
    var v58 int64 = 169
    fmt.Println(v3, v38, v27)
    v38 -= 0.1
    var v59 int16 = -32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v58
    _ = v59
}
