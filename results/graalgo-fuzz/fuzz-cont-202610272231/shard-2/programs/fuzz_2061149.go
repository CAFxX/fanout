package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int; F1 bool }
func genFn0(p0 string) uint {
    v7 := GenS1{F0: 311, F1: (("foo bar" + "hello") + p0)}
    if (227 != 2147483647) {
        var v8 int64 = -9223372036854775808
        v9 := GenS1{F0: 1, F1: p0}
        _ = v8
        _ = v9
    } else {
        switch p0 {
        case "hello":
            var v10 float64 = 0.1
            _ = v10
        default:
            v11 := GenS1{F0: 0, F1: "a"}
            _ = v11
        }
        v12 := GenS2{F0: 42, F1: 715}
        _ = v12
    }
    var v13 int8 = 1
    v14 := []int{627, 848}
    _ = p0
    _ = v13
    _ = v14
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 int, p1 int) int8 {
    if ((2 > 18446744073709551615) && true) {
        if (32767 < -897) {
            v15 := []int{42, p0, (p0 & (p1 << ((1) & 63)))}
            if (float64((3.14159 - float32(3.14159))) <= 0.0) {
                switch (p1 % (((p1 - p0)) | 1)) {
                case 134:
                    if (false && ((!((4294967295 != 65535) && (568 < 4294967295))) || ((9223372036854775807 == 1) || true))) {
                        if ((false && true) && true) {
                            switch "" {
                            case "a":
                                var v16 uint = 2
                                _ = v16
                            case "":
                                for fi17 := 0; fi17 < 16; fi17++ {
                                    switch (fi17 + fi17) {
                                    case -1:
                                        if len(v15) > 0 {
                                            v15[0] = (p0 % ((p1) | 1))
                                        }
                                    case -9223372036854775808:
                                        var v18 int32 = 0
                                        _ = v18
                                    default:
                                        for fi19 := 0; fi19 < 8; fi19++ {
                                            if false {
                                                var v20 float32 = 1e10
                                                var v21 uint16 = 255
                                                _ = v20
                                                _ = v21
                                            } else {
                                                _ = len(v15)
                                                var v22 bool = true
                                                var v23 int8 = -128
                                                fmt.Println(p1, v22, p0)
                                                v23 += 51
                                                fmt.Println(v22, fi19, v22)
                                                _ = v22
                                                _ = v23
                                            }
                                            _ = fi19
                                        }
                                    }
                                    _ = fi17
                                }
                            }
                        } else {
                            for fi24 := 0; fi24 < 14; fi24++ {
                                p1 -= -498
                                for fi25 := 0; fi25 < 15; fi25++ {
                                    p1 -= fi24
                                    _ = fi25
                                }
                                v26 := GenS2{F0: 0, F1: 291}
                                _ = fi24
                                _ = v26
                            }
                            v27 := GenS3{F0: p0, F1: (p0 <= p1)}
                            _ = v27
                        }
                    } else {
                        v15 = append(v15, (p0 << ((genFn0("")) & 63)))
                        p0 = p1
                    }
                default:
                    v28 := GenS2{F0: 1, F1: -356}
                    _ = v28
                }
                p1 = p0
            } else {
                for fi29 := 0; fi29 < 5; fi29++ {
                    for fi30 := 0; fi30 < 8; fi30++ {
                        _ = len(v15)
                        _ = fi30
                    }
                    for fi31 := 0; fi31 < 1; fi31++ {
                        fmt.Println(p1, p1)
                        if false {
                            var v32 float64 = float64((1e10 / (1.5 + float32(0.1))))
                            fmt.Println(p0)
                            _ = v32
                        } else {
                            p1 = fi31
                            var v33 int16 = -463
                            v34 := GenS3{F0: (fi31 % (((-(p1))) | 1)), F1: true}
                            _ = v33
                            _ = v34
                        }
                        p1 += 42
                        _ = fi31
                    }
                    switch (-(fi29)) {
                    case -9223372036854775808:
                        v35 := GenS3{F0: (fi29 + (-(p0))), F1: false}
                        _ = v35
                    }
                    _ = fi29
                }
                fmt.Println(p1, p0)
                switch (("hello" + ("\n\t" + ("a" + "hello"))) + "\n\t") {
                case "":
                    var v36 uint = 850
                    _ = v36
                }
            }
            v37 := GenS2{F0: 9223372036854775807, F1: 32767}
            _ = v15
            _ = v37
        } else {
            v38 := []int{p1, (^(p1))}
            p0 += p0
            _ = v38
        }
        p0 += 2147483647
        var v39 float32 = 3.14159
        _ = v39
    } else {
        var v40 bool = false
        for fi41 := 0; fi41 < 18; fi41++ {
            var v42 float64 = float64(((3.14159 * 0.0) * (1e10 / 1e10)))
            for fi43 := 0; fi43 < 11; fi43++ {
                var v44 int64 = 18
                fmt.Println(p0, p0)
                var v45 bool = false
                _ = fi43
                _ = v44
                _ = v45
            }
            _ = fi41
            _ = v42
        }
        v46 := []int{p1, p0, 32768}
        _ = v40
        _ = v46
    }
    p1 -= (p0 >> ((42) & 63))
    p1 += (p1 - p0)
    _ = p0
    _ = p1
    return -128
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[int]bool)
    var v3 int32 = 42
    v4 := make(map[string]bool)
    v5 := make(map[uint32]string)
    v6 := []string{"hello", ("" + "")}
    switch 1 {
    case 779:
        for fi47 := 0; fi47 < 16; fi47++ {
            v5[65535] = "\n\t"
            if true {
                var v48 uint = 65535
                var v49 float32 = 3.14159
                for fi50 := 0; fi50 < 17; fi50++ {
                    var v51 int8 = -108
                    _ = fi50
                    _ = v51
                }
                _ = v48
                _ = v49
            } else {
                var v52 float64 = 1e10
                v53 := GenS2{F0: 358, F1: (v1 | v1)}
                var v54 uint64 = 18446744073709551615
                _ = v52
                _ = v53
                _ = v54
            }
            _ = fi47
        }
    case 0:
        fmt.Println(v3, v1)
    }
    if len(v6) > 0 {
        v6[0] = "a"
    }
    switch "" {
    case "foo bar":
        v3 += -1
    case "hello":
        if len(v6) > 0 {
            v6[0] = "a"
        }
    }
    v1 = v1
    switch -1 {
    case -9223372036854775808:
        if true {
            var v55 uint32 = 2147483647
            for fi56 := 0; fi56 < 11; fi56++ {
                if false {
                    v3 -= -70
                } else {
                    v1 = 32767
                }
                v55 = 0
                for fi57 := 0; fi57 < 2; fi57++ {
                    var v58 float64 = -2.25
                    if len(v6) > 0 {
                        v6[0] = "a"
                    }
                    _ = fi57
                    _ = v58
                }
                _ = fi56
            }
            _ = v55
        }
    default:
        v3 += (v3 % (((v3 - (v3 | v3))) | 1))
    }
    if len(v6) > 0 {
        v6[0] = ((("a" + "hello") + "a") + "a")
    }
    _ = len(v4)
    v1 += (v1 & v1)
    v3 += 0
    v3 -= v3
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
