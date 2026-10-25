package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int64, p1 uint, p2 uint8) string {
    if (true && false) {
        p0 = (^(p0))
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 float32 = (0.1 + (3.14159 - 1e10))
        var v9 int8 = 127
        var v10 int = (^(fi7))
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    p2 += 42
    _ = p0
    _ = p1
    _ = p2
    return ""
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{}
    v3 := make(map[string]string)
    var v4 int32 = 218
    var v5 uint8 = 9
    var v6 uint8 = v5
    v5 -= v5
    v4 += -2147483648
    switch 1 {
    case -9223372036854775808:
        for fi11 := 0; fi11 < 11; fi11++ {
            if true {
                var v12 string = "\n\t"
                v4 -= (v4 ^ 1)
                var v13 float32 = (float32((1.5 * (3.14159 / -2.25))) / ((3.14159 / 1.0) + -2.25))
                _ = v12
                _ = v13
            }
            if false {
                v4 += (-(v4))
                v14 := GenS1{F0: ((("a" + (((("" + "\n\t") + "\n\t") + (("hello" + "\n\t") + ("\n\t" + "hello"))) + "")) + ("a" + (("\n\t" + "\n\t") + ("\n\t" + (("" + "\n\t") + "hello"))))) + "foo bar"), F1: false}
                switch "hello" {
                case "hello":
                    var v15 uint64 = 2
                    _ = v15
                default:
                    v5 += v5
                }
                _ = v14
            }
            delete(v3, "")
            _ = fi11
        }
    default:
        v5 = v6
    }
    v2 = append(v2, "foo bar")
    if false {
        if (4294967295 <= 4294967295) {
            switch 234 {
            case 127:
                if false {
                    fmt.Println(v4, v6)
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                } else {
                    v16 := GenS1{F0: "", F1: (false && true)}
                    var v17 int8 = -1
                    _ = v16
                    _ = v17
                }
            }
        }
        for fi18 := 0; fi18 < 2; fi18++ {
            for fi19 := 0; fi19 < 3; fi19++ {
                fmt.Println(fi18, fi18, fi18)
                v20 := GenS2{F0: 42, F1: false}
                _ = len(v2)
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
        var v21 uint8 = 83
        _ = v21
    } else {
        var v22 uint64 = 42
        _ = v22
    }
    fmt.Println(v4, v5)
    v23 := GenS2{F0: v5, F1: false}
    var v24 int = 42
    var v25 int8 = 1
    fmt.Println(v25, v25, v25)
    if (!(-32768 > -32768)) {
        var v26 string = "\n\t"
        var v27 float32 = (float32(((1.5 - 0.0) + ((((0.1 - ((1e10 / ((1.5 + 1.5) / 3.14159)) - -2.25)) * 1e10) + 0.0) / ((1e10 - 3.14159) / 1.0)))) - 3.14159)
        v26 = "foo bar"
        _ = v26
        _ = v27
    } else {
        v28 := GenS3{F0: 4294967295}
        _ = v28
    }
    v4 = (v4 << ((4294967295) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
