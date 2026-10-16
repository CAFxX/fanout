package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 int16 }
func genFn0(p0 int16, p1 bool) string {
    for fi5 := 0; fi5 < 1; fi5++ {
        for fi6 := 0; fi6 < 18; fi6++ {
            var v7 float32 = 0.0
            v7 += v7
            fmt.Println(p0)
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 uint32 = 4294967295
        var v10 int64 = 200
        if p1 {
            switch "\n\t" {
            case "foo bar":
                if false {
                    fmt.Println(fi8, v9, v10)
                    if p1 {
                        var v11 uint64 = 651
                        switch ("a" + "") {
                        case "hello":
                            var v12 int16 = (p0 * (p0 % ((p0) | 1)))
                            _ = v12
                        case "foo bar":
                            v13 := []int{fi8, (fi8 % ((fi8) | 1)), 778}
                            _ = v13
                        default:
                            var v14 int16 = (p0 >> ((4294967295) & 63))
                            _ = v14
                        }
                        var v15 string = "a"
                        _ = v11
                        _ = v15
                    }
                    if p1 {
                        v16 := []int{fi8, fi8, (fi8 + fi8)}
                        _ = v16
                    }
                } else {
                    v17 := []int{}
                    var v18 int64 = (v10 * v10)
                    if p1 {
                        for fi19 := 0; fi19 < 6; fi19++ {
                            var v20 bool = p1
                            if (v9 != (^(v9))) {
                                v21 := GenS1{F0: 2, F1: 2147483647}
                                var v22 uint8 = 1
                                if false {
                                    v22 -= uint8(v9)
                                    p0 = 42
                                } else {
                                    var v23 int = fi19
                                    var v24 bool = v20
                                    _ = v23
                                    _ = v24
                                }
                                p1 = false
                                if (false && (p1 || ((fi19 / (((fi19 << ((18446744073709551615) & 63))) | 1)) >= (-(fi8))))) {
                                    var v25 uint = 65535
                                    v10 = int64(v9)
                                    _ = v25
                                }
                                _ = v21
                                _ = v22
                            }
                            var v26 float64 = 1e10
                            _ = fi19
                            _ = v20
                            _ = v26
                        }
                        switch fi8 {
                        case 2147483647:
                            for fi27 := 0; fi27 < 3; fi27++ {
                                if ((p0 % (((p0 | (p0 >> ((2) & 63)))) | 1)) == 0) {
                                    p1 = (p1 && p1)
                                    v28 := GenS1{F0: 255, F1: int32(v10)}
                                    if len(v17) >= 2 {
                                        v17 = v17[1:len(v17)]
                                    }
                                    _ = v28
                                } else {
                                    switch -493 {
                                    case -675:
                                        fmt.Println(v10, p1, p0)
                                        var v29 int = 9223372036854775807
                                        _ = v29
                                    }
                                    fmt.Println(fi8)
                                }
                                if len(v17) > 0 {
                                    v17[0] = (-(fi27))
                                }
                                var v30 int8 = -37
                                _ = fi27
                                _ = v30
                            }
                        default:
                            var v31 int32 = 1
                            _ = v31
                        }
                        p0 = (p0 - (p0 / (((p0 | -735)) | 1)))
                    }
                    _ = v17
                    _ = v18
                }
            case "":
                for fi32 := 0; fi32 < 18; fi32++ {
                    v33 := []int{(-(fi8)), (^(fi8)), (fi8 | fi32)}
                    _ = fi32
                    _ = v33
                }
            }
            fmt.Println(v9, fi8)
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v34 := GenS1{F0: 0, F1: 0}
    var v35 float32 = (float32(3.14159) + ((0.1 / -2.25) - 0.1))
    _ = p0
    _ = p1
    _ = v34
    _ = v35
    return "foo bar"
}

func main() {
    var v1 uint16 = 65535
    var v2 uint8 = 1
    var v3 int64 = 390
    v4 := make(map[string]uint32)
    v3 = (v3 >> ((455) & 63))
    var v36 int8 = 42
    v37 := GenS1{F0: (v1 % (((^(v1))) | 1)), F1: 2147483647}
    if (0 >= 2147483648) {
        var v38 int64 = v3
        v2 -= (^(v2))
        _ = v38
    }
    fmt.Println(v36)
    switch int(v3) {
    case -2147483648:
        v39 := GenS2{F0: 1}
        _ = v39
    case 32767:
        var v40 float64 = 1.5
        _ = v40
    case -9223372036854775808:
        v3 = v3
    default:
        v3 = int64(v2)
    }
    v2 = (v2 - v2)
    if false {
        for fi41 := 0; fi41 < 2; fi41++ {
            fmt.Println(v3, v36)
            v42 := GenS1{F0: 2, F1: -1}
            var v43 int32 = -425
            _ = fi41
            _ = v42
            _ = v43
        }
        v44 := GenS3{F0: 2, F1: 1000}
        _ = v44
    }
    var v45 float64 = 1e10
    v46 := GenS2{F0: v1}
    var v47 int32 = -940
    for fi48 := 0; fi48 < 9; fi48++ {
        v49 := []int{(-(fi48)), 1}
        _ = fi48
        _ = v49
    }
    var v50 bool = (true || (true && false))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v50
}
