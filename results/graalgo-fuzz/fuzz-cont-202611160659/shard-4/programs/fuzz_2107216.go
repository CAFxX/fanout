package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 int32 }

func main() {
    v1 := []int{32768}
    var v2 uint64 = 212
    var v3 int64 = -1
    fmt.Println(v2, v3)
    v3 -= (v3 - v3)
    v3 = -9223372036854775808
    if true {
        v3 += v3
    }
    v4 := GenS1{F0: 1, F1: ("a" + ""), F2: (59 < 42)}
    for fi5 := 0; fi5 < 1; fi5++ {
        v3 = (v3 * v3)
        var v6 uint = 1
        v7 := GenS1{F0: 9223372036854775807, F1: ((("hello" + "foo bar") + "hello") + ""), F2: (!(v3 != v3))}
        _ = fi5
        _ = v6
        _ = v7
    }
    v1 = append(v1, int(v3))
    var v8 int32 = 42
    if (2 < 179) {
        _ = len(v1)
        if true {
            fmt.Println(v8)
        }
    } else {
        fmt.Println(v8)
    }
    switch "" {
    case "":
        v1 = append(v1, 32767)
    }
    _ = len(v1)
    v9 := GenS2{F0: 139}
    if len(v1) > 0 {
        v1[0] = 42
    }
    fmt.Println(v2, v8)
    if false {
        if true {
            if (!(false || true)) {
                for fi10 := 0; fi10 < 6; fi10++ {
                    v11 := GenS2{F0: 128}
                    fmt.Println(v2, v8)
                    _ = fi10
                    _ = v11
                }
            } else {
                v2 = v2
            }
            if true {
                var v12 string = "foo bar"
                v12 = ""
                var v13 string = (((v12 + ("hello" + (("\n\t" + v12) + (v12 + "hello")))) + v12) + v12)
                _ = v12
                _ = v13
            } else {
                fmt.Println(v3, v8, v2)
            }
            for fi14 := 0; fi14 < 4; fi14++ {
                if (2147483647 >= 0) {
                    v3 = (v3 | (v3 | v3))
                    if false {
                        fmt.Println(v3)
                    }
                    v2 = (v2 * 963)
                } else {
                    if false {
                        fmt.Println(v8)
                        var v15 uint32 = 474
                        _ = v15
                    } else {
                        v2 = v2
                        var v16 int64 = -689
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v16
                    }
                    if ((-(v3)) > -111) {
                        var v17 uint = 4294967295
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 17; fi18++ {
                        var v19 int16 = 1
                        v2 += v2
                        _ = fi18
                        _ = v19
                    }
                }
                if false {
                    if false {
                        v20 := GenS3{F0: 2147483647}
                        var v21 uint64 = (v2 & 134)
                        _ = len(v1)
                        _ = v20
                        _ = v21
                    }
                    if ((("a" + ("foo bar" + "foo bar")) == "foo bar") || (2 >= 324)) {
                        if ((v3 * int64(fi14)) != int64(v8)) {
                            var v22 string = "\n\t"
                            v23 := GenS2{F0: 128}
                            var v24 uint32 = 42
                            _ = v22
                            _ = v23
                            _ = v24
                        } else {
                            var v25 uint = 42
                            v26 := GenS2{F0: 251}
                            _ = v25
                            _ = v26
                        }
                        fmt.Println(fi14, v3, v3)
                        switch "\n\t" {
                        case "foo bar":
                            var v27 int64 = v3
                            _ = v27
                        default:
                            var v28 float64 = 1e10
                            _ = v28
                        }
                    }
                    _ = len(v1)
                } else {
                    if true {
                        switch (fi14 / ((fi14) | 1)) {
                        case 4294967295:
                            v2 = (v2 + v2)
                        case 0:
                            v1 = append(v1, fi14)
                        }
                    }
                    for fi29 := 0; fi29 < 13; fi29++ {
                        var v30 int16 = 32767
                        v3 -= v3
                        v2 += v2
                        v3 = -9223372036854775808
                        if (-2.25 != 0.0) {
                            if true {
                                var v31 uint16 = 32767
                                switch -965 {
                                case 419:
                                    v32 := GenS1{F0: (fi14 + -660), F1: "", F2: false}
                                    _ = v32
                                }
                                _ = v31
                            } else {
                                var v33 float32 = 0.1
                                var v34 uint16 = 0
                                v35 := GenS2{F0: 128}
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                            v36 := GenS4{F0: (v8 - v8), F1: v8}
                            var v37 int16 = (v30 + 549)
                            _ = v36
                            _ = v37
                        } else {
                            fmt.Println(v2, v3, fi14)
                            var v38 float64 = 1.5
                            v2 -= v2
                            _ = v38
                        }
                        _ = fi29
                        _ = v30
                    }
                    fmt.Println(fi14, fi14)
                }
                _ = fi14
            }
        } else {
            v3 = v3
        }
        fmt.Println(v2)
        if true {
            var v39 uint32 = 4294967295
            if false {
                switch 32768 {
                case -9223372036854775808:
                    v39 -= 2147483647
                case 127:
                    v2 += v2
                default:
                    v1 = append(v1, 655)
                }
            } else {
                if len(v1) > 0 {
                    v1[0] = 128
                }
                for fi40 := 0; fi40 < 7; fi40++ {
                    var v41 uint32 = v39
                    var v42 uint32 = (v41 % (((v39 << ((2) & 63))) | 1))
                    v3 = (v3 | v3)
                    _ = fi40
                    _ = v41
                    _ = v42
                }
            }
            _ = v39
        } else {
            _ = len(v1)
            v1 = append(v1, 32768)
        }
    }
    v3 += v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
