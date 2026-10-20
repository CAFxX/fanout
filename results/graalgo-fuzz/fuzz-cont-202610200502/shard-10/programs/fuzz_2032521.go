package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 string }
type GenS2 struct { F0 int; F1 uint32; F2 int16 }

func main() {
    v1 := []uint32{}
    v2 := make(map[string]uint32)
    var v3 int = -299
    _ = len(v1)
    v4 := GenS1{F0: 688, F1: 83, F2: "a"}
    switch "foo bar" {
    case "a":
        v3 = v3
    case "":
        v5 := GenS1{F0: 2, F1: 127, F2: (("a" + "foo bar") + ("" + "\n\t"))}
        _ = v5
    }
    _ = len(v2)
    v3 -= v3
    if false {
        if (42 > 0) {
            fmt.Println(v3)
            for fi6 := 0; fi6 < 16; fi6++ {
                if (true || (!(255 > 26))) {
                    v7 := GenS2{F0: v3, F1: 65535, F2: 1}
                    _ = len(v2)
                    _ = v7
                }
                _ = fi6
            }
            fmt.Println(v3)
        } else {
            for fi8 := 0; fi8 < 13; fi8++ {
                v3 -= fi8
                fmt.Println(v3, fi8)
                if false {
                    if (!false) {
                        var v9 int32 = -292
                        var v10 int = (v3 * (-(fi8)))
                        _ = v10
                        _ = v9
                    }
                }
                _ = fi8
            }
        }
    }
    var v11 int8 = 100
    v11 = 100
    for fi12 := 0; fi12 < 5; fi12++ {
        _ = len(v2)
        v11 = v11
        v13 := GenS2{F0: 9223372036854775807, F1: 65535, F2: 42}
        _ = fi12
        _ = v13
    }
    var v14 int64 = -1
    v15 := GenS1{F0: 42, F1: 127, F2: ""}
    var v16 uint = 18446744073709551615
    switch "hello" {
    case "hello":
        var v17 float32 = 1.5
        _ = v17
    case "a":
        if false {
            v18 := v2[("" + ((("a" + "") + "hello") + "\n\t"))]
            var v19 int64 = -9223372036854775808
            _ = v18
            _ = v19
        }
    case "":
        var v20 uint16 = uint16(v16)
        _ = v20
    default:
        _ = len(v2)
    }
    if true {
        fmt.Println(v16, v3)
        for fi21 := 0; fi21 < 7; fi21++ {
            for fi22 := 0; fi22 < 11; fi22++ {
                var v23 float32 = (0.0 / float32(1.5))
                _ = fi22
                _ = v23
            }
            for fi24 := 0; fi24 < 5; fi24++ {
                switch (-(fi24)) {
                case 1:
                    var v25 bool = (2 < 2147483647)
                    _ = v25
                case -2147483648:
                    var v26 uint64 = 889
                    _ = v26
                }
                v1 = append(v1, 316)
                v27 := GenS1{F0: 32767, F1: 134, F2: (("a" + "foo bar") + "\n\t")}
                _ = fi24
                _ = v27
            }
            switch (v3 & v3) {
            case -9223372036854775808:
                for fi28 := 0; fi28 < 10; fi28++ {
                    var v29 uint = v16
                    _ = fi28
                    _ = v29
                }
            case 32768:
                fmt.Println(v16, v3, v14)
            default:
                for fi30 := 0; fi30 < 2; fi30++ {
                    var v31 float64 = 3.14159
                    var v32 uint16 = 65535
                    var v33 float32 = 1e10
                    _ = fi30
                    _ = v31
                    _ = v32
                    _ = v33
                }
            }
            _ = fi21
        }
        v2["hello"] = 4294967295
    } else {
        var v34 uint = 4294967295
        _ = v34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
