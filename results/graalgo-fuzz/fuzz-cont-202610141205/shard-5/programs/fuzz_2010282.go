package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int; F1 bool; F2 uint32 }
type GenS3 struct { F0 int; F1 uint }
func genFn0(p0 int8, p1 int8) uint {
    v4 := []int{4294967295, 290, 4294967295, -149}
    v5 := GenS1{F0: 0, F1: -1}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v6 := GenS2{F0: -1, F1: false, F2: 4294967295}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 478
}
func genFn1() int16 {
    switch "\n\t" {
    case "hello":
        var v7 int16 = 32767
        _ = v7
    case "\n\t":
        var v8 float32 = ((0.1 / float32(((0.1 - -2.25) * (-2.25 + (0.1 * 1e10))))) - (3.14159 * (((0.1 / 1e10) + 1.5) / 0.1)))
        _ = v8
    case "foo bar":
        var v9 uint = genFn0(127, 127)
        _ = v9
    }
    v10 := []int{127, 726, 2147483647}
    _ = v10
    return -46
}

func main() {
    var v1 uint16 = 352
    var v2 int32 = -1
    var v3 uint64 = 1
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 uint = 2
        _ = fi11
        _ = v12
    }
    switch -229 {
    case -477:
        var v13 float64 = (1.5 / float64((1.5 + 3.14159)))
        _ = v13
    case 4294967295:
        v3 = 2
    case -9223372036854775808:
        switch -2147483648 {
        case -113:
            v14 := GenS1{F0: 0, F1: 970}
            _ = v14
        case 1:
            v15 := GenS3{F0: -9223372036854775808, F1: 912}
            _ = v15
        case -2147483648:
            var v16 string = "foo bar"
            _ = v16
        }
    default:
        var v17 uint64 = (v3 / ((0) | 1))
        _ = v17
    }
    v18 := GenS1{F0: v2, F1: -769}
    if (-1 < -1) {
        var v19 int16 = -32768
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 int16 = v19
            v22 := GenS3{F0: (fi20 + fi20), F1: genFn0(1, -1)}
            _ = fi20
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 6; fi23++ {
            v24 := []int{-2147483648, (fi23 - fi23), fi23, (fi23 ^ (fi23 | int(v3)))}
            v2 = v2
            _ = fi23
            _ = v24
        }
        _ = v19
    } else {
        v25 := GenS3{F0: 127, F1: 42}
        v26 := []int{-704}
        for fi27 := 0; fi27 < 19; fi27++ {
            var v28 string = "hello"
            _ = fi27
            _ = v28
        }
        _ = v25
        _ = v26
    }
    for fi29 := 0; fi29 < 4; fi29++ {
        switch ((("foo bar" + (("hello" + "hello") + "hello")) + "hello") + (("hello" + "a") + "")) {
        case "":
            v30 := GenS3{F0: fi29, F1: 4294967295}
            _ = v30
        case "foo bar":
            var v31 uint64 = (v3 + (v3 - 170))
            _ = v31
        default:
            var v32 int64 = -1
            _ = v32
        }
        _ = fi29
    }
    fmt.Println(v1)
    var v33 int64 = -9223372036854775808
    var v34 uint8 = 109
    var v35 float32 = float32(1.5)
    v36 := GenS3{F0: -9223372036854775808, F1: 255}
    for fi37 := 0; fi37 < 6; fi37++ {
        v38 := GenS2{F0: (^(fi37)), F1: (!((float32(0.1) / -2.25) == v35)), F2: 0}
        v39 := []int{(^(fi37)), 794, int(v34)}
        _ = fi37
        _ = v38
        _ = v39
    }
    for fi40 := 0; fi40 < 19; fi40++ {
        v2 = (v2 << ((822) & 63))
        var v41 int64 = (^(v33))
        v35 = v35
        _ = fi40
        _ = v41
    }
    v35 = (0.1 - float32(3.14159))
    v42 := []int{-717, -1}
    v43 := GenS1{F0: v2, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v42:", len(v42), v42)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v42
    _ = v43
}
