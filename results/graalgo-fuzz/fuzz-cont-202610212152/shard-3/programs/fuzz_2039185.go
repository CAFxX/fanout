package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint64 }

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 0
    var v3 int16 = v2
    var v4 uint = 65535
    var v5 uint8 = 29
    var v6 float64 = float64(0.0)
    v7 := v1[(("\n\t" + (("\n\t" + "\n\t") + ("\n\t" + "foo bar"))) + ("" + "hello"))]
    v8 := v1[v7]
    v9 := []int{}
    var v10 int32 = -1
    var v11 uint32 = 4294967295
    if false {
        var v12 uint8 = (v5 | uint8(v4))
        v5 = v5
        for fi13 := 0; fi13 < 15; fi13++ {
            switch 694 {
            case 4294967295:
                var v14 string = (((v7 + "foo bar") + v7) + ((v8 + (((v7 + v8) + "foo bar") + (v8 + v8))) + v8))
                _ = v14
            }
            _ = fi13
        }
        _ = v12
    }
    _ = len(v1)
    if len(v9) > 0 {
        v9[0] = 0
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        switch (fi15 | (fi15 << ((v4) & 63))) {
        case 327:
            var v16 float32 = -2.25
            _ = v16
        case 42:
            if len(v9) > 0 {
                v9[0] = (fi15 & int(v4))
            }
        default:
            v11 = v11
        }
        v17 := GenS1{F0: 136}
        var v18 uint16 = uint16(v5)
        _ = fi15
        _ = v17
        _ = v18
    }
    v5 -= (v5 ^ 42)
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 float32 = (1.5 + (0.1 / 1.0))
        var v21 int32 = (v10 << (((v4 & v4)) & 63))
        var v22 int8 = 1
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    fmt.Println(v7, v6)
    v3 += (v2 | v2)
    v23 := v1[v8]
    for fi24 := 0; fi24 < 15; fi24++ {
        if (1 != -128) {
            var v25 float32 = ((0.0 / -2.25) - (float32(v6) + (0.0 / (-2.25 + (1.5 - 0.1)))))
            switch (fi24 >> ((v4) & 63)) {
            case 943:
                var v26 string = ((v23 + v23) + "a")
                _ = v26
            case 32768:
                fmt.Println(v6)
            }
            if true {
                if (!false) {
                    v27 := v1["\n\t"]
                    var v28 uint8 = (v5 & (v5 >> (((v4 >> (((^(v4))) & 63))) & 63)))
                    _ = v27
                    _ = v28
                } else {
                    if true {
                        v7 = "a"
                        var v29 string = "a"
                        for fi30 := 0; fi30 < 2; fi30++ {
                            _ = len(v1)
                            switch 1 {
                            case 0:
                                if len(v9) > 0 {
                                    v9[0] = -1
                                }
                            case -814:
                                var v31 uint8 = (v5 % ((v5) | 1))
                                _ = v31
                            case 9223372036854775807:
                                var v32 uint16 = 255
                                _ = v32
                            default:
                                var v33 uint32 = (v11 >> ((v4) & 63))
                                _ = v33
                            }
                            _ = fi30
                        }
                        _ = v29
                    } else {
                        for fi34 := 0; fi34 < 4; fi34++ {
                            if false {
                                v25 -= (1.5 * 1e10)
                            }
                            _ = len(v9)
                            var v35 float32 = v25
                            _ = fi34
                            _ = v35
                        }
                        v36 := GenS2{F0: (("foo bar" + "\n\t") + (v8 + ("hello" + v8))), F1: uint64(v4)}
                        switch 427 {
                        case 4294967295:
                            var v37 float64 = v6
                            _ = v37
                        }
                        _ = v36
                    }
                    var v38 uint = 42
                    var v39 uint16 = 42
                    _ = v38
                    _ = v39
                }
                v2 = v2
            }
            _ = v25
        }
        fmt.Println(v10)
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = fi24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
