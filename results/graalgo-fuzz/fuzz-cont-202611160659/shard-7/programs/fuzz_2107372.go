package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 uint64 }

func main() {
    var v1 int = 285
    v2 := make(map[uint32]string)
    var v3 uint8 = 2
    var v4 string = "a"
    if (-1 > 1) {
        var v5 int32 = -1
        var v6 float64 = 3.14159
        v4 = ("hello" + v4)
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 2147483647}
    switch (("hello" + v4) + ("foo bar" + "foo bar")) {
    case "":
        var v8 uint32 = 1
        _ = v8
    }
    var v9 uint64 = 682
    var v10 bool = (42 != 102)
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := GenS2{F0: 0, F1: v4}
        v13 := GenS2{F0: (v3 % (((v3 ^ (^(v3)))) | 1)), F1: ("" + v4)}
        var v14 uint64 = v9
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    v10 = (893 > 2147483647)
    v9 += (v9 % ((318) | 1))
    v15 := []int{}
    v9 += v9
    if false {
        for fi16 := 0; fi16 < 10; fi16++ {
            fmt.Println(v3)
            _ = fi16
        }
        if v10 {
            var v17 int = (-(v1))
            _ = v17
        } else {
            if v10 {
                for fi18 := 0; fi18 < 11; fi18++ {
                    if (!v10) {
                        if len(v15) > 0 {
                            v15[0] = 32767
                        }
                        var v19 bool = (!(true && v10))
                        fmt.Println(v3)
                        _ = v19
                    } else {
                        v4 = v4
                    }
                    fmt.Println(v3)
                    if v10 {
                        var v20 float32 = 0.1
                        _ = v20
                    } else {
                        switch ("a" + v4) {
                        case "hello":
                            v2[2] = v4
                        }
                        fmt.Println(fi18, v9)
                    }
                    _ = fi18
                }
                fmt.Println(v4, v1, v9)
            }
        }
    } else {
        switch (v4 + "a") {
        case "\n\t":
            if (v10 || (((1e10 / (1.5 - -2.25)) > 3.14159) && (!((!(!false)) || (v4 < v4))))) {
                var v21 int32 = 140
                _ = v21
            }
        case "hello":
            var v22 int64 = 9223372036854775807
            _ = v22
        case "":
            switch "hello" {
            case "foo bar":
                v10 = (v10 && v10)
            case "hello":
                v10 = ((v3 + 128) <= v3)
            }
        default:
            if v10 {
                for fi23 := 0; fi23 < 15; fi23++ {
                    for fi24 := 0; fi24 < 20; fi24++ {
                        var v25 uint64 = uint64(v3)
                        fmt.Println(v25, v3)
                        v26 := GenS3{F0: -638, F1: 255}
                        var v27 uint = 255
                        _ = fi24
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    fmt.Println(v10)
                    _ = fi23
                }
                v15 = append(v15, 1)
            } else {
                var v28 uint16 = 65535
                _ = v28
            }
        }
        v29 := GenS4{F0: 1, F1: 666}
        var v30 int32 = 305
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 5; fi31++ {
        if (v3 != 168) {
            if ((v3 | (v3 ^ v3)) < v3) {
                switch "" {
                case "hello":
                    switch (v4 + (((v4 + v4) + v4) + v4)) {
                    case "hello":
                        var v32 uint8 = (^(v3))
                        _ = v32
                    case "a":
                        v33 := GenS3{F0: -213, F1: uint16(v3)}
                        _ = v33
                    case "foo bar":
                        v34 := GenS5{F0: (v9 << ((255) & 63))}
                        _ = v34
                    default:
                        var v35 uint8 = (^(v3))
                        _ = v35
                    }
                case "foo bar":
                    var v36 int32 = 110
                    _ = v36
                case "":
                    var v37 int16 = -589
                    _ = v37
                }
            } else {
                for fi38 := 0; fi38 < 6; fi38++ {
                    var v39 uint16 = 101
                    var v40 uint = 4294967295
                    fmt.Println(fi31, v1)
                    _ = fi38
                    _ = v39
                    _ = v40
                }
            }
            var v41 uint64 = (v9 - 2)
            _ = v41
        }
        for fi42 := 0; fi42 < 17; fi42++ {
            var v43 float64 = (1e10 / 1.0)
            _ = fi42
            _ = v43
        }
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
