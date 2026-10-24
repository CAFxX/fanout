package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 int; F2 bool }
type GenS4 struct { F0 int8; F1 int16; F2 uint8 }
type GenS5 struct { F0 uint16; F1 int8; F2 int }
type GenS6 struct { F0 uint; F1 int8; F2 uint64 }
type GenS7 struct { F0 int32; F1 uint64 }
type GenS8 struct { F0 bool; F1 int32; F2 uint16 }
type GenS9 struct { F0 int64; F1 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 int = -835
    var v3 bool = (false || (255 > 42))
    v4 := []int{}
    var v5 uint8 = 79
    delete(v1, 127)
    for fi6 := 0; fi6 < 7; fi6++ {
        if false {
            v3 = (!false)
            _ = len(v4)
        }
        for fi7 := 0; fi7 < 13; fi7++ {
            fmt.Println(v2, v5, v2)
            if false {
                v3 = (814 <= 65535)
            } else {
                v8 := GenS1{F0: 2147483648, F1: 2, F2: (^(v5))}
                v9 := v1[fi6]
                _ = v8
                _ = v9
            }
            v1[fi6] = 42
            var v10 uint16 = 32767
            if (v3 && false) {
                if len(v4) > 0 {
                    v4[0] = (v2 ^ (fi7 << ((2) & 63)))
                }
                var v11 uint16 = v10
                v4 = append(v4, (^(v2)))
                _ = v11
            }
            _ = fi7
            _ = v10
        }
        _ = fi6
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := GenS2{F0: 202}
        v4 = append(v4, fi12)
        _ = fi12
        _ = v13
    }
    if v3 {
        if false {
            for fi14 := 0; fi14 < 9; fi14++ {
                var v15 int16 = 1
                _ = fi14
                _ = v15
            }
        } else {
            var v16 float64 = ((float64(1.5) * 1e10) + (1e10 - 1.5))
            v17 := GenS3{F0: (723 <= 0), F1: v2, F2: ((!(0 >= 1)) && (("hello" == "foo bar") && v3))}
            if v3 {
                var v18 float64 = v16
                if ((-119 > -130) || (v3 || false)) {
                    for fi19 := 0; fi19 < 7; fi19++ {
                        var v20 int64 = -812
                        if false {
                            if v3 {
                                for fi21 := 0; fi21 < 11; fi21++ {
                                    if len(v4) > 0 {
                                        v4[0] = (fi19 / ((32768) | 1))
                                    }
                                    if false {
                                        switch (v2 | fi19) {
                                        case -2147483648:
                                            v22 := GenS4{F0: 127, F1: -1, F2: 247}
                                            _ = v22
                                        }
                                    }
                                    _ = fi21
                                }
                                v20 = 9223372036854775807
                            } else {
                                if true {
                                    fmt.Println(v2)
                                } else {
                                    var v23 uint32 = 856
                                    v23 = (v23 << ((18446744073709551615) & 63))
                                    _ = v23
                                }
                            }
                            fmt.Println(v20)
                        }
                        _ = fi19
                        _ = v20
                    }
                    v1[v2] = 2147483647
                } else {
                    for fi24 := 0; fi24 < 17; fi24++ {
                        var v25 int16 = -1
                        fmt.Println(v25, v16)
                        _ = fi24
                        _ = v25
                    }
                }
                var v26 uint32 = 699
                _ = v18
                _ = v26
            } else {
                var v27 uint16 = 42
                _ = v27
            }
            _ = v16
            _ = v17
        }
        v28 := GenS5{F0: 1, F1: 42, F2: (v2 & 42)}
        _ = v28
    }
    if len(v4) > 0 {
        v4[0] = (v2 % (((v2 ^ v2)) | 1))
    }
    v29 := GenS3{F0: v3, F1: (-(v2)), F2: (!((!(-838 > 0)) && v3))}
    for fi30 := 0; fi30 < 3; fi30++ {
        v31 := GenS6{F0: 613, F1: 1, F2: 388}
        fmt.Println(v5, v5)
        v32 := GenS7{F0: 2147483647, F1: 1}
        _ = fi30
        _ = v31
        _ = v32
    }
    _ = len(v1)
    v4 = append(v4, 885)
    switch v2 {
    case 9223372036854775807:
        v3 = v3
    default:
        v3 = (!(!(int(v5) >= 127)))
    }
    v2 += (v2 % ((v2) | 1))
    switch v2 {
    case 128:
        v33 := v1[v2]
        _ = v33
    case 127:
        _ = len(v1)
    default:
        v3 = (0 != 1)
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        v2 += (-(v2))
        for fi35 := 0; fi35 < 8; fi35++ {
            _ = len(v4)
            var v36 int32 = 1
            var v37 uint8 = 171
            if v3 {
                fmt.Println(v2)
                v38 := GenS8{F0: (65535 <= 385), F1: (v36 >> ((348) & 63)), F2: 74}
                switch fi34 {
                case 1:
                    if ((42 <= 179) || true) {
                        var v39 uint32 = 42
                        var v40 int = v2
                        switch (fi34 + 32767) {
                        case 191:
                            var v41 uint8 = (v5 >> ((255) & 63))
                            _ = v41
                        case 32767:
                            v42 := GenS8{F0: ((!v3) && (121 > -20)), F1: v36, F2: 774}
                            _ = v42
                        }
                        var v43 uint = 2
                        v44 := GenS9{F0: int64(fi35), F1: (644 > 180)}
                        var v45 uint = (v43 * v43)
                        _ = v39
                        _ = v40
                        _ = v43
                        _ = v44
                        _ = v45
                    }
                case 0:
                    var v46 int16 = 0
                    _ = v46
                }
                _ = v38
            } else {
                v2 -= fi34
                var v47 uint64 = 42
                _ = len(v1)
                _ = v47
            }
            v36 = (^(v36))
            _ = fi35
            _ = v36
            _ = v37
        }
        var v48 float32 = 1e10
        _ = fi34
        _ = v48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
