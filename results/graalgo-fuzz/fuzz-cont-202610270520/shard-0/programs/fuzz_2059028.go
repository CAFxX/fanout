package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 uint64 }
func genFn0(p0 bool, p1 int64) uint8 {
    v5 := []int{2147483647, 32767, -9223372036854775808, 128}
    p1 += (p1 * p1)
    _ = p0
    _ = p1
    _ = v5
    return 127
}
func genFn1(p0 int64, p1 uint32) uint16 {
    switch 128 {
    case -9223372036854775808:
        var v6 bool = (!true)
        _ = v6
    }
    v7 := []int{625, -9223372036854775808}
    p1 = p1
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    var v1 int8 = 0
    v2 := []string{}
    v3 := make(map[int]bool)
    var v4 string = "a"
    v8 := GenS1{F0: 4294967295}
    if (!false) {
        var v9 bool = true
        _ = v9
    }
    var v10 uint32 = 479
    v4 = "a"
    var v11 float64 = 1.5
    if false {
        v12 := GenS1{F0: (^(v10))}
        fmt.Println(v4)
        v13 := GenS1{F0: 0}
        _ = v12
        _ = v13
    } else {
        var v14 uint16 = genFn1(0, (v10 % ((v10) | 1)))
        v10 = (v10 - 65535)
        _ = v14
    }
    if false {
        for fi15 := 0; fi15 < 11; fi15++ {
            if len(v2) > 0 {
                v2[0] = "a"
            }
            fmt.Println(fi15, v10)
            var v16 int32 = 166
            _ = fi15
            _ = v16
        }
        v11 = v11
        v11 = (1.5 - v11)
    }
    v10 -= (v10 & (v10 << ((18446744073709551615) & 63)))
    for fi17 := 0; fi17 < 20; fi17++ {
        for fi18 := 0; fi18 < 1; fi18++ {
            v19 := GenS1{F0: (^(v10))}
            var v20 int64 = -49
            _ = fi18
            _ = v19
            _ = v20
        }
        if len(v2) > 0 {
            v2[0] = ((v4 + "foo bar") + "hello")
        }
        v21 := GenS1{F0: v10}
        _ = fi17
        _ = v21
    }
    if false {
        var v22 bool = true
        _ = v22
    }
    var v23 int32 = 2147483647
    for fi24 := 0; fi24 < 2; fi24++ {
        _ = len(v2)
        var v25 uint64 = 437
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        if (!true) {
            switch (fi26 / ((fi26) | 1)) {
            case 1:
                v27 := GenS2{F0: 255, F1: 42}
                _ = v27
            case 9223372036854775807:
                var v28 uint = 952
                _ = v28
            default:
                fmt.Println(fi26, v4, v4)
            }
            for fi29 := 0; fi29 < 3; fi29++ {
                switch -1 {
                case 4294967295:
                    fmt.Println(v11, fi29)
                    v10 -= v10
                    v30 := GenS3{F0: 128, F1: (v23 | (v23 * (v23 + -2147483648)))}
                    var v31 string = "\n\t"
                    _ = v30
                    _ = v31
                default:
                    if true {
                        var v32 int64 = 0
                        _ = v32
                    } else {
                        fmt.Println(v11, v10)
                        var v33 bool = ((255 == 4294967295) || (((2 != genFn1(75, (v10 | (v10 | 1)))) || false) || (!(795 >= genFn1(0, (v10 | 26))))))
                        _ = v33
                    }
                }
                v3[1] = false
                switch fi26 {
                case -1:
                    if false {
                        v4 = "a"
                        switch 415 {
                        case -9223372036854775808:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case -682:
                            v23 = (^(v23))
                        case -2147483648:
                            var v34 uint64 = 42
                            _ = v34
                        default:
                            fmt.Println(v4, fi29, fi26)
                        }
                    }
                case 2147483647:
                    var v35 uint64 = 667
                    _ = v35
                default:
                    v10 = (^(v10))
                }
                _ = fi29
            }
            v36 := GenS1{F0: (v10 % ((v10) | 1))}
            _ = v36
        } else {
            for fi37 := 0; fi37 < 12; fi37++ {
                if false {
                    switch v4 {
                    case "hello":
                        v1 = (v1 % (((v1 << ((42) & 63))) | 1))
                    case "a":
                        if ((true || true) || true) {
                            var v38 float32 = 3.14159
                            var v39 float64 = 1.5
                            _ = v38
                            _ = v39
                        } else {
                            v4 = v4
                        }
                    }
                } else {
                    if (false || true) {
                        v40 := GenS3{F0: 128, F1: (v23 - -2147483648)}
                        _ = v40
                    } else {
                        fmt.Println(fi37)
                        v41 := GenS4{F0: v1}
                        switch fi26 {
                        case 312:
                            if (true && true) {
                                v10 += (v10 << ((42) & 63))
                                fmt.Println(fi26, fi26, v11)
                            } else {
                                var v42 int = fi37
                                if len(v2) > 0 {
                                    v2[0] = v4
                                }
                                _ = v42
                            }
                        case 42:
                            switch v4 {
                            case "":
                                if (65535 > 18446744073709551615) {
                                    v1 = (-(v1))
                                    v11 = v11
                                }
                            case "a":
                                v10 += (v10 & v10)
                            case "foo bar":
                                v23 = (-(v23))
                            }
                        default:
                            var v43 int8 = -39
                            _ = v43
                        }
                        var v44 float64 = v11
                        var v45 int32 = v23
                        _ = v41
                        _ = v44
                        _ = v45
                    }
                    v46 := GenS3{F0: 129, F1: 913}
                    switch 42 {
                    case -2147483648:
                        v2 = append(v2, v4)
                    case 528:
                        v1 += (^(v1))
                    case 9223372036854775807:
                        v2 = append(v2, ("foo bar" + v4))
                    }
                    _ = v46
                }
                var v47 uint16 = 32767
                _ = fi37
                _ = v47
            }
            v48 := GenS5{F0: 0, F1: 2}
            var v49 uint = 280
            _ = v48
            _ = v49
        }
        _ = fi26
    }
    var v50 uint8 = 2
    v51 := GenS4{F0: (v1 << ((42) & 63))}
    v52 := v3[127]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v50
    _ = v51
    _ = v52
    _ = v8
}
