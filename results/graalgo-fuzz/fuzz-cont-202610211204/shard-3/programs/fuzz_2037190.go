package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint }
type GenS2 struct { F0 uint; F1 uint8; F2 int }
type GenS3 struct { F0 int16; F1 uint }

func main() {
    var v1 uint = 42
    v2 := []string{"foo bar", "hello", "a"}
    var v3 int32 = 42
    var v4 uint8 = 127
    v5 := make(map[int]bool)
    fmt.Println(v4, v4)
    v1 = v1
    v6 := GenS1{F0: v4, F1: ("foo bar" + ("a" + "hello")), F2: (v1 + v1)}
    for fi7 := 0; fi7 < 3; fi7++ {
        v3 = v3
        v3 += (v3 | 42)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi7
    }
    if false {
        if ((265 == (-(v3))) || false) {
            v8 := GenS1{F0: 12, F1: (("foo bar" + "\n\t") + ("" + "a")), F2: v1}
            _ = v8
        } else {
            for fi9 := 0; fi9 < 1; fi9++ {
                fmt.Println(fi9)
                for fi10 := 0; fi10 < 8; fi10++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v11 := GenS1{F0: v4, F1: "\n\t", F2: 1}
                    if true {
                        var v12 int64 = 0
                        v13 := GenS2{F0: (v1 & (v1 + (v1 / ((v1) | 1)))), F1: v4, F2: fi9}
                        _ = v12
                        _ = v13
                    }
                    _ = fi10
                    _ = v11
                }
                _ = fi9
            }
            _ = len(v5)
        }
    } else {
        v2 = append(v2, "\n\t")
        v14 := GenS1{F0: (v4 | v4), F1: ("a" + ""), F2: 18446744073709551615}
        _ = v14
    }
    switch 4294967295 {
    case -626:
        if (!true) {
            fmt.Println(v3)
            var v15 int8 = 100
            var v16 uint8 = 49
            _ = v15
            _ = v16
        }
    }
    var v17 uint16 = 25
    fmt.Println(v3, v17)
    var v18 int = 128
    v19 := GenS3{F0: 32767, F1: 4294967295}
    v1 -= (v1 & 978)
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 uint16 = 200
        _ = fi20
        _ = v21
    }
    _ = len(v5)
    if (0 >= 0) {
        v2 = append(v2, "foo bar")
    } else {
        for fi22 := 0; fi22 < 20; fi22++ {
            var v23 int = v18
            var v24 int32 = v3
            if false {
                var v25 uint16 = (v17 / ((1) | 1))
                var v26 float64 = 3.14159
                _ = v25
                _ = v26
            }
            _ = fi22
            _ = v23
            _ = v24
        }
        v18 = (v18 / (((v18 - (v18 * (v18 / ((int(v4)) | 1))))) | 1))
    }
    for fi27 := 0; fi27 < 7; fi27++ {
        var v28 int16 = int16(v1)
        var v29 float32 = ((-2.25 / 3.14159) + (3.14159 + -2.25))
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
