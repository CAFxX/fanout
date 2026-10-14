package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 uint32; F2 int32 }

func main() {
    v1 := []uint32{1}
    v2 := make(map[uint32]uint32)
    var v3 float64 = float64(1.5)
    var v4 int16 = 1000
    var v5 uint8 = 127
    v6 := make(map[int]string)
    fmt.Println(v4, v5)
    if false {
        var v7 uint16 = 0
        _ = v7
    } else {
        var v8 uint8 = v5
        v4 -= -32768
        _ = v8
    }
    fmt.Println(v4, v4, v4)
    v5 += 212
    delete(v6, 976)
    v9 := GenS1{F0: "foo bar", F1: 2147483647, F2: 1}
    if (149 == v5) {
        v10 := GenS1{F0: (("" + "a") + ""), F1: 969, F2: 2}
        _ = v10
    }
    v11 := GenS1{F0: "foo bar", F1: -2147483648, F2: 2147483647}
    var v12 int16 = v4
    fmt.Println(v5)
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    fmt.Println(v12)
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 uint32 = 0
        switch fi13 {
        case -409:
            var v15 float64 = (0.0 + 0.0)
            _ = v15
        case -489:
            if false {
                var v16 int = 2147483647
                _ = v16
            } else {
                for fi17 := 0; fi17 < 16; fi17++ {
                    var v18 uint64 = 0
                    fmt.Println(v5)
                    var v19 uint16 = 65535
                    var v20 int64 = -1
                    _ = fi17
                    _ = v18
                    _ = v19
                    _ = v20
                }
            }
        }
        _ = fi13
        _ = v14
    }
    v5 = (v5 & 47)
    if (42 >= 617) {
        v5 = 255
        if ((v5 <= v5) || ((((2147483647 < 892) || true) && (-1 > 1)) && false)) {
            v3 += v3
            switch ("hello" + "foo bar") {
            case "\n\t":
                if false {
                    _ = len(v1)
                }
            case "":
                v12 = v4
            }
        } else {
            var v21 string = ""
            v22 := GenS1{F0: v21, F1: 2147483647, F2: 4294967295}
            _ = v21
            _ = v22
        }
    } else {
        v23 := GenS1{F0: "foo bar", F1: -1, F2: 2147483647}
        v24 := GenS2{F0: 42}
        fmt.Println(v3, v5, v4)
        _ = v23
        _ = v24
    }
    switch "hello" {
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = 65535
        }
    default:
        v25 := GenS3{F0: 255, F1: 226, F2: -1}
        _ = v25
    }
    v3 = 1e10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
