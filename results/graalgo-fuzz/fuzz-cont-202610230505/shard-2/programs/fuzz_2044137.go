package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int32; F1 int32; F2 int8 }
type GenS3 struct { F0 uint; F1 bool; F2 string }
func genFn0() int32 {
    var v6 int16 = 171
    var v7 float32 = 0.1
    var v8 bool = ("\n\t" != "foo bar")
    v9 := []int{}
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 int32 = 0
    v2 := make(map[int]uint32)
    var v3 float64 = 0.1
    v4 := make(map[string]uint32)
    var v5 string = "a"
    v10 := GenS1{F0: -977, F1: (-9223372036854775808 <= 42)}
    v11 := []int{1, -245, 128, 128}
    v2[604] = 42
    var v12 float64 = v3
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 float32 = 0.1
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: -32768, F1: false}
    fmt.Println(v5, v1, v1)
    _ = len(v4)
    for fi16 := 0; fi16 < 13; fi16++ {
        if true {
            if (!false) {
                for fi17 := 0; fi17 < 4; fi17++ {
                    if (4294967295 == 2) {
                        if (true || true) {
                            if false {
                                if (4294967295 <= 4294967295) {
                                    v18 := GenS2{F0: (v1 / ((2147483647) | 1)), F1: genFn0(), F2: -128}
                                    var v19 uint32 = 2147483647
                                    if len(v11) > 0 {
                                        v11[0] = (fi16 >> ((255) & 63))
                                    }
                                    _ = v18
                                    _ = v19
                                }
                                var v20 int16 = 0
                                v20 = (v20 - (v20 / ((-32768) | 1)))
                                _ = v20
                            }
                            v12 -= ((v3 + v3) / 0.1)
                        } else {
                            v3 -= v3
                        }
                        v5 = (v5 + "")
                    } else {
                        fmt.Println(v3, fi16)
                    }
                    var v21 int16 = -32768
                    var v22 float64 = v12
                    _ = fi17
                    _ = v21
                    _ = v22
                }
                fmt.Println(v3, fi16)
            }
            for fi23 := 0; fi23 < 13; fi23++ {
                v24 := GenS2{F0: (v1 / ((v1) | 1)), F1: (v1 - v1), F2: 0}
                _ = fi23
                _ = v24
            }
            if (498 == 4294967295) {
                var v25 uint16 = 32767
                _ = v25
            } else {
                v4[v5] = 241
                for fi26 := 0; fi26 < 6; fi26++ {
                    var v27 float32 = ((1e10 + 1e10) / float32(v3))
                    if true {
                        v3 += 3.14159
                    }
                    _ = fi26
                    _ = v27
                }
            }
        } else {
            switch ("foo bar" + v5) {
            case "a":
                v12 += 0.0
            default:
                if (128 >= 53) {
                    v28 := GenS3{F0: 1, F1: true, F2: "hello"}
                    if true {
                        v29 := GenS2{F0: genFn0(), F1: (v1 % ((genFn0()) | 1)), F2: -75}
                        _ = v29
                    }
                    var v30 int = fi16
                    _ = v28
                    _ = v30
                } else {
                    fmt.Println(fi16)
                }
            }
            if (!true) {
                if true {
                    v31 := GenS2{F0: 42, F1: 42, F2: 100}
                    _ = v31
                }
                var v32 int64 = 42
                _ = v32
            }
            var v33 float64 = v3
            _ = v33
        }
        var v34 float32 = 0.1
        v11 = append(v11, -1)
        _ = fi16
        _ = v34
    }
    v35 := GenS1{F0: 42, F1: true}
    v36 := GenS1{F0: 1, F1: true}
    delete(v4, v5)
    var v37 int8 = -128
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    if (97 == 65535) {
        v3 -= v12
    } else {
        var v38 int64 = 0
        if (0 >= 455) {
            v5 = ""
        } else {
            v3 = v3
            fmt.Println(v1, v3)
            v37 -= (v37 % ((v37) | 1))
        }
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
