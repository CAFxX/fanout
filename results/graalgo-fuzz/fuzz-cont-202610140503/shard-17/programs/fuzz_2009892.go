package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 int64 }
func genFn0() uint64 {
    var v6 uint = 4294967295
    v6 -= v6
    fmt.Println(v6)
    for fi7 := 0; fi7 < 4; fi7++ {
        v6 += 711
        v6 += 4294967295
        _ = fi7
    }
    fmt.Println(v6)
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 uint = 2
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]bool)
    var v4 int16 = -283
    var v5 uint64 = 64
    if false {
        v8 := GenS1{F0: v5, F1: int32(v1)}
        var v9 int8 = -1
        _ = v8
        _ = v9
    } else {
        v5 = (^(v5))
        if false {
            var v10 int16 = (v4 ^ 32767)
            for fi11 := 0; fi11 < 10; fi11++ {
                v5 = (v5 * v5)
                _ = fi11
            }
            v10 += v10
            _ = v10
        } else {
            var v12 float32 = -2.25
            if (genFn0() > (^(v5))) {
                v12 = (float32(0.1) / 1.0)
                if (1 >= 42) {
                    for fi13 := 0; fi13 < 2; fi13++ {
                        switch ((((("\n\t" + "a") + ("a" + "")) + ("foo bar" + "foo bar")) + "foo bar") + "hello") {
                        case "":
                            for fi14 := 0; fi14 < 7; fi14++ {
                                v15 := GenS2{F0: (v4 & (^(v4))), F1: 1}
                                v12 = v12
                                v1 = (v1 & v1)
                                _ = fi14
                                _ = v15
                            }
                        default:
                            for fi16 := 0; fi16 < 11; fi16++ {
                                v17 := GenS2{F0: (v4 - v4), F1: 9223372036854775807}
                                v18 := GenS1{F0: genFn0(), F1: 0}
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                        }
                        for fi19 := 0; fi19 < 7; fi19++ {
                            v20 := GenS3{F0: (v1 * (v1 % ((v1) | 1))), F1: 120}
                            _ = fi19
                            _ = v20
                        }
                        _ = fi13
                    }
                    for fi21 := 0; fi21 < 7; fi21++ {
                        v22 := []int{fi21, (fi21 << ((186) & 63))}
                        _ = fi21
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 19; fi23++ {
                        if true {
                            switch 0 {
                            case 128:
                                v24 := GenS2{F0: (v4 / (((v4 + 42)) | 1)), F1: 9223372036854775807}
                                _ = v24
                            case 42:
                                v4 = -32768
                            default:
                                switch fi23 {
                                case -2147483648:
                                    var v25 uint32 = 443
                                    _ = v25
                                }
                            }
                        }
                        for fi26 := 0; fi26 < 3; fi26++ {
                            v5 -= (v5 >> ((v1) & 63))
                            _ = fi26
                        }
                        for fi27 := 0; fi27 < 3; fi27++ {
                            if false {
                                switch -837 {
                                case -933:
                                    v1 = v1
                                    v28 := GenS1{F0: v5, F1: 42}
                                    _ = v28
                                }
                                var v29 uint = 112
                                v29 = 64
                                _ = v29
                            }
                            var v30 uint = 18446744073709551615
                            _ = fi27
                            _ = v30
                        }
                        _ = fi23
                    }
                } else {
                    v31 := []int{2147483647}
                    for fi32 := 0; fi32 < 14; fi32++ {
                        _ = len(v31)
                        _ = fi32
                    }
                    _ = v31
                }
            } else {
                v33 := GenS2{F0: v4, F1: 45}
                _ = v33
            }
            _ = v12
        }
    }
    v5 += genFn0()
    v4 += (v4 * v4)
    v34 := []int{}
    if (!(-342 != 991)) {
        switch (("hello" + "a") + "hello") {
        case "":
            fmt.Println(v4, v5, v5)
        }
        var v35 int16 = (v4 & (v4 >> ((18446744073709551615) & 63)))
        _ = v35
    } else {
        if true {
            if true {
                var v36 uint32 = 2147483647
                _ = v36
            }
            v5 = (v5 << ((v1) & 63))
        }
        var v37 float32 = (3.14159 * -2.25)
        _ = v37
    }
    v38 := GenS2{F0: 1000, F1: -795}
    for fi39 := 0; fi39 < 16; fi39++ {
        var v40 uint64 = genFn0()
        v40 = (^(v5))
        _ = fi39
        _ = v40
    }
    v1 = v1
    var v41 uint = v1
    if (v5 < 913) {
        var v42 int64 = -235
        v3[65535] = true
        if (v42 != v42) {
            fmt.Println(v5, v5)
            switch "" {
            case "a":
                var v43 int = 1
                _ = v43
            case "foo bar":
                var v44 int16 = 32767
                _ = v44
            default:
                if ((872 > 4294967295) && true) {
                    var v45 uint32 = 0
                    _ = v45
                } else {
                    v1 = v41
                    var v46 int8 = -128
                    _ = v46
                }
            }
        } else {
            delete(v3, 178)
            v47 := GenS4{F0: (v42 & -1)}
            _ = v47
        }
        _ = v42
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v4
    _ = v41
    _ = v5
}
