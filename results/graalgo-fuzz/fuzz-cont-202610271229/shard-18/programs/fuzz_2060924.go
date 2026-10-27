package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]uint32)
    v3 := []uint32{2, 2147483647}
    var v4 uint = 0
    v4 = (v4 % (((^(v4))) | 1))
    if ((("a" + ("hello" + ("foo bar" + "a"))) + "\n\t") >= ("\n\t" + ("" + ("\n\t" + "hello")))) {
        v5 := v1[127]
        if false {
            v5 += 2
            v4 -= (^(v4))
        } else {
            v5 -= 146
            var v6 bool = (0 < 128)
            _ = v6
        }
        _ = v5
    } else {
        var v7 uint16 = 2
        _ = v7
    }
    var v8 int8 = 1
    v8 -= 12
    v4 = v4
    var v9 uint64 = 114
    delete(v1, int(v4))
    var v10 uint64 = 0
    v11 := v2[684]
    var v12 int16 = int16(v11)
    v13 := GenS1{F0: -34, F1: uint16(v11)}
    var v14 string = ""
    v14 = ((("a" + (("\n\t" + (((v14 + v14) + (v14 + v14)) + ((("" + ("\n\t" + (("\n\t" + ("hello" + "hello")) + v14))) + (v14 + (((("hello" + ((v14 + v14) + "foo bar")) + ((v14 + v14) + (((v14 + v14) + "foo bar") + v14))) + v14) + v14))) + "foo bar"))) + ((v14 + ("\n\t" + v14)) + "\n\t"))) + v14) + ("a" + "foo bar"))
    v14 = ("hello" + "")
    switch int(v12) {
    case 42:
        for fi15 := 0; fi15 < 14; fi15++ {
            _ = len(v1)
            v16 := GenS1{F0: (v8 - (v8 % ((v8) | 1))), F1: 32767}
            _ = fi15
            _ = v16
        }
    case 32767:
        _ = len(v2)
    default:
        if true {
            v9 = v10
        } else {
            var v17 int = 2147483647
            if true {
                if true {
                    var v18 int8 = -1
                    switch v14 {
                    case "a":
                        var v19 bool = (v11 < (^(v11)))
                        _ = v19
                    case "foo bar":
                        var v20 float64 = float64(0.0)
                        _ = v20
                    case "":
                        var v21 uint32 = 428
                        _ = v21
                    default:
                        if true {
                            v22 := GenS2{F0: 0, F1: 42}
                            for fi23 := 0; fi23 < 20; fi23++ {
                                v14 = ("foo bar" + (v14 + "\n\t"))
                                if len(v3) > 0 {
                                    v3[0] = v11
                                }
                                _ = fi23
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v22
                        } else {
                            switch (v14 + "") {
                            case "":
                                v10 = (v9 + v9)
                            case "hello":
                                v10 = 4294967295
                            default:
                                fmt.Println(v11, v17, v9)
                            }
                        }
                    }
                    v11 = 403
                    _ = v18
                } else {
                    for fi24 := 0; fi24 < 17; fi24++ {
                        v4 = v4
                        switch (v17 % ((v17) | 1)) {
                        case 227:
                            v9 -= v10
                        case 32768:
                            var v25 int8 = v8
                            _ = v25
                        default:
                            v26 := v2[uint32(v10)]
                            _ = v26
                        }
                        _ = fi24
                    }
                    if true {
                        v12 += 42
                        switch (v17 << ((v4) & 63)) {
                        case -1:
                            for fi27 := 0; fi27 < 12; fi27++ {
                                v28 := GenS2{F0: 2, F1: v4}
                                _ = fi27
                                _ = v28
                            }
                        case 4294967295:
                            var v29 int64 = -127
                            _ = v29
                        }
                    }
                    if (1 == 119) {
                        for fi30 := 0; fi30 < 3; fi30++ {
                            v31 := GenS3{F0: -91}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi30
                            _ = v31
                        }
                    } else {
                        if true {
                            v32 := GenS1{F0: int8(v4), F1: 913}
                            _ = v32
                        }
                        fmt.Println(v17, v8)
                    }
                }
                var v33 float32 = 3.14159
                _ = v33
            }
            v34 := GenS3{F0: -1}
            _ = v17
            _ = v34
        }
    }
    var v35 uint32 = (^(v11))
    var v36 int64 = int64(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v8
    _ = v9
}
