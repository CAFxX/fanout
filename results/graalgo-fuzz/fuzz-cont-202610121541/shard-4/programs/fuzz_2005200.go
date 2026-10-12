package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int8; F1 int16; F2 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = 0
    v3 := make(map[uint32]int)
    v4 := []string{"a", "", "\n\t"}
    var v5 string = ("\n\t" + "")
    v6 := make(map[int]string)
    v5 = v5
    var v7 int64 = v2
    if false {
        v4 = append(v4, v5)
        if ((-2.25 * -2.25) != 1.5) {
            v5 = ""
            v7 = 9223372036854775807
        } else {
            var v8 int8 = -80
            for fi9 := 0; fi9 < 14; fi9++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi9
            }
            _ = v8
        }
    }
    fmt.Println(v5)
    switch -656 {
    case 32767:
        v10 := GenS1{F0: 2, F1: (127 != 1)}
        _ = v10
    case -9223372036854775808:
        if (-128 <= 0) {
            v4 = append(v4, v5)
            v4 = append(v4, "")
        }
    case 42:
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 string = "hello"
            var v13 int = (fi11 + (fi11 + 32768))
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    v2 = 1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v4)
    v2 -= 243
    var v14 int = -2147483648
    var v15 bool = (float32(1e10) <= 0.1)
    if len(v4) > 0 {
        v4[0] = v5
    }
    v16 := v6[(v14 | v14)]
    for fi17 := 0; fi17 < 4; fi17++ {
        fmt.Println(v16)
        v18 := GenS2{F0: -1, F1: 32767, F2: -630}
        if (!v15) {
            var v19 int32 = 42
            var v20 int32 = v19
            if len(v4) > 0 {
                v4[0] = ((v16 + "") + v5)
            }
            _ = v19
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
