package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8; F1 uint; F2 int }
type GenS6 struct { F0 uint64; F1 uint64; F2 bool }
type GenS7 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 uint = 872
    var v2 int8 = 100
    var v3 bool = (1 == 2147483647)
    v4 := make(map[int]bool)
    var v5 uint32 = 412
    v6 := []int{354, -607, 127}
    v7 := GenS1{F0: 4294967295, F1: 834}
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 6; fi9++ {
            if true {
                v10 := GenS1{F0: 0, F1: 362}
                v2 -= v2
                _ = v10
            }
            v4[int(v1)] = false
            _ = fi9
        }
        if true {
            v6 = append(v6, (fi8 & 1))
            if (-32768 < 0) {
                var v11 uint16 = 32767
                v12 := GenS2{F0: uint64(v1)}
                v2 = int8(fi8)
                _ = v11
                _ = v12
            }
            var v13 bool = v3
            _ = v13
        }
        _ = fi8
    }
    v14 := GenS1{F0: 42, F1: 1}
    switch "\n\t" {
    case "foo bar":
        v15 := v4[232]
        _ = v15
    case "":
        v1 = (v1 >> ((v1) & 63))
    case "\n\t":
        if v3 {
            v3 = ((v5 + v5) <= (v5 / (((v5 << ((v1) & 63))) | 1)))
            for fi16 := 0; fi16 < 8; fi16++ {
                if (!(v3 || (!(!false)))) {
                    if false {
                        var v17 float32 = -2.25
                        for fi18 := 0; fi18 < 17; fi18++ {
                            var v19 int8 = v2
                            v20 := GenS3{F0: uint8(v1)}
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        _ = v17
                    } else {
                        if false {
                            if (!(v3 && true)) {
                                var v21 int8 = (^(v2))
                                v4[(fi16 >> ((18446744073709551615) & 63))] = v3
                                v1 = (^(v1))
                                _ = v21
                            } else {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                for fi22 := 0; fi22 < 8; fi22++ {
                                    if v3 {
                                        var v23 int64 = 42
                                        _ = v23
                                    }
                                    var v24 int32 = 0
                                    switch ((((("a" + "hello") + (("foo bar" + "a") + ("hello" + "a"))) + "") + "\n\t") + "foo bar") {
                                    case "":
                                        v24 = v24
                                    default:
                                        var v25 uint = v1
                                        _ = v25
                                    }
                                    if ((v2 - v2) < 1) {
                                        v1 = v1
                                        v26 := GenS4{F0: 2}
                                        _ = v26
                                    } else {
                                        var v27 int8 = 111
                                        if (0 < 2) {
                                            var v28 int = fi22
                                            _ = v28
                                        }
                                        _ = v27
                                    }
                                    fmt.Println(v1, fi16)
                                    _ = fi22
                                    _ = v24
                                }
                                var v29 int32 = -2147483648
                                _ = v29
                            }
                        } else {
                            v5 = (v5 * (v5 >> (((^(v1))) & 63)))
                            v30 := GenS3{F0: 255}
                            var v31 bool = false
                            _ = v30
                            _ = v31
                        }
                        for fi32 := 0; fi32 < 2; fi32++ {
                            _ = len(v6)
                            var v33 uint8 = 2
                            switch fi16 {
                            case 128:
                                fmt.Println(fi16, v5, fi32)
                            default:
                                v1 -= v1
                            }
                            _ = fi32
                            _ = v33
                        }
                    }
                } else {
                    if v3 {
                        v34 := GenS4{F0: 165}
                        switch (fi16 % (((fi16 << ((18446744073709551615) & 63))) | 1)) {
                        case 42:
                            v2 = v2
                        case 966:
                            switch ("" + ("" + "hello")) {
                            case "hello":
                                v35 := GenS5{F0: 42, F1: 941, F2: (fi16 * (fi16 + fi16))}
                                _ = v35
                            }
                        case 467:
                            var v36 float32 = 3.14159
                            _ = v36
                        }
                        if v3 {
                            v37 := GenS6{F0: 2, F1: uint64(v5), F2: false}
                            delete(v4, fi16)
                            if len(v6) > 0 {
                                v6[0] = (fi16 & (fi16 - fi16))
                            }
                            _ = v37
                        }
                        _ = v34
                    } else {
                        v38 := GenS7{F0: 1, F1: 229}
                        v4[4294967295] = (1 != int64(v1))
                        _ = v38
                    }
                    v2 += (v2 % ((v2) | 1))
                    v6 = append(v6, 2147483647)
                }
                fmt.Println(v3, v3)
                _ = fi16
            }
        }
    default:
        if len(v6) > 0 {
            v6[0] = 1
        }
    }
    var v39 float64 = 0.0
    v5 -= (^(v5))
    v3 = (-9223372036854775808 > 1)
    var v40 uint64 = uint64(v1)
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
}
