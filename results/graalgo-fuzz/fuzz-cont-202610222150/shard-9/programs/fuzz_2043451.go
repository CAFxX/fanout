package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }
func genFn0(p0 uint8, p1 int16, p2 int64) int64 {
    v5 := GenS1{F0: p2, F1: p0, F2: 696}
    p2 = (p2 * (p2 + (p2 >> ((18446744073709551615) & 63))))
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p2
}

func main() {
    var v1 int64 = 807
    v2 := []string{"", ("a" + ""), "\n\t"}
    var v3 uint8 = 128
    var v4 float64 = 0.0
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 int32 = -2147483648
            switch ("" + (("" + ("hello" + "")) + (("foo bar" + (("hello" + "") + "foo bar")) + ""))) {
            case "hello":
                v8 = v8
            case "\n\t":
                v8 += v8
            case "":
                fmt.Println(v1, fi7, v3)
                var v9 uint = 42
                var v10 int8 = 0
                _ = v10
                _ = v9
            }
            _ = fi7
            _ = v8
        }
        switch 2147483647 {
        case 42:
            if true {
                var v11 float64 = 1.5
                _ = v11
            }
        case 977:
            var v12 int32 = int32(v1)
            _ = v12
        case 181:
            v3 = (v3 % (((v3 % ((v3) | 1))) | 1))
        default:
            var v13 float32 = (float32(v4) * 0.0)
            _ = v13
        }
        _ = fi6
    }
    switch (("foo bar" + "hello") + "\n\t") {
    case "foo bar":
        v14 := GenS2{F0: 18446744073709551615}
        _ = v14
    case "a":
        if (("foo bar" != "foo bar") && (false || true)) {
            for fi15 := 0; fi15 < 8; fi15++ {
                v1 += (v1 & v1)
                if true {
                    fmt.Println(v3, fi15, fi15)
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                }
                _ = fi15
            }
        } else {
            v3 = v3
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
        }
    case "\n\t":
        v16 := GenS3{F0: -1}
        _ = v16
    default:
        fmt.Println(v4, v3, v3)
    }
    v2 = append(v2, ("" + "hello"))
    var v17 uint32 = 4294967295
    var v18 int32 = 567
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v19 bool = true
    var v20 float32 = 0.1
    v21 := GenS2{F0: 18446744073709551615}
    fmt.Println(v20, v1)
    fmt.Println(v4)
    var v22 float32 = float32(0.1)
    v18 = (v18 ^ 2147483647)
    var v23 int32 = v18
    if true {
        for fi24 := 0; fi24 < 10; fi24++ {
            v2 = append(v2, ("hello" + "hello"))
            _ = fi24
        }
    }
    v20 = v20
    fmt.Println(v22, v18)
    v19 = v19
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
