package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint32; F1 int32 }

func main() {
    v1 := []int{}
    var v2 float64 = 3.14159
    var v3 int8 = 100
    if true {
        if false {
            var v4 float64 = (3.14159 + (v2 + (v2 + (v2 - float64(1e10)))))
            _ = v4
        }
    }
    v3 += v3
    _ = len(v1)
    switch "foo bar" {
    case "\n\t":
        for fi5 := 0; fi5 < 5; fi5++ {
            v2 -= 3.14159
            switch ("" + ("foo bar" + (("\n\t" + ("\n\t" + "foo bar")) + "\n\t"))) {
            case "\n\t":
                for fi6 := 0; fi6 < 8; fi6++ {
                    if (580 < 42) {
                        if false {
                            var v7 int16 = -32768
                            if ("hello" > "hello") {
                                if false {
                                    if (false && (422 < 468)) {
                                        var v8 float64 = (v2 - 0.1)
                                        v7 = -32768
                                        _ = len(v1)
                                        _ = v8
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = fi6
                                    }
                                }
                            } else {
                                var v9 int8 = v3
                                v10 := GenS1{F0: 2}
                                _ = v10
                                _ = v9
                            }
                            _ = v7
                        }
                        v3 += -83
                    }
                    _ = fi6
                }
            case "":
                if false {
                    var v11 int16 = -651
                    var v12 string = "a"
                    _ = v11
                    _ = v12
                } else {
                    switch -996 {
                    case 444:
                        v3 = 0
                    case 2147483647:
                        v13 := GenS2{F0: 1, F1: (("a" + "") + ("foo bar" + (("hello" + "a") + "hello")))}
                        _ = v13
                    case 663:
                        v2 += (((((((0.1 - float64(0.1)) * v2) * 0.0) + float64(0.0)) - 0.0) / (v2 * 1.5)) * v2)
                    default:
                        v14 := GenS1{F0: 779}
                        _ = v14
                    }
                }
            default:
                for fi15 := 0; fi15 < 17; fi15++ {
                    v16 := GenS2{F0: 42, F1: "hello"}
                    _ = fi15
                    _ = v16
                }
            }
            v2 -= 1.5
            _ = fi5
        }
    case "foo bar":
        var v17 uint64 = 1
        _ = v17
    case "":
        v3 += -128
    }
    fmt.Println(v2, v2)
    v18 := GenS3{F0: 699, F1: 2147483647}
    v19 := GenS3{F0: 1, F1: 207}
    fmt.Println(v3, v2)
    v1 = append(v1, 32768)
    for fi20 := 0; fi20 < 9; fi20++ {
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 uint32 = 2
            fmt.Println(fi21, v2)
            var v23 uint16 = 0
            _ = fi21
            _ = v22
            _ = v23
        }
        var v24 uint64 = 656
        _ = fi20
        _ = v24
    }
    switch -2147483648 {
    case 2147483647:
        v2 -= v2
    case 32768:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case -2147483648:
        v2 += float64((-2.25 / (3.14159 * -2.25)))
    }
    var v25 uint64 = 967
    for fi26 := 0; fi26 < 16; fi26++ {
        v27 := GenS3{F0: 65535, F1: -730}
        for fi28 := 0; fi28 < 19; fi28++ {
            v3 += 100
            switch "foo bar" {
            case "":
                if len(v1) > 0 {
                    v1[0] = (^(fi26))
                }
            case "\n\t":
                var v29 int64 = 713
                _ = v29
            default:
                if (!(0.0 < float32(v2))) {
                    v2 = float64(((1.5 / ((0.0 / float32(v2)) + 1e10)) - float32(float64(-2.25))))
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v25 += 18446744073709551615
                } else {
                    v30 := GenS3{F0: 0, F1: 460}
                    _ = len(v1)
                    _ = v30
                }
            }
            _ = len(v1)
            _ = fi28
        }
        var v31 string = ((((("a" + (("\n\t" + "\n\t") + "")) + ("a" + "foo bar")) + (("" + (("a" + (("hello" + "\n\t") + "foo bar")) + ("\n\t" + "hello"))) + "a")) + "") + "foo bar")
        _ = fi26
        _ = v27
        _ = v31
    }
    v2 += ((0.1 + v2) / (3.14159 / (float64(0.0) / (v2 + v2))))
    if len(v1) > 0 {
        v1[0] = -545
    }
    v32 := GenS2{F0: 0, F1: "foo bar"}
    v33 := GenS2{F0: 727, F1: "foo bar"}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v33
}
