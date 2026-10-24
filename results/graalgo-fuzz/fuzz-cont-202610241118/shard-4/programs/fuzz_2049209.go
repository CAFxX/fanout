package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int }
type GenS2 struct { F0 int; F1 int32; F2 int }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 int64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 65535
    var v3 int8 = 35
    v4 := GenS1{F0: (v3 + 127), F1: false, F2: 1}
    if false {
        fmt.Println(v2)
    } else {
        v3 -= v3
    }
    switch 127 {
    case -1:
        v5 := []int{-365, -2147483648}
        _ = v5
    case -728:
        v6 := GenS2{F0: 4294967295, F1: -962, F2: -1}
        _ = v6
    case 127:
        var v7 float32 = -2.25
        _ = v7
    default:
        var v8 int8 = -1
        _ = v8
    }
    if (2 >= 2) {
        fmt.Println(v3, v3)
        for fi9 := 0; fi9 < 7; fi9++ {
            v3 += (v3 + v3)
            _ = fi9
        }
        v2 -= (v2 + (v2 % (((v2 & 4294967295)) | 1)))
    } else {
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 string = ("foo bar" + ("" + "\n\t"))
            for fi12 := 0; fi12 < 19; fi12++ {
                v3 -= v3
                _ = fi12
            }
            _ = fi10
            _ = v11
        }
        if (3.14159 < (-2.25 + 1e10)) {
            v2 = v2
        }
        v13 := GenS2{F0: 4294967295, F1: 2147483647, F2: 2147483647}
        _ = v13
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        if false {
            var v15 uint8 = 1
            _ = v15
        } else {
            v16 := GenS3{F0: false, F1: 464}
            _ = v16
        }
        if true {
            if true {
                var v17 int64 = 9223372036854775807
                var v18 float64 = 1.5
                var v19 int16 = 1
                _ = v17
                _ = v18
                _ = v19
            } else {
                v20 := []int{}
                _ = v20
            }
            var v21 uint64 = 0
            var v22 bool = ((v3 + -128) != (v3 ^ -1))
            _ = v21
            _ = v22
        } else {
            var v23 int64 = -1
            var v24 int32 = 1
            if false {
                for fi25 := 0; fi25 < 2; fi25++ {
                    var v26 int16 = 1
                    switch (-(fi14)) {
                    case 134:
                        if (!true) {
                            v3 += v3
                        }
                    case 159:
                        if false {
                            v26 = 848
                            if false {
                                var v27 float32 = 3.14159
                                v3 += (^(v3))
                                _ = v27
                            } else {
                                v28 := GenS4{F0: (-(v26)), F1: 65535}
                                if (0 == 128) {
                                    fmt.Println(v26)
                                    v24 = -1
                                    v29 := []int{1}
                                    _ = v29
                                } else {
                                    if false {
                                        v1[(-(fi14))] = v2
                                    } else {
                                        v30 := v1[fi14]
                                        var v31 int = fi25
                                        v31 = v31
                                        if (!true) {
                                            fmt.Println(v31, v30, v26)
                                            v24 = (v24 * (^(v24)))
                                        } else {
                                            switch fi25 {
                                            case 352:
                                                var v32 uint16 = 2
                                                _ = v32
                                            case -119:
                                                fmt.Println(v30, v31)
                                            default:
                                                var v33 int64 = (v23 % ((-1) | 1))
                                                _ = v33
                                            }
                                            var v34 int = -609
                                            if (v2 >= 791) {
                                                v35 := GenS2{F0: v34, F1: v24, F2: 127}
                                                _ = v35
                                            } else {
                                                if true {
                                                    v24 -= 1
                                                } else {
                                                    v36 := GenS5{F0: (v23 ^ int64(v30))}
                                                    _ = v36
                                                }
                                            }
                                            _ = v34
                                        }
                                        _ = v30
                                        _ = v31
                                    }
                                    var v37 uint8 = 255
                                    _ = v37
                                }
                                var v38 int32 = (^(v24))
                                _ = v28
                                _ = v38
                            }
                        }
                    }
                    _ = fi25
                    _ = v26
                }
                for fi39 := 0; fi39 < 11; fi39++ {
                    v40 := v1[fi39]
                    v41 := GenS1{F0: int8(fi39), F1: false, F2: fi39}
                    _ = fi39
                    _ = v40
                    _ = v41
                }
            }
            _ = v23
            _ = v24
        }
        _ = fi14
    }
    if (false && false) {
        for fi42 := 0; fi42 < 1; fi42++ {
            var v43 uint16 = 24
            var v44 int = (fi42 % ((-9223372036854775808) | 1))
            if false {
                v1[-9223372036854775808] = v2
                _ = len(v1)
                var v45 uint32 = v2
                _ = v45
            } else {
                for fi46 := 0; fi46 < 7; fi46++ {
                    var v47 int = v44
                    _ = fi46
                    _ = v47
                }
                v48 := GenS3{F0: (42 <= 142), F1: 32767}
                _ = v48
            }
            _ = fi42
            _ = v43
            _ = v44
        }
        fmt.Println(v3)
    }
    fmt.Println(v3, v3)
    v49 := GenS2{F0: int(v2), F1: 2147483647, F2: -1}
    v50 := GenS4{F0: 32767, F1: 255}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v50
}
