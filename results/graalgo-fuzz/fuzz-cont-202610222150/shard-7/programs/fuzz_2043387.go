package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{}
    v3 := make(map[uint32]bool)
    var v4 float32 = float32(1e10)
    v5 := []uint32{2}
    var v6 uint = 2
    switch 128 {
    case 2147483647:
        v6 = 65535
    default:
        v7 := GenS1{F0: -52, F1: 31, F2: 42}
        _ = v7
    }
    var v8 bool = true
    v8 = v8
    fmt.Println(v8)
    switch "\n\t" {
    case "hello":
        if true {
            var v9 uint8 = 109
            _ = v9
        }
    case "\n\t":
        v1[("" + "hello")] = v8
    case "foo bar":
        for fi10 := 0; fi10 < 5; fi10++ {
            switch "foo bar" {
            case "":
                v8 = (143 < 42)
            case "foo bar":
                for fi11 := 0; fi11 < 10; fi11++ {
                    v6 += (v6 << (((^(v6))) & 63))
                    v12 := GenS2{F0: 413, F1: 2}
                    fmt.Println(v4, fi10, fi11)
                    _ = fi11
                    _ = v12
                }
            case "hello":
                v13 := v3[42]
                _ = v13
            }
            if (false || v8) {
                var v14 bool = (v8 || (fi10 > fi10))
                _ = v14
            }
            _ = fi10
        }
    default:
        v6 += (v6 / ((v6) | 1))
    }
    var v15 string = "a"
    v6 += (v6 | v6)
    var v16 uint64 = 4294967295
    var v17 uint8 = 128
    var v18 int32 = 546
    if (-7 >= int8(v17)) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        for fi19 := 0; fi19 < 7; fi19++ {
            for fi20 := 0; fi20 < 13; fi20++ {
                if v8 {
                    v21 := GenS3{F0: 42}
                    var v22 uint = v6
                    _ = v21
                    _ = v22
                } else {
                    _ = len(v3)
                    var v23 int16 = -1
                    if len(v5) > 0 {
                        v5[0] = uint32(v17)
                    }
                    _ = v23
                }
                _ = fi20
            }
            if (128 >= v17) {
                for fi24 := 0; fi24 < 8; fi24++ {
                    v3[1] = ((((v15 + v15) + v15) + (v15 + v15)) <= v15)
                    v25 := GenS2{F0: 985, F1: 42}
                    var v26 string = v15
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                v2 = append(v2, 95)
            }
            _ = fi19
        }
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        v16 += (^(v16))
        _ = fi27
    }
    var v28 int16 = -32768
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
