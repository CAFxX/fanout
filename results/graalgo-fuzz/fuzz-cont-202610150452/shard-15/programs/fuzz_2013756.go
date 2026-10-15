package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint16; F2 int16 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint16) int32 {
    for fi6 := 0; fi6 < 2; fi6++ {
        switch fi6 {
        case 127:
            for fi7 := 0; fi7 < 18; fi7++ {
                var v8 string = (("" + ("" + (("foo bar" + "foo bar") + ((("hello" + "foo bar") + "") + "hello")))) + "")
                v8 = ("hello" + "a")
                v8 = (v8 + (v8 + v8))
                var v9 bool = (-9223372036854775808 > 0)
                if (!(700 == 42)) {
                    switch "" {
                    case "hello":
                        v10 := []int{(fi7 % ((fi7) | 1)), (fi7 & (fi7 % ((fi7) | 1))), fi6}
                        _ = v10
                    case "":
                        var v11 float32 = (0.1 / ((float32(3.14159) / 1e10) / (3.14159 / 1e10)))
                        _ = v11
                    case "\n\t":
                        fmt.Println(v8)
                    default:
                        v9 = (((p0 - 2) >= 1) || v9)
                    }
                    if true {
                        v9 = (false && (-2147483648 < fi7))
                    } else {
                        var v12 uint = 42
                        var v13 uint16 = uint16(v12)
                        var v14 int = fi7
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    var v15 uint64 = 604
                    _ = v15
                }
                _ = fi7
                _ = v8
                _ = v9
            }
        }
        _ = fi6
    }
    var v16 int = -1
    var v17 int32 = 0
    fmt.Println(v16)
    v18 := GenS1{F0: 0}
    _ = p0
    _ = v16
    _ = v17
    _ = v18
    return v17
}
func genFn1(p0 uint32) uint64 {
    v19 := GenS2{F0: 2, F1: 46, F2: 42}
    var v20 uint8 = 0
    for fi21 := 0; fi21 < 5; fi21++ {
        v22 := GenS3{F0: 18, F1: 842}
        _ = fi21
        _ = v22
    }
    _ = p0
    _ = v19
    _ = v20
    return 461
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    v3 := []string{"hello", "\n\t", ("a" + "a"), (("\n\t" + "a") + ("a" + "\n\t"))}
    v4 := []uint32{4294967295, 220}
    v5 := make(map[string]int)
    if ((2 < 1) || ((!true) && (65535 >= 897))) {
        var v23 int = 4294967295
        var v24 uint64 = genFn1(65535)
        if (2 > 2) {
            var v25 bool = (((v23 * 798) != (^(v23))) || true)
            _ = v25
        }
        _ = v23
        _ = v24
    } else {
        switch -1 {
        case 262:
            var v26 uint8 = 254
            _ = v26
        case 127:
            var v27 uint64 = 18446744073709551615
            _ = v27
        }
        var v28 int16 = 1000
        if true {
            if true {
                if false {
                    v29 := v2["foo bar"]
                    fmt.Println(v29, v28)
                    _ = v29
                } else {
                    for fi30 := 0; fi30 < 1; fi30++ {
                        if false {
                            for fi31 := 0; fi31 < 16; fi31++ {
                                if (true && false) {
                                    fmt.Println(fi30)
                                    var v32 int64 = 42
                                    _ = v32
                                }
                                _ = len(v3)
                                v28 += v28
                                _ = fi31
                            }
                        } else {
                            fmt.Println(fi30, fi30)
                            fmt.Println(v28)
                            var v33 int8 = 100
                            _ = v33
                        }
                        _ = fi30
                    }
                }
                var v34 uint8 = 29
                var v35 float64 = ((-2.25 + 1.5) / float64((float32(0.0) - 1.5)))
                _ = v34
                _ = v35
            } else {
                if (-2.25 == 0.1) {
                    if false {
                        for fi36 := 0; fi36 < 13; fi36++ {
                            v28 += v28
                            var v37 float32 = (1e10 - ((0.0 + 1.5) * 3.14159))
                            if true {
                                var v38 int16 = 0
                                _ = v38
                            } else {
                                switch (fi36 & (fi36 + 2147483647)) {
                                case 2147483647:
                                    for fi39 := 0; fi39 < 10; fi39++ {
                                        v4 = append(v4, 65535)
                                        _ = fi39
                                    }
                                case 32767:
                                    if (!(72 == fi36)) {
                                        fmt.Println(v37)
                                    }
                                default:
                                    v2[""] = true
                                }
                                var v40 uint64 = 0
                                if true {
                                    _ = len(v4)
                                    v41 := GenS4{F0: 42, F1: 42}
                                    _ = v41
                                }
                                _ = v40
                            }
                            _ = fi36
                            _ = v37
                        }
                    } else {
                        v4 = append(v4, 4294967295)
                        var v42 int8 = -128
                        if true {
                            if false {
                                v28 = v28
                            }
                            var v43 float32 = -2.25
                            _ = v43
                        }
                        _ = v42
                    }
                    v28 = v28
                    v28 = v28
                } else {
                    var v44 string = ("a" + "")
                    v44 = v44
                    fmt.Println(v44)
                    _ = v44
                }
                v28 += v28
                if (-424 < 176) {
                    v45 := GenS3{F0: 6, F1: 1}
                    for fi46 := 0; fi46 < 13; fi46++ {
                        var v47 int32 = -1
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v48 int8 = -128
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                    for fi49 := 0; fi49 < 19; fi49++ {
                        for fi50 := 0; fi50 < 9; fi50++ {
                            var v51 string = ("\n\t" + "a")
                            v52 := GenS4{F0: 2, F1: 642}
                            _ = fi50
                            _ = v51
                            _ = v52
                        }
                        v28 = v28
                        v53 := GenS4{F0: 65535, F1: 255}
                        _ = fi49
                        _ = v53
                    }
                    _ = v45
                }
            }
            v3 = append(v3, "")
        }
        _ = v28
    }
    var v54 uint32 = 2
    var v55 uint32 = 65535
    for fi56 := 0; fi56 < 15; fi56++ {
        v55 -= v55
        _ = fi56
    }
    v5["\n\t"] = 9223372036854775807
    v55 = v55
    v54 = 207
    v57 := GenS5{F0: v54}
    fmt.Println(v54)
    if false {
        switch 128 {
        case -722:
            for fi58 := 0; fi58 < 16; fi58++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v4 = append(v4, v54)
                _ = fi58
            }
        }
        for fi59 := 0; fi59 < 10; fi59++ {
            v54 = v54
            fmt.Println(v55, fi59, v54)
            _ = fi59
        }
    }
    delete(v1, v54)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v57
}
