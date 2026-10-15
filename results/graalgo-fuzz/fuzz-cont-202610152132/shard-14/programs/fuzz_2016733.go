package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int16; F1 int }
func genFn0(p0 uint) uint8 {
    v6 := GenS1{F0: p0, F1: 1}
    p0 -= p0
    fmt.Println(p0)
    _ = p0
    _ = v6
    return 228
}

func main() {
    v1 := []uint32{}
    v2 := []int{4294967295}
    v3 := make(map[int]bool)
    var v4 float32 = 3.14159
    var v5 float64 = 0.0
    v7 := GenS1{F0: 4294967295, F1: 9223372036854775807}
    v4 = 0.0
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 uint8 = 42
        for fi10 := 0; fi10 < 8; fi10++ {
            if (true || ("\n\t" < "hello")) {
                v4 = ((v4 * v4) * v4)
                switch (((("hello" + ("" + "foo bar")) + ("\n\t" + "\n\t")) + "a") + "foo bar") {
                case "a":
                    if true {
                        v11 := GenS1{F0: 1, F1: -498}
                        _ = v11
                    }
                case "\n\t":
                    var v12 int64 = 42
                    _ = v12
                case "hello":
                    if true {
                        var v13 bool = false
                        v1 = append(v1, 42)
                        if (((v4 / v4) + 3.14159) == float32((v5 + v5))) {
                            v14 := GenS1{F0: 976, F1: 844}
                            _ = v14
                        }
                        _ = v13
                    }
                }
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if false {
                fmt.Println(v4)
                var v15 float32 = -2.25
                var v16 float32 = (float32(1.5) - -2.25)
                _ = v15
                _ = v16
            }
            _ = fi10
        }
        if (-9223372036854775808 >= 42) {
            v4 = float32(v5)
            v4 = (v4 - 3.14159)
            var v17 string = "foo bar"
            _ = v17
        } else {
            v18 := GenS2{F0: v9, F1: "a"}
            var v19 uint64 = 42
            v5 += float64(v4)
            _ = v18
            _ = v19
        }
        _ = fi8
        _ = v9
    }
    var v20 uint16 = 42
    v5 = (v5 - v5)
    v21 := GenS2{F0: 23, F1: ("\n\t" + (("hello" + "foo bar") + ("\n\t" + (("\n\t" + "\n\t") + "hello"))))}
    v20 -= v20
    for fi22 := 0; fi22 < 10; fi22++ {
        for fi23 := 0; fi23 < 5; fi23++ {
            if true {
                v4 = (v4 - v4)
                if true {
                    var v24 uint8 = 240
                    if true {
                        v3[fi22] = (480 <= 697)
                        var v25 uint64 = 0
                        var v26 int32 = -969
                        _ = v25
                        _ = v26
                    } else {
                        if (0 > 15) {
                            var v27 uint16 = (v20 - (v20 << ((0) & 63)))
                            _ = v27
                        }
                        v5 = -2.25
                        fmt.Println(fi22)
                        var v28 int8 = 0
                        _ = v28
                    }
                    fmt.Println(fi23, fi22)
                    _ = v24
                } else {
                    v5 -= v5
                    switch (-(fi23)) {
                    case 2147483647:
                        fmt.Println(v5)
                    default:
                        var v29 float64 = 3.14159
                        _ = v29
                    }
                    fmt.Println(fi22, v20)
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                switch fi22 {
                case 1:
                    v30 := GenS1{F0: 18446744073709551615, F1: 1}
                    _ = v30
                case 32767:
                    var v31 uint64 = 1
                    _ = v31
                case 42:
                    switch fi22 {
                    case 32767:
                        v4 = 0.1
                    case -9223372036854775808:
                        v5 = (((v5 / (v5 - 1e10)) + float64(3.14159)) + v5)
                    case 4294967295:
                        switch "a" {
                        case "":
                            if (!(((("a" != ("\n\t" + ("a" + "foo bar"))) || (true || true)) || (false && false)) && true)) {
                                v32 := GenS2{F0: 36, F1: ""}
                                var v33 int16 = 1
                                _ = v32
                                _ = v33
                            }
                        case "a":
                            v20 += (v20 / ((816) | 1))
                        case "\n\t":
                            var v34 uint32 = 211
                            _ = v34
                        }
                    }
                default:
                    var v35 float32 = v4
                    _ = v35
                }
                v36 := GenS3{F0: 1, F1: fi23}
                fmt.Println(fi22, fi22, fi22)
                var v37 int32 = -1
                v4 += (1.5 + v4)
                var v38 bool = true
                _ = v36
                _ = v37
                _ = v38
            }
            v5 += -2.25
            fmt.Println(v5)
            v4 = 0.1
            _ = fi23
        }
        for fi39 := 0; fi39 < 20; fi39++ {
            fmt.Println(fi39, v20)
            var v40 float64 = float64(v4)
            if false {
                var v41 bool = (false && (true || true))
                _ = v41
            }
            _ = fi39
            _ = v40
        }
        _ = fi22
    }
    if true {
        var v42 uint8 = genFn0(1)
        _ = v42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
