package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS5 struct { F0 uint; F1 int16; F2 int64 }
type GenS6 struct { F0 int; F1 uint16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 bool }

func main() {
    v1 := make(map[int]int)
    var v2 int64 = -58
    v3 := make(map[int]int)
    fmt.Println(v2)
    var v4 uint16 = 1
    fmt.Println(v2, v4)
    if true {
        v5 := GenS1{F0: 0, F1: -150}
        fmt.Println(v2)
        for fi6 := 0; fi6 < 12; fi6++ {
            var v7 int64 = -933
            for fi8 := 0; fi8 < 14; fi8++ {
                v9 := GenS1{F0: 255, F1: -1}
                if (962 == (^(v4))) {
                    var v10 string = "\n\t"
                    _ = v10
                }
                if false {
                    if ((((3.14159 + (float64(0.1) * float64((0.1 / 0.1)))) * float64(-2.25)) * 3.14159) < (3.14159 * 3.14159)) {
                        var v11 uint8 = 144
                        if false {
                            v12 := GenS1{F0: 118, F1: 280}
                            _ = v12
                        }
                        _ = v11
                    }
                    delete(v3, 2147483647)
                }
                var v13 int16 = -32768
                _ = fi8
                _ = v13
                _ = v9
            }
            v14 := GenS1{F0: 128, F1: 42}
            _ = fi6
            _ = v14
            _ = v7
        }
        _ = v5
    } else {
        delete(v1, 127)
        if true {
            v15 := GenS1{F0: 202, F1: 450}
            v4 = 1
            _ = v15
        } else {
            if false {
                switch "hello" {
                case "foo bar":
                    fmt.Println(v2, v4)
                case "\n\t":
                    v16 := []int{-1}
                    _ = v16
                default:
                    if true {
                        v2 = v2
                        switch ((("" + "a") + "") + ("\n\t" + "foo bar")) {
                        case "a":
                            if false {
                                v17 := GenS2{F0: (v4 | v4), F1: (2 <= 122)}
                                v2 = v2
                                var v18 bool = true
                                _ = v17
                                _ = v18
                            }
                        case "foo bar":
                            v4 += v4
                        default:
                            v19 := GenS3{F0: -825}
                            _ = v19
                        }
                    } else {
                        v20 := GenS2{F0: (v4 & v4), F1: false}
                        _ = v20
                    }
                }
                v4 = (^(v4))
            }
            for fi21 := 0; fi21 < 16; fi21++ {
                delete(v3, fi21)
                _ = fi21
            }
            if false {
                v4 += v4
            } else {
                v22 := GenS4{F0: 712, F1: 55, F2: v2}
                fmt.Println(v2)
                _ = v22
            }
        }
    }
    if (3.14159 >= 1e10) {
        for fi23 := 0; fi23 < 4; fi23++ {
            for fi24 := 0; fi24 < 16; fi24++ {
                if (100 != 100) {
                    fmt.Println(v4)
                    var v25 uint64 = 4294967295
                    v26 := GenS1{F0: 2, F1: 2147483647}
                    var v27 uint8 = 255
                    if ((^(v25)) <= (^(v25))) {
                        v28 := GenS3{F0: (v2 << ((547) & 63))}
                        var v29 bool = true
                        switch "a" {
                        case "\n\t":
                            v2 = (v2 >> ((94) & 63))
                        case "":
                            var v30 int16 = 822
                            _ = v30
                        }
                        _ = v28
                        _ = v29
                    }
                    _ = v25
                    _ = v26
                    _ = v27
                }
                _ = len(v1)
                _ = fi24
            }
            for fi31 := 0; fi31 < 19; fi31++ {
                var v32 bool = (((("" + "a") + "") + "hello") != "hello")
                v33 := GenS5{F0: 255, F1: 42, F2: (v2 - v2)}
                var v34 int32 = 42
                var v35 uint32 = 65535
                _ = fi31
                _ = v32
                _ = v33
                _ = v34
                _ = v35
            }
            _ = fi23
        }
        v36 := GenS6{F0: -432, F1: (v4 % ((580) | 1))}
        _ = v36
    }
    v4 += v4
    v2 = v2
    if (false || (!true)) {
        v37 := GenS7{F0: 287}
        var v38 uint = 4294967295
        _ = v37
        _ = v38
    }
    var v39 int8 = -10
    v39 = (v39 % (((v39 << ((4294967295) & 63))) | 1))
    v40 := GenS8{F0: false}
    var v41 int8 = v39
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
