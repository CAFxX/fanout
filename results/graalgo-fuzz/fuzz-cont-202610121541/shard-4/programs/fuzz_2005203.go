package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint }

func main() {
    var v1 int32 = 42
    var v2 uint64 = 18446744073709551615
    v3 := make(map[string]uint32)
    var v4 int64 = 0
    v5 := []string{"foo bar"}
    v6 := []string{("\n\t" + "a"), "\n\t"}
    if (!false) {
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := v3["hello"]
            for fi9 := 0; fi9 < 14; fi9++ {
                switch ("\n\t" + "") {
                case "hello":
                    delete(v3, "foo bar")
                case "foo bar":
                    v10 := v3["foo bar"]
                    _ = len(v3)
                    v4 = (v4 % (((-(v4))) | 1))
                    _ = v10
                case "a":
                    v11 := GenS1{F0: 23}
                    _ = v11
                default:
                    if true {
                        var v12 uint8 = 60
                        v8 -= v8
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v12
                    } else {
                        v1 = 0
                    }
                }
                var v13 uint32 = v8
                v2 = (^(v2))
                var v14 int16 = 32767
                _ = fi9
                _ = v13
                _ = v14
            }
            _ = fi7
            _ = v8
        }
        var v15 int = -1
        if (v1 >= v1) {
            v16 := v3[((("a" + "\n\t") + ("foo bar" + "hello")) + ("a" + ("\n\t" + "")))]
            if false {
                var v17 float32 = 0.0
                switch v15 {
                case 904:
                    for fi18 := 0; fi18 < 1; fi18++ {
                        var v19 string = "\n\t"
                        fmt.Println(v2)
                        if len(v6) > 0 {
                            v6[0] = (v19 + (v19 + v19))
                        }
                        _ = fi18
                        _ = v19
                    }
                case 127:
                    var v20 int = (v15 / (((v15 >> ((65535) & 63))) | 1))
                    _ = v20
                case 32768:
                    var v21 bool = ((int16(v1) < 32767) || false)
                    _ = v21
                default:
                    var v22 int32 = int32(v16)
                    _ = v22
                }
                _ = v17
            }
            v6 = append(v6, "\n\t")
            _ = v16
        }
        _ = v15
    }
    if true {
        var v23 int16 = 746
        var v24 int64 = v4
        switch 0 {
        case -411:
            var v25 float64 = (float64(0.1) - (float64(0.0) * -2.25))
            _ = v25
        case 260:
            var v26 uint64 = v2
            _ = v26
        }
        _ = v23
        _ = v24
    } else {
        fmt.Println(v2, v1, v4)
    }
    switch 9223372036854775807 {
    case 4294967295:
        for fi27 := 0; fi27 < 16; fi27++ {
            var v28 int = fi27
            fmt.Println(v2, v28)
            _ = len(v3)
            _ = fi27
            _ = v28
        }
    }
    v29 := v3["hello"]
    var v30 uint = 4294967295
    v30 -= v30
    var v31 float32 = float32(((3.14159 - 0.0) + -2.25))
    v32 := GenS1{F0: v29}
    v2 -= (v2 / ((698) | 1))
    if false {
        v33 := GenS1{F0: (^(v29))}
        var v34 uint32 = 136
        if (0 < 255) {
            var v35 int16 = -1
            _ = v35
        } else {
            var v36 string = "\n\t"
            if false {
                v37 := GenS2{F0: (v30 << (((v30 << (((^(v30))) & 63))) & 63)), F1: 248, F2: (v29 % (((v34 - v29)) | 1))}
                var v38 int64 = (^(v4))
                var v39 uint8 = 42
                _ = v37
                _ = v38
                _ = v39
            }
            delete(v3, "")
            _ = v36
        }
        _ = v33
        _ = v34
    }
    v3["a"] = 0
    switch "hello" {
    case "hello":
        fmt.Println(v30, v31)
    }
    v40 := GenS3{F0: (v30 ^ 65535)}
    for fi41 := 0; fi41 < 9; fi41++ {
        for fi42 := 0; fi42 < 8; fi42++ {
            v43 := GenS2{F0: 42, F1: 128, F2: (v29 * v29)}
            v1 = v1
            if (!(true || false)) {
                v30 -= 1
                v44 := GenS4{F0: v30}
                v45 := GenS3{F0: v30}
                _ = v44
                _ = v45
            }
            _ = fi42
            _ = v43
        }
        _ = fi41
    }
    if (((1.5 * ((((float32(0.0) + 1.5) - (float32((1.5 / 0.1)) / v31)) / 0.1) / (v31 * v31))) / ((1e10 - v31) / (v31 / 1.0))) > 0.1) {
        fmt.Println(v29, v2)
        v30 = (v30 | (v30 >> (((v30 % ((v30) | 1))) & 63)))
        v46 := v3["hello"]
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
