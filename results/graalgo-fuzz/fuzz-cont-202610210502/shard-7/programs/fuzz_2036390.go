package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 int8; F1 int8; F2 string }
type GenS5 struct { F0 bool; F1 int64; F2 string }
type GenS6 struct { F0 int; F1 bool }
type GenS7 struct { F0 int }
func genFn0(p0 bool, p1 int, p2 int16) uint32 {
    if (213 > 224) {
        var v6 bool = p0
        var v7 float32 = 0.1
        _ = v6
        _ = v7
    } else {
        if p0 {
            p1 = (-(p1))
            v8 := GenS1{F0: 255, F1: -803, F2: 255}
            var v9 int8 = 42
            _ = v8
            _ = v9
        }
    }
    p0 = ((p2 >= (^(p2))) || ((3.14159 * 1.5) == (1.5 * 0.0)))
    v10 := []int{p1}
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 65535
}

func main() {
    var v1 uint8 = 128
    var v2 uint8 = 2
    v3 := make(map[uint32]string)
    v4 := []string{"foo bar", "\n\t", ("a" + "foo bar")}
    v5 := []string{"foo bar", ("hello" + "\n\t")}
    fmt.Println(v2, v1)
    if true {
        var v11 int = 4294967295
        _ = v11
    }
    v2 = (v1 / ((255) | 1))
    for fi12 := 0; fi12 < 18; fi12++ {
        if (1 != 168) {
            var v13 int64 = -607
            for fi14 := 0; fi14 < 20; fi14++ {
                _ = len(v5)
                fmt.Println(fi12)
                v13 += 1
                _ = fi14
            }
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 int64 = 42
                v17 := GenS2{F0: 5, F1: 2147483647}
                v18 := GenS3{F0: 2, F1: false}
                if true {
                    var v19 uint8 = v1
                    switch "\n\t" {
                    case "\n\t":
                        if (!true) {
                            if len(v5) > 0 {
                                v5[0] = "hello"
                            }
                        }
                    }
                    _ = v19
                } else {
                    v20 := GenS1{F0: 0, F1: 42, F2: 2}
                    v21 := GenS4{F0: 1, F1: 1, F2: ""}
                    _ = v20
                    _ = v21
                }
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            _ = v13
        }
        fmt.Println(v1, fi12)
        _ = fi12
    }
    var v22 uint32 = 824
    for fi23 := 0; fi23 < 15; fi23++ {
        v22 -= 433
        for fi24 := 0; fi24 < 13; fi24++ {
            v1 += (v2 + 0)
            if (!(542 >= 1)) {
                var v25 int = (fi23 % ((fi24) | 1))
                v26 := GenS5{F0: false, F1: 1, F2: "hello"}
                v27 := GenS3{F0: 4294967295, F1: ((v22 / (((v22 % ((v22) | 1))) | 1)) != v22)}
                var v28 float32 = 0.0
                if false {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    switch (-(fi24)) {
                    case 1:
                        v22 -= genFn0((65535 <= (v22 << ((60) & 63))), fi24, -441)
                    case 32767:
                        var v29 uint = 358
                        _ = v29
                    case -1:
                        var v30 int32 = 827
                        _ = v30
                    default:
                        if len(v4) > 0 {
                            v4[0] = ("" + "hello")
                        }
                    }
                } else {
                    v31 := GenS6{F0: (fi23 * (v25 - (^(fi24)))), F1: true}
                    v32 := GenS7{F0: v25}
                    v33 := GenS2{F0: 2, F1: v22}
                    _ = v31
                    _ = v32
                    _ = v33
                }
                _ = v25
                _ = v26
                _ = v27
                _ = v28
            } else {
                var v34 float64 = 0.1
                v2 -= v2
                _ = v34
            }
            v35 := GenS7{F0: fi23}
            _ = fi24
            _ = v35
        }
        _ = fi23
    }
    var v36 int64 = 714
    v1 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
