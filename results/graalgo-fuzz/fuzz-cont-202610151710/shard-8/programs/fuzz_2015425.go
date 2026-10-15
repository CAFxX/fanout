package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 uint64; F2 uint64 }
func genFn0() uint16 {
    v6 := GenS1{F0: 1}
    var v7 float32 = (1.5 / 3.14159)
    v8 := []int{0, 32768}
    v9 := GenS2{F0: -2147483648}
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 uint32 = 4294967295
        for fi12 := 0; fi12 < 8; fi12++ {
            v7 += v7
            var v13 uint16 = 42
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]bool)
    var v3 int64 = 42
    var v4 int32 = 981
    v5 := make(map[string]bool)
    v3 = 42
    for fi14 := 0; fi14 < 12; fi14++ {
        v5["foo bar"] = (true && (!false))
        v15 := []int{(fi14 & fi14), (fi14 | fi14), (-(fi14))}
        fmt.Println(v4, fi14)
        _ = fi14
        _ = v15
    }
    var v16 int = 128
    v17 := GenS1{F0: v4}
    v4 = (v4 / ((v4) | 1))
    v18 := v5[("foo bar" + ("foo bar" + ("a" + "hello")))]
    v18 = ((1 >= 2) || ((v3 % (((v3 ^ v3)) | 1)) > 0))
    var v19 int16 = 42
    v5[("" + "foo bar")] = v18
    switch "" {
    case "a":
        if (!(32767 != genFn0())) {
            for fi20 := 0; fi20 < 16; fi20++ {
                v18 = (v18 && v18)
                _ = fi20
            }
        } else {
            v19 = v19
        }
    case "hello":
        for fi21 := 0; fi21 < 12; fi21++ {
            v4 = (^(v4))
            _ = fi21
        }
    default:
        if (!v18) {
            for fi22 := 0; fi22 < 13; fi22++ {
                v23 := GenS2{F0: v4}
                _ = fi22
                _ = v23
            }
            if false {
                v24 := GenS1{F0: v4}
                var v25 string = ""
                switch v16 {
                case -156:
                    v26 := GenS1{F0: (^(v4))}
                    _ = v26
                }
                _ = v24
                _ = v25
            } else {
                var v27 int = v16
                _ = v27
            }
        } else {
            if true {
                if ((0.1 + 0.0) < (-2.25 / 0.1)) {
                    switch "hello" {
                    case "\n\t":
                        switch v16 {
                        case -9223372036854775808:
                            var v28 int = v16
                            _ = v28
                        case -288:
                            var v29 uint32 = 0
                            _ = v29
                        default:
                            var v30 string = (("" + "a") + ("" + "a"))
                            _ = v30
                        }
                    default:
                        v31 := GenS3{F0: v4, F1: v1, F2: (v1 % ((v1) | 1))}
                        _ = v31
                    }
                    delete(v5, "a")
                    v32 := []int{v16, (-(v16)), -2147483648}
                    _ = v32
                }
            } else {
                v16 -= v16
            }
        }
    }
    v19 = (v19 / ((v19) | 1))
    var v33 float32 = float32(-2.25)
    var v34 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
