package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() uint32 {
    var v4 bool = (4294967295 <= 18446744073709551615)
    var v5 uint8 = 128
    fmt.Println(v5, v5)
    v5 = 114
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := []int{0, fi6, (fi6 / ((fi6) | 1))}
        if v4 {
            var v8 uint32 = 1
            for fi9 := 0; fi9 < 20; fi9++ {
                v10 := GenS1{F0: 42}
                if len(v7) > 0 {
                    v7[0] = fi6
                }
                v8 += (v8 / (((^(v8))) | 1))
                _ = fi9
                _ = v10
            }
            var v11 int = fi6
            _ = v11
            _ = v8
        }
        fmt.Println(v5, v5)
        _ = fi6
        _ = v7
    }
    _ = v4
    _ = v5
    return 2
}
func genFn1() uint8 {
    var v12 int64 = 0
    v12 = v12
    _ = v12
    return 8
}

func main() {
    var v1 int = -694
    v2 := make(map[string]string)
    var v3 string = "hello"
    v13 := []int{4294967295, v1, v1, (v1 & v1)}
    fmt.Println(v1, v1)
    switch "\n\t" {
    case "":
        for fi14 := 0; fi14 < 6; fi14++ {
            if len(v13) > 0 {
                v13[0] = fi14
            }
            var v15 uint8 = 84
            _ = fi14
            _ = v15
        }
    case "foo bar":
        if (244 <= -167) {
            for fi16 := 0; fi16 < 14; fi16++ {
                var v17 int32 = 2147483647
                fmt.Println(v1, v1)
                _ = fi16
                _ = v17
            }
        }
    }
    var v18 int64 = -450
    for fi19 := 0; fi19 < 6; fi19++ {
        _ = len(v2)
        _ = fi19
    }
    v18 -= (v18 * (v18 << ((933) & 63)))
    v3 = v3
    v3 = v3
    var v20 int64 = v18
    fmt.Println(v18)
    if (true || false) {
        var v21 int16 = 233
        var v22 bool = true
        v18 -= v20
        _ = v21
        _ = v22
    } else {
        v1 = v1
    }
    var v23 string = v3
    if false {
        v20 -= (v18 | (^(v20)))
        switch "foo bar" {
        case "hello":
            var v24 string = v23
            _ = v24
        case "\n\t":
            switch (("hello" + v3) + "a") {
            case "":
                delete(v2, ("\n\t" + ("foo bar" + ("a" + ""))))
            case "hello":
                var v25 float64 = 0.0
                _ = v25
            }
        }
    }
    var v26 int8 = 100
    v27 := v2["a"]
    fmt.Println(v26)
    var v28 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v20
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
