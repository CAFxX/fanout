package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint64; F2 uint8 }
func genFn0(p0 uint8) int {
    switch "hello" {
    case "hello":
        if true {
            var v7 bool = (false || true)
            _ = v7
        }
    }
    switch ("foo bar" + ("hello" + "foo bar")) {
    case "a":
        var v8 uint32 = 65535
        _ = v8
    }
    _ = p0
    return 127
}

func main() {
    var v1 float64 = (0.0 * (((3.14159 - 0.1) * 0.0) - ((((float64(-2.25) / 1e10) - 3.14159) - 0.1) - 0.0)))
    v2 := make(map[int]int)
    v3 := []int{220, 0, 32768, -2147483648}
    var v4 uint64 = 42
    var v5 uint16 = 0
    var v6 int = 1
    var v9 uint = uint(v4)
    for fi10 := 0; fi10 < 12; fi10++ {
        v5 = (v5 * (v5 >> ((v9) & 63)))
        switch (((("\n\t" + "a") + "hello") + "\n\t") + "") {
        case "":
            fmt.Println(v9)
        default:
            if false {
                v11 := GenS1{F0: ""}
                _ = v11
            } else {
                v5 = v5
                v4 -= (v4 | (v4 & 2))
                fmt.Println(v5, fi10)
            }
        }
        _ = fi10
    }
    if false {
        fmt.Println(v4, v4)
        switch ("a" + ("foo bar" + "\n\t")) {
        case "foo bar":
            var v12 int32 = -411
            _ = v12
        case "":
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 int64 = 0
                v14 = (v14 / ((v14) | 1))
                var v15 int64 = (^(v14))
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        if ((v1 - ((v1 + v1) * v1)) < v1) {
            if true {
                v16 := v2[(v6 ^ (v6 << ((4294967295) & 63)))]
                _ = v16
            } else {
                var v17 int16 = 0
                switch 2147483647 {
                case 820:
                    v6 = (-(v6))
                default:
                    fmt.Println(v4, v4)
                }
                _ = v17
            }
            v3 = append(v3, (v6 * v6))
            for fi18 := 0; fi18 < 17; fi18++ {
                v19 := GenS2{F0: "hello", F1: v4, F2: 135}
                _ = fi18
                _ = v19
            }
        }
    }
    v20 := GenS1{F0: ("\n\t" + "hello")}
    fmt.Println(v1)
    fmt.Println(v5, v9, v4)
    fmt.Println(v9, v1)
    var v21 int16 = 32767
    var v22 uint32 = 253
    switch (v6 & genFn0(95)) {
    case 4294967295:
        fmt.Println(v1, v22)
    default:
        v23 := GenS2{F0: (("\n\t" + "a") + ""), F1: v4, F2: 128}
        _ = v23
    }
    if true {
        fmt.Println(v22, v22)
        var v24 uint8 = 32
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
