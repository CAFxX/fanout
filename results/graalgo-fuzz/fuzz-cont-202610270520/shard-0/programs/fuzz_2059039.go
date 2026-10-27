package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int }
func genFn0(p0 uint16, p1 string) uint8 {
    if (0.0 <= -2.25) {
        for fi6 := 0; fi6 < 19; fi6++ {
            fmt.Println(p0, p0)
            v7 := GenS1{F0: 0, F1: 294}
            var v8 string = p1
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    var v9 int = -1
    _ = p0
    _ = p1
    _ = v9
    return 127
}

func main() {
    var v1 float64 = float64((1.5 * (1e10 + 3.14159)))
    var v2 uint = 0
    var v3 int8 = -1
    v4 := make(map[string]string)
    var v5 string = ("hello" + "foo bar")
    fmt.Println(v2)
    var v10 uint16 = 32767
    v11 := v4[(v5 + (v5 + (v5 + ("\n\t" + v5))))]
    var v12 int8 = (v3 + (v3 << ((v2) & 63)))
    if true {
        if true {
            v3 = (v12 % (((^(v3))) | 1))
            v2 = 1
        } else {
            v13 := GenS2{F0: ((v11 + (v11 + v5)) + (((("foo bar" + v5) + v11) + v5) + ((v11 + v11) + (("" + "foo bar") + (v5 + (v11 + ("a" + (v5 + "")))))))), F1: uint32(v2)}
            delete(v4, v11)
            _ = v13
        }
        fmt.Println(v3)
        v12 = (v3 % ((v3) | 1))
    }
    fmt.Println(v12, v1, v1)
    for fi14 := 0; fi14 < 11; fi14++ {
        for fi15 := 0; fi15 < 18; fi15++ {
            if (v1 == 0.1) {
                var v16 float32 = ((float32((((v1 * v1) * float64(0.0)) - (v1 / v1))) + (0.1 - (0.0 + ((0.0 + (1.5 / 1.0)) * 3.14159)))) - 1e10)
                _ = v16
            } else {
                if (895 == 837) {
                    fmt.Println(fi14, v12, v5)
                    if (v3 != 1) {
                        if false {
                            v11 = ((v11 + "hello") + v5)
                        } else {
                            v17 := []int{(fi15 & fi15), (-(fi15)), 4294967295}
                            _ = v17
                        }
                        v10 = 648
                    }
                } else {
                    var v18 int32 = 0
                    _ = v18
                }
            }
            _ = len(v4)
            if (v5 != (v11 + (v5 + "hello"))) {
                var v19 uint = v2
                v12 = v12
                var v20 int64 = 42
                v21 := v4[(v5 + v11)]
                v22 := GenS3{F0: int(v20)}
                if (1 > 2) {
                    var v23 uint8 = 128
                    _ = v23
                }
                _ = v19
                _ = v20
                _ = v21
                _ = v22
            } else {
                delete(v4, v5)
                v4[(((((v5 + v11) + "a") + "foo bar") + v5) + v11)] = v5
            }
            fmt.Println(v12, v11)
            _ = fi15
        }
        _ = fi14
    }
    v2 = v2
    fmt.Println(v1)
    for fi24 := 0; fi24 < 3; fi24++ {
        var v25 int32 = 406
        _ = fi24
        _ = v25
    }
    var v26 string = v11
    fmt.Println(v10, v5, v2)
    fmt.Println(v1, v1)
    v10 = (v10 - 51)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
