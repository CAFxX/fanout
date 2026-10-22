package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int; F2 uint }
func genFn0() int {
    var v4 uint32 = 2
    v4 = v4
    switch "\n\t" {
    case "":
        v4 = 65535
    case "a":
        v4 -= 2
    case "foo bar":
        for fi5 := 0; fi5 < 4; fi5++ {
            switch (fi5 & fi5) {
            case 708:
                v4 = (^(v4))
            default:
                for fi6 := 0; fi6 < 13; fi6++ {
                    if true {
                        fmt.Println(v4, fi5, fi6)
                        v4 = (v4 + (^(v4)))
                    } else {
                        if true {
                            v7 := []int{(fi5 | (fi6 >> ((42) & 63))), fi5}
                            _ = v7
                        }
                        var v8 int = fi5
                        fmt.Println(fi5, v4, v4)
                        var v9 float64 = float64(0.1)
                        _ = v8
                        _ = v9
                    }
                    _ = fi6
                }
            }
            if true {
                if false {
                    fmt.Println(fi5)
                    if true {
                        v4 -= (v4 * 611)
                    } else {
                        if (-517 > 32767) {
                            fmt.Println(fi5, v4)
                            v10 := GenS1{F0: v4}
                            var v11 uint64 = 4294967295
                            _ = v10
                            _ = v11
                        } else {
                            v12 := []int{(fi5 + fi5), (fi5 ^ int(v4))}
                            v4 = v4
                            _ = v12
                        }
                        if true {
                            var v13 string = "\n\t"
                            for fi14 := 0; fi14 < 18; fi14++ {
                                if false {
                                    var v15 int64 = 1
                                    _ = v15
                                } else {
                                    if (0.0 >= 1e10) {
                                        var v16 bool = (false || (true || true))
                                        _ = v16
                                    }
                                    v17 := []int{fi14, fi14, (fi14 / ((fi14) | 1)), fi5}
                                    _ = v17
                                }
                                v4 -= v4
                                v18 := []int{fi5, fi5}
                                _ = fi14
                                _ = v18
                            }
                            _ = v13
                        } else {
                            var v19 int64 = -9223372036854775808
                            switch ("a" + ("foo bar" + "a")) {
                            case "":
                                switch ("\n\t" + "hello") {
                                case "hello":
                                    var v20 string = "foo bar"
                                    _ = v20
                                case "\n\t":
                                    v21 := GenS1{F0: v4}
                                    _ = v21
                                }
                            case "a":
                                v4 = (^(v4))
                            default:
                                v4 -= (v4 ^ (^(v4)))
                            }
                            _ = v19
                        }
                        for fi22 := 0; fi22 < 14; fi22++ {
                            v23 := GenS1{F0: (v4 << ((4294967295) & 63))}
                            _ = fi22
                            _ = v23
                        }
                    }
                    var v24 int = fi5
                    _ = v24
                }
                var v25 string = "hello"
                if false {
                    fmt.Println(v4)
                    v26 := GenS1{F0: (^(v4))}
                    _ = v26
                } else {
                    v4 -= v4
                    v27 := []int{(fi5 * -365)}
                    if (true && (1 >= 851)) {
                        if (4294967295 == 572) {
                            if len(v27) > 0 {
                                v27[0] = 32767
                            }
                        } else {
                            if true {
                                var v28 uint64 = 2
                                fmt.Println(v28, v28)
                                for fi29 := 0; fi29 < 11; fi29++ {
                                    v25 = "hello"
                                    _ = fi29
                                }
                                _ = v28
                            }
                            var v30 float32 = -2.25
                            _ = v30
                        }
                    } else {
                        if len(v27) > 0 {
                            v27[0] = (^(fi5))
                        }
                        v25 = (((v25 + ("a" + v25)) + v25) + v25)
                        v4 = (v4 % (((v4 ^ 4294967295)) | 1))
                    }
                    _ = v27
                }
                _ = v25
            }
            fmt.Println(fi5, fi5)
            _ = fi5
        }
    }
    _ = v4
    return 4294967295
}
func genFn1(p0 int) int8 {
    p0 = 127
    var v31 int8 = 117
    p0 += (^(p0))
    _ = p0
    _ = v31
    return 127
}

func main() {
    var v1 uint32 = 2147483647
    var v2 int8 = -50
    v3 := []int{132, -559, -67, 357}
    v32 := GenS2{F0: 1, F1: -9223372036854775808, F2: 42}
    var v33 uint32 = v1
    v3 = append(v3, int(v2))
    var v34 uint8 = 127
    fmt.Println(v1, v34)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v35 := GenS1{F0: v33}
    if len(v3) > 0 {
        v3[0] = 1
    }
    if (!(((42 >= -303) || false) && true)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        var v36 int32 = 42
        fmt.Println(v36)
        if len(v3) > 0 {
            v3[0] = 66
        }
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}
