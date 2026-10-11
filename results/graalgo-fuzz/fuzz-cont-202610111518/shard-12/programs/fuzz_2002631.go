package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 uint; F1 uint }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int16; F1 int16 }
type GenS9 struct { F0 uint; F1 bool; F2 uint }
func genFn0(p0 uint64, p1 uint16) int8 {
    v6 := []int{-224, -468, 4294967295, 9223372036854775807}
    var v7 float64 = 3.14159
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -128
}
func genFn1() bool {
    var v8 int16 = -32768
    var v9 int32 = 0
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 uint = 4294967295
    v2 := []int{56, 32768, 9223372036854775807, 823}
    var v3 float64 = 0.1
    v4 := make(map[uint32]string)
    v5 := []string{}
    var v10 int = 32768
    v11 := GenS1{F0: 993, F1: 1}
    if false {
        v12 := GenS1{F0: 0, F1: 0}
        _ = v12
    } else {
        var v13 uint8 = 128
        v14 := GenS1{F0: 0, F1: -32768}
        v15 := v4[42]
        _ = v13
        _ = v14
        _ = v15
    }
    v3 += 0.1
    if ((-128 <= 1) || (0.1 <= (3.14159 * 1.5))) {
        fmt.Println(v1)
    } else {
        var v16 float32 = 1e10
        for fi17 := 0; fi17 < 18; fi17++ {
            switch ("hello" + "") {
            case "hello":
                var v18 int32 = -940
                _ = v18
            case "a":
                v1 += v1
            }
            var v19 int32 = 1
            _ = fi17
            _ = v19
        }
        _ = v16
    }
    if false {
        switch "" {
        case "":
            var v20 int16 = -1
            _ = v20
        case "foo bar":
            var v21 uint32 = 4294967295
            _ = v21
        default:
            switch "foo bar" {
            case "a":
                for fi22 := 0; fi22 < 16; fi22++ {
                    switch "a" {
                    case "a":
                        if (!false) {
                            switch fi22 {
                            case -1:
                                v23 := GenS2{F0: 9223372036854775807}
                                _ = v23
                            case 0:
                                v24 := GenS3{F0: fi22}
                                _ = v24
                            }
                            fmt.Println(v3, v10, fi22)
                            fmt.Println(v1, v10)
                        }
                    case "hello":
                        var v25 uint = v1
                        _ = v25
                    }
                    _ = fi22
                }
            case "foo bar":
                switch (v10 >> (((v1 + (v1 % ((v1) | 1)))) & 63)) {
                case -856:
                    var v26 bool = false
                    _ = v26
                }
            case "hello":
                v27 := GenS4{F0: "a", F1: (2147483647 != -2147483648)}
                _ = v27
            default:
                v3 -= (v3 * -2.25)
            }
        }
        switch "a" {
        case "":
            if (!true) {
                var v28 int32 = -611
                _ = v28
            } else {
                v10 = v10
            }
        case "foo bar":
            var v29 int = 4294967295
            _ = v29
        }
        for fi30 := 0; fi30 < 10; fi30++ {
            for fi31 := 0; fi31 < 14; fi31++ {
                if (true || true) {
                    var v32 bool = false
                    _ = v32
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if ((2 < 18446744073709551615) || (255 >= 65535)) {
                        v33 := GenS5{F0: 94, F1: -2147483648}
                        if (32 < 42) {
                            switch "a" {
                            case "":
                                v34 := GenS6{F0: (v1 * (v1 >> ((331) & 63))), F1: v1}
                                var v35 int32 = 0
                                _ = v34
                                _ = v35
                            }
                            if true {
                                var v36 string = "hello"
                                fmt.Println(v10, v1, v36)
                                switch fi30 {
                                case 42:
                                    v4[512] = "hello"
                                case -1:
                                    v3 = ((v3 - (3.14159 + 0.1)) + 1.5)
                                }
                                _ = v36
                            } else {
                                var v37 int16 = -32768
                                v10 = 995
                                v38 := GenS4{F0: (("hello" + "a") + ""), F1: false}
                                v10 -= -2147483648
                                v10 -= fi30
                                _ = v37
                                _ = v38
                            }
                        } else {
                            v3 -= (0.0 + 3.14159)
                            var v39 int32 = 1
                            _ = v39
                        }
                        _ = v33
                    } else {
                        v1 += v1
                        var v40 int16 = int16(v1)
                        _ = v40
                    }
                    if (true || false) {
                        v41 := GenS6{F0: (v1 ^ (^(v1))), F1: v1}
                        v42 := GenS3{F0: fi30}
                        _ = v41
                        _ = v42
                    }
                    v43 := GenS7{F0: "hello"}
                    _ = v43
                }
                _ = fi31
            }
            v1 = v1
            _ = fi30
        }
    }
    fmt.Println(v10)
    fmt.Println(v1)
    v10 = (v10 & (v10 + v10))
    for fi44 := 0; fi44 < 20; fi44++ {
        if false {
            v1 = v1
            v3 -= (v3 + v3)
        }
        var v45 uint64 = 1
        v46 := GenS8{F0: 1000, F1: 605}
        _ = fi44
        _ = v45
        _ = v46
    }
    var v47 bool = false
    for fi48 := 0; fi48 < 1; fi48++ {
        for fi49 := 0; fi49 < 20; fi49++ {
            v50 := GenS3{F0: (-(fi48))}
            _ = fi49
            _ = v50
        }
        v51 := GenS3{F0: -589}
        v52 := GenS9{F0: 0, F1: v47, F2: 0}
        _ = fi48
        _ = v51
        _ = v52
    }
    var v53 float64 = (float64(((((3.14159 / 0.1) + 0.1) * 1.5) * (3.14159 * -2.25))) - (1.5 / (v3 * ((v3 - float64(3.14159)) - v3))))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v53
}
