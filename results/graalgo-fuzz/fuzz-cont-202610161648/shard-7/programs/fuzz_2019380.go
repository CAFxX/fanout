package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 int32; F2 uint8 }
type GenS5 struct { F0 int32 }

func main() {
    var v1 uint64 = 42
    var v2 string = "hello"
    var v3 float32 = ((float32(1.5) - 0.1) / 1.5)
    var v4 int8 = 100
    v5 := GenS1{F0: v1, F1: 2}
    v6 := GenS2{F0: 127}
    fmt.Println(v3, v1, v4)
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint = 2
        v2 = ("foo bar" + v2)
        _ = fi7
        _ = v8
    }
    v3 = (v3 * (v3 - 0.0))
    switch 4294967295 {
    case 433:
        for fi9 := 0; fi9 < 13; fi9++ {
            if false {
                v1 = 990
                fmt.Println(v3)
                v3 = (1.5 / v3)
            } else {
                for fi10 := 0; fi10 < 13; fi10++ {
                    v11 := GenS3{F0: int64(v1)}
                    _ = fi10
                    _ = v11
                }
                for fi12 := 0; fi12 < 11; fi12++ {
                    switch (v2 + v2) {
                    case "":
                        var v13 bool = true
                        if (uint8(v1) <= 0) {
                            if (v4 >= v4) {
                                fmt.Println(v2, v3, v1)
                            }
                            v4 -= v4
                        } else {
                            v14 := GenS1{F0: 827, F1: 65535}
                            switch "" {
                            case "hello":
                                v1 -= 2
                                v3 -= (-2.25 * 0.0)
                            case "foo bar":
                                v1 += v1
                            default:
                                if false {
                                    v15 := []int{fi12}
                                    fmt.Println(fi9, v13, v1)
                                    _ = v15
                                } else {
                                    switch 4294967295 {
                                    case -894:
                                        if v13 {
                                            var v16 bool = (937 <= 538)
                                            v2 = ""
                                            _ = v16
                                        }
                                    case 2147483647:
                                        var v17 float64 = 0.0
                                        _ = v17
                                    default:
                                        v1 += (v1 << ((65535) & 63))
                                    }
                                    v4 = 127
                                    v1 -= (v1 >> ((988) & 63))
                                }
                            }
                            _ = v14
                        }
                        v18 := []int{fi12, (fi9 - 42)}
                        if len(v18) > 0 {
                            v18[0] = (^(fi9))
                        }
                        v18 = append(v18, fi12)
                        fmt.Println(fi9, fi12)
                        fmt.Println(v1)
                        switch (-(fi12)) {
                        case 917:
                            var v19 int = fi12
                            _ = v19
                        case 435:
                            v20 := GenS2{F0: 2147483647}
                            _ = v20
                        default:
                            fmt.Println(v1, v1)
                        }
                        _ = v13
                        _ = v18
                    case "\n\t":
                        v3 -= float32(-2.25)
                    case "hello":
                        v2 = v2
                    default:
                        if (436 >= -1) {
                            var v21 string = v2
                            _ = v21
                        } else {
                            if false {
                                var v22 uint32 = 4294967295
                                var v23 uint8 = 128
                                v24 := []int{fi12, (fi9 - fi9), fi9, fi9}
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                            var v25 uint16 = 2
                            var v26 int = int(v4)
                            v27 := GenS1{F0: 0, F1: 2147483648}
                            var v28 float32 = v3
                            var v29 int16 = -1
                            _ = v25
                            _ = v26
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                    }
                    var v30 float32 = v3
                    v31 := []int{}
                    v1 += v1
                    _ = fi12
                    _ = v30
                    _ = v31
                }
                for fi32 := 0; fi32 < 12; fi32++ {
                    if (false && false) {
                        v33 := []int{-1, fi9, fi9}
                        if len(v33) >= 2 {
                            v33 = v33[1:len(v33)]
                        }
                        switch fi32 {
                        case 2147483647:
                            if true {
                                v3 += 1e10
                            }
                        case 32767:
                            v34 := GenS2{F0: -236}
                            _ = v34
                        case 629:
                            var v35 int = -9
                            _ = v35
                        }
                        var v36 uint16 = 2
                        _ = v33
                        _ = v36
                    } else {
                        if (!(32767 > 2147483647)) {
                            if ((fi32 * fi9) >= (fi9 % (((fi9 << ((42) & 63))) | 1))) {
                                v3 = (v3 * v3)
                                var v37 uint16 = 2
                                var v38 int32 = 0
                                var v39 int32 = -1
                                _ = v37
                                _ = v38
                                _ = v39
                            } else {
                                v4 = (v4 << ((65535) & 63))
                                var v40 uint32 = 42
                                v40 -= (v40 - 2147483647)
                                _ = v40
                            }
                            var v41 bool = false
                            _ = v41
                        }
                        if false {
                            var v42 bool = (!false)
                            if v42 {
                                v1 = 0
                                v43 := GenS4{F0: (-1 >= -32768), F1: 982, F2: 0}
                                v3 += (v3 + v3)
                                _ = v43
                            }
                            _ = v42
                        }
                        v4 = 100
                        v44 := GenS5{F0: -2147483648}
                        fmt.Println(v4)
                        _ = v44
                    }
                    _ = fi32
                }
            }
            v45 := GenS4{F0: true, F1: 42, F2: 128}
            _ = fi9
            _ = v45
        }
    default:
        var v46 int64 = 0
        _ = v46
    }
    v47 := []int{128, 994, -9223372036854775808, 42}
    if true {
        v3 += v3
        for fi48 := 0; fi48 < 16; fi48++ {
            var v49 uint32 = 65535
            _ = fi48
            _ = v49
        }
    }
    for fi50 := 0; fi50 < 13; fi50++ {
        v3 = v3
        for fi51 := 0; fi51 < 16; fi51++ {
            v1 -= 1
            var v52 uint = 291
            _ = fi51
            _ = v52
        }
        _ = fi50
    }
    v53 := GenS3{F0: 42}
    v2 = ("hello" + (v2 + v2))
    _ = len(v47)
    var v54 uint16 = 32767
    v47 = append(v47, -2147483648)
    fmt.Println(v1, v54, v2)
    switch v2 {
    case "hello":
        v1 -= (^(v1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v53
    _ = v54
    _ = v6
}
