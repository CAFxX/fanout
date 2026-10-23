package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS6 struct { F0 int32; F1 bool }

func main() {
    var v1 int32 = -633
    var v2 string = "hello"
    var v3 string = v2
    v4 := make(map[uint32]int)
    v2 = v3
    v5 := []int{}
    fmt.Println(v2, v1)
    var v6 float32 = (-2.25 - (1.5 * 1e10))
    var v7 bool = true
    v3 = ""
    for fi8 := 0; fi8 < 19; fi8++ {
        if v7 {
            v2 = "a"
            var v9 uint16 = 255
            if (int64(v9) >= 433) {
                var v10 string = "a"
                var v11 int = (fi8 + fi8)
                if (1 == 718) {
                    v10 = v10
                    v3 = "foo bar"
                    for fi12 := 0; fi12 < 1; fi12++ {
                        v10 = v10
                        var v13 uint8 = 42
                        _ = fi12
                        _ = v13
                    }
                }
                _ = v10
                _ = v11
            }
            _ = v9
        }
        v14 := GenS1{F0: 4294967295, F1: 0, F2: 1}
        _ = fi8
        _ = v14
    }
    v3 = "foo bar"
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 int32 = v1
        var v17 int8 = 127
        v18 := GenS1{F0: 4294967295, F1: 18446744073709551615, F2: 65535}
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    v5 = append(v5, -1)
    var v19 uint = 106
    if (255 >= 0) {
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 uint = v19
            _ = fi20
            _ = v21
        }
    } else {
        switch "foo bar" {
        case "hello":
            delete(v4, 4294967295)
        case "":
            var v22 int = 32768
            _ = v22
        case "foo bar":
            if (1 < 100) {
                v19 += v19
                fmt.Println(v6, v2)
            }
        default:
            var v23 bool = v7
            _ = v23
        }
        v24 := GenS1{F0: 121, F1: v19, F2: 42}
        _ = v24
    }
    var v25 uint32 = 0
    switch -106 {
    case 32768:
        for fi26 := 0; fi26 < 11; fi26++ {
            v27 := GenS1{F0: 666, F1: (v19 * 42), F2: 894}
            v28 := GenS2{F0: (v7 && (510 != 65535))}
            _ = fi26
            _ = v27
            _ = v28
        }
    case -2147483648:
        var v29 uint8 = 127
        _ = v29
    default:
        for fi30 := 0; fi30 < 3; fi30++ {
            switch fi30 {
            case -9223372036854775808:
                if v7 {
                    if ((v25 / ((v25) | 1)) > v25) {
                        switch v2 {
                        case "foo bar":
                            v4[(^(v25))] = 1
                        case "\n\t":
                            v6 += v6
                        case "":
                            if len(v5) > 0 {
                                v5[0] = fi30
                            }
                        default:
                            v31 := GenS3{F0: 403}
                            _ = v31
                        }
                    } else {
                        if (230 != 42) {
                            fmt.Println(v7, v2, v19)
                        }
                        switch "hello" {
                        case "":
                            _ = len(v5)
                        case "\n\t":
                            if v7 {
                                var v32 float32 = v6
                                _ = v32
                            }
                        }
                    }
                } else {
                    v33 := GenS4{F0: 84, F1: 42}
                    fmt.Println(v6, fi30, v25)
                    _ = v33
                }
            case 127:
                v34 := GenS1{F0: 42, F1: 18446744073709551615, F2: 18446744073709551615}
                _ = v34
            case -1:
                if v7 {
                    v35 := GenS5{F0: -128, F1: 2, F2: 42}
                    switch (fi30 % ((fi30) | 1)) {
                    case 32768:
                        var v36 uint16 = 57
                        _ = v36
                    case 32767:
                        v37 := GenS6{F0: v1, F1: true}
                        _ = v37
                    case 4294967295:
                        v38 := GenS3{F0: 701}
                        _ = v38
                    }
                    _ = v35
                }
            default:
                v25 = 2
            }
            fmt.Println(v1)
            _ = fi30
        }
    }
    for fi39 := 0; fi39 < 13; fi39++ {
        var v40 int64 = -1
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
