package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    var v3 int64 = 820
    var v4 int32 = -1
    var v5 uint16 = 32767
    for fi6 := 0; fi6 < 15; fi6++ {
        fmt.Println(v4, v4, fi6)
        _ = fi6
    }
    var v7 int8 = 127
    switch int(v4) {
    case 9223372036854775807:
        var v8 int32 = (v4 / (((v4 >> ((18446744073709551615) & 63))) | 1))
        _ = v8
    case 42:
        switch 0 {
        case 127:
            var v9 uint8 = 160
            _ = v9
        default:
            switch "\n\t" {
            case "\n\t":
                v4 = v4
            case "hello":
                v10 := []int{143, -9223372036854775808, -2147483648, 32768}
                _ = v10
            case "foo bar":
                v3 -= (v3 % (((^(v3))) | 1))
            }
        }
    }
    v4 = v4
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 int8 = 42
        v12 += 127
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 uint = 0
        v15 := []int{fi13, fi13}
        _ = fi13
        _ = v14
        _ = v15
    }
    v4 = v4
    var v16 int64 = v3
    v7 = (-(v7))
    var v17 int = 4294967295
    v17 += v17
    if false {
        v7 -= 1
    } else {
        if ((v5 > v5) || true) {
            var v18 float32 = 0.1
            if (int32(v16) > 2147483647) {
                switch (("a" + "hello") + "hello") {
                case "foo bar":
                    v19 := []int{v17, v17, v17}
                    _ = v19
                case "hello":
                    v20 := GenS1{F0: int8(v5)}
                    _ = v20
                default:
                    _ = len(v1)
                }
                for fi21 := 0; fi21 < 4; fi21++ {
                    var v22 int = -9223372036854775808
                    _ = fi21
                    _ = v22
                }
            } else {
                v23 := GenS1{F0: (-(v7))}
                _ = v23
            }
            _ = v18
        } else {
            for fi24 := 0; fi24 < 12; fi24++ {
                var v25 int32 = (v4 / ((2147483647) | 1))
                v26 := v2["\n\t"]
                _ = fi24
                _ = v25
                _ = v26
            }
        }
        fmt.Println(v3, v7)
    }
    v5 = (v5 * 32767)
    fmt.Println(v16, v17, v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
