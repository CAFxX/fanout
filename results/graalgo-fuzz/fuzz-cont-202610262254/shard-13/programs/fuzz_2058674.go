package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
func genFn0(p0 uint16, p1 int32) int32 {
    for fi4 := 0; fi4 < 20; fi4++ {
        p0 -= (p0 | p0)
        var v5 float64 = 0.1
        var v6 float64 = v5
        _ = fi4
        _ = v5
        _ = v6
    }
    if false {
        v7 := []int{32767, 9223372036854775807, 773}
        _ = len(v7)
        _ = v7
    } else {
        var v8 int8 = 1
        if (0 >= uint32(p0)) {
            var v9 uint16 = 154
            var v10 float32 = 1.5
            if false {
                v11 := GenS1{F0: (("foo bar" + ("\n\t" + "hello")) + "\n\t"), F1: 2, F2: 32767}
                p0 = p0
                v12 := GenS1{F0: "hello", F1: 0, F2: 1000}
                _ = v11
                _ = v12
            } else {
                fmt.Println(v8)
                v9 -= (p0 << ((42) & 63))
            }
            _ = v10
            _ = v9
        }
        _ = v8
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 string = (("" + "hello") + "")
    var v2 uint8 = 128
    var v3 uint8 = (v2 % ((255) | 1))
    if true {
        v13 := []int{-9223372036854775808, -918}
        var v14 uint64 = 0
        _ = v13
        _ = v14
    } else {
        v2 = 127
        var v15 float64 = 0.0
        _ = v15
    }
    var v16 uint = 18446744073709551615
    var v17 uint8 = v3
    if (!false) {
        v16 += (v16 & v16)
        if (-2.25 > (1.5 + -2.25)) {
            v17 = v2
            v18 := []int{}
            _ = v18
        }
        var v19 int8 = 85
        _ = v19
    }
    switch 0 {
    case -293:
        v20 := GenS1{F0: (((v1 + "a") + ("\n\t" + ("\n\t" + ("\n\t" + v1)))) + (v1 + "")), F1: v17, F2: 1000}
        _ = v20
    }
    v21 := GenS2{F0: "foo bar"}
    v1 = "foo bar"
    if false {
        if (float64(0.0) > (float64(1e10) - float64(((0.0 + ((-2.25 + 0.0) / 3.14159)) / 1.5)))) {
            if false {
                v22 := GenS2{F0: v1}
                v1 = "foo bar"
                if ((1e10 * 3.14159) <= 1.5) {
                    if false {
                        var v23 uint = (v16 - v16)
                        v24 := GenS1{F0: (v1 + v1), F1: (v2 & (v17 | (v2 - (v3 >> (((v16 & 2)) & 63))))), F2: 1000}
                        _ = v23
                        _ = v24
                    }
                    var v25 int64 = 1
                    if (v1 <= ((v1 + v1) + v1)) {
                        v17 = v17
                        if false {
                            v17 = v3
                            v26 := []int{}
                            _ = v26
                        } else {
                            for fi27 := 0; fi27 < 9; fi27++ {
                                v3 = (v17 | 2)
                                _ = fi27
                            }
                            fmt.Println(v2, v2)
                        }
                    }
                    _ = v25
                } else {
                    var v28 int16 = 75
                    _ = v28
                }
                _ = v22
            } else {
                var v29 bool = (!(-791 > 42))
                for fi30 := 0; fi30 < 4; fi30++ {
                    v31 := []int{(fi30 % ((fi30) | 1)), fi30, 1, 32767}
                    v3 -= (v3 & (v17 << ((v16) & 63)))
                    _ = fi30
                    _ = v31
                }
                _ = v29
            }
        } else {
            v32 := GenS3{F0: 9223372036854775807}
            v33 := GenS1{F0: "hello", F1: (v2 * (^(v3))), F2: -269}
            _ = v32
            _ = v33
        }
        var v34 bool = false
        fmt.Println(v2, v1)
        _ = v34
    }
    v35 := GenS1{F0: "foo bar", F1: (v3 >> (((^(v16))) & 63)), F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v35
}
