package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint16; F2 bool }
type GenS4 struct { F0 uint32; F1 int16; F2 int }
type GenS5 struct { F0 uint16; F1 uint64 }
type GenS6 struct { F0 bool; F1 string; F2 int8 }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 uint64; F1 uint8 }
type GenS9 struct { F0 int16; F1 uint32 }
func genFn0(p0 int32, p1 uint32, p2 uint64) uint64 {
    fmt.Println(p0)
    v6 := []int{}
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 int8 = -1
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        if (-128 == -27) {
            v10 := GenS1{F0: (p2 <= (p2 >> ((42) & 63)))}
            p2 = p2
            _ = v10
        }
        v11 := GenS2{F0: 2}
        p0 -= (p0 | 1)
        _ = fi9
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    var v3 float32 = -2.25
    var v4 uint = 18446744073709551615
    var v5 int16 = -456
    v12 := GenS3{F0: 42, F1: 98, F2: (-115 == 40)}
    var v13 uint16 = 358
    var v14 uint = 65535
    var v15 uint8 = 127
    var v16 bool = false
    v17 := []int{1, -284, 9223372036854775807}
    if (!true) {
        fmt.Println(v4, v13, v13)
        switch -9223372036854775808 {
        case -234:
            switch "a" {
            case "foo bar":
                if (true || false) {
                    v4 -= (v4 >> ((v4) & 63))
                    if v16 {
                        for fi18 := 0; fi18 < 16; fi18++ {
                            _ = len(v17)
                            _ = fi18
                        }
                        if ((v16 || (int32(v13) > 0)) || true) {
                            var v19 float64 = -2.25
                            v1[-1] = -607
                            _ = v19
                        } else {
                            if v16 {
                                v20 := GenS4{F0: 65535, F1: (v5 | 1000), F2: 42}
                                _ = v20
                            }
                            var v21 uint = v14
                            _ = v21
                        }
                    }
                    for fi22 := 0; fi22 < 10; fi22++ {
                        v23 := GenS4{F0: 822, F1: v5, F2: (fi22 + (fi22 | 2147483647))}
                        v24 := v2[uint32(v13)]
                        switch (fi22 % (((fi22 % (((fi22 + 42)) | 1))) | 1)) {
                        case 127:
                            for fi25 := 0; fi25 < 7; fi25++ {
                                if (18446744073709551615 != 4294967295) {
                                    _ = len(v2)
                                } else {
                                    v26 := GenS5{F0: v13, F1: 0}
                                    var v27 int64 = 42
                                    var v28 uint32 = 45
                                    var v29 uint8 = 42
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                                if (!v16) {
                                    switch (v24 + "\n\t") {
                                    case "":
                                        switch (fi22 | fi25) {
                                        case 895:
                                            var v30 int32 = 712
                                            _ = v30
                                        case 9223372036854775807:
                                            var v31 uint8 = v15
                                            _ = v31
                                        case -9223372036854775808:
                                            var v32 int8 = 100
                                            _ = v32
                                        default:
                                            var v33 bool = (v24 == v24)
                                            _ = v33
                                        }
                                    case "foo bar":
                                        fmt.Println(v15, fi25)
                                    default:
                                        fmt.Println(fi25)
                                    }
                                } else {
                                    v5 -= (v5 | -368)
                                }
                                _ = len(v2)
                                _ = fi25
                            }
                        default:
                            var v34 uint = (v14 | 232)
                            _ = v34
                        }
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                }
            default:
                v15 += v15
            }
        case 128:
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
        default:
            v13 -= (v13 / ((660) | 1))
        }
        v4 = (v4 % ((v14) | 1))
    }
    for fi35 := 0; fi35 < 13; fi35++ {
        if (1 <= 0) {
            fmt.Println(v16, v16)
        } else {
            for fi36 := 0; fi36 < 1; fi36++ {
                if ((int32(v13) <= -2147483648) && true) {
                    if (v15 <= (v15 % ((v15) | 1))) {
                        v16 = true
                        v37 := GenS6{F0: ((!false) && ((v14 & 42) > (v14 ^ v4))), F1: (("" + "hello") + ("hello" + ("foo bar" + ("" + "\n\t")))), F2: -128}
                        if true {
                            v38 := GenS7{F0: v15, F1: 726}
                            _ = v38
                        }
                        _ = v37
                    } else {
                        if v16 {
                            _ = len(v17)
                            var v39 uint64 = uint64(v14)
                            var v40 float64 = -2.25
                            _ = v39
                            _ = v40
                        }
                        var v41 int32 = -22
                        var v42 uint = v14
                        v43 := GenS3{F0: -103, F1: (v13 & v13), F2: v16}
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                }
                v44 := GenS8{F0: genFn0(-2147483648, 382, 18446744073709551615), F1: 0}
                _ = fi36
                _ = v44
            }
        }
        _ = fi35
    }
    if v16 {
        v45 := GenS9{F0: v5, F1: 1}
        var v46 uint64 = 2
        _ = v45
        _ = v46
    }
    v47 := GenS7{F0: (v15 * (v15 ^ v15)), F1: 239}
    var v48 int16 = 475
    for fi49 := 0; fi49 < 3; fi49++ {
        fmt.Println(v16, fi49, v16)
        var v50 float64 = 0.1
        v5 -= 1000
        _ = fi49
        _ = v50
    }
    v15 -= (v15 << ((v14) & 63))
    for fi51 := 0; fi51 < 4; fi51++ {
        for fi52 := 0; fi52 < 6; fi52++ {
            if false {
                v15 += v15
                v5 -= int16(v4)
                var v53 uint8 = v15
                _ = v53
            }
            var v54 uint32 = 42
            _ = fi52
            _ = v54
        }
        var v55 bool = (v3 <= v3)
        _ = fi51
        _ = v55
    }
    v15 = (v15 | v15)
    fmt.Println(v13, v14)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
}
