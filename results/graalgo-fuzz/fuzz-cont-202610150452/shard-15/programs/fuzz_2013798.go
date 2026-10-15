package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[string]string)
    var v3 uint32 = 2147483647
    v4 := []int{-2147483648, 128}
    var v5 bool = false
    var v6 int32 = -495
    var v7 float64 = float64(float32(0.0))
    v8 := GenS1{F0: 255}
    switch -280 {
    case -409:
        v3 += 0
    }
    v9 := GenS1{F0: uint8(v3)}
    v1 -= v1
    v10 := GenS1{F0: 128}
    v7 -= v7
    var v11 string = (("hello" + "hello") + "\n\t")
    var v12 bool = v5
    v12 = v5
    switch "\n\t" {
    case "":
        for fi13 := 0; fi13 < 4; fi13++ {
            v14 := GenS2{F0: (v6 ^ (v6 - -1))}
            for fi15 := 0; fi15 < 1; fi15++ {
                v16 := GenS3{F0: 92}
                fmt.Println(v6)
                if v12 {
                    var v17 float32 = 0.0
                    _ = v17
                } else {
                    _ = len(v4)
                }
                _ = fi15
                _ = v16
            }
            for fi18 := 0; fi18 < 5; fi18++ {
                switch (^(fi13)) {
                case -2147483648:
                    _ = len(v2)
                case 844:
                    if ((true || false) && true) {
                        var v19 int16 = 1
                        var v20 uint64 = 2
                        fmt.Println(v3, v5, v7)
                        v21 := GenS2{F0: int32(v19)}
                        var v22 int16 = (v19 * 696)
                        switch (fi18 % (((fi13 * 9223372036854775807)) | 1)) {
                        case 128:
                            fmt.Println(v22, v6, v12)
                        case -9223372036854775808:
                            v7 += v7
                        }
                        _ = v19
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                case 558:
                    var v23 uint64 = 0
                    _ = v23
                }
                _ = fi18
            }
            _ = fi13
            _ = v14
        }
    case "a":
        var v24 uint = 42
        _ = v24
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        for fi26 := 0; fi26 < 15; fi26++ {
            fmt.Println(v12, v7)
            var v27 bool = (0 < 255)
            v6 = (^(v6))
            _ = fi26
            _ = v27
        }
        v3 += (v3 & 2)
        _ = fi25
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        v29 := GenS1{F0: 2}
        for fi30 := 0; fi30 < 18; fi30++ {
            v31 := GenS4{F0: 253, F1: 65535}
            _ = fi30
            _ = v31
        }
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
