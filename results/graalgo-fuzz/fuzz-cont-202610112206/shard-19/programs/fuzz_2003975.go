package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int; F2 uint8 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 string; F1 uint64; F2 uint64 }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint8, p1 uint) int8 {
    p0 += (p0 * p0)
    v4 := GenS1{F0: 127}
    v5 := []int{32768, -9223372036854775808, 1}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -114
}

func main() {
    v1 := make(map[string]bool)
    v2 := []int{-254, 32768}
    var v3 uint64 = 18446744073709551615
    if true {
        var v6 int32 = 869
        v6 = -983
        var v7 int64 = 42
        _ = v6
        _ = v7
    } else {
        var v8 float64 = ((0.1 / (1.5 / (-2.25 - 0.0))) + (0.1 + ((0.0 / 1e10) / 0.1)))
        v3 = 2
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        fmt.Println(v3, fi9)
        if (!(!false)) {
            for fi10 := 0; fi10 < 5; fi10++ {
                switch "" {
                case "":
                    switch (^(fi9)) {
                    case 4294967295:
                        v3 -= (v3 + v3)
                    }
                    v3 -= v3
                    v2 = append(v2, fi9)
                    v1[(((("\n\t" + "\n\t") + "foo bar") + "foo bar") + "")] = false
                    var v11 int8 = -1
                    if false {
                        v12 := GenS2{F0: 156, F1: fi9, F2: 1}
                        _ = v12
                    }
                    v3 = v3
                    _ = v11
                case "a":
                    switch ("hello" + "hello") {
                    case "a":
                        var v13 float32 = (1e10 / 1e10)
                        _ = v13
                    case "\n\t":
                        if false {
                            fmt.Println(v3)
                            var v14 uint64 = v3
                            var v15 bool = false
                            v16 := GenS3{F0: 65535, F1: -32768}
                            _ = v14
                            _ = v15
                            _ = v16
                        } else {
                            fmt.Println(fi9, v3, v3)
                        }
                    case "hello":
                        switch ("hello" + "a") {
                        case "hello":
                            v17 := GenS2{F0: 127, F1: (fi10 - (fi9 % ((fi10) | 1))), F2: 0}
                            v2 = append(v2, (^(fi9)))
                            _ = v17
                        case "a":
                            switch "" {
                            case "\n\t":
                                var v18 int8 = -7
                                _ = v18
                            case "":
                                var v19 uint = 18446744073709551615
                                _ = v19
                            case "a":
                                var v20 int16 = int16(fi10)
                                _ = v20
                            }
                        default:
                            v21 := GenS4{F0: "hello", F1: (v3 | v3), F2: v3}
                            _ = v21
                        }
                    }
                default:
                    v3 = (^(v3))
                }
                _ = fi10
            }
            var v22 float32 = -2.25
            _ = v22
        } else {
            var v23 float64 = 3.14159
            _ = v23
        }
        _ = fi9
    }
    var v24 int8 = 1
    for fi25 := 0; fi25 < 19; fi25++ {
        _ = len(v1)
        for fi26 := 0; fi26 < 3; fi26++ {
            var v27 uint32 = 96
            if true {
                fmt.Println(v24, fi25, v3)
                if false {
                    if (-710 <= 0) {
                        if (0 >= -1) {
                            fmt.Println(fi26, v3)
                        } else {
                            var v28 int = (fi25 & fi25)
                            v24 -= (^(v24))
                            _ = v28
                        }
                    } else {
                        var v29 int32 = 42
                        var v30 float64 = (-2.25 + 3.14159)
                        _ = v29
                        _ = v30
                    }
                    fmt.Println(fi25, fi26)
                    var v31 int16 = 32767
                    _ = v31
                }
                var v32 uint = 4294967295
                _ = v32
            }
            _ = len(v1)
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    var v33 uint = 18446744073709551615
    v33 += v33
    var v34 uint32 = 1
    v35 := GenS5{F0: -9223372036854775808}
    v2 = append(v2, 128)
    var v36 uint32 = (v34 ^ (^(v34)))
    v34 -= 744
    var v37 bool = true
    if ((-2.25 != 3.14159) || true) {
        var v38 uint32 = 272
        if (!true) {
            fmt.Println(v34)
            v1["hello"] = true
        } else {
            v36 = 4294967295
        }
        _ = v38
    }
    v37 = (((128 < 42) && (!((!true) || ((("\n\t" + "hello") + ("\n\t" + "foo bar")) < "foo bar")))) || v37)
    if (v37 || (1000 == 42)) {
        var v39 string = "foo bar"
        _ = v39
    } else {
        for fi40 := 0; fi40 < 16; fi40++ {
            if true {
                var v41 int = fi40
                _ = v41
            } else {
                v3 = (v3 ^ v3)
            }
            _ = fi40
        }
        v24 -= v24
        var v42 uint16 = 42
        _ = v42
    }
    switch ("a" + "a") {
    case "\n\t":
        var v43 uint32 = v36
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}
