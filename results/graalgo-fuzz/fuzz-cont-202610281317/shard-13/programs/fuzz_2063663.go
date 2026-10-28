package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 bool }
type GenS2 struct { F0 int; F1 int32; F2 uint8 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 int8; F1 string; F2 uint8 }
func genFn0(p0 uint64) int64 {
    v4 := []int{-2147483648, -2147483648, 925, -9223372036854775808}
    v5 := GenS1{F0: 9223372036854775807, F1: 18446744073709551615, F2: (true && (false && true))}
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            v4 = append(v4, fi6)
            p0 = (p0 % (((p0 >> ((42) & 63))) | 1))
            _ = fi7
        }
        _ = fi6
    }
    p0 = p0
    p0 += (p0 << ((255) & 63))
    _ = p0
    _ = v4
    _ = v5
    return 9223372036854775807
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint16 = 1
    var v3 int = -2147483648
    if false {
        if false {
            for fi8 := 0; fi8 < 17; fi8++ {
                v1 = v1
                switch -515 {
                case 1:
                    var v9 int8 = -128
                    _ = v9
                case 32767:
                    v1 = v1
                default:
                    var v10 int64 = 9223372036854775807
                    _ = v10
                }
                _ = fi8
            }
            for fi11 := 0; fi11 < 20; fi11++ {
                fmt.Println(v1)
                var v12 uint8 = 128
                v12 -= 1
                _ = fi11
                _ = v12
            }
        }
        fmt.Println(v2)
    } else {
        fmt.Println(v1, v3)
    }
    var v13 int = v3
    if true {
        var v14 int64 = 426
        if (!(v2 > v2)) {
            var v15 uint32 = 4294967295
            v16 := []int{v3}
            fmt.Println(v2)
            _ = v15
            _ = v16
        } else {
            v17 := GenS1{F0: genFn0((v1 >> ((363) & 63))), F1: 65535, F2: true}
            _ = v17
        }
        _ = v14
    } else {
        v18 := GenS2{F0: (v3 - 127), F1: -2147483648, F2: 0}
        for fi19 := 0; fi19 < 7; fi19++ {
            v13 -= v3
            v20 := GenS3{F0: 127, F1: 1}
            if (18446744073709551615 <= 969) {
                for fi21 := 0; fi21 < 14; fi21++ {
                    v3 = v3
                    _ = fi21
                }
                if (!true) {
                    for fi22 := 0; fi22 < 4; fi22++ {
                        v13 -= (fi22 | (v3 % ((-736) | 1)))
                        v13 += (fi19 % ((fi19) | 1))
                        if false {
                            var v23 uint8 = 42
                            v24 := []int{(v3 - 32767)}
                            v23 -= (v23 / ((42) | 1))
                            _ = v23
                            _ = v24
                        } else {
                            fmt.Println(fi19, v1, v13)
                            var v25 int32 = 42
                            if (false || (v2 < 31)) {
                                var v26 uint16 = v2
                                v1 = 470
                                fmt.Println(v1)
                                _ = v26
                            }
                            _ = v25
                        }
                        _ = fi22
                    }
                } else {
                    var v27 int64 = 42
                    _ = v27
                }
                v3 = -868
            }
            _ = fi19
            _ = v20
        }
        _ = v18
    }
    var v28 bool = (1 <= 42)
    v29 := []int{(v3 | (v3 >> ((1) & 63))), v3}
    v29 = append(v29, 1)
    v2 += v2
    var v30 bool = (false || (!("foo bar" <= "a")))
    fmt.Println(v1)
    if (((float64(3.14159) - (-2.25 * 1e10)) > (0.0 - 1.5)) && (!("" <= ("a" + "foo bar")))) {
        v31 := GenS4{F0: 1, F1: "\n\t", F2: 89}
        _ = v31
    }
    v32 := GenS1{F0: 9223372036854775807, F1: 896, F2: true}
    v3 = (v3 << ((2) & 63))
    v3 = (v13 % ((619) | 1))
    var v33 uint16 = v2
    if v28 {
        switch "a" {
        case "hello":
            if len(v29) >= 2 {
                v29 = v29[1:len(v29)]
            }
        case "foo bar":
            var v34 uint8 = 1
            _ = v34
        case "":
            if true {
                fmt.Println(v30, v28, v3)
            } else {
                v2 += v33
                _ = len(v29)
                var v35 uint64 = v1
                _ = v35
            }
        }
        var v36 float32 = (3.14159 - 0.0)
        v37 := GenS1{F0: -935, F1: 7, F2: (2 == 4294967295)}
        _ = v36
        _ = v37
    } else {
        var v38 string = "hello"
        for fi39 := 0; fi39 < 9; fi39++ {
            fmt.Println(v13, v30, v3)
            var v40 int16 = -244
            if true {
                v33 += 65535
                if (2 > 255) {
                    var v41 bool = v28
                    _ = v41
                }
                if (true || v30) {
                    var v42 int16 = 683
                    v2 = 65535
                    var v43 int = (fi39 / ((v13) | 1))
                    _ = v42
                    _ = v43
                } else {
                    if v28 {
                        var v44 uint16 = v2
                        var v45 float32 = (-2.25 - 1.5)
                        _ = v44
                        _ = v45
                    }
                    fmt.Println(v30)
                }
            }
            _ = fi39
            _ = v40
        }
        _ = v38
    }
    v33 = 0
    var v46 uint16 = v2
    v1 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v46
}
