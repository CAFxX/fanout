package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 bool; F1 int32 }
func genFn0(p0 int8) uint64 {
    v6 := GenS1{F0: 255}
    p0 += p0
    _ = p0
    _ = v6
    return 2
}
func genFn1(p0 uint32, p1 int8, p2 uint32) uint {
    v7 := GenS1{F0: 209}
    var v8 int = -514
    v9 := []int{(^(v8))}
    v8 = (v8 + (v8 * v8))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 int32 = 534
    var v2 uint64 = 2
    var v3 uint32 = 565
    v4 := make(map[uint32]int)
    var v5 uint64 = 42
    var v10 uint8 = uint8(v5)
    v2 = v2
    var v11 int16 = -658
    for fi12 := 0; fi12 < 10; fi12++ {
        v13 := GenS1{F0: (v10 >> ((65535) & 63))}
        for fi14 := 0; fi14 < 15; fi14++ {
            fmt.Println(v5, fi14, fi12)
            if true {
                switch fi12 {
                case -637:
                    var v15 int64 = int64(fi14)
                    _ = v15
                default:
                    v2 = v5
                }
            }
            _ = fi14
        }
        for fi16 := 0; fi16 < 18; fi16++ {
            v5 = genFn0(-1)
            var v17 string = ("" + "\n\t")
            fmt.Println(fi16, v2)
            _ = fi16
            _ = v17
        }
        _ = fi12
        _ = v13
    }
    if false {
        v18 := []int{127}
        _ = v18
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        v20 := GenS1{F0: (v10 | (v10 + v10))}
        v1 = (v1 ^ (v1 | (v1 * v1)))
        _ = fi19
        _ = v20
    }
    if true {
        v21 := GenS1{F0: v10}
        _ = v21
    } else {
        v4[v3] = -122
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        v3 = (v3 - 4294967295)
        if (true || false) {
            v2 += (^(v2))
            for fi23 := 0; fi23 < 9; fi23++ {
                if false {
                    v2 += (v5 | (^(v2)))
                    v1 = (v1 << ((genFn1(v3, -7, (v3 >> ((4294967295) & 63)))) & 63))
                    var v24 float64 = (float64(3.14159) / (1.5 + (float64(1e10) * ((0.1 + 0.0) + 3.14159))))
                    _ = v24
                }
                _ = fi23
            }
        } else {
            var v25 int32 = v1
            _ = len(v4)
            _ = v25
        }
        _ = fi22
    }
    if false {
        if true {
            if true {
                fmt.Println(v3, v11)
                if (127 == 1) {
                    var v26 bool = (-1 > (v1 >> ((0) & 63)))
                    switch -2147483648 {
                    case 0:
                        var v27 float32 = 3.14159
                        _ = v27
                    case -2147483648:
                        delete(v4, 552)
                    default:
                        var v28 int32 = v1
                        _ = v28
                    }
                    var v29 int = -494
                    _ = v26
                    _ = v29
                } else {
                    var v30 uint = 2
                    _ = v30
                }
                for fi31 := 0; fi31 < 4; fi31++ {
                    for fi32 := 0; fi32 < 11; fi32++ {
                        v5 = genFn0(-71)
                        var v33 int32 = -2147483648
                        fmt.Println(v10, v10)
                        _ = fi32
                        _ = v33
                    }
                    switch 907 {
                    case 101:
                        v34 := GenS1{F0: (v10 / (((v10 + 0)) | 1))}
                        _ = v34
                    case -9223372036854775808:
                        for fi35 := 0; fi35 < 15; fi35++ {
                            switch "a" {
                            case "foo bar":
                                v10 += (v10 + 2)
                            case "hello":
                                fmt.Println(v1, v10, fi35)
                                if (fi35 > (fi31 % (((fi31 | 32768)) | 1))) {
                                    var v36 int8 = -128
                                    fmt.Println(fi31, fi35, v5)
                                    _ = v36
                                } else {
                                    v1 -= v1
                                    v1 = v1
                                }
                            case "a":
                                fmt.Println(v10, v10, fi35)
                                v37 := GenS2{F0: 0, F1: 421}
                                v5 -= (v2 / ((v2) | 1))
                                fmt.Println(fi35)
                                _ = v37
                            default:
                                fmt.Println(v2, fi31, v1)
                            }
                            fmt.Println(fi31)
                            v3 += 2147483648
                            v38 := GenS3{F0: (!(true || (!true))), F1: v1}
                            _ = fi35
                            _ = v38
                        }
                    default:
                        var v39 string = "a"
                        _ = v39
                    }
                    _ = fi31
                }
            } else {
                v4[v3] = 32768
                fmt.Println(v5, v3)
            }
            for fi40 := 0; fi40 < 4; fi40++ {
                delete(v4, 65535)
                var v41 uint64 = genFn0(3)
                _ = fi40
                _ = v41
            }
        }
        var v42 uint64 = v5
        for fi43 := 0; fi43 < 5; fi43++ {
            for fi44 := 0; fi44 < 11; fi44++ {
                if false {
                    v45 := GenS2{F0: 42, F1: 0}
                    v46 := GenS1{F0: v10}
                    v42 += (v5 << ((327) & 63))
                    v47 := GenS3{F0: (("a" + "") <= ((("foo bar" + ("hello" + "a")) + (("" + ("foo bar" + "hello")) + "hello")) + ("" + "foo bar"))), F1: v1}
                    v48 := GenS2{F0: 0, F1: -1}
                    _ = v45
                    _ = v46
                    _ = v47
                    _ = v48
                } else {
                    v49 := v4[v3]
                    v3 += 689
                    v50 := []int{(fi44 ^ (^(fi43)))}
                    _ = v49
                    _ = v50
                }
                _ = fi44
            }
            _ = fi43
        }
        _ = v42
    }
    if false {
        var v51 uint8 = 127
        _ = v51
    }
    var v52 int8 = 42
    switch -885 {
    case 127:
        if (true || (42 == (v10 * (v10 * v10)))) {
            if (-528 > -335) {
                fmt.Println(v1, v3, v2)
                v53 := GenS2{F0: 56, F1: 612}
                _ = v53
            }
            var v54 int32 = -1
            _ = v54
        }
    default:
        v2 = (^(v5))
    }
    for fi55 := 0; fi55 < 6; fi55++ {
        var v56 int8 = v52
        _ = len(v4)
        _ = fi55
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
}
