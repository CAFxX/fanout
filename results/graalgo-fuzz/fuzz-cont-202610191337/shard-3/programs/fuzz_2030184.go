package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint }
func genFn0(p0 uint64, p1 uint8) bool {
    if (!false) {
        fmt.Println(p1)
        if (uint64(p1) <= (p0 << ((607) & 63))) {
            switch "" {
            case "foo bar":
                v5 := GenS1{F0: -1, F1: (p1 % ((p1) | 1)), F2: -659}
                _ = v5
            case "":
                p0 -= (^(p0))
            default:
                p1 += 74
            }
        }
        p0 += (p0 - 25)
    }
    var v6 uint8 = (p1 ^ (p1 * p1))
    _ = p0
    _ = p1
    _ = v6
    return (!true)
}

func main() {
    var v1 int64 = -1
    var v2 int = -1
    var v3 int16 = -1
    var v4 int8 = 86
    v7 := GenS2{F0: 0, F1: 42, F2: 64}
    fmt.Println(v2, v1, v2)
    if true {
        var v8 uint16 = 255
        _ = v8
    } else {
        v9 := []int{v2, (v2 * (v2 + (v2 / (((v2 & (v2 >> ((18446744073709551615) & 63)))) | 1)))), v2}
        v10 := GenS2{F0: 949, F1: 817, F2: 32767}
        v4 += (v4 << ((634) & 63))
        _ = v10
        _ = v9
    }
    var v11 float64 = 0.0
    if false {
        v4 = -1
        switch "\n\t" {
        case "":
            v12 := []int{}
            _ = v12
        case "hello":
            v13 := GenS1{F0: 228, F1: 255, F2: -1}
            _ = v13
        case "a":
            for fi14 := 0; fi14 < 10; fi14++ {
                fmt.Println(fi14, v11, v11)
                var v15 int64 = v1
                if false {
                    if (false && false) {
                        var v16 float64 = v11
                        v11 = v16
                        _ = v16
                    } else {
                        var v17 uint16 = 65535
                        _ = v17
                    }
                    switch (v2 / (((^(v2))) | 1)) {
                    case 0:
                        fmt.Println(v3)
                    case 32767:
                        for fi18 := 0; fi18 < 16; fi18++ {
                            v11 = v11
                            if false {
                                var v19 string = ((("hello" + (("foo bar" + "hello") + "")) + ("" + ("\n\t" + ""))) + ("" + ("" + (("a" + "hello") + "foo bar"))))
                                _ = v19
                            }
                            v20 := []int{418, (v2 & (fi14 / (((-(fi18))) | 1))), fi18, (fi18 * fi14)}
                            _ = fi18
                            _ = v20
                        }
                    }
                }
                _ = fi14
                _ = v15
            }
        }
    } else {
        switch 1 {
        case 9223372036854775807:
            var v21 int = v2
            _ = v21
        case 0:
            for fi22 := 0; fi22 < 12; fi22++ {
                v23 := GenS2{F0: 576, F1: 1, F2: 1}
                switch fi22 {
                case -1:
                    for fi24 := 0; fi24 < 20; fi24++ {
                        switch 42 {
                        case -143:
                            fmt.Println(v1, fi24)
                        }
                        var v25 uint8 = 175
                        var v26 int8 = v4
                        v3 = (v3 * 42)
                        v27 := []int{4294967295, (fi24 - (fi24 % ((fi22) | 1))), 127}
                        _ = fi24
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                case 128:
                    fmt.Println(v3)
                }
                v28 := GenS2{F0: 42, F1: 441, F2: 255}
                _ = fi22
                _ = v23
                _ = v28
            }
        case 32767:
            switch "a" {
            case "hello":
                v29 := GenS2{F0: 0, F1: 65535, F2: 65535}
                _ = v29
            }
        default:
            var v30 float64 = v11
            _ = v30
        }
    }
    v11 += v11
    fmt.Println(v11, v4, v3)
    fmt.Println(v2)
    for fi31 := 0; fi31 < 13; fi31++ {
        if false {
            v32 := GenS1{F0: -2147483648, F1: 72, F2: -1}
            _ = v32
        }
        _ = fi31
    }
    fmt.Println(v3)
    v2 = (v2 >> ((42) & 63))
    v33 := GenS2{F0: 1, F1: 255, F2: 945}
    v34 := GenS3{F0: 741}
    v35 := GenS3{F0: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v7
}
