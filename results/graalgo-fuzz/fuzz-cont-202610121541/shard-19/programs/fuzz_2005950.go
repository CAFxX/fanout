package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int32 }
type GenS2 struct { F0 int64; F1 int32 }

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"a"}
    v3 := []string{"a", "\n\t", ("" + ((("hello" + "\n\t") + "\n\t") + ("" + "\n\t")))}
    for fi4 := 0; fi4 < 12; fi4++ {
        v5 := GenS1{F0: 1, F1: 0, F2: 664}
        var v6 int8 = -101
        var v7 string = ("foo bar" + "")
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 100, F1: -268, F2: 110}
    var v9 uint32 = 65535
    v10 := GenS1{F0: 88, F1: 0, F2: 0}
    v11 := v1[32767]
    for fi12 := 0; fi12 < 17; fi12++ {
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        _ = len(v2)
        _ = fi12
    }
    v9 += 734
    switch "a" {
    case "foo bar":
        fmt.Println(v9)
    default:
        v1[-277] = true
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v9)
    if v11 {
        if true {
            var v13 uint = 336
            if (((3.14159 / (float32(3.14159) - 1.5)) != 1e10) || (v11 && (0 < 2))) {
                v13 -= (v13 & (v13 << ((v13) & 63)))
                if v11 {
                    v14 := GenS2{F0: -9223372036854775808, F1: 1}
                    _ = v14
                } else {
                    v11 = (128 != 0)
                    var v15 string = ("\n\t" + (("hello" + "hello") + "a"))
                    _ = v15
                }
                if len(v2) > 0 {
                    v2[0] = (((("\n\t" + "\n\t") + ("" + "")) + "foo bar") + "")
                }
            } else {
                if len(v3) > 0 {
                    v3[0] = ("\n\t" + "\n\t")
                }
            }
            v9 -= 660
            _ = v13
        } else {
            if ((65535 < 255) && v11) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            v1[531] = (!(!(!(168 < 0))))
            var v16 uint16 = 255
            _ = v16
        }
        var v17 uint16 = 692
        for fi18 := 0; fi18 < 11; fi18++ {
            if (42 == 127) {
                v1[-2147483648] = v11
                if (!("\n\t" != ("a" + "\n\t"))) {
                    var v19 uint = 279
                    var v20 int32 = 2147483647
                    if (v20 <= 702) {
                        switch "a" {
                        case "hello":
                            v19 = (^(v19))
                        case "":
                            v19 -= 493
                        default:
                            _ = len(v1)
                        }
                    }
                    _ = v19
                    _ = v20
                } else {
                    v9 = (v9 - v9)
                    v21 := GenS1{F0: -90, F1: 9223372036854775807, F2: 374}
                    _ = v21
                }
            }
            for fi22 := 0; fi22 < 20; fi22++ {
                var v23 int64 = 1
                _ = fi22
                _ = v23
            }
            _ = fi18
        }
        _ = v17
    }
    var v24 uint32 = v9
    if true {
        var v25 float32 = float32(1e10)
        v11 = false
        v26 := GenS1{F0: 10, F1: 9223372036854775807, F2: 473}
        _ = v25
        _ = v26
    }
    if (v9 != (v24 - v24)) {
        v1[-750] = v11
    } else {
        var v27 uint64 = 42
        fmt.Println(v11, v11, v24)
        _ = v27
    }
    var v28 int32 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v8
    _ = v9
}
