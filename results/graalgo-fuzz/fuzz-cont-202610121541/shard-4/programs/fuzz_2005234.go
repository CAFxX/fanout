package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint32; F2 bool }
func genFn0(p0 int8) int32 {
    switch ("a" + "\n\t") {
    case "a":
        v4 := []int{-9223372036854775808, 49, 0}
        _ = v4
    case "foo bar":
        var v5 string = ""
        _ = v5
    case "\n\t":
        v6 := []int{0}
        _ = v6
    }
    p0 = (^(p0))
    _ = p0
    return -1
}

func main() {
    v1 := []uint32{}
    var v2 int = -9223372036854775808
    v3 := []string{"\n\t", ((("hello" + (("hello" + "foo bar") + ("\n\t" + "foo bar"))) + "") + "hello"), "a", ""}
    switch ("foo bar" + "a") {
    case "hello":
        v7 := GenS1{F0: (!(!(!false))), F1: "a"}
        _ = v7
    case "":
        v2 = (^(v2))
    }
    v2 = -1
    switch (("foo bar" + "a") + (("hello" + (("" + "a") + "hello")) + "foo bar")) {
    case "foo bar":
        var v8 int32 = -351
        _ = v8
    case "\n\t":
        v9 := GenS2{F0: 32767, F1: 958}
        _ = v9
    }
    if false {
        for fi10 := 0; fi10 < 4; fi10++ {
            switch (-(fi10)) {
            case -421:
                if (42 == 65535) {
                    var v11 int8 = 24
                    v2 = v2
                    v1 = append(v1, 65535)
                    _ = v11
                }
            case 9223372036854775807:
                v2 += 42
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch fi10 {
            case 0:
                var v12 bool = ((!((-9223372036854775808 >= 278) && (v2 != (fi10 * 32767)))) || true)
                _ = v12
            case -127:
                fmt.Println(v2, fi10)
            }
            _ = fi10
        }
    }
    var v13 int8 = -72
    v14 := GenS2{F0: -1, F1: 378}
    v15 := GenS3{F0: 128, F1: 42, F2: false}
    var v16 uint32 = 267
    for fi17 := 0; fi17 < 3; fi17++ {
        switch "hello" {
        case "foo bar":
            var v18 int32 = 908
            _ = v18
        }
        v13 = 38
        _ = fi17
    }
    if false {
        if (42 > 1000) {
            v16 += 2147483647
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v1 = append(v1, (v16 % (((^(v16))) | 1)))
        }
        if len(v1) > 0 {
            v1[0] = 83
        }
        var v19 int64 = 521
        _ = v19
    } else {
        var v20 uint = 4294967295
        _ = v20
    }
    var v21 bool = (-2147483648 < (v2 + 232))
    for fi22 := 0; fi22 < 2; fi22++ {
        v23 := GenS1{F0: ((v21 || v21) && (42 > 579)), F1: "hello"}
        _ = fi22
        _ = v23
    }
    v13 += v13
    var v24 uint = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v24
    _ = v3
}
