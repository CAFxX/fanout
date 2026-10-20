package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 string }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int8; F1 uint8 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[string]uint32)
    var v3 uint32 = 42
    v4 := []string{}
    var v5 int16 = 1000
    v6 := []uint32{240, (v3 << (((v1 | 42)) & 63))}
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS1{F0: 868, F1: 0, F2: (("" + "foo bar") + ("" + ("hello" + "")))}
        var v9 uint = 255
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v3)
    var v10 uint16 = 255
    v11 := GenS1{F0: 0, F1: -1, F2: "hello"}
    v12 := v2[("foo bar" + ("" + "\n\t"))]
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v13 string = "a"
    var v14 int16 = -816
    if (1 != 127) {
        for fi15 := 0; fi15 < 11; fi15++ {
            v12 = 4294967295
            fmt.Println(v12, v3, v3)
            _ = fi15
        }
        v16 := GenS2{F0: (v10 ^ v10), F1: -947}
        v17 := GenS3{F0: 50, F1: 189}
        _ = v16
        _ = v17
    } else {
        if (42 != 82) {
            v10 -= v10
            v13 = (v13 + "a")
        }
    }
    var v18 uint64 = 166
    fmt.Println(v1, v3)
    v19 := GenS1{F0: 2147483647, F1: 0, F2: ((v13 + v13) + "")}
    var v20 uint64 = v18
    var v21 uint = (v1 | v1)
    _ = len(v4)
    for fi22 := 0; fi22 < 11; fi22++ {
        if ("\n\t" > (v13 + v13)) {
            if (331 > -1) {
                v13 = ("\n\t" + v13)
                if len(v6) > 0 {
                    v6[0] = 731
                }
            } else {
                v23 := GenS1{F0: -1, F1: 897, F2: v13}
                _ = v23
            }
            v14 = (v14 & (v14 * -1))
            v14 += (-(v14))
        }
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
