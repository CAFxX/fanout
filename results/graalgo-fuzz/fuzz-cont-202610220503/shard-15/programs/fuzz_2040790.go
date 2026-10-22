package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[string]int)
    var v2 int8 = -1
    var v3 int64 = 1
    var v4 int32 = 239
    var v5 uint = 0
    switch ("foo bar" + "foo bar") {
    case "":
        var v6 int8 = v2
        _ = v6
    case "\n\t":
        if (4294967295 >= v5) {
            switch ((("\n\t" + ("a" + "hello")) + "\n\t") + "a") {
            case "a":
                for fi7 := 0; fi7 < 10; fi7++ {
                    v4 -= (v4 & (v4 % ((-9) | 1)))
                    var v8 uint64 = 42
                    var v9 uint8 = 0
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            case "":
                var v10 int = -9223372036854775808
                _ = v10
            case "hello":
                var v11 uint64 = 0
                _ = v11
            }
            v12 := v1["a"]
            fmt.Println(v5)
            _ = v12
        }
    default:
        var v13 float64 = 0.0
        _ = v13
    }
    fmt.Println(v2)
    switch (("" + "\n\t") + "") {
    case "hello":
        v14 := GenS1{F0: 0, F1: ((true || true) && true), F2: v5}
        _ = v14
    case "foo bar":
        v15 := []int{42, 32767, -255, -9223372036854775808}
        _ = v15
    default:
        if false {
            _ = len(v1)
            var v16 int64 = -1
            v17 := GenS1{F0: 220, F1: (((0.1 / 1.0) * (-2.25 + 1e10)) <= float32((0.1 / float64(1.5)))), F2: (v5 ^ 1)}
            _ = v16
            _ = v17
        } else {
            fmt.Println(v2)
            v2 = 42
            v5 -= (v5 >> (((v5 | (v5 ^ v5))) & 63))
        }
    }
    fmt.Println(v2)
    v18 := GenS1{F0: 65535, F1: (108 > 1), F2: (v5 / ((v5) | 1))}
    if false {
        var v19 uint32 = 2147483647
        v19 = v19
        var v20 float32 = (1.5 / 1.5)
        _ = v19
        _ = v20
    }
    if true {
        switch ("foo bar" + ("a" + (("hello" + "") + (("hello" + ("" + (("\n\t" + "") + "foo bar"))) + (("a" + "\n\t") + (("hello" + ("\n\t" + "a")) + "")))))) {
        case "hello":
            var v21 uint = (v5 - 373)
            _ = v21
        case "\n\t":
            v22 := GenS1{F0: 2, F1: (-2.25 == 1e10), F2: v5}
            _ = v22
        case "":
            v23 := []int{1}
            _ = v23
        default:
            var v24 uint16 = 0
            _ = v24
        }
        var v25 uint8 = 1
        _ = v25
    }
    v2 = -128
    if (!(!(!true))) {
        var v26 bool = (true || true)
        _ = v26
    }
    v2 = 1
    switch ("foo bar" + ("" + ("" + ((("foo bar" + "a") + "a") + "\n\t")))) {
    case "hello":
        v27 := v1["a"]
        _ = v27
    case "":
        v3 = v3
    default:
        v5 -= v5
    }
    v28 := GenS1{F0: 1, F1: (true && true), F2: v5}
    v5 -= (v5 << (((v5 << ((242) & 63))) & 63))
    v4 = v4
    v29 := []int{0, -1000}
    v2 = (v2 + (v2 * v2))
    if false {
        v30 := GenS2{F0: v2}
        switch -382 {
        case -2147483648:
            for fi31 := 0; fi31 < 7; fi31++ {
                for fi32 := 0; fi32 < 8; fi32++ {
                    var v33 uint = 65535
                    var v34 bool = true
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                if len(v29) >= 2 {
                    v29 = v29[1:len(v29)]
                }
                v4 = (v4 / (((v4 * v4)) | 1))
                _ = fi31
            }
        }
        if len(v29) > 0 {
            v29[0] = -1
        }
        _ = v30
    } else {
        var v35 uint = 0
        for fi36 := 0; fi36 < 9; fi36++ {
            delete(v1, "hello")
            v5 = (v35 - v5)
            v5 -= 160
            _ = fi36
        }
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
