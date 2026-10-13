package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 string; F2 int8 }
type GenS5 struct { F0 int16; F1 uint; F2 int32 }
type GenS6 struct { F0 uint }

func main() {
    var v1 string = ((("\n\t" + ("\n\t" + (("a" + "") + "a"))) + "a") + ("foo bar" + ("" + ((("" + "foo bar") + "") + "foo bar"))))
    var v2 float32 = 1e10
    var v3 int8 = -128
    var v4 int64 = 42
    v5 := GenS1{F0: 42, F1: 128}
    v4 += v4
    v4 += 0
    fmt.Println(v1)
    switch v1 {
    case "\n\t":
        switch "a" {
        case "a":
            for fi6 := 0; fi6 < 13; fi6++ {
                var v7 int = fi6
                if (true && (1 > 32767)) {
                    v1 = v1
                    v2 -= 0.1
                    v8 := []int{(-(v7)), (v7 / ((fi6) | 1))}
                    _ = v8
                } else {
                    v9 := GenS1{F0: -1, F1: -9223372036854775808}
                    v10 := []int{(v7 << ((255) & 63))}
                    if (((v2 + ((0.0 + v2) / v2)) * (1.5 / v2)) >= -2.25) {
                        if (!((false || ((1.5 > (v2 - (v2 / -2.25))) || (-1 == 32767))) && false)) {
                            switch 127 {
                            case -2147483648:
                                v3 += int8(v7)
                            }
                            v2 -= v2
                        }
                        if false {
                            var v11 uint8 = 237
                            _ = v11
                        }
                        v1 = "a"
                    }
                    _ = v10
                    _ = v9
                }
                _ = fi6
                _ = v7
            }
        case "hello":
            v12 := GenS1{F0: 1, F1: 1}
            _ = v12
        default:
            if (1 != -77) {
                switch ("hello" + v1) {
                case "":
                    v2 = float32(1.5)
                case "a":
                    v1 = "a"
                }
                var v13 float64 = 1.5
                _ = v13
            } else {
                for fi14 := 0; fi14 < 10; fi14++ {
                    var v15 int32 = 42
                    _ = fi14
                    _ = v15
                }
            }
        }
    default:
        var v16 int64 = v4
        _ = v16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        v2 += 0.0
        var v18 uint8 = 42
        _ = fi17
        _ = v18
    }
    var v19 int32 = -1
    switch 958 {
    case 195:
        v20 := []int{32768, -96, 127, 32767}
        _ = v20
    }
    v21 := GenS2{F0: 132, F1: (-(v19))}
    fmt.Println(v4, v19)
    if (799 <= 159) {
        if (0 >= 2147483647) {
            if ((665 <= 42) && true) {
                for fi22 := 0; fi22 < 10; fi22++ {
                    if true {
                        switch int(v3) {
                        case -9223372036854775808:
                            for fi23 := 0; fi23 < 17; fi23++ {
                                v3 = (v3 >> ((4294967295) & 63))
                                var v24 uint64 = 18446744073709551615
                                if true {
                                    var v25 uint8 = 128
                                    _ = v25
                                } else {
                                    v26 := []int{(fi22 % (((^(fi23))) | 1)), fi23, fi23, fi22}
                                    _ = v26
                                }
                                _ = fi23
                                _ = v24
                            }
                        case 42:
                            if true {
                                v27 := GenS1{F0: v19, F1: 42}
                                _ = v27
                            } else {
                                v1 = (v1 + ("\n\t" + "\n\t"))
                                fmt.Println(v2, v1, v19)
                                for fi28 := 0; fi28 < 18; fi28++ {
                                    v29 := []int{fi28}
                                    v30 := GenS3{F0: 32767}
                                    v29 = append(v29, fi28)
                                    v29 = append(v29, (fi22 + 42))
                                    if (v19 <= 656) {
                                        fmt.Println(v4, v1, fi22)
                                        v4 -= 0
                                        switch (fi22 >> ((637) & 63)) {
                                        case 32767:
                                            if len(v29) > 0 {
                                                v29[0] = 9223372036854775807
                                            }
                                        case -2147483648:
                                            var v31 string = v1
                                            _ = v31
                                        }
                                    }
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                            }
                        case -619:
                            for fi32 := 0; fi32 < 6; fi32++ {
                                var v33 int32 = (v19 / (((v19 % ((v19) | 1))) | 1))
                                _ = fi32
                                _ = v33
                            }
                        default:
                            v34 := []int{}
                            _ = v34
                        }
                        switch (fi22 + int(v3)) {
                        case -670:
                            v1 = (v1 + v1)
                        case 32767:
                            fmt.Println(v2, v1)
                        case 128:
                            v35 := GenS4{F0: 42, F1: "hello", F2: (v3 ^ (v3 % ((v3) | 1)))}
                            _ = v35
                        }
                        var v36 uint8 = 1
                        _ = v36
                    }
                    _ = fi22
                }
                if (((v19 + 1) <= v19) || (42 < 2147483648)) {
                    for fi37 := 0; fi37 < 9; fi37++ {
                        switch fi37 {
                        case 9223372036854775807:
                            var v38 float64 = 0.0
                            _ = v38
                        case 127:
                            v39 := GenS4{F0: 219, F1: v1, F2: v3}
                            _ = v39
                        }
                        _ = fi37
                    }
                }
                var v40 bool = true
                _ = v40
            } else {
                fmt.Println(v2, v2)
                var v41 uint16 = 629
                v42 := []int{-527, 1, 128, int(v4)}
                _ = v41
                _ = v42
            }
        }
    } else {
        for fi43 := 0; fi43 < 4; fi43++ {
            var v44 int16 = 137
            var v45 int16 = (v44 | v44)
            _ = fi43
            _ = v44
            _ = v45
        }
        var v46 int = 0
        v2 += -2.25
        _ = v46
    }
    v47 := []int{1, -665, 32767}
    if (65535 == 255) {
        v48 := GenS5{F0: 32767, F1: 429, F2: v19}
        _ = v48
    }
    var v49 int = 42
    switch v49 {
    case 127:
        v50 := GenS6{F0: 147}
        _ = v50
    default:
        v51 := GenS1{F0: v19, F1: (-(v49))}
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v47
    _ = v49
    _ = v5
}
