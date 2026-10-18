package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int; F2 string }
type GenS3 struct { F0 uint32; F1 uint; F2 uint }

func main() {
    var v1 uint = 1
    var v2 int = 9223372036854775807
    v3 := []int{v2, v2, (v2 - -916), (-(v2))}
    v4 := make(map[uint32]int)
    v5 := make(map[int]bool)
    v2 -= v2
    if true {
        switch "" {
        case "hello":
            var v6 uint = (^(v1))
            _ = v6
        case "foo bar":
            v2 = v2
        case "":
            fmt.Println(v2)
        }
    } else {
        for fi7 := 0; fi7 < 16; fi7++ {
            if false {
                v8 := GenS1{F0: true}
                _ = v8
            } else {
                v2 = 1
                var v9 uint = v1
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v9
            }
            v10 := GenS1{F0: (!(42 >= 199))}
            v11 := GenS1{F0: (!(59 >= 1))}
            _ = fi7
            _ = v10
            _ = v11
        }
        v12 := GenS1{F0: ((v1 & (v1 / (((v1 & v1)) | 1))) <= v1)}
        var v13 uint64 = 2
        _ = v12
        _ = v13
    }
    if (v2 <= v2) {
        v2 -= 2147483647
        var v14 int64 = 0
        _ = v14
    }
    fmt.Println(v2, v1)
    delete(v4, 2)
    v15 := GenS2{F0: 32767, F1: v2, F2: ("foo bar" + "a")}
    fmt.Println(v1, v1)
    v16 := GenS2{F0: -1, F1: v2, F2: ""}
    if len(v3) > 0 {
        v3[0] = (v2 ^ (v2 * (v2 - v2)))
    }
    var v17 int = v2
    var v18 int64 = -9223372036854775808
    switch (v17 ^ v17) {
    case 0:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2, v1)
    v19 := GenS2{F0: 32767, F1: v2, F2: (("" + ((("" + "foo bar") + "hello") + ("hello" + (("hello" + "foo bar") + "foo bar")))) + "a")}
    if (((false && false) || true) || ((0.1 + (1.5 / 1.5)) == 1e10)) {
        v20 := GenS3{F0: 925, F1: (v1 * v1), F2: v1}
        _ = len(v5)
        fmt.Println(v1, v17, v1)
        _ = v20
    }
    v18 = (v18 * (v18 % ((v18) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
