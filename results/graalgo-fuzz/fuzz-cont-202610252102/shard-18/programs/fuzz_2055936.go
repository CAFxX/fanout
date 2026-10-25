package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint; F1 int; F2 uint16 }

func main() {
    var v1 float32 = -2.25
    v2 := make(map[int]int)
    var v3 int16 = -1
    if (!true) {
        v1 = 3.14159
        v3 += (-(v3))
        if false {
            for fi4 := 0; fi4 < 9; fi4++ {
                var v5 int64 = -1
                v6 := []int{(fi4 % ((fi4) | 1)), (fi4 >> ((42) & 63)), 32767, (fi4 % (((fi4 | (fi4 | fi4))) | 1))}
                for fi7 := 0; fi7 < 16; fi7++ {
                    v8 := GenS1{F0: 734, F1: (("" + ("" + "hello")) + ("" + ""))}
                    var v9 uint32 = 692
                    switch (("\n\t" + (("" + "hello") + ((("foo bar" + "hello") + "\n\t") + "a"))) + "foo bar") {
                    case "foo bar":
                        var v10 uint = 998
                        _ = v10
                    }
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                _ = fi4
                _ = v5
                _ = v6
            }
            v1 -= 1.5
        } else {
            var v11 float32 = 1.5
            _ = len(v2)
            _ = v11
        }
    } else {
        fmt.Println(v1, v3)
    }
    switch 32767 {
    case 2147483647:
        fmt.Println(v1, v3)
    case 9223372036854775807:
        v1 = 0.1
    default:
        var v12 string = "a"
        _ = v12
    }
    v1 = (v1 + (v1 - v1))
    v1 = ((float32(3.14159) * v1) / v1)
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 uint32 = 2
        fmt.Println(v14)
        for fi15 := 0; fi15 < 1; fi15++ {
            if false {
                switch fi13 {
                case -9223372036854775808:
                    v3 -= (v3 ^ v3)
                }
                var v16 uint16 = 171
                _ = v16
            }
            if true {
                var v17 float32 = ((v1 / v1) / v1)
                _ = v17
            }
            v18 := GenS2{F0: 255, F1: 9223372036854775807, F2: 447}
            _ = fi15
            _ = v18
        }
        _ = fi13
        _ = v14
    }
    if false {
        v19 := GenS2{F0: 42, F1: 0, F2: 32767}
        v20 := GenS1{F0: 1, F1: ""}
        _ = v19
        _ = v20
    }
    v21 := v2[-2147483648]
    v2[(v21 | 42)] = 42
    for fi22 := 0; fi22 < 6; fi22++ {
        if ("foo bar" <= "hello") {
            var v23 int64 = 9223372036854775807
            _ = v23
        }
        for fi24 := 0; fi24 < 1; fi24++ {
            var v25 int8 = -1
            _ = fi24
            _ = v25
        }
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
}
