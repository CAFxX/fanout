package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int; F1 int64; F2 int }
func genFn0(p0 uint8, p1 int32, p2 uint8) int16 {
    fmt.Println(p0, p1)
    v6 := GenS1{F0: 9223372036854775807}
    var v7 uint8 = (p0 ^ 36)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := []int{16, 9223372036854775807}
    var v2 string = "\n\t"
    v3 := make(map[uint32]uint32)
    var v4 int64 = -468
    var v5 string = ("hello" + v2)
    v5 = "hello"
    v4 -= (-(v4))
    if (false || true) {
        v8 := GenS2{F0: -2147483648, F1: 42}
        var v9 uint64 = 18446744073709551615
        v9 -= v9
        _ = v8
        _ = v9
    } else {
        v5 = v5
        var v10 uint16 = 726
        _ = v10
    }
    fmt.Println(v2, v4, v4)
    for fi11 := 0; fi11 < 11; fi11++ {
        if (genFn0(255, -2147483648, 2) >= -1) {
            for fi12 := 0; fi12 < 2; fi12++ {
                v1 = append(v1, (fi11 << ((2) & 63)))
                _ = fi12
            }
            if len(v1) > 0 {
                v1[0] = 32767
            }
            switch (v2 + v2) {
            case "hello":
                v4 -= (v4 ^ 0)
            case "\n\t":
                if (((((((v2 + v5) + v5) + "") + "") + v2) + v5) != "a") {
                    v13 := GenS3{F0: fi11, F1: (v4 >> ((42) & 63)), F2: 4294967295}
                    var v14 uint32 = 194
                    var v15 int16 = 42
                    _ = v13
                    _ = v14
                    _ = v15
                } else {
                    _ = len(v3)
                    var v16 int64 = v4
                    fmt.Println(v5)
                    _ = v16
                }
            case "foo bar":
                v5 = "a"
            }
        } else {
            v5 = ((v2 + ("\n\t" + ("a" + (("\n\t" + v2) + ("\n\t" + "\n\t"))))) + v2)
            fmt.Println(fi11)
        }
        _ = fi11
    }
    v17 := GenS2{F0: -9223372036854775808, F1: 128}
    v18 := GenS3{F0: -656, F1: v4, F2: 9223372036854775807}
    for fi19 := 0; fi19 < 8; fi19++ {
        for fi20 := 0; fi20 < 8; fi20++ {
            var v21 int8 = -128
            _ = fi20
            _ = v21
        }
        if len(v1) > 0 {
            v1[0] = 32768
        }
        fmt.Println(v4, fi19, fi19)
        _ = fi19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
