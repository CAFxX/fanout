package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int8; F1 int; F2 int16 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 int32 }
func genFn0() int64 {
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := []int{(fi4 * (fi4 ^ fi4))}
        v6 := GenS1{F0: 2, F1: 42}
        var v7 float32 = -2.25
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 uint64 = 1
    v8 -= v8
    fmt.Println(v8)
    _ = v8
    return 42
}
func genFn1() int {
    var v9 string = "\n\t"
    var v10 uint8 = 8
    var v11 float32 = 1e10
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 int16 = 1000
        v13 = (v13 ^ (^(v13)))
        switch 0 {
        case 32768:
            for fi14 := 0; fi14 < 2; fi14++ {
                var v15 uint8 = 127
                _ = fi14
                _ = v15
            }
        case 2147483647:
            var v16 string = ("\n\t" + (((v9 + "") + "") + v9))
            _ = v16
        case 0:
            v17 := GenS2{F0: 1, F1: (fi12 << ((42) & 63)), F2: -884}
            _ = v17
        default:
            switch v9 {
            case "a":
                var v18 uint32 = 2147483648
                _ = v18
            case "\n\t":
                v11 = float32((-2.25 + 0.1))
            case "foo bar":
                if (0 <= v10) {
                    if false {
                        var v19 int64 = 962
                        var v20 float64 = 0.1
                        for fi21 := 0; fi21 < 9; fi21++ {
                            fmt.Println(fi21, v9)
                            v22 := []int{(fi21 >> ((4294967295) & 63)), fi12, (fi21 >> ((65535) & 63)), fi12}
                            _ = fi21
                            _ = v22
                        }
                        _ = v19
                        _ = v20
                    } else {
                        v13 -= (v13 >> ((18446744073709551615) & 63))
                        var v23 int = (fi12 + (-(fi12)))
                        switch v9 {
                        case "a":
                            if false {
                                var v24 uint32 = 2147483647
                                fmt.Println(v24, v23)
                                _ = v24
                            }
                        }
                        _ = v23
                    }
                } else {
                    switch "" {
                    case "a":
                        var v25 int8 = 42
                        _ = v25
                    case "":
                        v26 := []int{fi12, -2147483648}
                        _ = v26
                    }
                }
            default:
                var v27 uint = 361
                _ = v27
            }
        }
        _ = fi12
        _ = v13
    }
    _ = v10
    _ = v11
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 int16 = -603
    var v2 uint8 = 128
    var v3 int64 = 0
    v28 := []int{0, 42, 0}
    v1 += v1
    var v29 int32 = -768
    var v30 int = 691
    if false {
        switch (("a" + "") + "hello") {
        case "":
            for fi31 := 0; fi31 < 9; fi31++ {
                v32 := GenS1{F0: 1, F1: 51}
                v30 += (fi31 & 32768)
                v1 -= 644
                _ = fi31
                _ = v32
            }
        default:
            if ((v2 / (((^(v2))) | 1)) >= (v2 - 111)) {
                var v33 float32 = 0.0
                if (-117 == 1) {
                    var v34 int = 2147483647
                    if (277 < 798) {
                        v35 := GenS3{F0: v30, F1: "foo bar"}
                        for fi36 := 0; fi36 < 7; fi36++ {
                            v37 := GenS4{F0: 42, F1: true}
                            _ = fi36
                            _ = v37
                        }
                        _ = v35
                    }
                    v3 -= v3
                    _ = v34
                }
                _ = v33
            } else {
                _ = len(v28)
            }
        }
    } else {
        switch ("a" + "hello") {
        case "foo bar":
            if (0.1 >= (1.5 - -2.25)) {
                v38 := GenS5{F0: 380, F1: "a"}
                var v39 int64 = genFn0()
                _ = v38
                _ = v39
            } else {
                for fi40 := 0; fi40 < 6; fi40++ {
                    v28 = append(v28, -1)
                    _ = fi40
                }
                var v41 bool = (4294967295 < 214)
                var v42 string = ("foo bar" + "a")
                _ = v41
                _ = v42
            }
        case "\n\t":
            v1 += 0
        default:
            v43 := GenS3{F0: 32768, F1: ("\n\t" + "a")}
            _ = v43
        }
        v44 := GenS2{F0: 127, F1: (v30 ^ int(v2)), F2: (v1 << ((0) & 63))}
        var v45 uint8 = 2
        _ = v44
        _ = v45
    }
    v46 := GenS6{F0: v29}
    fmt.Println(v3, v1)
    if len(v28) > 0 {
        v28[0] = (-(v30))
    }
    var v47 float64 = (float64(float32((-2.25 + 0.0))) - 0.0)
    for fi48 := 0; fi48 < 4; fi48++ {
        v47 = v47
        v2 += 127
        var v49 float64 = v47
        _ = fi48
        _ = v49
    }
    v29 = v29
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v46
    _ = v47
}
