package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int16; F2 int }

func main() {
    var v1 uint32 = 0
    var v2 int64 = 249
    v3 := []uint32{v1, 349}
    var v4 bool = ((0 <= 42) || (1 != 220))
    v5 := []string{"\n\t", (("a" + "foo bar") + ("foo bar" + (("hello" + "foo bar") + ("hello" + "hello")))), "a"}
    var v6 uint64 = 245
    var v7 uint64 = v6
    for fi8 := 0; fi8 < 14; fi8++ {
        fmt.Println(v1, v4)
        fmt.Println(v2)
        v4 = v4
        _ = fi8
    }
    var v9 int32 = -1
    if (635 != 359) {
        if v4 {
            var v10 int32 = (v9 - (^(v9)))
            v5 = append(v5, "\n\t")
            _ = v10
        } else {
            var v11 bool = (!(32767 < 121))
            _ = v11
        }
        for fi12 := 0; fi12 < 9; fi12++ {
            v2 = (-(v2))
            fmt.Println(v7)
            _ = fi12
        }
    }
    v13 := GenS1{F0: (v2 | v2)}
    if (true || (v4 && (v4 && v4))) {
        fmt.Println(v9)
    }
    if len(v5) > 0 {
        v5[0] = (("a" + "foo bar") + "")
    }
    if len(v5) > 0 {
        v5[0] = ("foo bar" + ("" + ("\n\t" + ("\n\t" + ("" + ("a" + ("a" + ("hello" + ((("a" + "") + "") + "")))))))))
    }
    v7 -= 2
    for fi14 := 0; fi14 < 5; fi14++ {
        for fi15 := 0; fi15 < 15; fi15++ {
            v9 = (v9 / (((v9 & v9)) | 1))
            fmt.Println(fi14, v1)
            v1 -= v1
            _ = fi15
        }
        _ = fi14
    }
    var v16 int16 = 32767
    v17 := GenS2{F0: 255, F1: v16, F2: int(v16)}
    for fi18 := 0; fi18 < 6; fi18++ {
        v2 = 9223372036854775807
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
