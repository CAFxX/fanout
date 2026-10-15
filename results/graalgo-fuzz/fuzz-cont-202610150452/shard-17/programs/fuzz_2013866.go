package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 string; F2 int }
type GenS3 struct { F0 uint8; F1 uint64; F2 int32 }
func genFn0(p0 uint16, p1 uint16, p2 int32) string {
    fmt.Println(p1, p1)
    var v6 uint8 = 128
    v6 -= 42
    v7 := GenS1{F0: (p1 ^ 24)}
    v8 := []int{32767, int(p1), 490, -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return "a"
}
func genFn1(p0 int16, p1 uint16) uint16 {
    v9 := []int{464, 175, 4294967295}
    var v10 uint16 = (p1 / ((p1) | 1))
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return v10
}

func main() {
    v1 := []string{"a", "a", "hello"}
    v2 := []uint32{42, 296}
    var v3 float64 = -2.25
    var v4 string = ("a" + "")
    v5 := make(map[string]string)
    var v11 float32 = 0.1
    v3 = v3
    fmt.Println(v3, v3, v11)
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 10; fi13++ {
            if false {
                if false {
                    v4 = v4
                    var v14 bool = false
                    if v14 {
                        v3 = ((0.1 - (v3 * (v3 + -2.25))) - v3)
                        switch -417 {
                        case 9223372036854775807:
                            fmt.Println(fi12, fi13, v4)
                        case -2147483648:
                            var v15 uint = 101
                            _ = v15
                        default:
                            fmt.Println(v3)
                        }
                    } else {
                        v5[("hello" + v4)] = v4
                        var v16 uint = 0
                        v16 = v16
                        switch fi12 {
                        case 1:
                            fmt.Println(v4, v4, fi13)
                        case -2147483648:
                            if (((!(65535 == 410)) || v14) && (1 <= 0)) {
                                if v14 {
                                    v3 += (v3 + (v3 + -2.25))
                                } else {
                                    _ = len(v5)
                                    if (4294967295 == 0) {
                                        var v17 bool = (1000 > -341)
                                        _ = v17
                                    }
                                    _ = len(v1)
                                }
                                var v18 int8 = 42
                                var v19 float32 = 0.1
                                _ = v18
                                _ = v19
                            }
                        }
                        _ = v16
                    }
                    v4 = "a"
                    delete(v5, (("" + "hello") + ""))
                    var v20 int32 = 381
                    switch (fi12 ^ (fi13 & fi12)) {
                    case 32768:
                        v21 := v5[v4]
                        _ = v21
                    case 4294967295:
                        v20 += v20
                    }
                    _ = v14
                    _ = v20
                } else {
                    switch fi13 {
                    case 128:
                        if ((!true) || (((0.1 == (float64(3.14159) / float64(float32(v3)))) || (v11 != (v11 + v11))) && (63 >= 1))) {
                            v3 -= 0.1
                            var v22 uint16 = 779
                            switch (fi13 ^ fi13) {
                            case -93:
                                var v23 uint32 = 2
                                _ = v23
                            case -2147483648:
                                var v24 bool = false
                                v25 := GenS1{F0: (v22 * genFn1(1000, (v22 | (v22 >> ((255) & 63)))))}
                                _ = v24
                                _ = v25
                            }
                            _ = v22
                        } else {
                            fmt.Println(v11, v4, v4)
                        }
                    case 127:
                        v5[v4] = "foo bar"
                    case 2147483647:
                        v4 = (v4 + v4)
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    fmt.Println(v11, fi13, v11)
                    if (-246 > fi13) {
                        if ((fi13 / ((fi13) | 1)) != fi13) {
                            var v26 uint = 255
                            _ = v26
                        }
                        v27 := GenS1{F0: 32767}
                        v4 = v4
                        if true {
                            fmt.Println(fi12, v4, v3)
                            var v28 bool = (0 > -122)
                            _ = v28
                        } else {
                            var v29 uint32 = 2
                            _ = v29
                        }
                        if (1000 <= -320) {
                            v30 := GenS2{F0: 2, F1: (v4 + v4), F2: (fi13 % ((32767) | 1))}
                            _ = v30
                        }
                        if false {
                            v3 = 3.14159
                        } else {
                            fmt.Println(fi12, fi13)
                        }
                        _ = v27
                    } else {
                        if (808 > 0) {
                            v11 = v11
                            var v31 uint8 = 127
                            _ = v31
                        }
                        v32 := GenS3{F0: 1, F1: 830, F2: 0}
                        _ = v32
                    }
                    fmt.Println(fi13, v3)
                    v33 := v5[v4]
                    var v34 bool = (true || (false || true))
                    _ = v33
                    _ = v34
                }
            } else {
                if (!true) {
                    v4 = "hello"
                    var v35 float64 = 0.1
                    var v36 int = fi12
                    _ = v35
                    _ = v36
                }
                fmt.Println(v4)
                switch v4 {
                case "\n\t":
                    fmt.Println(v4)
                case "":
                    fmt.Println(v11)
                case "a":
                    var v37 uint8 = 2
                    _ = v37
                default:
                    if true {
                        var v38 bool = false
                        delete(v5, v4)
                        _ = v38
                    }
                }
            }
            _ = fi13
        }
        _ = fi12
    }
    var v39 int16 = 1000
    fmt.Println(v4)
    switch v4 {
    case "a":
        var v40 float32 = float32(-2.25)
        _ = v40
    }
    v4 = "a"
    _ = len(v5)
    v3 = (0.0 * 1e10)
    fmt.Println(v3)
    if true {
        v11 += (((1e10 + v11) + float32(float64(v11))) + float32(-2.25))
        v11 += v11
    }
    v41 := GenS2{F0: 42, F1: "", F2: -2147483648}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v41
    _ = v5
}
