package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int }
func genFn0() int64 {
    switch "foo bar" {
    case "foo bar":
        switch 439 {
        case 32767:
            var v7 int16 = -1
            _ = v7
        case -9223372036854775808:
            var v8 uint = 18446744073709551615
            _ = v8
        }
    default:
        var v9 int16 = 32767
        _ = v9
    }
    var v10 int16 = 1
    fmt.Println(v10)
    if false {
        v10 += v10
        v11 := GenS1{F0: 472, F1: 881}
        _ = v11
    }
    _ = v10
    return 42
}
func genFn1() int32 {
    var v12 float64 = 1e10
    v13 := GenS1{F0: 18446744073709551615, F1: 355}
    if ((!true) && ((-2147483648 != 1) && false)) {
        v12 -= float64(float32(3.14159))
        v12 = ((v12 * v12) * v12)
        v14 := []int{-2147483648}
        _ = v14
    } else {
        if false {
            if ((!(-156 > -79)) && (255 <= 42)) {
                switch 32768 {
                case -1:
                    var v15 float32 = 3.14159
                    _ = v15
                case 32:
                    v16 := GenS1{F0: 42, F1: 255}
                    _ = v16
                case 128:
                    v17 := []int{780, 127, 561}
                    _ = v17
                }
                v18 := GenS1{F0: 18446744073709551615, F1: 620}
                fmt.Println(v12)
                _ = v18
            } else {
                var v19 uint8 = 0
                _ = v19
            }
            v20 := GenS2{F0: 18446744073709551615}
            _ = v20
        } else {
            var v21 int16 = 1
            _ = v21
        }
        v12 += v12
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        v12 += float64((1.5 + 3.14159))
        _ = fi22
    }
    _ = v12
    _ = v13
    return 1
}

func main() {
    var v1 int8 = 0
    v2 := []string{"a", "foo bar", "foo bar"}
    var v3 uint16 = 51
    v4 := make(map[uint32]int)
    var v5 string = (("" + "hello") + "\n\t")
    var v6 int = 0
    var v23 float64 = float64(3.14159)
    if (v6 >= -2147483648) {
        fmt.Println(v5)
        v24 := GenS1{F0: 144, F1: 0}
        _ = v24
    }
    switch 4294967295 {
    case 0:
        v1 -= 42
    }
    var v25 int16 = -24
    fmt.Println(v23)
    var v26 string = ""
    var v27 bool = ((v6 << ((2) & 63)) == v6)
    v3 += v3
    var v28 uint64 = 199
    var v29 uint32 = 0
    v4[v29] = -2147483648
    if (0.0 >= 1e10) {
        var v30 int32 = genFn1()
        _ = v30
    } else {
        v4[v29] = (-(v6))
        if (!((v29 / (((^(v29))) | 1)) >= (v29 * 4294967295))) {
            v31 := GenS3{F0: v25, F1: (^(v28))}
            if false {
                var v32 int8 = v1
                v5 = ((((v26 + (("foo bar" + (v5 + v5)) + (v5 + v26))) + v5) + v26) + ((("\n\t" + v5) + "") + ("" + "\n\t")))
                if true {
                    switch v5 {
                    case "foo bar":
                        v33 := GenS3{F0: (v25 % ((v25) | 1)), F1: uint64(v3)}
                        _ = v33
                    default:
                        for fi34 := 0; fi34 < 3; fi34++ {
                            v35 := GenS2{F0: uint64(v3)}
                            v36 := GenS2{F0: 4294967295}
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                    }
                } else {
                    fmt.Println(v26, v6)
                    switch -93 {
                    case 1:
                        fmt.Println(v32)
                    default:
                        var v37 int = (-(v6))
                        _ = v37
                    }
                }
                _ = v32
            }
            var v38 int16 = v25
            _ = v31
            _ = v38
        }
    }
    var v39 bool = (v27 || (!true))
    switch ((((("" + v5) + v5) + "") + v5) + v26) {
    case "hello":
        v40 := GenS4{F0: -2147483648}
        _ = v40
    case "":
        v39 = (!v27)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
