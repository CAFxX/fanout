package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint }
type GenS2 struct { F0 bool; F1 uint }

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = ((((0.0 - 1.5) / 1.5) / 1e10) - 0.1)
    v3 := make(map[string]uint32)
    v4 := []uint32{548}
    switch "" {
    case "a":
        for fi5 := 0; fi5 < 3; fi5++ {
            v2 = v2
            v6 := v1[1]
            _ = fi5
            _ = v6
        }
    case "":
        if true {
            v2 -= v2
            fmt.Println(v2)
        } else {
            fmt.Println(v2)
        }
    default:
        v7 := v1[2147483648]
        _ = v7
    }
    var v8 int8 = -105
    if false {
        for fi9 := 0; fi9 < 19; fi9++ {
            switch ("foo bar" + "\n\t") {
            case "\n\t":
                fmt.Println(v2, fi9)
            case "":
                v10 := GenS1{F0: 1, F1: 1, F2: 255}
                _ = v10
            default:
                for fi11 := 0; fi11 < 20; fi11++ {
                    v2 += 1e10
                    switch fi9 {
                    case 4294967295:
                        var v12 uint64 = 0
                        _ = v12
                    }
                    v2 -= float32(float64(v2))
                    v2 = (0.0 / v2)
                    _ = fi11
                }
            }
            _ = fi9
        }
        delete(v3, "")
    } else {
        v13 := GenS1{F0: 751, F1: 42, F2: 33}
        var v14 string = "hello"
        var v15 float32 = 3.14159
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 int = -1
    if (!(0 == 0)) {
        fmt.Println(v16, v2)
        v17 := GenS1{F0: 65535, F1: 65535, F2: 665}
        v18 := GenS1{F0: 1, F1: 0, F2: 694}
        _ = v17
        _ = v18
    }
    switch (v16 * v16) {
    case 1:
        var v19 uint64 = 2
        _ = v19
    case 662:
        if (-1 <= -970) {
            for fi20 := 0; fi20 < 9; fi20++ {
                var v21 uint16 = 995
                if (!(65535 >= 604)) {
                    if ((false && true) || (770 != 0)) {
                        var v22 int16 = 42
                        for fi23 := 0; fi23 < 15; fi23++ {
                            v24 := GenS1{F0: 687, F1: uint(v21), F2: 754}
                            v25 := GenS1{F0: 65535, F1: 1, F2: 42}
                            var v26 uint8 = 42
                            _ = fi23
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 10; fi27++ {
                            v22 -= -353
                            var v28 int64 = 42
                            var v29 int32 = -1
                            fmt.Println(v2)
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        _ = v22
                    } else {
                        for fi30 := 0; fi30 < 19; fi30++ {
                            v2 = ((1e10 * v2) - (v2 / (v2 + v2)))
                            _ = fi30
                        }
                        v21 += 0
                    }
                }
                v21 -= (v21 * v21)
                _ = fi20
                _ = v21
            }
        } else {
            var v31 int64 = -1
            v31 = v31
            if ((-(v8)) > v8) {
                var v32 int = -9223372036854775808
                for fi33 := 0; fi33 < 20; fi33++ {
                    fmt.Println(v32)
                    _ = fi33
                }
                _ = v32
            } else {
                switch "hello" {
                case "hello":
                    switch 127 {
                    case 0:
                        if (!true) {
                            var v34 string = "\n\t"
                            if true {
                                v35 := GenS2{F0: true, F1: 4294967295}
                                _ = v35
                            } else {
                                var v36 uint32 = 357
                                _ = v36
                            }
                            _ = v34
                        }
                    case 1:
                        switch 32767 {
                        case 42:
                            fmt.Println(v16)
                        }
                    }
                case "":
                    v2 = (v2 * v2)
                }
            }
            _ = v31
        }
    default:
        v37 := v3[(("" + "") + "a")]
        _ = v37
    }
    if (-1 < -1) {
        var v38 int16 = -1
        _ = v38
    } else {
        if (9223372036854775807 > 9223372036854775807) {
            v2 = (float32((1e10 - float64((v2 + 0.1)))) + (float32(0.1) / v2))
            v16 += (v16 - (v16 << ((42) & 63)))
            v8 -= 100
        } else {
            switch v16 {
            case 1:
                var v39 string = ""
                _ = v39
            case 4294967295:
                if len(v4) > 0 {
                    v4[0] = 2147483647
                }
            }
            var v40 uint32 = 1
            _ = v40
        }
        v16 += v16
    }
    v2 -= (1e10 + 3.14159)
    for fi41 := 0; fi41 < 1; fi41++ {
        var v42 uint64 = 42
        _ = fi41
        _ = v42
    }
    v16 = (v16 & v16)
    v43 := GenS2{F0: true, F1: 255}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v8
}
