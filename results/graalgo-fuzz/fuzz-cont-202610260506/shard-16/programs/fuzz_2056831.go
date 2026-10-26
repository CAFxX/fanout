package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint16 }
type GenS2 struct { F0 int64; F1 int16; F2 uint16 }
func genFn0(p0 int) int16 {
    var v5 uint64 = 271
    if true {
        v5 += 1
    } else {
        var v6 int64 = -1
        _ = v6
    }
    _ = p0
    _ = v5
    return 42
}

func main() {
    var v1 int32 = 42
    var v2 int = 4294967295
    var v3 int16 = 297
    v4 := []uint32{1}
    var v7 float64 = 0.0
    if (127 != 127) {
        v3 += v3
    } else {
        if false {
            var v8 int32 = (v1 >> ((2) & 63))
            v4 = append(v4, 2)
            _ = v8
        } else {
            var v9 uint32 = 0
            _ = v9
        }
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v2)
        }
    }
    _ = len(v4)
    var v10 uint16 = 1
    v11 := GenS1{F0: "", F1: 207, F2: 32767}
    var v12 float32 = 1.5
    if true {
        v7 = (((1.5 + v7) - -2.25) / v7)
        v10 = v10
        var v13 bool = (false && true)
        _ = v13
    } else {
        for fi14 := 0; fi14 < 9; fi14++ {
            for fi15 := 0; fi15 < 19; fi15++ {
                var v16 int = (v2 >> ((1) & 63))
                _ = fi15
                _ = v16
            }
            if true {
                v2 = v2
                _ = len(v4)
                v10 += (^(v10))
            }
            _ = fi14
        }
        var v17 uint32 = 2
        for fi18 := 0; fi18 < 9; fi18++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi18
        }
        _ = v17
    }
    v2 += v2
    if (false && true) {
        switch ("foo bar" + ("a" + "")) {
        case "foo bar":
            var v19 string = "a"
            _ = v19
        case "":
            fmt.Println(v1)
        case "\n\t":
            if true {
                var v20 int8 = 95
                var v21 int16 = (v3 >> ((4294967295) & 63))
                fmt.Println(v10, v21)
                _ = v20
                _ = v21
            }
        default:
            fmt.Println(v10, v2, v3)
        }
        for fi22 := 0; fi22 < 14; fi22++ {
            _ = len(v4)
            if (((v7 / v7) >= (float64((0.1 - v12)) * (v7 - ((-2.25 * (v7 * v7)) * v7)))) && (false || true)) {
                fmt.Println(v1)
            }
            _ = fi22
        }
        for fi23 := 0; fi23 < 16; fi23++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi23
        }
    } else {
        if len(v4) > 0 {
            v4[0] = 2
        }
        for fi24 := 0; fi24 < 12; fi24++ {
            fmt.Println(v3, fi24, v10)
            var v25 int8 = 42
            _ = fi24
            _ = v25
        }
        if true {
            var v26 float32 = (v12 + v12)
            _ = v26
        }
    }
    v27 := GenS1{F0: "\n\t", F1: -9223372036854775808, F2: 647}
    switch "hello" {
    case "hello":
        v1 = (v1 ^ (v1 * v1))
    case "a":
        if (-737 == 1) {
            if (v1 == v1) {
                for fi28 := 0; fi28 < 12; fi28++ {
                    v29 := GenS2{F0: -9223372036854775808, F1: genFn0(fi28), F2: (v10 * v10)}
                    _ = fi28
                    _ = v29
                }
                for fi30 := 0; fi30 < 8; fi30++ {
                    var v31 int64 = 1
                    var v32 int8 = 0
                    for fi33 := 0; fi33 < 10; fi33++ {
                        var v34 int8 = v32
                        _ = fi33
                        _ = v34
                    }
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            }
            switch (v2 % (((v2 - (v2 << ((255) & 63)))) | 1)) {
            case 9223372036854775807:
                var v35 uint16 = (v10 & 896)
                _ = v35
            default:
                v12 -= 1e10
            }
            v4 = append(v4, 42)
        } else {
            v1 = v1
            v12 -= float32(-2.25)
            _ = len(v4)
        }
    case "\n\t":
        var v36 int64 = 0
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v7
}
