package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint32; F1 string }
func genFn0(p0 int64, p1 string, p2 uint16) int32 {
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := []int{fi7}
        if true {
            v8 = append(v8, fi7)
        } else {
            p2 = p2
            if true {
                v8 = append(v8, int(p0))
                for fi9 := 0; fi9 < 7; fi9++ {
                    fmt.Println(p0, p1, p2)
                    fmt.Println(p1, p2)
                    _ = fi9
                }
            } else {
                p1 = p1
                var v10 int16 = 130
                var v11 uint32 = 2
                _ = v10
                _ = v11
            }
        }
        var v12 int64 = 1
        _ = fi7
        _ = v12
        _ = v8
    }
    p0 = p0
    var v13 int8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return -887
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{-147, -2147483648}
    var v3 uint64 = 4294967295
    var v4 string = "hello"
    var v5 uint32 = uint32(v3)
    var v6 bool = false
    var v14 float64 = -2.25
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v15 int8 = 90
    var v16 uint = 1
    fmt.Println(v6, v15)
    fmt.Println(v4)
    if v6 {
        fmt.Println(v14, v15, v16)
        var v17 int64 = -9223372036854775808
        for fi18 := 0; fi18 < 10; fi18++ {
            if (v4 < v4) {
                var v19 float64 = v14
                if v6 {
                    fmt.Println(v19)
                    for fi20 := 0; fi20 < 20; fi20++ {
                        switch "hello" {
                        case "\n\t":
                            v17 += (v17 / (((v17 << ((463) & 63))) | 1))
                            if ((v3 / ((1) | 1)) >= (v3 << ((v16) & 63))) {
                                fmt.Println(fi18, v4)
                                var v21 uint16 = 355
                                if v6 {
                                    v22 := GenS1{F0: 352}
                                    v4 = ((("" + "foo bar") + ("hello" + ("hello" + "foo bar"))) + v4)
                                    _ = v22
                                }
                                _ = v21
                            }
                        case "hello":
                            v23 := GenS2{F0: 42, F1: 935, F2: (v16 + v16)}
                            _ = v23
                        case "":
                            if ((!v6) || ((v5 | (v5 << ((v16) & 63))) >= (^(v5)))) {
                                if (1 == 32767) {
                                    fmt.Println(fi20, fi20, v5)
                                    var v24 int8 = 42
                                    if (float64((1e10 * float32(v14))) == v19) {
                                        if (false || (!false)) {
                                            v25 := GenS3{F0: 1, F1: 0}
                                            v4 = ("a" + ((v4 + v4) + v4))
                                            _ = v25
                                        }
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = (fi18 ^ fi20)
                                        }
                                        var v26 int16 = -32768
                                        _ = v26
                                    }
                                    _ = v24
                                }
                            }
                        }
                        _ = fi20
                    }
                }
                fmt.Println(v16, fi18)
                _ = v19
            } else {
                var v27 int32 = -51
                _ = v27
            }
            _ = fi18
        }
        _ = v17
    } else {
        if v6 {
            if (v3 != v3) {
                v4 = "foo bar"
                var v28 int64 = -1
                var v29 uint32 = v5
                _ = v28
                _ = v29
            }
        }
        _ = len(v1)
        var v30 uint8 = 83
        _ = v30
    }
    if true {
        v31 := GenS4{F0: v5, F1: ""}
        var v32 uint64 = 18446744073709551615
        _ = v31
        _ = v32
    } else {
        v16 -= (v16 / (((v16 + 2)) | 1))
        v2 = append(v2, 42)
        for fi33 := 0; fi33 < 9; fi33++ {
            _ = len(v2)
            _ = fi33
        }
    }
    var v34 uint64 = (v3 ^ (^(v3)))
    var v35 uint = v16
    var v36 string = v4
    fmt.Println(v4)
    if len(v2) > 0 {
        v2[0] = -1
    }
    for fi37 := 0; fi37 < 14; fi37++ {
        switch (v4 + v36) {
        case "":
            v38 := GenS2{F0: 1, F1: 942, F2: v16}
            _ = v38
        }
        v4 = (v4 + ((v36 + ((((("a" + (v4 + "foo bar")) + "\n\t") + "a") + v4) + (("a" + "\n\t") + (v36 + v4)))) + ("a" + v4)))
        _ = fi37
    }
    var v39 int32 = genFn0(1, "", 2)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi40 := 0; fi40 < 5; fi40++ {
        var v41 float32 = 0.0
        _ = fi40
        _ = v41
    }
    v4 = (v36 + ((v4 + (v36 + "foo bar")) + (v36 + ("hello" + v4))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
