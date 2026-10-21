package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 string }
func genFn0() string {
    switch "hello" {
    case "a":
        if false {
            var v6 int64 = 9223372036854775807
            switch ("foo bar" + "a") {
            case "hello":
                if false {
                    var v7 int16 = 1000
                    _ = v7
                }
            }
            fmt.Println(v6)
            _ = v6
        }
    case "hello":
        var v8 int32 = 0
        _ = v8
    default:
        var v9 int16 = -1
        _ = v9
    }
    var v10 uint = 2
    v11 := GenS1{F0: 42, F1: 42}
    _ = v10
    _ = v11
    return ("hello" + ("\n\t" + "\n\t"))
}

func main() {
    var v1 uint16 = 0
    var v2 int = -722
    v3 := make(map[int]int)
    var v4 uint16 = v1
    var v5 uint32 = 1
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 int = (-(v2))
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 int8 = 100
        v16 := GenS1{F0: fi14, F1: 295}
        var v17 int16 = 0
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    var v18 string = (("" + (("a" + "\n\t") + ("" + ("a" + "foo bar")))) + "a")
    v19 := GenS1{F0: 128, F1: v1}
    if ((!true) && false) {
        v20 := GenS1{F0: -996, F1: 42}
        _ = v20
    } else {
        fmt.Println(v2, v5)
    }
    fmt.Println(v2, v2)
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 int = v2
        v23 := v3[(-(fi21))]
        _ = fi21
        _ = v22
        _ = v23
    }
    v1 += v4
    if ("\n\t" < v18) {
        v5 = (^(v5))
    } else {
        if (v18 <= ("" + v18)) {
            v5 = v5
            v24 := GenS2{F0: v18}
            v25 := GenS2{F0: v18}
            _ = v24
            _ = v25
        } else {
            v18 = "\n\t"
            v1 = v4
        }
        var v26 int16 = -32768
        var v27 uint64 = 139
        _ = v26
        _ = v27
    }
    if (v5 > 0) {
        var v28 uint64 = 462
        v3[v2] = 32767
        _ = v28
    } else {
        for fi29 := 0; fi29 < 8; fi29++ {
            var v30 bool = false
            _ = fi29
            _ = v30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
