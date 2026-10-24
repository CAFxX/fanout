package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int }
type GenS2 struct { F0 bool; F1 uint; F2 uint16 }
type GenS3 struct { F0 int; F1 uint64; F2 int32 }
func genFn0(p0 uint) int32 {
    p0 += (p0 >> (((^(p0))) & 63))
    var v5 uint8 = 255
    var v6 uint32 = 556
    v7 := []int{32768, 128, 1, 32768}
    switch -1 {
    case 434:
        var v8 uint8 = v5
        _ = v8
    case 127:
        for fi9 := 0; fi9 < 10; fi9++ {
            switch (fi9 ^ (fi9 / ((fi9) | 1))) {
            case 9223372036854775807:
                p0 += (p0 * 117)
            case -161:
                fmt.Println(fi9)
            case 1:
                fmt.Println(fi9, v5)
            default:
                switch fi9 {
                case -9223372036854775808:
                    if true {
                        v10 := GenS1{F0: (v6 | (^(v6))), F1: 32767, F2: (fi9 - (fi9 % (((fi9 << (((p0 - 1)) & 63))) | 1)))}
                        var v11 uint32 = v6
                        _ = v10
                        _ = v11
                    } else {
                        var v12 uint8 = v5
                        _ = v12
                    }
                }
            }
            v13 := GenS1{F0: (v6 | v6), F1: 32767, F2: fi9}
            v14 := GenS1{F0: 42, F1: 1000, F2: (^(fi9))}
            _ = fi9
            _ = v13
            _ = v14
        }
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 int64) int64 {
    fmt.Println(p0)
    switch ("hello" + ("hello" + (("\n\t" + "\n\t") + "a"))) {
    case "hello":
        p0 -= p0
    default:
        var v15 int = 9223372036854775807
        _ = v15
    }
    p0 -= -482
    var v16 int64 = (^(p0))
    _ = p0
    _ = v16
    return v16
}

func main() {
    var v1 int8 = -128
    v2 := []uint32{20, 936, 2147483648, 2147483648}
    v3 := make(map[string]bool)
    var v4 string = "hello"
    fmt.Println(v1)
    fmt.Println(v1)
    fmt.Println(v1, v1)
    for fi17 := 0; fi17 < 12; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            if true {
                v3["hello"] = (0.1 != 0.0)
                var v19 int32 = int32(fi17)
                _ = v19
            }
            v4 = ""
            fmt.Println(v4, fi17, v4)
            _ = fi18
        }
        _ = fi17
    }
    var v20 int32 = 42
    v21 := GenS1{F0: 2, F1: 1000, F2: 42}
    v20 = v20
    var v22 int8 = v1
    for fi23 := 0; fi23 < 20; fi23++ {
        switch fi23 {
        case 128:
            if true {
                var v24 float32 = 0.0
                _ = v24
            } else {
                if ((65535 < 944) || ((1 <= 4294967295) || true)) {
                    var v25 uint = 0
                    _ = v25
                } else {
                    fmt.Println(v1, fi23)
                }
                for fi26 := 0; fi26 < 4; fi26++ {
                    delete(v3, v4)
                    var v27 uint32 = 480
                    var v28 uint64 = 199
                    delete(v3, "")
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
        default:
            if len(v2) > 0 {
                v2[0] = 2
            }
        }
        for fi29 := 0; fi29 < 19; fi29++ {
            if false {
                var v30 uint32 = 2
                var v31 bool = ((v30 | v30) == v30)
                v32 := GenS1{F0: v30, F1: 42, F2: -526}
                _ = v30
                _ = v31
                _ = v32
            } else {
                if ((1.5 != float32((1e10 - (0.0 + (1.5 + ((-2.25 * 1e10) + 1.5)))))) && true) {
                    var v33 uint = 350
                    delete(v3, v4)
                    if (2 > 2147483647) {
                        var v34 int32 = (v20 | int32(v33))
                        fmt.Println(v1, fi23, v34)
                        v35 := GenS2{F0: (v20 != 42), F1: 4294967295, F2: 214}
                        switch ("" + (v4 + (v4 + "\n\t"))) {
                        case "\n\t":
                            var v36 uint32 = 782
                            _ = v36
                        default:
                            var v37 int32 = -602
                            _ = v37
                        }
                        _ = v34
                        _ = v35
                    }
                    _ = v33
                }
            }
            var v38 int8 = v22
            fmt.Println(v20)
            _ = fi29
            _ = v38
        }
        v22 += (v22 << ((0) & 63))
        _ = fi23
    }
    var v39 float64 = (1e10 / 1.5)
    for fi40 := 0; fi40 < 6; fi40++ {
        var v41 uint16 = 32767
        _ = fi40
        _ = v41
    }
    v22 = (v22 - -128)
    v42 := GenS1{F0: 2, F1: 1, F2: 666}
    if true {
        if false {
            v43 := GenS3{F0: 559, F1: 1, F2: (v20 << ((368) & 63))}
            _ = v43
        }
        fmt.Println(v20, v39)
        if (-1 >= 4294967295) {
            delete(v3, v4)
            v22 -= v1
        }
    } else {
        var v44 int32 = genFn0(65535)
        _ = v44
    }
    delete(v3, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v39
    _ = v4
    _ = v42
}
