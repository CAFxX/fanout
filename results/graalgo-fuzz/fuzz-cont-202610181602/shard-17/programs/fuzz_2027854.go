package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8 }
func genFn0() string {
    var v5 uint16 = 65535
    v5 += (v5 / (((v5 & (^(v5)))) | 1))
    v5 = (v5 / ((0) | 1))
    _ = v5
    return (("foo bar" + "a") + "hello")
}

func main() {
    var v1 int8 = 23
    v2 := make(map[int]string)
    v3 := []int{9223372036854775807}
    v4 := []string{"", "\n\t", "a", "foo bar"}
    switch -443 {
    case -896:
        v6 := GenS1{F0: false, F1: 1}
        _ = v6
    case 128:
        for fi7 := 0; fi7 < 2; fi7++ {
            _ = len(v2)
            _ = fi7
        }
    default:
        v1 = v1
    }
    var v8 int16 = 0
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) > 0 {
        v4[0] = ((((("a" + "foo bar") + (("hello" + ("hello" + "")) + ("" + ((("" + ("foo bar" + "foo bar")) + "a") + (("\n\t" + (("\n\t" + ("foo bar" + "\n\t")) + ("foo bar" + ("\n\t" + "foo bar")))) + ("a" + "")))))) + "") + "hello") + "\n\t")
    }
    var v9 float32 = float32(float64(3.14159))
    var v10 uint32 = 65535
    if false {
        if (2 <= 50) {
            fmt.Println(v8, v1, v10)
        }
        var v11 uint16 = 0
        var v12 uint8 = 42
        _ = v11
        _ = v12
    }
    if (42 <= 1) {
        if false {
            if (!(false || (!(-2.25 != 3.14159)))) {
                if true {
                    for fi13 := 0; fi13 < 16; fi13++ {
                        v1 += (v1 | v1)
                        _ = fi13
                    }
                    v4 = append(v4, "\n\t")
                    v14 := GenS2{F0: (float32(float64(v9)) > ((-2.25 - v9) * v9))}
                    _ = v14
                } else {
                    v15 := GenS3{F0: 1}
                    var v16 int = 4294967295
                    _ = v15
                    _ = v16
                }
            }
            v17 := GenS3{F0: -781}
            switch 4294967295 {
            case -9223372036854775808:
                var v18 uint64 = 299
                _ = v18
            case 4294967295:
                for fi19 := 0; fi19 < 14; fi19++ {
                    switch fi19 {
                    case -537:
                        v20 := GenS3{F0: fi19}
                        _ = v20
                    }
                    if true {
                        v8 = v8
                    }
                    _ = fi19
                }
            default:
                fmt.Println(v9)
            }
            _ = v17
        }
        for fi21 := 0; fi21 < 19; fi21++ {
            var v22 bool = (true && false)
            v23 := GenS4{F0: 18446744073709551615}
            _ = len(v3)
            _ = fi21
            _ = v22
            _ = v23
        }
        v8 -= v8
    }
    fmt.Println(v8, v10, v8)
    _ = len(v2)
    var v24 uint16 = 65535
    for fi25 := 0; fi25 < 8; fi25++ {
        v2[(fi25 << ((uint(v24)) & 63))] = "foo bar"
        var v26 uint = 4294967295
        for fi27 := 0; fi27 < 12; fi27++ {
            v2[fi27] = "a"
            var v28 uint8 = 250
            _ = fi27
            _ = v28
        }
        _ = fi25
        _ = v26
    }
    if false {
        v9 = (v9 + ((1e10 / float32(float64(v9))) * (float32((1e10 - 1e10)) * float32(1e10))))
        v29 := GenS3{F0: 127}
        switch 42 {
        case 9223372036854775807:
            var v30 int = 4294967295
            _ = v30
        case 127:
            v31 := GenS1{F0: true, F1: v8}
            _ = v31
        case 1:
            v1 += (v1 - (v1 / (((v1 ^ 30)) | 1)))
        }
        _ = v29
    } else {
        if ((("" == "a") && true) || (uint8(v10) <= 1)) {
            v1 = 1
            v24 += 0
            var v32 float64 = ((1.5 + ((1e10 / (3.14159 + 3.14159)) + 3.14159)) * -2.25)
            _ = v32
        } else {
            v1 -= 28
            v9 += (0.0 + (0.0 + (float32((1.5 + (0.1 + 1.5))) * v9)))
            switch "foo bar" {
            case "a":
                if (!false) {
                    v10 -= (v10 & 42)
                    v3 = append(v3, 4294967295)
                } else {
                    v33 := GenS5{F0: (v1 + v1)}
                    _ = v33
                }
            case "hello":
                _ = len(v2)
            }
        }
    }
    if true {
        var v34 int64 = 1
        for fi35 := 0; fi35 < 16; fi35++ {
            var v36 int64 = -9223372036854775808
            _ = fi35
            _ = v36
        }
        v37 := GenS3{F0: 521}
        _ = v34
        _ = v37
    }
    v24 -= (v24 | (v24 + (v24 >> ((4) & 63))))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
