package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int64, p1 uint64) int8 {
    var v5 uint16 = 32767
    var v6 int8 = -43
    v6 -= (v6 << ((1) & 63))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (-(v6))
}
func genFn1(p0 int8, p1 bool, p2 int32) uint16 {
    var v7 int64 = 42
    p2 -= (p2 % (((p2 + (p2 * p2))) | 1))
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{"\n\t"}
    var v3 float32 = 0.0
    var v4 int32 = 0
    var v8 uint64 = 2
    v8 = v8
    var v9 uint8 = 128
    if (532 >= 732) {
        v9 = v9
        v9 = 85
        if false {
            var v10 float64 = 3.14159
            var v11 string = ""
            v11 = v11
            _ = v10
            _ = v11
        } else {
            if true {
                if false {
                    var v12 uint64 = (^(v8))
                    v3 = float32(float64((0.1 * float32(0.1))))
                    delete(v1, "\n\t")
                    _ = v12
                } else {
                    var v13 int16 = 1000
                    _ = v13
                }
            } else {
                delete(v1, (("a" + "foo bar") + "\n\t"))
                var v14 uint64 = 1
                _ = v14
            }
            var v15 int64 = -1
            _ = v15
        }
    } else {
        v8 += v8
        var v16 uint16 = 65535
        _ = v16
    }
    v17 := GenS1{F0: 1, F1: v9, F2: 42}
    v18 := v1["\n\t"]
    if false {
        var v19 int = v18
        _ = v19
    } else {
        switch "" {
        case "foo bar":
            if ((1e10 / 0.1) < 0.0) {
                var v20 int16 = 477
                if (!(18446744073709551615 > 65535)) {
                    if (((v20 < v20) || false) || true) {
                        v4 += v4
                        if false {
                            var v21 int = 1
                            var v22 float32 = ((1e10 * (v3 * (0.0 + v3))) / 1.0)
                            v23 := GenS2{F0: 299}
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                    } else {
                        for fi24 := 0; fi24 < 1; fi24++ {
                            var v25 int = (-(v18))
                            var v26 int16 = v20
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    }
                    for fi27 := 0; fi27 < 16; fi27++ {
                        v3 -= 0.1
                        var v28 int64 = int64(v20)
                        _ = fi27
                        _ = v28
                    }
                } else {
                    var v29 string = "\n\t"
                    fmt.Println(v9)
                    if (true || true) {
                        var v30 uint32 = 4294967295
                        var v31 int = (v18 / ((v18) | 1))
                        _ = v30
                        _ = v31
                    } else {
                        var v32 uint = 2
                        v4 += 2147483647
                        v3 += v3
                        _ = v32
                    }
                    _ = v29
                }
                _ = v20
            }
        case "":
            var v33 int16 = 0
            _ = v33
        default:
            for fi34 := 0; fi34 < 11; fi34++ {
                var v35 uint8 = 137
                var v36 int16 = -32768
                _ = fi34
                _ = v35
                _ = v36
            }
        }
        for fi37 := 0; fi37 < 7; fi37++ {
            var v38 int64 = 1
            var v39 string = ("hello" + "")
            v39 = (v39 + (v39 + v39))
            _ = fi37
            _ = v38
            _ = v39
        }
        switch (v18 << ((793) & 63)) {
        case -1:
            _ = len(v1)
        case -785:
            if (!(783 == 0)) {
                fmt.Println(v18)
                for fi40 := 0; fi40 < 4; fi40++ {
                    v41 := GenS1{F0: 2, F1: v9, F2: -9223372036854775808}
                    if (117 < uint(v9)) {
                        fmt.Println(v3)
                    }
                    for fi42 := 0; fi42 < 19; fi42++ {
                        v8 = 541
                        v8 = (v8 - v8)
                        v8 = v8
                        _ = fi42
                    }
                    _ = fi40
                    _ = v41
                }
            } else {
                var v43 string = "foo bar"
                _ = v43
            }
        }
    }
    if true {
        v44 := GenS3{F0: 27}
        var v45 uint32 = 545
        _ = v44
        _ = v45
    }
    var v46 int64 = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v8
    _ = v9
}
