package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 bool }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int8 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 int64; F1 int16; F2 int }
type GenS6 struct { F0 uint; F1 int }
func genFn0(p0 int16) string {
    var v7 uint64 = 42
    v8 := GenS1{F0: 127, F1: "", F2: ((v7 * v7) != v7)}
    v7 -= (v7 - v7)
    switch 748 {
    case -450:
        for fi9 := 0; fi9 < 17; fi9++ {
            for fi10 := 0; fi10 < 9; fi10++ {
                fmt.Println(fi9, fi9)
                _ = fi10
            }
            _ = fi9
        }
    case 584:
        p0 = -32768
    case 32767:
        var v11 float64 = 0.1
        _ = v11
    default:
        fmt.Println(p0, v7)
    }
    _ = p0
    _ = v7
    _ = v8
    return ""
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = -869
    var v3 string = (("foo bar" + "") + (("" + ("\n\t" + ("" + ("\n\t" + ("" + "foo bar"))))) + ("foo bar" + "hello")))
    var v4 uint64 = 18446744073709551615
    v5 := []int{0}
    v6 := []int{-987, 9223372036854775807, 42}
    var v12 int16 = 841
    switch (v3 + v3) {
    case "hello":
        var v13 float64 = 1.5
        _ = v13
    case "\n\t":
        for fi14 := 0; fi14 < 4; fi14++ {
            v3 = ((("foo bar" + "\n\t") + (v3 + "foo bar")) + (v3 + ((v3 + "foo bar") + ("" + (v3 + "foo bar")))))
            switch (((v3 + v3) + v3) + (v3 + v3)) {
            case "a":
                if (652 >= (fi14 >> ((686) & 63))) {
                    if true {
                        if false {
                            v15 := GenS2{F0: 65535, F1: 65535}
                            _ = v15
                        }
                    } else {
                        for fi16 := 0; fi16 < 1; fi16++ {
                            switch "foo bar" {
                            case "\n\t":
                                if false {
                                    if ((!true) && false) {
                                        var v17 uint32 = 2
                                        var v18 uint8 = 127
                                        v17 = 4294967295
                                        fmt.Println(fi16)
                                        _ = v17
                                        _ = v18
                                    }
                                    var v19 int8 = -4
                                    v20 := GenS3{F0: (^(v4)), F1: 0, F2: v19}
                                    _ = v19
                                    _ = v20
                                } else {
                                    if (65535 >= 65535) {
                                        if (true || (127 != 42)) {
                                            if ((v2 - v2) >= (v2 + (-(v2)))) {
                                                delete(v1, (fi16 << ((18446744073709551615) & 63)))
                                                var v21 int = (fi14 & 9223372036854775807)
                                                _ = v21
                                            } else {
                                                v5 = append(v5, fi14)
                                            }
                                            switch fi14 {
                                            case 32767:
                                                var v22 bool = (0 == -128)
                                                _ = v22
                                            case -9223372036854775808:
                                                v23 := v1[fi14]
                                                _ = v23
                                            case -7:
                                                v24 := GenS4{F0: 541, F1: ((-32768 >= (-(v12))) || false)}
                                                _ = v24
                                            default:
                                                var v25 int32 = 0
                                                _ = v25
                                            }
                                        } else {
                                            var v26 int32 = 0
                                            _ = v26
                                        }
                                    }
                                    var v27 int64 = -9223372036854775808
                                    _ = v27
                                }
                            }
                            _ = fi16
                        }
                        v4 -= 4294967295
                    }
                }
            case "\n\t":
                if true {
                    _ = len(v6)
                } else {
                    v28 := GenS5{F0: (v2 + (v2 ^ v2)), F1: v12, F2: 2147483647}
                    fmt.Println(v2, fi14)
                    var v29 uint64 = 261
                    _ = v28
                    _ = v29
                }
            }
            var v30 uint = 600
            _ = fi14
            _ = v30
        }
    default:
        var v31 float64 = 1.5
        _ = v31
    }
    fmt.Println(v3, v3)
    v3 = (("foo bar" + "a") + v3)
    if true {
        var v32 uint32 = 988
        if true {
            v6 = append(v6, 127)
        } else {
            if (100 != -106) {
                for fi33 := 0; fi33 < 5; fi33++ {
                    v3 = v3
                    v2 -= v2
                    delete(v1, 32768)
                    _ = fi33
                }
                v6 = append(v6, 32768)
            } else {
                fmt.Println(v32, v12, v2)
            }
        }
        var v34 int32 = -118
        _ = v32
        _ = v34
    }
    v35 := v1[42]
    switch ((v35 + v3) + "\n\t") {
    case "":
        v2 = -1
    }
    v36 := GenS6{F0: 18446744073709551615, F1: 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
