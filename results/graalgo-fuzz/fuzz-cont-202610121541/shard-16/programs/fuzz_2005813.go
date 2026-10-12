package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int32 }
func genFn0() bool {
    var v7 int64 = -9223372036854775808
    if false {
        fmt.Println(v7)
    } else {
        for fi8 := 0; fi8 < 1; fi8++ {
            v9 := GenS1{F0: fi8}
            _ = fi8
            _ = v9
        }
    }
    _ = v7
    return ((105 >= 18446744073709551615) || true)
}
func genFn1(p0 bool) int32 {
    p0 = (p0 || p0)
    var v10 int = 4294967295
    p0 = (!true)
    _ = p0
    _ = v10
    return 1
}

func main() {
    var v1 int16 = 22
    var v2 uint8 = 128
    var v3 int = 1
    var v4 string = "\n\t"
    var v5 float32 = 0.1
    var v6 uint8 = (v2 << ((255) & 63))
    v1 -= (^(v1))
    v3 = v3
    v11 := GenS1{F0: v3}
    v2 = v6
    if (!false) {
        if true {
            switch v4 {
            case "":
                for fi12 := 0; fi12 < 10; fi12++ {
                    v13 := []int{(fi12 * v3), 42}
                    var v14 string = v4
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            case "a":
                fmt.Println(v3, v2, v4)
            case "\n\t":
                v5 -= v5
            default:
                var v15 float64 = (-2.25 / 0.1)
                _ = v15
            }
            fmt.Println(v2)
        }
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 bool = true
        v18 := []int{32768, (v3 & (fi16 - (v3 * v3)))}
        v4 = "foo bar"
        _ = fi16
        _ = v17
        _ = v18
    }
    switch 128 {
    case 32768:
        if (113 < 18446744073709551615) {
            switch v3 {
            case 0:
                for fi19 := 0; fi19 < 6; fi19++ {
                    v20 := GenS2{F0: 0, F1: (v1 << ((104) & 63))}
                    _ = fi19
                    _ = v20
                }
            case -846:
                var v21 float32 = v5
                _ = v21
            case 32767:
                var v22 float32 = 0.1
                _ = v22
            }
        }
    case 128:
        v23 := []int{(v3 | (^(v3))), v3, 42}
        _ = v23
    default:
        v24 := GenS3{F0: int32(v1)}
        _ = v24
    }
    var v25 int8 = 1
    switch v4 {
    case "":
        if true {
            v5 = float32((0.1 + (1e10 / float64(((v5 / float32(((1.5 * ((-2.25 - 1.5) - 1.5)) / -2.25))) / float32(-2.25))))))
        }
    case "foo bar":
        var v26 uint64 = 1
        _ = v26
    case "hello":
        v2 = v6
    }
    for fi27 := 0; fi27 < 14; fi27++ {
        fmt.Println(v5, v2, v3)
        var v28 int = (fi27 * (v3 >> ((2) & 63)))
        switch (v28 & v28) {
        case 9223372036854775807:
            if (2147483648 == 1) {
                fmt.Println(fi27, v25)
                v29 := GenS3{F0: -1}
                for fi30 := 0; fi30 < 2; fi30++ {
                    v25 = (v25 * v25)
                    if true {
                        v1 = (v1 ^ (^(v1)))
                        fmt.Println(fi27, v5)
                    }
                    _ = fi30
                }
                _ = v29
            }
        case 529:
            if true {
                switch (v28 * fi27) {
                case 286:
                    v2 = v2
                case 32768:
                    v5 += v5
                }
                var v31 string = v4
                v32 := []int{-960, (v3 | v3), fi27}
                _ = v31
                _ = v32
            }
        case 1:
            if false {
                switch (v28 % ((-726) | 1)) {
                case 682:
                    if ((v6 & 1) > 1) {
                        v1 += 0
                        for fi33 := 0; fi33 < 10; fi33++ {
                            v1 += 32767
                            if (740 <= 0) {
                                v25 += v25
                                v34 := GenS1{F0: (fi33 - v3)}
                                if (!true) {
                                    v1 -= -1
                                } else {
                                    var v35 bool = true
                                    fmt.Println(v25, v2)
                                    var v36 int8 = v25
                                    _ = v35
                                    _ = v36
                                }
                                v4 = ("a" + "foo bar")
                                _ = v34
                            }
                            _ = fi33
                        }
                    } else {
                        var v37 bool = false
                        _ = v37
                    }
                default:
                    for fi38 := 0; fi38 < 8; fi38++ {
                        v6 = 161
                        _ = fi38
                    }
                }
                v6 = v2
            }
        default:
            var v39 int64 = -434
            _ = v39
        }
        _ = fi27
        _ = v28
    }
    v4 = "foo bar"
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
