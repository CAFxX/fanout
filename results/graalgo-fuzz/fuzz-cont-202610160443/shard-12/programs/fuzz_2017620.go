package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 uint8; F2 int16 }
type GenS4 struct { F0 uint; F1 uint64; F2 int8 }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64; F1 int64 }
type GenS8 struct { F0 int64; F1 uint32 }
type GenS9 struct { F0 int32; F1 int32 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int64; F1 int }
func genFn0() bool {
    var v4 float32 = 0.1
    v5 := GenS1{F0: 42, F1: 179}
    v4 -= float32(3.14159)
    v4 -= 1.5
    if true {
        if ("\n\t" <= ((("hello" + "hello") + ("a" + ("a" + ((("foo bar" + "a") + "\n\t") + "")))) + ("foo bar" + (("" + "a") + ("hello" + ("" + ("a" + ("foo bar" + (("foo bar" + "a") + (("a" + "") + "foo bar")))))))))) {
            v4 -= ((float32(1.5) + ((3.14159 + (v4 - (0.1 / float32(1e10)))) + v4)) * 0.1)
            for fi6 := 0; fi6 < 3; fi6++ {
                for fi7 := 0; fi7 < 9; fi7++ {
                    if false {
                        v8 := []int{(fi7 % ((fi7) | 1))}
                        v9 := GenS2{F0: 65535}
                        v4 -= -2.25
                        _ = v8
                        _ = v9
                    }
                    _ = fi7
                }
                _ = fi6
            }
        } else {
            v4 -= v4
            var v10 uint32 = 883
            _ = v10
        }
        switch 42 {
        case -1:
            v11 := GenS1{F0: 151, F1: 74}
            _ = v11
        case 397:
            v12 := GenS3{F0: "", F1: 255, F2: 1000}
            _ = v12
        case 446:
            var v13 bool = true
            _ = v13
        }
        if true {
            v14 := GenS3{F0: "foo bar", F1: 0, F2: 0}
            v4 -= 0.1
            _ = v14
        }
    } else {
        for fi15 := 0; fi15 < 3; fi15++ {
            v4 -= ((((float32(0.0) / float32(((1.5 / -2.25) + 1e10))) - ((1e10 + 0.1) / v4)) - 1e10) / 0.1)
            v4 += v4
            _ = fi15
        }
        if (42 > 0) {
            var v16 int64 = 1
            var v17 int64 = (v16 << ((4294967295) & 63))
            _ = v16
            _ = v17
        } else {
            v4 = v4
            var v18 string = ("hello" + "foo bar")
            v19 := GenS4{F0: 548, F1: 2, F2: 0}
            _ = v18
            _ = v19
        }
        fmt.Println(v4)
    }
    _ = v4
    _ = v5
    return true
}

func main() {
    var v1 int16 = 0
    var v2 int = 32768
    v3 := []int{v2, v2, (v2 >> ((18446744073709551615) & 63))}
    var v20 int8 = -128
    for fi21 := 0; fi21 < 6; fi21++ {
        var v22 uint64 = 0
        var v23 float64 = -2.25
        _ = fi21
        _ = v22
        _ = v23
    }
    v2 = (v2 << ((255) & 63))
    v2 = v2
    v1 -= 32767
    switch 32768 {
    case -9223372036854775808:
        var v24 int8 = -128
        _ = v24
    default:
        var v25 int32 = -2147483648
        _ = v25
    }
    switch -165 {
    case -953:
        var v26 float64 = (1e10 + 1.5)
        _ = v26
    }
    var v27 int32 = 0
    var v28 int16 = (^(v1))
    var v29 int32 = v27
    if len(v3) > 0 {
        v3[0] = (v2 - v2)
    }
    fmt.Println(v2)
    v2 += v2
    if false {
        v27 += v27
        for fi30 := 0; fi30 < 10; fi30++ {
            switch -2147483648 {
            case 32768:
                v20 = 100
            case 32767:
                var v31 uint64 = 18446744073709551615
                _ = v31
            }
            switch v2 {
            case 32768:
                if ((v29 + (^(v27))) >= (^(v29))) {
                    for fi32 := 0; fi32 < 9; fi32++ {
                        var v33 int16 = 88
                        var v34 int16 = (v1 & int16(fi32))
                        v34 -= 1000
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                    var v35 uint32 = 0
                    for fi36 := 0; fi36 < 20; fi36++ {
                        v37 := GenS5{F0: ("" + (("" + ("hello" + "hello")) + "foo bar")), F1: "a"}
                        if (678 > 1) {
                            if true {
                                var v38 string = "a"
                                if false {
                                    if len(v3) > 0 {
                                        v3[0] = (fi36 | 1)
                                    }
                                    v39 := GenS2{F0: 255}
                                    _ = v39
                                } else {
                                    v20 = v20
                                }
                                _ = v38
                            } else {
                                v1 += (v28 & (^(v1)))
                                v27 += (v27 ^ (^(v29)))
                                var v40 int32 = v29
                                _ = v40
                            }
                            v27 += (v29 - 0)
                            var v41 uint32 = 42
                            _ = v41
                        }
                        _ = fi36
                        _ = v37
                    }
                    _ = v35
                } else {
                    v1 = (^(v28))
                    fmt.Println(v2, v1)
                }
            default:
                v42 := GenS6{F0: 42}
                _ = v42
            }
            _ = fi30
        }
    } else {
        var v43 uint64 = 878
        v44 := GenS7{F0: (v43 >> ((4294967295) & 63)), F1: 0}
        if (!true) {
            if (-135 == 9223372036854775807) {
                switch v2 {
                case 4294967295:
                    if len(v3) > 0 {
                        v3[0] = v2
                    }
                case 127:
                    var v45 int64 = 42
                    _ = v45
                case -9223372036854775808:
                    fmt.Println(v2, v27)
                }
                if (v28 == (v28 - v1)) {
                    switch ("\n\t" + "hello") {
                    case "a":
                        v1 = -32768
                    }
                }
            } else {
                fmt.Println(v29, v28, v28)
                for fi46 := 0; fi46 < 14; fi46++ {
                    if false {
                        switch "" {
                        case "":
                            v47 := GenS8{F0: 9223372036854775807, F1: 2147483647}
                            _ = v47
                        case "a":
                            var v48 uint32 = 0
                            _ = v48
                        case "foo bar":
                            v2 = (-(fi46))
                        }
                        fmt.Println(v43, v29, v27)
                        fmt.Println(v20, fi46, v1)
                    }
                    _ = fi46
                }
                var v49 uint8 = 2
                _ = v49
            }
            for fi50 := 0; fi50 < 5; fi50++ {
                switch v2 {
                case -2147483648:
                    v27 += v29
                case 0:
                    var v51 float64 = float64(1.5)
                    _ = v51
                case -473:
                    v52 := GenS4{F0: 301, F1: v43, F2: (v20 * 100)}
                    _ = v52
                }
                _ = fi50
            }
            v3 = append(v3, -1)
        } else {
            if len(v3) > 0 {
                v3[0] = int(v43)
            }
            switch v2 {
            case -2147483648:
                var v53 uint8 = 255
                _ = v53
            case 1:
                v54 := GenS3{F0: ("a" + "\n\t"), F1: 1, F2: -553}
                _ = v54
            case 127:
                v55 := GenS9{F0: (v27 << ((769) & 63)), F1: (^(v27))}
                _ = v55
            default:
                var v56 uint64 = (^(v43))
                _ = v56
            }
        }
        _ = v43
        _ = v44
    }
    if false {
        for fi57 := 0; fi57 < 17; fi57++ {
            fmt.Println(v27, fi57)
            v2 += -1
            if (96 != 128) {
                var v58 string = (((((("foo bar" + ("hello" + "hello")) + "foo bar") + ("" + "")) + "foo bar") + "a") + "")
                _ = v58
            }
            _ = fi57
        }
    } else {
        v59 := GenS10{F0: 2}
        if true {
            if true {
                v20 += (v20 * v20)
            }
            v60 := GenS11{F0: 1, F1: v2}
            if false {
                var v61 int8 = (-(v20))
                if (713 >= 18446744073709551615) {
                    for fi62 := 0; fi62 < 16; fi62++ {
                        var v63 float64 = (-2.25 / (1.5 - 1e10))
                        var v64 int8 = (v61 & v61)
                        v20 = v20
                        _ = fi62
                        _ = v63
                        _ = v64
                    }
                }
                v20 = int8(v29)
                _ = v61
            }
            _ = v60
        } else {
            var v65 uint = 65535
            _ = len(v3)
            _ = v65
        }
        v29 -= v27
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
