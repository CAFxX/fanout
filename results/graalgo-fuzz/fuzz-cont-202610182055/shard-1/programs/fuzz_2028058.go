package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := []string{}
    v2 := []string{(("foo bar" + "") + "foo bar"), "", ((("hello" + "a") + "foo bar") + "foo bar"), ""}
    var v3 string = "\n\t"
    var v4 float64 = -2.25
    v5 := make(map[int]bool)
    var v6 uint64 = 124
    if (false && (255 == 255)) {
        switch v3 {
        case "hello":
            delete(v5, -362)
        case "foo bar":
            v3 = v3
        default:
            if true {
                for fi7 := 0; fi7 < 19; fi7++ {
                    v6 = v6
                    _ = fi7
                }
                fmt.Println(v4)
            } else {
                fmt.Println(v4, v4)
            }
        }
        for fi8 := 0; fi8 < 4; fi8++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi8
        }
    } else {
        v4 += (v4 * 1e10)
        v9 := GenS1{F0: 42}
        _ = v9
    }
    var v10 int64 = -9223372036854775808
    var v11 float32 = (((3.14159 * (1e10 / 0.1)) - 3.14159) - ((1e10 + 0.1) / -2.25))
    for fi12 := 0; fi12 < 19; fi12++ {
        if true {
            v10 = (v10 / ((v10) | 1))
            v4 = (v4 * v4)
        } else {
            switch (v3 + v3) {
            case "\n\t":
                delete(v5, fi12)
            }
        }
        _ = fi12
    }
    v5[128] = (-907 == 421)
    v5[-2147483648] = (370 >= 42)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v13 uint64 = (v6 * 1)
    var v14 int16 = -456
    var v15 uint32 = 733
    var v16 uint64 = v13
    if true {
        var v17 float32 = 1.5
        if false {
            for fi18 := 0; fi18 < 8; fi18++ {
                v14 -= (v14 + -1)
                _ = len(v5)
                var v19 bool = false
                _ = fi18
                _ = v19
            }
            var v20 int32 = -406
            _ = v20
        } else {
            v21 := GenS2{F0: -454, F1: true, F2: v15}
            v22 := GenS2{F0: -89, F1: (18446744073709551615 != uint(v15)), F2: v15}
            v1 = append(v1, (("hello" + "hello") + (v3 + ("hello" + (v3 + v3)))))
            _ = v21
            _ = v22
        }
        v23 := GenS3{F0: (v15 * (^(v15))), F1: 1}
        _ = v17
        _ = v23
    } else {
        fmt.Println(v3, v14)
    }
    if (v4 == v4) {
        var v24 int16 = 32767
        v6 = (v16 >> ((941) & 63))
        _ = v24
    }
    v25 := GenS4{F0: -2147483648}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
