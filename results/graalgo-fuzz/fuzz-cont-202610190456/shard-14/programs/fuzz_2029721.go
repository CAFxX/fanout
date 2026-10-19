package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int32; F1 uint8; F2 int8 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[string]bool)
    var v3 uint = 0
    v4 := []int{42}
    var v5 uint64 = 4294967295
    var v6 int8 = -1
    var v7 int = 1
    var v8 int = (v7 >> ((uint(v5)) & 63))
    if (!("\n\t" > "hello")) {
        for fi9 := 0; fi9 < 7; fi9++ {
            if (1 != 48) {
                v8 = int(v3)
            } else {
                v7 += fi9
                if len(v4) > 0 {
                    v4[0] = v8
                }
                for fi10 := 0; fi10 < 18; fi10++ {
                    v1 -= (v1 | (v1 - -9223372036854775808))
                    var v11 float64 = (3.14159 / 1e10)
                    _ = fi10
                    _ = v11
                }
            }
            v12 := GenS1{F0: 209}
            switch v8 {
            case 669:
                v13 := GenS2{F0: 212, F1: 1}
                _ = v13
            }
            _ = fi9
            _ = v12
        }
        v14 := GenS1{F0: 0}
        if true {
            v15 := GenS3{F0: 1000, F1: 100}
            _ = v15
        } else {
            v16 := GenS3{F0: -384, F1: v6}
            _ = v16
        }
        _ = v14
    } else {
        for fi17 := 0; fi17 < 14; fi17++ {
            if true {
                var v18 int16 = 42
                var v19 int64 = (v1 | (v1 * v1))
                for fi20 := 0; fi20 < 4; fi20++ {
                    if (true || ("\n\t" > "\n\t")) {
                        switch 9223372036854775807 {
                        case 42:
                            _ = len(v4)
                            switch 4294967295 {
                            case -205:
                                v18 = (v18 << ((v3) & 63))
                            }
                        case -9223372036854775808:
                            var v21 uint32 = 2147483647
                            _ = v21
                        }
                        v19 -= v19
                        fmt.Println(v7, v18)
                    }
                    _ = fi20
                }
                _ = v18
                _ = v19
            } else {
                var v22 uint64 = (v5 << ((v3) & 63))
                _ = v22
            }
            v4 = append(v4, (fi17 << ((v3) & 63)))
            _ = fi17
        }
        var v23 int8 = (v6 % (((v6 % ((43) | 1))) | 1))
        _ = v23
    }
    v24 := GenS3{F0: 1000, F1: (^(v6))}
    var v25 bool = true
    v26 := GenS4{F0: 1, F1: 128, F2: v6}
    var v27 int8 = v6
    v2["hello"] = false
    var v28 int64 = -696
    v29 := GenS1{F0: 127}
    var v30 uint32 = 65535
    var v31 int8 = (v27 + v6)
    _ = len(v4)
    if (255 >= 864) {
        v8 = (^(v8))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
