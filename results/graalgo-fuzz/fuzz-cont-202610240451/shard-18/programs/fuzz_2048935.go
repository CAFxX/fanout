package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS5 struct { F0 uint8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 393
    var v3 int = 42
    v4 := make(map[string]bool)
    var v5 int32 = 2147483647
    var v6 int = 42
    fmt.Println(v3, v2)
    fmt.Println(v5, v3)
    v6 += -9223372036854775808
    var v7 int16 = 42
    var v8 uint8 = 19
    for fi9 := 0; fi9 < 2; fi9++ {
        v8 = 42
        _ = fi9
    }
    if false {
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 string = ("\n\t" + ((("a" + "foo bar") + "\n\t") + "\n\t"))
            v12 := GenS1{F0: -1}
            for fi13 := 0; fi13 < 13; fi13++ {
                if true {
                    switch (fi10 ^ int(v2)) {
                    case -774:
                        v14 := []int{v3, v3, v6, v3}
                        _ = v14
                    case 102:
                        v15 := GenS2{F0: (v8 / ((62) | 1))}
                        v16 := GenS3{F0: 1, F1: int(v7), F2: 2147483647}
                        _ = v15
                        _ = v16
                    }
                    switch v11 {
                    case "foo bar":
                        v17 := []int{(-(fi10)), (^(fi13)), v3, int(v8)}
                        v5 = (v5 + (v5 | (v5 ^ 2147483647)))
                        var v18 uint32 = 2147483647
                        _ = v17
                        _ = v18
                    case "\n\t":
                        var v19 float32 = ((0.0 / 1.0) + (-2.25 * (1.5 - 1e10)))
                        _ = v19
                    default:
                        if (-128 < 39) {
                            v20 := []int{fi10}
                            var v21 string = v11
                            v3 += (fi13 % (((fi13 % ((-503) | 1))) | 1))
                            _ = v20
                            _ = v21
                        } else {
                            v22 := []int{(fi10 ^ -495), (fi10 % ((32768) | 1)), 908, fi10}
                            v23 := GenS4{F0: 0, F1: 2, F2: (v8 ^ v8)}
                            v2 -= v2
                            _ = v22
                            _ = v23
                        }
                    }
                }
                _ = fi13
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        for fi24 := 0; fi24 < 5; fi24++ {
            v7 -= v7
            _ = fi24
        }
        v8 = (v8 << ((1) & 63))
    } else {
        switch (v6 % ((v3) | 1)) {
        case 2147483647:
            v25 := GenS5{F0: 139}
            _ = v25
        case 127:
            for fi26 := 0; fi26 < 7; fi26++ {
                v2 = v2
                delete(v4, ("a" + "\n\t"))
                _ = fi26
            }
        case -9223372036854775808:
            v27 := GenS1{F0: 0}
            _ = v27
        }
    }
    v5 = 3
    fmt.Println(v6)
    v5 = int32(v8)
    v28 := GenS3{F0: 2, F1: (v3 << ((65535) & 63)), F2: 4294967295}
    v3 = (v6 % (((v3 + v6)) | 1))
    var v29 string = (("\n\t" + "\n\t") + "foo bar")
    var v30 uint64 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
