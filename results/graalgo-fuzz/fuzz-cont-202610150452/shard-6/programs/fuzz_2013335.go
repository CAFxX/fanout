package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint, p1 bool) bool {
    fmt.Println(p1, p1)
    p1 = ((("a" >= "a") || (-1 >= 0)) || ((42 < -747) || (0.0 == -2.25)))
    if (true && p1) {
        for fi5 := 0; fi5 < 2; fi5++ {
            for fi6 := 0; fi6 < 18; fi6++ {
                var v7 float32 = 1e10
                v8 := GenS1{F0: "hello", F1: 2147483647, F2: -861}
                v9 := []int{0, 128, 32768, -1}
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
            if true {
                for fi10 := 0; fi10 < 2; fi10++ {
                    var v11 float64 = 0.1
                    v12 := []int{}
                    v11 += float64(-2.25)
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            } else {
                if (!(p1 && false)) {
                    fmt.Println(p1, p1)
                    v13 := []int{fi5, fi5, (fi5 << ((p0) & 63)), 32768}
                    var v14 uint = p0
                    _ = v13
                    _ = v14
                }
                for fi15 := 0; fi15 < 20; fi15++ {
                    switch "a" {
                    case "":
                        var v16 bool = (9223372036854775807 == 42)
                        _ = v16
                    default:
                        var v17 uint8 = 255
                        _ = v17
                    }
                    if ((("foo bar" + "hello") <= (("foo bar" + "") + "foo bar")) || ((fi5 | (fi15 - (fi15 * 32767))) < 128)) {
                        switch "\n\t" {
                        case "":
                            v18 := []int{}
                            _ = v18
                        case "a":
                            var v19 float32 = (-2.25 * (0.0 + 3.14159))
                            _ = v19
                        case "\n\t":
                            fmt.Println(fi5, p1)
                        default:
                            var v20 uint16 = 42
                            _ = v20
                        }
                    }
                    if ((660 <= -1) || (p1 && p1)) {
                        v21 := []int{(^(fi5)), (-(fi15)), fi5}
                        _ = v21
                    }
                    _ = fi15
                }
            }
            _ = fi5
        }
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 int32 = -2147483648
        _ = fi22
        _ = v23
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 float32 = -2.25
    var v2 int16 = -1
    var v3 uint16 = 255
    var v4 bool = ((v3 - (^(v3))) >= v3)
    var v24 int16 = 42
    v24 += 32767
    v25 := GenS2{F0: 65535}
    v26 := GenS2{F0: v3}
    if v4 {
        v27 := []int{-22}
        _ = v27
    } else {
        var v28 uint16 = (v3 ^ (v3 / (((v3 % ((100) | 1))) | 1)))
        v29 := GenS2{F0: v28}
        _ = v28
        _ = v29
    }
    v3 -= (v3 & 42)
    v1 = v1
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
