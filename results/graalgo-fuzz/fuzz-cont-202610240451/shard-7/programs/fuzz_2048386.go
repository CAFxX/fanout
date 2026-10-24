package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int8 }

func main() {
    v1 := []string{"", "\n\t", ("hello" + "\n\t")}
    v2 := []uint32{2}
    v3 := []int{-198}
    var v4 bool = true
    v5 := make(map[string]bool)
    v6 := []int{128, 32768, 188}
    v7 := GenS1{F0: 349}
    switch 42 {
    case 32767:
        for fi8 := 0; fi8 < 5; fi8++ {
            _ = len(v2)
            v4 = v4
            var v9 float64 = 3.14159
            _ = fi8
            _ = v9
        }
    case 127:
        var v10 uint64 = 706
        _ = v10
    default:
        v11 := GenS1{F0: 65535}
        _ = v11
    }
    if (("a" < ("foo bar" + "hello")) || (!v4)) {
        var v12 int64 = 42
        _ = v12
    } else {
        switch "" {
        case "a":
            for fi13 := 0; fi13 < 19; fi13++ {
                for fi14 := 0; fi14 < 10; fi14++ {
                    if v4 {
                        delete(v5, "\n\t")
                        var v15 uint64 = 0
                        _ = v15
                    } else {
                        v16 := GenS2{F0: -1, F1: 32767}
                        fmt.Println(v4, fi14)
                        var v17 bool = v4
                        _ = v16
                        _ = v17
                    }
                    var v18 int32 = 2147483647
                    _ = fi14
                    _ = v18
                }
                _ = fi13
            }
        }
        v4 = false
        fmt.Println(v4)
    }
    v4 = false
    v19 := GenS3{F0: 187, F1: 100, F2: 127}
    fmt.Println(v4)
    if false {
        var v20 int32 = 1
        v21 := v5["hello"]
        _ = v20
        _ = v21
    } else {
        if true {
            var v22 string = "foo bar"
            _ = v22
        }
        if len(v2) > 0 {
            v2[0] = 65535
        }
        v23 := GenS2{F0: 833, F1: 1000}
        _ = v23
    }
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        var v25 float32 = float32(float64(1.5))
        _ = fi24
        _ = v25
    }
    if false {
        for fi26 := 0; fi26 < 1; fi26++ {
            fmt.Println(v4)
            v27 := GenS3{F0: 42, F1: 1, F2: 127}
            _ = fi26
            _ = v27
        }
        v4 = true
        v5["hello"] = (-1 >= 32768)
    } else {
        for fi28 := 0; fi28 < 9; fi28++ {
            var v29 int64 = 219
            v29 += -9223372036854775808
            _ = fi28
            _ = v29
        }
    }
    v6 = append(v6, 9223372036854775807)
    v4 = v4
    for fi30 := 0; fi30 < 1; fi30++ {
        v4 = ((((fi30 >> ((0) & 63)) != (fi30 ^ fi30)) || v4) || (v4 || (410 < 0)))
        v4 = (-357 != 0)
        _ = fi30
    }
    var v31 int32 = 2147483647
    for fi32 := 0; fi32 < 12; fi32++ {
        if false {
            if v4 {
                v4 = (-956 >= 1000)
                switch fi32 {
                case -995:
                    switch (fi32 >> ((1) & 63)) {
                    case 2147483647:
                        var v33 uint64 = 1
                        _ = v33
                    case 42:
                        v31 = v31
                    case 127:
                        var v34 uint16 = 65535
                        _ = v34
                    default:
                        var v35 bool = ((3.14159 >= -2.25) || false)
                        _ = v35
                    }
                case 32767:
                    v31 -= 42
                case 733:
                    var v36 uint = 154
                    _ = v36
                default:
                    for fi37 := 0; fi37 < 15; fi37++ {
                        switch fi32 {
                        case 4294967295:
                            v38 := GenS1{F0: 589}
                            _ = v38
                        }
                        _ = fi37
                    }
                }
                var v39 uint8 = 2
                _ = v39
            }
            if v4 {
                v40 := GenS4{F0: 42}
                v41 := GenS3{F0: 42, F1: 255, F2: 128}
                v42 := GenS3{F0: 2, F1: 42, F2: 0}
                _ = v40
                _ = v41
                _ = v42
            }
        }
        _ = fi32
    }
    v43 := GenS1{F0: 42}
    var v44 int32 = v31
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v7
}
