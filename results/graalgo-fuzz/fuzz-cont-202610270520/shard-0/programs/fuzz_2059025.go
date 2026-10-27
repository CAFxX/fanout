package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
func genFn0(p0 uint32) uint64 {
    var v6 string = "a"
    if false {
        p0 -= (p0 ^ 2)
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 int32 = 2147483647
            fmt.Println(v6, v8, v8)
            p0 += 0
            _ = fi7
            _ = v8
        }
        if true {
            v6 = ((((("\n\t" + (v6 + v6)) + ("" + "hello")) + ("foo bar" + (v6 + v6))) + v6) + "hello")
        }
    } else {
        if (((-2.25 + ((0.1 + float32((0.1 + 1.5))) - 3.14159)) <= 1e10) && false) {
            var v9 string = v6
            v6 = ("hello" + v6)
            _ = v9
        }
        fmt.Println(v6, p0)
        if ((v6 != v6) || false) {
            p0 = (p0 << ((18446744073709551615) & 63))
        } else {
            var v10 string = v6
            if true {
                p0 += p0
                var v11 string = v6
                v12 := []int{1, -1, 127}
                _ = v11
                _ = v12
            } else {
                var v13 uint8 = 1
                for fi14 := 0; fi14 < 12; fi14++ {
                    switch "a" {
                    case "":
                        if (((0.0 * -2.25) + 3.14159) != 3.14159) {
                            p0 = p0
                            v10 = ((((("hello" + v10) + "") + v6) + v10) + v6)
                        }
                    case "a":
                        fmt.Println(p0, p0, v13)
                    default:
                        v6 = ((v6 + (v6 + (v10 + "hello"))) + "foo bar")
                    }
                    var v15 uint32 = p0
                    v13 = (v13 * 216)
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
            _ = v10
        }
    }
    fmt.Println(v6, p0)
    _ = p0
    _ = v6
    return 42
}

func main() {
    var v1 uint8 = 255
    var v2 int = -704
    v3 := make(map[string]uint32)
    var v4 int32 = -2147483648
    var v5 bool = true
    var v16 uint64 = 20
    var v17 float32 = 0.0
    v18 := GenS1{F0: 32767, F1: 1}
    v17 = float32((0.0 + 0.1))
    if (v17 == ((v17 - (v17 - 1e10)) * (v17 / v17))) {
        var v19 uint = 95
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 uint = (^(v19))
            v5 = v5
            _ = fi20
            _ = v21
        }
        if v5 {
            var v22 bool = (uint16(v1) != 70)
            _ = v22
        }
        _ = v19
    } else {
        if (v1 > (^(v1))) {
            v23 := GenS1{F0: -1, F1: 188}
            _ = v23
        }
        v24 := []int{v2, 0, (v2 << ((0) & 63)), -953}
        _ = v24
    }
    v25 := []int{}
    v16 -= 42
    switch (("hello" + ("" + "")) + ("hello" + "a")) {
    case "foo bar":
        switch "" {
        case "hello":
            v2 += (v2 % (((v2 << ((2) & 63))) | 1))
        }
    }
    switch (v2 % (((v2 << ((851) & 63))) | 1)) {
    case 127:
        switch "foo bar" {
        case "foo bar":
            var v26 int16 = -901
            _ = v26
        case "\n\t":
            var v27 int8 = 127
            _ = v27
        case "hello":
            v17 = float32(1e10)
        default:
            var v28 uint = 65535
            _ = v28
        }
    case 632:
        fmt.Println(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
