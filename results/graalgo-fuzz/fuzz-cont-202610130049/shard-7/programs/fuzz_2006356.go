package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint8; F1 int16; F2 int64 }
func genFn0() int32 {
    var v7 uint64 = 4294967295
    v8 := GenS1{F0: -24, F1: 128}
    var v9 uint16 = 1
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint64 = 42
    var v2 uint64 = (v1 - (v1 ^ 753))
    var v3 int64 = -9223372036854775808
    var v4 uint32 = 4294967295
    var v5 uint32 = (v4 - (v4 * 965))
    var v6 float64 = ((1.5 / 1.5) / (1e10 / 1e10))
    v5 = v4
    for fi10 := 0; fi10 < 9; fi10++ {
        if false {
            var v11 float32 = -2.25
            if false {
                v12 := GenS1{F0: 100, F1: 873}
                _ = v12
            }
            if (0 < -1) {
                if false {
                    switch ("\n\t" + (("foo bar" + "hello") + "hello")) {
                    case "a":
                        v3 += v3
                    case "hello":
                        v13 := []int{}
                        _ = v13
                    case "\n\t":
                        if true {
                            if ((!true) || true) {
                                for fi14 := 0; fi14 < 15; fi14++ {
                                    if true {
                                        v15 := GenS1{F0: 1, F1: (fi14 & (fi10 ^ fi14))}
                                        v3 += (v3 % ((-1) | 1))
                                        _ = v15
                                    }
                                    _ = fi14
                                }
                                v11 += float32((v6 - 1e10))
                            } else {
                                fmt.Println(v5, fi10)
                                v16 := GenS1{F0: 93, F1: int(v5)}
                                _ = v16
                            }
                        }
                    default:
                        v17 := GenS1{F0: 127, F1: fi10}
                        _ = v17
                    }
                    var v18 int8 = -95
                    v6 += v6
                    _ = v18
                }
                v19 := GenS1{F0: 127, F1: 42}
                v20 := GenS1{F0: -1, F1: -768}
                _ = v19
                _ = v20
            }
            _ = v11
        } else {
            v21 := []int{fi10, int(v3)}
            fmt.Println(v3, v2, v1)
            var v22 uint8 = 127
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 3; fi23++ {
            v24 := GenS1{F0: -128, F1: fi23}
            fmt.Println(fi23, v5, v6)
            v25 := GenS2{F0: 42, F1: 32767, F2: v3}
            _ = fi23
            _ = v24
            _ = v25
        }
        fmt.Println(v5)
        _ = fi10
    }
    var v26 uint8 = 1
    var v27 int8 = 127
    var v28 uint64 = (^(v2))
    fmt.Println(v6, v5)
    switch -9223372036854775808 {
    case 0:
        var v29 uint64 = (v1 | v2)
        _ = v29
    case 4294967295:
        for fi30 := 0; fi30 < 14; fi30++ {
            if true {
                for fi31 := 0; fi31 < 19; fi31++ {
                    v26 += (v26 ^ 171)
                    _ = fi31
                }
                for fi32 := 0; fi32 < 12; fi32++ {
                    var v33 uint8 = (v26 & (v26 >> ((uint(v2)) & 63)))
                    if true {
                        v34 := []int{}
                        _ = v34
                    }
                    fmt.Println(v5)
                    _ = fi32
                    _ = v33
                }
            } else {
                v5 -= 4294967295
                switch fi30 {
                case 42:
                    if false {
                        var v35 bool = (!true)
                        v36 := []int{fi30, 179, 4294967295, fi30}
                        if true {
                            if v35 {
                                v2 = v1
                            } else {
                                if (v26 <= (^(v26))) {
                                    if ((!(!(v35 && ((42 == 2147483647) || v35)))) && v35) {
                                        v36 = append(v36, fi30)
                                        v35 = (int16(v27) == 42)
                                        v26 += v26
                                    }
                                }
                                v37 := GenS1{F0: v27, F1: fi30}
                                _ = v37
                            }
                            if len(v36) > 0 {
                                v36[0] = (fi30 & fi30)
                            }
                            fmt.Println(v3, v1, v3)
                        } else {
                            v6 = (0.1 + v6)
                            v38 := GenS2{F0: (^(v26)), F1: -32768, F2: (v3 | 9223372036854775807)}
                            v4 -= v5
                            _ = v38
                        }
                        _ = v35
                        _ = v36
                    } else {
                        var v39 int = fi30
                        v27 = 1
                        if (false || (1e10 < 1.5)) {
                            v40 := []int{(v39 % ((fi30) | 1))}
                            v27 = (v27 << ((18446744073709551615) & 63))
                            _ = v40
                        } else {
                            v41 := GenS2{F0: (v26 % ((v26) | 1)), F1: -32768, F2: -465}
                            fmt.Println(v3, v39, v39)
                            fmt.Println(v28)
                            _ = v41
                        }
                        _ = v39
                    }
                case 582:
                    var v42 uint16 = 0
                    _ = v42
                default:
                    var v43 float32 = (0.0 * 0.0)
                    _ = v43
                }
                var v44 int8 = 1
                _ = v44
            }
            _ = fi30
        }
    case -1:
        v1 -= v28
    default:
        var v45 uint32 = v5
        _ = v45
    }
    v46 := []int{int(v26), 127, 128, -39}
    if len(v46) > 0 {
        v46[0] = 0
    }
    v6 = float64(0.0)
    switch 269 {
    case 0:
        switch -1 {
        case 1:
            v3 = (v3 & v3)
        }
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
