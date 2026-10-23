package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 bool; F1 uint8; F2 int8 }

func main() {
    var v1 bool = true
    var v2 uint64 = 1
    v3 := make(map[int]int)
    var v4 int8 = 1
    v5 := make(map[string]uint32)
    var v6 uint = 553
    fmt.Println(v1, v2)
    for fi7 := 0; fi7 < 12; fi7++ {
        switch "a" {
        case "foo bar":
            v8 := []int{-544, 0}
            _ = v8
        case "\n\t":
            for fi9 := 0; fi9 < 20; fi9++ {
                var v10 string = "hello"
                _ = fi9
                _ = v10
            }
        default:
            var v11 int8 = (v4 >> ((2) & 63))
            _ = v11
        }
        var v12 bool = (((^(fi7)) > 2147483647) && v1)
        _ = fi7
        _ = v12
    }
    v13 := GenS1{F0: 32767, F1: 4294967295}
    v2 = (v2 - (v2 + v2))
    v4 = 1
    v14 := GenS2{F0: 9223372036854775807}
    delete(v5, "foo bar")
    v15 := GenS2{F0: 945}
    switch 2147483647 {
    case -1:
        for fi16 := 0; fi16 < 7; fi16++ {
            v17 := GenS3{F0: 42, F1: 1}
            _ = fi16
            _ = v17
        }
    default:
        var v18 uint16 = 849
        _ = v18
    }
    _ = len(v5)
    v19 := []int{-690, 823, 42, 128}
    if false {
        v6 += 255
        if true {
            v1 = (1 != 65)
            _ = len(v19)
            if (!v1) {
                var v20 int16 = -1
                _ = v20
            } else {
                v4 += v4
                if (float32(((1e10 - 0.0) * 0.0)) == 0.0) {
                    for fi21 := 0; fi21 < 3; fi21++ {
                        v22 := GenS4{F0: ((("" + "foo bar") + "foo bar") > ("" + "\n\t")), F1: 42, F2: v4}
                        v6 += v6
                        v23 := GenS1{F0: 578, F1: 0}
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    switch "foo bar" {
                    case "a":
                        v4 -= 0
                    case "hello":
                        var v24 uint8 = 128
                        _ = v24
                    default:
                        var v25 float64 = (1e10 * 1e10)
                        _ = v25
                    }
                }
                for fi26 := 0; fi26 < 4; fi26++ {
                    if false {
                        for fi27 := 0; fi27 < 4; fi27++ {
                            if v1 {
                                v2 = 952
                                if false {
                                    var v28 bool = (fi26 <= 198)
                                    fmt.Println(v2)
                                    var v29 string = ""
                                    if v1 {
                                        if ((!(-405 != 1)) && (true || (v1 || v28))) {
                                            v30 := GenS4{F0: (0.0 > 3.14159), F1: 2, F2: v4}
                                            v4 = 100
                                            if false {
                                                var v31 string = v29
                                                if (0.0 < ((0.0 / 3.14159) * 3.14159)) {
                                                    if v1 {
                                                        var v32 float64 = 3.14159
                                                        if len(v19) >= 2 {
                                                            v19 = v19[1:len(v19)]
                                                        }
                                                        v29 = ((("hello" + "foo bar") + v29) + v29)
                                                        _ = v32
                                                    }
                                                    var v33 uint8 = 42
                                                    if len(v19) > 0 {
                                                        v19[0] = fi27
                                                    }
                                                    _ = v33
                                                }
                                                v6 += v6
                                                fmt.Println(v6)
                                                fmt.Println(v29)
                                                _ = v31
                                            }
                                            fmt.Println(v4, v6, v28)
                                            _ = v30
                                        } else {
                                            fmt.Println(fi26, v29)
                                            v29 = (v29 + "")
                                            v28 = false
                                        }
                                        v34 := GenS1{F0: 475, F1: 240}
                                        _ = v34
                                    }
                                    delete(v5, v29)
                                    v35 := v3[(-(fi27))]
                                    fmt.Println(v6)
                                    _ = v28
                                    _ = v29
                                    _ = v35
                                }
                                v2 = 18446744073709551615
                            } else {
                                v2 += v2
                                var v36 float32 = float32(1e10)
                                _ = v36
                            }
                            v1 = (287 >= (^(v2)))
                            switch ("" + "a") {
                            case "\n\t":
                                var v37 int8 = (v4 / ((int8(v2)) | 1))
                                _ = v37
                            }
                            _ = fi27
                        }
                        var v38 int = fi26
                        _ = v38
                    }
                    if len(v19) > 0 {
                        v19[0] = (-(fi26))
                    }
                    _ = fi26
                }
            }
        } else {
            v1 = true
            var v39 uint16 = 65535
            _ = v39
        }
        var v40 float32 = (1e10 * 0.0)
        _ = v40
    } else {
        if (true || (1.5 < float64((-2.25 / 1.0)))) {
            var v41 string = (("" + "foo bar") + "hello")
            _ = v41
        }
        v42 := v3[997]
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
