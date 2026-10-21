package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 int64; F1 int8; F2 uint }
type GenS5 struct { F0 int8; F1 uint32; F2 int }
type GenS6 struct { F0 int16; F1 int8; F2 int16 }
type GenS7 struct { F0 uint16 }

func main() {
    var v1 uint64 = 515
    var v2 float32 = float32(1.5)
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    v5 := []uint32{851, 2, 2}
    var v6 float64 = (((((-2.25 * 0.0) / (-2.25 / 1.0)) - (1.5 * float64(v2))) * 1.5) - float64(v2))
    var v7 string = "a"
    switch v7 {
    case "":
        for fi8 := 0; fi8 < 9; fi8++ {
            switch (fi8 << ((42) & 63)) {
            case -958:
                fmt.Println(v2, fi8)
            case -1:
                switch (fi8 << ((1) & 63)) {
                case 219:
                    v7 = v7
                }
            case 42:
                switch (^(fi8)) {
                case -1:
                    fmt.Println(fi8, v2, v6)
                }
            default:
                if (0 < 95) {
                    if true {
                        for fi9 := 0; fi9 < 15; fi9++ {
                            switch (fi9 & fi9) {
                            case -660:
                                v2 += v2
                            case 32768:
                                var v10 int8 = 24
                                v10 = v10
                                _ = v10
                            case -82:
                                var v11 bool = (1e10 > ((1e10 * v2) / v2))
                                var v12 float32 = (1.5 - 3.14159)
                                fmt.Println(v1, fi9)
                                _ = v11
                                _ = v12
                            }
                            v1 = (v1 | 0)
                            var v13 int32 = -269
                            _ = fi9
                            _ = v13
                        }
                    } else {
                        var v14 float64 = float64(v2)
                        var v15 int32 = 1
                        var v16 float32 = v2
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    v7 = v7
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            }
            v7 = "foo bar"
            _ = fi8
        }
    case "foo bar":
        v7 = v7
    default:
        if len(v5) > 0 {
            v5[0] = 360
        }
    }
    if len(v5) > 0 {
        v5[0] = 1
    }
    if (4294967295 <= 128) {
        switch 9223372036854775807 {
        case 1:
            v7 = v7
        default:
            v1 = 2
        }
    }
    var v17 int32 = 1
    v18 := GenS1{F0: 0}
    var v19 string = v7
    v2 = (v2 - 0.0)
    var v20 int16 = 104
    switch -357 {
    case 0:
        v17 = (-(v17))
    case 32767:
        if (true || false) {
            for fi21 := 0; fi21 < 15; fi21++ {
                var v22 string = (v7 + "\n\t")
                _ = fi21
                _ = v22
            }
            fmt.Println(v7, v1, v7)
        } else {
            v6 -= v6
            v23 := GenS2{F0: (v1 / ((v1) | 1)), F1: v1}
            _ = v23
        }
    case 4:
        v4[0] = 0
    }
    fmt.Println(v19, v17)
    switch ("hello" + v7) {
    case "\n\t":
        fmt.Println(v7, v20, v20)
    case "foo bar":
        var v24 uint8 = 255
        _ = v24
    case "":
        fmt.Println(v17, v2, v6)
    }
    switch -573 {
    case -799:
        if (true || true) {
            var v25 bool = false
            for fi26 := 0; fi26 < 11; fi26++ {
                var v27 float32 = v2
                for fi28 := 0; fi28 < 9; fi28++ {
                    if (true || v25) {
                        var v29 int64 = 695
                        v30 := GenS3{F0: 42, F1: 2147483647}
                        _ = v29
                        _ = v30
                    } else {
                        var v31 float32 = v27
                        var v32 uint64 = v1
                        _ = len(v4)
                        _ = v31
                        _ = v32
                    }
                    _ = fi28
                }
                var v33 uint8 = 0
                _ = fi26
                _ = v27
                _ = v33
            }
            _ = v25
        }
    }
    switch 0 {
    case 0:
        if true {
            var v34 float32 = (v2 * 0.0)
            if len(v5) > 0 {
                v5[0] = 42
            }
            delete(v3, int(v17))
            _ = v34
        } else {
            switch v7 {
            case "hello":
                if (true && false) {
                    fmt.Println(v7)
                }
            }
            v7 = v7
            v19 = v7
        }
    case -1:
        v7 = ""
    default:
        if false {
            if true {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v35 int = -9223372036854775808
                _ = v35
            } else {
                var v36 uint16 = 159
                fmt.Println(v19, v17, v2)
                if (!false) {
                    switch "\n\t" {
                    case "\n\t":
                        v6 = v6
                    default:
                        v37 := GenS4{F0: 934, F1: -1, F2: 4294967295}
                        _ = v37
                    }
                    v38 := GenS5{F0: -1, F1: 65535, F2: -1}
                    var v39 uint = 209
                    _ = v38
                    _ = v39
                } else {
                    v40 := GenS6{F0: 0, F1: 100, F2: (v20 << ((101) & 63))}
                    _ = v40
                }
                _ = v36
            }
        }
    }
    v1 += v1
    for fi41 := 0; fi41 < 5; fi41++ {
        var v42 uint8 = 127
        _ = fi41
        _ = v42
    }
    v43 := GenS7{F0: 805}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
}
