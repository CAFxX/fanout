package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint }
type GenS3 struct { F0 int32; F1 int64 }
func genFn0(p0 uint16, p1 int8) uint {
    v6 := []int{-9223372036854775808, int(p1), 42, -9223372036854775808}
    p0 = p0
    if len(v6) > 0 {
        v6[0] = 1
    }
    _ = p0
    _ = p1
    _ = v6
    return 65535
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 42
    v3 := make(map[string]bool)
    var v4 float64 = 0.0
    v5 := []uint32{0}
    switch "a" {
    case "hello":
        fmt.Println(v4, v2)
    case "a":
        v4 = v4
    }
    if false {
        v2 = v2
    }
    if (1 <= 0) {
        var v7 uint16 = 1
        _ = v7
    } else {
        switch -2147483648 {
        case 640:
            if (4294967295 == 0) {
                var v8 int8 = -100
                _ = v8
            }
        }
        var v9 uint32 = 0
        if (v4 == v4) {
            v10 := GenS1{F0: 127, F1: 42}
            v11 := GenS2{F0: 1, F1: 18, F2: 809}
            if (0.1 < (-2.25 / 1e10)) {
                for fi12 := 0; fi12 < 19; fi12++ {
                    if true {
                        for fi13 := 0; fi13 < 9; fi13++ {
                            v5 = append(v5, v9)
                            fmt.Println(v9)
                            _ = fi13
                        }
                    }
                    switch (fi12 * fi12) {
                    case -725:
                        fmt.Println(fi12, v9, v4)
                    }
                    _ = fi12
                }
                var v14 int8 = 100
                _ = v14
            } else {
                v2 = v2
                for fi15 := 0; fi15 < 12; fi15++ {
                    var v16 float32 = (1.5 * 1e10)
                    switch fi15 {
                    case 42:
                        for fi17 := 0; fi17 < 20; fi17++ {
                            delete(v1, "foo bar")
                            switch (("foo bar" + ("foo bar" + "\n\t")) + "\n\t") {
                            case "":
                                var v18 float32 = (((1e10 + (0.0 / v16)) * v16) / float32(v4))
                                _ = v18
                            case "\n\t":
                                v9 = (v9 ^ (v9 >> ((4294967295) & 63)))
                            case "a":
                                var v19 float64 = (v4 - (1e10 * -2.25))
                                _ = v19
                            }
                            var v20 int8 = 0
                            v16 += v16
                            _ = fi17
                            _ = v20
                        }
                    case -2147483648:
                        fmt.Println(v4, v4)
                    case 2147483647:
                        var v21 float32 = 1e10
                        _ = v21
                    }
                    _ = fi15
                    _ = v16
                }
            }
            _ = v10
            _ = v11
        }
        _ = v9
    }
    v4 += (1.5 * (v4 + (v4 / v4)))
    _ = len(v3)
    var v22 uint16 = 32767
    var v23 int = 997
    if true {
        _ = len(v3)
    } else {
        switch v23 {
        case 32768:
            v1[""] = "foo bar"
        default:
            if true {
                v4 = v4
                var v24 float64 = float64(3.14159)
                if false {
                    if false {
                        switch (-(v23)) {
                        case 32768:
                            if len(v5) > 0 {
                                v5[0] = 2147483648
                            }
                        case 408:
                            delete(v3, "\n\t")
                        }
                        if len(v5) > 0 {
                            v5[0] = 2147483647
                        }
                    } else {
                        if ("\n\t" < (((("a" + "\n\t") + "a") + ("a" + "\n\t")) + "hello")) {
                            fmt.Println(v23, v2, v4)
                            if (!(false && true)) {
                                var v25 uint64 = 4294967295
                                var v26 int64 = 1
                                for fi27 := 0; fi27 < 12; fi27++ {
                                    fmt.Println(v2, v22)
                                    _ = fi27
                                }
                                _ = v25
                                _ = v26
                            } else {
                                if (1 <= 611) {
                                    var v28 int16 = 42
                                    if false {
                                        v29 := GenS3{F0: 1, F1: -9223372036854775808}
                                        _ = v29
                                    }
                                    _ = v28
                                }
                                v30 := GenS1{F0: v23, F1: -2147483648}
                                _ = v30
                            }
                        } else {
                            v2 -= v2
                            var v31 string = "a"
                            _ = v31
                        }
                    }
                } else {
                    var v32 uint8 = v2
                    _ = v32
                }
                _ = v24
            }
        }
    }
    v33 := v3["a"]
    v2 = v2
    var v34 bool = (false || v33)
    v35 := v1[""]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
