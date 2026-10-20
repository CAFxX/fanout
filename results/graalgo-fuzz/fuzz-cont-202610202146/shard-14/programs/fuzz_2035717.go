package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int32 }
type GenS2 struct { F0 int16; F1 bool; F2 int32 }
type GenS3 struct { F0 uint32; F1 int64; F2 int16 }

func main() {
    v1 := []string{("a" + "hello"), "a", "hello"}
    var v2 uint = 42
    v3 := make(map[uint32]bool)
    v4 := make(map[string]bool)
    var v5 bool = false
    if v5 {
        fmt.Println(v2)
        v2 -= 184
    }
    if (185 == 65535) {
        _ = len(v1)
        switch -1 {
        case -2147483648:
            if v5 {
                if false {
                    v5 = v5
                    var v6 int64 = 1
                    _ = v6
                } else {
                    switch ("foo bar" + "") {
                    case "a":
                        for fi7 := 0; fi7 < 17; fi7++ {
                            if false {
                                var v8 float64 = -2.25
                                v1 = append(v1, "hello")
                                switch 128 {
                                case -9223372036854775808:
                                    if v5 {
                                        v5 = (-797 != 158)
                                    }
                                }
                                _ = v8
                            }
                            v2 = 761
                            _ = fi7
                        }
                    case "\n\t":
                        v2 = (v2 >> ((271) & 63))
                    default:
                        v5 = ((v2 - 1) <= 4294967295)
                    }
                }
            }
        case 0:
            for fi9 := 0; fi9 < 10; fi9++ {
                if v5 {
                    var v10 string = ("hello" + "a")
                    _ = v10
                } else {
                    v11 := GenS1{F0: 9223372036854775807, F1: 9223372036854775807, F2: -2147483648}
                    if true {
                        var v12 uint16 = 42
                        _ = v12
                    } else {
                        var v13 int = (fi9 + 0)
                        if (0.1 == -2.25) {
                            delete(v4, ((("" + "") + "hello") + ("\n\t" + (("hello" + ("" + "")) + ""))))
                        }
                        v14 := GenS1{F0: 0, F1: -9223372036854775808, F2: 0}
                        _ = v13
                        _ = v14
                    }
                    _ = v11
                }
                _ = fi9
            }
        case 42:
            fmt.Println(v2)
        }
        v2 = (v2 % ((2) | 1))
    } else {
        v5 = true
        v15 := GenS1{F0: -809, F1: 441, F2: 42}
        v5 = ((float64(0.1) + -2.25) >= 1e10)
        _ = v15
    }
    var v16 uint64 = 42
    if (!false) {
        if v5 {
            fmt.Println(v5)
            var v17 uint32 = 2147483648
            _ = v17
        } else {
            if v5 {
                var v18 uint32 = 701
                v19 := GenS1{F0: -840, F1: -841, F2: 0}
                v5 = false
                _ = v18
                _ = v19
            }
        }
        v20 := GenS2{F0: 0, F1: (((("foo bar" + "\n\t") + "a") + ("a" + ("\n\t" + "a"))) <= "\n\t"), F2: 0}
        for fi21 := 0; fi21 < 13; fi21++ {
            if true {
                v1 = append(v1, "foo bar")
                var v22 string = "\n\t"
                _ = v22
            } else {
                var v23 uint64 = 1
                _ = v23
            }
            v24 := GenS1{F0: 9223372036854775807, F1: -572, F2: 2147483647}
            v25 := GenS3{F0: 0, F1: 0, F2: -331}
            _ = fi21
            _ = v24
            _ = v25
        }
        _ = v20
    }
    v5 = (false && (false || (!v5)))
    var v26 uint8 = 255
    var v27 float32 = ((-2.25 / 1e10) * 0.1)
    v2 = (v2 - v2)
    var v28 int = -496
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
