package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int16 }
func genFn0(p0 bool, p1 bool) int {
    if ((127 <= 0) || (!true)) {
        p0 = ((false && ((904 == 255) || false)) && ((32767 < 32767) && p0))
        fmt.Println(p1, p1)
    } else {
        var v7 int = -1
        var v8 uint16 = 665
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            switch (("\n\t" + "a") + "\n\t") {
            case "foo bar":
                var v11 string = "a"
                _ = v11
            }
            var v12 uint64 = 300
            _ = fi10
            _ = v12
        }
        _ = fi9
    }
    var v13 uint64 = 107
    fmt.Println(v13, v13, v13)
    _ = p0
    _ = p1
    _ = v13
    return 42
}

func main() {
    var v1 int = 2147483647
    var v2 uint16 = 42
    var v3 float64 = 3.14159
    var v4 bool = false
    var v5 bool = true
    var v6 uint16 = (v2 * (v2 / (((^(v2))) | 1)))
    fmt.Println(v6, v3, v3)
    if v4 {
        v14 := GenS1{F0: "\n\t"}
        _ = v14
    }
    var v15 uint16 = 76
    if (-128 > -128) {
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 bool = (0 == 2147483647)
            v18 := []int{(^(v1))}
            v6 -= v15
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    if v5 {
        var v19 int32 = 0
        var v20 int8 = -91
        _ = v19
        _ = v20
    }
    if true {
        var v21 int16 = 32767
        _ = v21
    } else {
        for fi22 := 0; fi22 < 8; fi22++ {
            v23 := GenS2{F0: 1000, F1: -32768}
            v24 := GenS1{F0: "a"}
            v25 := []int{fi22, genFn0((1000 > 42), true), -2147483648, v1}
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
        var v26 uint = 0
        _ = v26
    }
    var v27 uint32 = 2147483647
    v6 += (^(v2))
    if (true && ((9223372036854775807 != -598) && ("\n\t" < ("" + ("foo bar" + "\n\t"))))) {
        var v28 uint8 = 165
        v29 := []int{(-(v1)), v1, v1, 32767}
        for fi30 := 0; fi30 < 16; fi30++ {
            _ = len(v29)
            if (4294967295 <= 42) {
                var v31 int16 = 785
                var v32 int64 = -938
                _ = v31
                _ = v32
            }
            _ = fi30
        }
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
