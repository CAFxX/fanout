package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int32) int8 {
    fmt.Println(p0)
    v7 := GenS1{F0: 255}
    for fi8 := 0; fi8 < 19; fi8++ {
        p0 = -935
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 int32 = (p0 >> ((4294967295) & 63))
            v10 = (v10 << ((1) & 63))
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    switch "a" {
    case "hello":
        var v11 int8 = 127
        _ = v11
    case "a":
        if (!(2147483648 != 2147483647)) {
            switch "a" {
            case "\n\t":
                v12 := []int{2147483647, 128, 305}
                _ = v12
            default:
                var v13 int8 = -128
                _ = v13
            }
            p0 = p0
        }
    case "":
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 float64 = 3.14159
            var v16 float32 = (float32(0.0) * 3.14159)
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    _ = p0
    _ = v7
    return 73
}

func main() {
    v1 := []string{}
    v2 := make(map[string]uint32)
    var v3 string = ((("\n\t" + "hello") + "foo bar") + "hello")
    var v4 int64 = -9223372036854775808
    var v5 uint = 874
    v6 := []uint32{611}
    if (32767 > 307) {
        if true {
            var v17 uint8 = 221
            var v18 int16 = 32767
            var v19 uint64 = 1
            _ = v17
            _ = v18
            _ = v19
        }
        fmt.Println(v3, v5)
    }
    v20 := GenS1{F0: 1}
    if len(v1) > 0 {
        v1[0] = (((((v3 + v3) + v3) + (v3 + ("\n\t" + v3))) + "hello") + ("foo bar" + v3))
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        for fi22 := 0; fi22 < 15; fi22++ {
            v5 -= v5
            switch fi21 {
            case 1:
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    fmt.Println(fi21, fi22, fi22)
                } else {
                    v23 := v2[v3]
                    _ = v23
                }
                _ = len(v2)
            }
            fmt.Println(v3, fi21, v4)
            _ = fi22
        }
        _ = fi21
    }
    fmt.Println(v5, v4, v3)
    v3 = v3
    fmt.Println(v3, v5, v5)
    v5 = 0
    v3 = ""
    v3 = v3
    if false {
        v6 = append(v6, 309)
        switch -81 {
        case -706:
            for fi24 := 0; fi24 < 5; fi24++ {
                for fi25 := 0; fi25 < 3; fi25++ {
                    v4 += (-(v4))
                    _ = fi25
                }
                _ = fi24
            }
        case 2147483647:
            for fi26 := 0; fi26 < 12; fi26++ {
                v27 := GenS1{F0: 42}
                if (-2147483648 <= 2147483647) {
                    v3 = v3
                }
                _ = fi26
                _ = v27
            }
        }
        v28 := GenS1{F0: uint8(v5)}
        _ = v28
    }
    var v29 bool = false
    for fi30 := 0; fi30 < 20; fi30++ {
        v29 = (-542 != 2147483647)
        _ = fi30
    }
    var v31 bool = (127 == 213)
    v4 = v4
    var v32 uint16 = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
