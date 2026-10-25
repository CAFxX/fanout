package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 int8; F1 int; F2 uint16 }
type GenS4 struct { F0 uint8; F1 uint32; F2 int32 }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{596, 0}
    var v3 string = "hello"
    var v4 int64 = -9223372036854775808
    var v5 float64 = 1.5
    v1[(v3 + "")] = false
    switch v3 {
    case "\n\t":
        var v6 uint16 = 2
        _ = v6
    case "foo bar":
        switch v3 {
        case "":
            v4 = (-(v4))
        default:
            v4 += v4
        }
    }
    if true {
        for fi7 := 0; fi7 < 17; fi7++ {
            delete(v1, v3)
            _ = fi7
        }
    } else {
        v8 := GenS1{F0: 225, F1: -1}
        _ = v8
    }
    if true {
        var v9 float32 = -2.25
        var v10 uint8 = 127
        if (42 < -66) {
            fmt.Println(v5, v3, v4)
            if true {
                v9 = 1.5
                var v11 int32 = -284
                v12 := GenS1{F0: 32767, F1: 0}
                _ = v11
                _ = v12
            } else {
                switch (v3 + v3) {
                case "\n\t":
                    v13 := GenS2{F0: v3, F1: 537}
                    _ = v13
                default:
                    v14 := GenS3{F0: 100, F1: -1, F2: 214}
                    _ = v14
                }
            }
            for fi15 := 0; fi15 < 17; fi15++ {
                if (!false) {
                    var v16 int8 = 74
                    v9 -= 1e10
                    _ = v16
                }
                _ = len(v2)
                _ = fi15
            }
        } else {
            for fi17 := 0; fi17 < 16; fi17++ {
                v2 = append(v2, (fi17 * 0))
                if ((v4 % ((v4) | 1)) >= (-(v4))) {
                    var v18 int16 = 1000
                    v19 := v1["a"]
                    _ = v18
                    _ = v19
                }
                _ = fi17
            }
        }
        _ = v10
        _ = v9
    } else {
        v20 := GenS1{F0: -636, F1: -1}
        _ = v20
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        v3 = v3
        v22 := GenS4{F0: 255, F1: 1, F2: -1}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi21
        _ = v22
    }
    fmt.Println(v3, v3, v3)
    fmt.Println(v3, v3)
    var v23 uint64 = 0
    switch v3 {
    case "\n\t":
        for fi24 := 0; fi24 < 15; fi24++ {
            if false {
                v3 = "hello"
            }
            var v25 float64 = v5
            if false {
                delete(v1, ("\n\t" + (v3 + v3)))
                v4 = v4
                for fi26 := 0; fi26 < 8; fi26++ {
                    switch v3 {
                    case "a":
                        if false {
                            var v27 int = 128
                            _ = v27
                        } else {
                            v3 = "foo bar"
                        }
                    default:
                        var v28 float64 = float64((0.1 + 0.0))
                        _ = v28
                    }
                    _ = fi26
                }
            }
            _ = fi24
            _ = v25
        }
    case "a":
        _ = len(v1)
    }
    var v29 int16 = 1
    fmt.Println(v23, v5)
    v5 = -2.25
    v5 = 3.14159
    if (!true) {
        v4 = v4
    } else {
        v30 := GenS2{F0: (v3 + ("" + "")), F1: 65535}
        fmt.Println(v23, v29)
        for fi31 := 0; fi31 < 5; fi31++ {
            switch -776 {
            case 32768:
                v32 := GenS4{F0: 0, F1: uint32(v23), F2: -2147483648}
                _ = v32
            case 4294967295:
                var v33 uint64 = 2
                _ = v33
            case -9223372036854775808:
                var v34 float32 = 0.0
                _ = v34
            }
            _ = fi31
        }
        _ = v30
    }
    var v35 uint8 = 127
    var v36 int16 = (v29 + (v29 / (((v29 << ((1) & 63))) | 1)))
    for fi37 := 0; fi37 < 13; fi37++ {
        fmt.Println(v4, v23, v3)
        v38 := GenS3{F0: -1, F1: 42, F2: 1}
        var v39 uint64 = 18446744073709551615
        _ = fi37
        _ = v38
        _ = v39
    }
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
