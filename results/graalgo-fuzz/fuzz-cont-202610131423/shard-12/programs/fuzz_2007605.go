package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16; F1 int }
type GenS7 struct { F0 int64; F1 string }
func genFn0(p0 int64, p1 string, p2 int32) int8 {
    var v5 int32 = (p2 << ((65535) & 63))
    switch p1 {
    case "hello":
        var v6 bool = (-9223372036854775808 < 32768)
        _ = v6
    case "\n\t":
        v7 := GenS1{F0: ("" + "foo bar"), F1: 42}
        _ = v7
    case "a":
        if true {
            if (0 <= 848) {
                var v8 uint16 = 65535
                for fi9 := 0; fi9 < 19; fi9++ {
                    var v10 int = (fi9 + 132)
                    if true {
                        var v11 string = p1
                        _ = v11
                    }
                    if true {
                        fmt.Println(v5, v8, fi9)
                        var v12 int = (fi9 - v10)
                        var v13 float32 = (0.0 - 0.0)
                        _ = v12
                        _ = v13
                    }
                    _ = fi9
                    _ = v10
                }
                _ = v8
            } else {
                fmt.Println(v5, p2)
                var v14 uint32 = 482
                if (-1 == 100) {
                    for fi15 := 0; fi15 < 10; fi15++ {
                        for fi16 := 0; fi16 < 15; fi16++ {
                            v17 := []int{}
                            switch "foo bar" {
                            case "a":
                                fmt.Println(p0, p0)
                            default:
                                p0 = -536
                            }
                            v18 := GenS2{F0: p0, F1: uint64(v14), F2: v14}
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                        _ = fi15
                    }
                } else {
                    var v19 int8 = 127
                    var v20 uint32 = (v14 * (v14 + (v14 << ((255) & 63))))
                    if (1 <= 2) {
                        v21 := []int{689, -9223372036854775808, 0}
                        v22 := GenS3{F0: (-(p0))}
                        var v23 int64 = (p0 * (p0 | p0))
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    _ = v19
                    _ = v20
                }
                _ = v14
            }
            p2 = 199
            var v24 uint16 = 65535
            _ = v24
        } else {
            if false {
                p2 -= (p2 | 375)
                if false {
                    v25 := []int{2147483647, 42}
                    v26 := GenS4{F0: 1, F1: p0}
                    _ = v25
                    _ = v26
                }
                var v27 int8 = 82
                _ = v27
            } else {
                var v28 int64 = (p0 / (((p0 ^ -982)) | 1))
                var v29 uint = 162
                fmt.Println(p1)
                _ = v28
                _ = v29
            }
            for fi30 := 0; fi30 < 14; fi30++ {
                for fi31 := 0; fi31 < 2; fi31++ {
                    fmt.Println(v5, p2, fi30)
                    _ = fi31
                }
                switch (fi30 << ((259) & 63)) {
                case 32767:
                    v32 := []int{(fi30 + (fi30 * fi30)), -1}
                    _ = v32
                case -607:
                    var v33 uint8 = 255
                    _ = v33
                case -9223372036854775808:
                    if true {
                        var v34 uint16 = 2
                        fmt.Println(p2, p2, v34)
                        if true {
                            p2 += v5
                            if (1 <= 1000) {
                                var v35 int32 = int32(p0)
                                switch (fi30 << ((262) & 63)) {
                                case 155:
                                    switch (fi30 ^ 128) {
                                    case 127:
                                        for fi36 := 0; fi36 < 17; fi36++ {
                                            v37 := GenS1{F0: p1, F1: 28}
                                            _ = fi36
                                            _ = v37
                                        }
                                    case -1:
                                        v38 := []int{fi30, fi30}
                                        _ = v38
                                    case 61:
                                        p1 = ""
                                    default:
                                        v39 := GenS5{F0: (p0 - (p0 % ((p0) | 1)))}
                                        _ = v39
                                    }
                                case 1:
                                    var v40 float64 = 3.14159
                                    _ = v40
                                case 128:
                                    fmt.Println(v35, v34)
                                default:
                                    var v41 int64 = 42
                                    _ = v41
                                }
                                v42 := GenS1{F0: p1, F1: 42}
                                _ = v35
                                _ = v42
                            } else {
                                if true {
                                    v43 := GenS6{F0: v34, F1: (fi30 * 2147483647)}
                                    _ = v43
                                }
                                p1 = "hello"
                            }
                        }
                        _ = v34
                    }
                default:
                    v44 := GenS7{F0: p0, F1: p1}
                    _ = v44
                }
                _ = fi30
            }
            for fi45 := 0; fi45 < 14; fi45++ {
                var v46 float64 = 0.1
                v47 := []int{32768, 98}
                _ = fi45
                _ = v46
                _ = v47
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 127
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = ""
    var v3 uint8 = 110
    var v4 uint64 = 2
    v48 := []int{}
    switch v2 {
    case "foo bar":
        if false {
            v2 = ""
            var v49 float64 = 0.0
            var v50 int8 = int8(v4)
            _ = v49
            _ = v50
        }
    case "a":
        if len(v48) >= 2 {
            v48 = v48[1:len(v48)]
        }
    default:
        var v51 uint = 42
        _ = v51
    }
    for fi52 := 0; fi52 < 16; fi52++ {
        var v53 int = -266
        v2 = ("" + (v2 + "a"))
        _ = fi52
        _ = v53
    }
    v54 := GenS1{F0: "", F1: -92}
    switch 32768 {
    case -352:
        if (936 <= 32767) {
            var v55 string = v2
            _ = v55
        } else {
            for fi56 := 0; fi56 < 9; fi56++ {
                if (v3 == (^(v3))) {
                    v1[1] = v2
                    v3 = v3
                    var v57 int16 = 1
                    _ = v57
                }
                _ = fi56
            }
        }
    case 127:
        fmt.Println(v3, v2)
    case 32768:
        var v58 uint32 = 42
        _ = v58
    }
    _ = len(v1)
    var v59 uint32 = 1
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v54
    _ = v59
}
