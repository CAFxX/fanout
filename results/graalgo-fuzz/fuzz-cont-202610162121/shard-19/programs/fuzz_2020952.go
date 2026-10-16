package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint; F2 int64 }
type GenS3 struct { F0 uint32; F1 string }
func genFn0(p0 string, p1 int64) uint32 {
    var v5 uint = 287
    for fi6 := 0; fi6 < 5; fi6++ {
        if (int16(v5) != 819) {
            v7 := []int{}
            _ = v7
        }
        var v8 string = p0
        var v9 uint64 = 1
        _ = fi6
        _ = v8
        _ = v9
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v5
    return 2147483648
}
func genFn1(p0 uint8, p1 int64) int32 {
    v10 := GenS1{F0: 65535}
    if true {
        p0 = p0
    }
    p0 -= (p0 >> ((65535) & 63))
    _ = p0
    _ = p1
    _ = v10
    return 2147483647
}

func main() {
    var v1 uint16 = 255
    var v2 string = "\n\t"
    var v3 uint16 = 32767
    var v4 int32 = 0
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 int16 = 1000
        v13 := []int{}
        v13 = append(v13, -1)
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := []int{611, 46, 127, 285}
    v1 -= v1
    fmt.Println(v4, v2)
    switch v2 {
    case "":
        v4 = v4
    case "foo bar":
        v15 := GenS1{F0: 461}
        _ = v15
    case "\n\t":
        var v16 int16 = 0
        _ = v16
    }
    var v17 bool = false
    switch 4294967295 {
    case -600:
        for fi18 := 0; fi18 < 17; fi18++ {
            if (((v3 % ((65535) | 1)) < (v1 + 255)) && ((false || (!v17)) && true)) {
                if len(v14) > 0 {
                    v14[0] = fi18
                }
                if v17 {
                    if true {
                        if len(v14) > 0 {
                            v14[0] = (fi18 | -2147483648)
                        }
                        if len(v14) > 0 {
                            v14[0] = (fi18 + fi18)
                        }
                    } else {
                        if v17 {
                            if len(v14) > 0 {
                                v14[0] = fi18
                            }
                            fmt.Println(v4, v1, fi18)
                        }
                    }
                    if v17 {
                        _ = len(v14)
                    } else {
                        switch fi18 {
                        case 2147483647:
                            var v19 string = v2
                            _ = v19
                        default:
                            var v20 bool = true
                            _ = v20
                        }
                    }
                }
                if len(v14) > 0 {
                    v14[0] = (-(fi18))
                }
            }
            _ = len(v14)
            for fi21 := 0; fi21 < 4; fi21++ {
                fmt.Println(v2, v3, v4)
                v14 = append(v14, (fi21 % ((fi21) | 1)))
                _ = fi21
            }
            _ = fi18
        }
    case 42:
        _ = len(v14)
    }
    if false {
        if v17 {
            var v22 int64 = 0
            v22 = (v22 << ((454) & 63))
            _ = v22
        }
    }
    var v23 float32 = (((1.5 / (1.5 * 1.5)) / -2.25) - (-2.25 / (0.0 + 3.14159)))
    v14 = append(v14, 127)
    if (0 != 4294967295) {
        switch 2147483647 {
        case 533:
            v24 := GenS2{F0: v1, F1: 42, F2: 0}
            _ = v24
        default:
            fmt.Println(v2)
        }
    } else {
        switch 2147483647 {
        case 32767:
            for fi25 := 0; fi25 < 18; fi25++ {
                if (v3 == (v1 + v3)) {
                    fmt.Println(v2, fi25, v3)
                    v26 := GenS3{F0: 2147483648, F1: ""}
                    if (true && (!false)) {
                        v1 = 0
                        v1 = (^(v3))
                    } else {
                        fmt.Println(v3, v23, v1)
                    }
                    _ = v26
                }
                v27 := GenS1{F0: 584}
                _ = fi25
                _ = v27
            }
        case -503:
            fmt.Println(v4, v17, v2)
        default:
            if (v17 && true) {
                var v28 int16 = 0
                v29 := GenS3{F0: 2, F1: "a"}
                if len(v14) > 0 {
                    v14[0] = 127
                }
                _ = v28
                _ = v29
            }
        }
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        var v31 float64 = float64(v23)
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
