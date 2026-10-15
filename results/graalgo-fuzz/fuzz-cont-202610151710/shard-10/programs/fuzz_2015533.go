package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int }

func main() {
    var v1 int16 = 32767
    v2 := []string{("hello" + "\n\t"), "", "hello"}
    var v3 uint16 = 65535
    var v4 float32 = float32((-2.25 * (((1e10 + (1e10 / (float64(-2.25) * 3.14159))) - 1e10) / (((0.0 * (-2.25 + 1e10)) + ((0.0 * 1.5) - float64(0.1))) / -2.25))))
    v5 := make(map[uint32]bool)
    var v6 int64 = -117
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    delete(v5, 0)
    for fi7 := 0; fi7 < 9; fi7++ {
        if false {
            var v8 int8 = -63
            _ = v8
        } else {
            v9 := GenS1{F0: v6, F1: v6}
            _ = v9
        }
        var v10 float32 = -2.25
        v5[42] = true
        _ = fi7
        _ = v10
    }
    fmt.Println(v6)
    var v11 int32 = 0
    var v12 uint = 0
    if (1.5 == v4) {
        v13 := GenS2{F0: 128}
        if (v3 <= (v3 << ((v12) & 63))) {
            var v14 int = int(v3)
            _ = v14
        } else {
            if false {
                for fi15 := 0; fi15 < 6; fi15++ {
                    v16 := GenS2{F0: (fi15 ^ (fi15 & fi15))}
                    var v17 int = (fi15 % ((32768) | 1))
                    v18 := v5[2147483647]
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                }
                v11 = 118
            }
            var v19 int32 = v11
            var v20 int16 = 1000
            _ = v19
            _ = v20
        }
        _ = v13
    }
    var v21 int16 = (-(v1))
    var v22 uint = (v12 / (((v12 * v12)) | 1))
    v1 = v21
    fmt.Println(v21, v3, v4)
    switch ((("\n\t" + "hello") + ("a" + "foo bar")) + ("" + "a")) {
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "":
        switch "hello" {
        case "a":
            for fi23 := 0; fi23 < 10; fi23++ {
                v2 = append(v2, "")
                var v24 int64 = v6
                v4 = v4
                _ = fi23
                _ = v24
            }
        case "\n\t":
            if false {
                v25 := GenS1{F0: (v6 - v6), F1: -297}
                var v26 int16 = -370
                if (!true) {
                    v27 := GenS1{F0: v6, F1: (v6 | (-(v6)))}
                    var v28 int = -9223372036854775808
                    _ = v27
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 4; fi29++ {
                        for fi30 := 0; fi30 < 9; fi30++ {
                            var v31 int32 = -706
                            if false {
                                v11 += 42
                            }
                            switch (fi30 >> ((509) & 63)) {
                            case -51:
                                if (560 <= uint32(v12)) {
                                    var v32 float64 = 1.5
                                    _ = v32
                                } else {
                                    var v33 int32 = (v11 << (((v12 | v12)) & 63))
                                    v3 = 50
                                    var v34 uint8 = 0
                                    _ = v33
                                    _ = v34
                                }
                            }
                            _ = fi30
                            _ = v31
                        }
                        _ = fi29
                    }
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                    for fi35 := 0; fi35 < 16; fi35++ {
                        if (true || (v26 <= 1000)) {
                            _ = len(v2)
                        }
                        v1 = (v21 - v26)
                        _ = fi35
                    }
                }
                _ = v25
                _ = v26
            }
        case "hello":
            var v36 string = ("foo bar" + "hello")
            _ = v36
        }
    default:
        var v37 float64 = 1e10
        _ = v37
    }
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
