package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 int16 }

func main() {
    var v1 int8 = 1
    var v2 int64 = -9223372036854775808
    v3 := make(map[string]uint32)
    var v4 int8 = v1
    fmt.Println(v4, v1, v1)
    v3[("hello" + "a")] = 138
    v1 = (v1 << ((143) & 63))
    fmt.Println(v1)
    v5 := []int{32767, 801, 1, 689}
    v2 -= (v2 / ((int64(v1)) | 1))
    delete(v3, "a")
    _ = len(v3)
    if false {
        v1 = 1
        var v6 bool = ((v2 & v2) <= -1)
        _ = v6
    } else {
        fmt.Println(v2)
        if (!(4294967295 > 4294967295)) {
            if false {
                for fi7 := 0; fi7 < 1; fi7++ {
                    v8 := GenS1{F0: -9223372036854775808, F1: (("" + "\n\t") + "a"), F2: (v1 * v1)}
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    switch "" {
                    case "foo bar":
                        for fi9 := 0; fi9 < 5; fi9++ {
                            var v10 uint = 341
                            _ = fi9
                            _ = v10
                        }
                    default:
                        var v11 uint8 = 1
                        _ = v11
                    }
                    _ = fi7
                    _ = v8
                }
                var v12 float64 = 3.14159
                if false {
                    if false {
                        var v13 int64 = (v2 % ((v2) | 1))
                        _ = v13
                    }
                } else {
                    _ = len(v5)
                }
                _ = v12
            } else {
                var v14 float32 = (float32(float64(((1e10 + 3.14159) * (3.14159 + 1.5)))) + 1e10)
                _ = v14
            }
            for fi15 := 0; fi15 < 15; fi15++ {
                var v16 int = 0
                v17 := GenS1{F0: 195, F1: (("" + "") + ("foo bar" + "")), F2: (-(v1))}
                _ = fi15
                _ = v16
                _ = v17
            }
            fmt.Println(v4, v1, v2)
        }
        v18 := GenS2{F0: 65535, F1: 202, F2: 1}
        _ = v18
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        v20 := GenS2{F0: 42, F1: 0, F2: 0}
        _ = fi19
        _ = v20
    }
    v21 := GenS1{F0: v2, F1: "a", F2: (v1 / ((v4) | 1))}
    var v22 bool = true
    if v22 {
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 bool = v22
            _ = fi23
            _ = v24
        }
        for fi25 := 0; fi25 < 10; fi25++ {
            for fi26 := 0; fi26 < 11; fi26++ {
                v27 := GenS1{F0: 0, F1: "hello", F2: 1}
                _ = fi26
                _ = v27
            }
            v28 := GenS3{F0: 1}
            fmt.Println(v2, v2)
            _ = fi25
            _ = v28
        }
        var v29 int16 = -244
        _ = v29
    } else {
        v30 := GenS3{F0: -123}
        var v31 int8 = -69
        if v22 {
            if (v22 || (true && ((2 == 575) || (!false)))) {
                v1 -= v31
                if (-710 < 2147483647) {
                    v31 += v31
                    var v32 string = "foo bar"
                    _ = v32
                }
                v33 := v3["a"]
                _ = v33
            } else {
                if (v1 <= 100) {
                    var v34 bool = (v22 || (!false))
                    switch 32768 {
                    case -102:
                        var v35 int32 = -1
                        _ = v35
                    default:
                        v31 = 2
                    }
                    _ = v34
                } else {
                    if v22 {
                        v36 := GenS4{F0: v1, F1: 32767}
                        v31 -= (v4 / (((v4 >> ((0) & 63))) | 1))
                        var v37 uint = 2
                        _ = v36
                        _ = v37
                    }
                }
                fmt.Println(v4, v4, v31)
            }
            v38 := GenS2{F0: 134, F1: 255, F2: 0}
            v31 = v4
            _ = v38
        }
        _ = v30
        _ = v31
    }
    v1 += (^(v4))
    fmt.Println(v1)
    fmt.Println(v22, v2)
    v3[("" + ("a" + "a"))] = 4294967295
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
