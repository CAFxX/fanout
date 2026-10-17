package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint; F1 uint32 }
func genFn0() int8 {
    var v6 int = 0
    var v7 int64 = int64(v6)
    v8 := []int{}
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 uint64 = 758
        _ = fi9
        _ = v10
    }
    v8 = append(v8, v6)
    _ = v6
    _ = v7
    _ = v8
    return int8(v6)
}

func main() {
    var v1 string = "hello"
    var v2 int32 = 1
    v3 := []string{(v1 + "\n\t"), (("foo bar" + "foo bar") + v1)}
    var v4 int32 = v2
    var v5 uint16 = 2
    if false {
        v11 := GenS1{F0: -32768, F1: 2147483647, F2: 288}
        v12 := GenS2{F0: 574, F1: 2147483647}
        _ = v11
        _ = v12
    } else {
        var v13 int64 = 1
        switch (v1 + v1) {
        case "foo bar":
            if false {
                v14 := GenS1{F0: 1000, F1: 1, F2: -32768}
                v3 = append(v3, "hello")
                var v15 int = 0
                _ = v14
                _ = v15
            }
        case "a":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "\n\t":
            v2 = v2
        }
        _ = v13
    }
    var v16 uint8 = 8
    var v17 uint32 = 1
    fmt.Println(v4, v5)
    fmt.Println(v5)
    fmt.Println(v16, v2, v17)
    var v18 int8 = 42
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 int32 = int32(v17)
        if ((3.14159 / 1.5) < (-2.25 + -2.25)) {
            v21 := GenS1{F0: 42, F1: v17, F2: -1}
            for fi22 := 0; fi22 < 12; fi22++ {
                if (false || false) {
                    if (0 != 18446744073709551615) {
                        var v23 uint64 = 2
                        v24 := GenS2{F0: 42, F1: (v17 << ((717) & 63))}
                        _ = v23
                        _ = v24
                    }
                    fmt.Println(v20, v20)
                }
                _ = fi22
            }
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    v25 := GenS1{F0: -1, F1: 4294967295, F2: 660}
    if (255 != 0) {
        v4 = v2
        var v26 int16 = -1
        _ = v26
    }
    if false {
        for fi27 := 0; fi27 < 4; fi27++ {
            v16 += v16
            fmt.Println(v17)
            _ = fi27
        }
    } else {
        v16 = v16
        v5 += (^(v5))
    }
    var v28 uint16 = 2
    fmt.Println(v2)
    v5 = (v5 - (v28 | v28))
    v17 += v17
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
