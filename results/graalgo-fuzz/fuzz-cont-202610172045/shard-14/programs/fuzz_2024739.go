package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 uint32; F1 uint32 }
func genFn0(p0 int32) int8 {
    var v7 float32 = ((float32(((1.5 * 1e10) * ((((1e10 / 1e10) - (((float64(1e10) + 1.5) + 3.14159) + 1.5)) - ((3.14159 / 0.1) - (((1.5 - 0.0) * (1e10 - 1.5)) + 0.1))) - float64(-2.25)))) / 1.5) / float32((3.14159 + 0.1)))
    for fi8 := 0; fi8 < 20; fi8++ {
        switch 2147483647 {
        case 0:
            p0 += p0
        case 154:
            v9 := []int{4294967295}
            _ = v9
        default:
            switch ("foo bar" + "a") {
            case "\n\t":
                v7 -= 0.1
            }
        }
        _ = fi8
    }
    _ = p0
    _ = v7
    return 127
}

func main() {
    var v1 bool = (-1 > 42)
    v2 := []string{}
    v3 := make(map[string]uint32)
    var v4 int = -196
    var v5 uint32 = 0
    v6 := []int{127, (v4 / ((int(v5)) | 1))}
    var v10 float32 = (3.14159 - 0.0)
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    v11 := GenS1{F0: ""}
    v6 = append(v6, (-(v4)))
    v12 := GenS1{F0: "foo bar"}
    var v13 uint64 = 42
    var v14 float64 = float64((v10 - (0.1 * -2.25)))
    var v15 int16 = 76
    if false {
        var v16 bool = v1
        if ((!v1) && v1) {
            v5 -= v5
            switch "\n\t" {
            case "hello":
                switch ("\n\t" + "\n\t") {
                case "hello":
                    var v17 float32 = v10
                    _ = v17
                }
            case "":
                var v18 uint = 181
                _ = v18
            case "a":
                fmt.Println(v14, v10, v13)
            }
        } else {
            for fi19 := 0; fi19 < 5; fi19++ {
                switch "a" {
                case "":
                    if (42 <= uint32(v13)) {
                        if v16 {
                            switch ("" + (("hello" + "\n\t") + (("foo bar" + ("" + "hello")) + "hello"))) {
                            case "a":
                                delete(v3, "\n\t")
                            }
                            var v20 uint = 18446744073709551615
                            v21 := GenS1{F0: (("a" + "hello") + ("foo bar" + "\n\t"))}
                            _ = v20
                            _ = v21
                        } else {
                            v5 += 2147483648
                            v16 = v1
                        }
                        var v22 uint = 18446744073709551615
                        for fi23 := 0; fi23 < 6; fi23++ {
                            var v24 uint32 = v5
                            fmt.Println(v1, v1, fi19)
                            var v25 uint32 = (v24 << (((v22 % ((65535) | 1))) & 63))
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        _ = v22
                    } else {
                        if v16 {
                            v26 := GenS2{F0: uint8(v5)}
                            _ = v26
                        }
                        fmt.Println(v5)
                        v3["a"] = 2
                    }
                case "\n\t":
                    var v27 int64 = -932
                    _ = v27
                case "hello":
                    v28 := GenS3{F0: 9223372036854775807, F1: -112}
                    _ = v28
                default:
                    for fi29 := 0; fi29 < 18; fi29++ {
                        var v30 uint32 = v5
                        _ = fi29
                        _ = v30
                    }
                }
                v31 := GenS1{F0: ""}
                for fi32 := 0; fi32 < 13; fi32++ {
                    v33 := GenS4{F0: v5, F1: (v5 << ((1) & 63))}
                    var v34 float64 = ((v14 * v14) / 3.14159)
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                _ = fi19
                _ = v31
            }
        }
        _ = v16
    }
    v35 := GenS2{F0: 2}
    var v36 int32 = -1
    v37 := GenS4{F0: v5, F1: (^(v5))}
    fmt.Println(v4, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
