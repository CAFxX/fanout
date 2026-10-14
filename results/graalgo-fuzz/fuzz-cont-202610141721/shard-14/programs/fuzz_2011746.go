package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32 }
func genFn0(p0 bool, p1 int8) uint8 {
    fmt.Println(p1, p0)
    p0 = (0 != -961)
    if (-2147483648 > 0) {
        var v4 int64 = -9223372036854775808
        v4 += 37
        _ = v4
    }
    var v5 uint32 = 4294967295
    switch "foo bar" {
    case "a":
        v6 := GenS1{F0: 1}
        _ = v6
    case "hello":
        if (65535 != 0) {
            if ((p1 << ((517) & 63)) > -128) {
                for fi7 := 0; fi7 < 5; fi7++ {
                    for fi8 := 0; fi8 < 1; fi8++ {
                        fmt.Println(p0, fi7)
                        p0 = ((42 <= 18446744073709551615) && false)
                        var v9 uint64 = 614
                        if p0 {
                            switch "" {
                            case "hello":
                                v10 := GenS1{F0: -104}
                                _ = v10
                            case "\n\t":
                                p1 -= p1
                            default:
                                p0 = ((p0 && (!(-2.25 <= (1e10 - 3.14159)))) && true)
                            }
                            var v11 uint64 = v9
                            _ = v11
                        }
                        if (!(p1 >= p1)) {
                            var v12 int16 = 1000
                            if (p0 || true) {
                                v5 -= 495
                                p1 += (p1 << ((4294967295) & 63))
                            } else {
                                if (21 != (fi7 * (^(fi8)))) {
                                    v5 = v5
                                } else {
                                    v12 += v12
                                    var v13 int8 = (p1 | (p1 * (p1 ^ p1)))
                                    v9 = (^(v9))
                                    _ = v13
                                }
                                p1 = (^(p1))
                            }
                            v14 := GenS1{F0: 0}
                            var v15 uint = 2
                            _ = v12
                            _ = v14
                            _ = v15
                        } else {
                            fmt.Println(fi7, v5, p0)
                            v16 := GenS1{F0: -2147483648}
                            v17 := []int{}
                            _ = v16
                            _ = v17
                        }
                        _ = fi8
                        _ = v9
                    }
                    switch fi7 {
                    case 855:
                        v18 := []int{813}
                        _ = v18
                    }
                    _ = fi7
                }
            } else {
                for fi19 := 0; fi19 < 15; fi19++ {
                    p1 = -67
                    _ = fi19
                }
                v20 := GenS2{F0: p0}
                var v21 int16 = -1
                _ = v20
                _ = v21
            }
            var v22 uint = 922
            _ = v22
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return 0
}

func main() {
    v1 := []uint32{4294967295}
    v2 := []string{("\n\t" + "\n\t"), "foo bar", "", "foo bar"}
    var v3 float32 = 0.1
    var v23 float64 = (1e10 * ((0.0 * (((3.14159 + 1e10) + (0.1 * 1.5)) - (((((float64(0.1) / 1e10) / 3.14159) + 0.1) / float64(v3)) + 0.1))) / (-2.25 / float64(float32(1.5)))))
    var v24 int = 1
    for fi25 := 0; fi25 < 20; fi25++ {
        var v26 int16 = 931
        _ = len(v1)
        _ = fi25
        _ = v26
    }
    var v27 int8 = 85
    v1 = append(v1, 65535)
    v28 := GenS3{F0: 2147483647}
    var v29 uint64 = 0
    var v30 int16 = -32768
    switch "" {
    case "":
        fmt.Println(v27, v29, v30)
    case "\n\t":
        v31 := GenS2{F0: (v27 != (v27 << ((18446744073709551615) & 63)))}
        _ = v31
    case "foo bar":
        for fi32 := 0; fi32 < 13; fi32++ {
            var v33 uint16 = 32767
            fmt.Println(v30, v24, v3)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi32
            _ = v33
        }
    default:
        if (1 == 42) {
            var v34 float64 = 0.1
            _ = v34
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v35 int8 = (v27 / (((v27 / (((v27 >> ((329) & 63))) | 1))) | 1))
    _ = len(v2)
    var v36 int8 = 127
    var v37 int = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v37
}
