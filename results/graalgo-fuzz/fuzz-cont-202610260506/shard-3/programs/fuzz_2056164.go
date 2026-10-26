package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = 727
    v3 := make(map[uint32]uint32)
    v4 := []uint32{0, 548}
    var v5 uint8 = 128
    var v6 int8 = int8(v5)
    fmt.Println(v6)
    switch "hello" {
    case "hello":
        var v7 int8 = 127
        _ = v7
    default:
        var v8 float64 = 1e10
        _ = v8
    }
    var v9 float32 = 3.14159
    v5 += (v5 << ((228) & 63))
    var v10 string = "a"
    switch -734 {
    case 128:
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := GenS1{F0: (v5 << ((18446744073709551615) & 63))}
            var v13 uint64 = 807
            _ = fi11
            _ = v12
            _ = v13
        }
    case 127:
        v14 := GenS2{F0: 0, F1: v5}
        _ = v14
    }
    v9 = (0.0 + ((1.5 - v9) + (1e10 / -2.25)))
    switch 32768 {
    case 32767:
        for fi15 := 0; fi15 < 8; fi15++ {
            v2 = v2
            switch ("hello" + (((v10 + v10) + v10) + ("" + v10))) {
            case "a":
                for fi16 := 0; fi16 < 12; fi16++ {
                    v17 := GenS2{F0: 1, F1: v5}
                    switch (fi15 >> ((18446744073709551615) & 63)) {
                    case 1:
                        if true {
                            var v18 int64 = 42
                            _ = v18
                        } else {
                            switch v10 {
                            case "":
                                v19 := GenS2{F0: 56, F1: (v5 >> ((1) & 63))}
                                var v20 int32 = 42
                                if (true && true) {
                                    if true {
                                        switch v10 {
                                        case "\n\t":
                                            fmt.Println(fi15, v10)
                                        default:
                                            fmt.Println(v10)
                                        }
                                        if true {
                                            var v21 int8 = 1
                                            _ = v21
                                        } else {
                                            v22 := GenS1{F0: 2}
                                            var v23 float64 = 0.1
                                            v9 = v9
                                            _ = v22
                                            _ = v23
                                        }
                                    }
                                    var v24 float64 = 1.5
                                    _ = v24
                                }
                                _ = v19
                                _ = v20
                            case "hello":
                                var v25 uint8 = (v5 % ((2) | 1))
                                _ = v25
                            case "a":
                                if false {
                                    var v26 uint = 1
                                    var v27 uint16 = 42
                                    v6 = (v6 >> ((v26) & 63))
                                    _ = v26
                                    _ = v27
                                } else {
                                    switch "" {
                                    case "hello":
                                        switch int(v5) {
                                        case 642:
                                            v2 = v2
                                        case 128:
                                            fmt.Println(fi16)
                                        }
                                    case "foo bar":
                                        if true {
                                            var v28 float32 = v9
                                            var v29 uint16 = 65535
                                            _ = v28
                                            _ = v29
                                        } else {
                                            if (42 >= 2) {
                                                if true {
                                                    v9 += 0.1
                                                    fmt.Println(fi16, fi16, v9)
                                                    _ = len(v1)
                                                } else {
                                                    v5 += (^(v5))
                                                    delete(v1, 65535)
                                                }
                                                fmt.Println(v5)
                                                fmt.Println(fi16, fi15)
                                                v30 := GenS2{F0: 1, F1: v5}
                                                _ = v30
                                            }
                                        }
                                    default:
                                        if true {
                                            if (1.5 == 1.5) {
                                                if true {
                                                    var v31 int = 4294967295
                                                    var v32 float32 = v9
                                                    _ = v31
                                                    _ = v32
                                                }
                                            } else {
                                                if true {
                                                    switch "hello" {
                                                    case "foo bar":
                                                        v9 -= ((0.0 + -2.25) - float32(((((1e10 * 0.0) * float64(float32(0.1))) + 1.5) / (((3.14159 / -2.25) - 1.5) + 0.0))))
                                                    case "":
                                                        switch fi16 {
                                                        case -1:
                                                            delete(v1, 441)
                                                        default:
                                                            if (!(v6 >= v6)) {
                                                                var v33 int16 = (v2 | v2)
                                                                _ = v33
                                                            } else {
                                                                if ((true && false) && (!(uint64(v5) < 0))) {
                                                                    v9 += (v9 + -2.25)
                                                                }
                                                                _ = len(v3)
                                                                var v34 string = (v10 + v10)
                                                                _ = v34
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                            var v35 float64 = 0.0
                                            var v36 float64 = v35
                                            _ = v35
                                            _ = v36
                                        }
                                    }
                                }
                            }
                            var v37 uint16 = 1
                            _ = v37
                        }
                    }
                    v2 = -601
                    if (65535 < 2) {
                        fmt.Println(fi15, v10)
                        fmt.Println(fi15, v10, v2)
                        v38 := GenS1{F0: 19}
                        _ = v38
                    } else {
                        switch v10 {
                        case "a":
                            v10 = ""
                        default:
                            fmt.Println(fi16, v10)
                        }
                        v5 += v5
                        var v39 bool = true
                        _ = v39
                    }
                    var v40 uint32 = 42
                    v9 = v9
                    v40 = (v40 & (^(v40)))
                    v41 := GenS2{F0: 874, F1: (v5 | 164)}
                    _ = fi16
                    _ = v17
                    _ = v40
                    _ = v41
                }
            }
            _ = fi15
        }
    case 809:
        if false {
            switch "\n\t" {
            case "hello":
                for fi42 := 0; fi42 < 10; fi42++ {
                    _ = len(v4)
                    var v43 string = (("" + ("hello" + v10)) + "foo bar")
                    var v44 float64 = float64(((v9 - (((v9 * -2.25) - v9) * -2.25)) / (v9 - v9)))
                    _ = fi42
                    _ = v43
                    _ = v44
                }
            case "a":
                v45 := GenS2{F0: 463, F1: (v5 + (^(v5)))}
                _ = v45
            case "foo bar":
                switch v10 {
                case "":
                    for fi46 := 0; fi46 < 1; fi46++ {
                        var v47 uint32 = 1
                        _ = fi46
                        _ = v47
                    }
                case "\n\t":
                    var v48 uint32 = 2147483647
                    _ = v48
                case "a":
                    delete(v3, 597)
                default:
                    var v49 uint64 = 18446744073709551615
                    _ = v49
                }
            default:
                var v50 int32 = 0
                _ = v50
            }
        }
    case 4294967295:
        var v51 uint16 = 255
        _ = v51
    }
    if (4294967295 > 9223372036854775807) {
        switch v10 {
        case "foo bar":
            v9 -= (v9 + (v9 - (v9 * ((v9 * 1e10) + (v9 / v9)))))
        case "\n\t":
            var v52 int32 = 2147483647
            _ = v52
        default:
            if ((false && (0 != 965)) || true) {
                var v53 float64 = 0.0
                _ = v53
            }
        }
        v54 := GenS2{F0: 199, F1: 128}
        _ = v54
    } else {
        for fi55 := 0; fi55 < 19; fi55++ {
            var v56 int8 = v6
            var v57 uint32 = 4294967295
            _ = fi55
            _ = v56
            _ = v57
        }
        _ = len(v4)
        var v58 int = 128
        _ = v58
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
