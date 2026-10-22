package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int16; F1 int; F2 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 int64; F2 uint64 }
type GenS5 struct { F0 int16; F1 int; F2 uint32 }
func genFn0(p0 int) int32 {
    fmt.Println(p0)
    fmt.Println(p0)
    fmt.Println(p0)
    v5 := []int{(p0 << ((42) & 63)), (-(p0))}
    var v6 int16 = -959
    _ = p0
    _ = v5
    _ = v6
    return -2147483648
}

func main() {
    v1 := []uint32{79, 42, 440, 4294967295}
    var v2 uint32 = 65535
    v3 := []int{-301, 32768, 42, 32767}
    var v4 uint16 = uint16(v2)
    v3 = append(v3, 32768)
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(v4, v4, v4)
        _ = fi7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        v4 -= v4
        v9 := GenS1{F0: -1, F1: 100}
        switch fi8 {
        case -945:
            var v10 int16 = 269
            _ = v10
        case -2147483648:
            v1 = append(v1, (v2 & v2))
        case -1:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            if len(v1) > 0 {
                v1[0] = v2
            }
        }
        _ = fi8
        _ = v9
    }
    if false {
        var v11 float32 = 1.5
        _ = len(v1)
        _ = v11
    } else {
        for fi12 := 0; fi12 < 10; fi12++ {
            v4 = uint16(v2)
            _ = fi12
        }
    }
    if (18446744073709551615 != 4294967295) {
        for fi13 := 0; fi13 < 16; fi13++ {
            if true {
                v14 := GenS2{F0: 0, F1: fi13, F2: 906}
                v15 := GenS1{F0: 226, F1: 127}
                _ = v14
                _ = v15
            }
            _ = fi13
        }
        if (-9223372036854775808 > 939) {
            v16 := GenS2{F0: 926, F1: -286, F2: 2147483647}
            switch "" {
            case "":
                var v17 float64 = ((((1e10 - 3.14159) + -2.25) / 3.14159) - (1e10 - 3.14159))
                _ = v17
            default:
                var v18 int16 = 1
                _ = v18
            }
            for fi19 := 0; fi19 < 20; fi19++ {
                var v20 float32 = -2.25
                _ = fi19
                _ = v20
            }
            _ = v16
        }
    }
    var v21 float64 = -2.25
    var v22 int32 = genFn0(-967)
    for fi23 := 0; fi23 < 20; fi23++ {
        if true {
            if (!false) {
                switch fi23 {
                case 2147483647:
                    for fi24 := 0; fi24 < 7; fi24++ {
                        switch (fi23 << ((984) & 63)) {
                        case 504:
                            var v25 uint = 4294967295
                            _ = v25
                        case 42:
                            switch (-(fi23)) {
                            case 245:
                                v2 -= v2
                            }
                        default:
                            switch "" {
                            case "foo bar":
                                v1 = append(v1, 2147483648)
                                if false {
                                    var v26 uint32 = v2
                                    _ = v26
                                }
                            case "a":
                                if len(v3) > 0 {
                                    v3[0] = fi23
                                }
                            case "hello":
                                v22 -= -2147483648
                            }
                        }
                        _ = fi24
                    }
                case 127:
                    var v27 string = ("foo bar" + "foo bar")
                    _ = v27
                }
            } else {
                if false {
                    if true {
                        v28 := GenS2{F0: 32767, F1: 4294967295, F2: (-(v22))}
                        _ = v28
                    } else {
                        if true {
                            for fi29 := 0; fi29 < 20; fi29++ {
                                fmt.Println(v21)
                                v30 := GenS3{F0: v22}
                                _ = fi29
                                _ = v30
                            }
                            fmt.Println(v2, fi23)
                            v21 = (v21 * (v21 - 0.1))
                        } else {
                            v31 := GenS1{F0: 157, F1: -39}
                            for fi32 := 0; fi32 < 1; fi32++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v33 string = "a"
                                v3 = append(v3, (fi23 | 1))
                                if (!(true || false)) {
                                    var v34 uint = 42
                                    if len(v3) > 0 {
                                        v3[0] = fi23
                                    }
                                    var v35 int = (fi32 - fi32)
                                    _ = v34
                                    _ = v35
                                } else {
                                    if false {
                                        switch (((v33 + (v33 + (v33 + (v33 + "\n\t")))) + v33) + "") {
                                        case "foo bar":
                                            fmt.Println(fi23, v21)
                                        }
                                    }
                                    var v36 string = ((((v33 + v33) + v33) + v33) + (v33 + v33))
                                    var v37 int = 9223372036854775807
                                    _ = v36
                                    _ = v37
                                }
                                fmt.Println(v21)
                                _ = fi32
                                _ = v33
                            }
                            _ = v31
                        }
                    }
                    v38 := GenS3{F0: genFn0((^(fi23)))}
                    _ = v38
                }
            }
            if true {
                for fi39 := 0; fi39 < 14; fi39++ {
                    v40 := GenS4{F0: 0, F1: 42, F2: 42}
                    var v41 int8 = 1
                    if (false || (fi39 >= -9223372036854775808)) {
                        v42 := GenS1{F0: -9223372036854775808, F1: (v41 - v41)}
                        if (v2 != (^(v2))) {
                            var v43 int16 = -654
                            _ = v43
                        }
                        var v44 uint16 = v4
                        _ = v42
                        _ = v44
                    }
                    _ = fi39
                    _ = v40
                    _ = v41
                }
                _ = len(v1)
            } else {
                v45 := GenS2{F0: 0, F1: -1, F2: (v22 << ((18446744073709551615) & 63))}
                v3 = append(v3, 2147483647)
                _ = v45
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            v46 := GenS5{F0: -246, F1: (fi23 ^ fi23), F2: (v2 ^ v2)}
            _ = len(v3)
            var v47 int64 = -1
            _ = v46
            _ = v47
        }
        _ = fi23
    }
    var v48 bool = (!true)
    switch "foo bar" {
    case "a":
        v2 += (^(v2))
    }
    var v49 bool = false
    if len(v3) > 0 {
        v3[0] = 542
    }
    var v50 uint8 = 147
    for fi51 := 0; fi51 < 15; fi51++ {
        v21 = v21
        var v52 float32 = 0.1
        var v53 float32 = float32(1e10)
        _ = fi51
        _ = v52
        _ = v53
    }
    fmt.Println(v4, v21, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v50
}
