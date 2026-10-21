package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint32 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint16; F1 uint; F2 string }

func main() {
    var v1 float64 = (3.14159 - 0.0)
    var v2 int64 = 1
    v3 := make(map[string]int)
    if (true && (602 != 25)) {
        var v4 int = -534
        v4 = (v4 >> ((842) & 63))
        _ = v4
    } else {
        v5 := GenS1{F0: (true && true), F1: -128, F2: 2147483647}
        v1 = v1
        v1 = float64(0.1)
        _ = v5
    }
    var v6 float64 = (v1 + (3.14159 + float64(-2.25)))
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 float32 = 3.14159
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 int32 = 709
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    var v11 uint = 1
    var v12 string = (((("hello" + ("foo bar" + "a")) + "\n\t") + "foo bar") + ("foo bar" + ((("a" + "hello") + "\n\t") + "\n\t")))
    v13 := GenS2{F0: -1, F1: 32768}
    if false {
        v1 += ((((v6 * v6) - v1) * float64((-2.25 + -2.25))) + v1)
        var v14 bool = true
        v1 = v6
        _ = v14
    }
    v15 := GenS2{F0: 42, F1: 0}
    var v16 int16 = 42
    delete(v3, v12)
    v11 = 2
    if true {
        for fi17 := 0; fi17 < 5; fi17++ {
            for fi18 := 0; fi18 < 4; fi18++ {
                v12 = v12
                if (868 >= 0) {
                    _ = len(v3)
                    fmt.Println(v11, v6)
                    v11 -= (v11 >> ((v11) & 63))
                    var v19 string = (v12 + v12)
                    _ = v19
                } else {
                    if true {
                        _ = len(v3)
                        v20 := GenS1{F0: false, F1: 42, F2: 0}
                        switch fi17 {
                        case 9223372036854775807:
                            v21 := GenS1{F0: true, F1: 0, F2: 547}
                            _ = v21
                        case -26:
                            v22 := []int{(fi17 / (((-(fi17))) | 1)), fi18, fi17, fi17}
                            _ = v22
                        case -428:
                            v2 -= (^(v2))
                        default:
                            fmt.Println(fi17, v16, v2)
                        }
                        _ = v20
                    } else {
                        if true {
                            switch v12 {
                            case "\n\t":
                                if (!false) {
                                    v1 += (float64((float32(v1) * ((0.0 / (1e10 * -2.25)) - (0.1 * float32(1e10))))) / 1.0)
                                    var v23 uint64 = 0
                                    _ = v23
                                } else {
                                    if true {
                                        var v24 uint32 = 0
                                        if true {
                                            v2 = v2
                                            var v25 int64 = -9223372036854775808
                                            _ = v25
                                        }
                                        _ = v24
                                    }
                                }
                            case "a":
                                var v26 int8 = -45
                                _ = v26
                            case "hello":
                                v2 = v2
                            }
                        }
                    }
                }
                _ = fi18
            }
            v27 := GenS3{F0: 0, F1: (v11 - (^(v11))), F2: (v12 + ("a" + ""))}
            _ = fi17
            _ = v27
        }
        v12 = ""
    }
    v3[v12] = -227
    var v28 int16 = (v16 + v16)
    v29 := []int{}
    v11 -= (^(v11))
    var v30 string = v12
    var v31 float32 = (float32(float64(-2.25)) + 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v6
}
