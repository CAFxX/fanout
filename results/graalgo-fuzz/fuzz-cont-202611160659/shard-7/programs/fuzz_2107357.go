package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 int16 }
func genFn0(p0 uint8) uint {
    var v7 int8 = 54
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := []int{(^(fi8))}
        var v10 int64 = -757
        _ = fi8
        _ = v10
        _ = v9
    }
    fmt.Println(v7, p0)
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := GenS1{F0: (^(p0))}
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = v7
    return 1
}

func main() {
    var v1 bool = (2147483647 < 128)
    var v2 uint16 = 1
    var v3 uint = 934
    var v4 bool = v1
    var v5 uint = v3
    v6 := []uint32{2147483647, 837}
    if len(v6) > 0 {
        v6[0] = 819
    }
    if v1 {
        if len(v6) > 0 {
            v6[0] = 2
        }
        var v13 int16 = 32767
        _ = v13
    }
    switch -330 {
    case 9223372036854775807:
        v14 := GenS1{F0: 128}
        _ = v14
    case -88:
        v15 := GenS1{F0: 167}
        _ = v15
    case 887:
        var v16 uint16 = v2
        _ = v16
    }
    _ = len(v6)
    v17 := GenS1{F0: 42}
    fmt.Println(v4, v3, v4)
    var v18 uint8 = 42
    v2 = (^(v2))
    if len(v6) > 0 {
        v6[0] = 635
    }
    var v19 float32 = (1e10 * float32(float64(3.14159)))
    v20 := GenS2{F0: ((("foo bar" + "") + "a") + "foo bar")}
    v21 := GenS3{F0: (104 == 42)}
    v22 := GenS4{F0: (v4 && ((v18 >> ((v3) & 63)) > (^(v18)))), F1: 458}
    for fi23 := 0; fi23 < 19; fi23++ {
        fmt.Println(v5, v1, v18)
        _ = fi23
    }
    v19 = (float32(3.14159) - v19)
    for fi24 := 0; fi24 < 16; fi24++ {
        v6 = append(v6, 65535)
        for fi25 := 0; fi25 < 4; fi25++ {
            v6 = append(v6, uint32(v5))
            if v4 {
                fmt.Println(v18, v3)
                var v26 uint8 = 0
                switch (fi25 - fi25) {
                case 1:
                    fmt.Println(v2)
                }
                if false {
                    v5 += 4294967295
                } else {
                    v6 = append(v6, 1)
                    var v27 uint32 = 0
                    v19 = (-2.25 - v19)
                    _ = v27
                }
                var v28 float32 = v19
                _ = v26
                _ = v28
            }
            v6 = append(v6, 2)
            _ = fi25
        }
        _ = fi24
    }
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
