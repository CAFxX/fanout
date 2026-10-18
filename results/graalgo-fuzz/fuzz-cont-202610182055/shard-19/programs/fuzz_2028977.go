package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS5 struct { F0 int; F1 int64; F2 uint32 }

func main() {
    v1 := []uint32{0}
    var v2 int32 = 258
    v3 := make(map[string]bool)
    var v4 bool = true
    v5 := make(map[uint32]uint32)
    for fi6 := 0; fi6 < 15; fi6++ {
        fmt.Println(fi6, v4)
        v3["foo bar"] = (428 == -1)
        if true {
            delete(v5, 2147483648)
            switch -2147483648 {
            case 172:
                if v4 {
                    v7 := GenS1{F0: 65535}
                    var v8 int32 = v2
                    _ = v7
                    _ = v8
                } else {
                    v9 := GenS1{F0: 2147483647}
                    if (v4 && v4) {
                        var v10 uint32 = 65535
                        var v11 uint8 = 42
                        for fi12 := 0; fi12 < 6; fi12++ {
                            var v13 uint32 = (v10 / ((v10) | 1))
                            v13 = v13
                            _ = fi12
                            _ = v13
                        }
                        _ = v10
                        _ = v11
                    } else {
                        fmt.Println(fi6, fi6)
                        v2 -= v2
                        v14 := GenS2{F0: 65535, F1: v2}
                        _ = v14
                    }
                    var v15 bool = v4
                    _ = v15
                    _ = v9
                }
            case -355:
                var v16 int8 = 100
                _ = v16
            case -629:
                fmt.Println(v4, v2)
            }
            if true {
                v4 = (!v4)
                var v17 float32 = (1.5 / 3.14159)
                _ = v17
            } else {
                v4 = (!false)
                fmt.Println(fi6)
            }
        } else {
            v18 := v5[2147483648]
            var v19 uint64 = 1
            for fi20 := 0; fi20 < 9; fi20++ {
                v21 := GenS3{F0: 127}
                v22 := GenS4{F0: 1, F1: v4, F2: (v19 ^ 42)}
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = v18
            _ = v19
        }
        _ = fi6
    }
    v4 = v4
    fmt.Println(v4, v4)
    for fi23 := 0; fi23 < 3; fi23++ {
        fmt.Println(fi23, v4, v4)
        v2 += v2
        _ = fi23
    }
    _ = len(v1)
    var v24 int64 = -311
    var v25 uint32 = 2
    v2 = (v2 << ((2) & 63))
    delete(v5, 65535)
    if v4 {
        if len(v1) > 0 {
            v1[0] = (v25 ^ v25)
        }
        _ = len(v1)
    } else {
        for fi26 := 0; fi26 < 16; fi26++ {
            var v27 uint8 = 128
            switch (fi26 / (((fi26 - fi26)) | 1)) {
            case 32767:
                if v4 {
                    for fi28 := 0; fi28 < 18; fi28++ {
                        fmt.Println(v27, v27, fi26)
                        _ = fi28
                    }
                }
            case -79:
                for fi29 := 0; fi29 < 16; fi29++ {
                    var v30 uint64 = 717
                    _ = fi29
                    _ = v30
                }
            case -738:
                var v31 uint32 = v25
                _ = v31
            default:
                switch fi26 {
                case 818:
                    v27 -= 128
                }
            }
            switch "foo bar" {
            case "":
                var v32 int32 = (v2 * 0)
                _ = v32
            case "hello":
                fmt.Println(v27, v2)
            case "foo bar":
                for fi33 := 0; fi33 < 3; fi33++ {
                    v5[(v25 >> ((65535) & 63))] = (v25 ^ (v25 / (((^(v25))) | 1)))
                    _ = fi33
                }
            default:
                v2 = v2
            }
            _ = fi26
            _ = v27
        }
    }
    v34 := GenS1{F0: v25}
    var v35 uint64 = 508
    if v4 {
        if false {
            fmt.Println(v2, v2, v25)
            switch "" {
            case "":
                if ("" != "\n\t") {
                    var v36 int = 0
                    var v37 int = 42
                    _ = v36
                    _ = v37
                }
            case "foo bar":
                switch "" {
                case "hello":
                    var v38 float32 = -2.25
                    _ = v38
                }
            }
            var v39 uint8 = 42
            _ = v39
        }
        for fi40 := 0; fi40 < 5; fi40++ {
            v24 = (v24 % ((v24) | 1))
            var v41 float64 = 3.14159
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi40
            _ = v41
        }
        for fi42 := 0; fi42 < 17; fi42++ {
            var v43 uint64 = v35
            _ = len(v5)
            _ = fi42
            _ = v43
        }
    } else {
        v44 := GenS1{F0: v25}
        v4 = v4
        var v45 uint16 = 765
        _ = v44
        _ = v45
    }
    switch 127 {
    case 32768:
        var v46 float32 = (1.5 - 0.0)
        _ = v46
    default:
        v47 := GenS2{F0: 65535, F1: v2}
        _ = v47
    }
    v48 := GenS5{F0: 32767, F1: (v24 * v24), F2: v25}
    _ = len(v3)
    fmt.Println(v35)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v48
    _ = v5
}
