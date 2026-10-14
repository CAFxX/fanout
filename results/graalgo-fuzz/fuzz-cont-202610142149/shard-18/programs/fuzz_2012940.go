package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint32 }
type GenS2 struct { F0 string; F1 uint }
func genFn0(p0 int8) int16 {
    p0 -= 113
    v7 := []int{-9223372036854775808}
    if (42 > 1) {
        var v8 int64 = 299
        if false {
            switch "\n\t" {
            case "foo bar":
                v8 -= (v8 - v8)
            case "a":
                v9 := GenS1{F0: 65535, F1: 4294967295, F2: 65535}
                _ = v9
            case "\n\t":
                var v10 uint32 = 2147483648
                _ = v10
            default:
                var v11 uint32 = 935
                _ = v11
            }
            if (3.14159 > ((-2.25 * -2.25) * 1.5)) {
                var v12 uint8 = 126
                v8 = (v8 | 0)
                _ = v12
            }
            v13 := GenS1{F0: 42, F1: 179, F2: 504}
            _ = v13
        } else {
            var v14 uint = 171
            var v15 int64 = v8
            v14 = (v14 % (((v14 * v14)) | 1))
            _ = v14
            _ = v15
        }
        _ = v8
    }
    p0 = -105
    var v16 int64 = 0
    _ = p0
    _ = v16
    _ = v7
    return 42
}

func main() {
    v1 := []uint32{4294967295, 42, 166, 1}
    var v2 bool = (!true)
    v3 := []string{"foo bar", "foo bar", (("hello" + ("a" + ("hello" + ""))) + "hello"), ("\n\t" + ("" + "\n\t"))}
    v4 := []string{"\n\t", "foo bar", "hello", "a"}
    v5 := make(map[uint32]int)
    var v6 int8 = 100
    for fi17 := 0; fi17 < 5; fi17++ {
        v18 := GenS2{F0: ("foo bar" + "\n\t"), F1: 0}
        _ = fi17
        _ = v18
    }
    if v2 {
        var v19 float64 = (-2.25 - (float64((0.1 * (0.0 - (0.1 + 0.1)))) * (1e10 + 0.0)))
        _ = v19
    } else {
        v6 += (v6 * (v6 % (((v6 << ((634) & 63))) | 1)))
        var v20 float64 = 0.1
        _ = v20
    }
    var v21 uint8 = 126
    var v22 int = 42
    var v23 float64 = 3.14159
    if (!v2) {
        fmt.Println(v6, v2)
        var v24 int = int(v6)
        for fi25 := 0; fi25 < 16; fi25++ {
            v26 := GenS2{F0: "", F1: 95}
            v27 := GenS1{F0: 467, F1: 18446744073709551615, F2: 2147483648}
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = v24
    } else {
        _ = len(v5)
        for fi28 := 0; fi28 < 18; fi28++ {
            v6 = v6
            v29 := GenS1{F0: 65535, F1: 0, F2: 42}
            if (-85 < 0) {
                v2 = v2
                v6 = 38
            }
            _ = fi28
            _ = v29
        }
    }
    v30 := GenS1{F0: 1, F1: 912, F2: 2147483648}
    v23 = 0.0
    if len(v3) > 0 {
        v3[0] = ((((("foo bar" + "hello") + "foo bar") + "\n\t") + "") + "hello")
    }
    v21 = v21
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
