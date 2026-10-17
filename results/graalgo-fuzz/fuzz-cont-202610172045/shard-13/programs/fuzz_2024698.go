package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int16; F2 int }
func genFn0() uint32 {
    var v4 int32 = 1
    switch 2147483647 {
    case 1:
        var v5 bool = true
        _ = v5
    case 0:
        var v6 uint16 = 171
        _ = v6
    default:
        var v7 float32 = -2.25
        _ = v7
    }
    if ((-128 < 127) || (0 == 9223372036854775807)) {
        v8 := GenS1{F0: 2147483647}
        var v9 uint8 = 42
        _ = v8
        _ = v9
    } else {
        v10 := []int{2147483647}
        _ = v10
    }
    var v11 bool = true
    if false {
        var v12 string = (("a" + (("foo bar" + "") + ((("hello" + (("a" + (("a" + "foo bar") + "a")) + ("\n\t" + "a"))) + "\n\t") + "\n\t"))) + ((("foo bar" + "a") + ("\n\t" + "")) + ("hello" + "foo bar")))
        _ = v12
    } else {
        fmt.Println(v11, v4)
    }
    _ = v11
    _ = v4
    return 961
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 182
    var v3 uint8 = 247
    v3 = v3
    v2 = (v2 ^ 199)
    if (!(!(!(!false)))) {
        var v13 float32 = (0.1 + ((0.0 / 3.14159) * 1.5))
        v13 = (-2.25 / (v13 / (v13 + v13)))
        var v14 bool = (1 >= int8(v2))
        _ = v13
        _ = v14
    }
    v2 -= v2
    var v15 bool = false
    v3 = v3
    switch 0 {
    case 128:
        switch "hello" {
        case "hello":
            for fi16 := 0; fi16 < 1; fi16++ {
                var v17 uint32 = genFn0()
                var v18 int32 = 2147483647
                var v19 int = fi16
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
        case "\n\t":
            var v20 uint64 = 4294967295
            _ = v20
        case "foo bar":
            v21 := []int{}
            _ = v21
        }
    case -400:
        var v22 uint16 = 65535
        _ = v22
    default:
        fmt.Println(v3)
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        if (1 <= 1) {
            switch (fi23 >> ((175) & 63)) {
            case -250:
                v24 := GenS1{F0: 326}
                _ = v24
            default:
                v15 = (!v15)
            }
            switch fi23 {
            case 137:
                v25 := []int{}
                _ = v25
            default:
                var v26 int64 = -9223372036854775808
                _ = v26
            }
        }
        _ = fi23
    }
    v27 := GenS1{F0: genFn0()}
    if false {
        switch 2147483647 {
        case 89:
            fmt.Println(v2, v15, v15)
        case 407:
            v1[48] = 348
        case 1:
            for fi28 := 0; fi28 < 3; fi28++ {
                if v15 {
                    v1[127] = 65535
                    v15 = ("hello" == "foo bar")
                } else {
                    v15 = (4294967295 < 1)
                    v1[fi28] = uint32(v2)
                    v29 := GenS1{F0: 2}
                    _ = v29
                }
                _ = fi28
            }
        }
        var v30 int16 = -1
        _ = v30
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        var v32 uint64 = 4294967295
        fmt.Println(v15, v3, v2)
        _ = fi31
        _ = v32
    }
    var v33 float32 = 0.1
    v3 += 1
    fmt.Println(v2, v33, v2)
    if ((!(1e10 == -2.25)) && (v2 < (^(v2)))) {
        var v34 int = -2147483648
        _ = v34
    }
    if ((9223372036854775807 > 32768) || true) {
        var v35 string = ""
        var v36 string = v35
        _ = v35
        _ = v36
    }
    if v15 {
        v37 := GenS2{F0: -9223372036854775808, F1: -1, F2: 774}
        _ = v37
    }
    var v38 uint32 = uint32(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v38
}
