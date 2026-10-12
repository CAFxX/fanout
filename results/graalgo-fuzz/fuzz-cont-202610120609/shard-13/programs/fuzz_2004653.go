package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int64; F1 string; F2 int16 }
type GenS3 struct { F0 int8; F1 int64 }
func genFn0(p0 bool, p1 uint16, p2 int16) int32 {
    var v7 int = -9223372036854775808
    v8 := GenS1{F0: 32767, F1: -9223372036854775808}
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 uint32 = 65535
        var v11 int64 = 9223372036854775807
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        p2 = -1
        if p0 {
            v13 := []int{(v7 % ((v7) | 1)), fi12, 32768}
            _ = v13
        }
        var v14 uint32 = 2147483647
        _ = fi12
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -2147483648
}

func main() {
    var v1 float64 = -2.25
    var v2 float64 = v1
    var v3 float64 = 1e10
    var v4 uint8 = 234
    var v5 float64 = float64(1e10)
    v6 := make(map[string]string)
    if false {
        if false {
            v1 = 3.14159
            v15 := GenS1{F0: 1, F1: 552}
            v1 -= (v5 / 1.5)
            _ = v15
        }
    } else {
        for fi16 := 0; fi16 < 15; fi16++ {
            v5 += (v2 + v5)
            if (((v1 - v3) != (v1 + (float64((((-2.25 - -2.25) / 1.0) * (float32((v2 - 0.1)) + ((3.14159 - -2.25) + 0.1)))) * v1))) && (!(false || (-2.25 != 0.0)))) {
                var v17 int = -532
                _ = v17
            }
            v5 -= ((-2.25 * (v5 / v3)) + v3)
            _ = fi16
        }
        var v18 uint32 = 2
        _ = v18
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 float64 = float64(((1e10 + 1.5) + float32(v2)))
            _ = fi20
            _ = v21
        }
        v3 += -2.25
        if ("a" == "\n\t") {
            v22 := v6[("foo bar" + "foo bar")]
            _ = v22
        }
        _ = fi19
    }
    v23 := []int{9223372036854775807}
    var v24 uint8 = (^(v4))
    for fi25 := 0; fi25 < 3; fi25++ {
        for fi26 := 0; fi26 < 6; fi26++ {
            v24 = v4
            fmt.Println(v5, fi26, v4)
            if (!(830 > 339)) {
                if (false || (863 < 255)) {
                    v4 = 128
                    v27 := GenS2{F0: -903, F1: "", F2: -1}
                    var v28 float32 = (-2.25 / 1.5)
                    _ = v27
                    _ = v28
                } else {
                    var v29 uint8 = (v4 % (((v24 - v24)) | 1))
                    if true {
                        v2 = v3
                        if true {
                            if (65535 != 36) {
                                v6["a"] = ("" + "")
                            }
                            v2 = ((1.5 / 3.14159) - v2)
                        } else {
                            v6["foo bar"] = "\n\t"
                        }
                        v3 = 0.1
                    }
                    _ = v29
                }
                if false {
                    v30 := GenS3{F0: -124, F1: 42}
                    delete(v6, "\n\t")
                    v4 = (v4 * v24)
                    _ = v30
                } else {
                    v1 += 1e10
                    v24 = 255
                }
                fmt.Println(fi25)
                var v31 int32 = 2147483647
                switch ((("\n\t" + "\n\t") + "") + "a") {
                case "a":
                    if len(v23) >= 2 {
                        v23 = v23[1:len(v23)]
                    }
                case "hello":
                    v2 -= -2.25
                case "\n\t":
                    switch (("foo bar" + ("hello" + "\n\t")) + "") {
                    case "\n\t":
                        switch 0 {
                        case 32768:
                            v5 = v1
                        case -2147483648:
                            var v32 int32 = v31
                            _ = v32
                        }
                    case "a":
                        var v33 string = "a"
                        _ = v33
                    }
                }
                _ = v31
            }
            _ = fi26
        }
        var v34 int8 = -1
        _ = fi25
        _ = v34
    }
    var v35 uint32 = 65535
    var v36 int32 = 0
    switch "hello" {
    case "foo bar":
        v3 = v5
    default:
        var v37 int16 = 42
        _ = v37
    }
    for fi38 := 0; fi38 < 20; fi38++ {
        v24 = v24
        v36 = 2147483647
        _ = fi38
    }
    var v39 int64 = -9223372036854775808
    var v40 uint16 = 1
    fmt.Println(v5)
    var v41 string = "foo bar"
    if true {
        var v42 int8 = 10
        var v43 float64 = v3
        _ = v42
        _ = v43
    }
    fmt.Println(v35, v5)
    if len(v23) > 0 {
        v23[0] = 7
    }
    if ((true || true) || false) {
        for fi44 := 0; fi44 < 1; fi44++ {
            fmt.Println(v40)
            _ = fi44
        }
        delete(v6, "foo bar")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
