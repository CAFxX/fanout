package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint }
type GenS4 struct { F0 int16; F1 int8; F2 bool }
type GenS5 struct { F0 uint8; F1 int }
func genFn0() string {
    var v4 uint16 = 441
    if (18446744073709551615 > 357) {
        for fi5 := 0; fi5 < 12; fi5++ {
            var v6 string = (((("" + "\n\t") + ("hello" + "")) + "a") + (("\n\t" + (("\n\t" + ("a" + (("hello" + "\n\t") + ((("a" + ("\n\t" + "")) + "foo bar") + "a")))) + "a")) + ""))
            var v7 int32 = 396
            if false {
                var v8 float32 = 0.1
                _ = v8
            } else {
                v9 := []int{(^(fi5))}
                _ = v9
            }
            _ = fi5
            _ = v6
            _ = v7
        }
        v4 = v4
        var v10 int = 774
        _ = v10
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        if (false && (true && (42 < -1))) {
            var v12 float32 = 3.14159
            switch (fi11 - (fi11 & (^(fi11)))) {
            case -2147483648:
                for fi13 := 0; fi13 < 15; fi13++ {
                    v12 = (((1e10 / 1.5) - (0.0 + 0.0)) / v12)
                    v14 := GenS1{F0: 0, F1: 533}
                    _ = fi13
                    _ = v14
                }
            case 32768:
                switch 127 {
                case 977:
                    var v15 float64 = 1.5
                    _ = v15
                case -55:
                    var v16 uint32 = 42
                    _ = v16
                }
            }
            for fi17 := 0; fi17 < 20; fi17++ {
                if true {
                    if false {
                        var v18 int32 = -2147483648
                        v12 -= (float32((((1.5 * (3.14159 - 1e10)) + (1e10 / 1e10)) - 1.5)) * (1.5 * (v12 + v12)))
                        v19 := []int{int(v4)}
                        if false {
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                            fmt.Println(fi17)
                        }
                        var v20 uint8 = 2
                        switch ("foo bar" + "hello") {
                        case "foo bar":
                            v20 = (v20 - v20)
                        case "":
                            v21 := GenS1{F0: fi17, F1: int(v20)}
                            _ = v21
                        case "a":
                            switch "\n\t" {
                            case "hello":
                                v12 = (((v12 / v12) + -2.25) * (v12 / v12))
                            case "\n\t":
                                v12 = -2.25
                            default:
                                var v22 int8 = 34
                                _ = v22
                            }
                        }
                        _ = v18
                        _ = v19
                        _ = v20
                    } else {
                        v12 += 0.1
                        if (269 >= 2) {
                            v23 := GenS2{F0: 42}
                            var v24 int32 = 42
                            v25 := GenS2{F0: 18446744073709551615}
                            v12 += 0.1
                            if (((v4 - (v4 | v4)) > v4) || true) {
                                v24 = v24
                                v26 := GenS3{F0: 2, F1: (^(v4)), F2: 243}
                                _ = v26
                            } else {
                                v24 = 0
                                v27 := GenS3{F0: 42, F1: 2, F2: 354}
                                var v28 int = (-(fi11))
                                _ = v27
                                _ = v28
                            }
                            fmt.Println(fi17, fi11, v24)
                            v29 := []int{fi17}
                            _ = v23
                            _ = v24
                            _ = v25
                            _ = v29
                        }
                        v30 := GenS1{F0: -9223372036854775808, F1: (fi17 + 127)}
                        var v31 uint = 18446744073709551615
                        fmt.Println(v4, v12, fi17)
                        _ = v30
                        _ = v31
                    }
                    var v32 uint16 = v4
                    _ = v32
                } else {
                    fmt.Println(v4, fi17)
                }
                var v33 int = fi17
                var v34 uint32 = 2147483648
                _ = fi17
                _ = v33
                _ = v34
            }
            _ = v12
        }
        for fi35 := 0; fi35 < 14; fi35++ {
            v36 := GenS2{F0: 18446744073709551615}
            v37 := GenS4{F0: 32767, F1: 81, F2: (!((fi11 >> ((764) & 63)) != fi35))}
            _ = fi35
            _ = v36
            _ = v37
        }
        v38 := []int{fi11, fi11, (fi11 & 9223372036854775807)}
        _ = fi11
        _ = v38
    }
    _ = v4
    return ("\n\t" + "foo bar")
}

func main() {
    var v1 int8 = 42
    var v2 uint16 = 32767
    v3 := []uint32{}
    v39 := GenS1{F0: -828, F1: -2147483648}
    var v40 float64 = 1e10
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    fmt.Println(v1, v40, v40)
    v41 := GenS2{F0: 4294967295}
    var v42 string = "hello"
    v43 := GenS5{F0: 0, F1: 1}
    fmt.Println(v2, v40, v1)
    v40 += 0.1
    v2 = v2
    v42 = (("foo bar" + v42) + v42)
    var v44 uint8 = 128
    var v45 int8 = v1
    v42 = ((v42 + v42) + ((v42 + (v42 + v42)) + v42))
    var v46 int = 32767
    if len(v3) > 0 {
        v3[0] = 0
    }
    var v47 string = ("a" + v42)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
}
