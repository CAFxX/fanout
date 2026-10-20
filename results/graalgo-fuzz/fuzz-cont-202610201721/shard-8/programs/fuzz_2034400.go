package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 uint16; F1 uint32; F2 string }
type GenS4 struct { F0 int; F1 uint }

func main() {
    var v1 uint64 = 42
    v2 := make(map[int]int)
    var v3 uint16 = 32767
    var v4 float64 = (0.0 - -2.25)
    v3 = 42
    v5 := GenS1{F0: 42}
    var v6 uint8 = 128
    v7 := GenS1{F0: -128}
    v3 = v3
    v8 := GenS2{F0: 100, F1: -2147483648}
    switch "hello" {
    case "hello":
        v9 := GenS1{F0: 92}
        _ = v9
    case "\n\t":
        v10 := GenS1{F0: 1}
        _ = v10
    }
    v11 := GenS3{F0: (v3 >> ((783) & 63)), F1: 2, F2: "foo bar"}
    var v12 uint32 = 2147483647
    if (!(!true)) {
        if ((!(("" > ("foo bar" + "\n\t")) || true)) || (!false)) {
            fmt.Println(v4, v12)
            var v13 int64 = -840
            _ = v13
        } else {
            var v14 string = ("foo bar" + "\n\t")
            if false {
                var v15 float64 = 1.5
                if (v14 <= v14) {
                    switch 444 {
                    case 42:
                        v1 -= v1
                    case 375:
                        fmt.Println(v3, v14, v12)
                    case -894:
                        v12 -= 2147483647
                    }
                } else {
                    _ = len(v2)
                    fmt.Println(v12, v4, v1)
                    delete(v2, 329)
                }
                fmt.Println(v12)
                _ = v15
            }
            _ = v14
        }
        switch "foo bar" {
        case "":
            v12 += 507
        case "hello":
            v16 := []int{}
            _ = v16
        default:
            fmt.Println(v3)
        }
    } else {
        v17 := []int{542}
        v12 = 212
        var v18 float64 = (1.5 + 1e10)
        _ = v17
        _ = v18
    }
    var v19 uint8 = 42
    for fi20 := 0; fi20 < 1; fi20++ {
        fmt.Println(v1, v1, v19)
        for fi21 := 0; fi21 < 18; fi21++ {
            v22 := GenS4{F0: fi21, F1: 255}
            v6 += 1
            var v23 float32 = 1e10
            _ = fi21
            _ = v22
            _ = v23
        }
        for fi24 := 0; fi24 < 3; fi24++ {
            v3 = v3
            switch ("\n\t" + ("a" + (("" + (("" + "hello") + "\n\t")) + "foo bar"))) {
            case "":
                var v25 uint = uint(v12)
                _ = v25
            }
            switch (fi20 - fi20) {
            case 9223372036854775807:
                v6 -= 0
            }
            _ = fi24
        }
        _ = fi20
    }
    if true {
        v26 := []int{9223372036854775807, 128, -1}
        _ = v26
    } else {
        var v27 float32 = 3.14159
        _ = v27
    }
    v28 := []int{-1}
    var v29 int64 = -9223372036854775808
    delete(v2, 127)
    v30 := GenS2{F0: -18, F1: 941}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
