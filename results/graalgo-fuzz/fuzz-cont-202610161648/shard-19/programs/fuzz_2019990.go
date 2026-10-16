package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 int64 }
func genFn0(p0 int) uint32 {
    v4 := GenS1{F0: ("hello" + "hello")}
    var v5 uint8 = 2
    var v6 int8 = 0
    fmt.Println(v5, v5)
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 965
}

func main() {
    var v1 string = "a"
    var v2 bool = true
    var v3 int = 967
    fmt.Println(v3)
    var v7 uint16 = 42
    v8 := []int{(v3 * (v3 ^ (v3 << ((595) & 63))))}
    v7 -= v7
    if true {
        if v2 {
            fmt.Println(v3, v1)
            if v2 {
                fmt.Println(v7)
                switch "a" {
                case "":
                    var v9 int8 = 42
                    _ = v9
                case "hello":
                    fmt.Println(v7, v1, v1)
                case "a":
                    fmt.Println(v7, v7, v1)
                }
                v7 = 65535
            } else {
                var v10 uint32 = 2147483648
                var v11 float64 = 1.5
                v1 = "a"
                _ = v10
                _ = v11
            }
        }
    }
    var v12 uint = 42
    if (1 > 0) {
        switch v1 {
        case "":
            v8 = append(v8, v3)
        case "a":
            v3 = v3
        }
        v3 = v3
        switch (v3 - 4294967295) {
        case -9223372036854775808:
            v13 := GenS1{F0: v1}
            _ = v13
        }
    } else {
        v14 := GenS2{F0: v3}
        for fi15 := 0; fi15 < 15; fi15++ {
            for fi16 := 0; fi16 < 13; fi16++ {
                v7 = (v7 ^ v7)
                v17 := GenS3{F0: v12, F1: ("foo bar" + v1)}
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
        if len(v8) > 0 {
            v8[0] = 0
        }
        _ = v14
    }
    v1 = ("a" + (v1 + v1))
    switch ("foo bar" + "a") {
    case "\n\t":
        if (609 != 42) {
            v1 = (v1 + (v1 + (v1 + v1)))
        } else {
            v18 := GenS4{F0: 466}
            var v19 float32 = (0.0 + 1.5)
            if (!false) {
                for fi20 := 0; fi20 < 17; fi20++ {
                    var v21 uint16 = v7
                    _ = fi20
                    _ = v21
                }
                v2 = v2
            } else {
                fmt.Println(v1, v19)
                v2 = (307 > 2)
                var v22 uint16 = (^(v7))
                _ = v22
            }
            _ = v18
            _ = v19
        }
    }
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
