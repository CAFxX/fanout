package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int32; F2 int32 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int32; F1 int16; F2 string }
func genFn0(p0 uint8, p1 string) int16 {
    switch -9223372036854775808 {
    case -1:
        var v6 float32 = 0.1
        _ = v6
    }
    var v7 int = -2147483648
    _ = p0
    _ = p1
    _ = v7
    return 1000
}

func main() {
    v1 := []string{(("foo bar" + "\n\t") + "\n\t"), "foo bar", (("a" + (("\n\t" + "hello") + "foo bar")) + "\n\t")}
    var v2 bool = (-2147483648 == -2147483648)
    var v3 uint8 = 1
    v4 := make(map[int]uint32)
    v5 := make(map[string]int)
    switch "\n\t" {
    case "hello":
        v2 = v2
    case "foo bar":
        v8 := GenS1{F0: 112}
        _ = v8
    }
    v9 := GenS2{F0: 42, F1: 1, F2: 2147483647}
    var v10 int64 = -193
    if (v10 != -9223372036854775808) {
        for fi11 := 0; fi11 < 19; fi11++ {
            v10 = (v10 & (v10 / (((v10 + (v10 / ((v10) | 1)))) | 1)))
            _ = fi11
        }
    }
    switch "\n\t" {
    case "foo bar":
        switch -9223372036854775808 {
        case 282:
            var v12 int16 = 0
            _ = v12
        case 2147483647:
            if false {
                var v13 int8 = 0
                _ = v13
            }
        }
    default:
        if (((v10 % ((v10) | 1)) == (v10 + 170)) && (100 > 1)) {
            if (("hello" + "a") < ("foo bar" + "")) {
                var v14 uint64 = 42
                _ = v14
            } else {
                var v15 uint8 = (v3 & v3)
                _ = v15
            }
            v16 := GenS1{F0: 1}
            v2 = true
            _ = v16
        }
    }
    if true {
        v2 = v2
        switch -343 {
        case 0:
            if false {
                var v17 uint = 2
                _ = v17
            }
        case 9223372036854775807:
            v2 = (!false)
        case 127:
            var v18 int32 = -560
            _ = v18
        }
    } else {
        switch 32768 {
        case 2147483647:
            v19 := GenS2{F0: 0, F1: 840, F2: 1}
            _ = v19
        case -551:
            switch (("hello" + ((("a" + "a") + "hello") + "hello")) + "hello") {
            case "a":
                v2 = (!(1 <= 2147483648))
            default:
                fmt.Println(v10)
            }
        default:
            var v20 int = 2147483647
            _ = v20
        }
        v3 = (v3 ^ 128)
        for fi21 := 0; fi21 < 18; fi21++ {
            var v22 uint8 = (^(v3))
            _ = fi21
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        v3 = (^(v3))
        _ = fi23
    }
    var v24 int64 = (v10 ^ -9223372036854775808)
    fmt.Println(v24, v3, v3)
    var v25 int16 = -177
    _ = len(v5)
    v26 := GenS3{F0: v3, F1: 42}
    for fi27 := 0; fi27 < 4; fi27++ {
        for fi28 := 0; fi28 < 19; fi28++ {
            var v29 float64 = float64(0.1)
            v1 = append(v1, ("" + ("foo bar" + "a")))
            _ = fi28
            _ = v29
        }
        v30 := GenS4{F0: 42, F1: 1000, F2: ("\n\t" + "hello")}
        _ = fi27
        _ = v30
    }
    v3 = (^(v3))
    for fi31 := 0; fi31 < 16; fi31++ {
        var v32 uint = 876
        _ = fi31
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
