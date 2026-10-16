package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint }
type GenS3 struct { F0 string; F1 uint16; F2 uint }
type GenS4 struct { F0 int }
func genFn0(p0 int32, p1 int16, p2 bool) bool {
    var v7 uint64 = 42
    fmt.Println(p0)
    v8 := []int{int(p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return false
}

func main() {
    v1 := make(map[string]string)
    var v2 int64 = -1
    var v3 int8 = 127
    v4 := []uint32{}
    v5 := make(map[uint32]uint32)
    var v6 float32 = 3.14159
    for fi9 := 0; fi9 < 12; fi9++ {
        if true {
            for fi10 := 0; fi10 < 10; fi10++ {
                v11 := GenS1{F0: v3, F1: ("a" + (("a" + "a") + ("hello" + "\n\t"))), F2: (fi10 + -9223372036854775808)}
                v12 := GenS2{F0: 548, F1: 1, F2: 1}
                if (128 < 255) {
                    if ((!(65535 < 0)) && (!(fi10 != (fi10 & 32768)))) {
                        if (true && false) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v13 uint64 = 260
                            _ = v13
                        } else {
                            v6 = float32((1.5 * float64(v6)))
                            fmt.Println(v2, fi9)
                        }
                        switch (fi9 ^ (fi9 % (((fi10 >> ((0) & 63))) | 1))) {
                        case 1:
                            switch (-(fi9)) {
                            case 1:
                                var v14 int = (fi10 + fi9)
                                _ = v14
                            case 128:
                                fmt.Println(v3, fi9, fi9)
                            }
                        }
                    }
                    var v15 uint = 843
                    _ = v15
                } else {
                    var v16 float64 = (float64(0.0) + (float64(v6) - (((float64(v6) - (float64(1.5) + 0.1)) - 1e10) * 0.0)))
                    _ = v16
                }
                _ = fi10
                _ = v11
                _ = v12
            }
            var v17 int32 = 1
            _ = v17
        } else {
            var v18 int64 = (v2 / ((9223372036854775807) | 1))
            _ = v18
        }
        _ = fi9
    }
    if true {
        v6 -= 0.1
        fmt.Println(v3)
    }
    var v19 uint64 = 18446744073709551615
    v6 += v6
    v20 := GenS3{F0: "", F1: 2, F2: 1}
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 uint = 255
        for fi23 := 0; fi23 < 15; fi23++ {
            v3 -= (v3 | (v3 % ((v3) | 1)))
            _ = fi23
        }
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
        _ = fi21
        _ = v22
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 float32 = v6
        _ = fi24
        _ = v25
    }
    v26 := GenS3{F0: ("a" + ("\n\t" + "")), F1: 0, F2: 596}
    switch ("" + "\n\t") {
    case "hello":
        for fi27 := 0; fi27 < 13; fi27++ {
            if (v19 <= v19) {
                v3 -= (v3 + v3)
                if true {
                    v19 += v19
                }
                if ((0 != 350) && (v19 == v19)) {
                    if ((("" + ("foo bar" + ("a" + "a"))) + "\n\t") > "") {
                        var v28 uint64 = (v19 * v19)
                        var v29 uint16 = 854
                        _ = v28
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 13; fi30++ {
                        var v31 int8 = v3
                        _ = fi30
                        _ = v31
                    }
                }
            } else {
                if true {
                    v2 += (v2 >> ((594) & 63))
                    v4 = append(v4, 655)
                } else {
                    _ = len(v1)
                }
                for fi32 := 0; fi32 < 2; fi32++ {
                    var v33 bool = true
                    _ = fi32
                    _ = v33
                }
            }
            _ = fi27
        }
    case "\n\t":
        fmt.Println(v19, v3, v19)
    }
    fmt.Println(v6, v2)
    v34 := v1["a"]
    v35 := GenS2{F0: 0, F1: 235, F2: 18446744073709551615}
    if ((v34 + "foo bar") <= v34) {
        v4 = append(v4, 2147483647)
        v36 := GenS4{F0: -9223372036854775808}
        switch 32768 {
        case 128:
            for fi37 := 0; fi37 < 1; fi37++ {
                var v38 string = v34
                v6 += v6
                _ = fi37
                _ = v38
            }
        case 127:
            v3 = (v3 + (v3 | -128))
        default:
            var v39 int32 = 0
            _ = v39
        }
        _ = v36
    } else {
        v6 = (v6 * v6)
        v5[687] = 2147483647
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
