package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := make(map[string]bool)
    var v2 bool = (true && false)
    var v3 int8 = 127
    for fi4 := 0; fi4 < 20; fi4++ {
        v3 -= v3
        _ = fi4
    }
    if ((v2 || v2) && (0 != 0)) {
        v5 := GenS1{F0: 1, F1: "a", F2: (((("" + "hello") + "hello") + "\n\t") + "\n\t")}
        var v6 float64 = 1.5
        _ = v5
        _ = v6
    }
    v2 = v2
    var v7 int16 = 0
    if ((3.14159 == 1e10) || false) {
        _ = len(v1)
        if (float32(0.1) <= 1.5) {
            v8 := GenS2{F0: int(v7)}
            v2 = v2
            _ = v8
        } else {
            v9 := GenS3{F0: 10}
            if (210 > 42) {
                switch "hello" {
                case "":
                    v10 := []int{}
                    _ = v10
                case "\n\t":
                    if true {
                        var v11 uint16 = 548
                        _ = v11
                    }
                }
                v12 := GenS2{F0: -9223372036854775808}
                for fi13 := 0; fi13 < 7; fi13++ {
                    v14 := GenS1{F0: 128, F1: ("" + "a"), F2: "foo bar"}
                    v7 = v7
                    for fi15 := 0; fi15 < 18; fi15++ {
                        v16 := v1["a"]
                        fmt.Println(v2, v3, v3)
                        var v17 uint8 = 92
                        v18 := []int{(fi15 - fi13), (fi15 ^ (fi13 ^ 637)), (fi13 % ((9223372036854775807) | 1))}
                        if ((!v16) || (!false)) {
                            var v19 uint16 = 32767
                            v20 := GenS2{F0: (fi13 << ((532) & 63))}
                            v17 += (v17 - 255)
                            _ = v19
                            _ = v20
                        } else {
                            var v21 uint16 = 1
                            var v22 float64 = 3.14159
                            switch (fi13 << ((uint(v17)) & 63)) {
                            case 0:
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                            case 458:
                                v7 += (v7 & int16(v17))
                            default:
                                v7 += -32768
                            }
                            _ = v21
                            _ = v22
                        }
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    _ = fi13
                    _ = v14
                }
                _ = v12
            }
            _ = v9
        }
    }
    var v23 string = ""
    if v2 {
        v24 := GenS2{F0: 4294967295}
        var v25 uint32 = 1
        var v26 uint = 18446744073709551615
        _ = v24
        _ = v25
        _ = v26
    }
    fmt.Println(v2, v7)
    v27 := []int{368}
    v28 := GenS2{F0: 128}
    v1[""] = ((1 > 2147483647) && false)
    v2 = false
    for fi29 := 0; fi29 < 11; fi29++ {
        if (3.14159 != ((((float64(float32(-2.25)) / float64(((0.1 + 1e10) * 0.1))) - 0.0) * ((1e10 * -2.25) / 1e10)) - 0.0)) {
            var v30 uint32 = 1
            var v31 uint16 = 42
            _ = v30
            _ = v31
        } else {
            var v32 int16 = (v7 + v7)
            v7 = 1
            _ = v32
        }
        _ = fi29
    }
    var v33 int32 = 1
    fmt.Println(v33, v23)
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v7
}
