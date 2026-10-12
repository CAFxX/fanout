package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint; F1 int32; F2 uint8 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 bool; F2 uint }

func main() {
    var v1 bool = true
    v2 := make(map[uint32]uint32)
    var v3 int = 650
    v4 := []int{}
    v5 := make(map[int]bool)
    fmt.Println(v1, v3)
    v1 = (false && v1)
    v6 := GenS1{F0: -9223372036854775808, F1: -32768}
    switch (v3 + (v3 + 172)) {
    case 32767:
        var v7 int8 = 42
        _ = v7
    default:
        fmt.Println(v3)
    }
    fmt.Println(v3, v3)
    if (305 < 42) {
        fmt.Println(v1, v1)
    } else {
        switch (v3 ^ v3) {
        case 0:
            v8 := GenS1{F0: (^(v3)), F1: -324}
            _ = v8
        case 127:
            var v9 int = v3
            _ = v9
        case 32768:
            for fi10 := 0; fi10 < 18; fi10++ {
                fmt.Println(fi10, v1)
                switch "a" {
                case "a":
                    if v1 {
                        var v11 int32 = -2147483648
                        _ = v11
                    }
                case "\n\t":
                    if len(v4) > 0 {
                        v4[0] = (v3 * v3)
                    }
                case "hello":
                    if v1 {
                        if (100 < -111) {
                            if v1 {
                                delete(v2, 65535)
                                var v12 int8 = -128
                                _ = v12
                            }
                            v3 -= fi10
                            for fi13 := 0; fi13 < 3; fi13++ {
                                v14 := GenS1{F0: fi13, F1: 130}
                                fmt.Println(v1, fi13)
                                v3 += (-(fi13))
                                _ = fi13
                                _ = v14
                            }
                        } else {
                            for fi15 := 0; fi15 < 12; fi15++ {
                                v2[1] = 2147483647
                                var v16 int64 = -744
                                fmt.Println(fi15, fi10, v3)
                                _ = fi15
                                _ = v16
                            }
                            v17 := GenS1{F0: (^(fi10)), F1: -838}
                            if (915 == 0) {
                                switch "\n\t" {
                                case "\n\t":
                                    for fi18 := 0; fi18 < 7; fi18++ {
                                        var v19 int8 = -1
                                        if v1 {
                                            fmt.Println(v19)
                                        } else {
                                            _ = len(v5)
                                            switch fi10 {
                                            case 379:
                                                if (!v1) {
                                                    if (true && (65535 >= 1)) {
                                                        fmt.Println(fi10)
                                                    } else {
                                                        if v1 {
                                                            fmt.Println(v1, v1)
                                                            var v20 float32 = -2.25
                                                            _ = v20
                                                        }
                                                        v3 = (fi18 << ((563) & 63))
                                                        var v21 uint = 1
                                                        _ = v21
                                                    }
                                                    v22 := GenS2{F0: 255, F1: 522, F2: 128}
                                                    _ = v22
                                                }
                                            case -886:
                                                fmt.Println(v19, fi10, fi10)
                                                fmt.Println(v3, fi10, v19)
                                            default:
                                                if len(v4) > 0 {
                                                    v4[0] = 1
                                                }
                                            }
                                            switch v3 {
                                            case -2147483648:
                                                v19 -= (-(v19))
                                                var v23 int16 = -1
                                                _ = v23
                                            case 745:
                                                v1 = v1
                                            default:
                                                v24 := GenS2{F0: 4294967295, F1: 1, F2: 127}
                                                _ = v24
                                            }
                                            v25 := GenS1{F0: (fi10 | fi10), F1: -32768}
                                            _ = v25
                                        }
                                        _ = fi18
                                        _ = v19
                                    }
                                case "foo bar":
                                    for fi26 := 0; fi26 < 14; fi26++ {
                                        v2[957] = 923
                                        _ = fi26
                                    }
                                default:
                                    var v27 int64 = 42
                                    _ = v27
                                }
                                switch (fi10 >> ((1) & 63)) {
                                case 42:
                                    _ = len(v5)
                                case 4294967295:
                                    _ = len(v4)
                                case 0:
                                    v1 = (264 == 0)
                                default:
                                    var v28 uint16 = 0
                                    _ = v28
                                }
                            }
                            _ = v17
                        }
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = len(v4)
                        v29 := GenS3{F0: 0, F1: fi10}
                        _ = v29
                    }
                }
                _ = fi10
            }
        default:
            var v30 bool = (!false)
            _ = v30
        }
    }
    var v31 int16 = 0
    var v32 string = "foo bar"
    v33 := GenS4{F0: (v31 | (v31 * 304))}
    v34 := GenS2{F0: 2, F1: -594, F2: 182}
    v31 -= (v31 ^ v31)
    delete(v2, 0)
    v1 = (!true)
    var v35 int64 = -1
    switch ("hello" + "a") {
    case "hello":
        var v36 int8 = 100
        _ = v36
    case "a":
        v37 := GenS2{F0: 2, F1: 1, F2: 0}
        _ = v37
    case "":
        v38 := GenS5{F0: 65535, F1: false, F2: 962}
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
