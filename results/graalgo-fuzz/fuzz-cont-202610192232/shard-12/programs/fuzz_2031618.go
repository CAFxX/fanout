package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int16; F2 string }
type GenS3 struct { F0 uint16; F1 int64 }
func genFn0(p0 int16, p1 uint16) bool {
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 int8 = 100
        _ = fi7
        _ = v8
    }
    var v9 bool = true
    _ = p0
    _ = p1
    _ = v9
    return (647 <= 1)
}

func main() {
    var v1 uint16 = 349
    v2 := []int{9223372036854775807, -1, 9223372036854775807, 303}
    var v3 string = ("hello" + "\n\t")
    v4 := []string{v3, (v3 + v3), v3, v3}
    v5 := []uint32{0, 1, 1, 225}
    v6 := []string{(v3 + v3), "a", (v3 + (((v3 + (v3 + v3)) + ((("" + "hello") + "foo bar") + "")) + (((("\n\t" + v3) + v3) + ((v3 + v3) + v3)) + "hello"))), (v3 + v3)}
    switch v3 {
    case "a":
        fmt.Println(v3, v3)
    default:
        if (1 == 104) {
            v10 := GenS1{F0: 228}
            if true {
                var v11 int8 = 0
                if ((^(v1)) == v1) {
                    if false {
                        var v12 uint32 = 1
                        if (!true) {
                            if true {
                                v13 := GenS1{F0: 18446744073709551615}
                                fmt.Println(v11, v3)
                                switch 128 {
                                case 2147483647:
                                    var v14 uint32 = v12
                                    _ = v14
                                default:
                                    var v15 int32 = -180
                                    _ = v15
                                }
                                _ = v13
                            } else {
                                v16 := GenS1{F0: 255}
                                var v17 uint8 = 128
                                _ = v16
                                _ = v17
                            }
                            var v18 int64 = 643
                            if (v3 == "") {
                                var v19 uint16 = (v1 << ((42) & 63))
                                var v20 bool = true
                                _ = v19
                                _ = v20
                            }
                            _ = v18
                        }
                        v21 := GenS1{F0: 4294967295}
                        _ = v12
                        _ = v21
                    } else {
                        fmt.Println(v3, v3, v11)
                        v22 := GenS1{F0: 902}
                        _ = len(v5)
                        _ = v22
                    }
                    fmt.Println(v11)
                    v23 := GenS2{F0: 613, F1: 32767, F2: (("foo bar" + (v3 + v3)) + "foo bar")}
                    _ = v23
                } else {
                    switch v3 {
                    case "hello":
                        if (-118 == -1) {
                            if true {
                                v24 := GenS3{F0: (v1 % (((v1 / ((1) | 1))) | 1)), F1: -1}
                                fmt.Println(v11, v1, v11)
                                v1 = (v1 + (^(v1)))
                                _ = v24
                            }
                            v3 = (v3 + ("foo bar" + v3))
                        } else {
                            var v25 float64 = -2.25
                            _ = v25
                        }
                    }
                }
                _ = v11
            }
            _ = v10
        } else {
            v26 := GenS1{F0: 147}
            _ = v26
        }
    }
    v3 = "a"
    var v27 uint16 = 42
    if false {
        var v28 int8 = -1
        if false {
            if (-529 > 32767) {
                _ = len(v5)
                v29 := GenS3{F0: 750, F1: 9223372036854775807}
                _ = v29
            } else {
                var v30 int32 = 165
                var v31 int = 748
                _ = v30
                _ = v31
            }
        } else {
            if false {
                _ = len(v5)
            }
        }
        _ = v28
    } else {
        v27 += v27
        if len(v2) > 0 {
            v2[0] = -429
        }
    }
    var v32 uint64 = 2
    if len(v4) > 0 {
        v4[0] = v3
    }
    var v33 int32 = 1
    var v34 string = "foo bar"
    var v35 bool = false
    _ = len(v6)
    for fi36 := 0; fi36 < 7; fi36++ {
        var v37 int16 = 0
        var v38 uint8 = 255
        var v39 float64 = 0.1
        _ = fi36
        _ = v37
        _ = v38
        _ = v39
    }
    fmt.Println(v3, v1, v1)
    switch v34 {
    case "hello":
        v40 := GenS2{F0: 207, F1: -32768, F2: (v3 + (("\n\t" + ("" + (v3 + "\n\t"))) + v3))}
        _ = v40
    case "":
        switch ("\n\t" + v34) {
        case "":
            var v41 int8 = 4
            _ = v41
        case "hello":
            var v42 uint16 = v1
            _ = v42
        }
    default:
        var v43 uint8 = 1
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
