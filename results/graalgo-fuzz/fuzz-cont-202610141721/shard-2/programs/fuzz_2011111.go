package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int16; F1 uint64; F2 string }
type GenS4 struct { F0 int8; F1 uint }

func main() {
    var v1 uint16 = 42
    var v2 bool = false
    var v3 uint16 = v1
    var v4 uint32 = 42
    var v5 uint16 = (v3 >> ((18446744073709551615) & 63))
    var v6 uint = 65535
    v7 := []int{32768, 127, 2147483647}
    v8 := GenS1{F0: 0}
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := GenS1{F0: 127}
            _ = fi10
            _ = v11
        }
        if (v2 && v2) {
            if (true || v2) {
                for fi12 := 0; fi12 < 16; fi12++ {
                    if len(v7) > 0 {
                        v7[0] = fi12
                    }
                    var v13 int = 128
                    v7 = append(v7, (-(v13)))
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    switch "hello" {
                    case "foo bar":
                        if len(v7) > 0 {
                            v7[0] = -2147483648
                        }
                    }
                    _ = fi12
                    _ = v13
                }
            }
        }
        _ = fi9
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        v3 -= (v1 % ((v3) | 1))
        var v15 uint32 = (^(v4))
        _ = fi14
        _ = v15
    }
    if (v2 || false) {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        if (!v2) {
            for fi16 := 0; fi16 < 3; fi16++ {
                v17 := GenS1{F0: -128}
                var v18 int32 = 2147483647
                v19 := GenS2{F0: 1, F1: v1}
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
        } else {
            if false {
                v20 := GenS2{F0: 663, F1: 2}
                var v21 uint8 = 127
                _ = v20
                _ = v21
            }
            var v22 uint64 = 64
            _ = v22
        }
    }
    var v23 int64 = 844
    v6 -= 2
    if v2 {
        v6 = v6
        var v24 int = 4294967295
        if (true || v2) {
            var v25 int8 = 42
            var v26 int32 = 42
            var v27 uint32 = 561
            _ = v25
            _ = v26
            _ = v27
        }
        _ = v24
    } else {
        v28 := GenS3{F0: -740, F1: 1, F2: "a"}
        _ = v28
    }
    var v29 int64 = (v23 & 42)
    var v30 int8 = 100
    v3 += v5
    v31 := GenS1{F0: 1}
    v6 += (v6 - v6)
    for fi32 := 0; fi32 < 1; fi32++ {
        v2 = v2
        for fi33 := 0; fi33 < 17; fi33++ {
            fmt.Println(v6, v1)
            v7 = append(v7, fi32)
            _ = fi33
        }
        v34 := GenS3{F0: 464, F1: 0, F2: "foo bar"}
        _ = fi32
        _ = v34
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        var v36 uint32 = (v4 | 1)
        v37 := GenS4{F0: v30, F1: v6}
        _ = fi35
        _ = v36
        _ = v37
    }
    if (!v2) {
        var v38 float32 = 3.14159
        _ = v38
    } else {
        var v39 float64 = 0.1
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        var v40 int16 = 1000
        _ = v39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v23
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
