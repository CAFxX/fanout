package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int8 }
func genFn0(p0 uint8) uint16 {
    if (649 <= 253) {
        v5 := GenS1{F0: "foo bar", F1: "hello", F2: -9223372036854775808}
        if ((42 <= 18446744073709551615) || (-128 != 42)) {
            v6 := []int{-964, 4294967295}
            fmt.Println(p0)
            _ = v6
        }
        v7 := []int{-2147483648, -2147483648}
        _ = v5
        _ = v7
    } else {
        var v8 uint8 = 228
        _ = v8
    }
    if true {
        p0 += p0
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS1{F0: "", F1: "a", F2: 982}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        if ((3.14159 / 1.5) < 0.0) {
            if (1e10 > 0.1) {
                if false {
                    p0 -= (p0 * (p0 << ((2) & 63)))
                    var v12 int32 = 1
                    _ = v12
                } else {
                    p0 = (p0 | 0)
                }
            } else {
                for fi13 := 0; fi13 < 20; fi13++ {
                    var v14 uint64 = 2
                    fmt.Println(fi11, p0)
                    if (float64(0.1) <= (float64(0.0) * 3.14159)) {
                        var v15 uint16 = 1
                        v16 := GenS1{F0: "foo bar", F1: "foo bar", F2: 9223372036854775807}
                        var v17 uint8 = (^(p0))
                        _ = v15
                        _ = v16
                        _ = v17
                    } else {
                        v14 = (v14 >> ((550) & 63))
                        v18 := []int{fi13, fi11, fi11}
                        _ = v18
                    }
                    _ = fi13
                    _ = v14
                }
            }
            v19 := []int{0, (fi11 ^ fi11)}
            _ = v19
        }
        p0 += (p0 / ((p0) | 1))
        _ = fi11
    }
    p0 = (p0 ^ 93)
    _ = p0
    return 73
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 3.14159
    v3 := make(map[string]bool)
    v4 := make(map[int]bool)
    v2 = v2
    v2 = ((3.14159 + 0.0) * v2)
    v20 := GenS2{F0: genFn0(254)}
    v2 = (v2 / (((v2 - (3.14159 + 0.0)) + v2) / v2))
    for fi21 := 0; fi21 < 19; fi21++ {
        if (false || false) {
            v2 += v2
            v2 = 1e10
        } else {
            if (true && false) {
                var v22 float64 = (1.5 * ((1e10 - 0.0) + (-2.25 - ((((1.5 / 3.14159) - 1e10) - (3.14159 / ((0.0 - (1e10 * 1e10)) / 3.14159))) - 0.1))))
                _ = v22
            } else {
                if false {
                    var v23 float32 = v2
                    if false {
                        fmt.Println(v23)
                        v3["a"] = false
                    }
                    _ = v23
                } else {
                    v24 := []int{fi21, 1, (fi21 | 32768)}
                    _ = v24
                }
                fmt.Println(fi21, fi21)
            }
        }
        v2 += 3.14159
        _ = fi21
    }
    var v25 int16 = 32767
    fmt.Println(v2, v2)
    v26 := GenS3{F0: 582, F1: 0}
    var v27 string = (("foo bar" + ("\n\t" + "")) + "foo bar")
    v27 = (((v27 + ("foo bar" + ((v27 + v27) + v27))) + (v27 + ("foo bar" + ((((((("hello" + v27) + "foo bar") + ((v27 + (("\n\t" + v27) + ("\n\t" + (((v27 + (v27 + v27)) + ((v27 + (v27 + ("a" + v27))) + ((((((v27 + v27) + "\n\t") + v27) + (((((v27 + (("a" + v27) + v27)) + ((v27 + (v27 + ("\n\t" + v27))) + v27)) + "foo bar") + (v27 + v27)) + "foo bar")) + "\n\t") + v27))) + v27)))) + (v27 + "foo bar"))) + v27) + ((v27 + "hello") + (("hello" + v27) + ""))) + v27) + v27)))) + (v27 + v27))
    var v28 string = "foo bar"
    switch ((v27 + v27) + "hello") {
    case "a":
        v25 = 42
    case "":
        fmt.Println(v28)
    }
    if false {
        var v29 uint = 255
        v30 := v4[-9223372036854775808]
        _ = v29
        _ = v30
    }
    if false {
        for fi31 := 0; fi31 < 4; fi31++ {
            for fi32 := 0; fi32 < 14; fi32++ {
                fmt.Println(fi31, v25, v27)
                _ = len(v4)
                v33 := v1["a"]
                var v34 float32 = v2
                _ = fi32
                _ = v33
                _ = v34
            }
            _ = fi31
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
