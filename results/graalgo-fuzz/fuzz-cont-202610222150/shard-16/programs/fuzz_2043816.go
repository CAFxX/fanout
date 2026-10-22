package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 int64; F2 uint8 }
func genFn0() int16 {
    v5 := GenS1{F0: 1000, F1: 255, F2: 2}
    switch (((("a" + "\n\t") + "foo bar") + "foo bar") + ("foo bar" + "a")) {
    case "hello":
        for fi6 := 0; fi6 < 17; fi6++ {
            fmt.Println(fi6)
            _ = fi6
        }
    default:
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := GenS1{F0: -967, F1: 0, F2: 92}
            _ = fi7
            _ = v8
        }
    }
    _ = v5
    return 42
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[string]string)
    var v3 uint16 = 255
    v4 := []uint32{2147483647}
    fmt.Println(v3, v1)
    var v9 uint = 255
    v1 += genFn0()
    switch ("" + "a") {
    case "":
        var v10 bool = false
        _ = v10
    }
    var v11 uint = (v9 / ((255) | 1))
    if true {
        var v12 bool = (!((v1 & 0) >= (v1 / ((-32768) | 1))))
        _ = v12
    }
    var v13 string = "hello"
    if (9223372036854775807 != -8) {
        if true {
            fmt.Println(v11, v1, v13)
            var v14 float32 = 0.0
            if len(v4) > 0 {
                v4[0] = 0
            }
            _ = v14
        }
        var v15 int32 = 2147483647
        if false {
            _ = len(v2)
            for fi16 := 0; fi16 < 13; fi16++ {
                for fi17 := 0; fi17 < 17; fi17++ {
                    switch fi17 {
                    case -392:
                        switch "foo bar" {
                        case "hello":
                            var v18 int64 = 254
                            _ = v18
                        case "foo bar":
                            var v19 float64 = 1.5
                            _ = v19
                        case "a":
                            v20 := GenS1{F0: -32768, F1: 4294967295, F2: 0}
                            if false {
                                var v21 float32 = ((-2.25 - float32(3.14159)) + 1.5)
                                _ = v21
                            }
                            var v22 int64 = -9223372036854775808
                            v13 = v13
                            _ = v20
                            _ = v22
                        }
                    case -764:
                        v11 -= 856
                    }
                    if false {
                        v23 := GenS1{F0: (v1 >> ((v9) & 63)), F1: (^(v9)), F2: 2}
                        _ = v23
                    }
                    _ = fi17
                }
                _ = fi16
            }
        } else {
            var v24 float64 = 1.5
            v1 -= v1
            fmt.Println(v9, v3)
            _ = v24
        }
        _ = v15
    }
    switch v13 {
    case "":
        fmt.Println(v3)
    case "a":
        if (true || (42 == 42)) {
            var v25 float32 = (3.14159 + 0.0)
            switch "hello" {
            case "a":
                for fi26 := 0; fi26 < 8; fi26++ {
                    v27 := GenS1{F0: 646, F1: 255, F2: 495}
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    for fi28 := 0; fi28 < 8; fi28++ {
                        var v29 int32 = -586
                        v30 := v2[v13]
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    _ = fi26
                    _ = v27
                }
            default:
                v31 := GenS1{F0: (v1 | (v1 & genFn0())), F1: v11, F2: 867}
                _ = v31
            }
            _ = v25
        } else {
            v32 := GenS1{F0: v1, F1: (v9 * 486), F2: 4294967295}
            if ((490 == 2147483647) && false) {
                v33 := GenS2{F0: "a"}
                var v34 uint32 = 561
                _ = v33
                _ = v34
            }
            _ = v32
        }
    case "\n\t":
        v3 += (v3 | (v3 >> (((v9 % (((v11 + (v9 ^ 0))) | 1))) & 63)))
    }
    v1 -= v1
    if (!((255 != 255) || true)) {
        if true {
            fmt.Println(v11, v9, v13)
            v35 := v2[(v13 + "\n\t")]
            var v36 string = v35
            _ = v35
            _ = v36
        } else {
            v4 = append(v4, 2147483648)
            var v37 string = "a"
            _ = v37
        }
        var v38 uint64 = 42
        fmt.Println(v1, v38)
        _ = v38
    } else {
        v39 := GenS3{F0: true}
        v40 := GenS2{F0: "hello"}
        _ = v39
        _ = v40
    }
    fmt.Println(v13, v11, v13)
    v9 = v9
    v41 := GenS4{F0: 124, F1: -128, F2: 248}
    if (1e10 == 0.1) {
        var v42 float64 = float64(3.14159)
        v42 -= v42
        _ = v42
    } else {
        fmt.Println(v3)
        if true {
            switch 127 {
            case -496:
                var v43 int = -745
                _ = v43
            default:
                for fi44 := 0; fi44 < 7; fi44++ {
                    var v45 bool = false
                    _ = fi44
                    _ = v45
                }
            }
        } else {
            var v46 int8 = 42
            _ = v46
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v9
}
