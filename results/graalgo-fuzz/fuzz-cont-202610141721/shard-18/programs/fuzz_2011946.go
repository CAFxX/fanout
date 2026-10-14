package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int32 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 uint64; F2 bool }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 uint64; F1 uint8; F2 uint64 }
func genFn0(p0 uint16) bool {
    p0 = (p0 | p0)
    var v4 int16 = -949
    v5 := GenS1{F0: 1, F1: 738, F2: 42}
    fmt.Println(p0)
    if ((-(v4)) >= (v4 + (v4 + (v4 & -153)))) {
        if false {
            if (!false) {
                v4 = v4
            } else {
                switch ("foo bar" + "a") {
                case "hello":
                    v6 := GenS1{F0: -9223372036854775808, F1: 42, F2: -2147483648}
                    _ = v6
                case "foo bar":
                    v4 = (v4 >> ((4294967295) & 63))
                }
            }
        } else {
            if (-1 > 9223372036854775807) {
                var v7 int16 = v4
                _ = v7
            }
            var v8 uint32 = 2147483647
            _ = v8
        }
        var v9 uint64 = 4294967295
        _ = v9
    } else {
        v4 = 987
        var v10 bool = false
        _ = v10
    }
    _ = p0
    _ = v4
    _ = v5
    return true
}

func main() {
    var v1 bool = false
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    v1 = v1
    v3[-2147483648] = 1
    for fi11 := 0; fi11 < 17; fi11++ {
        fmt.Println(fi11, v1)
        if ((100 != 42) || v1) {
            v1 = v1
            var v12 uint8 = 1
            _ = v12
        }
        _ = fi11
    }
    v13 := v2[-847]
    var v14 float32 = 1.5
    var v15 uint8 = 127
    for fi16 := 0; fi16 < 15; fi16++ {
        v14 = v14
        for fi17 := 0; fi17 < 2; fi17++ {
            fmt.Println(fi17)
            var v18 float32 = v14
            switch fi17 {
            case 0:
                v19 := GenS2{F0: 817, F1: v1}
                _ = v19
            default:
                v20 := GenS3{F0: fi16}
                _ = v20
            }
            v1 = true
            switch fi17 {
            case -9223372036854775808:
                v21 := GenS3{F0: (fi17 << ((18446744073709551615) & 63))}
                _ = v21
            default:
                if v1 {
                    v22 := GenS1{F0: -304, F1: 783, F2: -575}
                    switch "" {
                    case "hello":
                        fmt.Println(v15, fi17, fi16)
                        switch "\n\t" {
                        case "hello":
                            var v23 uint32 = v13
                            _ = v23
                        case "\n\t":
                            v2[(fi16 & fi17)] = v13
                        case "a":
                            var v24 int = (fi17 + (fi17 << ((255) & 63)))
                            _ = v24
                        default:
                            switch (fi16 ^ -36) {
                            case 0:
                                if v1 {
                                    var v25 uint64 = 573
                                    var v26 uint8 = (v15 >> ((599) & 63))
                                    _ = v25
                                    _ = v26
                                }
                            case -1:
                                var v27 int64 = 835
                                _ = v27
                            case -620:
                                var v28 float32 = 0.1
                                _ = v28
                            }
                        }
                        v29 := []int{fi16, (fi16 + 0), -859, -18}
                        _ = v29
                    case "":
                        v30 := []int{fi17}
                        _ = v30
                    default:
                        fmt.Println(v1, fi17, v18)
                    }
                    var v31 uint = 2
                    _ = v22
                    _ = v31
                }
            }
            _ = fi17
            _ = v18
        }
        if (v1 || v1) {
            if v1 {
                switch fi16 {
                case -9223372036854775808:
                    if false {
                        var v32 uint = 256
                        var v33 string = "foo bar"
                        _ = v32
                        _ = v33
                    }
                case -465:
                    fmt.Println(v14, v13, v15)
                default:
                    for fi34 := 0; fi34 < 7; fi34++ {
                        if v1 {
                            fmt.Println(v1, v13, v15)
                            v14 += 3.14159
                            v35 := GenS4{F0: (v15 + v15), F1: uint64(v15), F2: v1}
                            _ = v35
                        }
                        v13 = (v13 | 393)
                        _ = len(v2)
                        _ = fi34
                    }
                }
            }
            switch (fi16 ^ fi16) {
            case -979:
                switch (^(fi16)) {
                case 660:
                    if (false && v1) {
                        var v36 int16 = -1
                        var v37 bool = (!((v1 && ((-773 <= 9223372036854775807) || (0 > 1))) && (420 <= 749)))
                        _ = v36
                        _ = v37
                    } else {
                        for fi38 := 0; fi38 < 18; fi38++ {
                            fmt.Println(v14)
                            _ = fi38
                        }
                        for fi39 := 0; fi39 < 2; fi39++ {
                            var v40 uint32 = v13
                            v40 = v40
                            _ = fi39
                            _ = v40
                        }
                    }
                case 1:
                    v13 -= v13
                default:
                    if false {
                        v1 = v1
                    }
                }
            default:
                if false {
                    for fi41 := 0; fi41 < 10; fi41++ {
                        v13 = v13
                        v15 += v15
                        _ = fi41
                    }
                    v13 += (^(v13))
                }
            }
            v15 = v15
        }
        _ = fi16
    }
    var v42 float64 = 1.5
    fmt.Println(v15)
    v43 := GenS1{F0: -917, F1: 2, F2: -609}
    var v44 uint32 = (v13 ^ 4294967295)
    for fi45 := 0; fi45 < 15; fi45++ {
        if v1 {
            v46 := GenS5{F0: 731, F1: "hello"}
            for fi47 := 0; fi47 < 18; fi47++ {
                switch "\n\t" {
                case "hello":
                    var v48 int64 = 0
                    _ = v48
                case "a":
                    var v49 int8 = -70
                    _ = v49
                default:
                    v50 := GenS5{F0: v13, F1: "\n\t"}
                    _ = v50
                }
                var v51 float32 = v14
                switch (fi47 << ((18446744073709551615) & 63)) {
                case 875:
                    fmt.Println(v14, fi45, v42)
                default:
                    if true {
                        v52 := []int{fi47, 42, (-(fi47)), int(v15)}
                        _ = v52
                    } else {
                        v53 := GenS6{F0: 311, F1: (^(v15)), F2: 289}
                        _ = v53
                    }
                }
                _ = fi47
                _ = v51
            }
            var v54 int32 = 1
            _ = v46
            _ = v54
        }
        _ = fi45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v42
    _ = v43
    _ = v44
}
