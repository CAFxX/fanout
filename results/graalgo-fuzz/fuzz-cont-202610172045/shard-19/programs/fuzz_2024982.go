package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 uint32 }

func main() {
    var v1 int16 = 42
    var v2 uint32 = 65535
    var v3 int32 = 42
    var v4 int = -798
    fmt.Println(v1, v4, v2)
    v5 := GenS1{F0: 1}
    if true {
        var v6 int8 = 42
        v7 := []int{-499, v4}
        var v8 float64 = float64(1e10)
        _ = v6
        _ = v7
        _ = v8
    }
    switch (v4 % ((129) | 1)) {
    case -9223372036854775808:
        v4 += (v4 & 553)
    case -2147483648:
        v3 -= v3
    case 32768:
        v2 += v2
    }
    var v9 int = v4
    v10 := GenS1{F0: 4294967295}
    v2 -= 65535
    v11 := GenS2{F0: (v1 + v1), F1: 1}
    v4 += v9
    switch ("a" + "foo bar") {
    case "foo bar":
        if (((-61 <= -13) || false) || ((v4 << ((195) & 63)) != v9)) {
            for fi12 := 0; fi12 < 8; fi12++ {
                v3 += 910
                if ((278 == 721) || true) {
                    v2 -= (v2 | v2)
                    v13 := GenS2{F0: 1000, F1: 964}
                    switch (("foo bar" + ("foo bar" + ("a" + "hello"))) + "foo bar") {
                    case "a":
                        v1 = 32767
                    case "":
                        if true {
                            v14 := GenS2{F0: 977, F1: (^(v2))}
                            v15 := GenS1{F0: 18446744073709551615}
                            _ = v14
                            _ = v15
                        }
                    case "foo bar":
                        fmt.Println(v4, fi12, v1)
                    }
                    _ = v13
                }
                fmt.Println(v9)
                _ = fi12
            }
        }
    case "\n\t":
        for fi16 := 0; fi16 < 19; fi16++ {
            fmt.Println(v9, v1, v2)
            v3 = (v3 << ((1) & 63))
            if true {
                v2 -= (^(v2))
                v9 = v9
                if true {
                    v17 := []int{int(v1), fi16, (v4 << ((uint(v2)) & 63)), (fi16 * (v4 - (fi16 / ((fi16) | 1))))}
                    for fi18 := 0; fi18 < 7; fi18++ {
                        _ = len(v17)
                        _ = fi18
                    }
                    v19 := GenS2{F0: 1000, F1: (v2 * v2)}
                    _ = v17
                    _ = v19
                } else {
                    v20 := []int{(v9 & -868)}
                    var v21 int = -1
                    var v22 uint8 = 128
                    _ = v20
                    _ = v21
                    _ = v22
                }
            } else {
                fmt.Println(v1)
            }
            _ = fi16
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
