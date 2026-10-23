package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 uint8; F2 uint8 }
func genFn0(p0 int) int32 {
    var v4 uint16 = 2
    var v5 uint8 = 39
    v6 := GenS1{F0: p0, F1: 18446744073709551615}
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS2{F0: 350}
        _ = fi7
        _ = v8
    }
    switch ((("" + (("a" + "hello") + "hello")) + "\n\t") + "hello") {
    case "hello":
        v9 := GenS1{F0: p0, F1: 42}
        _ = v9
    case "foo bar":
        switch "a" {
        case "\n\t":
            v10 := []int{0, (p0 + 127), p0}
            _ = v10
        case "":
            v11 := []int{(p0 % (((p0 | p0)) | 1)), (^(p0))}
            _ = v11
        case "a":
            switch (("foo bar" + "") + "") {
            case "foo bar":
                v12 := []int{}
                _ = v12
            case "\n\t":
                if true {
                    v13 := []int{-2147483648}
                    _ = v13
                }
            case "a":
                var v14 int32 = 42
                _ = v14
            default:
                switch (((("foo bar" + "\n\t") + "a") + "\n\t") + "") {
                case "foo bar":
                    v15 := GenS3{F0: int64(v4)}
                    _ = v15
                }
            }
        }
    default:
        var v16 int = (p0 * p0)
        _ = v16
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1(p0 int) uint16 {
    if (true && true) {
        fmt.Println(p0)
    }
    p0 = (p0 | (p0 | (p0 & p0)))
    v17 := []int{p0}
    p0 = p0
    p0 = p0
    _ = p0
    _ = v17
    return 32767
}

func main() {
    var v1 float32 = ((1e10 + 0.1) / (1e10 / 3.14159))
    v2 := make(map[string]string)
    v3 := []string{}
    var v18 int16 = -32768
    delete(v2, "foo bar")
    v19 := v2["hello"]
    var v20 float64 = 3.14159
    var v21 int32 = 2147483647
    v22 := GenS1{F0: 144, F1: 950}
    v19 = v19
    if len(v3) > 0 {
        v3[0] = v19
    }
    switch v19 {
    case "":
        if len(v3) > 0 {
            v3[0] = v19
        }
    default:
        v1 = 3.14159
    }
    if (42 > 117) {
        v1 = 1.5
        v21 = v21
    } else {
        if (127 > 9223372036854775807) {
            v23 := GenS4{F0: 1, F1: 128, F2: 1}
            v18 = v18
            if true {
                if false {
                    if (42 != 42) {
                        if (("" <= (((v19 + ((((("foo bar" + ("foo bar" + v19)) + ((v19 + "foo bar") + v19)) + (("foo bar" + v19) + v19)) + v19) + "foo bar")) + "hello") + v19)) || (-32768 != 653)) {
                            var v24 uint64 = 2
                            _ = v24
                        }
                    } else {
                        fmt.Println(v19, v21, v18)
                    }
                    v20 += float64(1e10)
                }
            }
            _ = v23
        } else {
            v18 -= 1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
