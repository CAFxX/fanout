package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 int; F1 int; F2 uint16 }
type GenS7 struct { F0 int; F1 int64 }
type GenS8 struct { F0 int64; F1 uint8 }
func genFn0() int32 {
    v4 := []int{4294967295, 42, 7}
    for fi5 := 0; fi5 < 2; fi5++ {
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 float64 = 1e10
            _ = fi6
            _ = v7
        }
        var v8 bool = true
        _ = fi5
        _ = v8
    }
    _ = v4
    return -1
}

func main() {
    var v1 uint8 = 0
    var v2 int16 = 1
    v3 := make(map[int]int)
    if false {
        var v9 bool = false
        v2 += v2
        _ = v9
    }
    switch (("foo bar" + "") + ((("" + "foo bar") + "") + "foo bar")) {
    case "hello":
        v10 := GenS1{F0: (v2 << ((65535) & 63))}
        _ = v10
    case "foo bar":
        if true {
            var v11 int32 = -1
            switch 32 {
            case -857:
                v11 -= 0
            }
            _ = v11
        } else {
            var v12 uint16 = 65535
            _ = v12
        }
    default:
        v1 = v1
    }
    var v13 float64 = 0.1
    var v14 bool = (18446744073709551615 != 18446744073709551615)
    var v15 bool = false
    v16 := GenS2{F0: 1, F1: (true || v15)}
    v1 += 42
    fmt.Println(v13, v15, v2)
    fmt.Println(v1)
    switch -443 {
    case 32768:
        if (0 < 42) {
            switch "" {
            case "a":
                for fi17 := 0; fi17 < 15; fi17++ {
                    v18 := GenS2{F0: 1, F1: v15}
                    var v19 int = fi17
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            default:
                v20 := GenS3{F0: 1, F1: 42}
                _ = v20
            }
            var v21 int = 32768
            v14 = (v15 && false)
            _ = v21
        }
    case -9223372036854775808:
        switch int(v1) {
        case -9223372036854775808:
            var v22 int = 96
            _ = v22
        default:
            fmt.Println(v2, v2)
        }
    default:
        v13 -= 0.0
    }
    if v15 {
        v23 := GenS4{F0: 294}
        _ = v23
    } else {
        for fi24 := 0; fi24 < 2; fi24++ {
            v25 := GenS2{F0: uint16(v1), F1: v15}
            for fi26 := 0; fi26 < 9; fi26++ {
                fmt.Println(fi26)
                _ = fi26
            }
            _ = fi24
            _ = v25
        }
        for fi27 := 0; fi27 < 18; fi27++ {
            switch "" {
            case "foo bar":
                v15 = (v15 || false)
            }
            fmt.Println(v15, fi27, fi27)
            _ = fi27
        }
        var v28 uint8 = 127
        _ = v28
    }
    fmt.Println(v14, v14)
    if (!(uint16(v1) != 30)) {
        for fi29 := 0; fi29 < 10; fi29++ {
            var v30 uint = 255
            v31 := GenS5{F0: 42, F1: (-(v2))}
            switch "a" {
            case "foo bar":
                v32 := GenS6{F0: (fi29 + (fi29 << (((v30 & (v30 * v30))) & 63))), F1: fi29, F2: 308}
                _ = v32
            case "hello":
                v33 := GenS6{F0: fi29, F1: 9223372036854775807, F2: 32767}
                _ = v33
            case "\n\t":
                if v15 {
                    var v34 int8 = 100
                    _ = v34
                } else {
                    var v35 float32 = 3.14159
                    for fi36 := 0; fi36 < 12; fi36++ {
                        switch (fi29 / (((^(fi36))) | 1)) {
                        case -2147483648:
                            fmt.Println(v14, v30)
                        case -9223372036854775808:
                            v2 = 32767
                        case 42:
                            v30 = v30
                        }
                        if (1 >= 225) {
                            if ((3.14159 * v35) == v35) {
                                fmt.Println(v30)
                            } else {
                                delete(v3, -105)
                            }
                        }
                        v1 -= 0
                        if (("a" >= "\n\t") && v15) {
                            switch "a" {
                            case "":
                                switch 4294967295 {
                                case 4294967295:
                                    v3[fi29] = int(v30)
                                case -529:
                                    v15 = v14
                                case 42:
                                    switch ("a" + ("hello" + ("foo bar" + "\n\t"))) {
                                    case "a":
                                        var v37 int16 = 0
                                        _ = v37
                                    case "":
                                        v3[971] = fi36
                                    default:
                                        switch "a" {
                                        case "hello":
                                            if v15 {
                                                var v38 string = (("" + "hello") + "hello")
                                                _ = v38
                                            }
                                        }
                                    }
                                default:
                                    v1 = (v1 << ((v30) & 63))
                                }
                            case "foo bar":
                                switch int(v1) {
                                case 486:
                                    if v14 {
                                        v39 := []int{(fi29 / ((128) | 1)), (fi29 ^ int(v30))}
                                        _ = v39
                                    }
                                default:
                                    v2 = (v2 * 1000)
                                }
                            }
                            v13 += v13
                        } else {
                            v40 := GenS4{F0: 2147483648}
                            _ = v40
                        }
                        _ = fi36
                    }
                    var v41 int64 = 679
                    _ = v35
                    _ = v41
                }
            default:
                if true {
                    fmt.Println(v13, v1, v15)
                } else {
                    var v42 int16 = -978
                    _ = v42
                }
            }
            _ = fi29
            _ = v30
            _ = v31
        }
        var v43 int8 = -42
        _ = v43
    } else {
        v44 := []int{42}
        var v45 int64 = 1
        if (("" <= "") || ((genFn0() < -772) && false)) {
            v46 := GenS7{F0: 32767, F1: v45}
            switch -932 {
            case 42:
                v47 := GenS8{F0: -9223372036854775808, F1: 42}
                _ = v47
            default:
                delete(v3, 32768)
            }
            _ = v46
        }
        _ = v44
        _ = v45
    }
    v3[42] = 42
    _ = len(v3)
    for fi48 := 0; fi48 < 15; fi48++ {
        switch "\n\t" {
        case "\n\t":
            for fi49 := 0; fi49 < 20; fi49++ {
                fmt.Println(v1, fi49)
                v50 := GenS1{F0: (v2 | 42)}
                _ = fi49
                _ = v50
            }
        case "hello":
            var v51 float64 = v13
            _ = v51
        case "foo bar":
            _ = len(v3)
        default:
            switch (fi48 % (((fi48 % ((9223372036854775807) | 1))) | 1)) {
            case 413:
                var v52 int16 = -730
                _ = v52
            case -9223372036854775808:
                v13 = v13
            default:
                var v53 bool = v14
                _ = v53
            }
        }
        _ = fi48
    }
    for fi54 := 0; fi54 < 2; fi54++ {
        v14 = true
        _ = fi54
    }
    fmt.Println(v1, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
