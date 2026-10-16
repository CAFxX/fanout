package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint; F1 int }
func genFn0(p0 uint16, p1 int16) int32 {
    for fi6 := 0; fi6 < 20; fi6++ {
        p0 -= (p0 / (((p0 % (((^(p0))) | 1))) | 1))
        v7 := GenS1{F0: p1}
        _ = fi6
        _ = v7
    }
    p0 -= p0
    v8 := []int{600, 42}
    _ = p0
    _ = p1
    _ = v8
    return 0
}

func main() {
    var v1 float32 = (0.1 * float32((-2.25 * (float64((float32(3.14159) * 1e10)) - float64(1e10)))))
    v2 := make(map[int]int)
    var v3 bool = false
    var v4 float64 = ((0.1 - ((float64(v1) - 3.14159) * 0.1)) + ((3.14159 + 1.5) - 1.5))
    v5 := []uint32{1, 2147483648, 0}
    switch 424 {
    case 1:
        v9 := GenS1{F0: -32768}
        _ = v9
    case 32767:
        if v3 {
            if (v3 || v3) {
                v10 := GenS1{F0: -32768}
                switch "a" {
                case "hello":
                    var v11 float64 = (v4 + (v4 + (((0.1 - float64((((float32(v4) - float32(((0.0 + 1.5) / (v4 * v4)))) + v1) + v1))) + (float64(v1) + 0.1)) + -2.25)))
                    _ = v11
                case "foo bar":
                    for fi12 := 0; fi12 < 2; fi12++ {
                        var v13 uint = 496
                        v14 := GenS2{F0: (^(fi12)), F1: 42, F2: 128}
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                }
                for fi15 := 0; fi15 < 18; fi15++ {
                    var v16 uint64 = 1
                    var v17 uint = 404
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                _ = v10
            } else {
                fmt.Println(v4, v4)
                var v18 int64 = -9223372036854775808
                _ = v18
            }
        } else {
            var v19 float32 = (v1 * v1)
            var v20 int32 = 0
            _ = v19
            _ = v20
        }
    case -627:
        _ = len(v2)
    default:
        var v21 string = "\n\t"
        _ = v21
    }
    if v3 {
        for fi22 := 0; fi22 < 8; fi22++ {
            v1 -= (1e10 * (float32(float64(float32((v4 + v4)))) * (3.14159 / 0.1)))
            v23 := GenS1{F0: -32768}
            _ = fi22
            _ = v23
        }
        var v24 uint16 = 686
        _ = v24
    }
    var v25 uint = 255
    for fi26 := 0; fi26 < 20; fi26++ {
        switch "foo bar" {
        case "hello":
            var v27 int8 = 57
            _ = v27
        default:
            var v28 int64 = -1
            _ = v28
        }
        v25 = v25
        _ = fi26
    }
    var v29 uint32 = 694
    v25 -= v25
    var v30 int8 = 60
    fmt.Println(v30, v1)
    v31 := GenS1{F0: -32768}
    switch 0 {
    case -905:
        v32 := GenS2{F0: 4294967295, F1: -819, F2: 127}
        _ = v32
    case -2147483648:
        if (("a" + ("\n\t" + "foo bar")) == "") {
            var v33 uint8 = 42
            var v34 uint = v25
            _ = v33
            _ = v34
        } else {
            v3 = ("hello" != "")
        }
    default:
        v35 := GenS2{F0: 9223372036854775807, F1: 42, F2: 2}
        _ = v35
    }
    if v3 {
        v2[137] = 2147483647
    } else {
        for fi36 := 0; fi36 < 12; fi36++ {
            v25 = v25
            _ = len(v5)
            if ((fi36 & fi36) > fi36) {
                var v37 bool = v3
                switch "\n\t" {
                case "":
                    var v38 uint8 = 1
                    _ = v38
                case "a":
                    _ = len(v5)
                case "foo bar":
                    var v39 int8 = -14
                    _ = v39
                default:
                    if (!v3) {
                        if false {
                            for fi40 := 0; fi40 < 17; fi40++ {
                                if (v30 < (v30 % ((127) | 1))) {
                                    if ("hello" == "") {
                                        v41 := GenS3{F0: (v25 | v25), F1: -480}
                                        var v42 float64 = v4
                                        v43 := GenS3{F0: 0, F1: (fi36 << ((v25) & 63))}
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                    }
                                    var v44 uint16 = 56
                                    _ = v44
                                } else {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    v45 := GenS3{F0: v25, F1: fi40}
                                    fmt.Println(v3)
                                    _ = v45
                                }
                                _ = fi40
                            }
                        } else {
                            var v46 int32 = 389
                            v47 := GenS2{F0: (^(fi36)), F1: 42, F2: 255}
                            var v48 float32 = float32(1.5)
                            _ = v46
                            _ = v47
                            _ = v48
                        }
                    }
                }
                _ = v37
            } else {
                switch 32768 {
                case 32767:
                    var v49 bool = (true && (v3 && false))
                    _ = v49
                case 799:
                    _ = len(v5)
                }
                v1 = v1
            }
            _ = fi36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
