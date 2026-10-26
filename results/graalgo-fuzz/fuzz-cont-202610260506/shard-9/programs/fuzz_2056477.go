package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int16 }
type GenS2 struct { F0 int32; F1 uint16 }
func genFn0(p0 int16, p1 int64, p2 bool) uint {
    if true {
        p0 = (-(p0))
    } else {
        var v7 uint32 = 2147483647
        _ = v7
    }
    p2 = ((3.14159 > 0.0) && (p2 && (4294967295 >= 4294967295)))
    fmt.Println(p1)
    p0 += p0
    var v8 int = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 42
}

func main() {
    var v1 int8 = -1
    v2 := make(map[int]uint32)
    var v3 bool = (2147483647 <= 1)
    v4 := make(map[string]bool)
    var v5 int64 = 9223372036854775807
    var v6 float64 = 1.5
    for fi9 := 0; fi9 < 20; fi9++ {
        fmt.Println(fi9)
        _ = fi9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 uint16 = 32767
        _ = fi10
        _ = v11
    }
    var v12 int16 = 32767
    var v13 float32 = 1.5
    var v14 int64 = (v5 ^ v5)
    v3 = ((v3 && (1e10 > (v13 - (v13 / v13)))) || v3)
    var v15 int = 0
    var v16 float32 = v13
    if true {
        switch (-(v15)) {
        case 128:
            if (42 < 65535) {
                v12 = (-(v12))
            } else {
                for fi17 := 0; fi17 < 10; fi17++ {
                    v2[(fi17 % ((-1) | 1))] = 146
                    v16 = v16
                    v18 := []int{v15}
                    _ = fi17
                    _ = v18
                }
                v19 := GenS1{F0: 70, F1: -9223372036854775808, F2: v12}
                switch v15 {
                case -1:
                    var v20 float64 = (v6 - (0.0 * ((0.1 / (v6 - (((0.0 + 3.14159) / (v6 * v6)) * (v6 - ((v6 / v6) * -2.25))))) / 1.5)))
                    _ = v20
                case -9223372036854775808:
                    fmt.Println(v3)
                case 42:
                    for fi21 := 0; fi21 < 15; fi21++ {
                        for fi22 := 0; fi22 < 1; fi22++ {
                            v23 := []int{fi22}
                            v23 = append(v23, (^(fi22)))
                            v24 := GenS1{F0: -2147483648, F1: 1, F2: v12}
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        switch "\n\t" {
                        case "a":
                            fmt.Println(v1)
                        case "foo bar":
                            var v25 uint8 = 2
                            _ = v25
                        }
                        if (255 <= 2) {
                            v26 := []int{(fi21 & v15)}
                            var v27 bool = (!true)
                            _ = v26
                            _ = v27
                        } else {
                            fmt.Println(v15)
                        }
                        _ = fi21
                    }
                }
                _ = v19
            }
        case 1:
            v2[v15] = 403
        case 981:
            v28 := GenS1{F0: 42, F1: (v14 % ((v5) | 1)), F2: v12}
            _ = v28
        default:
            _ = len(v4)
        }
        for fi29 := 0; fi29 < 2; fi29++ {
            var v30 bool = ((!((255 <= 42) && (1 < 18446744073709551615))) && v3)
            _ = fi29
            _ = v30
        }
        fmt.Println(v13, v1, v16)
    }
    var v31 bool = false
    if v31 {
        for fi32 := 0; fi32 < 9; fi32++ {
            var v33 int16 = v12
            v34 := GenS2{F0: 2147483647, F1: 1}
            var v35 int16 = v12
            _ = fi32
            _ = v33
            _ = v34
            _ = v35
        }
        for fi36 := 0; fi36 < 12; fi36++ {
            var v37 int32 = -2147483648
            _ = fi36
            _ = v37
        }
    } else {
        fmt.Println(v3, v15)
    }
    switch v15 {
    case 32767:
        var v38 string = "\n\t"
        _ = v38
    case 127:
        var v39 string = ("a" + "")
        _ = v39
    default:
        for fi40 := 0; fi40 < 8; fi40++ {
            switch int(v14) {
            case 2147483647:
                var v41 uint8 = 110
                _ = v41
            default:
                v2[v15] = 2
            }
            _ = fi40
        }
    }
    v42 := GenS1{F0: -2147483648, F1: -423, F2: (v12 * (^(v12)))}
    switch (v15 - v15) {
    case -511:
        for fi43 := 0; fi43 < 18; fi43++ {
            var v44 uint64 = 42
            var v45 bool = (65535 != 985)
            v5 = 42
            _ = fi43
            _ = v44
            _ = v45
        }
    case 398:
        if false {
            if (32768 == (v15 ^ (v15 / ((int(v14)) | 1)))) {
                var v46 int = (v15 / ((-538) | 1))
                _ = v46
            }
            var v47 float32 = (0.1 * v13)
            _ = v47
        }
    default:
        v48 := GenS2{F0: 0, F1: 119}
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
