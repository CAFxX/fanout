package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 int64 = 0
    v2 := make(map[uint32]bool)
    var v3 int64 = (v1 << ((4294967295) & 63))
    var v4 uint8 = 42
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]uint32)
    var v7 string = "\n\t"
    if ((false && ("" >= v7)) || false) {
        fmt.Println(v3, v1)
    }
    switch v7 {
    case "a":
        for fi8 := 0; fi8 < 15; fi8++ {
            for fi9 := 0; fi9 < 20; fi9++ {
                v1 += v3
                if (42 <= 42) {
                    v10 := GenS1{F0: -1, F1: 4294967295}
                    fmt.Println(v4, fi8)
                    var v11 uint = 2
                    _ = v10
                    _ = v11
                }
                v12 := GenS1{F0: -32768, F1: 42}
                fmt.Println(fi9, fi9, v7)
                if (false || ((0 > -602) && false)) {
                    v13 := []int{}
                    switch int(v1) {
                    case 32767:
                        switch v7 {
                        case "":
                            if false {
                                fmt.Println(fi8, fi8)
                            }
                        case "foo bar":
                            if (!false) {
                                switch fi8 {
                                case 4294967295:
                                    v1 = (v1 * v1)
                                case 9223372036854775807:
                                    var v14 uint16 = 42
                                    _ = v14
                                }
                                _ = len(v13)
                                if ((-1 < 32767) && false) {
                                    v15 := GenS2{F0: 42}
                                    _ = v15
                                }
                            }
                        default:
                            v1 -= (v1 % (((v1 >> ((255) & 63))) | 1))
                            v16 := GenS1{F0: 42, F1: 4294967295}
                            _ = v16
                        }
                    }
                    _ = v13
                } else {
                    var v17 uint = 0
                    var v18 string = ""
                    _ = v17
                    _ = v18
                }
                _ = fi9
                _ = v12
            }
            for fi19 := 0; fi19 < 20; fi19++ {
                v7 = "hello"
                v7 = v7
                var v20 uint16 = 1
                var v21 float64 = 1e10
                _ = fi19
                _ = v20
                _ = v21
            }
            var v22 uint8 = 106
            _ = fi8
            _ = v22
        }
    case "foo bar":
        if true {
            var v23 uint = 65535
            var v24 float32 = 3.14159
            _ = v23
            _ = v24
        }
    case "":
        var v25 uint8 = 42
        _ = v25
    }
    fmt.Println(v3, v4)
    for fi26 := 0; fi26 < 7; fi26++ {
        var v27 float64 = (-2.25 + 0.0)
        var v28 string = (v7 + v7)
        for fi29 := 0; fi29 < 14; fi29++ {
            v30 := GenS1{F0: -1, F1: 4294967295}
            v27 = v27
            if true {
                var v31 int64 = v3
                switch "a" {
                case "foo bar":
                    var v32 uint64 = 940
                    _ = v32
                }
                v33 := GenS3{F0: (v31 ^ (-(v1)))}
                _ = v31
                _ = v33
            } else {
                fmt.Println(v3, v3)
            }
            _ = fi29
            _ = v30
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    for fi34 := 0; fi34 < 9; fi34++ {
        v7 = v7
        _ = fi34
    }
    v4 = (v4 - v4)
    switch v7 {
    case "hello":
        if (true && (!false)) {
            var v35 int64 = 9223372036854775807
            _ = v35
        }
    default:
        v7 = "hello"
    }
    delete(v5, -2147483648)
    var v36 uint8 = v4
    v36 = (v4 / ((128) | 1))
    var v37 int = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
