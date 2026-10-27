package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint16; F2 int }
type GenS3 struct { F0 string }

func main() {
    var v1 string = ""
    var v2 uint = 255
    v3 := []int{-310}
    var v4 int = 32767
    var v5 int = 128
    var v6 uint16 = 65535
    switch 128 {
    case 2147483647:
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := GenS1{F0: 474}
            if (((3.14159 / (0.1 / 0.1)) <= 1.5) && (1 < 100)) {
                v1 = v1
                for fi9 := 0; fi9 < 16; fi9++ {
                    if false {
                        v10 := GenS1{F0: 42}
                        _ = v10
                    }
                    if (-9223372036854775808 != v5) {
                        if (false && (int64(v2) != -1)) {
                            if (v2 < v2) {
                                var v11 uint16 = v6
                                _ = v11
                            }
                            switch (-(v5)) {
                            case 32767:
                                v12 := GenS1{F0: 2147483647}
                                _ = v12
                            }
                        } else {
                            fmt.Println(v1)
                            _ = len(v3)
                            v5 += fi7
                            if len(v3) > 0 {
                                v3[0] = -9223372036854775808
                            }
                        }
                        var v13 bool = false
                        v13 = v13
                        _ = v13
                    }
                    _ = fi9
                }
                var v14 int = v4
                _ = v14
            } else {
                var v15 bool = (true && false)
                for fi16 := 0; fi16 < 11; fi16++ {
                    if v15 {
                        if v15 {
                            var v17 uint16 = (v6 % ((74) | 1))
                            if ((597 < 4294967295) || v15) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if (42 >= -32768) {
                                    v6 = (^(v6))
                                    if (!(true || ("\n\t" >= "foo bar"))) {
                                        v17 += (v6 & (v17 / ((v6) | 1)))
                                        var v18 uint32 = 159
                                        _ = v18
                                    } else {
                                        v3 = append(v3, (fi16 | fi16))
                                    }
                                    var v19 int32 = -219
                                    _ = v19
                                }
                                v20 := GenS2{F0: 0, F1: (v6 << ((65535) & 63)), F2: fi16}
                                _ = v20
                            } else {
                                var v21 int32 = 2147483647
                                _ = v21
                            }
                            var v22 int32 = 440
                            v23 := GenS2{F0: -2147483648, F1: 255, F2: 128}
                            var v24 uint = uint(v6)
                            _ = v17
                            _ = v22
                            _ = v23
                            _ = v24
                        } else {
                            v2 += v2
                            v15 = (((-2.25 + 3.14159) + 1.5) <= (3.14159 / (float32(1.5) - 0.1)))
                            var v25 float32 = (0.0 - float32((-2.25 - 1.5)))
                            _ = v25
                        }
                    }
                    var v26 float64 = 3.14159
                    if ((-32768 >= -1) || (1 >= -955)) {
                        var v27 float32 = 1.5
                        _ = v27
                    } else {
                        if len(v3) > 0 {
                            v3[0] = fi16
                        }
                        v28 := GenS3{F0: ((v1 + ("foo bar" + ((v1 + v1) + ("foo bar" + v1)))) + ((v1 + v1) + "a"))}
                        var v29 uint64 = 2
                        _ = v28
                        _ = v29
                    }
                    _ = fi16
                    _ = v26
                }
                _ = v15
            }
            _ = fi7
            _ = v8
        }
    }
    v2 -= v2
    var v30 int32 = -2147483648
    v2 = (^(v2))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v4)
    var v31 uint16 = v6
    if len(v3) > 0 {
        v3[0] = (-(v5))
    }
    _ = len(v3)
    _ = len(v3)
    var v32 int8 = 100
    if len(v3) > 0 {
        v3[0] = -658
    }
    var v33 int64 = -9223372036854775808
    v4 = v5
    v4 -= (v4 ^ v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
