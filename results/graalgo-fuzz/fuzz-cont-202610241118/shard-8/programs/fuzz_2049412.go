package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 int; F1 int32; F2 string }
func genFn0(p0 int16, p1 uint16, p2 bool) uint {
    for fi5 := 0; fi5 < 10; fi5++ {
        if (uint8(p1) <= 42) {
            var v6 float64 = 0.0
            var v7 uint32 = 534
            _ = v6
            _ = v7
        } else {
            fmt.Println(p2)
        }
        var v8 uint8 = uint8(p1)
        _ = fi5
        _ = v8
    }
    if (-2.25 >= (1.5 - 3.14159)) {
        p2 = (!false)
        p1 = 65535
    }
    _ = p0
    _ = p1
    _ = p2
    return 930
}
func genFn1(p0 bool, p1 uint16, p2 int8) uint32 {
    if true {
        var v9 uint16 = (^(p1))
        _ = v9
    }
    var v10 float64 = (1.5 + (1.5 / (3.14159 - (((-2.25 - 3.14159) - 3.14159) + ((1.5 / 1.5) - ((float64((((float32(float64((3.14159 - ((-2.25 / (0.1 + 1.5)) / 1.0)))) * float32(((0.0 / 0.1) * 0.1))) + 1e10) + 0.1)) * 0.1) + 3.14159))))))
    for fi11 := 0; fi11 < 1; fi11++ {
        if (p2 <= (p2 ^ p2)) {
            v12 := []int{fi11, fi11}
            switch fi11 {
            case -421:
                var v13 int = (fi11 - fi11)
                _ = v13
            default:
                for fi14 := 0; fi14 < 8; fi14++ {
                    if len(v12) > 0 {
                        v12[0] = fi14
                    }
                    var v15 uint32 = 2147483648
                    var v16 uint = 1
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
            var v17 int8 = p2
            _ = v12
            _ = v17
        }
        if p0 {
            v18 := GenS1{F0: 65535, F1: 323}
            for fi19 := 0; fi19 < 20; fi19++ {
                var v20 int16 = 108
                fmt.Println(v10, v20, v10)
                _ = fi19
                _ = v20
            }
            _ = v18
        }
        var v21 uint16 = p1
        _ = fi11
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 414
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint32 = 107
    v3 := make(map[int]bool)
    var v4 uint = 440
    v22 := []int{-359, -857, 32767}
    v23 := GenS1{F0: (^(v2)), F1: 42}
    v24 := GenS2{F0: 275}
    v25 := GenS3{F0: "\n\t", F1: "\n\t"}
    v26 := GenS3{F0: "foo bar", F1: ((("foo bar" + "") + "") + "hello")}
    v27 := GenS1{F0: v2, F1: 2}
    var v28 int8 = -1
    v1 -= v1
    v29 := GenS2{F0: -9223372036854775808}
    switch -9223372036854775808 {
    case -2147483648:
        delete(v3, 42)
    case 1:
        var v30 bool = (123 < 127)
        _ = v30
    default:
        switch "hello" {
        case "foo bar":
            var v31 int16 = 205
            _ = v31
        case "\n\t":
            for fi32 := 0; fi32 < 5; fi32++ {
                v33 := GenS4{F0: (-(fi32)), F1: v1, F2: "hello"}
                var v34 uint = v4
                _ = fi32
                _ = v33
                _ = v34
            }
        case "a":
            var v35 int16 = 42
            _ = v35
        }
    }
    v1 = (^(v1))
    fmt.Println(v4, v1, v1)
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    for fi36 := 0; fi36 < 19; fi36++ {
        v37 := GenS1{F0: v2, F1: 933}
        _ = fi36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
