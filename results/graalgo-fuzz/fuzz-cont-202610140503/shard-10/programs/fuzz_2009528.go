package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 96
    v3 := []int{}
    var v4 float64 = 1e10
    for fi5 := 0; fi5 < 17; fi5++ {
        _ = len(v1)
        switch fi5 {
        case 2147483647:
            fmt.Println(v4)
        case 32767:
            v6 := GenS1{F0: 1, F1: (^(v2))}
            _ = v6
        case 9223372036854775807:
            v2 = 219
        }
        var v7 bool = false
        _ = fi5
        _ = v7
    }
    v4 += (v4 * (v4 * 1e10))
    var v8 uint64 = 683
    v9 := GenS1{F0: v8, F1: v2}
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 int16 = 0
        _ = fi10
        _ = v11
    }
    var v12 uint64 = uint64(v2)
    if true {
        v4 = v4
    } else {
        var v13 int16 = -208
        _ = v13
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        for fi15 := 0; fi15 < 6; fi15++ {
            v16 := GenS2{F0: 0, F1: (fi14 % ((127) | 1))}
            _ = fi15
            _ = v16
        }
        switch ("foo bar" + ("hello" + "hello")) {
        case "a":
            for fi17 := 0; fi17 < 11; fi17++ {
                if false {
                    _ = len(v3)
                } else {
                    if (1 == (v12 | v12)) {
                        fmt.Println(v12, fi17, v4)
                        switch ((("\n\t" + "\n\t") + "") + "") {
                        case "foo bar":
                            v4 = v4
                        case "\n\t":
                            v8 = v8
                        }
                    }
                    var v18 int64 = -9223372036854775808
                    var v19 int32 = 261
                    _ = v18
                    _ = v19
                }
                v4 -= 0.0
                _ = fi17
            }
        case "":
            for fi20 := 0; fi20 < 19; fi20++ {
                if (128 < 1) {
                    if (0 <= 769) {
                        _ = len(v3)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                }
                v1[4294967295] = (((-1 >= -1) || false) && (!true))
                switch ("a" + ("hello" + "foo bar")) {
                case "":
                    var v21 int16 = 557
                    _ = v21
                }
                var v22 uint8 = v2
                _ = fi20
                _ = v22
            }
        case "foo bar":
            switch (fi14 + -619) {
            case 9223372036854775807:
                for fi23 := 0; fi23 < 3; fi23++ {
                    var v24 int16 = 1
                    _ = fi23
                    _ = v24
                }
            case -1:
                for fi25 := 0; fi25 < 8; fi25++ {
                    switch int(v12) {
                    case 11:
                        var v26 bool = false
                        _ = v26
                    case 32768:
                        var v27 uint16 = 576
                        _ = v27
                    }
                    var v28 uint16 = 0
                    var v29 uint = 540
                    _ = fi25
                    _ = v28
                    _ = v29
                }
            }
        default:
            v4 += v4
        }
        for fi30 := 0; fi30 < 14; fi30++ {
            fmt.Println(fi14, v8)
            v12 = v12
            v8 += 18446744073709551615
            _ = fi30
        }
        _ = fi14
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        v4 -= (((v4 / ((1.5 + float64((float32(v4) * 0.0))) * 1.5)) / (float64(((-2.25 * ((3.14159 * 3.14159) + 3.14159)) / 1e10)) * 1e10)) * float64(0.0))
        v32 := GenS3{F0: ("foo bar" + "\n\t")}
        _ = fi31
        _ = v32
    }
    v4 = (v4 * v4)
    var v33 bool = true
    delete(v1, 618)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v8
    _ = v9
}
