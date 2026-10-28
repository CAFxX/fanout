package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 uint64 = 2
    var v2 string = "a"
    var v3 float32 = float32(1.5)
    var v4 int64 = 9223372036854775807
    var v5 bool = true
    var v6 int64 = v4
    var v7 int64 = 0
    var v8 uint64 = (v1 >> ((18446744073709551615) & 63))
    if v5 {
        for fi9 := 0; fi9 < 8; fi9++ {
            switch v2 {
            case "hello":
                v8 -= (^(v1))
            default:
                for fi10 := 0; fi10 < 18; fi10++ {
                    var v11 uint32 = 2147483648
                    _ = fi10
                    _ = v11
                }
            }
            _ = fi9
        }
    } else {
        var v12 bool = (297 <= 0)
        var v13 int8 = 1
        _ = v12
        _ = v13
    }
    v14 := []int{}
    var v15 int32 = -2147483648
    v16 := GenS1{F0: 255}
    var v17 uint64 = v1
    if ((v6 ^ v4) >= (v7 >> ((255) & 63))) {
        var v18 uint64 = 1
        _ = v18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        var v20 uint32 = 970
        switch v2 {
        case "hello":
            for fi21 := 0; fi21 < 11; fi21++ {
                var v22 float32 = float32(0.0)
                _ = fi21
                _ = v22
            }
        case "foo bar":
            var v23 int = (-(fi19))
            _ = v23
        }
        if (!v5) {
            switch (fi19 - (fi19 ^ fi19)) {
            case -439:
                var v24 uint16 = 141
                _ = v24
            }
            if true {
                v15 += int32(fi19)
                for fi25 := 0; fi25 < 14; fi25++ {
                    var v26 uint64 = (v1 + 278)
                    _ = fi25
                    _ = v26
                }
                v17 += v17
            }
            v27 := GenS1{F0: 56}
            _ = v27
        } else {
            v28 := GenS1{F0: 42}
            v8 = v1
            var v29 uint = 42
            _ = v28
            _ = v29
        }
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
