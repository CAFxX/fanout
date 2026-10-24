package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 string }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint }
func genFn0() bool {
    var v5 int8 = 42
    fmt.Println(v5)
    if (true || (255 != 65535)) {
        v6 := []int{2147483647, -336}
        _ = v6
    } else {
        v7 := GenS1{F0: 2, F1: 771, F2: ("a" + "foo bar")}
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := []int{fi8, fi8}
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            var v10 float64 = 0.0
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := []int{128, -421, 128, 1}
        _ = v11
        _ = v7
    }
    var v12 int64 = 42
    _ = v12
    _ = v5
    return (!false)
}

func main() {
    var v1 int64 = 0
    v2 := make(map[uint32]bool)
    v3 := make(map[string]bool)
    var v4 uint32 = 99
    var v13 uint8 = 255
    v14 := GenS2{F0: 32767, F1: 1, F2: 18446744073709551615}
    switch int(v4) {
    case 2147483647:
        switch int(v4) {
        case 127:
            var v15 int64 = int64(v4)
            _ = v15
        }
    case 42:
        var v16 bool = (-128 <= 46)
        _ = v16
    default:
        v13 -= v13
    }
    switch -9223372036854775808 {
    case -959:
        if false {
            var v17 uint32 = v4
            _ = v17
        }
    case 42:
        v18 := GenS1{F0: v4, F1: 1, F2: "\n\t"}
        _ = v18
    case 751:
        for fi19 := 0; fi19 < 1; fi19++ {
            var v20 float64 = (float64(0.1) - 3.14159)
            var v21 int32 = -2147483648
            _ = fi19
            _ = v20
            _ = v21
        }
    default:
        v22 := []int{840, int(v13), 1}
        _ = v22
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        var v24 uint16 = 42
        _ = fi23
        _ = v24
    }
    if (!(42 > 248)) {
        switch ("" + ("a" + "foo bar")) {
        case "hello":
            v13 += (v13 ^ v13)
        case "a":
            v1 = v1
        case "\n\t":
            var v25 uint16 = 738
            _ = v25
        default:
            v26 := GenS1{F0: v4, F1: 897, F2: "a"}
            _ = v26
        }
    }
    var v27 bool = true
    var v28 float64 = 0.0
    v28 = 3.14159
    fmt.Println(v4)
    for fi29 := 0; fi29 < 14; fi29++ {
        switch fi29 {
        case -9223372036854775808:
            v13 += (v13 ^ (^(v13)))
        case -857:
            var v30 string = "a"
            _ = v30
        case 128:
            var v31 uint64 = 18446744073709551615
            _ = v31
        }
        _ = fi29
    }
    if v27 {
        v32 := GenS2{F0: 0, F1: uint16(v13), F2: 2}
        var v33 float64 = (3.14159 + v28)
        v34 := []int{4294967295, 32768, 127, 224}
        _ = v32
        _ = v33
        _ = v34
    }
    if (("" + "hello") <= "foo bar") {
        v4 -= 0
        var v35 float32 = (float32(-2.25) + 3.14159)
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
