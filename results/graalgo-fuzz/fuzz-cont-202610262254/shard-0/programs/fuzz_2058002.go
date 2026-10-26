package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 uint16; F1 int64 }

func main() {
    v1 := []int{433, 128}
    var v2 bool = false
    var v3 float64 = 3.14159
    var v4 int = -1
    for fi5 := 0; fi5 < 3; fi5++ {
        if len(v1) > 0 {
            v1[0] = -478
        }
        if v2 {
            fmt.Println(v2, v4, v4)
            v4 -= v4
        }
        var v6 uint64 = 18446744073709551615
        _ = fi5
        _ = v6
    }
    var v7 uint32 = 1
    v8 := GenS1{F0: 100, F1: 4294967295, F2: 1}
    v9 := GenS1{F0: 42, F1: 65535, F2: 255}
    fmt.Println(v7, v2)
    if v2 {
        if ((v2 || false) && v2) {
            if false {
                if (!false) {
                    var v10 uint = 4294967295
                    var v11 int16 = 562
                    _ = v10
                    _ = v11
                }
                var v12 int8 = 127
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v12
            } else {
                v13 := GenS2{F0: 0}
                var v14 int16 = 1
                _ = v13
                _ = v14
            }
        } else {
            v15 := GenS3{F0: ((true || (!(171 == 219))) && v2), F1: (("\n\t" + ("foo bar" + "hello")) + "\n\t")}
            switch "a" {
            case "hello":
                for fi16 := 0; fi16 < 6; fi16++ {
                    if len(v1) > 0 {
                        v1[0] = fi16
                    }
                    v2 = (1000 < -499)
                    for fi17 := 0; fi17 < 10; fi17++ {
                        var v18 int8 = 1
                        _ = fi17
                        _ = v18
                    }
                    _ = fi16
                }
            case "":
                for fi19 := 0; fi19 < 15; fi19++ {
                    v7 = v7
                    if ((float64(((1e10 / 1.5) - 0.0)) - v3) <= (float64((3.14159 + 0.0)) - v3)) {
                        var v20 int = (^(fi19))
                        v3 = float64((3.14159 * (((0.1 / 1.5) * 1e10) / 1.0)))
                        fmt.Println(v4)
                        _ = v20
                    } else {
                        switch (v4 & 0) {
                        case 9223372036854775807:
                            for fi21 := 0; fi21 < 2; fi21++ {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v2 = false
                                if (!v2) {
                                    v2 = v2
                                }
                                _ = fi21
                            }
                        }
                        var v22 int = (v4 << ((0) & 63))
                        _ = v22
                    }
                    if true {
                        var v23 bool = true
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v23
                    }
                    _ = fi19
                }
            }
            _ = v15
        }
        switch "\n\t" {
        case "":
            if len(v1) > 0 {
                v1[0] = (v4 ^ (v4 ^ 57))
            }
        case "foo bar":
            fmt.Println(v3)
        default:
            for fi24 := 0; fi24 < 16; fi24++ {
                var v25 bool = false
                fmt.Println(v3, v2)
                _ = fi24
                _ = v25
            }
        }
        var v26 uint64 = 18446744073709551615
        _ = v26
    } else {
        var v27 uint32 = 42
        _ = v27
    }
    var v28 int64 = -880
    var v29 int8 = 100
    v28 += 0
    v3 += v3
    v1 = append(v1, (v4 % (((v4 + v4)) | 1)))
    v28 = (v28 + 103)
    var v30 float32 = (1.5 / 1.5)
    v31 := GenS4{F0: 255, F1: v28}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
