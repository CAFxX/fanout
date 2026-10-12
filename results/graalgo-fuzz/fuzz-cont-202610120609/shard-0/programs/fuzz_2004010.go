package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string }
func genFn0(p0 int16, p1 int32, p2 int32) uint64 {
    p1 = p2
    fmt.Println(p0, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    return 170
}
func genFn1() int8 {
    v7 := GenS1{F0: "hello", F1: 127, F2: 0}
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 string = (("\n\t" + ("hello" + "foo bar")) + "hello")
        v10 := GenS1{F0: (v9 + v9), F1: -115, F2: -9223372036854775808}
        var v11 int = fi8
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    v12 := []int{9223372036854775807}
    _ = v12
    _ = v7
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 2
    v3 := make(map[uint32]int)
    var v4 uint = 255
    v5 := make(map[int]uint32)
    v6 := []string{"foo bar", "\n\t"}
    var v13 uint64 = 18446744073709551615
    for fi14 := 0; fi14 < 17; fi14++ {
        v4 = 1
        _ = fi14
    }
    v4 = 42
    fmt.Println(v2, v2)
    for fi15 := 0; fi15 < 13; fi15++ {
        for fi16 := 0; fi16 < 19; fi16++ {
            var v17 int16 = -276
            _ = fi16
            _ = v17
        }
        var v18 uint16 = (v2 >> ((255) & 63))
        _ = fi15
        _ = v18
    }
    v19 := GenS1{F0: "a", F1: 102, F2: 1}
    v13 = genFn0(271, 1, -2147483648)
    delete(v3, 2147483648)
    v13 -= 0
    v20 := GenS1{F0: ("a" + "\n\t"), F1: int8(v2), F2: 42}
    v4 += v4
    v2 += v2
    fmt.Println(v4, v2)
    switch "hello" {
    case "\n\t":
        if (0 > 888) {
            var v21 bool = ((false || (1000 < -138)) || (v4 == (v4 << (((v4 ^ (v4 / ((v4) | 1)))) & 63))))
            v13 += 1
            _ = v21
        } else {
            v22 := GenS2{F0: int32(v4)}
            if ("" != (((("\n\t" + "foo bar") + "") + "foo bar") + "a")) {
                var v23 bool = (false || false)
                if (v23 || v23) {
                    v4 -= v4
                }
                _ = v23
            }
            _ = v22
        }
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        for fi25 := 0; fi25 < 10; fi25++ {
            var v26 float32 = 0.1
            if true {
                var v27 uint32 = 2147483648
                var v28 uint16 = 213
                v29 := GenS2{F0: -1}
                _ = v27
                _ = v28
                _ = v29
            }
            var v30 uint = v4
            v13 -= (^(v13))
            var v31 uint32 = 227
            var v32 int16 = 32767
            _ = fi25
            _ = v26
            _ = v30
            _ = v31
            _ = v32
        }
        for fi33 := 0; fi33 < 6; fi33++ {
            if false {
                switch "foo bar" {
                case "":
                    var v34 int16 = 0
                    _ = v34
                case "foo bar":
                    switch fi33 {
                    case 42:
                        v2 += v2
                    case 4294967295:
                        v13 = (^(v13))
                        if (0 > 1) {
                            fmt.Println(v13, v4, fi33)
                            v35 := GenS1{F0: ("" + "foo bar"), F1: 21, F2: 42}
                            _ = v35
                        } else {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            if true {
                                var v36 bool = (!true)
                                if (false && (("hello" + "a") >= ("foo bar" + "foo bar"))) {
                                    var v37 int32 = 42
                                    fmt.Println(v13, fi33)
                                    _ = v37
                                } else {
                                    delete(v3, 599)
                                }
                                v36 = (!true)
                                _ = v36
                            } else {
                                var v38 float32 = 0.1
                                _ = v38
                            }
                            fmt.Println(fi24)
                        }
                        if (-106 < 1) {
                            fmt.Println(fi33, v2, v13)
                        } else {
                            var v39 float32 = 1.5
                            v2 += (v2 * v2)
                            _ = v39
                        }
                        if (((float64((1e10 - -2.25)) - (0.1 - 0.0)) - 1e10) < 0.1) {
                            var v40 uint64 = (v13 / ((uint64(v2)) | 1))
                            _ = v40
                        } else {
                            var v41 uint16 = (v2 | v2)
                            _ = v41
                        }
                    default:
                        var v42 bool = false
                        _ = v42
                    }
                default:
                    v13 -= (v13 + 0)
                }
                var v43 uint = v4
                v44 := GenS3{F0: "a"}
                switch (fi24 >> ((v4) & 63)) {
                case 2147483647:
                    var v45 uint8 = uint8(v43)
                    _ = v45
                case 32767:
                    fmt.Println(v4, v4, v43)
                case 9223372036854775807:
                    var v46 int16 = 265
                    _ = v46
                default:
                    fmt.Println(v4, v13, v13)
                }
                v13 = uint64(v2)
                _ = v43
                _ = v44
            }
            _ = fi33
        }
        v13 -= (v13 & v13)
        _ = fi24
    }
    delete(v3, 257)
    fmt.Println(v13, v2, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
