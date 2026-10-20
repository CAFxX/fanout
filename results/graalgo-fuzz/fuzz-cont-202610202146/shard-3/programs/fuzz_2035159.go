package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int16; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 uint64 }

func main() {
    var v1 int32 = 42
    var v2 uint8 = 2
    v3 := []uint32{2147483648, 468}
    v4 := make(map[int]string)
    var v5 uint8 = (v2 & 128)
    var v6 int64 = 9223372036854775807
    v7 := GenS1{F0: v2, F1: ""}
    if true {
        var v8 uint8 = v2
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 uint32 = 52
            _ = fi10
            _ = v11
        }
        v12 := GenS1{F0: (v5 + (^(v2))), F1: "\n\t"}
        var v13 float64 = 3.14159
        _ = fi9
        _ = v12
        _ = v13
    }
    if (-76 < 127) {
        for fi14 := 0; fi14 < 20; fi14++ {
            v2 -= v2
            _ = fi14
        }
        var v15 uint8 = v2
        _ = v15
    } else {
        switch "a" {
        case "\n\t":
            var v16 uint8 = (v5 | v2)
            _ = v16
        case "foo bar":
            if false {
                switch 2147483647 {
                case 525:
                    var v17 uint16 = 1
                    _ = v17
                default:
                    switch "foo bar" {
                    case "a":
                        if (v1 < v1) {
                            fmt.Println(v6)
                            v18 := v4[-1]
                            _ = v18
                        } else {
                            for fi19 := 0; fi19 < 5; fi19++ {
                                var v20 float32 = -2.25
                                _ = fi19
                                _ = v20
                            }
                            v1 = v1
                        }
                    case "":
                        v2 -= v2
                    }
                }
                var v21 int32 = 1
                _ = v21
            } else {
                v6 = 9223372036854775807
                for fi22 := 0; fi22 < 14; fi22++ {
                    for fi23 := 0; fi23 < 19; fi23++ {
                        v24 := GenS2{F0: 0}
                        _ = fi23
                        _ = v24
                    }
                    v4[fi22] = ("a" + ("foo bar" + ("a" + "hello")))
                    for fi25 := 0; fi25 < 7; fi25++ {
                        v26 := GenS1{F0: v5, F1: "a"}
                        _ = fi25
                        _ = v26
                    }
                    _ = fi22
                }
                var v27 string = (("\n\t" + ("\n\t" + "")) + "foo bar")
                _ = v27
            }
        }
        var v28 uint8 = v5
        v29 := GenS3{F0: v1, F1: ""}
        _ = v28
        _ = v29
    }
    var v30 float64 = 1.5
    var v31 float64 = 0.0
    v32 := GenS1{F0: (v5 << ((4294967295) & 63)), F1: "\n\t"}
    var v33 float64 = (v31 + v31)
    var v34 int = 1
    for fi35 := 0; fi35 < 3; fi35++ {
        switch v34 {
        case 4294967295:
            v6 += v6
        case 1:
            v1 -= (v1 % (((v1 / ((v1) | 1))) | 1))
        case 127:
            if (!false) {
                fmt.Println(v33, v30, v30)
                switch 998 {
                case 128:
                    var v36 uint = 255
                    _ = v36
                }
                var v37 int8 = 1
                _ = v37
            }
        }
        for fi38 := 0; fi38 < 17; fi38++ {
            fmt.Println(v5, v31)
            v33 -= -2.25
            _ = fi38
        }
        if true {
            v39 := GenS4{F0: 0, F1: 100, F2: (v2 & v2)}
            var v40 string = "foo bar"
            if true {
                v6 = v6
            } else {
                var v41 int32 = (v1 + v1)
                var v42 uint64 = 723
                v43 := GenS5{F0: 212, F1: (v5 * v5)}
                _ = v41
                _ = v42
                _ = v43
            }
            _ = v39
            _ = v40
        }
        _ = fi35
    }
    var v44 float32 = 1e10
    var v45 int = (v34 & 42)
    v46 := GenS6{F0: 4294967295}
    var v47 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v7
}
