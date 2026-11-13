package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 int8; F1 int32 }

func main() {
    v1 := []uint32{4294967295, 0}
    var v2 uint16 = 255
    var v3 int16 = -935
    var v4 bool = true
    v5 := []int{1, 0}
    v6 := make(map[int]bool)
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi8
        }
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := v6[1]
            v4 = (0 == 189)
            if (uint32(v2) != 759) {
                delete(v6, (fi9 ^ (-(fi7))))
            }
            _ = fi9
            _ = v10
        }
        _ = fi7
    }
    v11 := GenS1{F0: 0}
    var v12 bool = ((672 != 0) && v4)
    v12 = (!v4)
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 uint = 780
        _ = fi13
        _ = v14
    }
    switch 32767 {
    case 32767:
        if false {
            v4 = (2147483647 > 0)
            for fi15 := 0; fi15 < 11; fi15++ {
                var v16 int16 = v3
                v17 := GenS2{F0: 91, F1: (v2 * 65535)}
                _ = fi15
                _ = v16
                _ = v17
            }
        } else {
            if ((false || false) && v4) {
                var v18 uint = 42
                v18 += v18
                for fi19 := 0; fi19 < 7; fi19++ {
                    switch 32767 {
                    case -681:
                        var v20 string = "foo bar"
                        _ = v20
                    case -2147483648:
                        var v21 uint8 = 1
                        _ = v21
                    case 2147483647:
                        v22 := GenS3{F0: -9223372036854775808, F1: -110}
                        _ = v22
                    default:
                        fmt.Println(v2, v2, v18)
                    }
                    _ = fi19
                }
                _ = v18
            }
            for fi23 := 0; fi23 < 9; fi23++ {
                switch ("a" + ("\n\t" + "foo bar")) {
                case "":
                    var v24 int8 = 100
                    _ = v24
                }
                _ = fi23
            }
        }
    case 127:
        v4 = false
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        v2 -= (v2 << ((0) & 63))
        var v26 float64 = float64((1e10 - 0.1))
        v12 = v12
        _ = fi25
        _ = v26
    }
    if true {
        if (4294967295 == 4294967295) {
            var v27 int32 = -2147483648
            if (0 != uint64(v2)) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v2 = v2
            }
            _ = v27
        } else {
            var v28 uint64 = 2
            var v29 uint8 = 255
            fmt.Println(v4, v12)
            _ = v28
            _ = v29
        }
        v30 := GenS4{F0: -128, F1: 133}
        _ = v30
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        var v32 uint = 255
        _ = fi31
        _ = v32
    }
    var v33 int16 = (v3 * (-(v3)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
