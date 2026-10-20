package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 uint; F2 bool }
type GenS4 struct { F0 int; F1 int; F2 int16 }
type GenS5 struct { F0 uint; F1 string; F2 int16 }
type GenS6 struct { F0 uint8; F1 int16 }
type GenS7 struct { F0 bool; F1 uint16 }
type GenS8 struct { F0 int; F1 int; F2 uint }
type GenS9 struct { F0 uint32 }
func genFn0(p0 bool) uint64 {
    v5 := GenS1{F0: -1}
    if p0 {
        p0 = (!(!false))
        var v6 float64 = 0.1
        _ = v6
    } else {
        p0 = (p0 && p0)
        switch -1 {
        case -1:
            var v7 uint32 = 2147483648
            _ = v7
        case 767:
            var v8 int64 = 42
            _ = v8
        default:
            p0 = p0
        }
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 int32 = 0
            for fi11 := 0; fi11 < 15; fi11++ {
                v10 = 258
                v12 := []int{(fi11 << ((86) & 63)), 2147483647, (fi9 | fi11), 2147483647}
                v10 += 177
                v13 := GenS1{F0: -9223372036854775808}
                fmt.Println(v10, p0)
                var v14 int32 = v10
                v15 := GenS2{F0: 65535}
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
                _ = v15
            }
            var v16 uint16 = 1
            _ = fi9
            _ = v10
            _ = v16
        }
    }
    v17 := []int{}
    _ = p0
    _ = v17
    _ = v5
    return 42
}
func genFn1(p0 int16) bool {
    v18 := []int{2147483647, 127, -600, 0}
    fmt.Println(p0)
    p0 -= p0
    p0 += -362
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    _ = p0
    _ = v18
    return (-894 != 1)
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{}
    v3 := make(map[string]uint32)
    var v4 int16 = 0
    for fi19 := 0; fi19 < 1; fi19++ {
        if true {
            fmt.Println(v4, v4)
            if (759 == 626) {
                if false {
                    var v20 string = "foo bar"
                    _ = v20
                }
                var v21 uint32 = 4294967295
                var v22 float64 = 0.1
                _ = v21
                _ = v22
            } else {
                for fi23 := 0; fi23 < 16; fi23++ {
                    if false {
                        switch "a" {
                        case "hello":
                            if (159 != 1) {
                                var v24 int32 = -1
                                v25 := GenS2{F0: 42}
                                if (v4 >= (v4 / ((v4) | 1))) {
                                    v26 := GenS1{F0: 857}
                                    fmt.Println(v24)
                                    v27 := GenS3{F0: v24, F1: 478, F2: false}
                                    _ = v26
                                    _ = v27
                                } else {
                                    fmt.Println(fi19, v4, fi23)
                                    fmt.Println(v24, fi23)
                                }
                                v24 = v24
                                v28 := GenS3{F0: 1, F1: 42, F2: true}
                                switch fi19 {
                                case -9223372036854775808:
                                    v29 := GenS4{F0: (fi19 * fi19), F1: -1, F2: v4}
                                    _ = v29
                                default:
                                    var v30 uint32 = 0
                                    _ = v30
                                }
                                var v31 int8 = -1
                                if (!(!true)) {
                                    v4 -= -1
                                    v2 = append(v2, "a")
                                    if true {
                                        v4 -= v4
                                        fmt.Println(fi23, v4)
                                        var v32 bool = true
                                        _ = v32
                                    } else {
                                        v24 += 1
                                        v33 := GenS5{F0: 1, F1: "a", F2: v4}
                                        _ = v33
                                    }
                                }
                                if true {
                                    var v34 float64 = (0.0 * 0.1)
                                    _ = v34
                                } else {
                                    v35 := GenS2{F0: 2147483647}
                                    _ = v35
                                }
                                _ = v24
                                _ = v25
                                _ = v28
                                _ = v31
                            } else {
                                v36 := GenS6{F0: 55, F1: 32767}
                                v4 -= v4
                                _ = len(v2)
                                switch (-(fi23)) {
                                case -717:
                                    v4 = v4
                                    if true {
                                        v4 = (v4 | v4)
                                        v4 -= 1
                                        switch (^(fi19)) {
                                        case 9223372036854775807:
                                            if true {
                                                v37 := GenS7{F0: true, F1: 300}
                                                if (!false) {
                                                    v38 := GenS3{F0: int32(v4), F1: 255, F2: false}
                                                    _ = v38
                                                }
                                                _ = v37
                                            }
                                        case 0:
                                            if true {
                                                var v39 int32 = 42
                                                v39 -= v39
                                                _ = v39
                                            }
                                        case 128:
                                            v4 += v4
                                        default:
                                            v40 := v3[("\n\t" + "hello")]
                                            _ = v40
                                        }
                                    }
                                default:
                                    v4 = 1
                                }
                                _ = v36
                            }
                        default:
                            v4 -= (^(v4))
                        }
                        var v41 float64 = ((-2.25 + (1.5 / 0.1)) - (0.1 / 3.14159))
                        fmt.Println(fi19, v4)
                        _ = v41
                    } else {
                        v3[("foo bar" + "\n\t")] = 719
                        if len(v2) > 0 {
                            v2[0] = "a"
                        }
                        fmt.Println(v4)
                    }
                    v2 = append(v2, "\n\t")
                    _ = fi23
                }
            }
            var v42 float64 = 1e10
            _ = v42
        } else {
            var v43 uint = 255
            _ = v43
        }
        var v44 int = fi19
        var v45 int = (v44 + (fi19 * (v44 % ((-97) | 1))))
        _ = fi19
        _ = v44
        _ = v45
    }
    switch "\n\t" {
    case "hello":
        v4 += (v4 - v4)
    case "\n\t":
        for fi46 := 0; fi46 < 10; fi46++ {
            v4 -= 0
            _ = fi46
        }
    case "a":
        for fi47 := 0; fi47 < 14; fi47++ {
            switch "hello" {
            case "\n\t":
                var v48 int64 = 1
                _ = v48
            case "hello":
                v49 := GenS8{F0: (fi47 & (fi47 & (fi47 >> ((65535) & 63)))), F1: (^(fi47)), F2: 912}
                _ = v49
            }
            v4 += (v4 + v4)
            if ((!false) || (!(2147483648 != 2147483648))) {
                if true {
                    v4 -= (v4 % (((v4 << ((18446744073709551615) & 63))) | 1))
                    var v50 int32 = -1
                    v51 := GenS9{F0: 42}
                    _ = v50
                    _ = v51
                }
                var v52 bool = true
                v53 := GenS1{F0: 1}
                _ = v52
                _ = v53
            }
            _ = fi47
        }
    default:
        fmt.Println(v4)
    }
    var v54 float64 = (3.14159 * ((((1.5 - float64((0.0 / 1.0))) * 3.14159) + 1.5) - ((-2.25 + ((1e10 + 0.1) * (0.1 / 0.1))) - (-2.25 + 3.14159))))
    var v55 uint16 = 209
    var v56 float32 = 3.14159
    var v57 int = 32767
    v57 += v57
    fmt.Println(v54)
    for fi58 := 0; fi58 < 4; fi58++ {
        v59 := GenS1{F0: -1}
        fmt.Println(v4)
        _ = fi58
        _ = v59
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v55, v4, v57)
    if (1.5 != 1.5) {
        v56 = v56
    }
    var v60 uint32 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v60
}
