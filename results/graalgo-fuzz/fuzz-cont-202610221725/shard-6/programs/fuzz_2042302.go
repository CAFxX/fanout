package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int64; F2 uint }
type GenS3 struct { F0 uint }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]uint32)
    var v3 bool = ((360 > 419) || true)
    var v4 int16 = 1000
    var v5 int64 = -1
    v6 := []uint32{1, 1, 704}
    v5 = v5
    v6 = append(v6, 2147483647)
    var v7 float64 = ((((0.1 + (0.1 + float64(((0.1 - (0.1 / 0.1)) * 1e10)))) + float64(1.5)) * 1e10) + (-2.25 - (1e10 * (1.5 / (1e10 + 0.0)))))
    var v8 uint16 = 42
    if v3 {
        var v9 uint16 = uint16(v1)
        if false {
            for fi10 := 0; fi10 < 4; fi10++ {
                var v11 uint = 204
                _ = fi10
                _ = v11
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        } else {
            switch (("foo bar" + ("hello" + "")) + ("hello" + "foo bar")) {
            case "hello":
                for fi12 := 0; fi12 < 4; fi12++ {
                    var v13 uint8 = 128
                    _ = fi12
                    _ = v13
                }
            case "a":
                _ = len(v2)
            case "foo bar":
                if false {
                    v14 := GenS1{F0: 9223372036854775807}
                    for fi15 := 0; fi15 < 9; fi15++ {
                        for fi16 := 0; fi16 < 4; fi16++ {
                            var v17 uint = 749
                            if ((v4 + v4) != 32767) {
                                var v18 int64 = (v5 - (v5 * (v5 + v5)))
                                v19 := GenS2{F0: v9, F1: 259, F2: (v17 % ((v17) | 1))}
                                _ = v18
                                _ = v19
                            } else {
                                v17 += 2
                            }
                            switch "" {
                            case "foo bar":
                                v20 := v2[(fi16 + fi16)]
                                _ = v20
                            case "\n\t":
                                var v21 uint64 = 2
                                _ = v21
                            case "hello":
                                if (1 >= 41) {
                                    fmt.Println(fi15, v7, v8)
                                    var v22 int = (fi15 % (((fi15 & 520)) | 1))
                                    _ = v22
                                } else {
                                    delete(v2, (fi16 * fi15))
                                }
                                v9 -= 1
                                fmt.Println(fi16, v1)
                                if v3 {
                                    var v23 bool = (false || true)
                                    var v24 int32 = -833
                                    var v25 float32 = (3.14159 * ((1.5 + -2.25) - 3.14159))
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                            }
                            var v26 float32 = 3.14159
                            _ = fi16
                            _ = v17
                            _ = v26
                        }
                        _ = fi15
                    }
                    _ = v14
                }
            default:
                _ = len(v2)
            }
        }
        v27 := GenS3{F0: 1}
        _ = v27
        _ = v9
    }
    var v28 uint8 = 2
    for fi29 := 0; fi29 < 19; fi29++ {
        if (fi29 == -548) {
            for fi30 := 0; fi30 < 14; fi30++ {
                var v31 float32 = -2.25
                var v32 uint64 = 4294967295
                var v33 uint = 0
                _ = fi30
                _ = v31
                _ = v32
                _ = v33
            }
            fmt.Println(v4)
        } else {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            for fi34 := 0; fi34 < 10; fi34++ {
                var v35 float32 = float32(float64(float32(v7)))
                _ = fi34
                _ = v35
            }
            if len(v6) > 0 {
                v6[0] = 42
            }
        }
        _ = fi29
    }
    fmt.Println(v8, v5, v1)
    var v36 uint16 = (v8 + (v8 * (v8 * (v8 + 0))))
    var v37 float32 = float32(((-2.25 * v7) * ((((0.0 / 3.14159) - 0.1) * ((((v7 * float64(3.14159)) - float64((0.1 / 1.5))) - v7) + ((0.1 / 3.14159) + v7))) - 0.1)))
    if false {
        v6 = append(v6, 1)
    } else {
        var v38 int = -805
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        switch "" {
        case "\n\t":
            var v39 int8 = 127
            _ = v39
        }
        _ = v38
    }
    v6 = append(v6, 2)
    var v40 int32 = 42
    v1 += v1
    for fi41 := 0; fi41 < 1; fi41++ {
        v28 -= (^(v28))
        var v42 uint32 = 2147483647
        fmt.Println(v3, v42)
        _ = fi41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
