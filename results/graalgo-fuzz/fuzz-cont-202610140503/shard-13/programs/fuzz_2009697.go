package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 uint64 }
func genFn0() int8 {
    var v6 uint64 = 4294967295
    switch 2147483647 {
    case 409:
        for fi7 := 0; fi7 < 20; fi7++ {
            v6 -= (v6 | (v6 % (((v6 & (v6 | (v6 % ((1) | 1))))) | 1)))
            _ = fi7
        }
    case 128:
        v8 := GenS1{F0: 0}
        _ = v8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        if false {
            v6 += (v6 | (v6 / ((v6) | 1)))
        } else {
            fmt.Println(v6, v6)
            v10 := []int{}
            v11 := GenS1{F0: fi9}
            _ = v10
            _ = v11
        }
        _ = fi9
    }
    fmt.Println(v6)
    _ = v6
    return 127
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]string)
    var v3 bool = (1 < 18446744073709551615)
    var v4 uint64 = 2
    var v5 int8 = 100
    var v12 float64 = (0.1 * 3.14159)
    var v13 int16 = -32768
    fmt.Println(v12)
    var v14 int16 = -32768
    v12 += v12
    if (!v3) {
        var v15 int32 = -1
        if false {
            if (-901 != -798) {
                v16 := GenS1{F0: 128}
                _ = v16
            } else {
                var v17 string = (v1 + (((v1 + (v1 + v1)) + "a") + (((("\n\t" + ((v1 + "a") + (v1 + "\n\t"))) + v1) + "a") + v1)))
                _ = v17
            }
            fmt.Println(v5, v15, v14)
        } else {
            for fi18 := 0; fi18 < 11; fi18++ {
                var v19 int8 = int8(v4)
                v20 := v2[(v1 + v1)]
                switch v20 {
                case "\n\t":
                    v4 = (v4 / ((v4) | 1))
                case "foo bar":
                    _ = len(v2)
                default:
                    switch v20 {
                    case "hello":
                        var v21 float32 = 1e10
                        _ = v21
                    case "\n\t":
                        v22 := v2[v20]
                        _ = v22
                    }
                }
                _ = fi18
                _ = v19
                _ = v20
            }
            v14 += 1000
        }
        fmt.Println(v5, v13)
        _ = v15
    }
    if (false || false) {
        if v3 {
            for fi23 := 0; fi23 < 3; fi23++ {
                v13 = int16(v5)
                _ = fi23
            }
            fmt.Println(v1)
            v13 = (v14 % ((v13) | 1))
        } else {
            v24 := GenS1{F0: 4294967295}
            _ = v24
        }
    } else {
        if (176 > 186) {
            fmt.Println(v4)
            var v25 float64 = ((float64(1e10) / 1e10) * float64(((1.5 * 3.14159) * 1.5)))
            if true {
                var v26 uint8 = 128
                if false {
                    v27 := v2[v1]
                    _ = v27
                } else {
                    switch (v1 + "") {
                    case "a":
                        v28 := GenS2{F0: (v4 & 42)}
                        _ = v28
                    case "hello":
                        switch 127 {
                        case 2147483647:
                            if (128 < (v26 ^ v26)) {
                                if v3 {
                                    if (3.14159 < (0.1 - v12)) {
                                        fmt.Println(v5, v14, v5)
                                        if v3 {
                                            for fi29 := 0; fi29 < 1; fi29++ {
                                                var v30 float32 = 0.0
                                                var v31 int = (fi29 ^ int(v5))
                                                v32 := GenS3{F0: v26, F1: v4}
                                                _ = fi29
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                            }
                                        }
                                    }
                                } else {
                                    fmt.Println(v3)
                                }
                            }
                        }
                    }
                    if v3 {
                        var v33 uint16 = 65535
                        _ = len(v2)
                        _ = v33
                    }
                }
                _ = v26
            }
            _ = v25
        }
        v13 += (v14 % ((v13) | 1))
    }
    v34 := GenS2{F0: v4}
    v2[(v1 + v1)] = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
