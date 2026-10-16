package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint16; F1 int8 }
func genFn0(p0 int16, p1 string) uint {
    v6 := GenS1{F0: 882, F1: 18446744073709551615}
    v7 := []int{127, -9223372036854775808}
    var v8 float64 = (0.0 - (1.5 / 1e10))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 65535
}

func main() {
    v1 := []int{}
    var v2 uint16 = 32767
    var v3 uint = 0
    v4 := make(map[int]bool)
    var v5 uint32 = 42
    var v9 uint = 2
    v9 += (^(v3))
    if (true || false) {
        var v10 uint16 = (v2 >> ((v9) & 63))
        if false {
            v2 -= (^(v10))
        }
        v2 -= v2
        _ = v10
    } else {
        v9 += v3
    }
    var v11 float64 = 1e10
    v12 := GenS1{F0: v5, F1: 192}
    fmt.Println(v3)
    if false {
        if false {
            v13 := GenS1{F0: 2147483647, F1: 244}
            _ = v13
        }
        for fi14 := 0; fi14 < 8; fi14++ {
            for fi15 := 0; fi15 < 11; fi15++ {
                v16 := GenS2{F0: ((false && (-2.25 != (v11 + (v11 - 3.14159)))) || true)}
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
        v4[0] = ((("foo bar" <= "a") && ("hello" != ((("a" + ("hello" + "hello")) + "foo bar") + "a"))) || false)
    } else {
        var v17 int32 = -1
        v1 = append(v1, -1)
        var v18 float64 = ((v11 / 1.5) / (v11 - v11))
        _ = v17
        _ = v18
    }
    fmt.Println(v3, v3, v11)
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 uint32 = (v5 << ((v9) & 63))
        _ = fi19
        _ = v20
    }
    if ((!(!((1e10 != (1e10 * (1e10 - 1e10))) && true))) || (false && true)) {
        for fi21 := 0; fi21 < 16; fi21++ {
            var v22 int = (fi21 >> ((4294967295) & 63))
            _ = fi21
            _ = v22
        }
        switch 128 {
        case 142:
            fmt.Println(v5, v3)
        case 2147483647:
            v23 := GenS2{F0: (true || true)}
            _ = v23
        case 32767:
            var v24 bool = true
            _ = v24
        }
        var v25 uint32 = v5
        _ = v25
    } else {
        v26 := GenS3{F0: 127, F1: 809}
        _ = v26
    }
    switch "a" {
    case "hello":
        v3 -= 15
    case "":
        if false {
            v5 = 783
        } else {
            var v27 int64 = 9223372036854775807
            _ = v27
        }
    case "a":
        if (!false) {
            delete(v4, int(v3))
            v3 -= (v9 - v3)
        } else {
            v5 = v5
        }
    default:
        if true {
            v3 = 295
            var v28 uint8 = 1
            for fi29 := 0; fi29 < 17; fi29++ {
                v30 := GenS1{F0: 1, F1: 18446744073709551615}
                var v31 int16 = 854
                if (!(2147483647 < -1)) {
                    v32 := GenS4{F0: (v2 >> (((^(v3))) & 63)), F1: -61}
                    fmt.Println(v3)
                    v28 = (^(v28))
                    _ = v32
                } else {
                    v9 += genFn0(263, "hello")
                }
                _ = fi29
                _ = v30
                _ = v31
            }
            _ = v28
        }
    }
    fmt.Println(v3)
    fmt.Println(v2, v11, v9)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
