package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 string }
func genFn0() string {
    v7 := []int{}
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := GenS1{F0: 315}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 uint64 = 0
        v12 := GenS1{F0: uint(v11)}
        _ = fi10
        _ = v11
        _ = v12
    }
    switch 32767 {
    case 491:
        if true {
            if len(v7) > 0 {
                v7[0] = 1
            }
            var v13 uint32 = 42
            _ = v13
        }
    }
    _ = v7
    return "hello"
}

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = 1
    v3 := make(map[string]uint32)
    v4 := []int{652}
    var v5 float64 = -2.25
    v6 := make(map[int]string)
    for fi14 := 0; fi14 < 12; fi14++ {
        v6[fi14] = "\n\t"
        v2 -= v2
        v5 = float64(0.1)
        _ = fi14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        v16 := GenS1{F0: 0}
        _ = fi15
        _ = v16
    }
    v5 -= v5
    var v17 uint = 42
    var v18 float64 = v5
    for fi19 := 0; fi19 < 10; fi19++ {
        if (true || true) {
            delete(v3, ("hello" + ""))
            var v20 int32 = -2147483648
            var v21 string = "hello"
            _ = v20
            _ = v21
        }
        for fi22 := 0; fi22 < 20; fi22++ {
            v23 := GenS2{F0: (v17 / ((65535) | 1)), F1: 0}
            v2 += 1
            v5 = 1.5
            _ = fi22
            _ = v23
        }
        var v24 int = (fi19 | 907)
        _ = fi19
        _ = v24
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        v18 += v18
        var v26 uint32 = 1
        v2 = v2
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        for fi28 := 0; fi28 < 2; fi28++ {
            fmt.Println(fi28)
            var v29 uint = 255
            var v30 int8 = 127
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = fi27
    }
    var v31 uint8 = 42
    switch "" {
    case "foo bar":
        var v32 int = 2147483647
        _ = v32
    case "\n\t":
        v2 += (^(v2))
    case "hello":
        for fi33 := 0; fi33 < 12; fi33++ {
            v2 += int32(v31)
            switch (-(fi33)) {
            case -1:
                switch fi33 {
                case 42:
                    v34 := GenS3{F0: 764, F1: 0}
                    _ = v34
                case -9223372036854775808:
                    v5 += float64(-2.25)
                default:
                    switch "a" {
                    case "":
                        v35 := GenS2{F0: v17, F1: -32768}
                        _ = v35
                    case "hello":
                        _ = len(v4)
                    }
                }
            }
            var v36 int16 = 32767
            _ = fi33
            _ = v36
        }
    default:
        var v37 uint32 = 2147483647
        _ = v37
    }
    for fi38 := 0; fi38 < 15; fi38++ {
        for fi39 := 0; fi39 < 11; fi39++ {
            var v40 float64 = v18
            _ = fi39
            _ = v40
        }
        v5 = float64(0.1)
        _ = fi38
    }
    switch -1 {
    case 4294967295:
        v41 := GenS3{F0: 0, F1: 2147483647}
        _ = v41
    case 9223372036854775807:
        if false {
            v42 := GenS4{F0: ""}
            switch 128 {
            case 9223372036854775807:
                _ = len(v4)
            case -9223372036854775808:
                for fi43 := 0; fi43 < 4; fi43++ {
                    v44 := GenS4{F0: ""}
                    _ = fi43
                    _ = v44
                }
            case 127:
                var v45 float32 = ((1e10 * 1e10) + 1e10)
                _ = v45
            default:
                fmt.Println(v5, v18, v18)
            }
            _ = v42
        } else {
            v4 = append(v4, -535)
            if (42 == -128) {
                var v46 int = 42
                _ = v46
            } else {
                _ = len(v4)
            }
            if (!(!(false || false))) {
                v47 := GenS3{F0: 92, F1: 2}
                v5 -= 0.1
                var v48 bool = true
                _ = v47
                _ = v48
            } else {
                var v49 int16 = 0
                fmt.Println(v17)
                if false {
                    if true {
                        for fi50 := 0; fi50 < 3; fi50++ {
                            for fi51 := 0; fi51 < 3; fi51++ {
                                switch "a" {
                                case "":
                                    var v52 bool = ((((("hello" + "hello") + "") + "foo bar") >= "foo bar") || ((false || (v49 == (v49 % ((v49) | 1)))) || false))
                                    _ = v52
                                case "foo bar":
                                    v53 := GenS3{F0: 42, F1: 65535}
                                    _ = v53
                                case "hello":
                                    v2 += (v2 - v2)
                                default:
                                    v54 := GenS2{F0: (v17 + v17), F1: 1}
                                    _ = v54
                                }
                                switch fi51 {
                                case 42:
                                    var v55 uint16 = 0
                                    _ = v55
                                default:
                                    if (!false) {
                                        var v56 uint8 = 0
                                        var v57 int16 = (v49 ^ (v49 | (v49 % ((v49) | 1))))
                                        v57 = 704
                                        _ = v56
                                        _ = v57
                                    } else {
                                        fmt.Println(v2, fi50, v2)
                                        var v58 uint64 = 0
                                        v59 := GenS3{F0: 42, F1: 443}
                                        _ = v58
                                        _ = v59
                                    }
                                }
                                _ = fi51
                            }
                            v49 += (v49 >> (((v17 * (v17 >> ((v17) & 63)))) & 63))
                            _ = fi50
                        }
                        var v60 uint = (v17 & (^(v17)))
                        _ = v60
                    }
                    var v61 uint16 = 1
                    v2 += v2
                    _ = v61
                }
                _ = v49
            }
        }
    case 32768:
        var v62 uint8 = (^(v31))
        _ = v62
    default:
        var v63 int = 128
        _ = v63
    }
    for fi64 := 0; fi64 < 7; fi64++ {
        for fi65 := 0; fi65 < 20; fi65++ {
            if true {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v66 string = "\n\t"
                _ = v66
            } else {
                v2 = v2
            }
            v17 += 243
            _ = fi65
        }
        var v67 uint8 = v31
        _ = fi64
        _ = v67
    }
    if (127 <= 32768) {
        fmt.Println(v2, v31)
        v18 = v18
        if (100 < 3) {
            var v68 bool = false
            _ = v68
        }
    } else {
        v5 = v18
        var v69 int8 = 1
        for fi70 := 0; fi70 < 15; fi70++ {
            for fi71 := 0; fi71 < 3; fi71++ {
                var v72 int32 = 289
                _ = fi71
                _ = v72
            }
            v17 = (^(v17))
            _ = fi70
        }
        _ = v69
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
