package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint; F2 int16 }
type GenS4 struct { F0 int; F1 uint64 }

func main() {
    var v1 uint32 = 873
    var v2 uint16 = 1
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]uint32)
    var v5 uint = 1
    var v6 int8 = 0
    v6 = (^(v6))
    if (128 <= 42) {
        v2 = (v2 >> ((uint(v1)) & 63))
        var v7 uint16 = v2
        _ = v7
    } else {
        v8 := []int{2147483647, int(v6), 42}
        _ = v8
    }
    var v9 int32 = 42
    if true {
        var v10 int = -2147483648
        if false {
            var v11 uint = (v5 << (((v5 / (((v5 << (((v5 + 951)) & 63))) | 1))) & 63))
            _ = v11
        } else {
            v9 += v9
            v12 := GenS1{F0: v1, F1: (v2 & 42)}
            v13 := []int{}
            _ = v12
            _ = v13
        }
        _ = v10
    } else {
        v14 := GenS2{F0: 50, F1: (v5 % ((482) | 1)), F2: (v6 - (v6 + (v6 & v6)))}
        v5 += v5
        if false {
            for fi15 := 0; fi15 < 12; fi15++ {
                v16 := GenS2{F0: 110, F1: 42, F2: v6}
                _ = len(v4)
                _ = fi15
                _ = v16
            }
        } else {
            var v17 int16 = 42
            if (0 < v9) {
                switch "foo bar" {
                case "foo bar":
                    v6 -= v6
                case "":
                    v5 -= v5
                }
                switch 4294967295 {
                case 824:
                    v6 = -36
                case -513:
                    v9 += (-(v9))
                }
                var v18 int8 = (v6 << (((v5 >> ((uint(v2)) & 63))) & 63))
                _ = v18
            }
            fmt.Println(v1)
            _ = v17
        }
        _ = v14
    }
    if false {
        switch ((("foo bar" + "a") + "") + "\n\t") {
        case "hello":
            var v19 uint = (v5 + (v5 + (v5 % ((65535) | 1))))
            _ = v19
        case "\n\t":
            _ = len(v3)
        }
    } else {
        for fi20 := 0; fi20 < 5; fi20++ {
            var v21 uint16 = 307
            if false {
                if true {
                    v21 += (v2 * (v2 - 1))
                    v5 -= 1
                }
                var v22 int16 = 1
                _ = v22
            } else {
                if true {
                    if (fi20 == (fi20 << (((v5 >> ((65535) & 63))) & 63))) {
                        v5 += v5
                        if ((fi20 ^ fi20) != (fi20 % (((fi20 | fi20)) | 1))) {
                            for fi23 := 0; fi23 < 10; fi23++ {
                                fmt.Println(v1, v21)
                                v24 := []int{}
                                _ = fi23
                                _ = v24
                            }
                            var v25 int64 = 1
                            _ = v25
                        } else {
                            for fi26 := 0; fi26 < 11; fi26++ {
                                if (0.1 < 0.1) {
                                    if true {
                                        var v27 float64 = 1e10
                                        var v28 uint8 = 127
                                        var v29 uint8 = 42
                                        _ = v27
                                        _ = v28
                                        _ = v29
                                    }
                                    var v30 int = 0
                                    _ = v30
                                }
                                v31 := GenS1{F0: uint32(v5), F1: v2}
                                _ = fi26
                                _ = v31
                            }
                        }
                        v9 += -2147483648
                    } else {
                        fmt.Println(v6, v2, v1)
                        _ = len(v4)
                        v32 := v4[4294967295]
                        _ = v32
                    }
                    v33 := []int{(fi20 >> ((v5) & 63)), (fi20 ^ (fi20 >> ((v5) & 63))), (-(fi20))}
                    _ = v33
                } else {
                    v34 := GenS3{F0: 255, F1: 255, F2: int16(fi20)}
                    if (!((!((v6 % (((v6 / ((100) | 1))) | 1)) != (v6 ^ v6))) && ((^(v5)) >= (v5 & (^(v5)))))) {
                        for fi35 := 0; fi35 < 13; fi35++ {
                            fmt.Println(fi35)
                            v36 := GenS3{F0: 127, F1: 946, F2: -532}
                            switch "foo bar" {
                            case "foo bar":
                                if true {
                                    v37 := []int{fi35}
                                    var v38 int32 = v9
                                    if len(v37) > 0 {
                                        v37[0] = fi20
                                    }
                                    _ = v37
                                    _ = v38
                                } else {
                                    if true {
                                        v5 = v5
                                    } else {
                                        var v39 uint64 = 2
                                        _ = v39
                                    }
                                }
                            case "\n\t":
                                _ = len(v3)
                            default:
                                var v40 uint32 = 0
                                _ = v40
                            }
                            _ = fi35
                            _ = v36
                        }
                    } else {
                        var v41 int = (fi20 + fi20)
                        _ = v41
                    }
                    v42 := GenS4{F0: fi20, F1: 4294967295}
                    _ = v34
                    _ = v42
                }
                var v43 int8 = v6
                for fi44 := 0; fi44 < 13; fi44++ {
                    v21 -= (v21 + (^(v2)))
                    _ = fi44
                }
                _ = v43
            }
            _ = fi20
            _ = v21
        }
        var v45 uint8 = uint8(v1)
        _ = v45
    }
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
