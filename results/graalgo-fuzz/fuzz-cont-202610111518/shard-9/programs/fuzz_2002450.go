package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint16; F1 int16; F2 int8 }

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]bool)
    var v3 uint64 = (v1 + 18446744073709551615)
    v4 := []uint32{42, 42, 0, 664}
    var v5 string = ""
    var v6 uint32 = 957
    v1 -= (v1 / ((v3) | 1))
    v3 -= 981
    var v7 int = -9223372036854775808
    var v8 int32 = -89
    var v9 uint16 = 32767
    if true {
        v3 = (v1 + v1)
        v10 := GenS1{F0: (true && (255 != 448))}
        _ = v10
    }
    v7 += v7
    if true {
        if (v7 >= (v7 & (v7 ^ v7))) {
            v11 := GenS1{F0: true}
            _ = v11
        }
        fmt.Println(v1, v8, v8)
        v12 := GenS2{F0: (v5 + "hello"), F1: 893}
        _ = v12
    } else {
        v2[(("hello" + (v5 + v5)) + v5)] = false
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        if true {
            switch (fi13 - fi13) {
            case 128:
                for fi14 := 0; fi14 < 14; fi14++ {
                    fmt.Println(v1, v7)
                    _ = fi14
                }
            default:
                var v15 uint32 = (v6 << ((42) & 63))
                _ = v15
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi16 := 0; fi16 < 3; fi16++ {
                if true {
                    var v17 int64 = 321
                    delete(v2, "a")
                    _ = v17
                }
                switch 917 {
                case 42:
                    var v18 uint64 = v1
                    _ = v18
                case -9223372036854775808:
                    v7 = fi16
                default:
                    v4 = append(v4, (v6 | v6))
                }
                var v19 int8 = -1
                v3 = v3
                _ = fi16
                _ = v19
            }
        } else {
            if ((v6 >> ((uint(v6)) & 63)) >= (v6 | (^(v6)))) {
                var v20 int = -769
                var v21 int16 = 1000
                _ = v20
                _ = v21
            } else {
                _ = len(v2)
                v3 = 308
                for fi22 := 0; fi22 < 6; fi22++ {
                    if false {
                        v23 := GenS3{F0: (^(v9)), F1: -32768, F2: -1}
                        _ = v23
                    } else {
                        var v24 uint8 = 255
                        if len(v4) > 0 {
                            v4[0] = 65535
                        }
                        v2[("hello" + ("hello" + "hello"))] = (!true)
                        _ = v24
                    }
                    _ = fi22
                }
            }
        }
        _ = fi13
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
