package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint64; F1 int32; F2 bool }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 string; F1 int32; F2 bool }
type GenS5 struct { F0 uint16; F1 uint }

func main() {
    v1 := []int{355, 42, -9223372036854775808}
    v2 := []int{376, 489}
    v3 := []string{}
    var v4 int8 = -1
    v4 += -67
    switch 1 {
    case 2147483647:
        if (42 >= 65535) {
            v5 := GenS1{F0: 843, F1: 18446744073709551615}
            var v6 bool = true
            _ = v5
            _ = v6
        } else {
            fmt.Println(v4)
        }
    case -228:
        for fi7 := 0; fi7 < 19; fi7++ {
            v4 += (v4 + v4)
            fmt.Println(v4, v4)
            _ = fi7
        }
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 uint8 = 2
        _ = fi8
        _ = v9
    }
    switch 752 {
    case 32767:
        if (1e10 < -2.25) {
            fmt.Println(v4)
            switch "" {
            case "hello":
                fmt.Println(v4)
            case "":
                v1 = append(v1, 2147483647)
            case "a":
                for fi10 := 0; fi10 < 16; fi10++ {
                    if false {
                        v11 := GenS2{F0: 245, F1: 42, F2: (!("a" == "hello"))}
                        for fi12 := 0; fi12 < 3; fi12++ {
                            v4 = (^(v4))
                            var v13 uint = 4294967295
                            v4 -= -70
                            _ = fi12
                            _ = v13
                        }
                        switch "hello" {
                        case "foo bar":
                            v4 -= v4
                        case "":
                            for fi14 := 0; fi14 < 10; fi14++ {
                                v4 = (-(v4))
                                var v15 int16 = 0
                                var v16 int64 = 1
                                _ = fi14
                                _ = v15
                                _ = v16
                            }
                        default:
                            fmt.Println(fi10, v4)
                        }
                        _ = v11
                    } else {
                        fmt.Println(fi10, fi10)
                        v4 -= (-(v4))
                        v4 += v4
                    }
                    _ = fi10
                }
            default:
                if false {
                    switch 373 {
                    case -739:
                        switch "foo bar" {
                        case "foo bar":
                            var v17 float64 = -2.25
                            _ = v17
                        case "a":
                            v3 = append(v3, "\n\t")
                        case "\n\t":
                            var v18 string = ("a" + "a")
                            _ = v18
                        default:
                            var v19 uint32 = 311
                            _ = v19
                        }
                    default:
                        var v20 int32 = -2147483648
                        _ = v20
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    fmt.Println(v4)
                    v4 -= (v4 >> ((42) & 63))
                }
            }
        } else {
            for fi21 := 0; fi21 < 13; fi21++ {
                fmt.Println(fi21, v4)
                _ = fi21
            }
        }
    case 2147483647:
        v4 += (-(v4))
    case 42:
        switch "\n\t" {
        case "":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "\n\t":
            if ((v4 >= -128) || (!false)) {
                v4 = (v4 ^ 127)
                for fi22 := 0; fi22 < 2; fi22++ {
                    var v23 float64 = 0.1
                    _ = len(v2)
                    _ = fi22
                    _ = v23
                }
            } else {
                for fi24 := 0; fi24 < 18; fi24++ {
                    v4 = 58
                    v25 := GenS1{F0: fi24, F1: 923}
                    fmt.Println(v4)
                    _ = fi24
                    _ = v25
                }
            }
        case "a":
            fmt.Println(v4)
        }
    }
    if false {
        for fi26 := 0; fi26 < 19; fi26++ {
            for fi27 := 0; fi27 < 8; fi27++ {
                var v28 int16 = 42
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
        switch 4294967295 {
        case -144:
            v4 = v4
        case 42:
            if ((!false) || ((v4 * 100) < (v4 - v4))) {
                var v29 bool = false
                _ = v29
            }
        }
        _ = len(v3)
    } else {
        var v30 uint64 = 42
        _ = v30
    }
    v31 := GenS3{F0: 2, F1: 441}
    v32 := GenS3{F0: 360, F1: 18446744073709551615}
    v4 = v4
    if (v4 > (v4 & 1)) {
        var v33 int64 = -356
        var v34 int8 = v4
        _ = v33
        _ = v34
    }
    if (true || false) {
        v35 := GenS4{F0: "hello", F1: 42, F2: false}
        v36 := GenS2{F0: 42, F1: 2147483647, F2: false}
        switch 1 {
        case 42:
            var v37 int8 = 20
            _ = v37
        case 127:
            v4 += (v4 % ((v4) | 1))
        }
        _ = v35
        _ = v36
    } else {
        switch ("a" + "") {
        case "hello":
            switch 4294967295 {
            case -341:
                var v38 float32 = 0.1
                _ = v38
            case 1:
                for fi39 := 0; fi39 < 16; fi39++ {
                    var v40 int16 = -32768
                    fmt.Println(v4, v4)
                    v4 += (v4 - v4)
                    _ = fi39
                    _ = v40
                }
            default:
                for fi41 := 0; fi41 < 18; fi41++ {
                    v1 = append(v1, (fi41 * (^(fi41))))
                    v4 = 127
                    _ = fi41
                }
            }
        default:
            v4 += v4
        }
        v42 := GenS5{F0: 2, F1: 65535}
        v43 := GenS3{F0: 2, F1: 323}
        _ = v42
        _ = v43
    }
    v44 := GenS2{F0: 439, F1: 2147483647, F2: false}
    var v45 uint32 = 42
    fmt.Println(v4, v4)
    switch "a" {
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v45 = v45
    v4 = (v4 | -105)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v44
    _ = v45
}
