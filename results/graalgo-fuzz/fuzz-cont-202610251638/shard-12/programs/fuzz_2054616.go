package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 uint32; F2 int }
func genFn0(p0 int8, p1 bool, p2 int64) int8 {
    v4 := []int{-806, 1, 2147483647, 2147483647}
    var v5 int64 = p2
    p0 = (p0 + (p0 % (((p0 << ((4294967295) & 63))) | 1)))
    switch 4294967295 {
    case 0:
        if (!(!(-486 > -1))) {
            var v6 int8 = p0
            v5 -= p2
            _ = v6
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 100
}
func genFn1(p0 uint) uint {
    v7 := []int{50}
    p0 -= p0
    v8 := GenS1{F0: 0}
    var v9 int64 = 9223372036854775807
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (^(p0))
}

func main() {
    v1 := []int{-9223372036854775808, 487, 32767}
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]int)
    if (42 == 1) {
        var v10 int16 = -1
        _ = v10
    } else {
        var v11 int = 128
        _ = v11
    }
    var v12 uint8 = 127
    v12 = 227
    var v13 uint16 = 0
    var v14 uint = 553
    v14 += v14
    _ = len(v2)
    var v15 int64 = -9223372036854775808
    v15 += v15
    if (526 <= uint64(v13)) {
        for fi16 := 0; fi16 < 1; fi16++ {
            if false {
                switch (fi16 % ((fi16) | 1)) {
                case 42:
                    v17 := GenS1{F0: 42}
                    _ = v17
                case 1:
                    if (!false) {
                        var v18 int16 = -636
                        switch "hello" {
                        case "hello":
                            var v19 float32 = 0.0
                            _ = v19
                        case "foo bar":
                            v20 := GenS1{F0: uint64(v13)}
                            _ = v20
                        case "a":
                            if true {
                                if true {
                                    v21 := GenS1{F0: 1}
                                    fmt.Println(v18, fi16, v18)
                                    var v22 int64 = v15
                                    _ = v21
                                    _ = v22
                                }
                            }
                        }
                        var v23 bool = true
                        _ = v18
                        _ = v23
                    }
                }
                var v24 float64 = float64(-2.25)
                _ = v24
            }
            switch (("hello" + "a") + "foo bar") {
            case "foo bar":
                var v25 uint64 = 112
                _ = v25
            case "a":
                if true {
                    switch (fi16 / ((974) | 1)) {
                    case 32767:
                        v26 := GenS2{F0: (v14 & v14), F1: uint32(v13), F2: fi16}
                        _ = v26
                    case 32768:
                        if false {
                            var v27 int32 = -983
                            _ = v27
                        }
                    case 368:
                        v28 := GenS1{F0: 103}
                        _ = v28
                    }
                    if false {
                        switch "foo bar" {
                        case "\n\t":
                            var v29 int8 = -1
                            _ = v29
                        case "foo bar":
                            var v30 int32 = 42
                            _ = v30
                        case "hello":
                            v2[("" + "\n\t")] = 1
                        }
                    } else {
                        _ = len(v2)
                        var v31 int64 = v15
                        switch ("foo bar" + ("" + ("\n\t" + ("\n\t" + "foo bar")))) {
                        case "":
                            switch fi16 {
                            case -526:
                                v13 -= (v13 * 255)
                            case 32767:
                                if true {
                                    var v32 float64 = ((0.1 / (((1e10 - 3.14159) * float64((-2.25 * 0.1))) + float64(1.5))) + ((1e10 / 0.1) * (1e10 / 1e10)))
                                    _ = v32
                                } else {
                                    fmt.Println(v13)
                                    var v33 uint = (v14 + 255)
                                    _ = v33
                                }
                            case 0:
                                v13 += (v13 ^ v13)
                            default:
                                v13 -= (v13 ^ (v13 & 32767))
                            }
                        default:
                            var v34 bool = (!(471 < 42))
                            _ = v34
                        }
                        _ = v31
                    }
                }
            case "":
                var v35 string = "\n\t"
                _ = v35
            default:
                v15 = (v15 + (v15 >> (((v14 << (((v14 % (((v14 % ((0) | 1))) | 1))) & 63))) & 63)))
            }
            _ = fi16
        }
        for fi36 := 0; fi36 < 11; fi36++ {
            for fi37 := 0; fi37 < 18; fi37++ {
                v15 += (v15 / ((0) | 1))
                _ = fi37
            }
            _ = fi36
        }
    }
    v14 -= v14
    var v38 int = 32767
    if ("\n\t" <= "hello") {
        var v39 uint64 = 0
        _ = v39
    }
    fmt.Println(v15)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v38
}
