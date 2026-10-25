package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() uint {
    v7 := []int{0, -1, 128, 961}
    var v8 uint64 = 18446744073709551615
    var v9 uint32 = 77
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int8, p1 uint16, p2 int8) uint64 {
    if true {
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 int64 = 9223372036854775807
            _ = fi10
            _ = v11
        }
        switch ("\n\t" + "hello") {
        case "":
            for fi12 := 0; fi12 < 11; fi12++ {
                p1 = p1
                _ = fi12
            }
        case "\n\t":
            if true {
                switch "hello" {
                case "foo bar":
                    p0 = p0
                default:
                    var v13 float32 = 3.14159
                    _ = v13
                }
            }
        case "hello":
            v14 := []int{934}
            _ = v14
        }
    } else {
        p0 -= (p2 * -46)
    }
    var v15 float32 = 0.1
    v16 := []int{128, 42, 646, -1}
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 2
}

func main() {
    var v1 int32 = 42
    v2 := make(map[uint32]string)
    v3 := []string{"foo bar"}
    v4 := make(map[int]bool)
    var v5 int64 = -9223372036854775808
    var v6 uint16 = 65535
    if false {
        var v17 string = "foo bar"
        var v18 int64 = (v5 ^ v5)
        _ = v17
        _ = v18
    }
    var v19 uint64 = 1
    v5 += 42
    var v20 uint8 = 255
    var v21 uint32 = 26
    _ = len(v4)
    var v22 uint64 = v19
    v22 -= v22
    switch (((("\n\t" + "a") + "foo bar") + "a") + ("" + ("hello" + ("\n\t" + "")))) {
    case "hello":
        if (4294967295 < 42) {
            var v23 uint64 = (v22 * (v19 * uint64(v6)))
            _ = v23
        } else {
            var v24 int64 = v5
            var v25 int32 = 42
            var v26 uint = 296
            _ = v24
            _ = v25
            _ = v26
        }
    case "\n\t":
        if false {
            var v27 string = ("" + "")
            var v28 int64 = (v5 - (^(v5)))
            fmt.Println(v22)
            _ = v27
            _ = v28
        } else {
            var v29 int = 2147483647
            switch int(v19) {
            case -590:
                v3 = append(v3, ("foo bar" + "\n\t"))
            case 9223372036854775807:
                v1 = 2147483647
            }
            _ = v29
        }
    }
    if len(v3) > 0 {
        v3[0] = (("" + "hello") + "hello")
    }
    if false {
        v2[2147483648] = (("hello" + ("" + "hello")) + (("foo bar" + ("hello" + "a")) + "a"))
    }
    if (42 <= 1000) {
        if true {
            var v30 bool = false
            _ = v30
        }
        if (0 != -128) {
            var v31 string = "a"
            _ = v31
        } else {
            v22 -= (v22 / ((1) | 1))
        }
        v32 := GenS1{F0: 4294967295}
        _ = v32
    } else {
        v1 += (v1 + (v1 / (((v1 / (((v1 >> ((0) & 63))) | 1))) | 1)))
    }
    v2[v21] = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
