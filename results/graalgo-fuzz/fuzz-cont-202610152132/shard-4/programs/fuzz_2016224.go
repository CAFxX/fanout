package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{42, 2, 2147483647, 939}
    v3 := make(map[int]uint32)
    var v4 int16 = 42
    var v5 uint64 = 432
    var v6 uint = 4294967295
    v4 -= 1000
    switch "\n\t" {
    case "hello":
        v7 := GenS1{F0: 898}
        _ = v7
    case "a":
        v6 -= (^(v6))
    }
    switch -654 {
    case -855:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 128:
        if false {
            v4 += (v4 + v4)
        } else {
            for fi8 := 0; fi8 < 6; fi8++ {
                v9 := GenS1{F0: (v5 >> ((uint(v5)) & 63))}
                _ = fi8
                _ = v9
            }
            var v10 uint32 = 579
            v11 := GenS1{F0: 4294967295}
            _ = v10
            _ = v11
        }
    case 32768:
        v2 = append(v2, 121)
    default:
        if (127 > 1) {
            v12 := GenS1{F0: 18446744073709551615}
            for fi13 := 0; fi13 < 3; fi13++ {
                switch fi13 {
                case -994:
                    v14 := v3[(^(fi13))]
                    _ = v14
                }
                var v15 float32 = -2.25
                switch (fi13 ^ fi13) {
                case 0:
                    v5 = (v5 | (v5 | 719))
                case 127:
                    v16 := GenS2{F0: 0}
                    _ = v16
                case -1:
                    switch fi13 {
                    case 32767:
                        v6 -= v6
                    case 9223372036854775807:
                        if (v4 > -32768) {
                            v17 := GenS1{F0: v5}
                            if ((fi13 + -1) <= fi13) {
                                for fi18 := 0; fi18 < 10; fi18++ {
                                    v5 = (v5 ^ v5)
                                    _ = fi18
                                }
                            } else {
                                switch ("hello" + "a") {
                                case "\n\t":
                                    switch int(v5) {
                                    case 32768:
                                        fmt.Println(v5)
                                    case 32767:
                                        delete(v1, 865)
                                    default:
                                        fmt.Println(fi13)
                                    }
                                }
                            }
                            _ = v17
                        }
                    case 2147483647:
                        _ = len(v2)
                    }
                }
                _ = fi13
                _ = v15
            }
            _ = v12
        }
    }
    var v19 uint64 = (^(v5))
    var v20 uint = (v6 >> ((4294967295) & 63))
    v20 -= v20
    var v21 int8 = 42
    delete(v1, 2147483648)
    v6 = v20
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
