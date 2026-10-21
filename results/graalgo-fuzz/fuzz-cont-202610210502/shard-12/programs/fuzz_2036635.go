package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }

func main() {
    var v1 uint = 2
    v2 := []uint32{0, 275}
    v3 := make(map[string]string)
    v4 := []string{"\n\t", "a", ("hello" + "hello"), "a"}
    if (64 > 940) {
        v5 := v3["foo bar"]
        if (1 > 1) {
            v5 = ""
            fmt.Println(v5, v5)
            var v6 int = 2147483647
            _ = v6
        } else {
            v7 := GenS1{F0: -292, F1: -71}
            var v8 bool = true
            v9 := v3[v5]
            _ = v7
            _ = v8
            _ = v9
        }
        _ = v5
    } else {
        v1 -= (v1 * 255)
    }
    fmt.Println(v1)
    for fi10 := 0; fi10 < 4; fi10++ {
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 uint16 = 2
            v13 := GenS2{F0: 1}
            var v14 uint16 = 255
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        v1 = (^(v1))
        _ = fi10
    }
    v3["foo bar"] = "\n\t"
    v1 = (v1 % (((v1 | v1)) | 1))
    var v15 uint16 = 1
    v15 += (v15 / ((v15) | 1))
    var v16 int32 = 42
    for fi17 := 0; fi17 < 13; fi17++ {
        var v18 float32 = 3.14159
        var v19 uint = v1
        _ = fi17
        _ = v18
        _ = v19
    }
    v1 += v1
    v20 := GenS2{F0: 42}
    if true {
        switch 32767 {
        case 325:
            v1 += (v1 & (v1 / ((v1) | 1)))
        default:
            v16 = (v16 * (v16 & 346))
        }
        if (v1 > (v1 % (((^(v1))) | 1))) {
            if false {
                if true {
                    if true {
                        v21 := GenS3{F0: v1}
                        var v22 int32 = (v16 << ((4294967295) & 63))
                        _ = v21
                        _ = v22
                    } else {
                        fmt.Println(v15, v15)
                        if (-32 < -1) {
                            v23 := GenS2{F0: 127}
                            v15 = v15
                            for fi24 := 0; fi24 < 9; fi24++ {
                                switch 128 {
                                case 2147483647:
                                    v16 -= v16
                                }
                                if true {
                                    var v25 string = "\n\t"
                                    _ = v25
                                }
                                v4 = append(v4, "foo bar")
                                _ = fi24
                            }
                            _ = v23
                        }
                        v16 += (v16 ^ 1)
                    }
                } else {
                    if true {
                        for fi26 := 0; fi26 < 7; fi26++ {
                            if ((128 == 203) && true) {
                                for fi27 := 0; fi27 < 8; fi27++ {
                                    v28 := GenS3{F0: (v1 & 255)}
                                    _ = fi27
                                    _ = v28
                                }
                            }
                            _ = fi26
                        }
                        v15 = v15
                        for fi29 := 0; fi29 < 7; fi29++ {
                            switch "a" {
                            case "a":
                                v15 += (v15 >> ((v1) & 63))
                            case "hello":
                                for fi30 := 0; fi30 < 4; fi30++ {
                                    if true {
                                        var v31 uint32 = 1
                                        _ = v31
                                    } else {
                                        if ("foo bar" >= "\n\t") {
                                            v32 := GenS2{F0: 42}
                                            var v33 int64 = 1
                                            v33 -= v33
                                            _ = v32
                                            _ = v33
                                        }
                                    }
                                    var v34 uint16 = 42
                                    v15 = v15
                                    _ = fi30
                                    _ = v34
                                }
                            }
                            _ = fi29
                        }
                    } else {
                        var v35 string = ("\n\t" + ("" + "hello"))
                        v4 = append(v4, (((v35 + "hello") + (v35 + (v35 + v35))) + "\n\t"))
                        _ = v35
                    }
                    v36 := GenS1{F0: 42, F1: 0}
                    _ = v36
                }
                var v37 uint = 4294967295
                _ = v37
            }
        } else {
            v16 = v16
        }
        if (((v15 / (((v15 - v15)) | 1)) < v15) && (!false)) {
            v15 -= (v15 * v15)
        }
    }
    fmt.Println(v16, v15)
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
