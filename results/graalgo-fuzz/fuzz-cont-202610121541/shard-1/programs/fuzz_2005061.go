package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint64; F1 int32; F2 int }

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 42
    v3 := []string{}
    if ((127 == 46) || false) {
        var v4 int64 = -1
        _ = v4
    } else {
        _ = len(v1)
        var v5 int32 = v2
        v6 := GenS1{F0: 1, F1: -1}
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 42, F1: -596}
    v1[2147483647] = 9223372036854775807
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        if (!true) {
            v2 += (v2 | 42)
            for fi9 := 0; fi9 < 12; fi9++ {
                switch (fi8 << ((42) & 63)) {
                case 186:
                    fmt.Println(fi9, fi9, fi9)
                case 42:
                    v2 = (v2 | 442)
                }
                var v10 uint64 = 4294967295
                _ = fi9
                _ = v10
            }
            var v11 bool = false
            _ = v11
        }
        for fi12 := 0; fi12 < 11; fi12++ {
            v2 = (v2 & (v2 & 469))
            if ((-10 == int16(fi8)) && (true || false)) {
                fmt.Println(fi8)
                if true {
                    switch "" {
                    case "\n\t":
                        var v13 string = ("" + "")
                        _ = v13
                    }
                }
            } else {
                var v14 int16 = 0
                if (2 != 2) {
                    fmt.Println(fi12)
                    var v15 uint32 = 2147483647
                    v16 := GenS1{F0: 42, F1: 9223372036854775807}
                    var v17 string = ((("a" + "a") + ("hello" + ("hello" + "a"))) + "foo bar")
                    _ = len(v3)
                    _ = v15
                    _ = v16
                    _ = v17
                }
                v18 := GenS2{F0: 18446744073709551615, F1: v2, F2: int(v2)}
                var v19 int16 = (-(v14))
                v2 = (v2 % ((v2) | 1))
                _ = v14
                _ = v18
                _ = v19
            }
            _ = fi12
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi8
    }
    v2 += (v2 ^ (^(v2)))
    var v20 int16 = 168
    v2 -= v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v7
}
