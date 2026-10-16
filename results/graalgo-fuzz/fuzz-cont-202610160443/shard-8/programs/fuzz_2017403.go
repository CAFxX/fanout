package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int8; F1 int64 }
func genFn0(p0 uint8) uint {
    v5 := []int{2147483647, -9223372036854775808, 2147483647}
    for fi6 := 0; fi6 < 3; fi6++ {
        if (42 <= -1) {
            var v7 int = (fi6 + (fi6 << ((4294967295) & 63)))
            v8 := GenS1{F0: 1, F1: 127}
            for fi9 := 0; fi9 < 18; fi9++ {
                var v10 uint64 = 303
                v11 := GenS1{F0: 430, F1: (p0 << ((4294967295) & 63))}
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = v7
            _ = v8
        }
        var v12 uint8 = (^(p0))
        if false {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if len(v5) > 0 {
                v5[0] = fi6
            }
        }
        _ = fi6
        _ = v12
    }
    fmt.Println(p0)
    _ = p0
    _ = v5
    return 65535
}
func genFn1(p0 int32) uint32 {
    v13 := GenS1{F0: (p0 % (((p0 >> ((884) & 63))) | 1)), F1: 128}
    v14 := []int{}
    p0 = 1
    var v15 uint = 660
    _ = p0
    _ = v13
    _ = v14
    _ = v15
    return 2147483647
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = (767 > 0)
    var v3 int64 = -106
    var v4 int16 = 497
    v16 := GenS2{F0: v2, F1: 0}
    for fi17 := 0; fi17 < 15; fi17++ {
        v4 = (-(v4))
        var v18 int16 = (v4 ^ (v4 << ((18446744073709551615) & 63)))
        _ = fi17
        _ = v18
    }
    switch 434 {
    case 0:
        for fi19 := 0; fi19 < 9; fi19++ {
            var v20 int8 = -1
            var v21 bool = (226 <= genFn0(70))
            v20 += (v20 - (v20 - (v20 * (v20 / ((127) | 1)))))
            _ = fi19
            _ = v20
            _ = v21
        }
    case 128:
        v4 = v4
    case 32767:
        v3 = v3
    }
    if ((("hello" + "hello") + "foo bar") <= "a") {
        v4 = 657
        v22 := GenS2{F0: true, F1: 818}
        _ = v22
    }
    v23 := GenS3{F0: -1, F1: v3}
    if v2 {
        var v24 uint64 = 42
        v3 += v3
        if ((("hello" + "\n\t") + "\n\t") < "a") {
            fmt.Println(v24, v2)
            v2 = true
            if v2 {
                var v25 string = ("foo bar" + "hello")
                v2 = v2
                _ = v25
            }
        }
        _ = v24
    } else {
        fmt.Println(v2)
        fmt.Println(v3, v2, v3)
    }
    fmt.Println(v2, v4, v2)
    switch 4294967295 {
    case 32767:
        if v2 {
            fmt.Println(v4, v3, v2)
        }
    default:
        for fi26 := 0; fi26 < 3; fi26++ {
            v4 = v4
            for fi27 := 0; fi27 < 5; fi27++ {
                v28 := []int{1, (fi27 >> ((4294967295) & 63)), fi26, (fi26 * (fi26 - fi27))}
                v29 := v1[(^(fi26))]
                if len(v28) >= 2 {
                    v28 = v28[1:len(v28)]
                }
                _ = fi27
                _ = v28
                _ = v29
            }
            _ = fi26
        }
    }
    switch "\n\t" {
    case "hello":
        var v30 uint = 65535
        _ = v30
    default:
        for fi31 := 0; fi31 < 19; fi31++ {
            var v32 float32 = (1.5 + 0.1)
            _ = fi31
            _ = v32
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
