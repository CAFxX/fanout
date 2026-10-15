package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int; F1 int8 }
func genFn0() uint64 {
    var v7 uint32 = 349
    v8 := GenS1{F0: (true && false), F1: 42, F2: 939}
    v9 := GenS1{F0: (-718 >= -7), F1: -32768, F2: (v7 >> ((1) & 63))}
    v10 := []int{-543}
    var v11 uint32 = 65535
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 int, p1 int64, p2 int32) int32 {
    switch (-(p0)) {
    case 4294967295:
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 uint64 = 301
            v14 := []int{p0, p0}
            switch "\n\t" {
            case "foo bar":
                var v15 float32 = (((1.5 + 3.14159) * (0.1 + 1.5)) - ((1.5 / (float32(1.5) - ((0.0 - 1.5) / 1e10))) - float32(0.1)))
                _ = v15
            case "hello":
                switch "\n\t" {
                case "hello":
                    fmt.Println(p1)
                default:
                    var v16 int16 = 1
                    _ = v16
                }
            default:
                p0 += p0
            }
            _ = fi12
            _ = v13
            _ = v14
        }
    default:
        for fi17 := 0; fi17 < 10; fi17++ {
            v18 := GenS1{F0: false, F1: -742, F2: 2}
            v19 := GenS1{F0: false, F1: 32767, F2: 837}
            var v20 int32 = (p2 & -134)
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
    }
    if (42 < 1) {
        v21 := []int{p0, p0}
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    return -876
}

func main() {
    var v1 int32 = 794
    v2 := []string{"hello"}
    v3 := []uint32{2147483647}
    v4 := make(map[string]int)
    v5 := make(map[int]int)
    var v6 string = "a"
    v2 = append(v2, (v6 + "\n\t"))
    v1 += v1
    if true {
        switch 32767 {
        case 32767:
            var v22 float32 = 0.1
            _ = v22
        case 32768:
            v6 = "a"
        case -9223372036854775808:
            v23 := GenS1{F0: (-9223372036854775808 < -9223372036854775808), F1: int16(v1), F2: 734}
            _ = v23
        default:
            v1 = v1
        }
        var v24 float64 = 1.5
        _ = v24
    } else {
        v25 := GenS1{F0: false, F1: -32768, F2: 4294967295}
        v6 = (v6 + (v6 + "foo bar"))
        if true {
            if false {
                v6 = v6
                v6 = ("\n\t" + ((v6 + v6) + ("" + ("a" + v6))))
            } else {
                for fi26 := 0; fi26 < 10; fi26++ {
                    for fi27 := 0; fi27 < 16; fi27++ {
                        if (v1 > 573) {
                            v28 := GenS2{F0: 781, F1: 10}
                            v29 := GenS3{F0: 65535, F1: 0}
                            var v30 uint64 = 4294967295
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                        _ = fi27
                    }
                    _ = len(v4)
                    _ = fi26
                }
                _ = len(v5)
                v31 := GenS2{F0: 2, F1: 42}
                _ = v31
            }
            for fi32 := 0; fi32 < 3; fi32++ {
                v6 = v6
                switch ((v6 + "") + v6) {
                case "hello":
                    fmt.Println(fi32, v6, v6)
                default:
                    switch (fi32 % ((-1) | 1)) {
                    case -2147483648:
                        v33 := GenS4{F0: (v1 + v1)}
                        _ = v33
                    case 0:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                }
                _ = fi32
            }
            for fi34 := 0; fi34 < 20; fi34++ {
                for fi35 := 0; fi35 < 13; fi35++ {
                    var v36 int8 = 121
                    v1 = (v1 << ((441) & 63))
                    _ = fi35
                    _ = v36
                }
                _ = fi34
            }
        } else {
            if ((203 < 0) || false) {
                if false {
                    fmt.Println(v1, v6)
                } else {
                    for fi37 := 0; fi37 < 10; fi37++ {
                        var v38 int16 = 807
                        var v39 uint8 = 128
                        if len(v3) > 0 {
                            v3[0] = 65535
                        }
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                    v40 := GenS2{F0: 255, F1: -83}
                    if (42 != 32768) {
                        var v41 int64 = 9223372036854775807
                        _ = v41
                    } else {
                        for fi42 := 0; fi42 < 2; fi42++ {
                            switch (((("hello" + ("a" + v6)) + v6) + (v6 + "\n\t")) + v6) {
                            case "":
                                if (1 <= 9223372036854775807) {
                                    var v43 uint = 4294967295
                                    _ = v43
                                } else {
                                    v1 = (v1 << ((4294967295) & 63))
                                    v6 = ""
                                }
                            case "foo bar":
                                v1 -= -372
                            default:
                                v6 = v6
                            }
                            var v44 uint8 = 126
                            v44 = 2
                            _ = fi42
                            _ = v44
                        }
                        fmt.Println(v6, v6)
                    }
                    _ = v40
                }
                var v45 bool = (2 != 65535)
                var v46 bool = v45
                _ = v45
                _ = v46
            } else {
                var v47 uint64 = 4294967295
                _ = v47
            }
        }
        _ = v25
    }
    fmt.Println(v1, v6)
    v1 = (v1 << ((42) & 63))
    fmt.Println(v1, v1)
    fmt.Println(v1, v6)
    var v48 bool = (false || (((((-32768 > 42) && (!true)) || (!false)) && (!((v1 * genFn1(32767, -9223372036854775808, v1)) == -1))) && false))
    switch "foo bar" {
    case "a":
        v49 := v5[32768]
        _ = v49
    default:
        if v48 {
            var v50 int = -2147483648
            v48 = (!(115 > 1))
            _ = v50
        } else {
            v51 := GenS5{F0: 32767, F1: 0}
            if v48 {
                if (v6 > (v6 + v6)) {
                    var v52 int32 = 1
                    var v53 bool = (127 < 128)
                    _ = v52
                    _ = v53
                } else {
                    v6 = v6
                }
                if (2 == 422) {
                    fmt.Println(v1, v6)
                    v54 := GenS4{F0: (v1 / (((-(v1))) | 1))}
                    v6 = v6
                    _ = v54
                }
                var v55 float64 = (((((0.1 / -2.25) + 3.14159) - float64(3.14159)) + 1.5) + 3.14159)
                _ = v55
            } else {
                v56 := v5[32768]
                _ = v56
            }
            var v57 int64 = 9223372036854775807
            _ = v51
            _ = v57
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
