package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 int64; F1 uint16; F2 int32 }
func genFn0() string {
    var v5 uint8 = 30
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 uint64 = 18446744073709551615
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    var v9 bool = true
    v9 = (v5 >= 0)
    fmt.Println(v5, v9)
    _ = v5
    _ = v9
    return (((("a" + ("a" + ("a" + ("" + "")))) + (("foo bar" + "\n\t") + "")) + "hello") + "\n\t")
}

func main() {
    var v1 int64 = 0
    v2 := []string{"foo bar", "hello", ("foo bar" + ""), "a"}
    var v3 string = "hello"
    v4 := make(map[string]bool)
    fmt.Println(v3, v3)
    if len(v2) > 0 {
        v2[0] = "a"
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := GenS1{F0: 32767}
        if false {
            v12 := GenS1{F0: -424}
            if false {
                v13 := GenS1{F0: 42}
                if false {
                    v14 := GenS1{F0: -210}
                    _ = v14
                }
                _ = v13
            } else {
                var v15 float32 = 0.0
                v15 = (((v15 + 0.1) - -2.25) + float32((((-2.25 / 3.14159) / float64((v15 - ((1.5 / (v15 - v15)) / v15)))) + (((float64(0.0) - 3.14159) + (-2.25 + (0.0 + 3.14159))) * (-2.25 + ((3.14159 - float64(v15)) - ((0.1 * (0.1 + 1e10)) - (1e10 * 0.0))))))))
                var v16 uint8 = 42
                _ = v15
                _ = v16
            }
            _ = v12
        } else {
            fmt.Println(v1)
            v17 := GenS1{F0: 32767}
            v18 := v4[v3]
            _ = v17
            _ = v18
        }
        v1 = v1
        _ = fi10
        _ = v11
    }
    if (0 == 127) {
        if (0 == 4) {
            var v19 uint8 = 42
            v2 = append(v2, v3)
            switch v3 {
            case "hello":
                if false {
                    if (3.14159 <= 0.1) {
                        v20 := GenS2{F0: (v1 / ((v1) | 1)), F1: 18446744073709551615, F2: 1}
                        fmt.Println(v3, v19)
                        for fi21 := 0; fi21 < 17; fi21++ {
                            if false {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v3 = "foo bar"
                                v22 := GenS1{F0: -32768}
                                _ = v22
                            } else {
                                v1 = 738
                            }
                            var v23 int8 = -75
                            v3 = ("foo bar" + (v3 + v3))
                            _ = fi21
                            _ = v23
                        }
                        _ = v20
                    }
                } else {
                    var v24 int64 = -1
                    for fi25 := 0; fi25 < 7; fi25++ {
                        for fi26 := 0; fi26 < 10; fi26++ {
                            switch (fi25 & 32768) {
                            case -304:
                                fmt.Println(v19, fi26, v24)
                            case 128:
                                v24 -= v1
                            default:
                                v27 := GenS2{F0: v1, F1: 1, F2: 0}
                                _ = v27
                            }
                            var v28 string = (v3 + "a")
                            _ = fi26
                            _ = v28
                        }
                        _ = fi25
                    }
                    var v29 uint64 = 42
                    _ = v24
                    _ = v29
                }
            case "a":
                if true {
                    _ = len(v2)
                }
            }
            _ = v19
        }
        v30 := GenS3{F0: v3, F1: 32767}
        _ = v30
    } else {
        delete(v4, v3)
    }
    var v31 uint8 = 196
    if len(v2) > 0 {
        v2[0] = v3
    }
    fmt.Println(v1, v31)
    var v32 float64 = 1e10
    var v33 int32 = 954
    v34 := GenS4{F0: int64(v31), F1: 0, F2: v33}
    v35 := GenS2{F0: 0, F1: 4294967295, F2: 448}
    fmt.Println(v3, v33)
    if ((!(0 != 0)) && false) {
        var v36 bool = (100 != 42)
        for fi37 := 0; fi37 < 17; fi37++ {
            fmt.Println(v31, fi37)
            _ = fi37
        }
        _ = v36
    } else {
        if (32767 == -1) {
            var v38 uint8 = v31
            var v39 int8 = 47
            switch (v3 + v3) {
            case "foo bar":
                var v40 uint64 = 2
                _ = v40
            case "\n\t":
                var v41 bool = ((-830 > -32768) && true)
                _ = v41
            case "hello":
                fmt.Println(v3)
            }
            _ = v38
            _ = v39
        }
        if len(v2) > 0 {
            v2[0] = v3
        }
        switch 752 {
        case 32768:
            fmt.Println(v1, v1)
        case 0:
            _ = len(v2)
        case -2147483648:
            v3 = ((v3 + "hello") + v3)
        default:
            var v42 uint32 = 0
            _ = v42
        }
    }
    fmt.Println(v3, v32, v32)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
