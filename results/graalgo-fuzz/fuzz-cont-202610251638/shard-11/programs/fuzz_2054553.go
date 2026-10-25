package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 uint8 }

func main() {
    v1 := []int{}
    var v2 uint = 255
    var v3 bool = (true || false)
    v4 := []string{"", ("foo bar" + "foo bar"), "foo bar"}
    v5 := []string{""}
    var v6 int = 0
    v6 = 32767
    v3 = v3
    v2 += 4294967295
    if v3 {
        var v7 float64 = 1.5
        var v8 int16 = 592
        _ = v7
        _ = v8
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if (1e10 > float64((1e10 - (1e10 - 3.14159)))) {
        v3 = v3
    } else {
        if ((^(v2)) > v2) {
            if (int64(v6) >= 1) {
                if len(v1) > 0 {
                    v1[0] = 127
                }
                v3 = v3
            }
            if false {
                if ((1.5 + 1e10) == (1.5 + (float32(((1.5 * 0.0) - 1.5)) / 3.14159))) {
                    v9 := GenS1{F0: 238, F1: 1, F2: int32(v2)}
                    var v10 bool = false
                    switch v6 {
                    case 621:
                        switch (v6 & v6) {
                        case -9223372036854775808:
                            var v11 string = "foo bar"
                            _ = v11
                        case 32768:
                            v10 = v3
                        default:
                            if true {
                                var v12 uint32 = 2147483648
                                for fi13 := 0; fi13 < 15; fi13++ {
                                    v12 += v12
                                    v3 = false
                                    _ = fi13
                                }
                                v14 := GenS2{F0: (("\n\t" + ("\n\t" + "foo bar")) + "a")}
                                _ = v12
                                _ = v14
                            }
                        }
                    case 9223372036854775807:
                        v3 = false
                    case 176:
                        if false {
                            v15 := GenS2{F0: ("a" + "a")}
                            var v16 int = (v6 >> ((v2) & 63))
                            v3 = true
                            _ = v15
                            _ = v16
                        }
                    default:
                        var v17 float64 = ((float64(0.0) / (0.1 * 1e10)) - (-2.25 * float64(3.14159)))
                        _ = v17
                    }
                    _ = v10
                    _ = v9
                } else {
                    if v3 {
                        v6 = (v6 - (v6 + 42))
                        v18 := GenS3{F0: v2, F1: 1}
                        _ = v18
                    }
                    if v3 {
                        v19 := GenS4{F0: (v2 << (((v2 | v2)) & 63)), F1: v3}
                        _ = v19
                    }
                    var v20 uint64 = 0
                    _ = v20
                }
                var v21 float64 = 3.14159
                _ = v21
            } else {
                var v22 string = "\n\t"
                for fi23 := 0; fi23 < 9; fi23++ {
                    v24 := GenS5{F0: 101}
                    for fi25 := 0; fi25 < 13; fi25++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi25
                    }
                    _ = fi23
                    _ = v24
                }
                var v26 int64 = -1
                _ = v22
                _ = v26
            }
        } else {
            var v27 float32 = 3.14159
            fmt.Println(v2, v6, v3)
            _ = v27
        }
        if (-9223372036854775808 > 838) {
            var v28 bool = (0 < -735)
            if true {
                v28 = (v28 || v28)
            }
            _ = v28
        }
    }
    switch "a" {
    case "hello":
        v3 = v3
    case "a":
        if (127 == 255) {
            v3 = false
        } else {
            switch v6 {
            case 874:
                v6 = 42
            }
        }
    default:
        if len(v5) > 0 {
            v5[0] = ((("a" + "hello") + ("foo bar" + "foo bar")) + "\n\t")
        }
    }
    v3 = (v3 && false)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v2)
    var v29 uint = v2
    fmt.Println(v3, v2, v29)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
