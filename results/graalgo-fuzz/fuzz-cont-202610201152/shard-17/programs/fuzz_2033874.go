package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := make(map[string]bool)
    var v2 int = -2147483648
    var v3 float32 = ((float32(0.1) - 1e10) * 1.5)
    var v4 string = "hello"
    var v5 uint64 = 4294967295
    v1["foo bar"] = true
    var v6 int = 127
    if (0 <= 42) {
        fmt.Println(v3, v3, v6)
        fmt.Println(v2, v6)
    }
    var v7 uint32 = 4294967295
    v3 += -2.25
    var v8 uint8 = 127
    v1["a"] = (!(65535 != 717))
    v2 -= (-(v2))
    if (!(true && false)) {
        v4 = "foo bar"
    }
    if (1.5 != ((1e10 * (1e10 * ((1e10 * (0.0 * 0.0)) - 1e10))) - (1e10 + float64((0.0 + 0.1))))) {
        v2 += (v2 * v6)
        for fi9 := 0; fi9 < 4; fi9++ {
            _ = len(v1)
            _ = fi9
        }
        v10 := GenS1{F0: 358}
        _ = v10
    } else {
        var v11 int32 = 803
        var v12 string = v4
        _ = v11
        _ = v12
    }
    if false {
        var v13 uint8 = v8
        v5 += (v5 - v5)
        v14 := GenS1{F0: 334}
        _ = v13
        _ = v14
    } else {
        for fi15 := 0; fi15 < 1; fi15++ {
            if (v4 != "hello") {
                var v16 bool = (((0.1 + float64(((float32((((1.5 + 1e10) * 0.0) - ((1.5 / 0.1) * 1.5))) - (v3 * float32(0.0))) * -2.25))) * 0.1) > 1e10)
                v17 := GenS2{F0: 1000}
                _ = v16
                _ = v17
            } else {
                for fi18 := 0; fi18 < 16; fi18++ {
                    v8 = 255
                    _ = fi18
                }
                _ = len(v1)
            }
            var v19 float32 = float32(1e10)
            v6 -= v6
            _ = fi15
            _ = v19
        }
    }
    if (true || (0 > -1)) {
        fmt.Println(v7, v5)
        var v20 uint64 = (^(v5))
        _ = v20
    } else {
        v21 := GenS3{F0: 0}
        var v22 uint = uint(v7)
        var v23 int32 = -1
        _ = v21
        _ = v22
        _ = v23
    }
    var v24 uint16 = 692
    v7 = uint32(v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
