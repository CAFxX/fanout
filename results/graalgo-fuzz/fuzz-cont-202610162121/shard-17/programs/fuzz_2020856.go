package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 int; F1 uint64; F2 int8 }
func genFn0() uint16 {
    for fi6 := 0; fi6 < 20; fi6++ {
        fmt.Println(fi6)
        _ = fi6
    }
    v7 := []int{32768, -1, 2147483647, 1}
    _ = v7
    return 32767
}

func main() {
    v1 := []string{("hello" + (("\n\t" + (("hello" + ("a" + "")) + ("hello" + (("hello" + "foo bar") + "a")))) + "\n\t"))}
    var v2 uint32 = 42
    var v3 int16 = 42
    var v4 float32 = float32(0.0)
    v5 := make(map[uint32]bool)
    if false {
        var v8 uint64 = 1
        _ = v8
    }
    v3 = v3
    var v9 float32 = ((float32(-2.25) * v4) + v4)
    var v10 uint16 = 255
    v4 = (3.14159 * 0.1)
    var v11 float64 = (3.14159 * (0.0 + (1.5 * (1.5 * 1e10))))
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        fmt.Println(v10, v2, v10)
        _ = fi12
    }
    _ = len(v1)
    var v13 float32 = v9
    var v14 int64 = 42
    if (476 <= 2) {
        fmt.Println(v14, v3, v3)
        if true {
            for fi15 := 0; fi15 < 3; fi15++ {
                for fi16 := 0; fi16 < 7; fi16++ {
                    var v17 string = "hello"
                    v4 += (v13 * float32(v11))
                    _ = fi16
                    _ = v17
                }
                var v18 uint16 = 846
                _ = fi15
                _ = v18
            }
            v9 = float32(float64(float32(((v11 * v11) * -2.25))))
        } else {
            for fi19 := 0; fi19 < 4; fi19++ {
                v20 := GenS1{F0: 1}
                _ = fi19
                _ = v20
            }
        }
        if false {
            if false {
                v21 := GenS1{F0: 42}
                var v22 uint = 4294967295
                fmt.Println(v22, v22, v13)
                _ = v21
                _ = v22
            } else {
                if (true && (!true)) {
                    fmt.Println(v4, v14, v2)
                } else {
                    switch -308 {
                    case 2147483647:
                        v23 := GenS2{F0: -2147483648, F1: -27}
                        _ = v23
                    case 0:
                        if (!(2147483647 == 42)) {
                            v14 -= v14
                        }
                    default:
                        v10 = v10
                    }
                }
                if len(v1) > 0 {
                    v1[0] = "a"
                }
            }
            delete(v5, v2)
        } else {
            if len(v1) > 0 {
                v1[0] = ""
            }
            v13 = v13
            v24 := GenS1{F0: 390}
            _ = v24
        }
    }
    if ((-(v3)) >= (v3 ^ (v3 << ((18446744073709551615) & 63)))) {
        v11 = -2.25
        v13 -= v13
    } else {
        v9 = v4
    }
    if true {
        v14 += (v14 / (((v14 | (v14 - (v14 ^ 1)))) | 1))
    }
    fmt.Println(v14, v14)
    if false {
        switch "hello" {
        case "\n\t":
            if (v4 >= (v4 - v4)) {
                for fi25 := 0; fi25 < 3; fi25++ {
                    delete(v5, v2)
                    _ = fi25
                }
                v2 = 0
                if len(v1) > 0 {
                    v1[0] = ((("" + ("\n\t" + "foo bar")) + ((("a" + "foo bar") + ((("foo bar" + "\n\t") + "\n\t") + "a")) + ((("a" + "\n\t") + (("a" + "foo bar") + ("" + ("a" + ((("" + "") + "hello") + "foo bar"))))) + "a"))) + (("a" + "hello") + "foo bar"))
                }
            } else {
                switch "hello" {
                case "foo bar":
                    v26 := GenS3{F0: (v3 & v3), F1: 0, F2: 2147483647}
                    _ = v26
                case "hello":
                    _ = len(v1)
                case "":
                    delete(v5, v2)
                }
                var v27 uint8 = 1
                if ((4294967295 < 517) || true) {
                    v28 := GenS4{F0: 4294967295}
                    for fi29 := 0; fi29 < 8; fi29++ {
                        for fi30 := 0; fi30 < 12; fi30++ {
                            switch 9223372036854775807 {
                            case -2147483648:
                                var v31 int32 = 725
                                _ = v31
                            case -18:
                                fmt.Println(v27)
                            case 4294967295:
                                v2 = v2
                            }
                            _ = fi30
                        }
                        _ = fi29
                    }
                    _ = v28
                } else {
                    for fi32 := 0; fi32 < 2; fi32++ {
                        var v33 int = (fi32 & fi32)
                        _ = fi32
                        _ = v33
                    }
                    var v34 int8 = 1
                    _ = v34
                }
                _ = v27
            }
        }
    }
    v35 := GenS5{F0: -9223372036854775808, F1: 0}
    if false {
        v36 := GenS5{F0: -554, F1: 368}
        if false {
            v37 := GenS6{F0: 9223372036854775807, F1: 157, F2: 57}
            _ = v37
        }
        var v38 int8 = int8(v10)
        _ = v36
        _ = v38
    } else {
        v4 = v9
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v9
}
