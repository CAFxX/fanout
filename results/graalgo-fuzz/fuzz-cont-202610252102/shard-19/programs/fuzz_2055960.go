package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint32; F1 int32; F2 int64 }
type GenS3 struct { F0 bool; F1 uint }
func genFn0(p0 uint64, p1 uint8, p2 int8) int32 {
    p0 = p0
    if true {
        p1 = p1
        v4 := GenS1{F0: true, F1: p2, F2: 310}
        _ = v4
    } else {
        v5 := []int{186}
        var v6 uint = 0
        switch 513 {
        case 0:
            fmt.Println(p0, p0)
        case 1:
            v7 := GenS1{F0: true, F1: (^(p2)), F2: uint32(p1)}
            _ = v7
        case 32768:
            p1 += p1
        default:
            if false {
                var v8 string = "foo bar"
                p0 = (p0 << (((v6 - 4294967295)) & 63))
                _ = v8
            } else {
                _ = len(v5)
                v9 := GenS1{F0: ("hello" <= ("\n\t" + "hello")), F1: (p2 << (((v6 - 295)) & 63)), F2: uint32(p0)}
                p2 = (p2 & (p2 ^ (p2 * 127)))
                _ = v9
            }
        }
        _ = v5
        _ = v6
    }
    switch ("foo bar" + ("foo bar" + "foo bar")) {
    case "a":
        fmt.Println(p2)
    case "hello":
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 int32 = 2147483647
    var v2 bool = false
    v3 := make(map[int]bool)
    v1 = (v1 - genFn0(1, 128, 111))
    v1 = (v1 / ((v1) | 1))
    v2 = v2
    switch (("\n\t" + ("" + ("\n\t" + ((("hello" + "") + "a") + ("foo bar" + "foo bar"))))) + "") {
    case "foo bar":
        v10 := []int{-565, -9223372036854775808, 128, -1}
        _ = v10
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := GenS2{F0: 4294967295, F1: genFn0(18446744073709551615, 1, -1), F2: 42}
        v13 := GenS1{F0: v2, F1: 127, F2: 139}
        _ = fi11
        _ = v12
        _ = v13
    }
    v2 = true
    v14 := []int{32767, 4294967295, 259, -694}
    var v15 uint16 = 987
    fmt.Println(v1)
    v1 += (-(v1))
    if true {
        v16 := GenS2{F0: 2147483647, F1: v1, F2: -9223372036854775808}
        for fi17 := 0; fi17 < 16; fi17++ {
            var v18 uint64 = 213
            _ = len(v14)
            _ = fi17
            _ = v18
        }
        var v19 float32 = (-2.25 / (0.0 + 1e10))
        _ = v16
        _ = v19
    } else {
        var v20 int32 = v1
        _ = len(v14)
        var v21 int8 = -128
        _ = v20
        _ = v21
    }
    v14 = append(v14, 1)
    var v22 uint8 = 1
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        if (false || (255 <= (v22 << ((18446744073709551615) & 63)))) {
            v1 += v1
            for fi23 := 0; fi23 < 7; fi23++ {
                v1 = (v1 + v1)
                v3[fi23] = true
                var v24 string = "\n\t"
                _ = fi23
                _ = v24
            }
        } else {
            for fi25 := 0; fi25 < 2; fi25++ {
                var v26 int = fi25
                _ = fi25
                _ = v26
            }
            fmt.Println(v1, v2)
        }
    case "":
        var v27 uint32 = uint32(v22)
        _ = v27
    case "hello":
        fmt.Println(v22, v2)
    default:
        for fi28 := 0; fi28 < 19; fi28++ {
            v29 := GenS1{F0: true, F1: 0, F2: 106}
            switch "\n\t" {
            case "foo bar":
                v30 := v3[(fi28 / ((127) | 1))]
                _ = v30
            case "hello":
                var v31 int32 = v1
                _ = v31
            default:
                v32 := GenS1{F0: (!v2), F1: -128, F2: 845}
                _ = v32
            }
            _ = fi28
            _ = v29
        }
    }
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    fmt.Println(v1, v15)
    v33 := GenS3{F0: (((0.1 + float32((float64(0.1) / 3.14159))) + 0.0) == 0.1), F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v33
}
