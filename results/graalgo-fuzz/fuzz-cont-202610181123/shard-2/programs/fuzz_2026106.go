package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int16, p1 uint, p2 int8) uint64 {
    v6 := GenS1{F0: 0}
    v7 := []int{}
    v8 := GenS2{F0: 1, F1: 0, F2: 499}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"a"}
    var v3 float64 = -2.25
    v4 := make(map[uint32]uint32)
    v5 := []int{9223372036854775807, 524, 32768, 4294967295}
    switch ("a" + "foo bar") {
    case "foo bar":
        if false {
            if true {
                switch -783 {
                case 2147483647:
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                case 32768:
                    _ = len(v5)
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                var v9 uint = 721
                _ = v9
            }
            var v10 uint = 493
            v11 := v1[-903]
            _ = v10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        v4[65535] = 2
        var v13 int32 = 0
        _ = fi12
        _ = v13
    }
    var v14 uint16 = 0
    v15 := GenS1{F0: 65535}
    var v16 float64 = (((1e10 - (v3 * v3)) + 0.0) + 0.1)
    var v17 uint16 = 65535
    if (42 >= 2) {
        var v18 uint8 = 2
        if true {
            if (!(false && ((2147483647 > -1) || true))) {
                if (((float64(0.0) * (1.5 / (1e10 - ((-2.25 / (v3 / v16)) / 1.0)))) != float64((((0.1 / (1e10 - -2.25)) + (-2.25 + (((1e10 - 0.0) + ((((1.5 * 3.14159) + (1e10 + 0.1)) - 3.14159) + ((((3.14159 * (0.1 + 0.0)) * 1.5) - 0.1) * ((1.5 + float32((v16 / (((3.14159 + float64((float32(((v16 * v16) * 1e10)) - float32(((0.0 / v3) * v3))))) * v16) / v16)))) + (3.14159 / 3.14159))))) + ((0.1 / 1e10) - 0.0)))) + 0.0))) || false) {
                    switch 1 {
                    case 42:
                        for fi19 := 0; fi19 < 11; fi19++ {
                            v20 := GenS2{F0: -9223372036854775808, F1: (v17 ^ (v17 / (((v17 << ((65535) & 63))) | 1))), F2: v14}
                            _ = fi19
                            _ = v20
                        }
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
            if true {
                v21 := v4[2]
                _ = v21
            } else {
                if true {
                    if (!(!false)) {
                        var v22 uint64 = 264
                        var v23 uint = 1
                        var v24 int16 = int16(v22)
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                }
                if false {
                    _ = len(v2)
                }
            }
            v16 = v16
        } else {
            for fi25 := 0; fi25 < 3; fi25++ {
                switch -2147483648 {
                case 4294967295:
                    for fi26 := 0; fi26 < 16; fi26++ {
                        v14 = v17
                        v27 := GenS2{F0: 0, F1: 65535, F2: (v14 << ((2) & 63))}
                        fmt.Println(fi26)
                        _ = fi26
                        _ = v27
                    }
                case 1:
                    var v28 uint32 = 641
                    _ = v28
                default:
                    var v29 int32 = -221
                    _ = v29
                }
                v30 := GenS3{F0: 2}
                v17 -= 567
                _ = fi25
                _ = v30
            }
        }
        fmt.Println(v16, v16, v14)
        _ = v18
    } else {
        var v31 uint64 = 961
        if ((v31 % ((789) | 1)) != (^(v31))) {
            if false {
                if false {
                    for fi32 := 0; fi32 < 11; fi32++ {
                        if true {
                            for fi33 := 0; fi33 < 13; fi33++ {
                                if true {
                                    switch (fi33 ^ (fi33 >> ((65535) & 63))) {
                                    case 0:
                                        if len(v5) > 0 {
                                            v5[0] = 1
                                        }
                                    case 794:
                                        fmt.Println(v3, v3, v16)
                                        v34 := GenS1{F0: 65535}
                                        var v35 uint64 = v31
                                        _ = v34
                                        _ = v35
                                    }
                                    var v36 string = (("a" + "foo bar") + (("\n\t" + "hello") + ""))
                                    _ = v36
                                }
                                var v37 string = ""
                                _ = fi33
                                _ = v37
                            }
                            for fi38 := 0; fi38 < 17; fi38++ {
                                if (920 > 0) {
                                    var v39 float32 = (1.5 + 1.5)
                                    _ = v39
                                } else {
                                    _ = len(v1)
                                }
                                var v40 uint16 = (v14 << ((2) & 63))
                                v16 -= 1.5
                                _ = fi38
                                _ = v40
                            }
                            switch "hello" {
                            case "foo bar":
                                var v41 uint64 = (v31 ^ (v31 * (v31 + genFn0(1000, 0, 100))))
                                _ = v41
                            case "hello":
                                var v42 bool = false
                                _ = v42
                            case "a":
                                var v43 int = fi32
                                _ = v43
                            }
                        } else {
                            v17 -= (^(v17))
                            v14 += (v17 >> ((4294967295) & 63))
                        }
                        _ = fi32
                    }
                }
                switch 544 {
                case 0:
                    fmt.Println(v17, v31)
                case 1:
                    for fi44 := 0; fi44 < 6; fi44++ {
                        _ = len(v2)
                        v17 = 1
                        v45 := GenS3{F0: (v31 | (v31 ^ 692))}
                        _ = fi44
                        _ = v45
                    }
                case -9223372036854775808:
                    v46 := GenS3{F0: v31}
                    _ = v46
                default:
                    if false {
                        v16 -= (0.1 / 3.14159)
                    } else {
                        v17 -= (^(v17))
                        v47 := v4[42]
                        _ = v47
                    }
                }
                for fi48 := 0; fi48 < 11; fi48++ {
                    fmt.Println(v16, v17)
                    for fi49 := 0; fi49 < 1; fi49++ {
                        _ = len(v2)
                        fmt.Println(v31, v3)
                        v14 = v14
                        var v50 int64 = 9223372036854775807
                        _ = fi49
                        _ = v50
                    }
                    if (!(1.5 >= 0.1)) {
                        for fi51 := 0; fi51 < 5; fi51++ {
                            var v52 uint16 = 32767
                            v17 = 255
                            if false {
                                switch "hello" {
                                case "":
                                    v3 = (-2.25 * (v3 - (v3 - float64(1e10))))
                                case "foo bar":
                                    fmt.Println(v3)
                                case "a":
                                    v53 := GenS4{F0: (v31 % ((genFn0(42, 65535, 49)) | 1))}
                                    _ = v53
                                default:
                                    v52 += v14
                                }
                            } else {
                                var v54 uint64 = (v31 + (v31 / ((genFn0(-1, 679, 75)) | 1)))
                                var v55 int = -141
                                _ = v54
                                _ = v55
                            }
                            _ = fi51
                            _ = v52
                        }
                    } else {
                        var v56 bool = false
                        var v57 uint8 = 1
                        _ = v56
                        _ = v57
                    }
                    _ = fi48
                }
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
        v58 := v1[2147483647]
        _ = v31
        _ = v58
    }
    v17 -= (v17 * (v14 ^ (^(v14))))
    fmt.Println(v16)
    var v59 int8 = 100
    v59 = v59
    v17 = (v17 * v14)
    var v60 float64 = (v16 - (v3 + (v3 - float64((1.5 - 3.14159)))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v60
}
