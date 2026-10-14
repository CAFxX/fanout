package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int8 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 bool }
func genFn0(p0 int8, p1 uint8) uint16 {
    var v4 string = "hello"
    p1 += (^(p1))
    _ = p0
    _ = p1
    _ = v4
    return 32767
}
func genFn1(p0 uint8, p1 int16, p2 uint) int {
    v5 := []int{-2147483648}
    if ((4294967295 != 65535) || true) {
        for fi6 := 0; fi6 < 19; fi6++ {
            p2 = (p2 & (p2 ^ 18446744073709551615))
            if true {
                for fi7 := 0; fi7 < 7; fi7++ {
                    v8 := GenS1{F0: -1, F1: 879, F2: 57}
                    _ = fi7
                    _ = v8
                }
                v9 := GenS2{F0: "", F1: p0}
                _ = v9
            }
            _ = fi6
        }
        var v10 string = ""
        var v11 uint = 403
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -258
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    var v3 float32 = 1e10
    switch 9223372036854775807 {
    case 4294967295:
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := []int{fi12, (fi12 >> ((255) & 63)), 467, fi12}
            _ = fi12
            _ = v13
        }
    case -9223372036854775808:
        v3 = 0.1
    case -1:
        for fi14 := 0; fi14 < 16; fi14++ {
            if (!true) {
                v15 := []int{genFn1(128, -1, 255), 9223372036854775807, (fi14 + (fi14 + (fi14 ^ fi14)))}
                for fi16 := 0; fi16 < 13; fi16++ {
                    if false {
                        v17 := GenS1{F0: -1, F1: 9223372036854775807, F2: 81}
                        _ = v17
                    }
                    v3 = 0.1
                    v3 = 0.0
                    _ = fi16
                }
                switch fi14 {
                case 127:
                    for fi18 := 0; fi18 < 8; fi18++ {
                        v3 = 0.1
                        _ = fi18
                    }
                case 4294967295:
                    var v19 int64 = 9223372036854775807
                    _ = v19
                default:
                    v3 += v3
                }
                _ = v15
            } else {
                for fi20 := 0; fi20 < 11; fi20++ {
                    v21 := GenS2{F0: "hello", F1: 108}
                    fmt.Println(fi20)
                    var v22 bool = false
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                if false {
                    v3 = (v3 * v3)
                    if true {
                        v23 := []int{fi14, (fi14 | fi14), 1}
                        _ = v23
                    } else {
                        var v24 string = "a"
                        _ = v24
                    }
                    v25 := GenS2{F0: "a", F1: 255}
                    _ = v25
                }
            }
            _ = fi14
        }
    }
    var v26 int32 = -2147483648
    if (2147483648 == 838) {
        var v27 float64 = -2.25
        var v28 uint64 = 0
        _ = v27
        _ = v28
    }
    if (4294967295 > 754) {
        if false {
            switch (("" + "\n\t") + "foo bar") {
            case "a":
                v26 = (v26 + (v26 + v26))
            case "hello":
                switch "foo bar" {
                case "":
                    fmt.Println(v26, v26)
                case "foo bar":
                    var v29 int32 = v26
                    _ = v29
                default:
                    v30 := v2[-352]
                    _ = v30
                }
            default:
                var v31 int16 = 0
                _ = v31
            }
            v32 := []int{}
            _ = v32
        }
        if (42 < 78) {
            if (!false) {
                v26 -= 42
            }
        } else {
            v26 += v26
        }
        v33 := v2[42]
        _ = v33
    } else {
        fmt.Println(v3, v26)
    }
    v26 = v26
    v34 := GenS2{F0: "\n\t", F1: 78}
    v2[856] = -9223372036854775808
    var v35 bool = ((!true) || (9223372036854775807 == genFn1(0, 1, 65535)))
    v26 = v26
    v36 := []int{-2147483648}
    var v37 uint64 = 4294967295
    if (!(4294967295 >= 65535)) {
        switch 409 {
        case 853:
            fmt.Println(v3, v3)
        default:
            var v38 int = 208
            _ = v38
        }
        v37 = (v37 | (v37 | 42))
        fmt.Println(v35)
    } else {
        fmt.Println(v37, v26)
        delete(v2, -2147483648)
        v3 += 0.1
    }
    switch ("foo bar" + "\n\t") {
    case "":
        if len(v36) >= 2 {
            v36 = v36[1:len(v36)]
        }
    case "foo bar":
        v26 = (v26 << ((165) & 63))
    case "\n\t":
        v3 = 3.14159
    default:
        var v39 int = genFn1(178, -1, 4294967295)
        _ = v39
    }
    var v40 uint = 871
    if (244 <= 0) {
        v41 := GenS3{F0: v35}
        _ = v41
    }
    var v42 bool = (!v35)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v40
    _ = v42
}
