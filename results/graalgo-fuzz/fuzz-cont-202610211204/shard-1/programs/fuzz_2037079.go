package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 uint8; F1 int }
func genFn0(p0 int8, p1 int16) uint16 {
    p0 = (p0 << ((255) & 63))
    var v5 int64 = -1
    var v6 float32 = 0.1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 float32 = 1e10
    var v2 uint16 = 32767
    v3 := make(map[string]bool)
    var v4 bool = false
    if v4 {
        v2 += genFn0(-128, 0)
    } else {
        for fi7 := 0; fi7 < 8; fi7++ {
            switch "a" {
            case "\n\t":
                switch (fi7 - 127) {
                case -9223372036854775808:
                    for fi8 := 0; fi8 < 6; fi8++ {
                        v4 = (v4 || ("foo bar" <= ("" + (("foo bar" + ("a" + "foo bar")) + ("a" + "foo bar")))))
                        var v9 uint8 = 1
                        v1 -= (v1 / 1e10)
                        _ = fi8
                        _ = v9
                    }
                }
            case "foo bar":
                if v4 {
                    var v10 int8 = 42
                    _ = v10
                } else {
                    if (fi7 != -1) {
                        if v4 {
                            switch (fi7 - 4294967295) {
                            case 32768:
                                v4 = (v4 || (1e10 >= float64(v1)))
                            default:
                                var v11 int32 = -557
                                _ = v11
                            }
                            if (65535 >= 255) {
                                switch (fi7 >> ((18446744073709551615) & 63)) {
                                case 2147483647:
                                    for fi12 := 0; fi12 < 11; fi12++ {
                                        var v13 int16 = -32768
                                        fmt.Println(fi12, v1)
                                        v14 := GenS1{F0: (^(v13)), F1: fi12}
                                        _ = fi12
                                        _ = v13
                                        _ = v14
                                    }
                                case 191:
                                    v1 -= v1
                                case 4294967295:
                                    if true {
                                        switch fi7 {
                                        case -1:
                                            if (518 == -2147483648) {
                                                var v15 float64 = 1e10
                                                if v4 {
                                                    for fi16 := 0; fi16 < 17; fi16++ {
                                                        if true {
                                                            var v17 int16 = -32768
                                                            v17 = (v17 / ((v17) | 1))
                                                            _ = v17
                                                        }
                                                        _ = fi16
                                                    }
                                                }
                                                _ = v15
                                            }
                                        case 0:
                                            v18 := []int{fi7, (fi7 << ((643) & 63))}
                                            _ = v18
                                        case 128:
                                            var v19 uint16 = genFn0(100, -1)
                                            _ = v19
                                        }
                                        v20 := GenS1{F0: -403, F1: (fi7 % (((fi7 >> ((uint(v2)) & 63))) | 1))}
                                        _ = v20
                                    } else {
                                        v2 = (v2 & genFn0(127, -986))
                                        if (!false) {
                                            v2 += (v2 + v2)
                                            v2 = (^(v2))
                                            var v21 int32 = -2147483648
                                            _ = v21
                                        }
                                        v22 := GenS1{F0: 1, F1: (-(fi7))}
                                        _ = v22
                                    }
                                default:
                                    fmt.Println(v1, v2, v1)
                                }
                                var v23 float32 = float32(-2.25)
                                _ = v23
                            } else {
                                var v24 uint8 = 2
                                _ = v24
                            }
                        }
                        delete(v3, "a")
                        v2 = (v2 % (((v2 << ((65535) & 63))) | 1))
                    } else {
                        v25 := []int{(-(fi7)), (fi7 % (((^(fi7))) | 1)), 32768}
                        _ = v25
                    }
                    fmt.Println(v1, v4, fi7)
                }
            case "":
                v4 = true
            default:
                if v4 {
                    v26 := []int{fi7, fi7, (fi7 | fi7)}
                    var v27 uint = 42
                    v28 := GenS1{F0: -1, F1: (fi7 * 859)}
                    _ = v26
                    _ = v27
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 6; fi29++ {
                        var v30 bool = v4
                        var v31 uint32 = 65535
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                }
            }
            _ = fi7
        }
        var v32 int16 = 504
        _ = v32
    }
    v33 := GenS1{F0: 0, F1: 127}
    switch 128 {
    case 32768:
        v4 = v4
    case -19:
        var v34 uint = 65535
        _ = v34
    case 2147483647:
        v4 = v4
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        fmt.Println(fi35, v2, fi35)
        _ = fi35
    }
    var v36 int16 = 1
    var v37 float64 = -2.25
    switch 32768 {
    case -2147483648:
        var v38 float64 = v37
        _ = v38
    case 2147483647:
        v39 := GenS2{F0: 226, F1: 123}
        _ = v39
    case -733:
        fmt.Println(v37, v4)
    }
    v40 := GenS3{F0: 0, F1: 9223372036854775807}
    v41 := GenS1{F0: v36, F1: 127}
    if v4 {
        v42 := GenS1{F0: (v36 & -32768), F1: 0}
        _ = v42
    } else {
        v2 += v2
        v36 -= 1000
        for fi43 := 0; fi43 < 18; fi43++ {
            if ((((false || v4) || (!(255 == 127))) && (4294967295 <= 18446744073709551615)) || (v1 == ((-2.25 + 0.0) / v1))) {
                var v44 bool = (!false)
                _ = v44
            } else {
                if (42 >= 2147483648) {
                    var v45 int64 = 42
                    _ = v45
                }
            }
            _ = fi43
        }
    }
    if v4 {
        fmt.Println(v2)
        v4 = (!(true || false))
        var v46 string = ("a" + "hello")
        _ = v46
    }
    v47 := []int{9223372036854775807, 417, 9223372036854775807}
    v48 := GenS3{F0: 127, F1: 127}
    var v49 int = 9223372036854775807
    v37 = ((1e10 / v37) - v37)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v41
    _ = v47
    _ = v48
    _ = v49
}
