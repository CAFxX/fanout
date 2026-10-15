package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 uint64; F2 string }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 int64; F1 string }
func genFn0(p0 int16) string {
    if (false || ("" >= "")) {
        v5 := GenS1{F0: (^(p0)), F1: (p0 | p0)}
        v6 := GenS1{F0: p0, F1: p0}
        _ = v5
        _ = v6
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(p0)
    case "foo bar":
        if (false || true) {
            p0 = p0
            for fi7 := 0; fi7 < 8; fi7++ {
                v8 := GenS1{F0: p0, F1: p0}
                p0 = (p0 >> ((42) & 63))
                switch 2147483647 {
                case -555:
                    v9 := []int{1, (fi7 << ((65535) & 63))}
                    _ = v9
                case -295:
                    v10 := GenS2{F0: (fi7 | fi7)}
                    _ = v10
                }
                _ = fi7
                _ = v8
            }
            var v11 float64 = 3.14159
            _ = v11
        }
    case "hello":
        if true {
            var v12 int16 = 42
            v13 := []int{42, -1, 2147483647}
            _ = v12
            _ = v13
        }
    default:
        if (!(!false)) {
            for fi14 := 0; fi14 < 2; fi14++ {
                v15 := GenS1{F0: (p0 * p0), F1: (p0 - 362)}
                _ = fi14
                _ = v15
            }
            fmt.Println(p0)
            v16 := GenS3{F0: 0}
            _ = v16
        } else {
            for fi17 := 0; fi17 < 8; fi17++ {
                v18 := []int{(-(fi17))}
                if false {
                    fmt.Println(p0, fi17)
                } else {
                    v19 := GenS3{F0: 0}
                    fmt.Println(p0, p0)
                    _ = v19
                }
                _ = fi17
                _ = v18
            }
            var v20 uint = 42
            if true {
                v21 := GenS1{F0: (p0 + 1000), F1: (p0 >> ((473) & 63))}
                _ = v21
            } else {
                for fi22 := 0; fi22 < 11; fi22++ {
                    var v23 int64 = -454
                    _ = fi22
                    _ = v23
                }
                if (false && true) {
                    v24 := GenS1{F0: (p0 % (((p0 << (((v20 / (((^(v20))) | 1))) & 63))) | 1)), F1: 112}
                    p0 += p0
                    p0 = (p0 | p0)
                    _ = v24
                } else {
                    fmt.Println(p0, v20)
                    var v25 int32 = 0
                    fmt.Println(p0)
                    _ = v25
                }
            }
            _ = v20
        }
    }
    for fi26 := 0; fi26 < 10; fi26++ {
        v27 := []int{(fi26 & fi26), 516, fi26}
        _ = fi26
        _ = v27
    }
    _ = p0
    return ("" + "\n\t")
}
func genFn1(p0 uint64, p1 bool, p2 uint) uint64 {
    v28 := []int{9223372036854775807, 1}
    p2 = (p2 + uint(p0))
    var v29 uint64 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    _ = v29
    return v29
}

func main() {
    v1 := []uint32{}
    var v2 float32 = (3.14159 * ((-2.25 / -2.25) + float32(3.14159)))
    v3 := []string{"\n\t", ("hello" + "foo bar"), ((("hello" + "\n\t") + "foo bar") + ""), "\n\t"}
    var v4 int32 = -2147483648
    for fi30 := 0; fi30 < 3; fi30++ {
        if true {
            var v31 int32 = 1
            var v32 int16 = -51
            v1 = append(v1, 947)
            _ = v31
            _ = v32
        } else {
            v4 = v4
            var v33 uint = 18446744073709551615
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v33
        }
        v34 := GenS1{F0: 1000, F1: -233}
        for fi35 := 0; fi35 < 9; fi35++ {
            var v36 bool = true
            if v36 {
                if v36 {
                    v2 += (0.1 / v2)
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if (4294967295 < 977) {
                    switch 2147483647 {
                    case -1:
                        var v37 int8 = -71
                        _ = v37
                    case 0:
                        v38 := GenS4{F0: 616, F1: 2, F2: ("hello" + ("foo bar" + ""))}
                        _ = v38
                    case 512:
                        v39 := GenS2{F0: (-(fi35))}
                        _ = v39
                    }
                    var v40 bool = (1e10 == (v2 + 0.1))
                    _ = v40
                }
                var v41 uint16 = 32767
                _ = v41
            }
            _ = fi35
            _ = v36
        }
        _ = fi30
        _ = v34
    }
    fmt.Println(v4)
    for fi42 := 0; fi42 < 17; fi42++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi43 := 0; fi43 < 16; fi43++ {
            fmt.Println(fi42, fi42, v2)
            v4 = -2147483648
            v44 := GenS5{F0: 2, F1: 42}
            _ = fi43
            _ = v44
        }
        v4 = 2147483647
        _ = fi42
    }
    if (((true || (true || (float32(1.5) <= v2))) || true) && (255 >= 2)) {
        v1 = append(v1, 705)
        var v45 uint32 = 1
        fmt.Println(v2)
        _ = v45
    }
    v46 := GenS2{F0: 32768}
    if false {
        v2 = (1.5 / 1.0)
        if (-2147483648 != 42) {
            if (-32768 == -1) {
                for fi47 := 0; fi47 < 14; fi47++ {
                    for fi48 := 0; fi48 < 5; fi48++ {
                        v49 := GenS6{F0: 1, F1: fi48}
                        var v50 int32 = 42
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                    if false {
                        fmt.Println(fi47, v2)
                    } else {
                        v4 = v4
                        if (!(!(!(!(58 < 0))))) {
                            for fi51 := 0; fi51 < 2; fi51++ {
                                v4 += (v4 | v4)
                                fmt.Println(v4, v2, v4)
                                fmt.Println(v4)
                                _ = fi51
                            }
                            var v52 bool = (118 >= 2147483648)
                            fmt.Println(v4, fi47, v2)
                            _ = v52
                        }
                    }
                    var v53 float64 = -2.25
                    _ = fi47
                    _ = v53
                }
                for fi54 := 0; fi54 < 17; fi54++ {
                    fmt.Println(v4)
                    for fi55 := 0; fi55 < 4; fi55++ {
                        _ = len(v1)
                        var v56 uint32 = 475
                        _ = fi55
                        _ = v56
                    }
                    _ = fi54
                }
            }
            var v57 int = -903
            var v58 uint8 = 44
            _ = v57
            _ = v58
        }
    } else {
        fmt.Println(v4, v2)
    }
    var v59 uint8 = 0
    v60 := GenS7{F0: -9223372036854775808, F1: "\n\t"}
    v61 := GenS5{F0: 8, F1: 0}
    var v62 int32 = int32(v59)
    v2 = ((float32(0.0) * float32((3.14159 / 1.0))) * (1.5 - v2))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v59:", v59)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v59
    _ = v60
    _ = v61
    _ = v62
}
