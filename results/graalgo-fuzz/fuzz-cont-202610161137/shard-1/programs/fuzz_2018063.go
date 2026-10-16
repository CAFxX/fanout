package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int64; F1 string; F2 uint32 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint8; F1 int16 }
func genFn0() string {
    var v6 int = 32767
    var v7 uint8 = 1
    v7 = (^(v7))
    _ = v6
    _ = v7
    return ("" + "foo bar")
}
func genFn1(p0 uint64, p1 uint32, p2 uint16) string {
    var v8 uint8 = 0
    v9 := []int{316}
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return "a"
}

func main() {
    var v1 bool = true
    var v2 string = ("foo bar" + "\n\t")
    var v3 int32 = -2147483648
    var v4 string = ""
    v5 := []uint32{2147483647, 1, 1, 437}
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        v2 = ""
        var v11 float32 = float32(float64(1e10))
        if v1 {
            v12 := GenS1{F0: 0, F1: (v3 | (v3 * -1))}
            v13 := GenS2{F0: -1, F1: v4, F2: 42}
            _ = v12
            _ = v13
        } else {
            for fi14 := 0; fi14 < 17; fi14++ {
                var v15 uint64 = 169
                if (v15 == 733) {
                    var v16 uint64 = 1
                    switch v4 {
                    case "a":
                        var v17 bool = (fi14 <= int(v3))
                        _ = v17
                    case "foo bar":
                        var v18 float32 = (-2.25 + (0.1 - v11))
                        _ = v18
                    default:
                        v2 = ""
                    }
                    v16 += (v15 / ((v15) | 1))
                    _ = v16
                }
                _ = fi14
                _ = v15
            }
            for fi19 := 0; fi19 < 9; fi19++ {
                v20 := GenS2{F0: -9223372036854775808, F1: "a", F2: 4294967295}
                _ = fi19
                _ = v20
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        _ = fi10
        _ = v11
    }
    v1 = (("" + v4) < v4)
    v3 += (v3 / (((-(v3))) | 1))
    var v21 uint = 1
    v22 := GenS3{F0: 432, F1: -2147483648}
    fmt.Println(v4, v4, v4)
    switch -2147483648 {
    case 870:
        v1 = true
    case -212:
        v23 := GenS1{F0: 932, F1: v3}
        _ = v23
    case -297:
        if false {
            v1 = (("hello" + "hello") <= v2)
            v21 += (v21 % (((v21 >> (((v21 % (((v21 << ((523) & 63))) | 1))) & 63))) | 1))
            for fi24 := 0; fi24 < 3; fi24++ {
                var v25 int32 = v3
                switch (fi24 & (fi24 >> (((v21 - v21)) & 63))) {
                case 42:
                    for fi26 := 0; fi26 < 5; fi26++ {
                        var v27 string = "foo bar"
                        var v28 uint16 = 65535
                        v29 := GenS2{F0: 42, F1: ("\n\t" + "foo bar"), F2: 614}
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                case -372:
                    var v30 uint64 = 1
                    _ = v30
                }
                var v31 uint64 = 4294967295
                _ = fi24
                _ = v25
                _ = v31
            }
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            fmt.Println(v1)
        }
    }
    var v32 uint16 = 255
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v33 uint16 = 1
    v1 = ((float64(-2.25) + (1e10 * (-2.25 / 1.5))) < float64(0.0))
    v1 = (!v1)
    switch 959 {
    case -1:
        if len(v5) > 0 {
            v5[0] = 2
        }
    case 9223372036854775807:
        var v34 uint32 = 2
        _ = v34
    default:
        for fi35 := 0; fi35 < 3; fi35++ {
            var v36 int64 = 9223372036854775807
            v21 = 293
            _ = fi35
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 18; fi37++ {
        if (!false) {
            for fi38 := 0; fi38 < 5; fi38++ {
                switch (^(fi38)) {
                case 2147483647:
                    switch v4 {
                    case "":
                        fmt.Println(fi38, fi37, v33)
                    }
                }
                if false {
                    switch v4 {
                    case "hello":
                        if true {
                            switch "" {
                            case "a":
                                v39 := GenS2{F0: 42, F1: "\n\t", F2: 159}
                                _ = v39
                            default:
                                v33 = (v32 + (v32 / ((v33) | 1)))
                            }
                        }
                    case "foo bar":
                        v4 = ((("a" + v4) + "foo bar") + "a")
                    }
                } else {
                    v40 := GenS2{F0: 9223372036854775807, F1: v4, F2: 646}
                    if (40 >= v32) {
                        var v41 int64 = -533
                        fmt.Println(v2)
                        _ = v41
                    } else {
                        v5 = append(v5, 1)
                    }
                    _ = v40
                }
                if true {
                    if len(v5) > 0 {
                        v5[0] = 2
                    }
                    v2 = (v4 + ("" + ("" + v2)))
                    var v42 uint64 = 742
                    _ = v42
                } else {
                    v1 = false
                }
                _ = fi38
            }
            v4 = "\n\t"
        }
        fmt.Println(v4, v21)
        _ = fi37
    }
    var v43 int32 = (v3 >> ((v21) & 63))
    fmt.Println(v33, v1, v32)
    v44 := GenS4{F0: 42, F1: 90}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
