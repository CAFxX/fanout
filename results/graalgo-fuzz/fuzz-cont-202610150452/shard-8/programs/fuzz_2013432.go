package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int16; F2 int64 }
func genFn0(p0 int16) string {
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := []int{fi6}
        _ = fi6
        _ = v7
    }
    p0 = (p0 << ((919) & 63))
    var v8 int64 = 42
    v9 := GenS1{F0: p0, F1: "hello", F2: -128}
    _ = p0
    _ = v8
    _ = v9
    return "hello"
}
func genFn1() uint32 {
    var v10 int16 = 721
    if (false && true) {
        var v11 int32 = 2147483647
        var v12 int8 = 1
        _ = v11
        _ = v12
    }
    _ = v10
    return 2147483647
}

func main() {
    var v1 int32 = -2147483648
    v2 := []string{((("a" + "foo bar") + "a") + "\n\t")}
    var v3 float64 = 3.14159
    v4 := []string{"foo bar", "a"}
    v5 := []string{((("" + "") + "foo bar") + "\n\t")}
    v13 := GenS1{F0: 914, F1: "a", F2: 100}
    switch "\n\t" {
    case "hello":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "foo bar":
        fmt.Println(v1)
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v14 uint16 = 65535
    var v15 float32 = 1e10
    v16 := GenS1{F0: 354, F1: "hello", F2: 42}
    switch 1 {
    case 32767:
        v17 := GenS2{F0: "hello"}
        _ = v17
    case -1:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        if (true && false) {
            v14 -= v14
            var v18 string = ("\n\t" + ("\n\t" + ""))
            _ = v18
        }
    }
    if (!((((1.5 / v15) != 1e10) || (-1 != 9223372036854775807)) || (251 >= -812))) {
        v15 += v15
        if (float64(v15) != (((0.1 - v3) * v3) * 1e10)) {
            v19 := GenS2{F0: ("hello" + ("a" + ""))}
            if true {
                var v20 int8 = 42
                var v21 int16 = 0
                _ = v20
                _ = v21
            } else {
                var v22 uint32 = 888
                var v23 int = 127
                _ = v22
                _ = v23
            }
            switch "" {
            case "":
                v14 -= v14
            default:
                if (!false) {
                    fmt.Println(v3, v15, v14)
                    if false {
                        v24 := GenS2{F0: "\n\t"}
                        v25 := GenS1{F0: 0, F1: "\n\t", F2: -47}
                        var v26 float64 = v3
                        _ = v24
                        _ = v25
                        _ = v26
                    } else {
                        _ = len(v2)
                        fmt.Println(v14)
                    }
                    if (0.1 < -2.25) {
                        v5 = append(v5, ("foo bar" + "foo bar"))
                        v3 = 3.14159
                    }
                }
            }
            _ = v19
        } else {
            switch 4294967295 {
            case 2147483647:
                var v27 int8 = 0
                _ = v27
            }
        }
    } else {
        if len(v5) > 0 {
            v5[0] = "\n\t"
        }
        if false {
            var v28 uint16 = (v14 % ((v14) | 1))
            _ = v28
        } else {
            fmt.Println(v15)
            var v29 string = "hello"
            switch -672 {
            case -894:
                for fi30 := 0; fi30 < 17; fi30++ {
                    for fi31 := 0; fi31 < 18; fi31++ {
                        v3 = v3
                        v1 = int32(v14)
                        v29 = (("\n\t" + "\n\t") + v29)
                        _ = fi31
                    }
                    _ = fi30
                }
            case -924:
                v32 := GenS1{F0: 544, F1: v29, F2: 112}
                _ = v32
            }
            _ = v29
        }
    }
    _ = len(v5)
    var v33 uint8 = 42
    v34 := GenS3{F0: 1, F1: 706, F2: -9223372036854775808}
    var v35 int = -1
    var v36 uint = 65535
    if len(v2) > 0 {
        v2[0] = ""
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
