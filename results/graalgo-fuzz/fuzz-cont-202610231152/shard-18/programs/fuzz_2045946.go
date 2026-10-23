package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint64; F2 int32 }
type GenS3 struct { F0 int16; F1 int32; F2 uint }
func genFn0(p0 bool, p1 int32) uint64 {
    p0 = false
    p0 = (!(375 != 375))
    _ = p0
    _ = p1
    return 4294967295
}
func genFn1(p0 int, p1 uint64) int8 {
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 uint = 65535
        _ = fi5
        _ = v6
    }
    p1 = (^(p1))
    _ = p0
    _ = p1
    return 127
}

func main() {
    v1 := []int{0, 212, -758, 9223372036854775807}
    v2 := make(map[uint32]bool)
    var v3 uint = 255
    v4 := make(map[string]bool)
    v3 -= v3
    v3 += v3
    var v7 uint = v3
    _ = len(v1)
    v7 -= v3
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            if len(v1) > 0 {
                v1[0] = -78
            }
            switch -703 {
            case -2147483648:
                fmt.Println(fi9, fi9, fi8)
            default:
                fmt.Println(v7, v3, v7)
            }
            fmt.Println(v3, v7, v3)
            _ = fi9
        }
        _ = fi8
    }
    if (2 > 255) {
        var v10 float32 = 3.14159
        v7 -= (^(v7))
        fmt.Println(v3, v7)
        _ = v10
    } else {
        if (((^(v7)) > v7) && false) {
            fmt.Println(v7, v3)
            var v11 int64 = -1
            _ = v11
        } else {
            switch -2147483648 {
            case -9223372036854775808:
                v3 -= v3
            }
            switch "foo bar" {
            case "":
                for fi12 := 0; fi12 < 1; fi12++ {
                    v3 += 1
                    fmt.Println(v7, v7, v7)
                    fmt.Println(v7)
                    _ = fi12
                }
            case "hello":
                v3 = (^(v7))
            case "foo bar":
                if (!(128 == -341)) {
                    var v13 uint32 = 2147483648
                    _ = len(v1)
                    _ = v13
                }
            }
        }
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 int16 = -405
        fmt.Println(fi14)
        _ = fi14
        _ = v15
    }
    var v16 int8 = 100
    v17 := GenS1{F0: (("" + ("a" + "")) + ""), F1: 2}
    for fi18 := 0; fi18 < 12; fi18++ {
        switch (fi18 ^ (fi18 % ((fi18) | 1))) {
        case 32768:
            if false {
                v3 += 0
            } else {
                if (32767 >= -32768) {
                    v19 := v2[2147483647]
                    if (!false) {
                        v19 = (v19 || ((v3 >> (((v3 * v3)) & 63)) != v3))
                        var v20 float64 = 1e10
                        _ = v20
                    }
                    if v19 {
                        if (2 == 1) {
                            v21 := GenS2{F0: (genFn1(fi18, 332) >= v16), F1: genFn0((float64((-2.25 / (3.14159 - -2.25))) > 0.1), -176), F2: 0}
                            v22 := v2[1]
                            _ = v21
                            _ = v22
                        } else {
                            v23 := GenS3{F0: 0, F1: 2147483647, F2: 918}
                            _ = v23
                        }
                    }
                    _ = v19
                } else {
                    if false {
                        if ((!false) && ((v3 << ((366) & 63)) < v3)) {
                            v16 += v16
                        } else {
                            var v24 float64 = 1.5
                            _ = v24
                        }
                    }
                }
            }
        case -70:
            switch ((("" + "foo bar") + "a") + "foo bar") {
            case "hello":
                var v25 float32 = -2.25
                _ = v25
            case "\n\t":
                v26 := GenS1{F0: "foo bar", F1: 255}
                _ = v26
            default:
                v7 -= v7
            }
        }
        switch ("hello" + "foo bar") {
        case "hello":
            var v27 int16 = -1
            _ = v27
        default:
            for fi28 := 0; fi28 < 18; fi28++ {
                if (0 == -976) {
                    if (559 >= -616) {
                        var v29 int = -521
                        v7 = v3
                        _ = v29
                    }
                }
                var v30 int32 = -2147483648
                _ = fi28
                _ = v30
            }
        }
        _ = fi18
    }
    var v31 string = ((((("\n\t" + "foo bar") + ((("hello" + (("" + "hello") + "foo bar")) + "") + "\n\t")) + "") + "\n\t") + ("a" + ""))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v7
}
