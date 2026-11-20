package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 string }
func genFn0() int16 {
    switch "foo bar" {
    case "\n\t":
        var v7 uint8 = 1
        _ = v7
    case "hello":
        if true {
            v8 := GenS1{F0: -1, F1: "a"}
            _ = v8
        } else {
            var v9 float32 = ((0.1 / 1.0) - (0.0 / ((1e10 + 0.0) * 0.1)))
            var v10 uint64 = 18446744073709551615
            _ = v10
            _ = v9
        }
    case "":
        for fi11 := 0; fi11 < 14; fi11++ {
            fmt.Println(fi11)
            _ = fi11
        }
    }
    v12 := []int{0, 42, -349}
    _ = v12
    return 754
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 3.14159
    var v3 uint64 = 476
    var v4 float32 = v2
    v5 := make(map[uint32]string)
    var v6 int = -1
    v13 := []int{v6, v6, 322, 1}
    if true {
        v13 = append(v13, v6)
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 uint = 0
            switch -65 {
            case 128:
                v13 = append(v13, (v6 ^ (fi14 % ((v6) | 1))))
            case 9223372036854775807:
                switch "a" {
                case "\n\t":
                    _ = len(v13)
                default:
                    v16 := GenS1{F0: -88, F1: "foo bar"}
                    _ = v16
                }
            }
            if false {
                switch fi14 {
                case 127:
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                default:
                    switch 4294967295 {
                    case -430:
                        if false {
                            for fi17 := 0; fi17 < 4; fi17++ {
                                var v18 float32 = v2
                                var v19 bool = true
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                            var v20 bool = true
                            _ = v20
                        }
                    case 158:
                        v2 = v2
                    }
                }
                for fi21 := 0; fi21 < 14; fi21++ {
                    v6 -= -2147483648
                    if ((-2.25 - -2.25) <= (0.0 - float64(((((v2 / ((1.5 * float32((-2.25 * 0.0))) / v4)) - (v4 / ((v4 * 3.14159) * v2))) + (v2 * v2)) + 0.0)))) {
                        v3 = (v3 - 60)
                        if ((v15 - v15) >= v15) {
                            v22 := GenS2{F0: 255, F1: 891}
                            fmt.Println(fi21, v4)
                            _ = v22
                        }
                    }
                    v2 = (1.5 / v4)
                    fmt.Println(v6, v4, v3)
                    _ = fi21
                }
            }
            _ = fi14
            _ = v15
        }
    }
    var v23 float32 = 3.14159
    v13 = append(v13, v6)
    fmt.Println(v23, v4, v6)
    fmt.Println(v3, v23, v4)
    v2 += v23
    if true {
        switch ("foo bar" + "") {
        case "a":
            if true {
                for fi24 := 0; fi24 < 15; fi24++ {
                    switch "\n\t" {
                    case "":
                        switch fi24 {
                        case 32768:
                            v6 = v6
                        }
                    case "hello":
                        v6 += (v6 << ((255) & 63))
                    case "\n\t":
                        for fi25 := 0; fi25 < 3; fi25++ {
                            v26 := GenS1{F0: 1, F1: (("" + "a") + "")}
                            fmt.Println(v6, v2, v6)
                            _ = fi25
                            _ = v26
                        }
                    }
                    _ = fi24
                }
                for fi27 := 0; fi27 < 3; fi27++ {
                    var v28 int16 = 0
                    _ = fi27
                    _ = v28
                }
            }
        default:
            var v29 int = v6
            _ = v29
        }
        v30 := GenS3{F0: (v3 ^ (v3 & v3)), F1: false}
        _ = v30
    } else {
        v6 += 127
        if ((((v4 * (v2 - v23)) * 0.0) + (v4 + (0.1 - (1.5 * ((float32(((-2.25 + -2.25) + 1.5)) - 3.14159) + ((0.0 + 0.1) + v2)))))) > (float32(0.0) / 1.0)) {
            switch v6 {
            case 481:
                var v31 uint64 = v3
                _ = v31
            case -2147483648:
                for fi32 := 0; fi32 < 14; fi32++ {
                    v33 := GenS2{F0: 106, F1: -9223372036854775808}
                    _ = len(v13)
                    if (-2147483648 <= 42) {
                        switch "hello" {
                        case "a":
                            var v34 float32 = v4
                            _ = v34
                        }
                        v23 -= float32(1e10)
                        switch (v6 ^ fi32) {
                        case -1:
                            var v35 uint8 = 96
                            _ = v35
                        case -442:
                            var v36 int8 = -1
                            _ = v36
                        case 128:
                            v23 = v2
                        default:
                            v4 -= (v23 * 1e10)
                        }
                    }
                    _ = fi32
                    _ = v33
                }
            }
            v37 := GenS2{F0: 255, F1: 364}
            _ = v37
        } else {
            v13 = append(v13, v6)
        }
    }
    v4 += v23
    switch (v6 * (v6 & (^(v6)))) {
    case -134:
        v23 = 1e10
    case 0:
        fmt.Println(v6, v4)
    default:
        fmt.Println(v23, v2)
    }
    v38 := v1[901]
    v39 := GenS4{F0: (("foo bar" + "hello") + "a")}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
