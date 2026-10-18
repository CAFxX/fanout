package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 bool }
func genFn0(p0 uint16, p1 uint32, p2 int) bool {
    fmt.Println(p1, p2, p1)
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return false
}
func genFn1(p0 uint64) uint16 {
    v5 := GenS1{F0: 236}
    v6 := []int{321, 42}
    switch -9223372036854775808 {
    case 4294967295:
        p0 = (p0 | 4294967295)
    case -693:
        fmt.Println(p0)
    default:
        if false {
            for fi7 := 0; fi7 < 14; fi7++ {
                fmt.Println(fi7, p0)
                _ = fi7
            }
            p0 += (p0 % ((p0) | 1))
            for fi8 := 0; fi8 < 5; fi8++ {
                p0 = (p0 >> ((4294967295) & 63))
                v9 := GenS2{F0: "a"}
                _ = fi8
                _ = v9
            }
        } else {
            var v10 uint16 = 255
            _ = v10
        }
    }
    _ = p0
    _ = v5
    _ = v6
    return 588
}

func main() {
    var v1 bool = false
    v2 := []int{42, -1}
    var v3 int64 = -9223372036854775808
    v4 := make(map[uint32]bool)
    var v11 bool = v1
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 uint = 42
        v11 = (0 <= 1)
        _ = fi12
        _ = v13
    }
    v14 := GenS3{F0: true}
    if v11 {
        for fi15 := 0; fi15 < 12; fi15++ {
            if v11 {
                var v16 int = fi15
                v1 = v11
                for fi17 := 0; fi17 < 1; fi17++ {
                    if (v1 && v11) {
                        v18 := GenS1{F0: 898}
                        if len(v2) > 0 {
                            v2[0] = v16
                        }
                        switch (fi17 ^ v16) {
                        case 32767:
                            _ = len(v2)
                        case 0:
                            var v19 float32 = (float32(-2.25) / 1.0)
                            _ = v19
                        }
                        _ = v18
                    } else {
                        var v20 uint32 = 2147483647
                        _ = v20
                    }
                    var v21 int8 = 100
                    if (346 == -9223372036854775808) {
                        var v22 string = "foo bar"
                        switch "foo bar" {
                        case "":
                            var v23 uint16 = genFn1(0)
                            _ = v23
                        case "foo bar":
                            var v24 uint16 = 32767
                            var v25 uint64 = 18446744073709551615
                            if v11 {
                                if (0.0 == (1.5 - -2.25)) {
                                    switch 127 {
                                    case 233:
                                        switch v22 {
                                        case "foo bar":
                                            if (v1 || v11) {
                                                var v26 int8 = v21
                                                fmt.Println(v22)
                                                _ = v26
                                            }
                                        case "a":
                                            v16 = v16
                                        }
                                    }
                                    if (((v25 | (v25 << ((127) & 63))) < (v25 ^ (^(v25)))) || (!((!v1) || true))) {
                                        v2 = append(v2, int(v25))
                                        if (4294967295 < 65535) {
                                            v27 := GenS4{F0: -942, F1: (v11 || v11)}
                                            v3 -= -1
                                            _ = v27
                                        } else {
                                            v24 = v24
                                        }
                                    } else {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v28 int32 = 0
                                        v1 = (!false)
                                        _ = v28
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = (v16 * (fi15 | (v16 ^ fi15)))
                                    }
                                }
                            } else {
                                if false {
                                    fmt.Println(v16)
                                }
                                var v29 int64 = v3
                                v30 := GenS4{F0: -733, F1: v11}
                                _ = v29
                                _ = v30
                            }
                            _ = v24
                            _ = v25
                        case "\n\t":
                            var v31 string = ((v22 + v22) + (("hello" + v22) + "\n\t"))
                            _ = v31
                        }
                        v1 = false
                        _ = v22
                    }
                    fmt.Println(v21)
                    if (!(-2.25 == 3.14159)) {
                        _ = len(v2)
                    }
                    if (-2147483648 == -1) {
                        if v11 {
                            var v32 bool = false
                            v21 += (^(v21))
                            _ = v32
                        } else {
                            v33 := v4[42]
                            v33 = v33
                            var v34 string = ("hello" + (("" + "hello") + "\n\t"))
                            v16 += v16
                            v35 := GenS2{F0: v34}
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                        fmt.Println(v16, v21)
                        if (v11 || ((-32768 <= 0) && (!v11))) {
                            var v36 string = ""
                            _ = v36
                        } else {
                            fmt.Println(fi17, fi17)
                            v11 = true
                        }
                    }
                    var v37 float32 = ((0.0 - 3.14159) + -2.25)
                    fmt.Println(fi17, v21)
                    _ = fi17
                    _ = v21
                    _ = v37
                }
                _ = v16
            } else {
                v3 = v3
                if v11 {
                    var v38 float32 = (-2.25 / 0.1)
                    _ = v38
                } else {
                    delete(v4, 4294967295)
                }
                _ = len(v2)
            }
            var v39 float64 = (1e10 / 1.5)
            v11 = (-1 != v3)
            _ = fi15
            _ = v39
        }
    }
    for fi40 := 0; fi40 < 4; fi40++ {
        var v41 bool = true
        _ = fi40
        _ = v41
    }
    var v42 int = 128
    v3 = -578
    _ = len(v4)
    var v43 float64 = 1e10
    var v44 int = 197
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
}
