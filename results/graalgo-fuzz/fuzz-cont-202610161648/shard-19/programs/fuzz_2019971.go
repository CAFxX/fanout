package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]string)
    var v3 float32 = 0.1
    switch 1 {
    case 252:
        var v4 bool = false
        _ = v4
    }
    var v5 int64 = 9223372036854775807
    switch (("a" + ("\n\t" + "\n\t")) + "hello") {
    case "":
        v5 = v5
    case "\n\t":
        if false {
            var v6 int = -9223372036854775808
            var v7 int32 = -1
            _ = v6
            _ = v7
        } else {
            v8 := GenS1{F0: v5, F1: 512}
            v9 := v2[1]
            fmt.Println(v9, v5)
            _ = v8
            _ = v9
        }
    default:
        var v10 float64 = -2.25
        _ = v10
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 float32 = v3
        v13 := GenS1{F0: (v5 & (v5 | v5)), F1: 32767}
        v5 -= v5
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: (v5 >> ((1) & 63)), F1: 0}
    var v15 uint64 = 4294967295
    if true {
        var v16 int16 = -1
        v15 -= v15
        var v17 uint16 = 0
        _ = v16
        _ = v17
    }
    v18 := GenS2{F0: "a"}
    if (true && (true || false)) {
        v19 := []int{30, 1, -9223372036854775808, -30}
        delete(v2, 0)
        _ = v19
    }
    if false {
        switch (((("foo bar" + ((("hello" + "a") + "\n\t") + "a")) + ("hello" + "a")) + "foo bar") + "foo bar") {
        case "hello":
            switch 2147483647 {
            case 32767:
                var v20 int16 = 1
                _ = v20
            case 2147483647:
                for fi21 := 0; fi21 < 1; fi21++ {
                    for fi22 := 0; fi22 < 16; fi22++ {
                        var v23 int16 = 827
                        var v24 string = ""
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    _ = fi21
                }
            case -760:
                switch "" {
                case "hello":
                    v25 := []int{4294967295}
                    _ = v25
                case "foo bar":
                    var v26 uint32 = 56
                    _ = v26
                case "a":
                    if true {
                        v15 -= (^(v15))
                    } else {
                        fmt.Println(v5, v5, v3)
                        v27 := []int{127, -9223372036854775808, 841, 2147483647}
                        _ = v27
                    }
                }
            }
        case "\n\t":
            v15 = (v15 - 656)
        }
        fmt.Println(v3)
        var v28 int = -1
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v5
}
