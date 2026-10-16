package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS4 struct { F0 int; F1 int; F2 uint16 }
func genFn0(p0 int8, p1 int32) uint32 {
    p0 += p0
    v5 := GenS1{F0: 1, F1: (!false), F2: "foo bar"}
    v6 := GenS1{F0: -689, F1: true, F2: "\n\t"}
    p0 += 0
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 string) int32 {
    var v7 string = p0
    if (false && (false || ((("foo bar" + "foo bar") + (v7 + p0)) <= ""))) {
        v8 := []int{32768}
        _ = v8
    }
    if true {
        var v9 int64 = 9223372036854775807
        v10 := GenS1{F0: -951, F1: (733 >= -32768), F2: v7}
        _ = v10
        _ = v9
    } else {
        var v11 int64 = 1
        _ = v11
    }
    if (!true) {
        v12 := GenS1{F0: -234, F1: false, F2: ((v7 + v7) + v7)}
        _ = v12
    } else {
        var v13 uint8 = 128
        var v14 bool = false
        p0 = ("a" + p0)
        _ = v13
        _ = v14
    }
    _ = p0
    _ = v7
    return -1
}

func main() {
    v1 := []string{("hello" + ("a" + ""))}
    v2 := []int{2147483647, 128}
    var v3 float64 = 0.1
    v4 := []string{"hello", "hello", (((("\n\t" + ("" + "\n\t")) + "a") + ("" + (("\n\t" + ("foo bar" + ((("" + "foo bar") + ("" + "")) + (("foo bar" + "") + ("" + ("hello" + "foo bar")))))) + "a"))) + "hello"), ""}
    var v15 int8 = 119
    for fi16 := 0; fi16 < 3; fi16++ {
        fmt.Println(v3, v15)
        fmt.Println(fi16, v3)
        switch "foo bar" {
        case "":
            for fi17 := 0; fi17 < 20; fi17++ {
                var v18 float32 = -2.25
                if false {
                    switch (fi16 * fi16) {
                    case 9223372036854775807:
                        var v19 uint32 = 48
                        _ = v19
                    }
                }
                var v20 uint32 = genFn0((v15 >> ((406) & 63)), genFn1("foo bar"))
                _ = fi17
                _ = v18
                _ = v20
            }
        case "foo bar":
            for fi21 := 0; fi21 < 12; fi21++ {
                var v22 uint64 = 18446744073709551615
                v22 = 485
                _ = fi21
                _ = v22
            }
        default:
            for fi23 := 0; fi23 < 15; fi23++ {
                var v24 int32 = -2147483648
                _ = len(v4)
                var v25 uint8 = 1
                _ = fi23
                _ = v24
                _ = v25
            }
        }
        _ = fi16
    }
    v26 := GenS2{F0: 2147483647}
    var v27 int32 = 0
    switch "foo bar" {
    case "":
        var v28 int64 = 42
        _ = v28
    case "\n\t":
        v27 += 1
    }
    switch -639 {
    case -2147483648:
        v29 := GenS3{F0: 297, F1: 449, F2: genFn1("\n\t")}
        _ = v29
    case 127:
        v30 := GenS3{F0: 83, F1: 32767, F2: 2147483647}
        _ = v30
    case -1:
        fmt.Println(v15, v27)
    }
    fmt.Println(v3, v27)
    var v31 uint16 = 2
    for fi32 := 0; fi32 < 17; fi32++ {
        var v33 bool = true
        _ = fi32
        _ = v33
    }
    var v34 uint = 255
    fmt.Println(v27, v3, v3)
    v35 := GenS4{F0: 1, F1: 32768, F2: v31}
    v34 = v34
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v4
}
