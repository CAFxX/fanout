package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int; F2 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS6 struct { F0 int64; F1 uint }

func main() {
    var v1 int32 = 720
    var v2 uint16 = 65535
    var v3 string = "hello"
    var v4 uint8 = 113
    var v5 string = (v3 + "foo bar")
    for fi6 := 0; fi6 < 2; fi6++ {
        fmt.Println(v5, v5)
        v1 = (v1 * -665)
        if false {
            for fi7 := 0; fi7 < 2; fi7++ {
                v5 = v5
                var v8 uint64 = 1
                _ = fi7
                _ = v8
            }
            v9 := GenS1{F0: (127 > (^(v4))), F1: 42}
            _ = v9
        }
        _ = fi6
    }
    fmt.Println(v5, v4)
    var v10 int = 32767
    switch v3 {
    case "\n\t":
        if (!true) {
            var v11 float32 = 1e10
            _ = v11
        } else {
            var v12 uint8 = v4
            if (204 < 4294967295) {
                var v13 int64 = 42
                _ = v13
            } else {
                if (-1 == 100) {
                    for fi14 := 0; fi14 < 6; fi14++ {
                        if (!(4294967295 <= 1)) {
                            v15 := []int{fi14, -550, fi14}
                            v16 := GenS1{F0: (false || (2 < 1)), F1: 402}
                            _ = v15
                            _ = v16
                        } else {
                            v5 = v3
                        }
                        v12 -= (v12 ^ (v12 >> ((255) & 63)))
                        _ = fi14
                    }
                    var v17 int16 = 1
                    _ = v17
                } else {
                    fmt.Println(v12)
                    if true {
                        v10 -= v10
                        v18 := GenS2{F0: (v10 | (v10 >> ((804) & 63)))}
                        if (true || (int16(v4) <= 0)) {
                            if false {
                                fmt.Println(v3)
                                v19 := []int{}
                                for fi20 := 0; fi20 < 17; fi20++ {
                                    _ = len(v19)
                                    var v21 int = (fi20 | (v10 ^ fi20))
                                    v22 := GenS3{F0: uint64(v2), F1: int(v2), F2: 769}
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                                _ = v19
                            } else {
                                v23 := GenS4{F0: 279}
                                fmt.Println(v4, v2)
                                v24 := GenS5{F0: 4294967295, F1: (v2 / ((65535) | 1)), F2: 1}
                                _ = v23
                                _ = v24
                            }
                            v25 := GenS5{F0: 1, F1: (v2 << ((0) & 63)), F2: 558}
                            switch (v10 + v10) {
                            case 0:
                                for fi26 := 0; fi26 < 12; fi26++ {
                                    for fi27 := 0; fi27 < 19; fi27++ {
                                        var v28 uint32 = 42
                                        _ = fi27
                                        _ = v28
                                    }
                                    v29 := []int{(fi26 & (fi26 / (((fi26 - (v10 - int(v4)))) | 1))), 128, (v10 + v10)}
                                    _ = fi26
                                    _ = v29
                                }
                            case 9223372036854775807:
                                v3 = (v5 + "hello")
                            }
                            _ = v25
                        }
                        _ = v18
                    } else {
                        v30 := []int{(v10 & v10), (v10 % (((^(v10))) | 1))}
                        var v31 float64 = 0.1
                        _ = v30
                        _ = v31
                    }
                }
            }
            switch v10 {
            case 574:
                if (v5 >= (v5 + "hello")) {
                    v32 := []int{1, v10, (v10 - int(v4)), v10}
                    if true {
                        if len(v32) >= 2 {
                            v32 = v32[1:len(v32)]
                        }
                        var v33 uint64 = 2
                        switch 42 {
                        case 128:
                            v34 := GenS1{F0: false, F1: (v33 * (v33 >> ((2) & 63)))}
                            _ = v34
                        case 2147483647:
                            v2 += v2
                        case 1:
                            v3 = ("a" + v3)
                        }
                        _ = v33
                    }
                    v4 -= v12
                    _ = v32
                }
            case 2147483647:
                fmt.Println(v3)
            case 42:
                for fi35 := 0; fi35 < 18; fi35++ {
                    v5 = (v3 + v3)
                    _ = fi35
                }
            }
            _ = v12
        }
    default:
        v36 := GenS2{F0: v10}
        _ = v36
    }
    if (0.1 != -2.25) {
        var v37 float32 = -2.25
        var v38 float64 = 0.0
        _ = v37
        _ = v38
    } else {
        for fi39 := 0; fi39 < 7; fi39++ {
            var v40 float32 = 1e10
            _ = fi39
            _ = v40
        }
        var v41 int16 = 42
        v42 := GenS3{F0: 1, F1: 441, F2: 1}
        _ = v41
        _ = v42
    }
    var v43 int = v10
    fmt.Println(v10, v3)
    if (-128 == 1) {
        var v44 int64 = -9223372036854775808
        _ = v44
    }
    fmt.Println(v1, v1, v3)
    if true {
        v45 := GenS6{F0: 9223372036854775807, F1: 4294967295}
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
