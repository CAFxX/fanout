package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int64 }
func genFn0() uint8 {
    switch "" {
    case "a":
        v6 := GenS1{F0: 128, F1: 26}
        _ = v6
    case "\n\t":
        if (!((false && (127 == 600)) && (0 <= 42))) {
            var v7 int64 = 9223372036854775807
            var v8 float32 = ((float32(1.5) + 1e10) + 1.5)
            v9 := GenS2{F0: -949, F1: 1, F2: 228}
            _ = v7
            _ = v8
            _ = v9
        }
    default:
        switch 32768 {
        case -626:
            if true {
                if (!true) {
                    var v10 int64 = 42
                    var v11 int32 = 877
                    _ = v10
                    _ = v11
                }
                switch 1 {
                case -423:
                    if false {
                        for fi12 := 0; fi12 < 3; fi12++ {
                            var v13 int = (fi12 ^ fi12)
                            var v14 float64 = 3.14159
                            if true {
                                if true {
                                    v15 := []int{}
                                    var v16 uint64 = 1
                                    _ = v15
                                    _ = v16
                                }
                                if (("\n\t" + (("foo bar" + "a") + "hello")) <= "") {
                                    var v17 int32 = 42
                                    _ = v17
                                } else {
                                    v18 := []int{-9223372036854775808, (^(v13)), v13}
                                    _ = v18
                                }
                            }
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                    } else {
                        var v19 float64 = ((1e10 + 0.0) + 3.14159)
                        _ = v19
                    }
                case 881:
                    for fi20 := 0; fi20 < 10; fi20++ {
                        fmt.Println(fi20)
                        _ = fi20
                    }
                default:
                    var v21 uint32 = 65535
                    _ = v21
                }
                var v22 int64 = 42
                _ = v22
            }
        case 1:
            for fi23 := 0; fi23 < 14; fi23++ {
                v24 := []int{2147483647}
                fmt.Println(fi23)
                _ = fi23
                _ = v24
            }
        }
    }
    var v25 int16 = 1000
    var v26 int8 = 72
    switch ((("\n\t" + "foo bar") + "") + "hello") {
    case "":
        var v27 string = ((("\n\t" + "\n\t") + "") + ("\n\t" + "foo bar"))
        _ = v27
    }
    _ = v25
    _ = v26
    return 42
}
func genFn1(p0 uint32) int32 {
    v28 := GenS1{F0: 729, F1: p0}
    var v29 bool = (127 > 101)
    for fi30 := 0; fi30 < 11; fi30++ {
        p0 = p0
        _ = fi30
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        var v32 int = (-(fi31))
        _ = fi31
        _ = v32
    }
    v33 := []int{0, -9223372036854775808}
    _ = p0
    _ = v28
    _ = v29
    _ = v33
    return -497
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]bool)
    v3 := make(map[string]int)
    var v4 uint32 = 1
    var v5 int8 = 10
    _ = len(v2)
    v4 += (v4 >> ((1) & 63))
    v4 = v4
    v5 += v5
    if (0 != 42) {
        v4 -= (v4 & 2147483647)
        fmt.Println(v4, v4)
        v34 := GenS3{F0: -1}
        _ = v34
    } else {
        delete(v2, 127)
        for fi35 := 0; fi35 < 13; fi35++ {
            v1[v4] = "foo bar"
            var v36 float32 = (0.1 / -2.25)
            _ = fi35
            _ = v36
        }
        var v37 int64 = -1
        _ = v37
    }
    v38 := []int{128, -9223372036854775808}
    _ = len(v38)
    fmt.Println(v5)
    _ = len(v3)
    var v39 string = (("a" + (("foo bar" + (("\n\t" + "a") + "foo bar")) + "hello")) + (("\n\t" + "hello") + (("foo bar" + "\n\t") + "foo bar")))
    delete(v3, (v39 + v39))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
