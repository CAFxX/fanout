package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 uint16; F1 uint }
func genFn0(p0 uint16) uint64 {
    switch (("hello" + "") + ("foo bar" + "a")) {
    case "":
        fmt.Println(p0)
    case "hello":
        p0 += p0
    case "a":
        for fi6 := 0; fi6 < 13; fi6++ {
            var v7 int = (fi6 % ((42) | 1))
            _ = fi6
            _ = v7
        }
    default:
        v8 := []int{-2147483648, -854}
        _ = v8
    }
    if (true && (((2 <= 2) || false) && false)) {
        if (p0 != (p0 ^ (p0 ^ p0))) {
            fmt.Println(p0)
        } else {
            v9 := []int{0, -345, -551, 534}
            var v10 bool = (!true)
            for fi11 := 0; fi11 < 16; fi11++ {
                switch fi11 {
                case -2147483648:
                    v10 = true
                }
                _ = fi11
            }
            _ = v10
            _ = v9
        }
        p0 = 32767
        switch "" {
        case "\n\t":
            p0 = (p0 + p0)
        }
    }
    v12 := []int{}
    v13 := GenS1{F0: 255, F1: 2}
    p0 -= p0
    _ = p0
    _ = v12
    _ = v13
    return 0
}

func main() {
    var v1 bool = true
    var v2 string = "\n\t"
    var v3 uint64 = 42
    v4 := make(map[int]uint32)
    var v5 float32 = 0.1
    switch v2 {
    case "\n\t":
        switch v2 {
        case "a":
            for fi14 := 0; fi14 < 11; fi14++ {
                v15 := GenS1{F0: 1, F1: 107}
                v16 := []int{32767, -178, fi14}
                var v17 uint64 = v3
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
        }
    case "hello":
        switch v2 {
        case "hello":
            v3 = v3
        case "\n\t":
            v1 = (v5 < ((float32(0.1) - v5) + v5))
        case "a":
            v2 = (v2 + (v2 + "hello"))
        }
    default:
        if (((0.1 / v5) * v5) < float32(1e10)) {
            fmt.Println(v5, v2)
        }
    }
    if v1 {
        if (((1 != 65535) || ((v5 + v5) >= (v5 + v5))) && v1) {
            var v18 bool = (65535 > 255)
            _ = v18
        } else {
            for fi19 := 0; fi19 < 14; fi19++ {
                var v20 string = v2
                _ = fi19
                _ = v20
            }
        }
    } else {
        var v21 int = 0
        _ = v21
    }
    delete(v4, -308)
    switch v2 {
    case "a":
        if v1 {
            if true {
                v1 = false
                var v22 uint32 = 2147483648
                _ = v22
            }
            var v23 float64 = 1.5
            v24 := GenS1{F0: 2, F1: 32767}
            _ = v23
            _ = v24
        } else {
            v3 -= 707
        }
    case "\n\t":
        fmt.Println(v1)
    default:
        v25 := GenS1{F0: 186, F1: 255}
        _ = v25
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        fmt.Println(v3)
        _ = fi26
    }
    v2 = ((("a" + "") + v2) + "foo bar")
    var v27 int64 = -1
    var v28 float32 = float32(((float64((v5 / v5)) * (1e10 + float64(v5))) - 0.0))
    if false {
        v5 -= 0.1
        fmt.Println(v2, v5, v3)
        for fi29 := 0; fi29 < 6; fi29++ {
            v3 = 117
            for fi30 := 0; fi30 < 5; fi30++ {
                var v31 string = v2
                switch fi29 {
                case -2147483648:
                    var v32 bool = v1
                    _ = v32
                case 128:
                    if (!(0.0 != (float64(float32(-2.25)) * 0.0))) {
                        var v33 uint64 = 1
                        var v34 uint = uint(v33)
                        if ((1.5 * (float32(1e10) / (v28 * 0.1))) <= 0.0) {
                            v2 = ((v2 + "hello") + "hello")
                            v35 := GenS2{F0: (fi29 % (((fi29 >> (((v34 - (v34 + 458))) & 63))) | 1)), F1: 255, F2: 1}
                            _ = v35
                        } else {
                            v36 := GenS1{F0: 46, F1: 0}
                            switch (v31 + "foo bar") {
                            case "\n\t":
                                var v37 uint64 = v3
                                _ = v37
                            case "foo bar":
                                v4[fi29] = 190
                            case "":
                                v3 = 458
                                var v38 float64 = ((((float64(v28) * 0.1) + 0.0) / (0.1 / 3.14159)) + -2.25)
                                _ = v38
                            default:
                                v33 = (v3 ^ genFn0(uint16(v34)))
                            }
                            v39 := []int{}
                            _ = v36
                            _ = v39
                        }
                        _ = v33
                        _ = v34
                    }
                case 9223372036854775807:
                    var v40 int8 = 88
                    _ = v40
                }
                _ = fi30
                _ = v31
            }
            v41 := GenS1{F0: 86, F1: 65535}
            _ = fi29
            _ = v41
        }
    } else {
        var v42 uint16 = 32767
        var v43 float64 = (0.0 - 3.14159)
        var v44 uint32 = 2147483647
        _ = v42
        _ = v43
        _ = v44
    }
    v45 := GenS2{F0: -9223372036854775808, F1: 10, F2: 2}
    v46 := GenS3{F0: 134, F1: (v2 + (v2 + ("\n\t" + "hello")))}
    v47 := GenS4{F0: 255, F1: 2}
    fmt.Println(v5)
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
}
