package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint16 }
func genFn0() bool {
    v4 := []int{-696, -9223372036854775808, 4294967295}
    switch ("foo bar" + "") {
    case "a":
        for fi5 := 0; fi5 < 12; fi5++ {
            if true {
                var v6 uint64 = 18446744073709551615
                var v7 string = ("hello" + (("" + ("\n\t" + "\n\t")) + "foo bar"))
                var v8 uint8 = 0
                _ = v6
                _ = v7
                _ = v8
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v9 uint = 0
            _ = fi5
            _ = v9
        }
    default:
        v4 = append(v4, -745)
    }
    var v10 uint64 = 340
    var v11 int8 = 42
    if true {
        v12 := GenS1{F0: "hello"}
        _ = v12
    } else {
        v13 := GenS2{F0: v11, F1: 32767}
        fmt.Println(v11)
        _ = v13
    }
    _ = v10
    _ = v11
    _ = v4
    return false
}

func main() {
    v1 := []string{"foo bar", (((("\n\t" + ("foo bar" + (("hello" + "foo bar") + "foo bar"))) + "foo bar") + "a") + ""), "\n\t", ((("a" + ("\n\t" + "")) + "\n\t") + "\n\t")}
    var v2 int = -9223372036854775808
    v3 := make(map[string]int)
    v14 := GenS1{F0: "a"}
    v15 := GenS2{F0: -128, F1: 2}
    var v16 float32 = 1e10
    switch (^(v2)) {
    case -690:
        switch v2 {
        case 4294967295:
            if (!false) {
                var v17 int = v2
                _ = v17
            }
        case 584:
            var v18 string = "\n\t"
            _ = v18
        case 320:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    default:
        v19 := GenS2{F0: -128, F1: 347}
        _ = v19
    }
    v20 := GenS1{F0: "a"}
    if true {
        v21 := GenS1{F0: ""}
        _ = v21
    } else {
        if (-75 < -1) {
            var v22 uint = 65535
            _ = v22
        } else {
            var v23 uint = 1
            switch -381 {
            case 1:
                v2 = (v2 - v2)
            case 128:
                var v24 float32 = (float32(0.1) - v16)
                _ = v24
            case 2147483647:
                var v25 uint8 = uint8(v23)
                _ = v25
            }
            fmt.Println(v16, v2, v2)
            _ = v23
        }
    }
    v26 := GenS1{F0: "foo bar"}
    switch ("hello" + "hello") {
    case "\n\t":
        var v27 int = v2
        _ = v27
    case "foo bar":
        _ = len(v3)
    case "":
        fmt.Println(v2)
    }
    var v28 uint = 1
    fmt.Println(v28)
    var v29 int64 = -724
    v30 := GenS2{F0: 42, F1: 32767}
    var v31 int64 = v29
    for fi32 := 0; fi32 < 15; fi32++ {
        var v33 float32 = v16
        _ = fi32
        _ = v33
    }
    v34 := GenS1{F0: "foo bar"}
    if (v2 < (v2 & (v2 ^ 1))) {
        var v35 uint8 = 149
        _ = v35
    }
    v29 = -564
    for fi36 := 0; fi36 < 5; fi36++ {
        switch "foo bar" {
        case "\n\t":
            for fi37 := 0; fi37 < 20; fi37++ {
                v16 -= v16
                v31 -= (v29 * -9223372036854775808)
                _ = fi37
            }
        case "hello":
            var v38 int8 = 100
            _ = v38
        default:
            fmt.Println(v16, v31)
        }
        _ = fi36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
}
