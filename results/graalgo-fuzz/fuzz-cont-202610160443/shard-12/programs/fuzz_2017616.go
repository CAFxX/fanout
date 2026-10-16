package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64; F1 bool }

func main() {
    v1 := []string{("a" + "")}
    v2 := []uint32{4294967295}
    var v3 int8 = 127
    var v4 uint32 = 2147483647
    switch 9223372036854775807 {
    case 4294967295:
        for fi5 := 0; fi5 < 10; fi5++ {
            v2 = append(v2, v4)
            for fi6 := 0; fi6 < 14; fi6++ {
                var v7 uint = 255
                _ = fi6
                _ = v7
            }
            _ = fi5
        }
    case -203:
        var v8 int8 = v3
        _ = v8
    case 32767:
        if ((!true) && (!((true && (0 > 99)) || (3.14159 > (1.5 + (0.0 - 1.5)))))) {
            var v9 int16 = -402
            switch -149 {
            case 4294967295:
                v10 := GenS1{F0: v9, F1: 324, F2: 0}
                _ = v10
            case 0:
                fmt.Println(v4, v4)
            case -1:
                for fi11 := 0; fi11 < 1; fi11++ {
                    for fi12 := 0; fi12 < 1; fi12++ {
                        v13 := GenS1{F0: (v9 >> ((970) & 63)), F1: 859, F2: 255}
                        _ = fi12
                        _ = v13
                    }
                    switch "foo bar" {
                    case "a":
                        _ = len(v1)
                    case "\n\t":
                        var v14 int8 = v3
                        _ = v14
                    case "foo bar":
                        for fi15 := 0; fi15 < 4; fi15++ {
                            v3 += v3
                            var v16 int16 = 0
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi15
                            _ = v16
                        }
                    default:
                        if len(v2) > 0 {
                            v2[0] = v4
                        }
                    }
                    if (0 > 1) {
                        switch (fi11 & -457) {
                        case 4294967295:
                            v9 -= v9
                        case 32768:
                            var v17 uint64 = 182
                            _ = v17
                        case 127:
                            v9 += (-(v9))
                        }
                    } else {
                        var v18 uint8 = 43
                        _ = v18
                    }
                    _ = fi11
                }
            default:
                v19 := GenS1{F0: v9, F1: 2, F2: 65535}
                _ = v19
            }
            switch -9223372036854775808 {
            case -225:
                switch 9223372036854775807 {
                case 1:
                    if false {
                        var v20 int32 = 0
                        var v21 uint64 = 4294967295
                        v22 := GenS1{F0: 597, F1: 1, F2: 226}
                        _ = v20
                        _ = v21
                        _ = v22
                    } else {
                        for fi23 := 0; fi23 < 6; fi23++ {
                            if true {
                                v24 := GenS2{F0: fi23, F1: fi23}
                                v3 = v3
                                _ = v24
                            } else {
                                switch "foo bar" {
                                case "foo bar":
                                    switch fi23 {
                                    case -254:
                                        var v25 int16 = (-(v9))
                                        _ = v25
                                    case 1:
                                        v26 := GenS1{F0: (v9 + (v9 + v9)), F1: 2, F2: 42}
                                        _ = v26
                                    default:
                                        fmt.Println(v4)
                                    }
                                default:
                                    if (false || false) {
                                        var v27 float32 = -2.25
                                        var v28 int = (fi23 ^ (fi23 ^ (fi23 << ((370) & 63))))
                                        for fi29 := 0; fi29 < 16; fi29++ {
                                            if true {
                                                v30 := GenS1{F0: (v9 ^ (-(v9))), F1: 18446744073709551615, F2: 255}
                                                switch (^(fi23)) {
                                                case 96:
                                                    switch "" {
                                                    case "\n\t":
                                                        v31 := GenS3{F0: 65535, F1: 4294967295}
                                                        _ = v31
                                                    default:
                                                        v27 = (v27 * (v27 + ((float32(0.1) * v27) / (1e10 - v27))))
                                                    }
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = v30
                                            }
                                            var v32 int = fi23
                                            _ = fi29
                                            _ = v32
                                        }
                                        _ = v27
                                        _ = v28
                                    }
                                }
                            }
                            v9 = 1000
                            _ = fi23
                        }
                        v9 = v9
                        var v33 int16 = v9
                        _ = v33
                    }
                case 42:
                    var v34 float32 = 3.14159
                    _ = v34
                default:
                    var v35 bool = true
                    _ = v35
                }
            case 128:
                var v36 uint16 = 32767
                _ = v36
            default:
                v1 = append(v1, "\n\t")
            }
            _ = v9
        } else {
            v4 -= v4
            v37 := GenS3{F0: 1, F1: uint64(v4)}
            v3 = v3
            _ = v37
        }
    default:
        v38 := GenS2{F0: 32768, F1: 0}
        _ = v38
    }
    if true {
        v39 := GenS3{F0: 2, F1: 817}
        fmt.Println(v3, v3)
        switch "a" {
        case "a":
            v3 -= v3
        case "\n\t":
            v40 := GenS4{F0: v3}
            _ = v40
        }
        _ = v39
    }
    var v41 int32 = -1
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v4)
        v41 += (v41 * (v41 + v41))
    }
    v41 += 162
    v4 += (v4 << ((65535) & 63))
    var v42 uint16 = 20
    var v43 int32 = -1
    switch 4294967295 {
    case 67:
        var v44 uint64 = 18446744073709551615
        _ = v44
    case -351:
        var v45 int = 42
        _ = v45
    default:
        v46 := GenS5{F0: 0, F1: (724 == 512)}
        _ = v46
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
}
