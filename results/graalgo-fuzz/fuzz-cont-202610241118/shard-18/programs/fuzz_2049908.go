package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint; F2 int32 }
type GenS3 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS4 struct { F0 int; F1 int; F2 int64 }
type GenS5 struct { F0 int32; F1 bool }

func main() {
    var v1 float64 = (float64((float32(((float64(-2.25) + -2.25) / -2.25)) * 1.5)) + -2.25)
    var v2 uint16 = 0
    var v3 string = "hello"
    v4 := []string{(("\n\t" + v3) + v3), ((((v3 + "") + "a") + v3) + v3)}
    var v5 uint = 18446744073709551615
    var v6 float64 = 0.0
    v1 = v1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) > 0 {
        v4[0] = ""
    }
    fmt.Println(v1, v5, v1)
    v7 := GenS1{F0: 1}
    if (0.0 > 0.1) {
        v8 := GenS2{F0: v3, F1: 42, F2: -1}
        if (51 == 233) {
            var v9 float32 = -2.25
            _ = v9
        } else {
            v10 := GenS2{F0: (v3 + ((((v3 + v3) + v3) + "a") + "a")), F1: (v5 / (((v5 ^ (^(v5)))) | 1)), F2: 835}
            v11 := GenS2{F0: (((((((((v3 + ("foo bar" + v3)) + v3) + "\n\t") + ("hello" + v3)) + ("a" + "a")) + ((v3 + "a") + v3)) + (v3 + "a")) + v3) + ((v3 + v3) + "\n\t")), F1: (v5 + (v5 * (v5 - 255))), F2: 42}
            if ((-2.25 + 3.14159) < -2.25) {
                for fi12 := 0; fi12 < 3; fi12++ {
                    fmt.Println(v1, v1, v1)
                    for fi13 := 0; fi13 < 3; fi13++ {
                        var v14 string = "a"
                        fmt.Println(v3, fi12)
                        _ = fi13
                        _ = v14
                    }
                    _ = fi12
                }
                v3 = ("hello" + v3)
            } else {
                if (42 < 32767) {
                    if (!(-1 >= 9223372036854775807)) {
                        v5 = 477
                    } else {
                        for fi15 := 0; fi15 < 2; fi15++ {
                            for fi16 := 0; fi16 < 13; fi16++ {
                                v6 += (((v1 * v1) - -2.25) + v6)
                                if ((991 <= 0) || true) {
                                    var v17 int8 = -128
                                    var v18 uint = (v5 / ((v5) | 1))
                                    _ = v17
                                    _ = v18
                                } else {
                                    v5 -= 157
                                    if false {
                                        if (!true) {
                                            fmt.Println(fi15)
                                        }
                                        v19 := GenS2{F0: "hello", F1: v5, F2: 0}
                                        fmt.Println(v5)
                                        if true {
                                            if len(v4) > 0 {
                                                v4[0] = (v3 + (v3 + (((v3 + v3) + "a") + "a")))
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = v3
                                            }
                                        } else {
                                            v6 -= (v1 - (((v1 * float64(float32(float64(1e10)))) * v6) - float64(1.5)))
                                            var v20 uint = v5
                                            v3 = v3
                                            _ = v20
                                        }
                                        _ = v19
                                    } else {
                                        var v21 uint64 = 2
                                        var v22 int8 = -1
                                        v23 := GenS3{F0: -2147483648, F1: v21, F2: 0}
                                        _ = v21
                                        _ = v22
                                        _ = v23
                                    }
                                    if (v3 == v3) {
                                        v6 = (float64(0.0) * 0.0)
                                    }
                                }
                                _ = fi16
                            }
                            v24 := GenS1{F0: 42}
                            fmt.Println(v2)
                            _ = fi15
                            _ = v24
                        }
                        for fi25 := 0; fi25 < 5; fi25++ {
                            if len(v4) > 0 {
                                v4[0] = v3
                            }
                            _ = fi25
                        }
                    }
                }
            }
            _ = v10
            _ = v11
        }
        var v26 bool = true
        _ = v26
        _ = v8
    } else {
        v2 = (v2 % (((v2 - v2)) | 1))
        if false {
            var v27 uint32 = 835
            v28 := GenS4{F0: 4294967295, F1: 128, F2: 0}
            _ = v27
            _ = v28
        }
    }
    v29 := GenS1{F0: 427}
    v30 := GenS5{F0: -1, F1: (!((!(127 >= 127)) || (293 <= 424)))}
    v31 := GenS2{F0: v3, F1: v5, F2: -2147483648}
    v32 := GenS1{F0: 0}
    v33 := GenS4{F0: 32767, F1: 177, F2: 866}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
