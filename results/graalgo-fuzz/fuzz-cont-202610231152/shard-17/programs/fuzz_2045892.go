package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }

func main() {
    v1 := []string{"a", ("hello" + "")}
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 32767
    var v4 uint64 = 4294967295
    v5 := make(map[uint32]string)
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 bool = (0 == 794)
        _ = fi6
        _ = v7
    }
    if true {
        switch "" {
        case "hello":
            v4 -= 0
        default:
            v3 = 65535
        }
    } else {
        fmt.Println(v4)
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        switch -57 {
        case 32768:
            for fi9 := 0; fi9 < 13; fi9++ {
                v10 := GenS1{F0: fi9, F1: 100}
                var v11 int8 = 127
                fmt.Println(v4)
                _ = fi9
                _ = v10
                _ = v11
            }
        case 128:
            if false {
                v12 := GenS1{F0: (fi8 & 32768), F1: 72}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v12
            }
        }
        for fi13 := 0; fi13 < 2; fi13++ {
            v3 = (v3 & (v3 | 65535))
            var v14 uint8 = 120
            _ = fi13
            _ = v14
        }
        var v15 float32 = 3.14159
        _ = fi8
        _ = v15
    }
    v16 := GenS1{F0: -1, F1: -48}
    v17 := GenS1{F0: -2147483648, F1: 100}
    for fi18 := 0; fi18 < 5; fi18++ {
        v19 := GenS1{F0: (^(fi18)), F1: -128}
        v2[2147483648] = 942
        _ = fi18
        _ = v19
    }
    var v20 uint = 0
    if ("hello" >= "hello") {
        switch -940 {
        case 686:
            var v21 int32 = -212
            _ = v21
        }
    }
    v22 := GenS1{F0: 367, F1: -1}
    v20 += 344
    var v23 float32 = 3.14159
    for fi24 := 0; fi24 < 19; fi24++ {
        fmt.Println(v4, v20, v20)
        var v25 int64 = -181
        v1 = append(v1, "a")
        _ = fi24
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
