package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 int8 }
func genFn0(p0 int8) string {
    v4 := []int{164, -68, -1}
    for fi5 := 0; fi5 < 1; fi5++ {
        switch "foo bar" {
        case "\n\t":
            if (true || (!false)) {
                for fi6 := 0; fi6 < 1; fi6++ {
                    if ((!((!(("" > ("hello" + ("hello" + ("foo bar" + ("\n\t" + ("\n\t" + "a")))))) && (!(!true)))) && (1 < 0))) && (true || (3.14159 <= 0.1))) {
                        var v7 uint = 10
                        if ((0.1 > (float32(3.14159) / (0.0 + 0.1))) || true) {
                            switch ("" + ("\n\t" + "")) {
                            case "hello":
                                if len(v4) > 0 {
                                    v4[0] = (fi6 | 9223372036854775807)
                                }
                            case "foo bar":
                                var v8 uint16 = 32767
                                v7 += (v7 & uint(v8))
                                _ = v8
                            default:
                                var v9 int32 = -1
                                _ = v9
                            }
                            v10 := GenS1{F0: 42}
                            p0 = (p0 & p0)
                            _ = v10
                        } else {
                            var v11 bool = true
                            _ = len(v4)
                            switch "\n\t" {
                            case "hello":
                                v4 = append(v4, fi6)
                            case "a":
                                switch fi5 {
                                case 32767:
                                    if true {
                                        var v12 uint64 = 0
                                        switch 32767 {
                                        case -1:
                                            v4 = append(v4, (-(fi6)))
                                        default:
                                            v7 = (^(v7))
                                        }
                                        _ = v12
                                    }
                                case -608:
                                    _ = len(v4)
                                case 128:
                                    fmt.Println(fi6, fi5)
                                }
                                v13 := GenS2{F0: (v7 % ((777) | 1))}
                                _ = v13
                            }
                            v11 = ((float32((-2.25 / 1.0)) != -2.25) || v11)
                            if v11 {
                                var v14 uint16 = 255
                                _ = v14
                            }
                            _ = v11
                        }
                        var v15 uint32 = 2
                        _ = v15
                        _ = v7
                    }
                    _ = fi6
                }
                var v16 int64 = 0
                switch "a" {
                case "hello":
                    var v17 int64 = v16
                    _ = v17
                case "foo bar":
                    v16 -= v16
                case "":
                    switch (fi5 ^ fi5) {
                    case 62:
                        fmt.Println(p0, fi5, v16)
                    }
                default:
                    var v18 uint32 = 65535
                    _ = v18
                }
                _ = v16
            }
        }
        _ = fi5
    }
    v19 := GenS2{F0: 42}
    for fi20 := 0; fi20 < 2; fi20++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi20
    }
    _ = p0
    _ = v19
    _ = v4
    return (("foo bar" + ("hello" + "a")) + "hello")
}
func genFn1(p0 int32, p1 int32, p2 int) int64 {
    var v21 string = "foo bar"
    v21 = v21
    for fi22 := 0; fi22 < 13; fi22++ {
        p1 = p1
        _ = fi22
    }
    var v23 uint32 = 2147483648
    p1 += 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v23
    return 1
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    var v3 bool = true
    var v24 uint32 = 187
    v3 = (((0.1 - 0.1) > (0.0 / 1.0)) && v3)
    if v3 {
        if (!v3) {
            v3 = v3
        }
    } else {
        fmt.Println(v24, v3)
    }
    var v25 int8 = 42
    var v26 int8 = (v25 | 0)
    v24 += 4294967295
    v1[(("\n\t" + ("a" + "hello")) + (("" + "foo bar") + (("hello" + ("a" + "a")) + "hello")))] = "a"
    var v27 bool = (-2.25 != 0.0)
    v3 = v27
    v28 := GenS2{F0: 922}
    if v3 {
        for fi29 := 0; fi29 < 16; fi29++ {
            if ((!(!true)) || v27) {
                var v30 uint64 = 227
                _ = v30
            }
            var v31 int32 = -2147483648
            _ = fi29
            _ = v31
        }
        v27 = false
    }
    v25 += -74
    v32 := []int{9223372036854775807}
    v33 := GenS3{F0: 1, F1: (v25 + (^(v25)))}
    switch -2147483648 {
    case 9223372036854775807:
        fmt.Println(v26, v26, v26)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v33
}
