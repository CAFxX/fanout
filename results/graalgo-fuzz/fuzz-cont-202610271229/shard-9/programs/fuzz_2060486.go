package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 int32 }
func genFn0(p0 int32) string {
    var v4 bool = (-1 != -1)
    v5 := []int{}
    fmt.Println(p0, v4)
    fmt.Println(v4)
    if (p0 >= (p0 ^ (p0 & (p0 + p0)))) {
        v6 := GenS1{F0: 1, F1: 472}
        _ = v6
    }
    _ = p0
    _ = v4
    _ = v5
    return "\n\t"
}
func genFn1() uint32 {
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint64 = 18446744073709551615
        _ = fi7
        _ = v8
    }
    if true {
        v9 := []int{1, -371, 0, -788}
        _ = v9
    } else {
        if false {
            for fi10 := 0; fi10 < 8; fi10++ {
                v11 := GenS2{F0: 1000}
                fmt.Println(fi10)
                _ = fi10
                _ = v11
            }
        } else {
            var v12 int16 = 0
            var v13 int32 = -649
            _ = v12
            _ = v13
        }
    }
    switch 32767 {
    case -9223372036854775808:
        switch "foo bar" {
        case "\n\t":
            var v14 string = "\n\t"
            _ = v14
        case "a":
            var v15 uint64 = 2
            _ = v15
        case "hello":
            switch "" {
            case "":
                var v16 uint = 255
                _ = v16
            case "\n\t":
                var v17 int8 = -98
                _ = v17
            case "foo bar":
                var v18 string = "foo bar"
                _ = v18
            }
        }
    case 0:
        v19 := []int{2147483647, 681, 32768, 32767}
        _ = v19
    default:
        switch ((("a" + ("hello" + "\n\t")) + ("foo bar" + (("" + "hello") + ("\n\t" + "hello")))) + "foo bar") {
        case "":
            v20 := []int{}
            _ = v20
        case "a":
            var v21 int8 = 0
            _ = v21
        }
    }
    return 42
}

func main() {
    var v1 uint16 = 1
    v2 := []int{1}
    v3 := []string{("a" + "\n\t"), ("a" + "foo bar"), (("hello" + "foo bar") + ("" + "")), ("foo bar" + "hello")}
    fmt.Println(v1)
    var v22 int16 = -740
    if len(v2) > 0 {
        v2[0] = 128
    }
    if len(v2) > 0 {
        v2[0] = 1
    }
    if ((^(v1)) > (^(v1))) {
        fmt.Println(v22, v1)
        v23 := GenS2{F0: (v22 << ((2) & 63))}
        v24 := GenS2{F0: -1}
        _ = v23
        _ = v24
    } else {
        fmt.Println(v1, v1)
    }
    fmt.Println(v22, v1)
    if (("\n\t" >= (("" + ((("\n\t" + "foo bar") + "\n\t") + ("" + (((("hello" + "a") + ("foo bar" + "hello")) + "a") + ("\n\t" + (("foo bar" + "hello") + ("" + "hello"))))))) + "a")) && true) {
        fmt.Println(v22, v1)
        var v25 uint8 = 12
        v26 := GenS3{F0: (!(1e10 < 1.5))}
        _ = v25
        _ = v26
    }
    v27 := GenS4{F0: uint64(v1), F1: 84}
    v2 = append(v2, -1)
    v22 = 42
    fmt.Println(v1)
    if false {
        var v28 uint8 = 119
        _ = v28
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        var v30 int8 = -22
        for fi31 := 0; fi31 < 14; fi31++ {
            v22 += (v22 & v22)
            v30 = 100
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi31
        }
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v3
}
