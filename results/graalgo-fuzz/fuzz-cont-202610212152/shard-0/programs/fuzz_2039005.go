package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32; F1 int; F2 int8 }
func genFn0(p0 uint8, p1 bool) uint {
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := []int{fi7, fi7, (-(fi7)), -1}
        v9 := GenS1{F0: -57}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{0}
    switch "foo bar" {
    case "\n\t":
        switch 0 {
        case -9223372036854775808:
            v11 := GenS2{F0: 255}
            _ = v11
        case 0:
            var v12 uint32 = 2
            _ = v12
        default:
            v13 := GenS3{F0: 42, F1: 42}
            _ = v13
        }
    default:
        var v14 uint32 = 42
        _ = v14
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 int8 = -128
        var v17 uint = 18446744073709551615
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v10
    return 891
}
func genFn1(p0 int64, p1 int64) uint {
    fmt.Println(p0, p1)
    switch -9223372036854775808 {
    case 160:
        var v18 uint32 = 1
        _ = v18
    case 1:
        var v19 uint8 = 119
        _ = v19
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 919
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]string)
    v3 := []int{1, 2147483647, 127, 344}
    var v4 uint32 = 926
    v5 := make(map[uint32]uint32)
    var v6 float64 = -2.25
    v2[42] = (("\n\t" + "hello") + "a")
    v20 := GenS1{F0: (^(v1))}
    _ = len(v2)
    fmt.Println(v4, v6, v6)
    switch -555 {
    case 2147483647:
        var v21 float64 = 1e10
        _ = v21
    }
    if len(v3) > 0 {
        v3[0] = 128
    }
    var v22 string = ""
    var v23 int32 = v1
    v2[9223372036854775807] = ("" + v22)
    for fi24 := 0; fi24 < 9; fi24++ {
        fmt.Println(v4, v23)
        _ = fi24
    }
    if true {
        for fi25 := 0; fi25 < 5; fi25++ {
            for fi26 := 0; fi26 < 17; fi26++ {
                v23 = (v23 << ((1) & 63))
                if false {
                    v4 = v4
                } else {
                    switch fi25 {
                    case 1:
                        if (-2.25 != 1e10) {
                            switch (v22 + (v22 + ("a" + "a"))) {
                            case "":
                                var v27 string = (v22 + (v22 + "foo bar"))
                                _ = v27
                            case "\n\t":
                                _ = len(v2)
                            case "foo bar":
                                delete(v5, 0)
                            default:
                                fmt.Println(v22)
                            }
                            if (0 == 1) {
                                v28 := GenS4{F0: 275}
                                v29 := GenS4{F0: 0}
                                _ = v28
                                _ = v29
                            }
                            var v30 int64 = 1
                            delete(v5, v4)
                            if false {
                                if false {
                                    var v31 uint16 = 32767
                                    _ = v31
                                }
                            } else {
                                v3 = append(v3, (^(fi26)))
                                v22 = v22
                                _ = len(v3)
                            }
                            v32 := GenS5{F0: v23, F1: 32767, F2: 127}
                            switch ((v22 + ((v22 + v22) + "foo bar")) + v22) {
                            case "foo bar":
                                if false {
                                    if (-2.25 > (v6 * (v6 * -2.25))) {
                                        var v33 int32 = (v1 + (v1 % (((v1 - (v1 & v1))) | 1)))
                                        v34 := GenS4{F0: genFn0(42, false)}
                                        v22 = "hello"
                                        _ = v33
                                        _ = v34
                                    }
                                    var v35 int8 = -1
                                    _ = v35
                                } else {
                                    fmt.Println(v23)
                                    fmt.Println(v4, fi25, fi26)
                                    if false {
                                        var v36 uint32 = (v4 << ((0) & 63))
                                        v37 := GenS1{F0: v23}
                                        v38 := v5[1]
                                        var v39 int = 179
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                    }
                                }
                            case "":
                                v6 += v6
                            default:
                                fmt.Println(v30, v1)
                            }
                            if false {
                                var v40 float64 = (v6 + v6)
                                var v41 uint32 = (v4 + v4)
                                _ = v40
                                _ = v41
                            }
                            _ = v30
                            _ = v32
                        } else {
                            var v42 uint32 = 2147483648
                            v6 += v6
                            _ = v42
                        }
                    case 42:
                        v43 := GenS1{F0: 42}
                        _ = v43
                    case 128:
                        v23 -= v23
                    }
                }
                var v44 uint = 4294967295
                _ = fi26
                _ = v44
            }
            for fi45 := 0; fi45 < 8; fi45++ {
                var v46 int = -427
                _ = fi45
                _ = v46
            }
            _ = fi25
        }
        for fi47 := 0; fi47 < 1; fi47++ {
            fmt.Println(v23)
            _ = fi47
        }
    }
    for fi48 := 0; fi48 < 13; fi48++ {
        var v49 int8 = 1
        _ = fi48
        _ = v49
    }
    v22 = "\n\t"
    var v50 uint32 = 2147483648
    v50 -= v50
    v4 = (v4 / (((^(v50))) | 1))
    switch 32767 {
    case 1:
        fmt.Println(v1)
    default:
        fmt.Println(v4, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
}
