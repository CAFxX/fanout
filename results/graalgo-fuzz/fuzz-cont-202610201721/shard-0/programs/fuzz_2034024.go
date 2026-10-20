package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 int16; F1 uint16 }
type GenS8 struct { F0 uint16; F1 string }
type GenS9 struct { F0 uint32; F1 int64 }
func genFn0(p0 int, p1 string) uint32 {
    var v5 uint = 228
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := []int{(^(p0)), p0}
        _ = fi6
        _ = v7
    }
    if (42 != -1) {
        if true {
            if false {
                fmt.Println(p0, v5, p0)
            } else {
                if false {
                    p1 = p1
                    if false {
                        if (!(654 >= 2147483647)) {
                            v8 := []int{(p0 | p0)}
                            if false {
                                var v9 uint64 = 971
                                for fi10 := 0; fi10 < 9; fi10++ {
                                    var v11 uint16 = 255
                                    v12 := GenS1{F0: v5, F1: -769}
                                    switch (fi10 >> ((v5) & 63)) {
                                    case 32767:
                                        for fi13 := 0; fi13 < 10; fi13++ {
                                            v8 = append(v8, fi13)
                                            v5 = 42
                                            _ = fi13
                                        }
                                    case 1:
                                        fmt.Println(p0)
                                    }
                                    _ = fi10
                                    _ = v11
                                    _ = v12
                                }
                                _ = v9
                            }
                            fmt.Println(p0, v5)
                            _ = v8
                        }
                        var v14 float64 = 0.1
                        for fi15 := 0; fi15 < 7; fi15++ {
                            p0 = p0
                            _ = fi15
                        }
                        _ = v14
                    } else {
                        v16 := GenS1{F0: v5, F1: 42}
                        for fi17 := 0; fi17 < 5; fi17++ {
                            var v18 int32 = 2147483647
                            v19 := GenS2{F0: int16(v18), F1: 653}
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        if (2 != 255) {
                            var v20 int = (p0 ^ (p0 + (p0 & (p0 / ((-1) | 1)))))
                            _ = v20
                        } else {
                            p1 = (p1 + p1)
                        }
                        _ = v16
                    }
                }
                p1 = p1
            }
            v21 := []int{-121, (^(p0)), (p0 + p0), p0}
            _ = v21
        }
        p1 = p1
        var v22 float32 = (((float32(1.5) / -2.25) - 1e10) * (0.0 + float32(1e10)))
        _ = v22
    } else {
        p1 = p1
        v5 = v5
    }
    _ = p0
    _ = p1
    _ = v5
    return 86
}

func main() {
    var v1 int32 = 1
    var v2 float32 = 3.14159
    v3 := []string{"a"}
    v4 := make(map[uint32]uint32)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2, v2)
    var v23 int = -9223372036854775808
    var v24 int16 = -1
    switch v23 {
    case -9223372036854775808:
        v25 := GenS3{F0: genFn0((v23 ^ -2147483648), (("hello" + "foo bar") + "")), F1: 65535}
        _ = v25
    case -2147483648:
        v26 := GenS2{F0: (v24 & (^(v24))), F1: 1}
        _ = v26
    case 9223372036854775807:
        if (false && false) {
            v23 = 809
            for fi27 := 0; fi27 < 19; fi27++ {
                fmt.Println(v1)
                v28 := GenS4{F0: -694}
                _ = fi27
                _ = v28
            }
            v29 := GenS5{F0: 42, F1: -9223372036854775808}
            _ = v29
        }
    default:
        if true {
            fmt.Println(v2, v2)
            var v30 uint16 = 500
            for fi31 := 0; fi31 < 10; fi31++ {
                if (1 > (v30 >> ((297) & 63))) {
                    var v32 int8 = 0
                    _ = v32
                } else {
                    if true {
                        for fi33 := 0; fi33 < 5; fi33++ {
                            var v34 bool = true
                            if true {
                                var v35 string = "hello"
                                v34 = (!(65535 != 4294967295))
                                v35 = "foo bar"
                                _ = v35
                            } else {
                                var v36 int32 = v1
                                _ = v36
                            }
                            v37 := GenS6{F0: fi31, F1: 9223372036854775807}
                            if ((!(("foo bar" + "") != "foo bar")) || (384 <= 42)) {
                                v4[2] = 65535
                            }
                            if ("foo bar" != "a") {
                                v3 = append(v3, "foo bar")
                                if v34 {
                                    v38 := GenS3{F0: 2147483647, F1: v30}
                                    _ = v38
                                }
                            }
                            v39 := GenS3{F0: 42, F1: 32767}
                            switch v23 {
                            case 646:
                                var v40 uint16 = (v30 & v30)
                                _ = v40
                            case 127:
                                v41 := v4[2147483648]
                                _ = v41
                            case -9223372036854775808:
                                if false {
                                    _ = len(v4)
                                    v34 = v34
                                    v34 = ((!v34) && ((^(v30)) >= 65535))
                                }
                            }
                            _ = fi33
                            _ = v34
                            _ = v37
                            _ = v39
                        }
                        v3 = append(v3, "\n\t")
                        _ = len(v4)
                    } else {
                        for fi42 := 0; fi42 < 5; fi42++ {
                            fmt.Println(v30, fi31)
                            _ = fi42
                        }
                        v43 := GenS7{F0: (v24 - -32768), F1: v30}
                        if (42 >= 90) {
                            for fi44 := 0; fi44 < 7; fi44++ {
                                _ = len(v3)
                                v2 -= float32(1.5)
                                v45 := GenS8{F0: (v30 | 0), F1: "hello"}
                                _ = fi44
                                _ = v45
                            }
                            fmt.Println(v24, v30)
                        }
                        _ = v43
                    }
                    v23 = v23
                    if (4294967295 == 327) {
                        var v46 int64 = 344
                        switch -697 {
                        case -9223372036854775808:
                            for fi47 := 0; fi47 < 6; fi47++ {
                                var v48 float32 = (v2 / (v2 * v2))
                                v46 = v46
                                if false {
                                    if len(v3) > 0 {
                                        v3[0] = ("\n\t" + "")
                                    }
                                    v49 := GenS9{F0: 2147483648, F1: (^(v46))}
                                    _ = v49
                                } else {
                                    v23 -= (v23 ^ (fi31 | 0))
                                }
                                _ = fi47
                                _ = v48
                            }
                        case -2147483648:
                            var v50 int8 = 92
                            _ = v50
                        case -889:
                            v23 -= v23
                        }
                        var v51 int = (v23 - fi31)
                        _ = v46
                        _ = v51
                    }
                }
                v4[2147483648] = 68
                _ = fi31
            }
            _ = v30
        } else {
            v52 := GenS9{F0: 2, F1: -9223372036854775808}
            v24 = v24
            v53 := GenS4{F0: -1}
            _ = v52
            _ = v53
        }
    }
    for fi54 := 0; fi54 < 7; fi54++ {
        v1 += v1
        var v55 string = "a"
        _ = fi54
        _ = v55
    }
    var v56 uint64 = 4294967295
    for fi57 := 0; fi57 < 9; fi57++ {
        v58 := GenS9{F0: 447, F1: -1}
        delete(v4, 694)
        _ = fi57
        _ = v58
    }
    for fi59 := 0; fi59 < 12; fi59++ {
        if (1 != -33) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi60 := 0; fi60 < 14; fi60++ {
                v56 = 0
                v2 = v2
                v1 += v1
                var v61 uint = 1
                _ = fi60
                _ = v61
            }
        } else {
            var v62 int = (-(v23))
            _ = v62
        }
        _ = fi59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v56
}
