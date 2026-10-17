package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[string]bool)
    var v3 int64 = -430
    var v4 int64 = (-(v3))
    var v5 uint32 = 65535
    v6 := make(map[uint32]uint32)
    v1 = (v1 * v1)
    var v7 float64 = (3.14159 + float64(1e10))
    if (42 < 1) {
        v4 = v4
        if (!false) {
            v8 := GenS1{F0: v3, F1: 182}
            var v9 int32 = (v1 >> ((2) & 63))
            _ = v8
            _ = v9
        }
        if false {
            v6[4294967295] = (v5 | 322)
            var v10 uint64 = 780
            _ = v10
        } else {
            for fi11 := 0; fi11 < 15; fi11++ {
                if true {
                    v12 := GenS1{F0: (v4 ^ 0), F1: 65}
                    _ = v12
                } else {
                    if false {
                        v1 += v1
                        for fi13 := 0; fi13 < 1; fi13++ {
                            v3 += v4
                            v14 := []int{(fi11 * fi13)}
                            _ = fi13
                            _ = v14
                        }
                        v15 := []int{-26, 756, fi11, 1}
                        _ = v15
                    } else {
                        switch "hello" {
                        case "\n\t":
                            if false {
                                if (1.5 >= (float64(0.1) / v7)) {
                                    var v16 bool = (true || true)
                                    _ = v16
                                } else {
                                    for fi17 := 0; fi17 < 10; fi17++ {
                                        v18 := v2[("a" + "a")]
                                        v19 := GenS1{F0: v4, F1: 128}
                                        v3 += -375
                                        _ = fi17
                                        _ = v18
                                        _ = v19
                                    }
                                    if true {
                                        v5 = (^(v5))
                                        switch (fi11 ^ fi11) {
                                        case 260:
                                            var v20 int = fi11
                                            _ = v20
                                        case 128:
                                            fmt.Println(v4, v1, v5)
                                        default:
                                            switch fi11 {
                                            case 0:
                                                v4 += v3
                                            case 161:
                                                var v21 uint8 = 0
                                                _ = v21
                                            }
                                        }
                                    } else {
                                        for fi22 := 0; fi22 < 2; fi22++ {
                                            v23 := []int{fi22, fi11, fi11, (fi22 + (fi22 - fi22))}
                                            if len(v23) >= 2 {
                                                v23 = v23[1:len(v23)]
                                            }
                                            _ = fi22
                                            _ = v23
                                        }
                                        for fi24 := 0; fi24 < 11; fi24++ {
                                            var v25 uint = 0
                                            _ = fi24
                                            _ = v25
                                        }
                                        if false {
                                            var v26 uint32 = 1
                                            v1 += 0
                                            _ = v26
                                        }
                                    }
                                }
                                if (false && true) {
                                    var v27 float64 = v7
                                    _ = v27
                                }
                                var v28 uint16 = 32767
                                _ = v28
                            }
                        case "hello":
                            _ = len(v6)
                        case "":
                            v4 += v3
                        default:
                            var v29 uint64 = 901
                            _ = v29
                        }
                        v1 = (^(v1))
                        for fi30 := 0; fi30 < 6; fi30++ {
                            var v31 int8 = 1
                            v1 = v1
                            _ = fi30
                            _ = v31
                        }
                    }
                    v7 = ((-2.25 - v7) + (-2.25 + 0.0))
                    for fi32 := 0; fi32 < 20; fi32++ {
                        switch (fi11 * fi32) {
                        case 0:
                            if true {
                                v5 = 468
                                var v33 int8 = 127
                                if false {
                                    v3 = v3
                                    v4 = v3
                                    var v34 uint = 512
                                    _ = v34
                                }
                                _ = v33
                            }
                        case -9223372036854775808:
                            fmt.Println(fi11, fi11)
                        default:
                            v4 = int64(v1)
                        }
                        v35 := []int{fi32, fi11}
                        _ = fi32
                        _ = v35
                    }
                }
                var v36 int16 = -1
                _ = fi11
                _ = v36
            }
        }
    } else {
        v3 = -1
        v4 += int64(v5)
    }
    var v37 uint8 = 195
    v5 -= v5
    var v38 string = "\n\t"
    v39 := GenS1{F0: -9223372036854775808, F1: (v37 << ((0) & 63))}
    v38 = "hello"
    if ((v5 & (v5 / (((v5 << ((65535) & 63))) | 1))) > 149) {
        var v40 int8 = 0
        for fi41 := 0; fi41 < 6; fi41++ {
            var v42 float32 = 0.1
            v43 := GenS1{F0: 1, F1: v37}
            v44 := []int{fi41}
            _ = fi41
            _ = v42
            _ = v43
            _ = v44
        }
        v45 := GenS2{F0: 1}
        _ = v40
        _ = v45
    }
    delete(v2, v38)
    v46 := []int{178, 42, 4294967295}
    var v47 uint32 = (v5 / ((42) | 1))
    v2[((((v38 + v38) + v38) + "\n\t") + "a")] = ((v37 >> ((42) & 63)) < 170)
    var v48 int32 = 2147483647
    var v49 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
    _ = v7
}
