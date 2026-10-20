package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int }
type GenS2 struct { F0 uint32; F1 int; F2 int64 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 uint }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    var v4 int = 32768
    v5 := []uint32{42}
    v4 = v4
    for fi6 := 0; fi6 < 7; fi6++ {
        if (!true) {
            v7 := GenS1{F0: 42, F1: -1, F2: v4}
            _ = v7
        }
        _ = fi6
    }
    _ = len(v5)
    switch (v4 % ((v4) | 1)) {
    case 32767:
        if (true && (255 < 2)) {
            var v8 float64 = 1.5
            fmt.Println(v8, v4)
            _ = v8
        } else {
            for fi9 := 0; fi9 < 19; fi9++ {
                for fi10 := 0; fi10 < 10; fi10++ {
                    if (-9223372036854775808 >= 9223372036854775807) {
                        var v11 int8 = -128
                        if true {
                            v12 := GenS2{F0: 2147483648, F1: fi10, F2: -316}
                            var v13 uint32 = 2
                            v11 = (-(v11))
                            _ = v12
                            _ = v13
                        }
                        _ = v11
                    }
                    var v14 uint64 = 260
                    _ = fi10
                    _ = v14
                }
                _ = fi9
            }
        }
    default:
        if true {
            v4 = (v4 ^ 127)
            v4 += v4
            var v15 int16 = 739
            _ = v15
        } else {
            if true {
                if (0 != 0) {
                    var v16 int64 = 315
                    _ = v16
                }
                fmt.Println(v4)
            } else {
                fmt.Println(v4)
            }
            var v17 uint16 = 32767
            _ = len(v1)
            _ = v17
        }
    }
    v4 = v4
    if false {
        v4 = 4294967295
    } else {
        switch v4 {
        case 32768:
            switch 32768 {
            case -972:
                if (((-2.25 - (-2.25 + 0.0)) >= (float32(((float64((float32(1.5) - ((1e10 * 1.5) / 3.14159))) - 0.1) * (-2.25 + 3.14159))) + 0.0)) || false) {
                    for fi18 := 0; fi18 < 4; fi18++ {
                        var v19 int16 = 728
                        for fi20 := 0; fi20 < 18; fi20++ {
                            switch "hello" {
                            case "":
                                var v21 uint32 = 704
                                _ = v21
                            case "hello":
                                v5 = append(v5, 0)
                            case "\n\t":
                                fmt.Println(v4)
                            default:
                                var v22 float64 = 3.14159
                                _ = v22
                            }
                            v23 := GenS3{F0: 0, F1: 0}
                            _ = fi20
                            _ = v23
                        }
                        _ = fi18
                        _ = v19
                    }
                }
            case 128:
                for fi24 := 0; fi24 < 17; fi24++ {
                    _ = len(v2)
                    _ = fi24
                }
            }
        case 127:
            _ = len(v2)
        case -911:
            fmt.Println(v4)
        default:
            fmt.Println(v4)
        }
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        switch "\n\t" {
        case "\n\t":
            for fi26 := 0; fi26 < 17; fi26++ {
                switch "foo bar" {
                case "a":
                    var v27 uint32 = 401
                    _ = v27
                case "":
                    var v28 uint = 2
                    _ = v28
                }
                if false {
                    var v29 uint8 = 1
                    v4 += v4
                    delete(v3, "\n\t")
                    _ = v29
                }
                _ = fi26
            }
        case "hello":
            for fi30 := 0; fi30 < 15; fi30++ {
                v31 := GenS1{F0: 42, F1: 42, F2: -9223372036854775808}
                _ = fi30
                _ = v31
            }
        case "":
            for fi32 := 0; fi32 < 11; fi32++ {
                var v33 uint32 = 1
                _ = fi32
                _ = v33
            }
        }
        _ = fi25
    }
    v5 = append(v5, 4294967295)
    var v34 uint16 = 255
    v34 += 195
    v34 += v34
    switch ("\n\t" + "a") {
    case "a":
        _ = len(v2)
    case "\n\t":
        v35 := GenS1{F0: 0, F1: 1000, F2: (v4 ^ 267)}
        _ = v35
    case "foo bar":
        v34 += v34
    default:
        v4 = v4
    }
    if true {
        var v36 uint32 = 729
        if (-881 >= 0) {
            for fi37 := 0; fi37 < 5; fi37++ {
                var v38 string = "foo bar"
                _ = fi37
                _ = v38
            }
            v39 := GenS3{F0: 13, F1: 23}
            if (389 > 0) {
                for fi40 := 0; fi40 < 8; fi40++ {
                    v36 += (v36 | (^(v36)))
                    var v41 float32 = ((3.14159 / ((1.5 * 1e10) / (0.1 - (-2.25 + 0.0)))) / 3.14159)
                    _ = len(v5)
                    _ = fi40
                    _ = v41
                }
            } else {
                var v42 int16 = -774
                fmt.Println(v42, v36, v42)
                var v43 int32 = -1
                _ = v42
                _ = v43
            }
            _ = v39
        } else {
            if (42 < 2) {
                if true {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                v44 := GenS4{F0: 967}
                _ = v44
            }
            var v45 int8 = 42
            _ = v45
        }
        _ = v36
    } else {
        var v46 int64 = 42
        var v47 int32 = -826
        v48 := GenS3{F0: 1, F1: 42}
        _ = v46
        _ = v47
        _ = v48
    }
    fmt.Println(v4, v4)
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
