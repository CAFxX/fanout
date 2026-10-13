package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{443}
    var v3 uint16 = 42
    v4 := make(map[int]uint32)
    var v5 uint8 = 255
    switch -9223372036854775808 {
    case 32767:
        v6 := GenS1{F0: v3}
        _ = v6
    case 127:
        v3 = 0
    }
    var v7 bool = false
    var v8 int64 = 9223372036854775807
    var v9 int8 = -23
    _ = len(v1)
    switch "" {
    case "":
        fmt.Println(v8)
    case "foo bar":
        var v10 int16 = 32767
        _ = v10
    case "\n\t":
        for fi11 := 0; fi11 < 6; fi11++ {
            if ((fi11 << ((255) & 63)) >= (fi11 & fi11)) {
                for fi12 := 0; fi12 < 4; fi12++ {
                    switch fi12 {
                    case 42:
                        var v13 bool = true
                        _ = v13
                    case 314:
                        var v14 float32 = 0.0
                        _ = v14
                    case 4294967295:
                        delete(v4, (fi12 * int(v5)))
                    default:
                        if (!true) {
                            var v15 uint64 = 18446744073709551615
                            v5 = uint8(v15)
                            _ = v15
                        } else {
                            var v16 bool = false
                            _ = v16
                        }
                    }
                    var v17 uint16 = (v3 % (((^(v3))) | 1))
                    _ = fi12
                    _ = v17
                }
                v8 -= (v8 ^ v8)
            }
            if (!((v8 << ((18446744073709551615) & 63)) != 640)) {
                var v18 int64 = (v8 + v8)
                _ = v18
            }
            if false {
                v19 := GenS1{F0: 2}
                fmt.Println(fi11, v8, v9)
                _ = v19
            } else {
                var v20 int = 4294967295
                _ = v20
            }
            _ = fi11
        }
    }
    if false {
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 int16 = int16(fi21)
            var v23 uint32 = 756
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    switch 32768 {
    case 4294967295:
        var v24 string = "\n\t"
        _ = v24
    case 127:
        v5 = v5
    default:
        var v25 string = "\n\t"
        _ = v25
    }
    v9 = (v9 % (((^(v9))) | 1))
    for fi26 := 0; fi26 < 1; fi26++ {
        var v27 uint = 18446744073709551615
        _ = fi26
        _ = v27
    }
    v28 := GenS1{F0: (^(v3))}
    var v29 uint32 = 908
    for fi30 := 0; fi30 < 1; fi30++ {
        for fi31 := 0; fi31 < 12; fi31++ {
            var v32 uint = 757
            v5 = (v5 & 220)
            _ = fi31
            _ = v32
        }
        v33 := GenS2{F0: 0, F1: -1, F2: (v29 + (v29 << ((600) & 63)))}
        v34 := v1[v29]
        _ = fi30
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 19; fi35++ {
        for fi36 := 0; fi36 < 17; fi36++ {
            v37 := GenS3{F0: (v29 | 1), F1: -2147483648, F2: (^(v29))}
            _ = fi36
            _ = v37
        }
        _ = fi35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
