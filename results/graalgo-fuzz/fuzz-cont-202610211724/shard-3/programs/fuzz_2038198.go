package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 bool }
type GenS2 struct { F0 int; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint; F1 int32; F2 bool }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 int64; F1 uint64 }
func genFn0(p0 string) int {
    v4 := []int{}
    var v5 uint16 = 65535
    v5 = (^(v5))
    for fi6 := 0; fi6 < 19; fi6++ {
        if (!(!(!true))) {
            if (892 > uint32(v5)) {
                v7 := GenS1{F0: 127, F1: 0, F2: false}
                var v8 uint32 = 822
                _ = v7
                _ = v8
            } else {
                fmt.Println(fi6)
                for fi9 := 0; fi9 < 14; fi9++ {
                    v10 := GenS1{F0: 10, F1: 1, F2: true}
                    _ = len(v4)
                    if len(v4) > 0 {
                        v4[0] = fi9
                    }
                    _ = fi9
                    _ = v10
                }
            }
            fmt.Println(fi6, fi6)
        }
        if true {
            var v11 uint16 = (v5 << ((0) & 63))
            if len(v4) > 0 {
                v4[0] = fi6
            }
            v12 := GenS2{F0: 0, F1: 127, F2: v5}
            _ = v11
            _ = v12
        } else {
            var v13 int = fi6
            if true {
                v4 = append(v4, v13)
                for fi14 := 0; fi14 < 10; fi14++ {
                    _ = len(v4)
                    _ = fi14
                }
            }
            v13 -= fi6
            _ = v13
        }
        _ = fi6
    }
    v5 -= v5
    _ = p0
    _ = v4
    _ = v5
    return 4294967295
}
func genFn1(p0 int16, p1 int8, p2 uint8) int8 {
    for fi15 := 0; fi15 < 14; fi15++ {
        p0 += p0
        _ = fi15
    }
    if false {
        p1 = (p1 >> ((uint(p2)) & 63))
        if ((952 == 662) || (42 == 42)) {
            p1 = p1
            var v16 int = -2147483648
            if true {
                p1 += p1
            }
            _ = v16
        } else {
            switch "\n\t" {
            case "hello":
                p2 += p2
            default:
                v17 := []int{32768, 728, 9223372036854775807, -748}
                _ = v17
            }
        }
        var v18 uint32 = 868
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 127
    var v3 int64 = -9223372036854775808
    if false {
        for fi19 := 0; fi19 < 9; fi19++ {
            v20 := GenS3{F0: 65535, F1: -683, F2: (true || (false || ((3.14159 > -2.25) || (false || true))))}
            delete(v1, 146)
            _ = fi19
            _ = v20
        }
        switch "foo bar" {
        case "a":
            if false {
                for fi21 := 0; fi21 < 9; fi21++ {
                    v22 := []int{genFn0("\n\t")}
                    _ = fi21
                    _ = v22
                }
            }
        default:
            for fi23 := 0; fi23 < 17; fi23++ {
                v24 := []int{genFn0(""), 920, fi23}
                fmt.Println(v2)
                _ = fi23
                _ = v24
            }
        }
    }
    v25 := v1[2147483647]
    v26 := []int{}
    var v27 int = genFn0(("a" + (v25 + (v25 + v25))))
    switch v27 {
    case -2147483648:
        var v28 float32 = -2.25
        _ = v28
    case 1:
        v29 := GenS3{F0: 1, F1: -111, F2: false}
        _ = v29
    }
    if false {
        v30 := GenS4{F0: 1, F1: 0}
        switch v25 {
        case "\n\t":
            fmt.Println(v25, v25, v2)
        case "foo bar":
            var v31 uint64 = 2
            _ = v31
        }
        _ = v30
    }
    var v32 int16 = 179
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    var v33 uint16 = 617
    var v34 int = (v27 >> ((18446744073709551615) & 63))
    v35 := GenS4{F0: 255, F1: 42}
    switch genFn0((v25 + "foo bar")) {
    case 4294967295:
        var v36 int32 = 1
        _ = v36
    case -424:
        var v37 uint = uint(v33)
        _ = v37
    case -254:
        if false {
            if false {
                if false {
                    v3 -= 0
                }
                if ((^(v33)) != (v33 >> ((720) & 63))) {
                    switch v25 {
                    case "":
                        _ = len(v26)
                    default:
                        if ((!(false && false)) || false) {
                            if len(v26) > 0 {
                                v26[0] = v27
                            }
                            var v38 string = ("\n\t" + v25)
                            v26 = append(v26, genFn0((v25 + ((v25 + v38) + v38))))
                            _ = v38
                        }
                    }
                    fmt.Println(v32, v2, v2)
                    var v39 int8 = genFn1((^(v32)), genFn1(v32, (-(v2)), uint8(v33)), 127)
                    _ = v39
                }
                v3 = v3
            }
            v33 += 255
            for fi40 := 0; fi40 < 4; fi40++ {
                v2 = -77
                _ = fi40
            }
        } else {
            switch v25 {
            case "":
                for fi41 := 0; fi41 < 18; fi41++ {
                    v42 := GenS1{F0: v2, F1: (v32 + (v32 & v32)), F2: false}
                    _ = fi41
                    _ = v42
                }
            }
            if len(v26) > 0 {
                v26[0] = (v34 / (((^(v27))) | 1))
            }
        }
    default:
        for fi43 := 0; fi43 < 7; fi43++ {
            if (42 > 221) {
                for fi44 := 0; fi44 < 11; fi44++ {
                    v27 = (v27 & (v27 >> ((515) & 63)))
                    _ = fi44
                }
            } else {
                if true {
                    switch v25 {
                    case "foo bar":
                        var v45 int = (v34 | (fi43 << ((uint(v33)) & 63)))
                        _ = v45
                    case "":
                        v2 -= (v2 - 100)
                    }
                    _ = len(v26)
                    for fi46 := 0; fi46 < 13; fi46++ {
                        switch v25 {
                        case "\n\t":
                            var v47 uint8 = 42
                            v25 = "a"
                            var v48 int32 = 2147483647
                            v49 := GenS2{F0: genFn0("\n\t"), F1: v47, F2: (v33 * v33)}
                            var v50 float32 = 3.14159
                            _ = v47
                            _ = v48
                            _ = v49
                            _ = v50
                        case "":
                            var v51 bool = (1 >= 1)
                            _ = v51
                        case "hello":
                            if (42 == 128) {
                                if (323 == 0) {
                                    v52 := GenS1{F0: 45, F1: (v32 & (^(v32))), F2: (!true)}
                                    var v53 uint16 = v33
                                    if false {
                                        if len(v26) >= 2 {
                                            v26 = v26[1:len(v26)]
                                        }
                                        fmt.Println(fi46, v2, v53)
                                    }
                                    _ = v52
                                    _ = v53
                                } else {
                                    fmt.Println(fi46, v27)
                                }
                            } else {
                                if len(v26) > 0 {
                                    v26[0] = genFn0(v25)
                                }
                            }
                        default:
                            var v54 string = "foo bar"
                            _ = v54
                        }
                        v34 = (-(fi46))
                        _ = fi46
                    }
                }
                for fi55 := 0; fi55 < 14; fi55++ {
                    var v56 uint32 = 683
                    _ = fi55
                    _ = v56
                }
            }
            _ = fi43
        }
    }
    var v57 int64 = int64(v2)
    fmt.Println(v32)
    switch (^(v27)) {
    case 0:
        _ = len(v1)
    case 128:
        v58 := GenS3{F0: 602, F1: 42, F2: (1e10 <= float64(0.1))}
        _ = v58
    case 32768:
        v59 := GenS5{F0: (-(v3)), F1: 992}
        _ = v59
    default:
        var v60 int64 = 42
        _ = v60
    }
    v1[2147483647] = "\n\t"
    for fi61 := 0; fi61 < 7; fi61++ {
        delete(v1, 0)
        v27 += genFn0(v25)
        _ = fi61
    }
    var v62 uint32 = 585
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v57:", v57)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v57
    _ = v62
}
