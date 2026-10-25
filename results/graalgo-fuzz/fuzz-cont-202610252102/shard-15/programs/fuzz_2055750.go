package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint; F1 uint8 }
func genFn0(p0 int8) int64 {
    for fi4 := 0; fi4 < 19; fi4++ {
        v5 := []int{(fi4 / (((fi4 ^ fi4)) | 1))}
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: (("a" + (("" + "a") + ("hello" + "foo bar"))) + "foo bar")}
    if true {
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 uint8 = 42
            switch (fi7 >> ((2) & 63)) {
            case 128:
                var v9 int64 = -1
                _ = v9
            case 38:
                v8 -= 183
            case 585:
                v10 := []int{32768, fi7}
                _ = v10
            default:
                if (1.5 != (-2.25 / float64(-2.25))) {
                    if (!(-2.25 != 1e10)) {
                        for fi11 := 0; fi11 < 10; fi11++ {
                            fmt.Println(v8)
                            v12 := GenS1{F0: "hello"}
                            _ = fi11
                            _ = v12
                        }
                        if (!false) {
                            for fi13 := 0; fi13 < 19; fi13++ {
                                v8 = v8
                                var v14 string = "a"
                                fmt.Println(p0)
                                _ = fi13
                                _ = v14
                            }
                            for fi15 := 0; fi15 < 15; fi15++ {
                                p0 = p0
                                _ = fi15
                            }
                            v16 := GenS2{F0: 37}
                            _ = v16
                        } else {
                            switch "hello" {
                            case "":
                                fmt.Println(v8, p0)
                            case "foo bar":
                                if true {
                                    fmt.Println(v8, v8, fi7)
                                    switch ("\n\t" + (((("" + "hello") + "a") + (("hello" + "") + "")) + (((("\n\t" + "hello") + (("a" + "\n\t") + "a")) + "\n\t") + ("foo bar" + "hello")))) {
                                    case "foo bar":
                                        switch (fi7 / (((fi7 ^ -112)) | 1)) {
                                        case -9223372036854775808:
                                            v17 := []int{(^(fi7)), (fi7 + (fi7 << ((1000) & 63))), fi7, fi7}
                                            _ = v17
                                        case -61:
                                            var v18 uint64 = 0
                                            _ = v18
                                        }
                                    case "\n\t":
                                        fmt.Println(v8, fi7)
                                    default:
                                        v8 += (v8 ^ v8)
                                    }
                                } else {
                                    if ("foo bar" > "\n\t") {
                                        v8 += (^(v8))
                                        v19 := GenS3{F0: -1, F1: 32767}
                                        _ = v19
                                    }
                                }
                            }
                        }
                        for fi20 := 0; fi20 < 19; fi20++ {
                            v21 := GenS4{F0: 255, F1: 255}
                            _ = fi20
                            _ = v21
                        }
                    }
                }
            }
            for fi22 := 0; fi22 < 4; fi22++ {
                v23 := []int{(fi7 | fi7)}
                v24 := GenS1{F0: "hello"}
                p0 = p0
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = fi7
            _ = v8
        }
    }
    v25 := GenS4{F0: 0, F1: 169}
    _ = p0
    _ = v25
    _ = v6
    return 0
}

func main() {
    var v1 uint = 65535
    var v2 int = 4294967295
    var v3 float64 = 0.0
    v1 -= (v1 & 308)
    if (false || false) {
        v26 := []int{-271, v2}
        v1 -= v1
        _ = v26
    } else {
        v27 := GenS2{F0: 2}
        for fi28 := 0; fi28 < 3; fi28++ {
            var v29 uint8 = 127
            fmt.Println(v29, v29, fi28)
            if false {
                if false {
                    var v30 uint16 = 0
                    v31 := GenS3{F0: 42, F1: 1}
                    _ = v30
                    _ = v31
                }
            } else {
                var v32 uint32 = 0
                if false {
                    v32 = (v32 ^ v32)
                    var v33 uint32 = (v32 % ((v32) | 1))
                    _ = v33
                }
                v3 = (v3 + 1e10)
                _ = v32
            }
            _ = fi28
            _ = v29
        }
        var v34 uint32 = 527
        _ = v27
        _ = v34
    }
    for fi35 := 0; fi35 < 17; fi35++ {
        v36 := GenS1{F0: "a"}
        _ = fi35
        _ = v36
    }
    var v37 uint16 = 65535
    v38 := GenS3{F0: 768, F1: -898}
    for fi39 := 0; fi39 < 12; fi39++ {
        v40 := GenS1{F0: "\n\t"}
        _ = fi39
        _ = v40
    }
    var v41 float32 = 1e10
    v42 := GenS2{F0: 127}
    v43 := []int{(v2 * (-(v2))), -973}
    v37 -= (v37 ^ 0)
    var v44 uint16 = v37
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
