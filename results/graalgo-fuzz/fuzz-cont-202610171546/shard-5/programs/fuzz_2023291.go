package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int32, p1 int8, p2 int8) uint64 {
    for fi6 := 0; fi6 < 18; fi6++ {
        p2 += (-(p1))
        v7 := []int{-400, (fi6 - -9223372036854775808), 4294967295, (fi6 << ((2) & 63))}
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 1}
    v9 := GenS1{F0: 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 748
}

func main() {
    var v1 float64 = 0.1
    var v2 uint8 = 153
    v3 := []uint32{2147483647}
    var v4 int16 = 225
    var v5 string = "foo bar"
    v4 -= -579
    v1 += v1
    var v10 int16 = v4
    var v11 uint16 = 32767
    var v12 uint = 255
    v13 := GenS2{F0: uint8(v11)}
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 uint16 = (v11 % ((42) | 1))
        _ = fi14
        _ = v15
    }
    if false {
        v4 = v10
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if false {
            var v16 uint16 = v11
            for fi17 := 0; fi17 < 19; fi17++ {
                for fi18 := 0; fi18 < 14; fi18++ {
                    if false {
                        var v19 int16 = (v10 % ((0) | 1))
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v20 := GenS2{F0: uint8(v12)}
                        v21 := GenS3{F0: (v11 + (^(v11)))}
                        v2 += v2
                        _ = v19
                        _ = v20
                        _ = v21
                    } else {
                        switch v5 {
                        case "\n\t":
                            var v22 int64 = 42
                            _ = v22
                        case "a":
                            var v23 int64 = 1
                            _ = v23
                        }
                    }
                    _ = fi18
                }
                v3 = append(v3, uint32(v2))
                _ = fi17
            }
            _ = v16
        } else {
            fmt.Println(v11)
            for fi24 := 0; fi24 < 11; fi24++ {
                if (2 > (^(v2))) {
                    v2 -= (^(v2))
                }
                for fi25 := 0; fi25 < 8; fi25++ {
                    var v26 int = fi25
                    var v27 uint8 = v2
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                if false {
                    var v28 string = (v5 + v5)
                    _ = v28
                }
                _ = fi24
            }
        }
    } else {
        if ((true || (!false)) || true) {
            var v29 float32 = ((-2.25 / 3.14159) - (float32((v1 - (float64(3.14159) - v1))) / -2.25))
            _ = v29
        }
        if ((v1 - v1) == (0.0 + 0.1)) {
            for fi30 := 0; fi30 < 9; fi30++ {
                v5 = ("foo bar" + v5)
                fmt.Println(v2)
                _ = fi30
            }
        } else {
            v10 += (-(v4))
            v11 = 1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
