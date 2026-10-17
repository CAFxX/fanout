package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 string }
func genFn0() int32 {
    var v7 uint = 42
    v7 = v7
    var v8 int32 = -152
    v7 -= (^(v7))
    _ = v7
    _ = v8
    return -2147483648
}
func genFn1(p0 int64, p1 uint32) bool {
    var v9 bool = (!false)
    p0 = (p0 * p0)
    _ = p0
    _ = p1
    _ = v9
    return v9
}

func main() {
    var v1 string = (((("foo bar" + "\n\t") + ("\n\t" + "a")) + "") + "a")
    v2 := []int{-755, 42, 425}
    var v3 float32 = float32((0.0 * 3.14159))
    var v4 uint32 = 2147483648
    var v5 uint32 = (v4 - 2)
    v6 := make(map[string]string)
    v10 := GenS1{F0: false, F1: 2, F2: 0}
    var v11 uint32 = v4
    v4 -= (v4 & (v4 - (v4 ^ (v11 << ((65535) & 63)))))
    for fi12 := 0; fi12 < 2; fi12++ {
        v6[v1] = v1
        _ = fi12
    }
    v13 := GenS1{F0: true, F1: 128, F2: 42}
    var v14 bool = true
    v6[v1] = ("hello" + (("a" + v1) + "\n\t"))
    v14 = (42 >= 1)
    for fi15 := 0; fi15 < 11; fi15++ {
        v4 = v4
        if v14 {
            if len(v2) > 0 {
                v2[0] = fi15
            }
        }
        _ = fi15
    }
    var v16 uint32 = 4294967295
    for fi17 := 0; fi17 < 9; fi17++ {
        fmt.Println(fi17, fi17)
        if false {
            v2 = append(v2, (fi17 & (fi17 | -2147483648)))
            for fi18 := 0; fi18 < 17; fi18++ {
                var v19 uint32 = (v11 % ((v5) | 1))
                v4 = v11
                switch (fi17 >> ((18446744073709551615) & 63)) {
                case 1:
                    v20 := GenS2{F0: (fi18 + 336)}
                    _ = v20
                case -375:
                    if (494 >= 0) {
                        _ = len(v2)
                        switch v1 {
                        case "":
                            fmt.Println(v14, v3)
                        case "hello":
                            delete(v6, ("hello" + v1))
                        case "foo bar":
                            _ = len(v2)
                        default:
                            var v21 int8 = 110
                            _ = v21
                        }
                    } else {
                        v5 += v16
                        fmt.Println(v5, v11)
                        v14 = ((v1 + v1) >= (v1 + (("hello" + v1) + v1)))
                    }
                case 32768:
                    v22 := GenS2{F0: (fi17 + fi18)}
                    var v23 bool = (v5 >= v16)
                    if true {
                        var v24 uint8 = 128
                        switch fi18 {
                        case -77:
                            var v25 uint32 = (v5 & v19)
                            _ = v25
                        default:
                            if ((v24 & v24) > v24) {
                                var v26 uint32 = (^(v5))
                                var v27 float32 = 3.14159
                                _ = len(v2)
                                _ = v26
                                _ = v27
                            }
                        }
                        if false {
                            v28 := GenS1{F0: v14, F1: (v24 >> ((461) & 63)), F2: 255}
                            var v29 uint64 = 1
                            _ = v28
                            _ = v29
                        }
                        var v30 int8 = 1
                        _ = v24
                        _ = v30
                    } else {
                        var v31 float32 = (3.14159 * (v3 + (v3 / 1e10)))
                        _ = v31
                    }
                    _ = v22
                    _ = v23
                default:
                    switch v1 {
                    case "\n\t":
                        switch v1 {
                        case "":
                            v3 = 3.14159
                        }
                    case "":
                        v16 -= v16
                    case "foo bar":
                        v32 := GenS3{F0: 0, F1: ("a" + v1)}
                        _ = v32
                    }
                    v2 = append(v2, 9223372036854775807)
                    var v33 uint16 = 208
                    _ = v33
                }
                v34 := GenS1{F0: false, F1: 0, F2: 42}
                _ = fi18
                _ = v19
                _ = v34
            }
        }
        for fi35 := 0; fi35 < 9; fi35++ {
            if v14 {
                v2 = append(v2, (^(fi35)))
            }
            _ = fi35
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
