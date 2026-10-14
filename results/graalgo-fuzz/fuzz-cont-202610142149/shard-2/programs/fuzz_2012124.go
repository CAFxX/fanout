package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int32; F2 uint }
type GenS4 struct { F0 string; F1 int64 }
func genFn0() int16 {
    for fi4 := 0; fi4 < 15; fi4++ {
        v5 := []int{fi4, (fi4 / ((42) | 1)), 1, fi4}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 uint16 = 2
            var v9 uint = 42
            _ = fi7
            _ = v8
            _ = v9
        }
        if (-2147483648 >= -1) {
            var v10 bool = (1.5 >= 0.0)
            for fi11 := 0; fi11 < 10; fi11++ {
                var v12 uint8 = 1
                var v13 uint8 = (v12 * (^(v12)))
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = v10
        } else {
            var v14 int8 = 34
            _ = v14
        }
        _ = fi6
    }
    return -32768
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = 42
    var v3 uint = 0
    switch 1 {
    case 251:
        delete(v1, 2147483648)
    default:
        _ = len(v1)
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        for fi16 := 0; fi16 < 6; fi16++ {
            v3 -= (^(v3))
            _ = fi16
        }
        if (128 != 67) {
            if true {
                v17 := GenS1{F0: 42, F1: -147}
                v18 := GenS2{F0: 1}
                _ = v17
                _ = v18
            } else {
                v2 = (v2 / ((genFn0()) | 1))
                v19 := v1[42]
                for fi20 := 0; fi20 < 15; fi20++ {
                    if false {
                        var v21 uint8 = 2
                        v19 = v19
                        if true {
                            var v22 uint32 = 65535
                            v3 += (v3 << (((v3 & (v3 + v3))) & 63))
                            fmt.Println(fi20, v2, fi15)
                            _ = v22
                        } else {
                            v23 := GenS2{F0: 1}
                            _ = v23
                        }
                        _ = v21
                    }
                    delete(v1, 739)
                    var v24 uint32 = 65535
                    _ = fi20
                    _ = v24
                }
                _ = v19
            }
        }
        _ = fi15
    }
    if ((v3 & (v3 & (v3 >> ((42) & 63)))) != 65535) {
        fmt.Println(v3, v3)
        for fi25 := 0; fi25 < 3; fi25++ {
            for fi26 := 0; fi26 < 11; fi26++ {
                delete(v1, 352)
                v3 -= (v3 & v3)
                if false {
                    if true {
                        fmt.Println(v3, v3)
                        switch 1 {
                        case 4294967295:
                            v3 = v3
                        case 2147483647:
                            switch ("" + "") {
                            case "":
                                if true {
                                    fmt.Println(fi26, v2, v3)
                                    v27 := GenS1{F0: (-(fi25)), F1: -1}
                                    v28 := GenS2{F0: 162}
                                    _ = v27
                                    _ = v28
                                }
                            case "foo bar":
                                v29 := GenS2{F0: 2147483647}
                                _ = v29
                            case "hello":
                                v30 := []int{fi26, 2147483647, -805}
                                _ = v30
                            }
                        case 32768:
                            if (true && (("\n\t" + "\n\t") != "foo bar")) {
                                if (352 <= 0) {
                                    v31 := GenS1{F0: (fi26 * fi26), F1: 9223372036854775807}
                                    _ = v31
                                }
                                v3 -= 2
                            } else {
                                var v32 float32 = float32(1.5)
                                v3 += (v3 ^ v3)
                                _ = v32
                            }
                        }
                    } else {
                        fmt.Println(v2)
                        v33 := GenS3{F0: -1, F1: 42, F2: v3}
                        _ = v33
                    }
                    if false {
                        v34 := GenS4{F0: (("a" + ("hello" + "foo bar")) + "\n\t"), F1: -946}
                        v35 := GenS3{F0: 0, F1: -2147483648, F2: (^(v3))}
                        switch (((("\n\t" + "") + "foo bar") + "\n\t") + ("a" + ("foo bar" + "a"))) {
                        case "hello":
                            var v36 uint8 = 2
                            _ = v36
                        case "":
                            v2 = (^(v2))
                        }
                        _ = v34
                        _ = v35
                    } else {
                        v2 += genFn0()
                    }
                }
                _ = fi26
            }
            v2 -= (v2 | int16(fi25))
            v2 -= (v2 & genFn0())
            _ = fi25
        }
        var v37 uint8 = 213
        _ = v37
    }
    v38 := []int{2147483647, 554, 28}
    v38 = append(v38, 42)
    for fi39 := 0; fi39 < 3; fi39++ {
        if true {
            var v40 float64 = (((0.1 - (0.1 - float64(((-2.25 - (1.5 / 1.5)) + float32((0.1 + (3.14159 + float64(float32(((0.0 + 0.1) - 0.0)))))))))) - float64((0.1 + (-2.25 / 3.14159)))) * 3.14159)
            _ = v40
        } else {
            var v41 int64 = 28
            var v42 float32 = (((-2.25 - 0.1) + (float32((0.1 * 3.14159)) + ((0.1 - float32(3.14159)) - 3.14159))) + ((1e10 - ((0.1 * 3.14159) + (1e10 + 1.5))) * (1e10 * 0.0)))
            delete(v1, 42)
            _ = v41
            _ = v42
        }
        _ = fi39
    }
    var v43 bool = (true || true)
    _ = len(v38)
    var v44 int16 = 0
    if ((v44 == (v2 / (((-(v44))) | 1))) || ((float64(-2.25) * float64(0.1)) <= 1.5)) {
        var v45 int16 = (^(v44))
        fmt.Println(v43, v45)
        _ = v45
    } else {
        v44 = v44
    }
    var v46 float64 = ((float64((1.5 * 3.14159)) + ((0.0 * 1.5) * (float64((1e10 / (0.1 + (((1.5 * 3.14159) + 0.1) + 1e10)))) + (0.0 + float64(-2.25))))) - (1e10 * -2.25))
    var v47 float32 = -2.25
    v48 := GenS3{F0: -475, F1: 42, F2: v3}
    var v49 int16 = (v2 * (v2 % (((^(v2))) | 1)))
    v50 := GenS1{F0: 0, F1: 42}
    var v51 uint32 = 407
    var v52 int32 = 693
    if (42 != int16(v3)) {
        if v43 {
            var v53 uint16 = 2
            var v54 float64 = v46
            _ = v53
            _ = v54
        } else {
            for fi55 := 0; fi55 < 6; fi55++ {
                fmt.Println(v49, v46, v47)
                _ = fi55
            }
            v56 := v1[0]
            var v57 bool = (42 <= 1)
            _ = v56
            _ = v57
        }
        fmt.Println(v52, v49, v2)
        fmt.Println(v43, v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v43
    _ = v44
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v50
    _ = v51
    _ = v52
}
