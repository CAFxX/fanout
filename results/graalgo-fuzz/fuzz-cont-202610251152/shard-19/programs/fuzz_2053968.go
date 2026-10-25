package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int }
func genFn0(p0 uint32) uint8 {
    var v7 uint16 = 42
    for fi8 := 0; fi8 < 11; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            p0 = (p0 / (((p0 ^ p0)) | 1))
            v7 = (v7 ^ v7)
            _ = fi9
        }
        _ = fi8
    }
    if (!false) {
        v10 := GenS1{F0: 32768}
        _ = v10
    } else {
        p0 -= (p0 >> ((18446744073709551615) & 63))
        v7 = (^(v7))
    }
    var v11 uint16 = 32767
    v7 -= (v7 >> ((0) & 63))
    _ = p0
    _ = v11
    _ = v7
    return 42
}

func main() {
    v1 := []uint32{0}
    v2 := []uint32{2147483647, 391}
    var v3 int8 = 42
    v4 := []int{-9223372036854775808, -708}
    v5 := []string{"hello", ("" + "\n\t"), "hello"}
    v6 := []string{"a"}
    for fi12 := 0; fi12 < 20; fi12++ {
        if ((1 >= 42) && true) {
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 int32 = -2147483648
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
                _ = fi13
                _ = v14
            }
            switch -1 {
            case 202:
                var v15 uint64 = 4294967295
                _ = v15
            }
            v3 -= v3
        }
        v3 += (v3 + (v3 % ((v3) | 1)))
        _ = fi12
    }
    v3 += (v3 | v3)
    var v16 int64 = 0
    v16 = int64(v3)
    var v17 uint8 = 1
    v16 = v16
    var v18 uint64 = 1
    v19 := GenS1{F0: 2147483647}
    v20 := GenS1{F0: 32767}
    _ = len(v1)
    v18 = v18
    switch 42 {
    case 128:
        v18 -= v18
    case -201:
        var v21 int64 = (v16 << ((573) & 63))
        _ = v21
    case 0:
        if (1 != 0) {
            v5 = append(v5, "hello")
            if (false || true) {
                if (191 > 1) {
                    v22 := GenS1{F0: 1}
                    _ = v22
                }
                var v23 int16 = 0
                _ = v23
            } else {
                var v24 bool = (2 < 97)
                if v24 {
                    v24 = v24
                }
                v25 := GenS2{F0: 42, F1: 127}
                _ = v24
                _ = v25
            }
            if ((v17 >> ((18446744073709551615) & 63)) == (^(v17))) {
                fmt.Println(v16)
                _ = len(v2)
                var v26 int = 9223372036854775807
                _ = v26
            } else {
                for fi27 := 0; fi27 < 14; fi27++ {
                    var v28 string = "hello"
                    v28 = ""
                    if false {
                        v28 = v28
                        v29 := GenS2{F0: 0, F1: (fi27 << ((4294967295) & 63))}
                        fmt.Println(v17, fi27, v3)
                        _ = v29
                    } else {
                        v16 = (v16 - 128)
                        var v30 uint32 = uint32(v17)
                        _ = v30
                    }
                    _ = fi27
                    _ = v28
                }
                v31 := GenS1{F0: 127}
                if false {
                    switch "\n\t" {
                    case "a":
                        var v32 bool = false
                        _ = v32
                    case "":
                        if len(v2) > 0 {
                            v2[0] = 1
                        }
                    default:
                        var v33 int32 = 1
                        _ = v33
                    }
                    v16 -= v16
                }
                _ = v31
            }
        }
    default:
        var v34 uint64 = (v18 / ((v18) | 1))
        _ = v34
    }
    v35 := GenS1{F0: 128}
    var v36 int32 = -1
    var v37 uint64 = v18
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
