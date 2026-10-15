package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int, p1 uint16) int32 {
    var v7 string = "a"
    var v8 bool = (!(!(((p1 - p1) > p1) || (-6 <= -128))))
    if ((0.1 / 0.1) != (-2.25 / (float32(float64(float32(float64(float32(0.0))))) + 1e10))) {
        if true {
            v9 := GenS1{F0: v7, F1: 42}
            v10 := []int{0, 0, 32768, 32767}
            _ = v10
            _ = v9
        }
    }
    var v11 uint16 = p1
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 string) int16 {
    switch p0 {
    case "a":
        if true {
            fmt.Println(p0)
            for fi12 := 0; fi12 < 7; fi12++ {
                v13 := GenS1{F0: p0, F1: 2147483647}
                v14 := []int{(fi12 | fi12)}
                _ = fi12
                _ = v13
                _ = v14
            }
            if (true && (false && (4294967295 == 1))) {
                switch p0 {
                case "foo bar":
                    var v15 int64 = 1
                    _ = v15
                }
            } else {
                var v16 int32 = 0
                v17 := []int{4294967295, -9223372036854775808, 44}
                _ = v16
                _ = v17
            }
        } else {
            v18 := []int{128, -516, -9223372036854775808}
            p0 = p0
            _ = v18
        }
    case "foo bar":
        p0 = "\n\t"
    case "hello":
        fmt.Println(p0)
    default:
        switch (p0 + "\n\t") {
        case "\n\t":
            var v19 uint64 = 4294967295
            _ = v19
        }
    }
    p0 = ((p0 + p0) + "\n\t")
    _ = p0
    return -598
}

func main() {
    var v1 string = ""
    var v2 float32 = 3.14159
    v3 := make(map[string]int)
    var v4 int16 = -1
    var v5 string = ("hello" + v1)
    var v6 uint16 = 65535
    v6 = 65535
    v4 -= v4
    v5 = ("" + "")
    v2 = v2
    v2 = ((((v2 - -2.25) - (v2 + v2)) / v2) - (v2 - (0.1 * 0.0)))
    switch -766 {
    case 42:
        fmt.Println(v6)
    case 127:
        v5 = ("foo bar" + (v1 + (((v1 + (v1 + v5)) + "") + v1)))
    default:
        switch "" {
        case "a":
            v1 = v1
        }
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        for fi21 := 0; fi21 < 3; fi21++ {
            v22 := GenS2{F0: 4294967295}
            v6 -= v6
            var v23 int8 = 0
            v2 += ((v2 / 1e10) + 0.1)
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi20
    }
    if (!(v5 > ("a" + v1))) {
        fmt.Println(v4)
    }
    v2 = v2
    for fi24 := 0; fi24 < 6; fi24++ {
        if false {
            var v25 float64 = 3.14159
            var v26 float64 = v25
            _ = v25
            _ = v26
        }
        if true {
            var v27 int = (fi24 / (((-(fi24))) | 1))
            _ = v27
        }
        _ = fi24
    }
    fmt.Println(v5)
    v2 += ((v2 / -2.25) / (((v2 + 3.14159) - v2) + v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
