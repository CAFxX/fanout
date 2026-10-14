package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int8 }
func genFn0() uint8 {
    var v4 float32 = 0.1
    v4 += float32((0.0 * 0.1))
    var v5 float64 = (((float64(3.14159) - 0.0) + 1.5) / 3.14159)
    v6 := []int{51, 4294967295, -494, 0}
    switch 32767 {
    case 32768:
        var v7 uint64 = 0
        _ = v7
    case -836:
        v8 := GenS1{F0: -9223372036854775808, F1: true, F2: -5}
        _ = v8
    default:
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 string = "hello"
            for fi11 := 0; fi11 < 16; fi11++ {
                var v12 uint64 = 2
                _ = fi11
                _ = v12
            }
            var v13 int64 = 961
            _ = fi9
            _ = v10
            _ = v13
        }
    }
    _ = v4
    _ = v5
    _ = v6
    return 128
}
func genFn1() int64 {
    if (!false) {
        var v14 int = 9223372036854775807
        _ = v14
    }
    var v15 uint64 = 42
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 uint8 = 128
        var v18 int = (fi16 ^ -9223372036854775808)
        switch ("hello" + "a") {
        case "\n\t":
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 float32 = 0.0
                _ = fi19
                _ = v20
            }
        case "":
            v21 := []int{128}
            _ = v21
        default:
            var v22 int64 = 1
            _ = v22
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    _ = v15
    return -9223372036854775808
}

func main() {
    v1 := []string{}
    var v2 uint16 = 2
    v3 := make(map[string]bool)
    v2 = 2
    if true {
        if false {
            v2 = (^(v2))
            if (165 <= 42) {
                v2 = v2
                fmt.Println(v2)
            }
        } else {
            for fi23 := 0; fi23 < 15; fi23++ {
                v2 = (^(v2))
                var v24 bool = ((!true) || false)
                v24 = true
                _ = fi23
                _ = v24
            }
            v25 := GenS1{F0: 42, F1: false, F2: 127}
            var v26 uint32 = 965
            _ = v25
            _ = v26
        }
        var v27 float32 = 3.14159
        _ = v27
    }
    fmt.Println(v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v2)
    _ = len(v3)
    if ((false && true) && false) {
        fmt.Println(v2)
        v2 = 2
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v28 uint16 = v2
    var v29 uint32 = 395
    v28 = v2
    switch 32767 {
    case 4294967295:
        fmt.Println(v2)
    }
    var v30 float64 = float64(float32(3.14159))
    for fi31 := 0; fi31 < 11; fi31++ {
        v3["a"] = true
        _ = fi31
    }
    fmt.Println(v30, v2)
    switch ("" + ("hello" + "a")) {
    case "":
        fmt.Println(v29, v2)
    case "\n\t":
        for fi32 := 0; fi32 < 18; fi32++ {
            var v33 uint16 = (v2 + 1)
            _ = fi32
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 12; fi34++ {
        if (((((("hello" + ("a" + "foo bar")) + ("a" + "")) + "") + ("hello" + "a")) != "") && false) {
            v35 := GenS1{F0: 9223372036854775807, F1: (18446744073709551615 == 4294967295), F2: 100}
            _ = v35
        } else {
            var v36 float64 = (v30 / float64(1.5))
            v30 -= (1e10 * 1.5)
            _ = v36
        }
        v30 -= v30
        switch "hello" {
        case "hello":
            v29 = (v29 >> ((18446744073709551615) & 63))
        case "a":
            var v37 uint16 = v2
            _ = v37
        case "foo bar":
            if (!false) {
                v29 -= v29
                v30 = v30
                v30 = (1e10 - 3.14159)
            } else {
                fmt.Println(v29, v28)
                var v38 uint32 = (v29 * 1)
                v3[""] = true
                _ = v38
            }
        }
        _ = fi34
    }
    var v39 uint = 4294967295
    v2 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v39
}
