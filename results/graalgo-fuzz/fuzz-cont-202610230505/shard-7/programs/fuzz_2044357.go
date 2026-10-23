package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 string; F2 int8 }
func genFn0() int32 {
    var v7 int = 42
    if (0.0 >= 0.1) {
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := []int{v7, (-(v7)), -48}
            for fi10 := 0; fi10 < 18; fi10++ {
                v11 := GenS1{F0: 2147483647, F1: (false && false), F2: "\n\t"}
                _ = fi10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
        var v12 int16 = int16(v7)
        switch "hello" {
        case "":
            v12 = 0
        case "hello":
            v13 := GenS1{F0: 1, F1: true, F2: "foo bar"}
            _ = v13
        }
        _ = v12
    } else {
        v7 = v7
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 bool = true
            _ = fi14
            _ = v15
        }
        v16 := []int{9223372036854775807, v7, v7}
        _ = v16
    }
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = true
    v3 := make(map[string]string)
    v4 := make(map[string]int)
    var v5 int64 = -1
    var v6 string = (("a" + "hello") + "hello")
    switch -502 {
    case 4294967295:
        for fi17 := 0; fi17 < 3; fi17++ {
            v18 := []int{(-(fi17)), (fi17 ^ (^(fi17))), 42, (^(fi17))}
            _ = fi17
            _ = v18
        }
    case 32767:
        var v19 uint32 = 2
        _ = v19
    }
    fmt.Println(v2)
    v2 = false
    fmt.Println(v6)
    var v20 float64 = ((1e10 / ((float64(float32(1e10)) - 1.5) / 3.14159)) - (1e10 * (float64(1.5) - -2.25)))
    for fi21 := 0; fi21 < 12; fi21++ {
        if (v6 != ("" + "foo bar")) {
            var v22 int32 = 533
            var v23 uint64 = 1
            v24 := GenS2{F0: v23}
            _ = v22
            _ = v23
            _ = v24
        }
        v20 = v20
        _ = fi21
    }
    var v25 uint64 = 795
    var v26 int8 = -40
    var v27 int64 = (v5 << ((907) & 63))
    fmt.Println(v26, v26)
    var v28 uint16 = 65535
    switch -1 {
    case 489:
        var v29 string = ((("foo bar" + (v6 + "foo bar")) + (v6 + v6)) + ((("a" + v6) + (v6 + v6)) + "foo bar"))
        _ = v29
    case 1:
        v20 = v20
    default:
        var v30 uint64 = v25
        _ = v30
    }
    v31 := GenS3{F0: 100, F1: v6, F2: (v26 << ((18446744073709551615) & 63))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
