package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 uint16; F2 uint64 }
func genFn0() bool {
    v6 := []int{}
    switch "" {
    case "hello":
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 uint8 = 42
            if (1.5 > 0.1) {
                for fi9 := 0; fi9 < 11; fi9++ {
                    v8 = (^(v8))
                    _ = fi9
                }
                for fi10 := 0; fi10 < 3; fi10++ {
                    fmt.Println(v8, v8, v8)
                    _ = fi10
                }
            }
            _ = fi7
            _ = v8
        }
    default:
        var v11 uint16 = 465
        _ = v11
    }
    var v12 int = 428
    _ = v12
    _ = v6
    return true
}
func genFn1() bool {
    v13 := GenS1{F0: 18446744073709551615, F1: 88}
    var v14 int16 = -1
    var v15 bool = (!false)
    switch 9223372036854775807 {
    case 2147483647:
        var v16 int64 = 0
        _ = v16
    case -241:
        var v17 uint8 = 86
        _ = v17
    }
    v18 := []int{128, -1, 9223372036854775807}
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    return (("a" + "hello") >= "foo bar")
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 647
    var v3 int = -1
    var v4 int16 = 1
    v5 := []string{}
    if len(v5) > 0 {
        v5[0] = ""
    }
    var v19 uint64 = v2
    var v20 uint64 = 0
    var v21 float64 = 3.14159
    v22 := GenS1{F0: v20, F1: 42}
    var v23 string = "a"
    v23 = v23
    var v24 int32 = 1
    var v25 int8 = 0
    fmt.Println(v19, v21, v4)
    for fi26 := 0; fi26 < 19; fi26++ {
        v20 -= (v2 | (v2 * 344))
        switch fi26 {
        case 9223372036854775807:
            var v27 int = v3
            _ = v27
        case -9223372036854775808:
            var v28 uint = 65535
            _ = v28
        }
        if false {
            var v29 string = ("" + v23)
            if false {
                for fi30 := 0; fi30 < 18; fi30++ {
                    if ((false && (true || (!(false && true)))) && (v23 == "foo bar")) {
                        var v31 bool = (41 == 42)
                        if (!v31) {
                            if (v31 && false) {
                                var v32 float64 = ((v21 - v21) / v21)
                                if (true || v31) {
                                    v33 := GenS1{F0: (v2 + (v2 ^ (^(v19)))), F1: 2}
                                    _ = v33
                                }
                                if v31 {
                                    if (!(v29 <= ("\n\t" + v29))) {
                                        fmt.Println(v21, v19)
                                        v23 = v23
                                    }
                                    v25 += -101
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                                _ = v32
                            } else {
                                if len(v5) > 0 {
                                    v5[0] = ("a" + v29)
                                }
                                if v31 {
                                    v34 := GenS1{F0: v20, F1: 191}
                                    v29 = "foo bar"
                                    var v35 uint = 2
                                    _ = v34
                                    _ = v35
                                } else {
                                    switch "" {
                                    case "":
                                        fmt.Println(v24, v24, v21)
                                    }
                                    v3 += fi30
                                }
                                v36 := GenS2{F0: 4, F1: v19, F2: -748}
                                _ = v36
                            }
                        }
                        _ = v31
                    } else {
                        switch (v23 + "") {
                        case "a":
                            var v37 uint = 42
                            _ = v37
                        case "\n\t":
                            v20 -= v20
                        }
                        if (v20 >= 4294967295) {
                            fmt.Println(v23, v21, v4)
                            _ = len(v1)
                            delete(v1, -792)
                            var v38 uint8 = 2
                            _ = v38
                        } else {
                            fmt.Println(v19)
                            switch int(v20) {
                            case 9223372036854775807:
                                var v39 uint32 = 2
                                _ = v39
                            case -477:
                                v40 := GenS3{F0: (v25 << ((1) & 63))}
                                _ = v40
                            case -9223372036854775808:
                                v1[(-(fi26))] = true
                            default:
                                v3 = (^(v3))
                            }
                        }
                        var v41 int64 = 49
                        _ = v41
                    }
                    v20 -= v2
                    var v42 int = (-(v3))
                    _ = fi30
                    _ = v42
                }
                if (false && (-9223372036854775808 >= int64(v20))) {
                    v43 := GenS4{F0: 0, F1: 65535, F2: v20}
                    var v44 uint64 = 996
                    _ = v43
                    _ = v44
                }
            }
            v23 = ("\n\t" + "a")
            _ = v29
        }
        _ = fi26
    }
    switch v3 {
    case -347:
        for fi45 := 0; fi45 < 12; fi45++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi45
        }
    case 2147483647:
        fmt.Println(v2, v2, v2)
    }
    for fi46 := 0; fi46 < 3; fi46++ {
        v4 -= v4
        _ = len(v5)
        _ = fi46
    }
    var v47 float32 = ((float32(v21) + -2.25) + 0.1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
