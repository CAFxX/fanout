package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int }
type GenS2 struct { F0 int16; F1 string; F2 string }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 uint8; F1 string; F2 int8 }
type GenS6 struct { F0 uint16; F1 uint64 }
func genFn0() int64 {
    if (!false) {
        if false {
            var v5 int32 = 42
            for fi6 := 0; fi6 < 12; fi6++ {
                switch "a" {
                case "\n\t":
                    for fi7 := 0; fi7 < 15; fi7++ {
                        if true {
                            v8 := []int{fi6, 2147483647, (fi6 >> ((4294967295) & 63))}
                            _ = v8
                        }
                        _ = fi7
                    }
                default:
                    fmt.Println(fi6)
                }
                v9 := []int{-2147483648, (fi6 & fi6)}
                v5 = -2147483648
                _ = fi6
                _ = v9
            }
            for fi10 := 0; fi10 < 6; fi10++ {
                fmt.Println(fi10, fi10)
                v11 := []int{fi10, (fi10 / ((fi10) | 1)), -1}
                _ = fi10
                _ = v11
            }
            _ = v5
        }
    }
    v12 := GenS1{F0: -125, F1: "", F2: 834}
    v13 := GenS1{F0: -9223372036854775808, F1: ("foo bar" + "hello"), F2: -225}
    for fi14 := 0; fi14 < 4; fi14++ {
        v15 := GenS1{F0: -9223372036854775808, F1: ("" + "\n\t"), F2: 9223372036854775807}
        for fi16 := 0; fi16 < 9; fi16++ {
            if (1000 >= 792) {
                var v17 uint32 = 822
                if true {
                    switch "" {
                    case "foo bar":
                        v17 += v17
                        if (uint16(v17) >= 0) {
                            v17 -= (v17 << ((4294967295) & 63))
                            if false {
                                v17 += (v17 + v17)
                                v17 = (v17 >> ((0) & 63))
                                var v18 uint = 693
                                _ = v18
                            }
                            var v19 float64 = 1e10
                            _ = v19
                        } else {
                            if false {
                                fmt.Println(v17, v17)
                            }
                            var v20 string = "a"
                            v17 -= 69
                            _ = v20
                        }
                    }
                    if false {
                        var v21 uint16 = 6
                        v17 += 2
                        v22 := GenS2{F0: 32767, F1: (("" + "a") + ("foo bar" + "\n\t")), F2: "hello"}
                        _ = v21
                        _ = v22
                    }
                    v17 -= 981
                }
                var v23 int16 = 1
                _ = v17
                _ = v23
            } else {
                var v24 uint32 = 951
                if true {
                    v25 := []int{fi16}
                    if len(v25) > 0 {
                        v25[0] = 4294967295
                    }
                    _ = v25
                }
                v26 := GenS1{F0: -9223372036854775808, F1: "foo bar", F2: fi14}
                _ = v24
                _ = v26
            }
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    var v27 uint8 = 127
    _ = v12
    _ = v13
    _ = v27
    return -9223372036854775808
}
func genFn1(p0 string) int64 {
    v28 := GenS1{F0: -757, F1: "\n\t", F2: 1}
    v29 := GenS2{F0: 1000, F1: ((p0 + "hello") + p0), F2: (((p0 + ((("a" + "a") + ("hello" + p0)) + ("" + "foo bar"))) + ("a" + "hello")) + "hello")}
    fmt.Println(p0)
    if (4294967295 != 18446744073709551615) {
        p0 = p0
    } else {
        for fi30 := 0; fi30 < 18; fi30++ {
            p0 = "foo bar"
            var v31 float64 = 1e10
            _ = fi30
            _ = v31
        }
        v32 := GenS3{F0: p0, F1: 814}
        fmt.Println(p0)
        _ = v32
    }
    if (!(149 < 65535)) {
        v33 := []int{2147483647, -333, 42}
        fmt.Println(p0)
        _ = v33
    }
    _ = p0
    _ = v28
    _ = v29
    return -9223372036854775808
}

func main() {
    v1 := []int{32767, 9223372036854775807}
    var v2 int8 = -128
    var v3 int = 32767
    var v4 int16 = 838
    v34 := GenS4{F0: "a", F1: (^(v4))}
    switch "a" {
    case "a":
        v35 := GenS5{F0: 48, F1: ((("" + ("a" + "")) + "\n\t") + "foo bar"), F2: v2}
        _ = v35
    }
    if len(v1) > 0 {
        v1[0] = v3
    }
    if (false && false) {
        v4 -= -564
        if (!(-9223372036854775808 >= (v3 & (^(v3))))) {
            _ = len(v1)
        } else {
            v36 := GenS2{F0: v4, F1: "\n\t", F2: ""}
            _ = v36
        }
        for fi37 := 0; fi37 < 18; fi37++ {
            _ = len(v1)
            _ = fi37
        }
    } else {
        var v38 bool = false
        _ = len(v1)
        _ = v38
    }
    v39 := GenS6{F0: 121, F1: 42}
    var v40 float32 = (3.14159 / 3.14159)
    v41 := GenS3{F0: "foo bar", F1: 2}
    fmt.Println(v3, v40)
    fmt.Println(v3)
    v3 = (-(v3))
    if len(v1) > 0 {
        v1[0] = v3
    }
    fmt.Println(v3, v4)
    var v42 uint64 = 0
    for fi43 := 0; fi43 < 1; fi43++ {
        v2 = v2
        _ = fi43
    }
    v42 += 4294967295
    var v44 int = (-(v3))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v44
}
