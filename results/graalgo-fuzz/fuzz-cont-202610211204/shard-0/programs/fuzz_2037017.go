package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }

func main() {
    var v1 uint8 = 127
    var v2 uint16 = 32767
    var v3 uint8 = v1
    var v4 uint = 255
    var v5 uint64 = 24
    var v6 float32 = (((1e10 / 1.0) * 1e10) + float32(3.14159))
    fmt.Println(v1)
    if false {
        v7 := []int{-2147483648, 42, 32768, 1}
        _ = v7
    }
    var v8 int32 = -408
    if false {
        v8 = v8
    }
    if false {
        var v9 uint32 = 1
        _ = v9
    } else {
        v2 = 2
        v10 := GenS1{F0: true, F1: 42}
        switch -196 {
        case 32767:
            var v11 uint = (v4 / ((18446744073709551615) | 1))
            _ = v11
        default:
            var v12 int = 32767
            _ = v12
        }
        _ = v10
    }
    var v13 uint = (v4 % ((v4) | 1))
    fmt.Println(v5, v13, v1)
    if (float32(0.1) >= 0.1) {
        var v14 float32 = ((v6 + v6) + -2.25)
        var v15 int16 = 32767
        _ = v14
        _ = v15
    } else {
        v6 -= 1.5
    }
    if (0 <= -1) {
        switch -2147483648 {
        case -9223372036854775808:
            var v16 float32 = v6
            _ = v16
        }
        v17 := GenS1{F0: (0 >= v13), F1: (v2 | (v2 % ((32767) | 1)))}
        _ = v17
    } else {
        var v18 uint32 = 2147483648
        if true {
            v19 := []int{42}
            _ = v19
        } else {
            fmt.Println(v6, v2, v18)
            v5 = uint64(v3)
        }
        _ = v18
    }
    v20 := GenS1{F0: true, F1: v2}
    switch "hello" {
    case "":
        v21 := []int{1, 32768}
        _ = v21
    case "\n\t":
        switch 32767 {
        case 32767:
            for fi22 := 0; fi22 < 19; fi22++ {
                for fi23 := 0; fi23 < 9; fi23++ {
                    v24 := []int{fi22, (fi23 % (((^(fi22))) | 1)), fi22}
                    var v25 string = ("" + "foo bar")
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                _ = fi22
            }
        }
    }
    switch -229 {
    case 4294967295:
        if (1 == 1) {
            for fi26 := 0; fi26 < 8; fi26++ {
                v27 := []int{int(v8), -382, fi26}
                _ = fi26
                _ = v27
            }
        } else {
            var v28 uint64 = 42
            v29 := GenS1{F0: (float64(v6) < 0.0), F1: (v2 % (((^(v2))) | 1))}
            v30 := []int{-954, 32768, 32767, 4294967295}
            _ = v28
            _ = v29
            _ = v30
        }
    case -275:
        for fi31 := 0; fi31 < 8; fi31++ {
            var v32 int32 = (v8 & (v8 >> ((v13) & 63)))
            _ = fi31
            _ = v32
        }
    case 1:
        var v33 int8 = -1
        _ = v33
    }
    v5 -= 18446744073709551615
    var v34 int = 826
    v34 = -656
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
