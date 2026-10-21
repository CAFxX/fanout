package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 float32 = float32(1.5)
    v2 := []uint32{}
    var v3 uint32 = 4294967295
    v4 := make(map[string]int)
    var v5 bool = (v3 < 2147483648)
    v3 += (v3 * 65535)
    switch "" {
    case "a":
        v3 = v3
    case "\n\t":
        for fi6 := 0; fi6 < 2; fi6++ {
            v5 = v5
            _ = fi6
        }
    }
    _ = len(v4)
    fmt.Println(v3, v5)
    if ((v1 <= v1) && (v1 < v1)) {
        var v7 uint32 = 65535
        for fi8 := 0; fi8 < 3; fi8++ {
            if v5 {
                v1 = float32(0.1)
                v9 := GenS1{F0: uint8(v3), F1: (^(v3)), F2: 0}
                if (v5 && (1 > 844)) {
                    v7 -= v3
                    for fi10 := 0; fi10 < 11; fi10++ {
                        var v11 uint16 = 631
                        _ = fi10
                        _ = v11
                    }
                } else {
                    var v12 float64 = 1e10
                    fmt.Println(v3, v7)
                    v13 := v4[("" + "")]
                    _ = v12
                    _ = v13
                }
                _ = v9
            } else {
                var v14 int = (fi8 / ((fi8) | 1))
                for fi15 := 0; fi15 < 16; fi15++ {
                    var v16 float64 = -2.25
                    if v5 {
                        v5 = (!v5)
                        var v17 int = 42
                        v18 := GenS2{F0: "hello", F1: ""}
                        _ = v17
                        _ = v18
                    } else {
                        if (true && (818 != 42)) {
                            var v19 uint8 = 14
                            _ = v19
                        }
                        v4[""] = -638
                    }
                    fmt.Println(fi8, v5, fi8)
                    _ = fi15
                    _ = v16
                }
                var v20 string = "hello"
                _ = v14
                _ = v20
            }
            if (-32768 == 238) {
                if (2 == 4294967295) {
                    _ = len(v2)
                    var v21 uint64 = 858
                    v5 = v5
                    _ = v21
                } else {
                    v1 += v1
                    v7 -= v3
                }
            }
            _ = len(v2)
            _ = fi8
        }
        _ = v7
    }
    fmt.Println(v3, v3)
    v4["\n\t"] = 304
    for fi22 := 0; fi22 < 4; fi22++ {
        switch (fi22 / ((fi22) | 1)) {
        case -743:
            if false {
                v4[("\n\t" + (("hello" + "hello") + "hello"))] = 32768
                for fi23 := 0; fi23 < 19; fi23++ {
                    v3 -= 1
                    if (497 < (^(v3))) {
                        v5 = (32767 == 255)
                        switch "foo bar" {
                        case "hello":
                            var v24 uint64 = uint64(v3)
                            _ = v24
                        default:
                            var v25 string = ("\n\t" + "foo bar")
                            _ = v25
                        }
                    } else {
                        v26 := v4["\n\t"]
                        v3 += v3
                        _ = v26
                    }
                    _ = fi23
                }
                var v27 int16 = -982
                _ = v27
            } else {
                var v28 int64 = -9223372036854775808
                var v29 float64 = 0.1
                _ = v28
                _ = v29
            }
        case 42:
            for fi30 := 0; fi30 < 7; fi30++ {
                v31 := GenS3{F0: 4294967295, F1: 1}
                _ = fi30
                _ = v31
            }
        case -1:
            var v32 uint64 = 1
            _ = v32
        }
        v3 -= 628
        if true {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if len(v2) > 0 {
                v2[0] = 405
            }
        }
        _ = fi22
    }
    v33 := GenS4{F0: 350}
    if len(v2) > 0 {
        v2[0] = 32
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
