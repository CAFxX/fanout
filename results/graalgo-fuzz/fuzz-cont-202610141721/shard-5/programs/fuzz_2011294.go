package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint8 }
func genFn0() uint64 {
    var v4 uint32 = 290
    for fi5 := 0; fi5 < 14; fi5++ {
        v4 = v4
        var v6 bool = true
        switch fi5 {
        case 128:
            for fi7 := 0; fi7 < 16; fi7++ {
                v6 = (!v6)
                fmt.Println(v6)
                _ = fi7
            }
        case 633:
            if ((float64((3.14159 + 1e10)) >= ((1.5 + 3.14159) * float64(-2.25))) || false) {
                v4 = v4
                var v8 uint8 = 237
                var v9 int32 = 0
                _ = v8
                _ = v9
            } else {
                var v10 int64 = 42
                for fi11 := 0; fi11 < 8; fi11++ {
                    fmt.Println(v10, v4)
                    _ = fi11
                }
                var v12 float64 = float64(float32((0.1 - 0.0)))
                _ = v10
                _ = v12
            }
        case 32767:
            fmt.Println(fi5)
        }
        _ = fi5
        _ = v6
    }
    _ = v4
    return 4294967295
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 2
    var v3 uint16 = (v2 / ((v2) | 1))
    v13 := GenS1{F0: ((("a" + ("\n\t" + "foo bar")) + "hello") + "hello"), F1: (^(v3)), F2: 0}
    if false {
        v14 := []int{-2147483648, 0}
        _ = v14
    }
    v15 := []int{32768, 32768, -2147483648}
    for fi16 := 0; fi16 < 15; fi16++ {
        v3 += v3
        v2 += v3
        _ = fi16
    }
    v1[""] = 0
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        _ = len(v15)
        _ = fi17
    }
    if true {
        if false {
            var v18 int16 = 894
            if (float32((float64(3.14159) * 1e10)) >= float32(float64(float32(-2.25)))) {
                v18 = (v18 + v18)
                var v19 int64 = -9223372036854775808
                v3 -= v3
                _ = v19
            } else {
                var v20 int = 1
                _ = v20
            }
            if false {
                if false {
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    fmt.Println(v18, v3, v3)
                }
            }
            _ = v18
        } else {
            v21 := GenS1{F0: "\n\t", F1: (v2 << ((255) & 63)), F2: 1}
            v2 = v2
            if false {
                v1["\n\t"] = 42
            } else {
                if ("hello" == ("a" + ("" + ("" + ("a" + (("foo bar" + "foo bar") + "\n\t")))))) {
                    if true {
                        v22 := GenS1{F0: "\n\t", F1: (v2 / (((v3 * (^(v3)))) | 1)), F2: 128}
                        v2 = (v3 % ((v3) | 1))
                        for fi23 := 0; fi23 < 16; fi23++ {
                            for fi24 := 0; fi24 < 10; fi24++ {
                                v3 = 42
                                var v25 int = fi24
                                v26 := GenS1{F0: "hello", F1: v2, F2: 255}
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                            var v27 int = (fi23 | -764)
                            var v28 int64 = 431
                            _ = fi23
                            _ = v27
                            _ = v28
                        }
                        _ = v22
                    } else {
                        fmt.Println(v3, v2)
                        switch ((("\n\t" + "foo bar") + "hello") + "a") {
                        case "foo bar":
                            for fi29 := 0; fi29 < 1; fi29++ {
                                var v30 int32 = -981
                                _ = fi29
                                _ = v30
                            }
                        }
                        if false {
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                        }
                    }
                } else {
                    if false {
                        v2 = v3
                    }
                    fmt.Println(v2, v2)
                    v2 -= v3
                }
                v2 = 32767
                var v31 int = 795
                _ = v31
            }
            _ = v21
        }
    }
    var v32 int = 1
    fmt.Println(v3, v2)
    _ = len(v1)
    v33 := GenS1{F0: "", F1: v2, F2: 255}
    v15 = append(v15, -2147483648)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v33
}
