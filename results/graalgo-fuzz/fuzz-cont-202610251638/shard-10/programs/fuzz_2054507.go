package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint64 }
func genFn0() int64 {
    var v5 uint32 = 2147483648
    var v6 bool = ((0 == 42) || false)
    var v7 int16 = -1
    var v8 int = 32767
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 uint64 = 42
        if false {
            switch -9223372036854775808 {
            case 42:
                v11 := GenS1{F0: 107, F1: -128, F2: 42}
                _ = v11
            default:
                v12 := []int{}
                _ = v12
            }
            switch (("\n\t" + "hello") + "\n\t") {
            case "a":
                if v6 {
                    v13 := []int{(fi9 >> ((18446744073709551615) & 63)), (v8 >> ((367) & 63))}
                    v5 += v5
                    if (!(128 <= 87)) {
                        for fi14 := 0; fi14 < 16; fi14++ {
                            var v15 uint16 = 591
                            _ = fi14
                            _ = v15
                        }
                        v16 := GenS1{F0: -128, F1: 100, F2: 0}
                        v6 = v6
                        _ = v16
                    } else {
                        _ = len(v13)
                        for fi17 := 0; fi17 < 20; fi17++ {
                            _ = len(v13)
                            if (!v6) {
                                var v18 int64 = -9223372036854775808
                                v10 += (^(v10))
                                _ = v18
                            } else {
                                fmt.Println(fi9, v8)
                                switch 2147483647 {
                                case 0:
                                    var v19 int = fi17
                                    _ = v19
                                }
                            }
                            fmt.Println(v5, fi9)
                            _ = fi17
                        }
                        var v20 int8 = -128
                        _ = v20
                    }
                    _ = v13
                } else {
                    var v21 float32 = (0.1 / 3.14159)
                    _ = v21
                }
            case "hello":
                v22 := GenS1{F0: 1, F1: 127, F2: 42}
                _ = v22
            default:
                v6 = v6
            }
        } else {
            v5 = 940
            for fi23 := 0; fi23 < 2; fi23++ {
                var v24 int64 = 1
                fmt.Println(fi9)
                v24 += int64(v8)
                _ = fi23
                _ = v24
            }
        }
        var v25 uint = 0
        _ = fi9
        _ = v10
        _ = v25
    }
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 bool) uint16 {
    p0 = (42 < 32767)
    var v26 uint = 2
    _ = p0
    _ = v26
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 1
    v3 := make(map[string]bool)
    var v4 float64 = 3.14159
    v27 := GenS1{F0: 100, F1: 0, F2: 64}
    fmt.Println(v4, v2)
    for fi28 := 0; fi28 < 6; fi28++ {
        var v29 float32 = 1e10
        if false {
            v2 += -1
        }
        var v30 float32 = v29
        _ = fi28
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 18; fi31++ {
        v2 -= v2
        _ = fi31
    }
    v4 = ((v4 / v4) / 0.1)
    v32 := []int{800, 1, 4294967295, 830}
    v33 := GenS1{F0: 100, F1: 1, F2: 2}
    if len(v32) > 0 {
        v32[0] = 32767
    }
    var v34 string = "a"
    switch v34 {
    case "\n\t":
        _ = len(v1)
    }
    var v35 int32 = 1
    v35 += v35
    v34 = v34
    v1[433] = 2147483647
    if len(v32) >= 2 {
        v32 = v32[1:len(v32)]
    }
    if (!true) {
        v36 := GenS1{F0: -1, F1: -128, F2: 1}
        v35 = (v35 | v35)
        _ = v36
    } else {
        var v37 int64 = (v2 - genFn0())
        _ = v37
    }
    switch v34 {
    case "a":
        for fi38 := 0; fi38 < 3; fi38++ {
            if (0 >= fi38) {
                if true {
                    if false {
                        var v39 int8 = 42
                        _ = v39
                    }
                    if true {
                        delete(v1, 644)
                    }
                    var v40 bool = true
                    _ = v40
                }
                for fi41 := 0; fi41 < 8; fi41++ {
                    if len(v32) > 0 {
                        v32[0] = fi41
                    }
                    _ = fi41
                }
                _ = len(v3)
            }
            v42 := GenS1{F0: -95, F1: 1, F2: 18446744073709551615}
            _ = fi38
            _ = v42
        }
    default:
        var v43 uint8 = 231
        _ = v43
    }
    v44 := GenS1{F0: -1, F1: 27, F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v44
}
