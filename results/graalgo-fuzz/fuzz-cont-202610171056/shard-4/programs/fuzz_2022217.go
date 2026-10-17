package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int32 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS3 struct { F0 int32; F1 int32 }
func genFn0(p0 int64, p1 uint16, p2 uint) uint64 {
    v5 := []int{-486, 9223372036854775807, int(p0)}
    var v6 int32 = 761
    for fi7 := 0; fi7 < 5; fi7++ {
        p1 = (p1 << (((p2 + p2)) & 63))
        fmt.Println(fi7)
        var v8 bool = (!((4294967295 >= 0) && true))
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1(p0 uint) bool {
    p0 += p0
    switch "a" {
    case "\n\t":
        var v9 uint = (p0 - p0)
        _ = v9
    }
    v10 := GenS1{F0: 65535, F1: -2147483648, F2: -248}
    fmt.Println(p0)
    _ = p0
    _ = v10
    return (!false)
}

func main() {
    v1 := []int{-366, 127}
    var v2 int64 = 1
    v3 := make(map[string]int)
    v4 := make(map[string]uint32)
    var v11 uint64 = 4294967295
    switch (("\n\t" + "a") + ("\n\t" + "\n\t")) {
    case "a":
        v2 -= v2
    }
    v11 += 18446744073709551615
    if (((4294967295 > -719) && (2 <= 991)) && (2147483648 > 315)) {
        v2 = v2
        v11 -= v11
        var v12 bool = (true || true)
        _ = v12
    } else {
        if true {
            for fi13 := 0; fi13 < 17; fi13++ {
                v4["\n\t"] = 579
                v14 := GenS2{F0: 2147483647, F1: 100, F2: (v11 | 4294967295)}
                _ = fi13
                _ = v14
            }
            v11 = v11
            for fi15 := 0; fi15 < 4; fi15++ {
                v3["foo bar"] = fi15
                var v16 uint32 = 2147483647
                _ = fi15
                _ = v16
            }
        }
    }
    var v17 uint64 = 0
    switch 127 {
    case 2147483647:
        fmt.Println(v11)
    case -1:
        switch "a" {
        case "foo bar":
            fmt.Println(v11, v11, v2)
        case "\n\t":
            v11 += v17
        case "a":
            var v18 uint8 = 6
            _ = v18
        default:
            if ((false && false) && true) {
                v17 += genFn0(v2, 29, 18446744073709551615)
                v11 = (^(v11))
            } else {
                v11 = v17
                for fi19 := 0; fi19 < 13; fi19++ {
                    v11 += 42
                    fmt.Println(v2)
                    var v20 int = (fi19 ^ (fi19 / ((32768) | 1)))
                    _ = fi19
                    _ = v20
                }
                for fi21 := 0; fi21 < 6; fi21++ {
                    if true {
                        var v22 float64 = 3.14159
                        v2 = v2
                        fmt.Println(v2, v22)
                        _ = v22
                    } else {
                        v2 += v2
                    }
                    var v23 int16 = -1
                    for fi24 := 0; fi24 < 14; fi24++ {
                        var v25 uint8 = 127
                        _ = fi24
                        _ = v25
                    }
                    _ = fi21
                    _ = v23
                }
            }
        }
    }
    if false {
        v11 = 200
        v4["\n\t"] = 1
    }
    fmt.Println(v2)
    if (v17 < v17) {
        var v26 int8 = 0
        _ = v26
    } else {
        if ("hello" > "") {
            switch "" {
            case "\n\t":
                v27 := GenS1{F0: 4294967295, F1: -142, F2: -1}
                _ = v27
            }
            _ = len(v4)
        }
        v28 := GenS2{F0: 42, F1: 100, F2: (v11 % ((42) | 1))}
        for fi29 := 0; fi29 < 11; fi29++ {
            v2 += v2
            _ = fi29
        }
        _ = v28
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        if false {
            if true {
                for fi31 := 0; fi31 < 2; fi31++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v32 uint64 = genFn0((-(v2)), 382, 255)
                    _ = fi31
                    _ = v32
                }
                fmt.Println(v2, v2)
                for fi33 := 0; fi33 < 15; fi33++ {
                    v34 := GenS2{F0: 120, F1: 1, F2: v17}
                    var v35 uint64 = (v11 & (v17 | v11))
                    v35 -= (^(v17))
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            } else {
                v36 := GenS2{F0: 0, F1: 127, F2: v17}
                v17 = 1
                _ = v36
            }
            v37 := GenS3{F0: 42, F1: -2147483648}
            _ = v37
        } else {
            fmt.Println(v11, v2)
            _ = len(v1)
            v11 = v17
        }
        _ = fi30
    }
    var v38 int = -392
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v4
}
