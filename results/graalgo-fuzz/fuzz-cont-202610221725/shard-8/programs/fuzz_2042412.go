package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 string; F1 int8; F2 int16 }
func genFn0(p0 bool, p1 bool, p2 uint8) int {
    p2 = p2
    var v4 int16 = 32767
    for fi5 := 0; fi5 < 8; fi5++ {
        v6 := GenS1{F0: p2, F1: uint16(p2)}
        var v7 uint8 = (p2 << ((18446744073709551615) & 63))
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: (^(p2)), F1: uint16(p2)}
    var v9 int8 = -76
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v8
    _ = v9
    return 127
}
func genFn1(p0 string) bool {
    p0 = "\n\t"
    p0 = ((("hello" + p0) + (p0 + "foo bar")) + "hello")
    var v10 bool = (30 <= 252)
    for fi11 := 0; fi11 < 13; fi11++ {
        p0 = ("\n\t" + ("\n\t" + p0))
        var v12 int = fi11
        var v13 uint32 = 2
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = v10
    return false
}

func main() {
    var v1 int8 = -128
    v2 := make(map[int]string)
    var v3 uint16 = 612
    var v14 uint16 = 345
    var v15 uint8 = 0
    v3 = v3
    if false {
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 uint64 = 2
            if false {
                fmt.Println(v3, v3, v14)
                if ((v1 != (v1 >> ((42) & 63))) || ((^(v17)) == v17)) {
                    v18 := GenS1{F0: (^(v15)), F1: (v3 + (v3 + v3))}
                    _ = v18
                }
                v19 := GenS1{F0: (v15 & v15), F1: 666}
                _ = v19
            }
            if true {
                for fi20 := 0; fi20 < 5; fi20++ {
                    if false {
                        v3 -= (v3 & 821)
                    } else {
                        if (!false) {
                            if ((true || true) && true) {
                                v15 = 127
                            } else {
                                switch ("hello" + "hello") {
                                case "":
                                    v21 := []int{}
                                    _ = v21
                                case "hello":
                                    var v22 int64 = 9223372036854775807
                                    if (0.0 <= float64(1e10)) {
                                        var v23 int64 = (v22 | (^(v22)))
                                        v24 := GenS2{F0: (("\n\t" + ("\n\t" + "\n\t")) + "foo bar"), F1: (-(v1)), F2: 1000}
                                        v25 := v2[(-(fi20))]
                                        v26 := GenS1{F0: v15, F1: v14}
                                        v27 := GenS1{F0: 42, F1: 42}
                                        v3 -= v3
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                    } else {
                                        fmt.Println(fi16, fi20)
                                        switch (fi16 / (((fi20 >> ((2) & 63))) | 1)) {
                                        case 477:
                                            var v28 uint64 = uint64(v15)
                                            _ = v28
                                        case 128:
                                            fmt.Println(v1, v17, fi20)
                                        }
                                        switch fi20 {
                                        case 127:
                                            v29 := GenS2{F0: "", F1: (^(v1)), F2: -612}
                                            _ = v29
                                        case 86:
                                            if (-2.25 <= 1e10) {
                                                var v30 bool = true
                                                v31 := []int{}
                                                fmt.Println(v17, v30)
                                                _ = v30
                                                _ = v31
                                            } else {
                                                var v32 bool = ((-(v1)) >= (v1 & v1))
                                                fmt.Println(v3, v15, v14)
                                                _ = v32
                                            }
                                        case 32768:
                                            v33 := v2[(fi16 | -1)]
                                            _ = v33
                                        }
                                    }
                                    _ = v22
                                default:
                                    switch fi20 {
                                    case 42:
                                        if ((v1 * v1) >= v1) {
                                            var v34 uint16 = (v14 % ((v3) | 1))
                                            _ = v34
                                        }
                                    default:
                                        v3 -= v3
                                    }
                                }
                            }
                            v35 := GenS1{F0: v15, F1: v14}
                            if ((!(((v1 % ((0) | 1)) != (v1 * 100)) || false)) && (!((v1 | v1) != (v1 ^ v1)))) {
                                v36 := []int{fi20, 9223372036854775807, 1, -104}
                                _ = v36
                            }
                            _ = v35
                        }
                    }
                    _ = fi20
                }
                v37 := []int{fi16, (fi16 * genFn0(true, true, (v15 + (v15 * (v15 & v15))))), genFn0(true, false, 1)}
                _ = v37
            }
            _ = fi16
            _ = v17
        }
    }
    for fi38 := 0; fi38 < 10; fi38++ {
        for fi39 := 0; fi39 < 11; fi39++ {
            v1 = v1
            v15 += 42
            _ = fi39
        }
        v40 := GenS2{F0: "\n\t", F1: (v1 % (((v1 ^ 42)) | 1)), F2: int16(v14)}
        _ = fi38
        _ = v40
    }
    var v41 uint64 = 4294967295
    v42 := GenS1{F0: v15, F1: 0}
    v14 = v14
    v43 := GenS1{F0: (v15 & v15), F1: (v3 % ((v14) | 1))}
    v44 := []int{4294967295, 32767, 42}
    var v45 float32 = -2.25
    for fi46 := 0; fi46 < 7; fi46++ {
        v1 += v1
        switch (("a" + "a") + "hello") {
        case "a":
            v41 -= v41
        case "\n\t":
            var v47 int32 = -1
            _ = v47
        }
        if false {
            v48 := GenS2{F0: "a", F1: v1, F2: -32768}
            v49 := GenS1{F0: (v15 >> ((uint(v15)) & 63)), F1: 582}
            v14 += (v3 / ((v14) | 1))
            _ = v48
            _ = v49
        } else {
            v41 = (v41 ^ v41)
            _ = len(v2)
        }
        _ = fi46
    }
    if len(v44) > 0 {
        v44[0] = 864
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
