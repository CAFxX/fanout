package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 uint32; F2 bool }
type GenS3 struct { F0 uint32; F1 uint64; F2 string }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 uint; F1 string; F2 uint16 }
type GenS6 struct { F0 uint32; F1 uint16; F2 uint }
func genFn0(p0 uint, p1 uint32, p2 int64) bool {
    if true {
        var v5 float32 = 0.1
        v6 := GenS1{F0: 888}
        v7 := []int{128, 4294967295, 4294967295, -105}
        _ = v5
        _ = v6
        _ = v7
    } else {
        var v8 string = "\n\t"
        _ = v8
    }
    p2 = (-(p2))
    fmt.Println(p0, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return true
}
func genFn1(p0 int16) int16 {
    v9 := GenS1{F0: 4294967295}
    var v10 bool = true
    var v11 uint64 = 4294967295
    fmt.Println(v11, v10, p0)
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return p0
}

func main() {
    v1 := []int{-321, 613, 2147483647, 1}
    v2 := make(map[string]uint32)
    var v3 uint64 = 951
    var v4 int64 = 0
    switch 1 {
    case 32767:
        v2[("hello" + "a")] = 65535
    case 907:
        for fi12 := 0; fi12 < 11; fi12++ {
            v4 += (v4 / ((v4) | 1))
            _ = fi12
        }
    case 0:
        var v13 bool = (!true)
        _ = v13
    default:
        for fi14 := 0; fi14 < 13; fi14++ {
            v3 = 173
            v15 := v2["\n\t"]
            _ = fi14
            _ = v15
        }
    }
    v16 := v2["hello"]
    v16 = v16
    fmt.Println(v16, v16, v4)
    fmt.Println(v3, v3)
    switch (("\n\t" + (("foo bar" + "foo bar") + "\n\t")) + ("\n\t" + ("a" + "\n\t"))) {
    case "hello":
        v4 -= -9223372036854775808
    default:
        v17 := GenS2{F0: (443 == 1), F1: (v16 - 990), F2: false}
        _ = v17
    }
    fmt.Println(v16)
    v1 = append(v1, 32768)
    for fi18 := 0; fi18 < 17; fi18++ {
        for fi19 := 0; fi19 < 1; fi19++ {
            v4 += (^(v4))
            _ = fi19
        }
        switch fi18 {
        case 0:
            if false {
                v16 = (v16 << ((2) & 63))
                var v20 int32 = 1
                var v21 int64 = v4
                _ = v20
                _ = v21
            } else {
                if (v4 == (-(v4))) {
                    if false {
                        fmt.Println(fi18)
                        v16 = v16
                        if (255 <= 2) {
                            v22 := v2["hello"]
                            var v23 int8 = 1
                            fmt.Println(fi18, v22, v4)
                            _ = v22
                            _ = v23
                        }
                    } else {
                        v24 := v2["\n\t"]
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v25 := GenS3{F0: 2147483648, F1: v3, F2: "hello"}
                        _ = v24
                        _ = v25
                    }
                    v3 = v3
                    v16 -= (^(v16))
                }
            }
        case 32767:
            for fi26 := 0; fi26 < 15; fi26++ {
                v27 := GenS2{F0: true, F1: (v16 >> ((0) & 63)), F2: ((!(65535 >= 0)) && true)}
                v3 -= (v3 + (v3 - (v3 << ((2) & 63))))
                var v28 uint8 = uint8(v3)
                _ = fi26
                _ = v27
                _ = v28
            }
        case -9223372036854775808:
            v29 := GenS4{F0: "foo bar", F1: (1.5 == (0.1 + 1e10))}
            _ = v29
        }
        for fi30 := 0; fi30 < 19; fi30++ {
            if false {
                if true {
                    switch "foo bar" {
                    case "":
                        switch (fi30 - -278) {
                        case 79:
                            fmt.Println(v4)
                        case -236:
                            v4 += (v4 / ((-1) | 1))
                        case 127:
                            if false {
                                v3 = 768
                                v3 = 1
                                fmt.Println(fi18)
                            }
                        }
                    case "hello":
                        if (1e10 <= (3.14159 + (1e10 / 1e10))) {
                            v31 := GenS5{F0: 65535, F1: "hello", F2: 347}
                            var v32 float32 = 1e10
                            _ = len(v2)
                            _ = v31
                            _ = v32
                        } else {
                            v33 := v2["foo bar"]
                            fmt.Println(fi18)
                            _ = v33
                        }
                    default:
                        v16 = v16
                    }
                    _ = len(v2)
                } else {
                    if false {
                        var v34 int8 = int8(v4)
                        v16 = (v16 * (v16 + (^(v16))))
                        _ = v34
                    }
                    v35 := GenS2{F0: false, F1: v16, F2: false}
                    _ = v35
                }
                v36 := GenS3{F0: 65535, F1: (v3 / ((v3) | 1)), F2: "hello"}
                v37 := GenS6{F0: 42, F1: 540, F2: 46}
                _ = v36
                _ = v37
            }
            fmt.Println(v4, v16, v3)
            _ = len(v2)
            _ = fi30
        }
        _ = fi18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
