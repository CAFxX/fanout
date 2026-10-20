package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 string; F2 int64 }
type GenS3 struct { F0 uint8; F1 int32 }

func main() {
    var v1 float32 = 1e10
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 128
    var v4 uint16 = 255
    v5 := make(map[uint32]int)
    v6 := make(map[int]int)
    switch "a" {
    case "foo bar":
        switch "foo bar" {
        case "a":
            v3 = (v3 - (v3 | v3))
        case "hello":
            for fi7 := 0; fi7 < 13; fi7++ {
                if (696 != 42) {
                    var v8 string = (("\n\t" + "foo bar") + "\n\t")
                    v9 := []int{}
                    if (-219 >= -1) {
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        fmt.Println(v4, v1)
                        switch v8 {
                        case "":
                            var v10 uint16 = 255
                            _ = v10
                        case "foo bar":
                            switch v8 {
                            case "hello":
                                var v11 uint8 = (v3 & (v3 | 2))
                                _ = v11
                            }
                        case "\n\t":
                            for fi12 := 0; fi12 < 18; fi12++ {
                                var v13 uint = uint(v4)
                                _ = fi12
                                _ = v13
                            }
                        }
                    } else {
                        if (!false) {
                            v1 = v1
                            fmt.Println(fi7, v3, v8)
                        } else {
                            v1 = (-2.25 / v1)
                        }
                        if ((false || false) || (!true)) {
                            v3 -= 0
                            var v14 int32 = 42
                            _ = v14
                        }
                        fmt.Println(v1, v8)
                    }
                    _ = v8
                    _ = v9
                }
                for fi15 := 0; fi15 < 2; fi15++ {
                    fmt.Println(v1, v1)
                    v4 -= v4
                    _ = fi15
                }
                _ = fi7
            }
        }
    case "":
        switch 913 {
        case -401:
            if (-968 > -9223372036854775808) {
                v3 += 42
                switch ("hello" + "") {
                case "foo bar":
                    var v16 string = ""
                    _ = v16
                case "":
                    v4 -= (v4 + 255)
                default:
                    for fi17 := 0; fi17 < 11; fi17++ {
                        fmt.Println(v3)
                        _ = fi17
                    }
                }
                v18 := GenS1{F0: 4294967295}
                _ = v18
            } else {
                v19 := []int{-9223372036854775808, 525, -1, 4294967295}
                _ = v19
            }
        default:
            v4 += (v4 % (((v4 + 65535)) | 1))
        }
    case "\n\t":
        v1 += -2.25
    default:
        fmt.Println(v3, v1, v3)
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        fmt.Println(v3)
        var v21 int = (fi20 ^ fi20)
        v4 = 42
        _ = fi20
        _ = v21
    }
    if false {
        for fi22 := 0; fi22 < 7; fi22++ {
            v23 := []int{(fi22 - fi22), fi22}
            _ = fi22
            _ = v23
        }
        v24 := []int{}
        _ = v24
    } else {
        v4 = 0
        var v25 uint = 483
        _ = v25
    }
    fmt.Println(v1)
    delete(v6, 9223372036854775807)
    var v26 int = 447
    var v27 uint8 = v3
    switch "a" {
    case "foo bar":
        var v28 int64 = -1
        _ = v28
    case "hello":
        _ = len(v2)
    case "\n\t":
        for fi29 := 0; fi29 < 17; fi29++ {
            v3 -= (v3 / ((127) | 1))
            v30 := GenS2{F0: 2, F1: "\n\t", F2: -9223372036854775808}
            _ = fi29
            _ = v30
        }
    default:
        v31 := GenS3{F0: 137, F1: int32(v4)}
        _ = v31
    }
    switch 42 {
    case -9223372036854775808:
        v4 += v4
    case -156:
        if true {
            var v32 bool = (false || (3.14159 >= v1))
            var v33 int64 = -571
            _ = v32
            _ = v33
        }
    default:
        fmt.Println(v27, v27)
    }
    var v34 bool = (!true)
    for fi35 := 0; fi35 < 16; fi35++ {
        v36 := GenS1{F0: 824}
        v34 = (42 >= 1)
        v37 := GenS2{F0: 18446744073709551615, F1: "hello", F2: 9223372036854775807}
        _ = fi35
        _ = v36
        _ = v37
    }
    var v38 string = "a"
    var v39 int16 = -105
    if (!false) {
        for fi40 := 0; fi40 < 12; fi40++ {
            for fi41 := 0; fi41 < 11; fi41++ {
                switch (v38 + v38) {
                case "a":
                    switch -549 {
                    case 32767:
                        v38 = "\n\t"
                    }
                }
                _ = len(v6)
                switch (fi41 << ((0) & 63)) {
                case 1:
                    v34 = (v34 || (-86 < int8(v27)))
                default:
                    var v42 uint64 = 4294967295
                    _ = v42
                }
                _ = fi41
            }
            v34 = (109 >= 2)
            var v43 int8 = -106
            _ = fi40
            _ = v43
        }
        v1 -= (v1 / (((v1 - (1.5 - (v1 / (v1 - 0.1)))) / 1.0) / ((v1 + v1) - v1)))
        var v44 int8 = 127
        _ = v44
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
