package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS4 struct { F0 uint32; F1 int32; F2 uint64 }

func main() {
    var v1 int = -1
    var v2 float32 = 3.14159
    v3 := make(map[string]bool)
    var v4 int32 = 456
    v5 := make(map[uint32]uint32)
    var v6 int32 = v4
    switch (v1 % ((v1) | 1)) {
    case 32767:
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := GenS1{F0: v4}
            v9 := GenS2{F0: 291, F1: 1, F2: 58}
            switch 128 {
            case -464:
                if true {
                    for fi10 := 0; fi10 < 3; fi10++ {
                        var v11 uint32 = 42
                        _ = fi10
                        _ = v11
                    }
                } else {
                    if ((!true) || false) {
                        v12 := GenS2{F0: 0, F1: 2, F2: int64(v6)}
                        _ = v12
                    } else {
                        v13 := []int{128}
                        if (true || (129 != 1)) {
                            v14 := GenS1{F0: v4}
                            v2 += (v2 - (float32(1.5) - (v2 * v2)))
                            _ = v14
                        }
                        _ = v13
                    }
                }
            case -9223372036854775808:
                var v15 string = ""
                _ = v15
            case 42:
                v16 := []int{-661, (fi7 + fi7), (v1 - v1), (fi7 >> ((375) & 63))}
                _ = v16
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    default:
        v3["foo bar"] = (16 < 198)
    }
    v17 := []int{2147483647, (v1 >> ((4294967295) & 63)), (-(v1))}
    var v18 int16 = 1
    v19 := GenS1{F0: (v6 >> ((42) & 63))}
    switch 32767 {
    case 320:
        if (((0.0 * 1.5) - -2.25) <= 1e10) {
            for fi20 := 0; fi20 < 15; fi20++ {
                var v21 uint = 217
                v1 = v1
                v4 = v6
                _ = fi20
                _ = v21
            }
            v1 -= (v1 + 32767)
        } else {
            if false {
                switch "" {
                case "\n\t":
                    _ = len(v17)
                default:
                    var v22 uint64 = 42
                    _ = v22
                }
                v17 = append(v17, (v1 << ((790) & 63)))
            }
            v6 += (v6 | (v6 + v4))
        }
    default:
        v5[330] = 351
    }
    _ = len(v17)
    v6 = (v4 - v4)
    if (("a" + ("" + "foo bar")) < ("" + ("a" + ((("foo bar" + ("a" + ("foo bar" + ("\n\t" + "\n\t")))) + "a") + "foo bar")))) {
        if true {
            v23 := GenS3{F0: -9223372036854775808, F1: 42, F2: 1}
            _ = v23
        } else {
            fmt.Println(v18, v6)
            if (true || false) {
                v4 += v6
            } else {
                if len(v17) > 0 {
                    v17[0] = (-(v1))
                }
                v24 := v3["a"]
                var v25 uint8 = 2
                _ = v24
                _ = v25
            }
        }
    } else {
        if (false && false) {
            v1 += -125
        }
        var v26 float64 = (0.1 - float64((0.1 - (3.14159 - (((0.0 * 1e10) * v2) * 1.5)))))
        var v27 int16 = v18
        _ = v26
        _ = v27
    }
    v4 -= 42
    var v28 bool = true
    v29 := GenS4{F0: 1, F1: -1, F2: 42}
    v18 += 1
    v30 := GenS3{F0: -1, F1: 1, F2: 1}
    v18 = 1000
    if v28 {
        v2 = 0.1
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        switch 4294967295 {
        case -770:
            switch ("\n\t" + (((("hello" + "") + "foo bar") + "\n\t") + ((("\n\t" + "") + ("\n\t" + ("foo bar" + ""))) + "hello"))) {
            case "hello":
                _ = len(v3)
            case "\n\t":
                v32 := GenS3{F0: -265, F1: 929, F2: 0}
                _ = v32
            }
        }
        var v33 uint8 = 42
        _ = fi31
        _ = v33
    }
    var v34 float32 = v2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
