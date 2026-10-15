package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 int; F1 bool; F2 int16 }
type GenS5 struct { F0 uint64 }

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 102
    var v3 int64 = -992
    v4 := []uint32{}
    v5 := GenS1{F0: (v3 | (v3 << ((129) & 63))), F1: 2, F2: (v2 - (v2 - 32767))}
    fmt.Println(v2, v2)
    for fi6 := 0; fi6 < 11; fi6++ {
        fmt.Println(v3, fi6, v2)
        for fi7 := 0; fi7 < 18; fi7++ {
            if (12 == fi6) {
                switch "a" {
                case "hello":
                    fmt.Println(v3)
                }
            }
            v8 := GenS1{F0: 9223372036854775807, F1: 42, F2: (v2 / (((v2 * v2)) | 1))}
            _ = fi7
            _ = v8
        }
        v3 += (v3 * (v3 & v3))
        _ = fi6
    }
    if (98 < 255) {
        v3 = -947
        fmt.Println(v3)
        v9 := GenS1{F0: (v3 >> ((969) & 63)), F1: 2, F2: v2}
        _ = v9
    } else {
        if (true && false) {
            delete(v1, "a")
            switch -1 {
            case 128:
                v10 := GenS2{F0: v2}
                _ = v10
            default:
                v11 := v1["a"]
                _ = v11
            }
            for fi12 := 0; fi12 < 4; fi12++ {
                var v13 uint64 = 868
                _ = fi12
                _ = v13
            }
        } else {
            var v14 int64 = (v3 << ((2) & 63))
            v3 += 0
            _ = v14
        }
    }
    var v15 int16 = 422
    fmt.Println(v2)
    if (1 > 127) {
        for fi16 := 0; fi16 < 14; fi16++ {
            v17 := GenS3{F0: 42, F1: 1}
            for fi18 := 0; fi18 < 15; fi18++ {
                fmt.Println(v2)
                v19 := GenS4{F0: 0, F1: false, F2: v15}
                _ = fi18
                _ = v19
            }
            _ = fi16
            _ = v17
        }
        for fi20 := 0; fi20 < 12; fi20++ {
            for fi21 := 0; fi21 < 1; fi21++ {
                v15 = 185
                v15 += (v15 / ((v15) | 1))
                _ = fi21
            }
            if false {
                v22 := GenS3{F0: 283, F1: 42}
                for fi23 := 0; fi23 < 10; fi23++ {
                    v15 = v15
                    fmt.Println(v3)
                    _ = fi23
                }
                _ = v22
            }
            _ = fi20
        }
    }
    switch 1 {
    case 130:
        for fi24 := 0; fi24 < 10; fi24++ {
            v15 += (v15 >> ((255) & 63))
            _ = len(v4)
            _ = fi24
        }
    case 127:
        if (false && ((float64(1e10) * (3.14159 - float64(0.0))) == (0.0 + (0.0 / -2.25)))) {
            var v25 uint8 = 0
            _ = v25
        }
    case 4294967295:
        for fi26 := 0; fi26 < 14; fi26++ {
            if false {
                for fi27 := 0; fi27 < 9; fi27++ {
                    var v28 float64 = float64(0.0)
                    _ = len(v1)
                    _ = fi27
                    _ = v28
                }
                v29 := GenS2{F0: v2}
                _ = v29
            }
            var v30 uint16 = v2
            switch ("hello" + "hello") {
            case "\n\t":
                v31 := GenS2{F0: v30}
                _ = v31
            case "foo bar":
                var v32 int16 = v15
                _ = v32
            case "":
                v33 := GenS5{F0: 352}
                _ = v33
            }
            _ = fi26
            _ = v30
        }
    }
    switch "hello" {
    case "foo bar":
        for fi34 := 0; fi34 < 3; fi34++ {
            _ = len(v1)
            for fi35 := 0; fi35 < 15; fi35++ {
                var v36 string = ""
                _ = len(v4)
                if (42 >= -128) {
                    v3 += v3
                }
                if (false || ((!false) && (65535 < 1))) {
                    v3 = (v3 + (-(v3)))
                    var v37 float64 = (-2.25 / 1.0)
                    v36 = v36
                    _ = v37
                }
                _ = fi35
                _ = v36
            }
            switch "" {
            case "foo bar":
                v3 += (^(v3))
            case "a":
                v2 = (v2 & 644)
            default:
                v38 := GenS2{F0: v2}
                _ = v38
            }
            _ = fi34
        }
    default:
        v15 += 32767
    }
    if (!true) {
        var v39 bool = false
        v15 -= 730
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
