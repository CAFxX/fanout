package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint; F1 uint64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{428, 2147483648, 966, 2147483647}
    var v3 int8 = -1
    var v4 float64 = 0.0
    var v5 int16 = 647
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 float32 = 0.0
        _ = fi6
        _ = v7
    }
    var v8 int = 4294967295
    _ = len(v1)
    v8 = (v8 - (-(v8)))
    v3 = -128
    switch (v8 | v8) {
    case 32767:
        var v9 uint8 = 128
        _ = v9
    case -9223372036854775808:
        for fi10 := 0; fi10 < 11; fi10++ {
            v1[v8] = 2147483647
            var v11 uint16 = 1
            var v12 uint = 18446744073709551615
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    fmt.Println(v5, v3)
    v4 -= (float64(float32(1e10)) * (0.1 + ((3.14159 / v4) * 1e10)))
    fmt.Println(v4, v8, v3)
    if true {
        for fi13 := 0; fi13 < 2; fi13++ {
            for fi14 := 0; fi14 < 19; fi14++ {
                v4 += 1e10
                if len(v2) > 0 {
                    v2[0] = 38
                }
                if true {
                    if false {
                        v5 += v5
                    } else {
                        v15 := GenS1{F0: 807}
                        fmt.Println(v8, v4, v4)
                        if false {
                            v16 := GenS2{F0: 816}
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v16
                        }
                        _ = v15
                    }
                    fmt.Println(fi13, fi13, v3)
                    fmt.Println(v5)
                } else {
                    v17 := GenS1{F0: (fi13 / (((fi14 + -605)) | 1))}
                    var v18 uint64 = 149
                    _ = v17
                    _ = v18
                }
                _ = fi14
            }
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
            _ = fi13
        }
        var v19 int64 = 0
        v20 := GenS3{F0: 362, F1: (42 != -650)}
        _ = v19
        _ = v20
    } else {
        if (int16(v8) != (v5 << ((345) & 63))) {
            for fi21 := 0; fi21 < 11; fi21++ {
                v22 := GenS2{F0: 2147483647}
                _ = fi21
                _ = v22
            }
            switch (v8 + v8) {
            case 9223372036854775807:
                for fi23 := 0; fi23 < 7; fi23++ {
                    switch v8 {
                    case 87:
                        for fi24 := 0; fi24 < 3; fi24++ {
                            v25 := GenS4{F0: 1, F1: 0}
                            _ = fi24
                            _ = v25
                        }
                    }
                    _ = fi23
                }
            }
        }
        fmt.Println(v8, v8)
        switch ("" + "hello") {
        case "foo bar":
            v3 = v3
        case "":
            v26 := GenS2{F0: 0}
            _ = v26
        default:
            for fi27 := 0; fi27 < 19; fi27++ {
                if true {
                    var v28 int32 = 42
                    var v29 float64 = 0.1
                    _ = v28
                    _ = v29
                }
                v5 = 42
                _ = fi27
            }
        }
    }
    var v30 uint = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v8
}
