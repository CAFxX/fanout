package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint32; F1 int64 }

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    v4 := []uint32{2, 1, 65535}
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 uint16 = 961
        v1 -= 2
        _ = fi5
        _ = v6
    }
    var v7 uint16 = 0
    if false {
        v8 := GenS1{F0: 4294967295, F1: 0}
        v1 = v1
        for fi9 := 0; fi9 < 15; fi9++ {
            for fi10 := 0; fi10 < 13; fi10++ {
                var v11 int32 = 2147483647
                var v12 int = 1
                var v13 uint = 255
                var v14 string = ("foo bar" + ("" + ""))
                var v15 string = v14
                if false {
                    v1 = v1
                    fmt.Println(v7)
                } else {
                    v11 += 42
                    if true {
                        var v16 float32 = 3.14159
                        var v17 int8 = 42
                        switch fi9 {
                        case 127:
                            v13 -= uint(v7)
                            var v18 bool = true
                            v13 += v13
                            _ = v18
                        case 231:
                            v14 = v14
                        default:
                            fmt.Println(v13)
                        }
                        _ = v16
                        _ = v17
                    } else {
                        v7 -= 941
                        v15 = v14
                        var v19 int = (v12 | (-(fi10)))
                        _ = v19
                    }
                    if true {
                        v13 -= v13
                        fmt.Println(v11, v13, fi10)
                    } else {
                        v1 += (^(v1))
                    }
                    if (false && (166 <= 2)) {
                        if false {
                            var v20 int = fi10
                            var v21 uint64 = (v1 * v1)
                            v7 += v7
                            _ = v20
                            _ = v21
                        } else {
                            v15 = "\n\t"
                        }
                        v14 = (v14 + v15)
                        v11 = -251
                    }
                    if (-320 <= 0) {
                        v7 -= v7
                        v13 -= v13
                        var v22 int64 = -9223372036854775808
                        _ = v22
                    } else {
                        var v23 uint = 42
                        v4 = append(v4, uint32(v1))
                        fmt.Println(v12, fi9)
                        if (3.14159 > float64(0.1)) {
                            fmt.Println(v12, v15)
                        } else {
                            v24 := GenS1{F0: 888, F1: 1}
                            _ = len(v3)
                            if false {
                                if (v15 > v15) {
                                    switch v14 {
                                    case "foo bar":
                                        fmt.Println(v12)
                                    case "a":
                                        v11 += v11
                                    }
                                    if (2 != 103) {
                                        var v25 int16 = -32768
                                        var v26 int32 = -306
                                        _ = v25
                                        _ = v26
                                    }
                                    if false {
                                        if true {
                                            v1 = v1
                                            var v27 int16 = -1
                                            if (1 > (^(v12))) {
                                                var v28 bool = false
                                                _ = v28
                                            }
                                            _ = v27
                                        }
                                        v11 -= v11
                                        var v29 int32 = -2147483648
                                        _ = v29
                                    }
                                }
                                v2[2] = (v12 - 418)
                                v30 := GenS1{F0: (^(v1)), F1: 1}
                                _ = v30
                            }
                            _ = v24
                        }
                        _ = v23
                    }
                }
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
                _ = v14
                _ = v15
            }
            _ = fi9
        }
        _ = v8
    } else {
        switch 4294967295 {
        case -746:
            var v31 int = 1
            _ = v31
        case 9223372036854775807:
            v1 = v1
        case 2147483647:
            v32 := v3[2147483647]
            _ = v32
        default:
            var v33 int64 = -1
            _ = v33
        }
        fmt.Println(v1)
    }
    v34 := GenS2{F0: int(v1), F1: 1}
    if (-701 >= 42) {
        v1 += v1
        var v35 uint8 = 0
        _ = v35
    }
    if (!false) {
        fmt.Println(v1, v7)
        for fi36 := 0; fi36 < 15; fi36++ {
            v1 += (^(v1))
            _ = fi36
        }
        v37 := GenS3{F0: 255, F1: -9223372036854775808}
        _ = v37
    } else {
        if true {
            var v38 int8 = 42
            for fi39 := 0; fi39 < 2; fi39++ {
                var v40 int32 = 1
                v40 -= 2147483647
                var v41 int32 = (v40 << ((586) & 63))
                _ = fi39
                _ = v40
                _ = v41
            }
            var v42 bool = true
            _ = v38
            _ = v42
        }
        var v43 float64 = 1.5
        v7 -= (v7 + 483)
        _ = v43
    }
    delete(v3, -99)
    v44 := GenS1{F0: v1, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v44
    _ = v7
}
