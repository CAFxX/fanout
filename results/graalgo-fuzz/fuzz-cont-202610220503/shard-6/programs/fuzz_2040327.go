package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 int16; F2 int8 }

func main() {
    var v1 uint16 = 894
    v2 := []uint32{0, 829}
    var v3 int64 = -1
    var v4 string = "foo bar"
    v5 := []uint32{}
    v6 := make(map[uint32]bool)
    v3 = -9223372036854775808
    if true {
        if (("" != v4) || true) {
            switch v4 {
            case "hello":
                var v7 uint32 = 2147483648
                _ = v7
            default:
                var v8 int32 = int32(v1)
                _ = v8
            }
            for fi9 := 0; fi9 < 6; fi9++ {
                switch v4 {
                case "a":
                    for fi10 := 0; fi10 < 13; fi10++ {
                        switch ("\n\t" + "\n\t") {
                        case "a":
                            v11 := GenS1{F0: -1, F1: v3}
                            _ = v11
                        }
                        var v12 int16 = 32767
                        _ = fi10
                        _ = v12
                    }
                case "foo bar":
                    switch 9223372036854775807 {
                    case -142:
                        switch (((v4 + "\n\t") + (v4 + v4)) + (v4 + "a")) {
                        case "":
                            var v13 string = (v4 + "a")
                            _ = v13
                        case "\n\t":
                            if false {
                                var v14 uint32 = 0
                                fmt.Println(v1, v1)
                                _ = v14
                            } else {
                                if ("foo bar" == v4) {
                                    var v15 int = fi9
                                    _ = v15
                                } else {
                                    _ = len(v6)
                                }
                            }
                        default:
                            var v16 float32 = 3.14159
                            _ = v16
                        }
                    }
                default:
                    for fi17 := 0; fi17 < 16; fi17++ {
                        fmt.Println(fi9, fi17, v1)
                        _ = fi17
                    }
                }
                switch "" {
                case "a":
                    var v18 float64 = (1.5 - 0.1)
                    _ = v18
                case "hello":
                    var v19 bool = (false || false)
                    _ = v19
                }
                _ = fi9
            }
            fmt.Println(v3)
        }
        if false {
            var v20 bool = (!(!true))
            _ = v20
        }
    } else {
        var v21 uint32 = 778
        _ = v21
    }
    switch (((v4 + "") + v4) + ("\n\t" + v4)) {
    case "a":
        v4 = "a"
    case "":
        for fi22 := 0; fi22 < 8; fi22++ {
            v23 := v6[0]
            v4 = (v4 + v4)
            v24 := GenS2{F0: 42}
            _ = fi22
            _ = v23
            _ = v24
        }
    case "\n\t":
        v25 := GenS2{F0: 42}
        _ = v25
    default:
        v3 += v3
    }
    v4 = v4
    _ = len(v2)
    switch 4294967295 {
    case 164:
        for fi26 := 0; fi26 < 13; fi26++ {
            v27 := GenS2{F0: 2147483647}
            var v28 float32 = 3.14159
            if false {
                fmt.Println(v3, v1, v3)
            } else {
                fmt.Println(v3, v3)
            }
            _ = fi26
            _ = v27
            _ = v28
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi29 := 0; fi29 < 6; fi29++ {
        v6[4294967295] = (false || true)
        _ = fi29
    }
    for fi30 := 0; fi30 < 2; fi30++ {
        v31 := GenS3{F0: v1, F1: 135, F2: -4}
        var v32 float64 = (float64(-2.25) + (3.14159 + (3.14159 * float64((((1.5 - (1e10 - 0.1)) + ((0.0 + 0.0) + ((1.5 * (1e10 + (-2.25 + float32(((3.14159 / (float64(-2.25) - (-2.25 / (-2.25 * 0.1)))) * -2.25))))) / 3.14159))) + 1e10)))))
        _ = fi30
        _ = v31
        _ = v32
    }
    v3 += v3
    var v33 int8 = 1
    var v34 int64 = v3
    v35 := GenS1{F0: 0, F1: v34}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
