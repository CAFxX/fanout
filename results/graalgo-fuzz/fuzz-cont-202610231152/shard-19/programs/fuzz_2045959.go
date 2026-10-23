package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint64 }
func genFn0() string {
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := GenS1{F0: 1}
        v7 := []int{fi5, fi5, 2147483647, (fi5 >> ((0) & 63))}
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int64 = -352
    _ = v8
    return "\n\t"
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = ((-2.25 / ((float64(0.0) / 1.5) - (0.1 - (float64(0.0) - -2.25)))) + (-2.25 / float64(3.14159)))
    var v3 float32 = 3.14159
    var v4 string = ("\n\t" + "foo bar")
    fmt.Println(v3, v4, v4)
    v9 := GenS2{F0: 42, F1: 100}
    for fi10 := 0; fi10 < 11; fi10++ {
        if true {
            v11 := GenS1{F0: 253}
            v12 := v1[(fi10 % ((4294967295) | 1))]
            _ = v11
            _ = v12
        } else {
            delete(v1, fi10)
        }
        _ = fi10
    }
    fmt.Println(v2, v3)
    fmt.Println(v2, v4, v4)
    for fi13 := 0; fi13 < 2; fi13++ {
        v4 = (v4 + v4)
        fmt.Println(v4)
        _ = fi13
    }
    if (v4 > "foo bar") {
        switch "foo bar" {
        case "\n\t":
            _ = len(v1)
        }
        var v14 uint16 = 0
        if (0 == -64) {
            v2 = (v2 * (-2.25 + v2))
            var v15 float64 = float64(1.5)
            v2 += 3.14159
            _ = v15
        }
        _ = v14
    }
    var v16 uint8 = 128
    if (18446744073709551615 >= 0) {
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 uint32 = 2
            fmt.Println(v3, v18, fi17)
            v19 := []int{-841, (-(fi17)), (fi17 + fi17), fi17}
            _ = fi17
            _ = v18
            _ = v19
        }
        var v20 uint64 = 2
        _ = v20
    } else {
        switch v4 {
        case "foo bar":
            if ((18446744073709551615 != 2) || ((((v4 + (v4 + v4)) + (("a" + "a") + "\n\t")) + v4) >= v4)) {
                v21 := GenS2{F0: -284, F1: -50}
                if ((4294967295 == 4294967295) && (-1 > 389)) {
                    var v22 uint16 = 32767
                    if (18446744073709551615 != 4294967295) {
                        for fi23 := 0; fi23 < 5; fi23++ {
                            v24 := []int{}
                            _ = fi23
                            _ = v24
                        }
                        v25 := GenS3{F0: 4294967295}
                        var v26 int8 = 0
                        _ = v25
                        _ = v26
                    }
                    _ = v22
                } else {
                    v27 := GenS2{F0: 42, F1: 127}
                    _ = v27
                }
                if false {
                    v28 := GenS3{F0: 4294967295}
                    _ = v28
                } else {
                    v29 := GenS3{F0: 42}
                    v30 := []int{128, -9223372036854775808, 127}
                    if false {
                        var v31 uint64 = 1
                        if false {
                            v1[int(v31)] = 4294967295
                        } else {
                            if false {
                                if false {
                                    var v32 uint64 = v31
                                    _ = v32
                                }
                                switch -1 {
                                case 79:
                                    for fi33 := 0; fi33 < 20; fi33++ {
                                        fmt.Println(v4, fi33, v16)
                                        if (4294967295 == 41) {
                                            if true {
                                                if (!false) {
                                                    v16 += v16
                                                    fmt.Println(v3)
                                                } else {
                                                    v31 = 55
                                                    var v34 float32 = v3
                                                    v35 := GenS2{F0: 0, F1: -128}
                                                    _ = v34
                                                    _ = v35
                                                }
                                            }
                                            var v36 int8 = 16
                                            _ = v36
                                        }
                                        var v37 float32 = v3
                                        _ = fi33
                                        _ = v37
                                    }
                                default:
                                    v30 = append(v30, 1)
                                }
                                if (1 > (v31 | v31)) {
                                    v16 = 128
                                }
                            }
                        }
                        _ = v31
                    }
                    _ = v29
                    _ = v30
                }
                _ = v21
            } else {
                v38 := GenS2{F0: 1000, F1: 42}
                switch v4 {
                case "\n\t":
                    switch "" {
                    case "":
                        v39 := []int{2147483647, 2147483647, -386, -1}
                        _ = v39
                    }
                default:
                    v2 = (v2 * float64(-2.25))
                }
                _ = v38
            }
        case "\n\t":
            var v40 int32 = 1
            _ = v40
        }
    }
    v41 := v1[317]
    var v42 int = (v41 - (v41 << ((1) & 63)))
    v41 += 0
    var v43 float32 = (0.0 / (v3 + v3))
    var v44 int16 = -661
    var v45 uint32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v9
}
