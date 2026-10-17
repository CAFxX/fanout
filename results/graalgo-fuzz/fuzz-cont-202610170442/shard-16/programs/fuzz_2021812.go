package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint }
func genFn0(p0 int64) int32 {
    for fi7 := 0; fi7 < 6; fi7++ {
        p0 = -9223372036854775808
        p0 = 0
        v8 := GenS1{F0: 1, F1: 886}
        _ = fi7
        _ = v8
    }
    p0 = p0
    _ = p0
    return -2147483648
}
func genFn1(p0 string, p1 int32, p2 string) int16 {
    v9 := GenS1{F0: 65535, F1: 541}
    var v10 int16 = 0
    switch -160 {
    case 9223372036854775807:
        v11 := []int{32767, -9223372036854775808, -277, 698}
        _ = v11
    case 2147483647:
        v12 := []int{2147483647, 128, 0, -9223372036854775808}
        _ = v12
    case 32767:
        switch 680 {
        case 9223372036854775807:
            v13 := []int{int(p1), 253}
            _ = v13
        case 32768:
            fmt.Println(p2)
        }
    default:
        var v14 int32 = int32(v10)
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return v10
}

func main() {
    var v1 int16 = -310
    v2 := make(map[int]bool)
    v3 := []uint32{42, 0}
    var v4 int16 = (v1 % ((v1) | 1))
    var v5 float64 = -2.25
    var v6 float32 = (((((1.5 - float32(v5)) / 1.0) - 1.5) - 1.5) - -2.25)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v15 := GenS2{F0: 9223372036854775807, F1: 161}
    v4 = (v4 << ((0) & 63))
    fmt.Println(v6)
    v5 = (0.0 - v5)
    v5 = v5
    v4 = v4
    v3 = append(v3, 42)
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    if false {
        switch ((("foo bar" + "foo bar") + (("" + "") + "hello")) + (("foo bar" + (("\n\t" + (("a" + "\n\t") + "\n\t")) + ("a" + "foo bar"))) + "foo bar")) {
        case "foo bar":
            if (0 <= 1) {
                v16 := GenS2{F0: 77, F1: 18446744073709551615}
                fmt.Println(v6)
                if ("a" != "foo bar") {
                    var v17 string = "foo bar"
                    if true {
                        fmt.Println(v4, v17, v17)
                        fmt.Println(v6)
                        v18 := GenS1{F0: 2, F1: 4294967295}
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 4; fi19++ {
                        var v20 string = "a"
                        _ = fi19
                        _ = v20
                    }
                    _ = v17
                }
                _ = v16
            }
        case "":
            v21 := GenS2{F0: 9223372036854775807, F1: 796}
            _ = v21
        }
        v22 := GenS2{F0: 42, F1: 117}
        _ = v22
    }
    v5 -= 1e10
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
