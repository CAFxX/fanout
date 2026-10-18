package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int16; F1 bool; F2 int8 }
type GenS3 struct { F0 int8; F1 uint8; F2 int32 }
func genFn0(p0 int16, p1 uint, p2 uint16) uint64 {
    var v4 int16 = (p0 / ((1) | 1))
    var v5 uint16 = 0
    var v6 int32 = 42
    v7 := GenS1{F0: -1, F1: 799, F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1() int16 {
    for fi8 := 0; fi8 < 14; fi8++ {
        fmt.Println(fi8)
        var v9 int64 = 1
        if true {
            for fi10 := 0; fi10 < 16; fi10++ {
                v11 := []int{fi8, (fi10 >> ((42) & 63)), fi10}
                _ = fi10
                _ = v11
            }
            if true {
                v12 := []int{fi8, 32767}
                _ = v12
            } else {
                if (2 < 4294967295) {
                    v9 -= (v9 + v9)
                } else {
                    fmt.Println(v9, v9)
                }
            }
        } else {
            if false {
                var v13 uint16 = 1
                if (!false) {
                    if (127 < 105) {
                        for fi14 := 0; fi14 < 17; fi14++ {
                            v13 += v13
                            v15 := GenS1{F0: v9, F1: 159, F2: 85}
                            var v16 int64 = v9
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    }
                }
                _ = v13
            }
        }
        _ = fi8
        _ = v9
    }
    v17 := []int{128, 127, -119, 2147483647}
    var v18 float64 = ((3.14159 / float64((3.14159 * 0.1))) * 1.5)
    if true {
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 uint64 = genFn0(-1, 18446744073709551615, 32767)
            var v21 bool = (("\n\t" == "hello") && false)
            switch (-(fi19)) {
            case 752:
                if (242 <= 16) {
                    v20 -= v20
                    for fi22 := 0; fi22 < 17; fi22++ {
                        var v23 uint = 0
                        v21 = (18446744073709551615 >= (v23 - 0))
                        _ = fi22
                        _ = v23
                    }
                } else {
                    v18 = float64(float32(v18))
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    if (uint(v20) >= 478) {
                        var v24 int64 = 9223372036854775807
                        if false {
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                        } else {
                            v25 := GenS2{F0: 1, F1: true, F2: -23}
                            _ = v25
                        }
                        _ = v24
                    }
                }
            }
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    _ = v17
    _ = v18
    return 1
}

func main() {
    var v1 uint32 = 42
    v2 := make(map[string]string)
    var v3 int16 = 32767
    v1 = 307
    v3 = v3
    v3 += (v3 * v3)
    for fi26 := 0; fi26 < 12; fi26++ {
        switch (fi26 - (fi26 & 2147483647)) {
        case 2147483647:
            var v27 uint16 = 8
            _ = v27
        case -2147483648:
            fmt.Println(v1, fi26)
        case -1:
            var v28 int16 = 1000
            _ = v28
        }
        if (false && (42 <= 2)) {
            v1 -= v1
            v29 := GenS1{F0: -1, F1: 4294967295, F2: 4294967295}
            _ = v29
        }
        _ = fi26
    }
    v2["hello"] = "hello"
    v30 := GenS1{F0: -220, F1: 2, F2: v1}
    v3 -= genFn1()
    var v31 uint8 = 182
    var v32 uint64 = 261
    if (!(115 != (^(v32)))) {
        for fi33 := 0; fi33 < 19; fi33++ {
            for fi34 := 0; fi34 < 9; fi34++ {
                v35 := GenS3{F0: -128, F1: (v31 ^ (v31 + (v31 & 50))), F2: 1}
                var v36 int64 = 9223372036854775807
                v36 = (-(v36))
                _ = fi34
                _ = v35
                _ = v36
            }
            v2[("a" + "")] = ""
            _ = fi33
        }
        var v37 float64 = 0.1
        _ = v37
    } else {
        for fi38 := 0; fi38 < 9; fi38++ {
            v31 = (v31 << ((0) & 63))
            if (!false) {
                fmt.Println(v31)
            }
            _ = fi38
        }
        v1 += (v1 | 2147483647)
        for fi39 := 0; fi39 < 14; fi39++ {
            var v40 int16 = -32768
            for fi41 := 0; fi41 < 20; fi41++ {
                var v42 int8 = 96
                _ = fi41
                _ = v42
            }
            switch fi39 {
            case 9223372036854775807:
                v43 := GenS3{F0: 39, F1: 146, F2: int32(v1)}
                _ = v43
            default:
                for fi44 := 0; fi44 < 20; fi44++ {
                    var v45 int8 = 112
                    v45 = (^(v45))
                    if false {
                        v32 -= v32
                    } else {
                        var v46 float64 = -2.25
                        v45 += 106
                        fmt.Println(v31, v32, v45)
                        _ = v46
                    }
                    _ = fi44
                    _ = v45
                }
            }
            _ = fi39
            _ = v40
        }
    }
    switch "a" {
    case "foo bar":
        var v47 uint = 42
        _ = v47
    case "\n\t":
        v48 := []int{}
        _ = v48
    default:
        v3 = v3
    }
    var v49 uint8 = v31
    v49 = uint8(v32)
    if (false && true) {
        var v50 int32 = 525
        var v51 int = -2147483648
        var v52 int8 = -84
        _ = v50
        _ = v51
        _ = v52
    } else {
        switch -412 {
        case 134:
            v53 := GenS2{F0: v3, F1: ((v3 + 49) > (v3 - (v3 * 866))), F2: 100}
            _ = v53
        }
    }
    fmt.Println(v3, v3)
    v54 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v54:", len(v54), v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v49
    _ = v54
}
