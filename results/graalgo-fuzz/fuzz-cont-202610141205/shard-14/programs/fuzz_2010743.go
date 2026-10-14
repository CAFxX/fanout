package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 int16 = 42
    var v2 string = ("\n\t" + "a")
    v3 := make(map[uint32]uint32)
    v4 := []uint32{0, 0, 65535, 2}
    v5 := make(map[int]string)
    v6 := make(map[int]int)
    for fi7 := 0; fi7 < 19; fi7++ {
        v1 -= (v1 ^ int16(fi7))
        _ = fi7
    }
    v8 := GenS1{F0: 1}
    if (121 > 255) {
        if len(v4) > 0 {
            v4[0] = 0
        }
        v9 := GenS2{F0: v1}
        if len(v4) > 0 {
            v4[0] = 300
        }
        _ = v9
    }
    v10 := GenS2{F0: v1}
    if (100 <= 42) {
        v11 := GenS1{F0: 167}
        _ = v11
    }
    if false {
        v2 = v2
        _ = len(v4)
        var v12 int16 = 942
        _ = v12
    } else {
        v2 = v2
        v2 = (v2 + v2)
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        v1 -= (^(v1))
        _ = fi13
    }
    var v14 int8 = -10
    if false {
        for fi15 := 0; fi15 < 4; fi15++ {
            if false {
                v14 += v14
                var v16 bool = (255 == 255)
                var v17 float64 = 0.0
                _ = v16
                _ = v17
            }
            if (-659 != 42) {
                if false {
                    v14 = 100
                    v2 = v2
                }
                v2 = v2
            }
            var v18 string = "\n\t"
            _ = fi15
            _ = v18
        }
        switch 42 {
        case 2147483647:
            if (233 != 200) {
                var v19 float32 = (1.5 + 3.14159)
                var v20 int = 32767
                _ = v19
                _ = v20
            } else {
                var v21 float64 = 0.0
                fmt.Println(v14)
                v14 -= -128
                _ = v21
            }
        case 0:
            var v22 uint16 = 32767
            _ = v22
        case -280:
            v14 -= v14
        default:
            v14 = v14
        }
        for fi23 := 0; fi23 < 11; fi23++ {
            v14 -= (-(v14))
            v24 := GenS1{F0: 200}
            if (-2.25 >= ((-2.25 / 1.0) / 0.1)) {
                if true {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v25 int = (fi23 * (fi23 ^ (fi23 * fi23)))
                    _ = len(v6)
                    _ = v25
                }
                for fi26 := 0; fi26 < 6; fi26++ {
                    v14 = v14
                    var v27 uint8 = 117
                    if false {
                        _ = len(v4)
                    } else {
                        var v28 uint64 = 2
                        _ = v28
                    }
                    _ = fi26
                    _ = v27
                }
            }
            _ = fi23
            _ = v24
        }
    } else {
        var v29 float32 = 0.1
        _ = v29
    }
    var v30 string = v2
    fmt.Println(v2, v14)
    var v31 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
