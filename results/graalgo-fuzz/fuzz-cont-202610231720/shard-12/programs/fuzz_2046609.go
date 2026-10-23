package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 int; F2 string }

func main() {
    var v1 uint16 = 0
    var v2 int = 78
    v3 := make(map[uint32]uint32)
    var v4 int8 = 42
    v5 := make(map[uint32]bool)
    if true {
        if true {
            if false {
                for fi6 := 0; fi6 < 17; fi6++ {
                    v7 := []int{}
                    var v8 int16 = 0
                    _ = fi6
                    _ = v7
                    _ = v8
                }
                v9 := GenS1{F0: 304}
                for fi10 := 0; fi10 < 20; fi10++ {
                    var v11 uint8 = 235
                    v4 += v4
                    _ = fi10
                    _ = v11
                }
                _ = v9
            } else {
                v12 := GenS1{F0: 42}
                v1 = (v1 * 935)
                _ = v12
            }
        } else {
            for fi13 := 0; fi13 < 4; fi13++ {
                switch (^(fi13)) {
                case -2147483648:
                    var v14 int64 = 0
                    _ = v14
                case 128:
                    var v15 int32 = 1
                    _ = v15
                case 572:
                    var v16 float64 = (3.14159 * 0.0)
                    _ = v16
                }
                v1 = (v1 & v1)
                _ = fi13
            }
        }
        if (1 == 667) {
            var v17 float64 = -2.25
            _ = len(v3)
            _ = v17
        } else {
            var v18 int64 = -9223372036854775808
            switch (v2 + v2) {
            case 42:
                var v19 uint32 = 403
                _ = v19
            case -1:
                v4 = 1
            case -378:
                v20 := []int{}
                _ = v20
            }
            _ = v18
        }
    } else {
        switch "foo bar" {
        case "foo bar":
            v4 += (v4 | v4)
        }
    }
    _ = len(v5)
    v21 := GenS2{F0: 1, F1: 0}
    var v22 string = "\n\t"
    var v23 uint = 255
    for fi24 := 0; fi24 < 19; fi24++ {
        if false {
            if (uint16(v23) != v1) {
                fmt.Println(v22)
                if (805 >= 1000) {
                    v4 = (v4 << ((v23) & 63))
                } else {
                    for fi25 := 0; fi25 < 1; fi25++ {
                        v26 := GenS3{F0: 42}
                        _ = fi25
                        _ = v26
                    }
                    v22 = "hello"
                }
                v2 = (fi24 / (((fi24 - (fi24 | -176))) | 1))
            }
            v27 := GenS1{F0: v23}
            fmt.Println(v4)
            _ = v27
        }
        _ = len(v3)
        _ = fi24
    }
    for fi28 := 0; fi28 < 8; fi28++ {
        if (1000 <= -32768) {
            var v29 int8 = (v4 / (((-(v4))) | 1))
            v30 := GenS4{F0: 65535, F1: 514, F2: ""}
            _ = v29
            _ = v30
        }
        for fi31 := 0; fi31 < 7; fi31++ {
            v4 -= (v4 + (v4 ^ v4))
            _ = fi31
        }
        _ = len(v3)
        _ = fi28
    }
    var v32 bool = true
    v1 -= v1
    for fi33 := 0; fi33 < 1; fi33++ {
        if (0 > (-(v4))) {
            v32 = v32
            fmt.Println(v4)
            var v34 uint64 = 762
            _ = v34
        } else {
            if (v32 && (1 == 1)) {
                if (((0.0 * ((1e10 / 3.14159) * 0.1)) * float64(-2.25)) > float64((1e10 - 1e10))) {
                    if ((v1 >> (((v23 * 4294967295)) & 63)) != 781) {
                        v23 = (^(v23))
                    }
                }
                if (1e10 != (1e10 / 0.1)) {
                    var v35 int32 = 2147483647
                    _ = v35
                } else {
                    var v36 float32 = (float32(0.1) + 0.0)
                    v37 := GenS1{F0: (v23 + (v23 | 74))}
                    fmt.Println(v4, v22, v32)
                    _ = v36
                    _ = v37
                }
            } else {
                switch ("\n\t" + v22) {
                case "a":
                    var v38 float32 = (0.1 + 1.5)
                    _ = v38
                case "hello":
                    fmt.Println(fi33)
                case "\n\t":
                    var v39 float32 = 1.5
                    _ = v39
                default:
                    for fi40 := 0; fi40 < 11; fi40++ {
                        v4 += v4
                        _ = len(v3)
                        switch "" {
                        case "foo bar":
                            fmt.Println(v22, v4, v4)
                            v41 := GenS2{F0: -572, F1: 4294967295}
                            _ = v41
                        case "\n\t":
                            var v42 uint = v23
                            _ = v42
                        default:
                            if v32 {
                                if false {
                                    v23 -= v23
                                }
                                fmt.Println(v22, fi33, v1)
                            }
                        }
                        _ = fi40
                    }
                }
                fmt.Println(v1)
                switch (v2 | v2) {
                case 1:
                    switch "" {
                    case "a":
                        var v43 float32 = (0.1 * ((0.0 / 3.14159) - (0.0 + (0.1 + (float32((1e10 - ((((1.5 * 1.5) - ((0.1 * 0.1) / 0.1)) / 1e10) / 1e10))) + ((0.1 + 1.5) * (1e10 / (0.0 - 0.1))))))))
                        _ = v43
                    default:
                        delete(v5, 2)
                    }
                default:
                    var v44 string = ((v22 + v22) + "\n\t")
                    _ = v44
                }
            }
        }
        v45 := GenS2{F0: 2147483647, F1: 2}
        var v46 float32 = float32(((0.0 / 3.14159) + float64(-2.25)))
        _ = fi33
        _ = v45
        _ = v46
    }
    if (v32 || v32) {
        v47 := GenS4{F0: (v1 >> ((539) & 63)), F1: (-(v2)), F2: ("hello" + ((v22 + "\n\t") + v22))}
        var v48 float32 = float32(((0.1 + 1.5) / 0.1))
        _ = v47
        _ = v48
    }
    v32 = v32
    v49 := GenS1{F0: v23}
    v2 -= (v2 - -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v32
    _ = v4
    _ = v49
    _ = v5
}
