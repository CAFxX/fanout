package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int64; F2 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint64; F1 int; F2 uint32 }
func genFn0(p0 uint16) uint {
    var v7 int = -139
    v8 := GenS1{F0: (^(p0))}
    if true {
        if (true && ("a" >= ((((("hello" + "") + "foo bar") + "") + "") + ("a" + "a")))) {
            if (316 != 2) {
                if (false || false) {
                    v9 := GenS2{F0: 42, F1: int64(p0), F2: "a"}
                    _ = v9
                } else {
                    var v10 uint16 = (p0 & p0)
                    _ = v10
                }
            } else {
                if (32767 == 769) {
                    p0 += (p0 + (p0 << ((249) & 63)))
                    fmt.Println(p0, p0)
                    v7 = v7
                }
                p0 -= (^(p0))
                v11 := GenS3{F0: 1}
                _ = v11
            }
            for fi12 := 0; fi12 < 10; fi12++ {
                if (1.5 != -2.25) {
                    if true {
                        var v13 uint = 0
                        _ = v13
                    } else {
                        v7 = (fi12 ^ (v7 + -2147483648))
                        p0 = 134
                    }
                }
                var v14 uint16 = 0
                _ = fi12
                _ = v14
            }
            var v15 int64 = -425
            _ = v15
        } else {
            fmt.Println(v7, v7)
            if (363 <= -789) {
                v16 := GenS1{F0: 1}
                var v17 string = "a"
                if true {
                    v18 := GenS1{F0: 836}
                    var v19 int = v7
                    _ = v18
                    _ = v19
                }
                _ = v16
                _ = v17
            } else {
                v7 -= v7
            }
            fmt.Println(v7, p0)
        }
        fmt.Println(v7, v7)
    }
    var v20 uint64 = 152
    fmt.Println(v20)
    _ = p0
    _ = v20
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 float64 = float64(float32(3.14159))
    var v2 string = ("foo bar" + ("\n\t" + "a"))
    var v3 int64 = 0
    v4 := make(map[string]string)
    var v5 int32 = -1
    v6 := make(map[int]bool)
    for fi21 := 0; fi21 < 14; fi21++ {
        switch (^(fi21)) {
        case -9223372036854775808:
            v22 := v4[v2]
            _ = v22
        case 0:
            if false {
                var v23 bool = false
                _ = v23
            } else {
                fmt.Println(v2, fi21)
            }
        case -781:
            v5 = 0
        default:
            for fi24 := 0; fi24 < 1; fi24++ {
                if false {
                    v25 := GenS4{F0: 46}
                    var v26 uint32 = 65535
                    _ = v25
                    _ = v26
                } else {
                    var v27 bool = true
                    fmt.Println(fi21, v5)
                    _ = v27
                }
                var v28 uint64 = 4294967295
                _ = fi24
                _ = v28
            }
        }
        if ((-2.25 - float32(0.0)) >= -2.25) {
            var v29 uint = 680
            _ = v29
        } else {
            v30 := []int{-2147483648}
            _ = v30
        }
        _ = fi21
    }
    v1 = v1
    for fi31 := 0; fi31 < 10; fi31++ {
        if (false && (1.5 == 1.5)) {
            v1 = v1
            v3 = 42
        }
        _ = fi31
    }
    fmt.Println(v5)
    if (!true) {
        var v32 bool = (0.0 == 1e10)
        _ = v32
    } else {
        v33 := []int{129, 42, -992, -1}
        _ = v33
    }
    if true {
        v3 = v3
        var v34 int = 260
        for fi35 := 0; fi35 < 10; fi35++ {
            fmt.Println(v34, v3)
            v36 := GenS5{F0: 32767, F1: ("\n\t" + "foo bar")}
            _ = fi35
            _ = v36
        }
        _ = v34
    } else {
        for fi37 := 0; fi37 < 18; fi37++ {
            for fi38 := 0; fi38 < 8; fi38++ {
                delete(v6, (-(fi38)))
                v3 -= (v3 << ((384) & 63))
                _ = fi38
            }
            var v39 int16 = -376
            for fi40 := 0; fi40 < 5; fi40++ {
                v41 := GenS1{F0: 471}
                fmt.Println(v39, v2)
                v42 := []int{-9223372036854775808, (fi37 << ((65535) & 63)), fi37, (fi37 / ((fi40) | 1))}
                _ = fi40
                _ = v41
                _ = v42
            }
            _ = fi37
            _ = v39
        }
        var v43 int = 732
        _ = v43
    }
    for fi44 := 0; fi44 < 18; fi44++ {
        var v45 int64 = v3
        v1 -= v1
        _ = fi44
        _ = v45
    }
    v46 := GenS3{F0: v5}
    v3 -= 0
    if true {
        v47 := GenS2{F0: 3, F1: int64(v5), F2: v2}
        _ = v47
    }
    if ((v2 + "a") > "") {
        var v48 uint16 = 836
        v6[-989] = true
        v49 := GenS6{F0: 0}
        _ = v48
        _ = v49
    } else {
        v5 -= (v5 * v5)
        var v50 float64 = (v1 * float64(1.5))
        switch v2 {
        case "\n\t":
            v2 = v2
        }
        _ = v50
    }
    for fi51 := 0; fi51 < 2; fi51++ {
        var v52 uint64 = 1
        _ = fi51
        _ = v52
    }
    for fi53 := 0; fi53 < 10; fi53++ {
        if true {
            var v54 int16 = -918
            _ = v54
        } else {
            var v55 int8 = 0
            _ = v55
        }
        _ = fi53
    }
    for fi56 := 0; fi56 < 18; fi56++ {
        v1 = ((v1 + v1) + (-2.25 * float64(1.5)))
        v3 = (v3 * -9223372036854775808)
        _ = fi56
    }
    v57 := GenS7{F0: true}
    for fi58 := 0; fi58 < 13; fi58++ {
        v3 += v3
        for fi59 := 0; fi59 < 16; fi59++ {
            switch fi58 {
            case -9223372036854775808:
                var v60 float64 = 1.5
                v60 -= float64((-2.25 / (0.1 / 1.5)))
                if (42 >= 1) {
                    if true {
                        v61 := GenS7{F0: false}
                        delete(v6, (fi58 << ((255) & 63)))
                        var v62 uint = 255
                        var v63 int = -2147483648
                        var v64 uint8 = 1
                        v65 := GenS8{F0: (v62 & (v62 / ((v62) | 1)))}
                        _ = v61
                        _ = v62
                        _ = v63
                        _ = v64
                        _ = v65
                    } else {
                        v3 += (v3 * v3)
                        var v66 bool = true
                        if ((1 != -408) && true) {
                            if v66 {
                                var v67 int = fi58
                                v5 += 372
                                fmt.Println(v60, v5)
                                var v68 string = (v2 + "")
                                var v69 uint64 = 0
                                var v70 uint16 = 0
                                _ = v67
                                _ = v68
                                _ = v69
                                _ = v70
                            }
                        }
                        _ = v66
                    }
                    var v71 int8 = -38
                    var v72 float32 = (3.14159 + 3.14159)
                    _ = v71
                    _ = v72
                }
                _ = v60
            case -533:
                if (v2 > v2) {
                    v73 := GenS6{F0: 115}
                    _ = v73
                }
            }
            _ = fi59
        }
        v5 += (-(v5))
        _ = fi58
    }
    v74 := GenS9{F0: 353, F1: -9223372036854775808, F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v57
    _ = v6
    _ = v74
}
