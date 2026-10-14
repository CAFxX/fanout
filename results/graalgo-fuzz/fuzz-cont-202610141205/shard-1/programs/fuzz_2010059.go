package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]bool)
    var v3 int32 = 620
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := GenS1{F0: -11, F1: "a", F2: 2}
        var v6 bool = true
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := []int{2147483647, (fi7 & (fi4 / (((fi4 * 128)) | 1))), (fi7 ^ (fi4 - (fi7 ^ -9223372036854775808))), 4294967295}
            var v9 int64 = 9223372036854775807
            v9 = v9
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    v3 = (v3 + (v3 - (v3 / (((v3 + -2147483648)) | 1))))
    for fi10 := 0; fi10 < 4; fi10++ {
        v3 = (-(v3))
        _ = fi10
    }
    v3 = v3
    v11 := []int{-9223372036854775808, 4294967295, 32767}
    v3 = (-(v3))
    var v12 float64 = float64(1.5)
    v12 += (0.0 * 1.5)
    v12 = v12
    for fi13 := 0; fi13 < 15; fi13++ {
        if (((0.1 * 1.5) >= (-2.25 - 0.1)) || true) {
            fmt.Println(v3, v12, v12)
        } else {
            var v14 int8 = 14
            if false {
                var v15 int16 = 42
                _ = v15
            } else {
                v3 = (v3 - (v3 + (v3 & int32(v14))))
                for fi16 := 0; fi16 < 15; fi16++ {
                    if (197 != 2) {
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        fmt.Println(fi13, v14)
                    } else {
                        fmt.Println(v14, v14, v3)
                        if ("" > ("foo bar" + ("\n\t" + "a"))) {
                            v17 := GenS1{F0: v14, F1: ("\n\t" + "\n\t"), F2: 784}
                            v14 += v14
                            v3 = -107
                            v12 = v12
                            v3 += (^(v3))
                            _ = v17
                        } else {
                            var v18 uint64 = 174
                            var v19 uint64 = v18
                            _ = v18
                            _ = v19
                        }
                    }
                    _ = fi16
                }
                v3 = (^(v3))
            }
            _ = v14
        }
        switch -430 {
        case 127:
            for fi20 := 0; fi20 < 12; fi20++ {
                if ((false && ((("\n\t" + ("hello" + (((("hello" + "\n\t") + "\n\t") + "hello") + ("" + "foo bar")))) <= "foo bar") && true)) && ((v3 << ((267) & 63)) <= (v3 + v3))) {
                    var v21 float32 = 0.1
                    _ = v21
                } else {
                    var v22 int16 = 1
                    v22 += v22
                    _ = v22
                }
                _ = fi20
            }
        case -9223372036854775808:
            _ = len(v11)
        }
        v3 = (-(v3))
        _ = fi13
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        _ = len(v11)
        delete(v2, fi23)
        _ = fi23
    }
    v3 = (v3 * -303)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
}
