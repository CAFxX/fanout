package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 string; F1 int32; F2 uint8 }
type GenS3 struct { F0 int64; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 string; F2 bool }
func genFn0(p0 uint8, p1 bool) int64 {
    var v4 uint16 = 255
    v5 := GenS1{F0: -404, F1: 2}
    v6 := GenS2{F0: "\n\t", F1: 2147483647, F2: p0}
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 float32 = float32(1.5)
        var v9 uint16 = (v4 << ((878) & 63))
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint = 705
    _ = p0
    _ = p1
    _ = v10
    _ = v4
    _ = v5
    _ = v6
    return -1
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    var v11 int8 = 42
    switch 609 {
    case 32767:
        v12 := GenS3{F0: 26, F1: -47, F2: 217}
        _ = v12
    }
    v13 := v3[2147483647]
    for fi14 := 0; fi14 < 3; fi14++ {
        v11 = v11
        var v15 int32 = 0
        v16 := []int{(fi14 << ((1) & 63)), (fi14 & (fi14 / (((fi14 + (fi14 + (^(fi14))))) | 1)))}
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 uint = 1
    var v18 string = "hello"
    if ((255 < 1) && false) {
        for fi19 := 0; fi19 < 20; fi19++ {
            var v20 bool = (!true)
            if false {
                v21 := GenS3{F0: 0, F1: v11, F2: 0}
                v20 = v20
                _ = v21
            } else {
                for fi22 := 0; fi22 < 15; fi22++ {
                    if (0.0 >= 1e10) {
                        var v23 int64 = 863
                        v1[(fi19 | 9223372036854775807)] = (0 != 793)
                        var v24 uint64 = 18446744073709551615
                        _ = v23
                        _ = v24
                    } else {
                        if v20 {
                            var v25 uint64 = uint64(v13)
                            _ = v25
                        } else {
                            v18 = v18
                            switch 127 {
                            case 32767:
                                v26 := []int{fi19}
                                _ = v26
                            case -254:
                                var v27 int64 = 0
                                _ = v27
                            default:
                                if true {
                                    var v28 uint16 = 265
                                    v18 = ""
                                    _ = v28
                                } else {
                                    if (((9223372036854775807 > 0) || (1 == v11)) && (232 <= 65535)) {
                                        v20 = ((^(fi22)) > (fi22 - -452))
                                        v18 = (("hello" + v18) + v18)
                                    }
                                    v29 := v1[(fi22 / ((-9223372036854775808) | 1))]
                                    _ = v29
                                }
                            }
                            v20 = (-32768 > 32767)
                        }
                        var v30 int32 = -2147483648
                        _ = v30
                    }
                    _ = fi22
                }
            }
            switch fi19 {
            case 128:
                v31 := GenS3{F0: 169, F1: v11, F2: 918}
                _ = v31
            }
            _ = fi19
            _ = v20
        }
        fmt.Println(v11, v17)
    } else {
        var v32 int = -196
        _ = v32
    }
    for fi33 := 0; fi33 < 2; fi33++ {
        v17 += (v17 | 18446744073709551615)
        _ = fi33
    }
    v34 := []int{128}
    v11 -= (v11 >> (((v17 >> ((280) & 63))) & 63))
    switch v18 {
    case "hello":
        var v35 int64 = -177
        _ = v35
    case "a":
        if true {
            var v36 uint8 = 239
            _ = v36
        } else {
            fmt.Println(v13, v11, v18)
        }
    case "foo bar":
        switch 1 {
        case 772:
            var v37 uint64 = 2
            _ = v37
        case -2147483648:
            v11 += v11
        default:
            var v38 int8 = (v11 * -1)
            _ = v38
        }
    default:
        switch 2147483647 {
        case 32768:
            var v39 uint32 = v13
            _ = v39
        case 799:
            v13 += v13
        case 127:
            v40 := GenS4{F0: 42}
            _ = v40
        }
    }
    if true {
        switch v18 {
        case "":
            for fi41 := 0; fi41 < 1; fi41++ {
                v42 := v2[v18]
                _ = fi41
                _ = v42
            }
        case "\n\t":
            for fi43 := 0; fi43 < 5; fi43++ {
                if (true || (false || (18446744073709551615 > 18446744073709551615))) {
                    v13 -= v13
                    fmt.Println(fi43, v18)
                }
                if true {
                    v11 = v11
                    var v44 uint32 = v13
                    _ = v44
                }
                var v45 uint16 = 827
                _ = fi43
                _ = v45
            }
        case "hello":
            v34 = append(v34, 2147483647)
        }
        for fi46 := 0; fi46 < 8; fi46++ {
            v47 := GenS4{F0: (v17 & 65535)}
            _ = fi46
            _ = v47
        }
    }
    v18 = (v18 + "foo bar")
    delete(v1, -478)
    var v48 uint32 = (v13 / (((v13 / (((v13 / ((0) | 1))) | 1))) | 1))
    v49 := GenS4{F0: 0}
    for fi50 := 0; fi50 < 8; fi50++ {
        if true {
            for fi51 := 0; fi51 < 11; fi51++ {
                _ = len(v34)
                _ = fi51
            }
        } else {
            delete(v2, v18)
        }
        _ = fi50
    }
    v52 := GenS5{F0: (!true), F1: ("\n\t" + "\n\t"), F2: (false && (3.14159 != 0.0))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v11
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v34
    _ = v48
    _ = v49
    _ = v52
}
