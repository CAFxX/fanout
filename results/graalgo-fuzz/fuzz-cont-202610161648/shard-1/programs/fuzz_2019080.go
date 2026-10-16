package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 int32 = -726
    v2 := []int{}
    var v3 uint16 = 2
    v4 := make(map[uint32]int)
    var v5 string = ""
    var v6 uint8 = 128
    v7 := GenS1{F0: v1}
    if ((v3 ^ v3) >= v3) {
        v5 = (v5 + "hello")
        for fi8 := 0; fi8 < 19; fi8++ {
            switch "a" {
            case "\n\t":
                var v9 uint32 = 2147483647
                _ = v9
            case "hello":
                var v10 int = (fi8 - (fi8 & (fi8 & 4294967295)))
                _ = v10
            default:
                v11 := GenS1{F0: (v1 | (v1 | (v1 >> ((65535) & 63))))}
                _ = v11
            }
            _ = fi8
        }
        v12 := GenS1{F0: v1}
        _ = v12
    } else {
        var v13 uint16 = v3
        v3 += (^(v3))
        _ = v13
    }
    v6 -= uint8(v3)
    if ((32767 != 42) || true) {
        v6 += v6
        for fi14 := 0; fi14 < 17; fi14++ {
            v1 -= -2147483648
            var v15 int = 127
            _ = fi14
            _ = v15
        }
        if ((868 > 0) && false) {
            fmt.Println(v3, v5, v5)
        } else {
            switch "foo bar" {
            case "hello":
                v3 = (v3 * uint16(v6))
            case "":
                var v16 string = (v5 + "hello")
                _ = v16
            }
            v3 += (v3 | 32767)
            var v17 int64 = 9223372036854775807
            _ = v17
        }
    }
    var v18 uint8 = (v6 + 127)
    if (!(255 <= 2)) {
        delete(v4, 2147483648)
        delete(v4, 209)
        _ = len(v4)
    } else {
        var v19 uint16 = v3
        v18 += v6
        _ = v19
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        for fi21 := 0; fi21 < 6; fi21++ {
            var v22 int16 = -1
            _ = fi21
            _ = v22
        }
        fmt.Println(v18, fi20, v5)
        fmt.Println(fi20, v1, v18)
        _ = fi20
    }
    if ((2147483648 < 0) || ((("" + v5) + v5) != (v5 + v5))) {
        v18 -= v6
    } else {
        var v23 string = (("a" + "a") + "foo bar")
        v18 -= v18
        _ = v23
    }
    if false {
        var v24 int64 = -932
        v25 := GenS2{F0: (v18 ^ uint8(v3))}
        _ = v24
        _ = v25
    }
    v3 += (v3 | v3)
    if false {
        if (true || true) {
            fmt.Println(v6, v5, v6)
            v1 -= (v1 & v1)
        }
        fmt.Println(v1)
        _ = len(v2)
    }
    v26 := GenS3{F0: 1}
    v3 -= (v3 >> ((65535) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
