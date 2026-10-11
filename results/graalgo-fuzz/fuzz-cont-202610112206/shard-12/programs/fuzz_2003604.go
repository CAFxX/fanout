package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int64; F1 uint; F2 uint32 }
func genFn0() int64 {
    if (true && true) {
        var v4 bool = true
        _ = v4
    } else {
        var v5 string = "a"
        _ = v5
    }
    if false {
        v6 := GenS1{F0: false}
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: false}
            fmt.Println(fi7)
            _ = fi7
            _ = v8
        }
        _ = v6
    } else {
        v9 := GenS1{F0: true}
        var v10 float64 = -2.25
        _ = v10
        _ = v9
    }
    var v11 float32 = (0.1 * float32(3.14159))
    var v12 string = ""
    _ = v11
    _ = v12
    return 9223372036854775807
}
func genFn1() int {
    if (-445 <= 128) {
        switch 587 {
        case 1:
            switch 4294967295 {
            case 42:
                var v13 uint32 = 0
                _ = v13
            case -733:
                var v14 int8 = -1
                _ = v14
            default:
                v15 := []int{-967, 4294967295, 2147483647, -989}
                _ = v15
            }
        case 671:
            var v16 uint32 = 2147483647
            _ = v16
        }
        var v17 int64 = 0
        _ = v17
    }
    v18 := []int{128, 128, 1, 9223372036854775807}
    _ = v18
    return 0
}

func main() {
    var v1 string = "foo bar"
    var v2 int32 = 1
    var v3 int64 = 0
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v3, v3)
    case "":
        v19 := []int{}
        _ = v19
    default:
        v20 := []int{1}
        _ = v20
    }
    if false {
        fmt.Println(v2, v3)
        fmt.Println(v3, v3, v1)
    } else {
        switch -1 {
        case -173:
            if (554 == 4294967295) {
                if true {
                    v21 := GenS1{F0: false}
                    _ = v21
                }
                v3 = genFn0()
            }
        case 0:
            v2 -= 42
        }
        switch ((v1 + (("\n\t" + "\n\t") + v1)) + v1) {
        case "\n\t":
            v22 := GenS1{F0: (67 == 78)}
            _ = v22
        }
    }
    var v23 uint8 = 42
    switch (v1 + v1) {
    case "hello":
        for fi24 := 0; fi24 < 15; fi24++ {
            var v25 float64 = 1e10
            v26 := GenS1{F0: (484 >= 1000)}
            v27 := GenS2{F0: v1, F1: v23, F2: 763}
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
    case "a":
        v1 = "\n\t"
    default:
        v3 -= v3
    }
    for fi28 := 0; fi28 < 14; fi28++ {
        var v29 int = (fi28 % ((fi28) | 1))
        _ = fi28
        _ = v29
    }
    fmt.Println(v2, v3)
    for fi30 := 0; fi30 < 7; fi30++ {
        v31 := GenS3{F0: 570, F1: 65535, F2: 65535}
        switch (^(fi30)) {
        case 0:
            var v32 bool = (18446744073709551615 < 665)
            _ = v32
        }
        var v33 uint64 = 42
        _ = fi30
        _ = v31
        _ = v33
    }
    var v34 int16 = 0
    var v35 uint32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v34
    _ = v35
}
