package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint32 }
func genFn0() uint32 {
    if ((2147483647 == 694) && false) {
        var v4 int32 = 0
        _ = v4
    } else {
        v5 := []int{-1, 42, 4294967295}
        var v6 int8 = 127
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    var v8 int16 = 32767
    if (18446744073709551615 <= 42) {
        var v9 string = "a"
        var v10 uint = 42
        _ = v10
        _ = v9
    }
    v8 += (^(v8))
    _ = v8
    return 500
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = false
    v3 := make(map[uint32]int)
    var v11 int8 = 1
    v12 := []int{-448, -165}
    if true {
        if false {
            var v13 int16 = 32767
            v14 := GenS1{F0: ("a" + "hello"), F1: 42}
            for fi15 := 0; fi15 < 14; fi15++ {
                var v16 uint = 864
                var v17 int = fi15
                for fi18 := 0; fi18 < 16; fi18++ {
                    var v19 string = ""
                    _ = fi18
                    _ = v19
                }
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = v13
            _ = v14
        } else {
            if false {
                var v20 float32 = 0.0
                var v21 uint32 = 2147483647
                v21 += 42
                _ = v20
                _ = v21
            }
            if len(v12) > 0 {
                v12[0] = 3
            }
        }
        var v22 int32 = 2147483647
        if true {
            if (!(909 >= 42)) {
                fmt.Println(v22, v11, v2)
            } else {
                if (-805 != -279) {
                    v2 = true
                    v23 := GenS2{F0: 127, F1: 0, F2: v22}
                    if (754 <= 2) {
                        for fi24 := 0; fi24 < 14; fi24++ {
                            fmt.Println(v22, v2, v11)
                            v25 := GenS2{F0: 67, F1: 147, F2: (v22 << ((472) & 63))}
                            _ = fi24
                            _ = v25
                        }
                        fmt.Println(v22)
                    } else {
                        v12 = append(v12, 9223372036854775807)
                        var v26 bool = (!(-2.25 <= (0.1 - float64(1.5))))
                        for fi27 := 0; fi27 < 10; fi27++ {
                            v28 := GenS2{F0: 255, F1: 112, F2: 256}
                            var v29 uint = 1
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        _ = v26
                    }
                    _ = v23
                }
                v30 := GenS2{F0: 0, F1: 2, F2: (v22 / (((v22 - v22)) | 1))}
                for fi31 := 0; fi31 < 5; fi31++ {
                    switch "" {
                    case "hello":
                        for fi32 := 0; fi32 < 7; fi32++ {
                            v33 := GenS3{F0: 1, F1: 841, F2: 65535}
                            var v34 uint16 = 985
                            v22 = (^(v22))
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                    case "a":
                        if (!false) {
                            var v35 float32 = -2.25
                            v22 -= v22
                            var v36 uint64 = 42
                            _ = v35
                            _ = v36
                        }
                    case "foo bar":
                        v11 += 42
                    default:
                        v11 = 127
                    }
                    _ = fi31
                }
                _ = v30
            }
        }
        _ = v22
    }
    var v37 float32 = -2.25
    v38 := GenS3{F0: 128, F1: 654, F2: 910}
    var v39 int = 0
    var v40 bool = (!(0 >= 1))
    for fi41 := 0; fi41 < 3; fi41++ {
        switch (fi41 * (fi41 << ((0) & 63))) {
        case 32767:
            _ = len(v1)
        case 32768:
            if v2 {
                var v42 uint16 = 903
                _ = v42
            } else {
                if (!((!false) && true)) {
                    var v43 uint16 = 42
                    var v44 uint8 = 128
                    var v45 uint16 = (v43 & v43)
                    _ = v43
                    _ = v44
                    _ = v45
                }
            }
        case 2147483647:
            v40 = (0 >= 707)
        default:
            v39 = v39
        }
        fmt.Println(fi41, v2, v40)
        _ = fi41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v40
}
