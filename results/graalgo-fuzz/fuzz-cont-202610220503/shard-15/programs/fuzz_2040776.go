package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 int8; F2 uint64 }
func genFn0(p0 int64) uint {
    fmt.Println(p0)
    if (255 >= 2) {
        if false {
            p0 = p0
        } else {
            var v5 uint8 = 1
            _ = v5
        }
        switch "hello" {
        case "hello":
            var v6 int64 = 1
            _ = v6
        case "a":
            p0 = p0
        }
    }
    if (!(!(("hello" <= ("" + "a")) && false))) {
        var v7 int32 = -1
        for fi8 := 0; fi8 < 18; fi8++ {
            for fi9 := 0; fi9 < 12; fi9++ {
                switch "foo bar" {
                case "hello":
                    if ((!(283 < 234)) && true) {
                        switch ("foo bar" + "foo bar") {
                        case "foo bar":
                            if false {
                                if false {
                                    if true {
                                        if (float64((((((float32(1.5) - 0.1) * float32(float64(0.1))) - (0.1 / float32(1e10))) - 1.5) - 0.0)) != 1e10) {
                                            var v10 float64 = ((-2.25 + 3.14159) * ((0.0 - 0.1) + float64((-2.25 + ((((1e10 + 3.14159) - 3.14159) - 0.0) + 0.0)))))
                                            v10 += 0.1
                                            _ = v10
                                        } else {
                                            var v11 uint16 = 490
                                            v7 = v7
                                            _ = v11
                                        }
                                        var v12 bool = (!true)
                                        _ = v12
                                    } else {
                                        v7 += int32(p0)
                                    }
                                    if false {
                                        var v13 int = 32767
                                        p0 = -1
                                        _ = v13
                                    }
                                }
                            } else {
                                fmt.Println(fi9)
                                var v14 float64 = float64((float32(float64(0.0)) + 1e10))
                                var v15 float64 = v14
                                var v16 int32 = v7
                                v17 := []int{(fi9 | (fi8 / ((fi9) | 1))), fi9, (fi8 * fi9)}
                                _ = v14
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                        case "":
                            var v18 bool = ((127 == 100) && false)
                            _ = v18
                        case "a":
                            v19 := []int{4294967295, (fi9 >> ((2) & 63)), fi9, -1}
                            _ = v19
                        default:
                            var v20 bool = (697 >= 2147483648)
                            _ = v20
                        }
                    }
                case "\n\t":
                    var v21 int64 = (p0 << ((980) & 63))
                    _ = v21
                }
                v22 := []int{}
                _ = fi9
                _ = v22
            }
            _ = fi8
        }
        v7 = (v7 & v7)
        _ = v7
    }
    _ = p0
    return 42
}
func genFn1(p0 bool, p1 uint8) bool {
    var v23 int64 = -1
    v24 := []int{}
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    return p0
}

func main() {
    var v1 uint8 = 183
    var v2 uint8 = (v1 / (((v1 * v1)) | 1))
    v3 := make(map[string]string)
    var v4 int16 = 237
    v25 := []int{int(v4), 4294967295, 128, 98}
    v2 += 2
    for fi26 := 0; fi26 < 7; fi26++ {
        if false {
            v27 := GenS1{F0: 18}
            _ = v27
        }
        if (true || false) {
            v4 += v4
        }
        _ = fi26
    }
    _ = len(v25)
    v28 := v3[((("" + ("foo bar" + (("foo bar" + "\n\t") + "foo bar"))) + "") + ("hello" + "\n\t"))]
    switch (v28 + "hello") {
    case "foo bar":
        if false {
            var v29 uint64 = uint64(v1)
            _ = v29
        }
    }
    var v30 int = -532
    v25 = append(v25, 664)
    var v31 uint8 = v2
    v32 := GenS2{F0: 85, F1: 100, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
