package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint; F1 string; F2 uint16 }
type GenS3 struct { F0 string; F1 int32 }

func main() {
    var v1 uint8 = 42
    v2 := make(map[int]string)
    var v3 float64 = 0.1
    switch ("hello" + ((((("foo bar" + "hello") + "foo bar") + "") + "a") + ((("a" + "\n\t") + ("a" + "")) + "a"))) {
    case "a":
        if true {
            v1 += (^(v1))
            fmt.Println(v3)
        } else {
            v4 := GenS1{F0: 4294967295, F1: -1}
            _ = v4
        }
    case "":
        v5 := GenS2{F0: 255, F1: "foo bar", F2: 42}
        _ = v5
    default:
        var v6 int8 = 42
        _ = v6
    }
    switch "" {
    case "":
        v1 -= (v1 * (v1 << ((65535) & 63)))
    case "hello":
        if (408 > 255) {
            var v7 float32 = 0.0
            var v8 int32 = 2147483647
            v7 += -2.25
            _ = v7
            _ = v8
        } else {
            delete(v2, -300)
        }
    }
    var v9 int32 = -500
    if true {
        v9 += v9
        fmt.Println(v9, v1)
    }
    switch -449 {
    case -9223372036854775808:
        v10 := GenS3{F0: ("" + ("\n\t" + (("foo bar" + "") + ("" + ("a" + ("\n\t" + ("\n\t" + "foo bar"))))))), F1: (v9 | v9)}
        _ = v10
    case 32767:
        var v11 uint64 = 0
        _ = v11
    case -1:
        if false {
            var v12 uint = 928
            if false {
                fmt.Println(v9, v12)
            } else {
                fmt.Println(v3, v3, v1)
                if (42 == -1) {
                    var v13 int8 = -38
                    switch -51 {
                    case 32768:
                        v14 := GenS2{F0: 0, F1: ("\n\t" + (("hello" + ("hello" + "")) + "hello")), F2: 2}
                        _ = v14
                    default:
                        if false {
                            v3 = 1.5
                        } else {
                            for fi15 := 0; fi15 < 3; fi15++ {
                                v16 := GenS3{F0: "foo bar", F1: 2147483647}
                                _ = fi15
                                _ = v16
                            }
                            var v17 int32 = (-(v9))
                            _ = v17
                        }
                    }
                    var v18 uint8 = 0
                    _ = v13
                    _ = v18
                }
                v9 = (v9 - v9)
            }
            _ = v12
        }
    default:
        for fi19 := 0; fi19 < 3; fi19++ {
            v9 -= (v9 ^ -149)
            fmt.Println(v9, v9, v3)
            _ = fi19
        }
    }
    v20 := GenS1{F0: 836, F1: 1}
    _ = len(v2)
    v2[-773] = "\n\t"
    switch ("foo bar" + "hello") {
    case "foo bar":
        var v21 uint32 = 4294967295
        _ = v21
    case "a":
        v1 = (^(v1))
    default:
        for fi22 := 0; fi22 < 2; fi22++ {
            delete(v2, fi22)
            var v23 int16 = -862
            var v24 int64 = 621
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    v9 = -115
    v9 = (v9 % ((v9) | 1))
    var v25 int = 32767
    v26 := GenS2{F0: 65535, F1: "\n\t", F2: 638}
    var v27 string = "hello"
    v28 := []int{v25, v25}
    var v29 int16 = -32768
    if (!(false && (false || (65535 > 845)))) {
        fmt.Println(v29, v1, v25)
        v2[42] = "\n\t"
        if (((((1e10 - v3) - 1.5) + ((v3 / v3) / (3.14159 - ((v3 * v3) + (((-2.25 - float64(3.14159)) - v3) - 0.0))))) / float64(3.14159)) == -2.25) {
            switch v27 {
            case "\n\t":
                fmt.Println(v29, v27, v25)
            case "hello":
                var v30 float32 = ((3.14159 - 1.5) + (3.14159 + (3.14159 - 0.1)))
                _ = v30
            default:
                v31 := v2[(v25 - (v25 & 42))]
                _ = v31
            }
            v27 = (v27 + ((v27 + ("foo bar" + "foo bar")) + v27))
        }
    }
    delete(v2, v25)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v9
}
