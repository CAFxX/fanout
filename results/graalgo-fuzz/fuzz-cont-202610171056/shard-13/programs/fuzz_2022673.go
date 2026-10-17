package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint32 }
func genFn0() uint8 {
    var v7 uint8 = 83
    var v8 float32 = (((0.1 * -2.25) * 3.14159) - ((((((1.5 + float32((-2.25 - (-2.25 - 3.14159)))) - float32(-2.25)) - float32(-2.25)) * float32((float64(0.1) - 0.0))) / 1.5) + float32(0.1)))
    var v9 uint = 255
    v9 = (v9 + 335)
    if (583 <= 0) {
        var v10 uint32 = 557
        fmt.Println(v7, v8)
        v10 += v10
        _ = v10
    } else {
        v7 += v7
        if true {
            var v11 int16 = -126
            for fi12 := 0; fi12 < 11; fi12++ {
                v13 := []int{(^(fi12)), (fi12 / (((fi12 << (((v9 << (((v9 * (v9 + (v9 / (((v9 >> ((4294967295) & 63))) | 1))))) & 63))) & 63))) | 1)), (fi12 % ((-9223372036854775808) | 1))}
                v8 += v8
                _ = fi12
                _ = v13
            }
            _ = v11
        }
    }
    _ = v7
    _ = v8
    _ = v9
    return v7
}
func genFn1() uint {
    switch 32768 {
    case -9223372036854775808:
        var v14 float64 = 3.14159
        _ = v14
    case 196:
        v15 := GenS1{F0: 42, F1: 2147483648, F2: 42}
        _ = v15
    case 208:
        if (656 <= 828) {
            var v16 uint32 = 65535
            _ = v16
        }
    }
    var v17 bool = (!(false || (-32768 != -32768)))
    if true {
        v18 := GenS1{F0: -2147483648, F1: 198, F2: 9223372036854775807}
        if (((float32(0.1) * 0.0) * 3.14159) == 1e10) {
            v19 := []int{1, 9223372036854775807, 1, 32767}
            _ = v19
        }
        if ((true && ("hello" <= ((("" + "hello") + (("" + "hello") + "hello")) + "foo bar"))) || (-1 > -128)) {
            var v20 int64 = 42
            v17 = (42 > 2)
            var v21 uint = 0
            _ = v20
            _ = v21
        }
        _ = v18
    }
    _ = v17
    return 4294967295
}

func main() {
    var v1 float32 = float32(0.1)
    var v2 uint32 = 807
    var v3 int8 = 0
    var v4 int64 = -1
    var v5 float64 = float64(0.0)
    v6 := make(map[string]int)
    var v22 uint32 = (v2 - v2)
    switch -9223372036854775808 {
    case -9223372036854775808:
        switch -733 {
        case 745:
            v23 := GenS1{F0: -1, F1: v2, F2: v4}
            _ = v23
        case 42:
            switch ("" + "hello") {
            case "\n\t":
                switch int(v4) {
                case 447:
                    delete(v6, ("\n\t" + (("\n\t" + "a") + "a")))
                case 9223372036854775807:
                    for fi24 := 0; fi24 < 4; fi24++ {
                        switch fi24 {
                        case 199:
                            v25 := []int{}
                            _ = v25
                        case 1:
                            v3 = v3
                        case 2147483647:
                            switch "" {
                            case "foo bar":
                                var v26 uint8 = 162
                                _ = v26
                            }
                        default:
                            var v27 int64 = 0
                            _ = v27
                        }
                        switch (fi24 * fi24) {
                        case 1:
                            v4 = v4
                        case 4294967295:
                            var v28 uint = 2
                            _ = v28
                        }
                        _ = fi24
                    }
                case 976:
                    v5 = 0.0
                }
            case "a":
                v29 := GenS1{F0: 42, F1: (v22 << ((646) & 63)), F2: v4}
                _ = v29
            case "":
                for fi30 := 0; fi30 < 8; fi30++ {
                    fmt.Println(fi30, fi30, v5)
                    if (!(!false)) {
                        var v31 float32 = v1
                        for fi32 := 0; fi32 < 18; fi32++ {
                            if (genFn0() == 0) {
                                var v33 bool = true
                                _ = v33
                            }
                            if true {
                                v34 := GenS2{F0: -32768}
                                _ = v34
                            }
                            v2 = (v2 + v2)
                            v4 = v4
                            _ = fi32
                        }
                        _ = v31
                    } else {
                        var v35 int = fi30
                        switch "a" {
                        case "hello":
                            v2 = v22
                        }
                        _ = v35
                    }
                    _ = fi30
                }
            }
        default:
            v4 = int64(v22)
        }
    }
    for fi36 := 0; fi36 < 4; fi36++ {
        v37 := GenS1{F0: 42, F1: (v2 >> ((2) & 63)), F2: (v4 % (((v4 | (v4 * (v4 >> ((255) & 63))))) | 1))}
        var v38 uint64 = 42
        _ = fi36
        _ = v37
        _ = v38
    }
    v39 := []int{0, -182}
    for fi40 := 0; fi40 < 9; fi40++ {
        v1 += float32(1.5)
        v41 := GenS3{F0: (0 < 1000), F1: 42}
        _ = fi40
        _ = v41
    }
    for fi42 := 0; fi42 < 1; fi42++ {
        if false {
            v22 = (v22 ^ (^(v22)))
            var v43 string = "hello"
            v43 = "a"
            _ = v43
        }
        fmt.Println(v1, v2)
        for fi44 := 0; fi44 < 3; fi44++ {
            v22 = v2
            var v45 uint32 = v22
            _ = fi44
            _ = v45
        }
        _ = fi42
    }
    var v46 int32 = -2147483648
    var v47 float64 = 0.0
    for fi48 := 0; fi48 < 11; fi48++ {
        v5 -= 1.5
        _ = fi48
    }
    var v49 int32 = v46
    v50 := v6[("a" + ("\n\t" + ""))]
    for fi51 := 0; fi51 < 4; fi51++ {
        v2 += (v22 | (v22 - (v2 | (^(v2)))))
        fmt.Println(v1, v47, v22)
        _ = fi51
    }
    if false {
        v52 := GenS4{F0: 4294967295}
        var v53 int32 = -1
        var v54 float32 = 1e10
        _ = v52
        _ = v53
        _ = v54
    } else {
        fmt.Println(v50)
        var v55 uint8 = 128
        _ = v55
    }
    fmt.Println(v2, v4)
    for fi56 := 0; fi56 < 20; fi56++ {
        v50 += (v50 >> ((0) & 63))
        for fi57 := 0; fi57 < 2; fi57++ {
            _ = len(v6)
            v58 := GenS1{F0: (v49 % ((v49) | 1)), F1: 335, F2: v4}
            var v59 int32 = 0
            _ = fi57
            _ = v58
            _ = v59
        }
        v5 = v5
        _ = fi56
    }
    _ = len(v6)
    var v60 uint32 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v39
    _ = v4
    _ = v46
    _ = v47
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v60
}
