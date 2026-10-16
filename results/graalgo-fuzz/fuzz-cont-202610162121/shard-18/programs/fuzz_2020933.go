package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 string }
type GenS2 struct { F0 uint; F1 uint; F2 uint64 }
type GenS3 struct { F0 uint }
func genFn0(p0 uint16, p1 int) bool {
    p1 = (p1 >> ((238) & 63))
    fmt.Println(p0, p1)
    var v5 int64 = 42
    v6 := []int{p1}
    var v7 int = p1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return false
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]string)
    var v3 int = 32767
    v4 := make(map[uint32]bool)
    var v8 float32 = 3.14159
    v9 := []int{v3}
    if true {
        for fi10 := 0; fi10 < 10; fi10++ {
            switch fi10 {
            case 32768:
                var v11 uint16 = 42
                _ = v11
            default:
                if (((1 > 1) && (1 == 0)) || (!(-348 >= 42))) {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    switch (fi10 & fi10) {
                    case -2147483648:
                        if ((v8 * (v8 - (3.14159 * 3.14159))) > float32((0.0 - 1e10))) {
                            var v12 int = v3
                            if (!true) {
                                var v13 uint8 = 225
                                var v14 int16 = 67
                                if len(v9) > 0 {
                                    v9[0] = (v3 >> ((2) & 63))
                                }
                                _ = v13
                                _ = v14
                            } else {
                                if true {
                                    var v15 int = v3
                                    for fi16 := 0; fi16 < 20; fi16++ {
                                        v15 += fi16
                                        _ = fi16
                                    }
                                    v3 = (v3 - v3)
                                    _ = v15
                                } else {
                                    v12 = 42
                                    v17 := GenS1{F0: (42 >= 32767), F1: 0, F2: "hello"}
                                    if true {
                                        if ((1e10 == (float64((0.0 + 1e10)) * ((0.0 - 0.0) - -2.25))) || ((((128 > 255) || false) && false) || (39 < -128))) {
                                            var v18 int8 = 96
                                            v3 -= fi10
                                            _ = v18
                                        } else {
                                            if false {
                                                var v19 int = (fi10 >> ((65535) & 63))
                                                v19 -= -9223372036854775808
                                                _ = v19
                                            }
                                            v3 += -374
                                        }
                                    }
                                    _ = v17
                                }
                            }
                            var v20 int = -769
                            _ = v12
                            _ = v20
                        } else {
                            v3 -= 265
                        }
                    }
                } else {
                    var v21 int64 = 42
                    _ = v21
                }
            }
            v22 := GenS2{F0: 620, F1: 65535, F2: 0}
            if true {
                v3 = v3
                var v23 int8 = 42
                _ = v23
            }
            _ = fi10
            _ = v22
        }
    } else {
        if false {
            v24 := v4[2147483647]
            switch "hello" {
            case "hello":
                fmt.Println(v24, v8)
            case "foo bar":
                var v25 bool = v24
                _ = v25
            default:
                var v26 string = "hello"
                _ = v26
            }
            switch v3 {
            case 127:
                v3 += v3
            case -431:
                v8 -= (1e10 / v8)
            case -465:
                for fi27 := 0; fi27 < 12; fi27++ {
                    v3 -= v3
                    _ = fi27
                }
            }
            _ = v24
        }
        v8 += v8
    }
    v3 = 418
    v28 := GenS3{F0: 481}
    v8 = 1.5
    for fi29 := 0; fi29 < 16; fi29++ {
        fmt.Println(v8)
        _ = fi29
    }
    if (42 == 2147483647) {
        var v30 int = (v3 | (v3 << ((18446744073709551615) & 63)))
        _ = v30
    }
    switch "" {
    case "\n\t":
        v8 = 1e10
    }
    switch v3 {
    case -771:
        delete(v4, 799)
    case -2147483648:
        switch (("hello" + "a") + ("hello" + "a")) {
        case "":
            var v31 int = 32767
            _ = v31
        case "\n\t":
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
    default:
        if false {
            var v32 int32 = 2147483647
            var v33 int32 = -1
            _ = v32
            _ = v33
        }
    }
    fmt.Println(v8, v3)
    var v34 int16 = -1
    if true {
        fmt.Println(v3, v34, v8)
        v35 := v2[""]
        _ = len(v4)
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v34
    _ = v4
    _ = v8
    _ = v9
}
