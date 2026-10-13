package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int32; F2 int8 }
func genFn0(p0 uint8, p1 uint64, p2 uint16) uint64 {
    switch 9223372036854775807 {
    case 42:
        p0 = p0
    }
    switch 32768 {
    case 42:
        var v7 uint32 = 65535
        _ = v7
    case 32767:
        var v8 string = "a"
        _ = v8
    default:
        p1 = p1
    }
    switch 2147483647 {
    case -1:
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 << ((1) & 63))
}

func main() {
    var v1 uint = 1
    var v2 uint16 = 255
    var v3 int = 74
    var v4 int8 = 0
    var v5 string = "foo bar"
    v6 := []string{(("foo bar" + v5) + v5), "a"}
    var v9 float32 = 3.14159
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 int = (fi10 / ((v3) | 1))
        v1 -= v1
        v2 += (v2 | v2)
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        for fi13 := 0; fi13 < 1; fi13++ {
            if true {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            } else {
                switch int(v4) {
                case 32767:
                    v14 := GenS1{F0: v5}
                    _ = v14
                case 1:
                    switch "hello" {
                    case "":
                        v9 -= v9
                    case "\n\t":
                        fmt.Println(v2, v9, v4)
                    case "foo bar":
                        v15 := GenS1{F0: v5}
                        _ = v15
                    }
                case -9223372036854775808:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                default:
                    switch (fi12 + -2147483648) {
                    case -276:
                        v2 = 65535
                    }
                }
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            v4 = -32
            if (v5 <= ("hello" + ("\n\t" + v5))) {
                v2 += (v2 / (((^(v2))) | 1))
                if true {
                    var v16 bool = true
                    var v17 uint16 = v2
                    _ = v16
                    _ = v17
                } else {
                    var v18 int32 = -245
                    _ = v18
                }
                var v19 float64 = 1.5
                var v20 float32 = 0.0
                _ = v19
                _ = v20
            }
            _ = fi13
        }
        _ = fi12
    }
    v9 -= (v9 + 3.14159)
    v21 := GenS2{F0: 0, F1: 945, F2: -18}
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 uint32 = 791
        var v24 uint32 = (v23 * (v23 - (v23 << (((v1 >> ((v1) & 63))) & 63))))
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 int64 = 950
    v2 = (v2 * (v2 >> ((18446744073709551615) & 63)))
    v2 = v2
    var v26 int = v3
    v4 += (v4 ^ v4)
    var v27 uint16 = v2
    var v28 int = (-(v26))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
