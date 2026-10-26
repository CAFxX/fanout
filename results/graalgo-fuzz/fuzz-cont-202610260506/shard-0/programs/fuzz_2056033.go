package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 bool; F1 int16; F2 int64 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 uint64; F1 bool }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 0
    var v3 int16 = 379
    v4 := []int{}
    v3 -= 698
    v5 := GenS1{F0: 1, F1: 114}
    if (857 < 181) {
        fmt.Println(v2)
        v2 = 1
    }
    var v6 uint32 = 293
    for fi7 := 0; fi7 < 18; fi7++ {
        if (65535 >= 42) {
            if false {
                v8 := GenS2{F0: false, F1: (v3 * v3), F2: 42}
                v6 = (^(v6))
                v1[410] = (v6 << ((2) & 63))
                _ = v8
            } else {
                delete(v1, v6)
                var v9 int8 = 100
                _ = v9
            }
            if true {
                v10 := GenS1{F0: (v3 / ((0) | 1)), F1: 128}
                for fi11 := 0; fi11 < 14; fi11++ {
                    switch fi7 {
                    case 379:
                        v2 = v2
                    case 128:
                        switch ("\n\t" + "\n\t") {
                        case "a":
                            var v12 uint8 = 255
                            _ = v12
                        }
                    default:
                        var v13 uint64 = 2
                        v13 -= (^(v13))
                        _ = v13
                    }
                    v6 = v6
                    _ = fi11
                }
                fmt.Println(v2, v2, v6)
                _ = v10
            }
            switch ("hello" + ("a" + ("foo bar" + ("a" + (("foo bar" + "hello") + "\n\t"))))) {
            case "":
                v14 := GenS3{F0: v3, F1: (int16(v2) < 42)}
                _ = v14
            case "\n\t":
                v2 -= v2
            }
        }
        var v15 int32 = int32(v6)
        _ = fi7
        _ = v15
    }
    var v16 string = (("foo bar" + "\n\t") + "hello")
    switch 9223372036854775807 {
    case 127:
        fmt.Println(v2)
    default:
        v1[(v6 % ((0) | 1))] = 160
    }
    switch 1 {
    case 127:
        switch v16 {
        case "foo bar":
            fmt.Println(v16)
        default:
            v6 = v6
        }
    case -919:
        if true {
            if false {
                switch v16 {
                case "\n\t":
                    if (false && false) {
                        var v17 uint32 = (v6 % ((v6) | 1))
                        _ = v17
                    } else {
                        fmt.Println(v3)
                        var v18 uint16 = v2
                        _ = v18
                    }
                case "a":
                    v3 += v3
                case "hello":
                    var v19 int64 = -9223372036854775808
                    _ = v19
                }
                if true {
                    v20 := GenS1{F0: v3, F1: 212}
                    _ = v20
                } else {
                    var v21 int8 = 100
                    v22 := GenS2{F0: true, F1: (v3 / (((v3 | 480)) | 1)), F2: -839}
                    _ = v21
                    _ = v22
                }
            } else {
                switch v16 {
                case "a":
                    for fi23 := 0; fi23 < 20; fi23++ {
                        v24 := GenS2{F0: (!false), F1: 1000, F2: int64(v3)}
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi23
                        _ = v24
                    }
                }
                v3 += (v3 / (((v3 ^ v3)) | 1))
            }
        } else {
            v25 := GenS3{F0: v3, F1: (false && (!false))}
            _ = v25
        }
    }
    v3 -= (^(v3))
    for fi26 := 0; fi26 < 12; fi26++ {
        if (true || false) {
            v27 := GenS4{F0: 0, F1: false}
            _ = v27
        } else {
            delete(v1, (v6 - (^(v6))))
        }
        for fi28 := 0; fi28 < 20; fi28++ {
            v4 = append(v4, fi28)
            var v29 int16 = v3
            v16 = v16
            _ = fi28
            _ = v29
        }
        _ = fi26
    }
    var v30 uint8 = 40
    v6 = (v6 * (v6 ^ (^(v6))))
    v31 := GenS1{F0: (-(v3)), F1: (v30 % ((v30) | 1))}
    if (1e10 <= 0.0) {
        var v32 int = 42
        if false {
            for fi33 := 0; fi33 < 8; fi33++ {
                var v34 int64 = 0
                var v35 float32 = 1.5
                _ = fi33
                _ = v34
                _ = v35
            }
        } else {
            v30 = (v30 % ((255) | 1))
        }
        var v36 uint = 883
        _ = v32
        _ = v36
    } else {
        var v37 int16 = (v3 | v3)
        fmt.Println(v2, v30, v2)
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
