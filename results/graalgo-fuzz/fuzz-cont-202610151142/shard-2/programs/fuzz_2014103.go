package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 bool; F1 uint; F2 uint8 }
func genFn0(p0 uint, p1 bool) int32 {
    for fi6 := 0; fi6 < 5; fi6++ {
        if (559 < 42) {
            for fi7 := 0; fi7 < 3; fi7++ {
                var v8 int16 = 32767
                _ = fi7
                _ = v8
            }
            p0 = p0
        } else {
            v9 := GenS1{F0: false}
            v10 := GenS2{F0: p1, F1: 0}
            _ = v10
            _ = v9
        }
        v11 := []int{fi6, (fi6 * fi6), (fi6 >> (((p0 / ((p0) | 1))) & 63))}
        var v12 float32 = 0.1
        _ = fi6
        _ = v11
        _ = v12
    }
    var v13 int16 = 32767
    v14 := GenS2{F0: p1, F1: 72}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 1
}
func genFn1(p0 uint32, p1 uint32) int32 {
    var v15 int32 = genFn0(42, (1.5 < (1e10 / (float64(1.5) + (1.5 / 0.1)))))
    fmt.Println(p1)
    for fi16 := 0; fi16 < 12; fi16++ {
        switch ("foo bar" + "foo bar") {
        case "hello":
            v17 := []int{-2147483648, (fi16 * (fi16 & fi16)), (-(fi16))}
            _ = v17
        case "a":
            fmt.Println(v15, p1)
        }
        v18 := []int{-1, fi16}
        v15 += v15
        _ = fi16
        _ = v18
    }
    var v19 bool = ((0.1 > 0.0) || true)
    _ = p0
    _ = p1
    _ = v15
    _ = v19
    return (v15 % (((v15 & v15)) | 1))
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{1}
    v3 := []string{"", "\n\t"}
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]uint32)
    var v20 bool = true
    for fi21 := 0; fi21 < 8; fi21++ {
        delete(v5, 65535)
        v20 = v20
        v20 = (18446744073709551615 < 170)
        _ = fi21
    }
    switch -1 {
    case 128:
        for fi22 := 0; fi22 < 19; fi22++ {
            v20 = v20
            v20 = true
            _ = fi22
        }
    case -9223372036854775808:
        v20 = v20
    case 32768:
        var v23 int = -2147483648
        _ = v23
    }
    for fi24 := 0; fi24 < 17; fi24++ {
        v20 = (1e10 >= 3.14159)
        _ = fi24
    }
    if v20 {
        fmt.Println(v20)
    }
    _ = len(v1)
    if (-1 < 1) {
        var v25 float32 = (3.14159 / 1.0)
        _ = v25
    }
    if v20 {
        v26 := GenS2{F0: v20, F1: -128}
        if (-2147483648 < -184) {
            if v20 {
                if v20 {
                    for fi27 := 0; fi27 < 2; fi27++ {
                        var v28 uint8 = 1
                        var v29 float64 = 1.5
                        if true {
                            var v30 uint8 = (^(v28))
                            for fi31 := 0; fi31 < 14; fi31++ {
                                var v32 bool = v20
                                _ = fi31
                                _ = v32
                            }
                            var v33 int32 = 2147483647
                            _ = v30
                            _ = v33
                        } else {
                            v20 = (354 <= 742)
                            if (((0.1 / float32(float64(1e10))) / 1.0) >= -2.25) {
                                fmt.Println(v20, v28, v29)
                            }
                            v28 -= (v28 >> ((379) & 63))
                        }
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                }
                fmt.Println(v20)
                if v20 {
                    if v20 {
                        if v20 {
                            v20 = (!(307 != 32767))
                        }
                    } else {
                        var v34 int8 = -53
                        _ = len(v5)
                        var v35 uint32 = 2147483647
                        _ = v34
                        _ = v35
                    }
                    var v36 uint16 = 65535
                    v20 = ("\n\t" < "hello")
                    _ = v36
                }
            } else {
                if true {
                    if v20 {
                        v20 = (255 == 255)
                    }
                }
                for fi37 := 0; fi37 < 11; fi37++ {
                    if (!((1 < 0) || (false && (!v20)))) {
                        v20 = false
                        v38 := GenS3{F0: (!(!v20)), F1: 255, F2: 70}
                        _ = v38
                    } else {
                        var v39 uint64 = 0
                        var v40 float64 = 0.1
                        fmt.Println(fi37, v39, v39)
                        _ = v39
                        _ = v40
                    }
                    _ = fi37
                }
                v41 := GenS1{F0: (3.14159 == ((float64(0.1) * float64(float32(0.1))) / 1.0))}
                _ = v41
            }
        }
        _ = v26
    }
    var v42 uint64 = 279
    if false {
        var v43 int64 = -9223372036854775808
        var v44 int32 = 42
        _ = v43
        _ = v44
    }
    v20 = v20
    for fi45 := 0; fi45 < 18; fi45++ {
        v42 -= (v42 % ((v42) | 1))
        _ = fi45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
