package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 int; F2 uint64 }
type GenS4 struct { F0 bool; F1 int8; F2 int }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 int8 }
func genFn0(p0 bool) string {
    var v7 uint = 782
    if true {
        v8 := []int{}
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 int8 = 100
            for fi11 := 0; fi11 < 9; fi11++ {
                v12 := GenS1{F0: 0}
                _ = fi11
                _ = v12
            }
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 int8 = 127
        v15 := []int{-138}
        fmt.Println(p0)
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(v7)
    v7 += v7
    _ = p0
    _ = v7
    return (("\n\t" + ("\n\t" + (("" + "a") + "\n\t"))) + "hello")
}
func genFn1(p0 int32, p1 int32, p2 uint32) bool {
    p2 += (p2 | p2)
    v16 := GenS1{F0: 594}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return true
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[string]uint32)
    var v3 uint64 = 1
    v4 := []string{("\n\t" + "hello"), "\n\t", "foo bar"}
    v5 := make(map[string]string)
    v6 := []int{404, 42}
    v5[("" + "foo bar")] = "hello"
    var v17 uint32 = 4294967295
    for fi18 := 0; fi18 < 19; fi18++ {
        if false {
            v17 -= (v17 / (((v17 / (((v17 ^ v17)) | 1))) | 1))
            for fi19 := 0; fi19 < 17; fi19++ {
                var v20 uint = 65535
                var v21 int16 = 772
                _ = fi19
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 7; fi22++ {
                _ = len(v2)
                var v23 int16 = 42
                _ = fi22
                _ = v23
            }
        } else {
            v3 = 307
        }
        _ = fi18
    }
    v24 := GenS2{F0: 910}
    fmt.Println(v3, v17, v17)
    if false {
        _ = len(v6)
        if (!true) {
            v3 -= uint64(v17)
            fmt.Println(v3)
            v25 := GenS3{F0: true, F1: -327, F2: (^(v3))}
            _ = v25
        } else {
            v3 = (v3 ^ v1)
            for fi26 := 0; fi26 < 17; fi26++ {
                var v27 float64 = 1e10
                _ = fi26
                _ = v27
            }
        }
    } else {
        if (float64((0.1 / 1.0)) < 1e10) {
            delete(v5, "hello")
        } else {
            if (-1 > 1000) {
                v28 := GenS4{F0: false, F1: -128, F2: 32767}
                v29 := GenS5{F0: (1 == (v1 >> ((0) & 63))), F1: v1}
                for fi30 := 0; fi30 < 16; fi30++ {
                    v31 := GenS4{F0: false, F1: -128, F2: fi30}
                    v32 := v2["a"]
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                _ = v28
                _ = v29
            }
            v33 := GenS6{F0: -97}
            for fi34 := 0; fi34 < 8; fi34++ {
                for fi35 := 0; fi35 < 8; fi35++ {
                    switch (("hello" + "a") + "\n\t") {
                    case "a":
                        v36 := GenS6{F0: -128}
                        _ = v36
                    case "hello":
                        v37 := GenS4{F0: true, F1: 1, F2: (fi35 * (fi35 - fi35))}
                        _ = v37
                    case "":
                        v17 = (v17 | v17)
                    default:
                        if (!true) {
                            switch fi34 {
                            case -723:
                                var v38 uint16 = 255
                                _ = v38
                            case 2147483647:
                                v3 += v3
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v39 := v2["hello"]
                                v17 += (v17 % (((v17 / (((v17 - 65535)) | 1))) | 1))
                                if (1 >= 42) {
                                    var v40 int8 = 100
                                    _ = v40
                                } else {
                                    switch "hello" {
                                    case "hello":
                                        var v41 int64 = -172
                                        _ = v41
                                    default:
                                        var v42 uint = 522
                                        _ = v42
                                    }
                                    if false {
                                        v6 = append(v6, fi34)
                                        fmt.Println(v17)
                                        if ((!(float32((1.5 + (float64(0.0) / float64(1.5)))) > 0.1)) && ((!(false && true)) || true)) {
                                            if (((fi34 * fi34) > (fi34 << ((uint(v3)) & 63))) && false) {
                                                v43 := v5[(("hello" + "hello") + "foo bar")]
                                                v44 := GenS6{F0: 1}
                                                _ = v43
                                                _ = v44
                                            }
                                        } else {
                                            v6 = append(v6, 307)
                                            v17 = v39
                                            var v45 uint = 0
                                            _ = v45
                                        }
                                    }
                                }
                                _ = v39
                            case 0:
                                v3 = (v3 << ((0) & 63))
                            }
                            switch fi34 {
                            case -2147483648:
                                var v46 int8 = -128
                                _ = v46
                            case 1:
                                if len(v4) > 0 {
                                    v4[0] = "a"
                                }
                            case 582:
                                var v47 int = -9223372036854775808
                                _ = v47
                            default:
                                v3 += (v3 & v3)
                            }
                        }
                    }
                    v1 = v3
                    _ = fi35
                }
                _ = fi34
            }
            _ = v33
        }
        switch "" {
        case "hello":
            var v48 float64 = -2.25
            _ = v48
        case "a":
            v17 += (v17 / ((v17) | 1))
        case "foo bar":
            v17 = v17
        default:
            _ = len(v5)
        }
    }
    if (4294967295 > 4294967295) {
        var v49 bool = ((-2.25 + 1e10) < (1e10 * float64(0.0)))
        v3 -= v1
        v49 = v49
        _ = v49
    } else {
        if ((255 != 255) && (32767 <= -32768)) {
            v1 -= v3
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v2)
    switch "a" {
    case "\n\t":
        var v50 uint64 = v3
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
