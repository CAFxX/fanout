package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint }
func genFn0(p0 int8, p1 uint16, p2 uint64) uint32 {
    v6 := []int{3, -404, 0}
    p0 = p0
    _ = len(v6)
    fmt.Println(p2, p0)
    var v7 float64 = ((0.0 * -2.25) * (-2.25 + 0.1))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 2147483647
}
func genFn1() int32 {
    for fi8 := 0; fi8 < 16; fi8++ {
        fmt.Println(fi8)
        _ = fi8
    }
    var v9 uint16 = 32767
    v10 := GenS1{F0: (-1 < 954), F1: 930}
    v9 += 0
    _ = v10
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = 0.1
    var v3 uint = 255
    var v4 int = 32768
    var v5 int = (-(v4))
    var v11 int = 0
    var v12 int8 = -128
    if false {
        if (!(0.1 >= 0.1)) {
            var v13 bool = true
            switch v4 {
            case 128:
                if v13 {
                    if false {
                        fmt.Println(v13, v12)
                        var v14 int32 = genFn1()
                        delete(v1, "hello")
                        _ = v14
                    }
                    v1[""] = (v4 + (v4 >> ((v3) & 63)))
                    for fi15 := 0; fi15 < 18; fi15++ {
                        if (v13 || (true || (537 != 0))) {
                            if (v5 != 1) {
                                v16 := GenS2{F0: (!v13)}
                                var v17 int8 = v12
                                var v18 string = "hello"
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                        } else {
                            v19 := GenS2{F0: v13}
                            switch "a" {
                            case "":
                                v20 := GenS3{F0: 100, F1: v3}
                                _ = v20
                            }
                            _ = v19
                        }
                        v11 += (v4 & v5)
                        _ = fi15
                    }
                }
            }
            _ = v13
        }
        var v21 uint32 = 120
        _ = v21
    }
    v22 := []int{(v4 / (((v5 ^ v11)) | 1)), (v11 & v11), -9223372036854775808, (-(v5))}
    v3 -= v3
    for fi23 := 0; fi23 < 3; fi23++ {
        v11 = v11
        _ = fi23
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        for fi25 := 0; fi25 < 8; fi25++ {
            v12 += v12
            switch (v4 % ((v4) | 1)) {
            case -9223372036854775808:
                fmt.Println(fi24)
            case -32:
                fmt.Println(v12)
            case 128:
                v2 += v2
            default:
                var v26 uint = 0
                _ = v26
            }
            _ = fi25
        }
        v22 = append(v22, v5)
        var v27 int64 = 9223372036854775807
        _ = fi24
        _ = v27
    }
    var v28 int16 = -32768
    var v29 uint32 = 453
    switch "hello" {
    case "\n\t":
        for fi30 := 0; fi30 < 6; fi30++ {
            var v31 uint8 = 212
            switch ((("" + "\n\t") + "foo bar") + ("foo bar" + "\n\t")) {
            case "a":
                v12 = (v12 + 0)
            default:
                v11 -= v5
            }
            _ = fi30
            _ = v31
        }
    }
    if false {
        fmt.Println(v5, v12)
    } else {
        var v32 float64 = 1e10
        _ = v32
    }
    fmt.Println(v28, v4)
    var v33 uint16 = 1
    if false {
        v3 = v3
        if true {
            var v34 int16 = v28
            v35 := GenS2{F0: true}
            _ = v34
            _ = v35
        }
        var v36 float32 = 1e10
        _ = v36
    }
    v4 -= v11
    var v37 uint8 = 127
    for fi38 := 0; fi38 < 11; fi38++ {
        fmt.Println(v33, v4)
        for fi39 := 0; fi39 < 3; fi39++ {
            var v40 uint64 = 484
            switch "" {
            case "a":
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
            case "\n\t":
                fmt.Println(fi38)
            default:
                switch -9223372036854775808 {
                case -279:
                    fmt.Println(v11, v28)
                default:
                    if (true || (!(uint8(v33) != v37))) {
                        var v41 bool = false
                        var v42 float64 = 3.14159
                        switch ("" + "\n\t") {
                        case "":
                            v43 := GenS3{F0: v12, F1: 241}
                            _ = v43
                        case "foo bar":
                            var v44 string = "\n\t"
                            _ = v44
                        }
                        _ = v41
                        _ = v42
                    }
                }
            }
            _ = fi39
            _ = v40
        }
        _ = fi38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v37
    _ = v4
    _ = v5
}
