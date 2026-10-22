package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint; F2 int16 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 int32 }
func genFn0(p0 uint8, p1 int, p2 bool) int {
    p2 = (!(0 != 0))
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := []int{}
            switch fi8 {
            case 32768:
                var v10 float64 = 3.14159
                _ = v10
            case 128:
                fmt.Println(p1)
            case -925:
                p2 = (((float64(0.1) / (1e10 * -2.25)) + 0.0) <= -2.25)
            }
            if len(v9) > 0 {
                v9[0] = (fi8 + (fi7 & -2147483648))
            }
            _ = fi8
            _ = v9
        }
        for fi11 := 0; fi11 < 3; fi11++ {
            fmt.Println(fi11)
            _ = fi11
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 >> ((328) & 63))
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    v4 := make(map[string]string)
    v5 := []string{("" + "foo bar"), "a", "foo bar"}
    var v6 bool = false
    var v12 int32 = 447
    var v13 int16 = 32767
    if v6 {
        fmt.Println(v13, v12, v12)
        for fi14 := 0; fi14 < 16; fi14++ {
            v13 = 1000
            _ = fi14
        }
    }
    var v15 uint32 = 92
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 int = fi16
        for fi18 := 0; fi18 < 9; fi18++ {
            switch "" {
            case "foo bar":
                if (65535 == 4294967295) {
                    v13 += v13
                    delete(v2, (("hello" + "a") + "hello"))
                    v19 := GenS1{F0: 42}
                    _ = v19
                }
            default:
                var v20 int64 = -9223372036854775808
                _ = v20
            }
            _ = fi18
        }
        if v6 {
            fmt.Println(v12, v13)
            switch (("" + "a") + "hello") {
            case "hello":
                if (0.0 > 3.14159) {
                    v13 = v1
                    var v21 float32 = 1e10
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 14; fi22++ {
                        switch (fi22 << ((255) & 63)) {
                        case -412:
                            v23 := GenS2{F0: 0, F1: 231, F2: 452}
                            _ = v23
                        case 32768:
                            if true {
                                v24 := GenS3{F0: 42, F1: (-(v12))}
                                _ = v24
                            } else {
                                var v25 int16 = v1
                                _ = v25
                            }
                        case 128:
                            v6 = v6
                        }
                        switch fi22 {
                        case 265:
                            var v26 string = ("" + "")
                            _ = v26
                        case 127:
                            v27 := GenS3{F0: 2, F1: (v12 << ((0) & 63))}
                            if (1 > 1) {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v5 = append(v5, "\n\t")
                            } else {
                                if (128 < 0) {
                                    var v28 bool = (42 > -22)
                                    v28 = (4294967295 < 910)
                                    _ = v28
                                } else {
                                    v29 := GenS4{F0: 100, F1: "a"}
                                    v30 := GenS3{F0: 0, F1: v12}
                                    var v31 int16 = (v13 - (v1 % ((v1) | 1)))
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                }
                            }
                            v1 = -32768
                            fmt.Println(fi16, v1)
                            _ = v27
                        case -2147483648:
                            _ = len(v2)
                        default:
                            v32 := v4[""]
                            _ = v32
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi22
                    }
                }
            case "\n\t":
                fmt.Println(v17, v6, fi16)
            }
        }
        _ = fi16
        _ = v17
    }
    for fi33 := 0; fi33 < 5; fi33++ {
        var v34 string = ""
        var v35 int16 = 1
        fmt.Println(v15)
        _ = fi33
        _ = v34
        _ = v35
    }
    v36 := GenS5{F0: v12}
    v6 = v6
    v37 := v3[0]
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
