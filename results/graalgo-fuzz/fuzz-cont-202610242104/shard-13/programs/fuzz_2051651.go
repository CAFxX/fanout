package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 string; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint64) int8 {
    fmt.Println(p0)
    switch -569 {
    case 32767:
        var v4 float32 = float32((1.5 - 3.14159))
        _ = v4
    case 0:
        var v5 int = 42
        _ = v5
    default:
        var v6 uint = 2
        _ = v6
    }
    var v7 uint64 = p0
    if (127 >= 127) {
        fmt.Println(p0, v7)
        if (9223372036854775807 <= 1) {
            v7 = p0
            if (3.14159 <= 1.5) {
                v8 := GenS1{F0: -128}
                if false {
                    p0 += v7
                    for fi9 := 0; fi9 < 6; fi9++ {
                        var v10 int16 = -32768
                        _ = fi9
                        _ = v10
                    }
                } else {
                    for fi11 := 0; fi11 < 8; fi11++ {
                        for fi12 := 0; fi12 < 2; fi12++ {
                            p0 = v7
                            p0 -= p0
                            fmt.Println(fi11, fi12)
                            var v13 bool = (!false)
                            _ = fi12
                            _ = v13
                        }
                        _ = fi11
                    }
                    p0 = (p0 % ((v7) | 1))
                }
                v14 := GenS2{F0: p0, F1: 0}
                _ = v14
                _ = v8
            } else {
                if false {
                    var v15 int = -2147483648
                    _ = v15
                }
                v7 += (p0 % ((p0) | 1))
            }
            for fi16 := 0; fi16 < 12; fi16++ {
                var v17 int64 = -9223372036854775808
                _ = fi16
                _ = v17
            }
        } else {
            for fi18 := 0; fi18 < 2; fi18++ {
                p0 -= 42
                _ = fi18
            }
        }
    }
    var v19 int16 = -1
    _ = p0
    _ = v19
    _ = v7
    return -128
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 18446744073709551615
    v3 := make(map[string]bool)
    fmt.Println(v2)
    for fi20 := 0; fi20 < 11; fi20++ {
        if true {
            v21 := GenS1{F0: 0}
            if (1 >= 508) {
                fmt.Println(fi20)
                v2 = v2
            } else {
                var v22 float32 = (-2.25 / 0.1)
                for fi23 := 0; fi23 < 9; fi23++ {
                    _ = len(v3)
                    if (true && (!((("\n\t" + "hello") == "hello") || (v2 <= v2)))) {
                        v22 = (1.5 + v22)
                        fmt.Println(fi20)
                        if false {
                            v2 = (v2 & (v2 << ((65535) & 63)))
                        }
                        switch ((("a" + "\n\t") + "hello") + "a") {
                        case "":
                            var v24 float32 = float32(3.14159)
                            _ = v24
                        case "hello":
                            if (genFn0((v2 ^ 564)) >= 1) {
                                v25 := v3["hello"]
                                _ = v25
                            }
                        }
                        v22 += 3.14159
                    } else {
                        v26 := []int{(fi23 ^ -2147483648), 1, fi20, (fi23 & fi20)}
                        switch ((("" + "\n\t") + "a") + "foo bar") {
                        case "a":
                            v2 = 4294967295
                        case "foo bar":
                            v22 = 0.0
                        case "hello":
                            var v27 uint16 = 65535
                            _ = v27
                        default:
                            _ = len(v26)
                        }
                        fmt.Println(fi23, v2, fi20)
                        _ = v26
                    }
                    if (true && false) {
                        v28 := GenS3{F0: ("" + "a"), F1: 1, F2: 65535}
                        _ = v28
                    }
                    _ = fi23
                }
                _ = v22
            }
            var v29 string = ""
            _ = v21
            _ = v29
        }
        _ = fi20
    }
    switch -1 {
    case 127:
        var v30 uint16 = 79
        _ = v30
    case 42:
        v2 += (^(v2))
    default:
        var v31 float64 = -2.25
        _ = v31
    }
    if (0 < 191) {
        v2 -= v2
    } else {
        var v32 uint64 = v2
        if (0.0 < 1e10) {
            v33 := []int{9223372036854775807, 9223372036854775807, 127, 9223372036854775807}
            v34 := v3["hello"]
            if len(v33) >= 2 {
                v33 = v33[1:len(v33)]
            }
            _ = v33
            _ = v34
        } else {
            v1["\n\t"] = ""
            v3[""] = false
        }
        v35 := GenS3{F0: "", F1: 65535, F2: 42}
        _ = v32
        _ = v35
    }
    var v36 uint16 = 42
    var v37 int32 = 285
    for fi38 := 0; fi38 < 12; fi38++ {
        var v39 bool = (true || (1 <= uint8(v2)))
        _ = fi38
        _ = v39
    }
    v40 := GenS4{F0: 9223372036854775807}
    v41 := GenS2{F0: (v2 - (v2 + v2)), F1: 1}
    var v42 bool = (4294967295 == 183)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v40
    _ = v41
    _ = v42
}
