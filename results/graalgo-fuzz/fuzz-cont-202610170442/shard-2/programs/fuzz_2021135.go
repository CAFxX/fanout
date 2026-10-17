package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint16; F1 uint; F2 int32 }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 uint; F1 uint8 }
func genFn0(p0 int) int16 {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := GenS1{F0: 0, F1: 77, F2: -886}
        if (!((fi5 >> ((1) & 63)) > (p0 << ((0) & 63)))) {
            switch (-(fi5)) {
            case 9223372036854775807:
                if false {
                    v7 := GenS1{F0: 205, F1: 128, F2: -32768}
                    _ = v7
                } else {
                    switch "hello" {
                    case "":
                        if true {
                            var v8 uint8 = 255
                            _ = v8
                        } else {
                            p0 = (fi5 ^ -897)
                            p0 = (^(fi5))
                            p0 = -2147483648
                        }
                    case "foo bar":
                        for fi9 := 0; fi9 < 6; fi9++ {
                            var v10 int16 = 896
                            switch (p0 % ((9223372036854775807) | 1)) {
                            case 128:
                                v11 := []int{(-(fi9))}
                                _ = v11
                            case -492:
                                var v12 int32 = 2147483647
                                _ = v12
                            case 0:
                                v13 := []int{2147483647, 517, fi5}
                                _ = v13
                            }
                            if (-96 >= 0) {
                                v14 := []int{42, 2147483647}
                                _ = v14
                            } else {
                                v15 := GenS2{F0: true, F1: (("\n\t" + "\n\t") + "a")}
                                v16 := GenS1{F0: 18, F1: 2, F2: v10}
                                var v17 int16 = v10
                                p0 += (^(fi5))
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                            switch (p0 ^ fi9) {
                            case 42:
                                var v18 uint64 = 347
                                _ = v18
                            case -2147483648:
                                v19 := []int{2147483647, int(v10)}
                                _ = v19
                            }
                            v20 := GenS2{F0: (!true), F1: (((("foo bar" + (("" + ("hello" + "foo bar")) + "")) + "a") + (("" + "a") + "\n\t")) + "foo bar")}
                            _ = fi9
                            _ = v10
                            _ = v20
                        }
                    case "a":
                        var v21 int16 = -327
                        _ = v21
                    default:
                        var v22 int = (fi5 ^ (fi5 % ((p0) | 1)))
                        _ = v22
                    }
                }
            case -2147483648:
                var v23 float32 = 1e10
                _ = v23
            case 2147483647:
                v24 := []int{42, 2147483647, 314, 2147483647}
                _ = v24
            }
            p0 += (fi5 ^ (fi5 / ((fi5) | 1)))
        } else {
            v25 := []int{}
            _ = v25
        }
        _ = fi5
        _ = v6
    }
    _ = p0
    return -253
}

func main() {
    var v1 uint32 = 0
    var v2 uint64 = 857
    var v3 int8 = 22
    var v4 float64 = float64(0.0)
    if (498 < 0) {
        if false {
            v26 := []int{-2147483648, 32767, 1}
            for fi27 := 0; fi27 < 4; fi27++ {
                var v28 uint16 = 1
                v1 -= (^(v1))
                _ = fi27
                _ = v28
            }
            _ = v26
        } else {
            v2 = 321
        }
    }
    switch -249 {
    case 128:
        v29 := GenS1{F0: 255, F1: 78, F2: 762}
        _ = v29
    case 2147483647:
        v30 := GenS3{F0: 65535, F1: 4294967295, F2: -519}
        _ = v30
    case -1:
        v3 -= 0
    default:
        fmt.Println(v1)
    }
    for fi31 := 0; fi31 < 14; fi31++ {
        fmt.Println(v3)
        _ = fi31
    }
    switch "foo bar" {
    case "a":
        if false {
            v2 += 42
            v1 -= uint32(v2)
            var v32 bool = true
            _ = v32
        } else {
            var v33 int = -196
            var v34 int32 = 42
            _ = v33
            _ = v34
        }
    case "hello":
        var v35 uint32 = 65535
        _ = v35
    case "":
        v36 := []int{-2147483648, 32768, -47}
        _ = v36
    default:
        v4 += v4
    }
    var v37 uint = 255
    if (uint32(v37) != (^(v1))) {
        for fi38 := 0; fi38 < 10; fi38++ {
            for fi39 := 0; fi39 < 20; fi39++ {
                v40 := GenS1{F0: 255, F1: 42, F2: 42}
                v41 := GenS4{F0: "foo bar", F1: v2}
                _ = fi39
                _ = v40
                _ = v41
            }
            _ = fi38
        }
        var v42 uint16 = 485
        v1 = (v1 / (((v1 + (v1 & (^(v1))))) | 1))
        _ = v42
    } else {
        var v43 uint32 = v1
        v44 := []int{1, 2147483647}
        v45 := GenS5{F0: (v37 % ((2) | 1)), F1: 42}
        _ = v43
        _ = v44
        _ = v45
    }
    fmt.Println(v37, v1)
    var v46 uint = v37
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v46
}
