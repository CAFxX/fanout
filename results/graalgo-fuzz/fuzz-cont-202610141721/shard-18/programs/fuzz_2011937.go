package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 string; F2 uint64 }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 2147483647
    var v3 int16 = 32767
    v4 := GenS1{F0: 220, F1: 746}
    for fi5 := 0; fi5 < 13; fi5++ {
        if (false || true) {
            switch int(v3) {
            case 0:
                fmt.Println(v2)
            }
            if true {
                for fi6 := 0; fi6 < 9; fi6++ {
                    switch "a" {
                    case "a":
                        fmt.Println(v2)
                    }
                    v7 := []int{fi5}
                    _ = fi6
                    _ = v7
                }
                var v8 int8 = -1
                _ = v8
            } else {
                var v9 uint64 = 18446744073709551615
                _ = v9
            }
        }
        if (-1 != v3) {
            var v10 uint64 = 42
            _ = v10
        } else {
            fmt.Println(fi5)
            fmt.Println(fi5, v3)
            var v11 uint8 = 77
            _ = v11
        }
        _ = fi5
    }
    v12 := GenS1{F0: 93, F1: 0}
    for fi13 := 0; fi13 < 16; fi13++ {
        if true {
            var v14 float64 = 1e10
            _ = v14
        }
        if ("\n\t" == "") {
            if (v3 <= v3) {
                v15 := []int{fi13, 32767, (fi13 ^ fi13)}
                _ = v15
            }
            var v16 int32 = 0
            _ = v16
        } else {
            v17 := GenS2{F0: 2}
            var v18 int8 = int8(v2)
            _ = v17
            _ = v18
        }
        for fi19 := 0; fi19 < 11; fi19++ {
            fmt.Println(v2, v3, fi13)
            v3 = 1000
            _ = fi19
        }
        _ = fi13
    }
    fmt.Println(v2)
    v20 := GenS1{F0: 128, F1: 65535}
    v2 += int32(v3)
    switch 727 {
    case 0:
        var v21 uint32 = 142
        _ = v21
    default:
        var v22 string = ("hello" + "")
        _ = v22
    }
    if false {
        var v23 string = "\n\t"
        _ = v23
    }
    v2 += (v2 + -2147483648)
    v24 := []int{32767}
    if (!false) {
        v25 := GenS1{F0: 128, F1: 32767}
        _ = v25
    } else {
        var v26 uint64 = 17
        v27 := GenS3{F0: 2147483647, F1: "hello", F2: (v26 / (((v26 | 754)) | 1))}
        _ = v26
        _ = v27
    }
    var v28 int = 1
    fmt.Println(v2, v3, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v24
    _ = v28
    _ = v3
    _ = v4
}
