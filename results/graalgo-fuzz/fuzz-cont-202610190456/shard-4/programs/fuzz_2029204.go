package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
func genFn0(p0 int16) uint32 {
    var v6 int8 = 1
    if (1 <= 437) {
        switch (("hello" + "foo bar") + ("\n\t" + "\n\t")) {
        case "":
            switch 112 {
            case 32768:
                v7 := GenS1{F0: v6, F1: 42}
                _ = v7
            case 56:
                fmt.Println(p0, p0)
            default:
                fmt.Println(p0)
            }
        case "a":
            var v8 int32 = -719
            _ = v8
        case "hello":
            fmt.Println(p0, v6)
        default:
            var v9 uint16 = 65535
            _ = v9
        }
        if (false && (236 != 0)) {
            var v10 uint16 = 0
            _ = v10
        }
    } else {
        v11 := []int{-702, 1, -732, -914}
        _ = v11
    }
    _ = p0
    _ = v6
    return 2
}

func main() {
    var v1 float64 = 1.5
    var v2 string = "a"
    var v3 uint8 = 127
    var v4 bool = (65535 != 0)
    var v5 uint64 = 4294967295
    v12 := GenS1{F0: 0, F1: 2147483648}
    v5 = v5
    v13 := GenS1{F0: 0, F1: 0}
    for fi14 := 0; fi14 < 18; fi14++ {
        fmt.Println(v2)
        v3 = v3
        switch fi14 {
        case -9223372036854775808:
            v15 := []int{fi14, fi14, fi14}
            _ = v15
        case 9223372036854775807:
            for fi16 := 0; fi16 < 9; fi16++ {
                var v17 int8 = 42
                _ = fi16
                _ = v17
            }
        }
        _ = fi14
    }
    if ((-83 >= -128) || (!true)) {
        v5 = v5
        v5 = v5
    }
    v18 := []int{}
    if v4 {
        v19 := GenS1{F0: -1, F1: 2}
        _ = v19
    }
    fmt.Println(v5)
    var v20 float64 = v1
    var v21 bool = v4
    v5 = 39
    if v21 {
        var v22 float32 = 3.14159
        var v23 int16 = -32768
        _ = v22
        _ = v23
    } else {
        fmt.Println(v1, v4)
        fmt.Println(v3)
        v1 = v20
    }
    if false {
        for fi24 := 0; fi24 < 19; fi24++ {
            fmt.Println(v2)
            switch (^(fi24)) {
            case 127:
                var v25 int16 = 32767
                _ = v25
            case 884:
                var v26 uint = 2
                _ = v26
            case 1:
                v5 = (v5 - v5)
            default:
                for fi27 := 0; fi27 < 18; fi27++ {
                    if true {
                        if ((!v21) && (42 == 1)) {
                            var v28 uint64 = (v5 + v5)
                            if len(v18) >= 2 {
                                v18 = v18[1:len(v18)]
                            }
                            var v29 string = v2
                            _ = v28
                            _ = v29
                        }
                    }
                    _ = fi27
                }
            }
            _ = fi24
        }
        if (999 < 2147483647) {
            var v30 int64 = 42
            _ = v30
        }
        _ = len(v18)
    } else {
        var v31 uint32 = 2147483647
        _ = v31
    }
    var v32 int16 = -32768
    fmt.Println(v1, v20, v32)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
