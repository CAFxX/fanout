package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int, p1 int) int16 {
    switch ((((("hello" + "") + "\n\t") + ("\n\t" + "")) + "a") + "hello") {
    case "":
        p1 -= p1
    case "hello":
        var v6 float64 = 0.0
        _ = v6
    default:
        var v7 int = (p0 & 951)
        _ = v7
    }
    switch -2147483648 {
    case 127:
        if (!true) {
            var v8 uint16 = 65535
            _ = v8
        } else {
            if false {
                for fi9 := 0; fi9 < 14; fi9++ {
                    fmt.Println(p1, fi9)
                    _ = fi9
                }
            } else {
                for fi10 := 0; fi10 < 11; fi10++ {
                    if (-1 < -1) {
                        v11 := GenS1{F0: p1}
                        var v12 float32 = 1e10
                        _ = v11
                        _ = v12
                    } else {
                        if true {
                            v13 := GenS1{F0: 1}
                            _ = v13
                        }
                    }
                    _ = fi10
                }
                p0 = (p0 | 32767)
            }
            var v14 int16 = -32768
            _ = v14
        }
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = 1e10
    v3 := []uint32{2147483648, 0}
    v4 := []int{843}
    var v5 float32 = -2.25
    v2 += 1e10
    for fi15 := 0; fi15 < 19; fi15++ {
        for fi16 := 0; fi16 < 4; fi16++ {
            switch "\n\t" {
            case "foo bar":
                var v17 int16 = int16(fi15)
                _ = v17
            case "hello":
                var v18 int64 = -1
                _ = v18
            default:
                v5 = 1e10
            }
            var v19 uint8 = 255
            _ = fi16
            _ = v19
        }
        v5 -= v5
        var v20 uint16 = 1
        _ = fi15
        _ = v20
    }
    v3 = append(v3, 2147483648)
    if false {
        switch 477 {
        case 1:
            var v21 uint32 = 705
            _ = v21
        case -809:
            delete(v1, 213)
        default:
            var v22 bool = false
            _ = v22
        }
        fmt.Println(v5, v5)
    }
    v23 := GenS1{F0: 0}
    if true {
        var v24 int8 = -82
        _ = v24
    }
    v25 := GenS1{F0: 780}
    v2 = -2.25
    for fi26 := 0; fi26 < 1; fi26++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        switch fi26 {
        case -35:
            switch fi26 {
            case 42:
                v27 := GenS1{F0: (fi26 / ((fi26) | 1))}
                _ = v27
            case 32767:
                switch 520 {
                case 128:
                    v28 := v1[65535]
                    _ = v28
                default:
                    v2 = -2.25
                }
            case -1:
                if (!(v2 < (((v2 / 1e10) - 0.1) * v2))) {
                    _ = len(v1)
                    for fi29 := 0; fi29 < 20; fi29++ {
                        switch fi29 {
                        case 2147483647:
                            switch fi29 {
                            case 2147483647:
                                v1[0] = 0
                            case 32768:
                                v30 := GenS1{F0: fi29}
                                _ = v30
                            default:
                                v5 += 1.5
                            }
                        case -168:
                            if ((527 == 42) || true) {
                                v5 = (0.0 + -2.25)
                            } else {
                                var v31 int64 = 9223372036854775807
                                v5 = (v5 / v5)
                                if (1 != 42) {
                                    if false {
                                        v31 = -958
                                    } else {
                                        v2 -= (v2 - v2)
                                        v31 += v31
                                        v2 = v2
                                    }
                                    v32 := GenS1{F0: (fi29 + fi26)}
                                    v5 -= float32((v2 + (((v2 - v2) - v2) * 0.0)))
                                    _ = v32
                                } else {
                                    delete(v1, 752)
                                    var v33 uint32 = 65535
                                    _ = v33
                                }
                                _ = v31
                            }
                        default:
                            _ = len(v3)
                        }
                        var v34 int32 = 42
                        _ = fi29
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 7; fi35++ {
                        switch "" {
                        case "\n\t":
                            var v36 int64 = 9223372036854775807
                            _ = v36
                        }
                        v5 += float32(v2)
                        _ = fi35
                    }
                } else {
                    fmt.Println(v5)
                }
            }
        case 42:
            fmt.Println(v5, fi26, v5)
        }
        _ = fi26
    }
    var v37 float64 = float64(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
