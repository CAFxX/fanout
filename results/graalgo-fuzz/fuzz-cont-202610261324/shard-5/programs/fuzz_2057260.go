package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 string; F2 int8 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 bool; F2 uint }
type GenS6 struct { F0 int8; F1 int }
type GenS7 struct { F0 uint64; F1 uint8 }
type GenS8 struct { F0 int64; F1 uint64 }
func genFn0() bool {
    if (253 > 1) {
        v6 := GenS1{F0: 2}
        _ = v6
    } else {
        v7 := []int{-2147483648, 374}
        _ = v7
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 uint8 = 42
            fmt.Println(fi8, fi8)
            v10 = v10
            var v11 uint16 = 1
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 float64 = 1.5
        _ = fi8
        _ = v12
    }
    v13 := GenS2{F0: "", F1: ("hello" + (("" + "hello") + "\n\t")), F2: 100}
    if (((3.14159 / float32((0.1 + (3.14159 - (0.1 + 0.1))))) >= (3.14159 - float32(float64(((3.14159 + 1.5) + 1e10))))) || false) {
        var v14 int = -9223372036854775808
        _ = v14
    }
    _ = v13
    return (true && true)
}

func main() {
    var v1 bool = true
    v2 := make(map[string]int)
    var v3 int32 = 0
    var v4 int16 = 1
    v5 := make(map[uint32]int)
    var v15 float64 = float64(3.14159)
    v16 := []int{1}
    if ((v3 ^ -236) < v3) {
        switch "" {
        case "foo bar":
            if (((v1 || (v1 && v1)) || v1) || v1) {
                v1 = true
                for fi17 := 0; fi17 < 15; fi17++ {
                    if v1 {
                        v18 := GenS3{F0: 1, F1: (-(v4))}
                        _ = v18
                    } else {
                        if v1 {
                            v19 := GenS4{F0: 42}
                            _ = v19
                        } else {
                            v15 = (v15 * v15)
                            v20 := GenS5{F0: fi17, F1: v1, F2: 905}
                            delete(v5, 430)
                            _ = v20
                        }
                        v4 -= v4
                        fmt.Println(v3, v1)
                    }
                    _ = fi17
                }
                v1 = (v1 && (0 < 670))
            }
        case "a":
            v21 := GenS1{F0: 207}
            _ = v21
        case "":
            fmt.Println(v15, v4)
        default:
            v1 = (!(true || v1))
        }
    }
    v22 := v2["\n\t"]
    var v23 int8 = -128
    var v24 uint8 = 0
    if len(v16) > 0 {
        v16[0] = (v22 / (((v22 ^ v22)) | 1))
    }
    if true {
        if v1 {
            for fi25 := 0; fi25 < 10; fi25++ {
                fmt.Println(v4, v1)
                _ = fi25
            }
            var v26 uint64 = 2
            _ = v26
        } else {
            v15 -= v15
        }
        for fi27 := 0; fi27 < 17; fi27++ {
            if ((true && (!v1)) || ("foo bar" == ("" + ((("foo bar" + "") + "a") + "a")))) {
                var v28 int64 = 0
                fmt.Println(v4)
                switch (fi27 | (fi27 << ((65535) & 63))) {
                case 32768:
                    var v29 string = "foo bar"
                    _ = v29
                default:
                    for fi30 := 0; fi30 < 13; fi30++ {
                        var v31 float32 = 3.14159
                        var v32 uint = 4294967295
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                }
                _ = v28
            } else {
                var v33 uint32 = 65535
                for fi34 := 0; fi34 < 8; fi34++ {
                    v35 := GenS4{F0: v33}
                    if (v1 && (v1 && v1)) {
                        var v36 int32 = (v3 << ((633) & 63))
                        _ = v36
                    } else {
                        v37 := v2["a"]
                        _ = v37
                    }
                    switch "foo bar" {
                    case "\n\t":
                        var v38 uint16 = 1
                        _ = v38
                    default:
                        fmt.Println(v23, v15, v23)
                    }
                    _ = fi34
                    _ = v35
                }
                for fi39 := 0; fi39 < 7; fi39++ {
                    var v40 int8 = (v23 % ((v23) | 1))
                    v4 += v4
                    _ = fi39
                    _ = v40
                }
                _ = v33
            }
            _ = fi27
        }
        delete(v2, "a")
    }
    v1 = true
    if v1 {
        v23 -= 100
        v22 -= -682
        var v41 bool = false
        _ = v41
    } else {
        if true {
            v15 -= (v15 - ((float64(((0.1 * -2.25) * 0.1)) * (1.5 / v15)) - -2.25))
            var v42 int64 = 9223372036854775807
            _ = v42
        } else {
            for fi43 := 0; fi43 < 12; fi43++ {
                var v44 string = "a"
                _ = fi43
                _ = v44
            }
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        }
        var v45 int64 = 42
        v22 += (v22 % ((v22) | 1))
        _ = v45
    }
    fmt.Println(v4)
    for fi46 := 0; fi46 < 19; fi46++ {
        switch ("a" + "hello") {
        case "foo bar":
            var v47 uint16 = 866
            _ = v47
        case "hello":
            var v48 uint32 = 4294967295
            _ = v48
        case "a":
            for fi49 := 0; fi49 < 12; fi49++ {
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                _ = fi49
            }
        }
        switch (fi46 % ((v22) | 1)) {
        case -2147483648:
            v22 += (v22 ^ (fi46 >> ((1) & 63)))
        case 228:
            fmt.Println(v1)
        case 127:
            for fi50 := 0; fi50 < 2; fi50++ {
                v51 := GenS6{F0: (v23 + (v23 << ((18446744073709551615) & 63))), F1: 908}
                switch 9223372036854775807 {
                case -976:
                    v1 = ((65535 > 2147483647) && (v1 || ((((^(v23)) == v23) && (2 >= 2)) && (v3 != (v3 % (((v3 << ((2) & 63))) | 1))))))
                case 42:
                    if ((1.5 + 3.14159) >= 3.14159) {
                        v1 = false
                        v2[(("a" + "a") + "a")] = fi46
                        v52 := GenS1{F0: 2}
                        v24 = 127
                        _ = v52
                    } else {
                        switch 779 {
                        case 9223372036854775807:
                            v15 -= (float64((1e10 + 1.5)) + (((v15 / (v15 + v15)) + v15) + v15))
                        case 1:
                            switch 688 {
                            case 127:
                                var v53 uint32 = 65535
                                _ = v53
                            case 9223372036854775807:
                                v24 += v24
                            default:
                                v54 := GenS7{F0: 509, F1: 250}
                                _ = v54
                            }
                        default:
                            if (v23 <= v23) {
                                var v55 uint8 = 2
                                switch (("a" + "hello") + "a") {
                                case "\n\t":
                                    v22 += (fi46 - fi46)
                                    if (!((float64(3.14159) / 1.0) <= float64(1e10))) {
                                        var v56 float32 = -2.25
                                        fmt.Println(fi50)
                                        var v57 int8 = (v23 ^ 42)
                                        var v58 int8 = (v23 / (((^(v57))) | 1))
                                        var v59 string = "\n\t"
                                        _ = v56
                                        _ = v57
                                        _ = v58
                                        _ = v59
                                    }
                                case "":
                                    v60 := GenS8{F0: 42, F1: 18446744073709551615}
                                    if len(v16) >= 2 {
                                        v16 = v16[1:len(v16)]
                                    }
                                    _ = v60
                                default:
                                    v61 := GenS3{F0: 0, F1: (v4 % ((v4) | 1))}
                                    _ = v61
                                }
                                v4 = v4
                                v24 = (v24 * (v55 >> ((440) & 63)))
                                _ = v55
                            }
                        }
                        var v62 string = (("a" + ((((("foo bar" + ("foo bar" + (("a" + "") + "foo bar"))) + (("hello" + ("foo bar" + "a")) + "hello")) + "\n\t") + "") + "a")) + ("a" + "hello"))
                        _ = v62
                    }
                }
                _ = fi50
                _ = v51
            }
        }
        _ = fi46
    }
    v5[1] = 128
    switch "" {
    case "hello":
        for fi63 := 0; fi63 < 13; fi63++ {
            switch "a" {
            case "hello":
                v4 += v4
            }
            switch -301 {
            case 4294967295:
                fmt.Println(v3)
            case 1:
                if false {
                    fmt.Println(v4, v23)
                    v64 := v5[707]
                    _ = v64
                }
            default:
                var v65 uint16 = 356
                _ = v65
            }
            var v66 int32 = 0
            _ = fi63
            _ = v66
        }
    }
    var v67 int16 = (-(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v67:", v67)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v67
}
