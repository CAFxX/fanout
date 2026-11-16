package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 int; F1 uint16 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 uint16; F1 uint; F2 int16 }

func main() {
    var v1 bool = true
    var v2 int32 = 1
    v3 := make(map[string]uint32)
    fmt.Println(v1)
    v1 = (0 < 1)
    for fi4 := 0; fi4 < 19; fi4++ {
        v2 -= (v2 % ((-2147483648) | 1))
        var v5 uint64 = 1
        _ = fi4
        _ = v5
    }
    var v6 uint16 = 32767
    v7 := GenS1{F0: 255, F1: (v6 << ((42) & 63)), F2: -1}
    var v8 uint = 65535
    v9 := []int{2147483647, 110}
    switch (("hello" + "foo bar") + (("hello" + "\n\t") + "a")) {
    case "foo bar":
        v10 := GenS2{F0: v8, F1: v6}
        _ = v10
    case "\n\t":
        if (-1 < v2) {
            v11 := GenS3{F0: uint8(v6), F1: -128}
            switch 128 {
            case 0:
                _ = len(v9)
            case 32768:
                _ = len(v3)
            default:
                var v12 uint16 = (v6 % ((v6) | 1))
                _ = v12
            }
            v8 = (v8 << ((275) & 63))
            _ = v11
        }
    default:
        v6 -= v6
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 bool = true
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        switch ("foo bar" + "") {
        case "foo bar":
            v9 = append(v9, -2147483648)
        case "\n\t":
            for fi15 := 0; fi15 < 10; fi15++ {
                if false {
                    fmt.Println(v14, v8, v6)
                    v16 := GenS3{F0: 1, F1: 44}
                    if false {
                        var v17 bool = (!(-1 < -80))
                        v18 := GenS2{F0: (v8 + 88), F1: 32767}
                        if (v14 || ((uint64(v8) <= 97) || v1)) {
                            v14 = true
                            v19 := GenS4{F0: 518, F1: -9223372036854775808}
                            var v20 string = ("hello" + "foo bar")
                            _ = v19
                            _ = v20
                        } else {
                            v14 = true
                            v14 = v17
                            fmt.Println(v17, fi15)
                        }
                        _ = v17
                        _ = v18
                    } else {
                        if true {
                            v21 := GenS5{F0: (fi13 + fi13), F1: (v6 - v6)}
                            _ = v21
                        }
                    }
                    v2 -= v2
                    v1 = (false || (!v1))
                    _ = v16
                } else {
                    v22 := GenS6{F0: 42, F1: 726}
                    _ = v22
                }
                v14 = v14
                _ = fi15
            }
        default:
            v23 := GenS7{F0: (v6 + 886), F1: (v8 - (^(v8))), F2: 32767}
            _ = v23
        }
        _ = fi13
        _ = v14
    }
    v24 := GenS7{F0: v6, F1: 42, F2: 0}
    for fi25 := 0; fi25 < 8; fi25++ {
        fmt.Println(fi25)
        for fi26 := 0; fi26 < 13; fi26++ {
            var v27 float32 = ((float32(1e10) / 1e10) + 1.5)
            _ = fi26
            _ = v27
        }
        v6 += 1
        _ = fi25
    }
    for fi28 := 0; fi28 < 8; fi28++ {
        if (!v1) {
            v3["\n\t"] = uint32(v8)
        }
        _ = fi28
    }
    v2 = (v2 % ((2147483647) | 1))
    var v29 float32 = float32(float64((float32(1e10) / 1.0)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
