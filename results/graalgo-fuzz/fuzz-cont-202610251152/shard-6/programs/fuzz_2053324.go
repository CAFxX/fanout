package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 string; F2 uint }
type GenS3 struct { F0 string; F1 int16; F2 uint }
func genFn0(p0 uint64) uint16 {
    var v5 uint32 = 644
    p0 -= p0
    var v6 int32 = -1
    fmt.Println(v6, v6)
    _ = p0
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 uint64 = 4294967295
    var v2 bool = true
    var v3 float32 = 0.1
    var v4 float32 = (v3 + ((v3 + v3) * -2.25))
    v2 = (2 == 127)
    switch 9223372036854775807 {
    case -2147483648:
        var v7 uint = 2
        _ = v7
    case 128:
        if ((((0.1 - (((0.0 - ((1.5 - 1.5) * -2.25)) * 3.14159) * 0.1)) * 3.14159) == ((3.14159 * 3.14159) / 3.14159)) || (!(1 <= -9223372036854775808))) {
            var v8 uint = 843
            if (!(v2 && v2)) {
                if (int16(v1) >= 42) {
                    for fi9 := 0; fi9 < 7; fi9++ {
                        v4 -= float32(0.0)
                        v10 := GenS1{F0: 2}
                        _ = fi9
                        _ = v10
                    }
                    v11 := []int{9223372036854775807, 735, -94}
                    _ = v11
                }
            } else {
                v12 := []int{}
                if v2 {
                    v3 -= v3
                    var v13 uint = v8
                    var v14 float64 = (-2.25 + ((1.5 / 1.5) - -2.25))
                    _ = v13
                    _ = v14
                } else {
                    v12 = append(v12, 636)
                    fmt.Println(v2, v1, v3)
                }
                v4 = (float32((-2.25 - (0.1 + 0.0))) - (float32(3.14159) / (float32(0.0) + 0.1)))
                _ = v12
            }
            v2 = (v4 > v4)
            _ = v8
        }
    case 547:
        var v15 bool = v2
        _ = v15
    default:
        var v16 uint32 = 111
        _ = v16
    }
    v1 += (v1 | 1)
    switch ("a" + "a") {
    case "a":
        for fi17 := 0; fi17 < 5; fi17++ {
            v18 := GenS1{F0: 875}
            if (true || v2) {
                if v2 {
                    var v19 uint = 957
                    fmt.Println(v4, v1)
                    switch 127 {
                    case 128:
                        switch 2147483647 {
                        case 155:
                            v19 += v19
                        case 2147483647:
                            v20 := []int{}
                            _ = v20
                        case -2147483648:
                            v21 := GenS1{F0: uint32(v1)}
                            _ = v21
                        }
                    default:
                        switch "foo bar" {
                        case "\n\t":
                            switch -23 {
                            case -9223372036854775808:
                                v2 = true
                            case 572:
                                for fi22 := 0; fi22 < 6; fi22++ {
                                    v3 += v4
                                    v23 := []int{(fi17 << ((0) & 63))}
                                    _ = fi22
                                    _ = v23
                                }
                            }
                        case "":
                            v24 := []int{-579, fi17, (fi17 % ((2147483647) | 1)), (fi17 / (((fi17 << ((v19) & 63))) | 1))}
                            _ = v24
                        case "a":
                            var v25 float64 = 0.1
                            _ = v25
                        }
                    }
                    _ = v19
                } else {
                    for fi26 := 0; fi26 < 9; fi26++ {
                        var v27 int8 = 42
                        v28 := []int{1}
                        v1 = (v1 / (((v1 + v1)) | 1))
                        var v29 int = fi26
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    var v30 uint64 = 1
                    _ = v30
                }
                v31 := GenS1{F0: 2147483647}
                _ = v31
            }
            v3 -= (v4 * v3)
            _ = fi17
            _ = v18
        }
    case "":
        for fi32 := 0; fi32 < 18; fi32++ {
            v33 := GenS2{F0: 52, F1: "hello", F2: 4294967295}
            _ = fi32
            _ = v33
        }
    default:
        switch 127 {
        case -9223372036854775808:
            v34 := GenS1{F0: 2147483647}
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 20; fi35++ {
        v2 = false
        for fi36 := 0; fi36 < 7; fi36++ {
            v37 := []int{(fi36 / (((fi35 >> ((65535) & 63))) | 1)), (fi35 / ((6) | 1))}
            _ = fi36
            _ = v37
        }
        fmt.Println(v1, v3)
        _ = fi35
    }
    var v38 int8 = 127
    var v39 uint = 1
    var v40 string = "a"
    v40 = ((v40 + v40) + ("" + v40))
    for fi41 := 0; fi41 < 16; fi41++ {
        for fi42 := 0; fi42 < 9; fi42++ {
            fmt.Println(v1, fi41)
            _ = fi42
        }
        if ((v2 && (3.14159 != (0.0 + ((3.14159 * 3.14159) - (0.0 / ((((3.14159 / float64(v3)) * 3.14159) - 1e10) / 1e10)))))) && (v1 <= (v1 << (((v39 % ((v39) | 1))) & 63)))) {
            v43 := GenS1{F0: 2}
            _ = v43
        } else {
            v1 -= (^(v1))
        }
        if false {
            var v44 uint16 = genFn0((v1 % ((v1) | 1)))
            v4 = float32(((float64(v3) - (0.1 * float64(v3))) * (0.1 + 1.5)))
            switch (fi41 + fi41) {
            case 395:
                var v45 uint16 = v44
                _ = v45
            case 128:
                v46 := GenS1{F0: 274}
                _ = v46
            case 32768:
                fmt.Println(v3, v3, v4)
            default:
                v47 := []int{-96, -818, (fi41 >> (((v39 >> ((346) & 63))) & 63)), -9223372036854775808}
                _ = v47
            }
            _ = v44
        }
        _ = fi41
    }
    if (!v2) {
        v48 := []int{2147483647}
        _ = v48
    }
    if (v2 && (!false)) {
        v49 := []int{0, -1}
        _ = v49
    }
    v50 := GenS1{F0: 65535}
    v4 += 1e10
    v51 := []int{32767, 128}
    v52 := GenS3{F0: v40, F1: -1, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v51:", len(v51), v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v50
    _ = v51
    _ = v52
}
