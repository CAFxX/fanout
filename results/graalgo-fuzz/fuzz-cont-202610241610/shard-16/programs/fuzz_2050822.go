package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 bool; F2 uint16 }
func genFn0(p0 uint32, p1 uint8, p2 int64) int64 {
    var v6 float64 = 0.1
    var v7 uint8 = (p1 + p1)
    var v8 float32 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return p2
}
func genFn1() string {
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 bool = true
        v11 := []int{-819, (fi9 ^ (fi9 | (fi9 / (((fi9 << ((18446744073709551615) & 63))) | 1))))}
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 string = "foo bar"
    _ = v12
    return v12
}

func main() {
    var v1 uint8 = 0
    v2 := []string{("\n\t" + "a"), "", ("" + "foo bar"), "\n\t"}
    var v3 int64 = -713
    var v4 uint32 = 2147483648
    v5 := []string{("hello" + ""), "a"}
    v3 = (v3 | v3)
    if (-959 != 11) {
        var v13 int64 = v3
        if (v4 > v4) {
            if len(v2) > 0 {
                v2[0] = ""
            }
            v14 := GenS1{F0: 0}
            if false {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            _ = v14
        } else {
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 uint64 = 768
                v3 -= v3
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi15
                _ = v16
            }
            v3 = (v13 * (v3 ^ (v3 / (((v13 >> ((321) & 63))) | 1))))
            if (true || false) {
                v5 = append(v5, "foo bar")
                v17 := GenS1{F0: -147}
                var v18 bool = false
                _ = v17
                _ = v18
            }
        }
        _ = v13
    } else {
        var v19 int32 = 1
        var v20 uint16 = 0
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        v1 += v1
        for fi22 := 0; fi22 < 11; fi22++ {
            if false {
                if true {
                    var v23 int8 = 0
                    if false {
                        var v24 int = fi22
                        if (!true) {
                            var v25 float32 = (0.0 / 0.1)
                            v26 := GenS1{F0: 1}
                            switch (fi21 >> ((4294967295) & 63)) {
                            case -904:
                                v23 += (v23 << ((2) & 63))
                            case 805:
                                var v27 bool = true
                                _ = v27
                            default:
                                v5 = append(v5, "foo bar")
                            }
                            _ = v25
                            _ = v26
                        }
                        var v28 float64 = (1e10 * (0.1 * (((3.14159 + 0.1) / 1.0) / (-2.25 + -2.25))))
                        v29 := GenS2{F0: 872}
                        var v30 float64 = v28
                        _ = v24
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                    if false {
                        v23 = v23
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        if true {
                            fmt.Println(v23, fi22, fi22)
                        }
                    }
                    _ = v23
                }
                fmt.Println(fi22, v3)
            }
            _ = fi22
        }
        v4 += 652
        _ = fi21
    }
    v5 = append(v5, "\n\t")
    for fi31 := 0; fi31 < 5; fi31++ {
        switch fi31 {
        case 4294967295:
            v3 = 9223372036854775807
        }
        v3 += -9223372036854775808
        v32 := GenS3{F0: 42, F1: false, F2: 32767}
        _ = fi31
        _ = v32
    }
    if (!(!false)) {
        v4 = v4
    }
    var v33 uint64 = 42
    fmt.Println(v4)
    v3 = v3
    v3 = v3
    for fi34 := 0; fi34 < 4; fi34++ {
        var v35 uint16 = 0
        _ = fi34
        _ = v35
    }
    var v36 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v4
    _ = v5
}
