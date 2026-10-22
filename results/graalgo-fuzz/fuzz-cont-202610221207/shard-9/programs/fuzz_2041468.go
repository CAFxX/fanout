package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint }

func main() {
    v1 := []uint32{277, 65535, 4294967295, 772}
    v2 := make(map[string]uint32)
    var v3 uint16 = 674
    for fi4 := 0; fi4 < 11; fi4++ {
        for fi5 := 0; fi5 < 19; fi5++ {
            fmt.Println(fi5, fi5)
            var v6 int = 128
            switch fi4 {
            case -9223372036854775808:
                var v7 int = fi4
                _ = v7
            case 1:
                v3 = v3
            default:
                if (595 > 472) {
                    var v8 int64 = 42
                    _ = v8
                } else {
                    v3 = (^(v3))
                    var v9 int64 = -1
                    var v10 uint32 = 310
                    _ = v10
                    _ = v9
                }
            }
            _ = fi5
            _ = v6
        }
        v3 -= v3
        v3 = (^(v3))
        _ = fi4
    }
    switch 9223372036854775807 {
    case 827:
        v1 = append(v1, 159)
    }
    var v11 uint64 = 831
    fmt.Println(v3)
    if true {
        var v12 uint32 = 42
        switch -909 {
        case -747:
            for fi13 := 0; fi13 < 9; fi13++ {
                v14 := GenS1{F0: 0, F1: -782}
                _ = fi13
                _ = v14
            }
        case 0:
            v3 -= 1
        default:
            switch -1 {
            case -70:
                v15 := GenS2{F0: 4294967295}
                _ = v15
            case 1:
                if false {
                    switch ("a" + "\n\t") {
                    case "":
                        for fi16 := 0; fi16 < 8; fi16++ {
                            if false {
                                v17 := GenS1{F0: 2147483647, F1: 2147483647}
                                v18 := v2["\n\t"]
                                _ = v17
                                _ = v18
                            }
                            fmt.Println(v12, v3)
                            _ = fi16
                        }
                    case "a":
                        v3 = (^(v3))
                    }
                } else {
                    fmt.Println(v12, v11)
                    v19 := GenS1{F0: -1, F1: 749}
                    switch "\n\t" {
                    case "hello":
                        for fi20 := 0; fi20 < 2; fi20++ {
                            v12 = v12
                            _ = fi20
                        }
                    }
                    _ = v19
                }
            default:
                if (true && false) {
                    fmt.Println(v3, v12, v12)
                    if len(v1) > 0 {
                        v1[0] = v12
                    }
                }
            }
        }
        switch int(v11) {
        case 9223372036854775807:
            var v21 uint32 = v12
            _ = v21
        case 444:
            for fi22 := 0; fi22 < 14; fi22++ {
                if false {
                    var v23 float32 = ((3.14159 * (float32(3.14159) + float32(((1e10 * (0.1 * -2.25)) * (1e10 / ((1e10 / 1.5) / 1.0)))))) / 1.5)
                    _ = v23
                }
                _ = fi22
            }
        }
        _ = v12
    }
    var v24 uint8 = 0
    v24 -= 255
    switch 0 {
    case 4294967295:
        var v25 int32 = 727
        _ = v25
    case 42:
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = len(v1)
            var v26 int8 = -128
            _ = v26
        } else {
            v27 := GenS2{F0: 255}
            _ = v27
        }
    case 128:
        v24 = 128
    }
    v11 = v11
    v2[""] = 1
    v28 := v2[""]
    var v29 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
}
