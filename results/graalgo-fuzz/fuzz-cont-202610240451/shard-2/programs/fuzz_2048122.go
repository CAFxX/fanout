package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int16; F1 int64 }
func genFn0(p0 int, p1 int8, p2 uint32) uint64 {
    var v6 int16 = -906
    switch p0 {
    case 4294967295:
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := []int{fi7, 0, p0, fi7}
            var v9 int32 = 0
            _ = fi7
            _ = v8
            _ = v9
        }
    default:
        v10 := GenS1{F0: 42}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}

func main() {
    var v1 int = 86
    var v2 int8 = -128
    v3 := []string{}
    var v4 uint16 = 0
    var v5 uint8 = 0
    if false {
        v11 := GenS1{F0: 42}
        _ = v11
    }
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    var v12 uint16 = v4
    for fi13 := 0; fi13 < 15; fi13++ {
        if (!true) {
            if (v5 != 1) {
                for fi14 := 0; fi14 < 3; fi14++ {
                    var v15 float64 = -2.25
                    v5 += (v5 >> ((18446744073709551615) & 63))
                    var v16 uint32 = 4294967295
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            } else {
                var v17 uint8 = uint8(v4)
                v18 := GenS2{F0: -1, F1: (v17 + 255)}
                _ = v17
                _ = v18
            }
            for fi19 := 0; fi19 < 6; fi19++ {
                v12 -= 897
                _ = fi19
            }
        }
        var v20 int64 = 42
        var v21 int32 = 42
        _ = fi13
        _ = v20
        _ = v21
    }
    v2 += (v2 - (v2 / ((-128) | 1)))
    v4 -= v4
    for fi22 := 0; fi22 < 9; fi22++ {
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 int16 = int16(v12)
            var v25 float64 = float64(-2.25)
            var v26 uint8 = 0
            fmt.Println(v12, fi22)
            v27 := GenS2{F0: 0, F1: v26}
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
            _ = v27
        }
        _ = fi22
    }
    if true {
        v28 := GenS1{F0: 214}
        if len(v3) > 0 {
            v3[0] = "a"
        }
        if true {
            if (((v12 > 42) && (4294967295 > uint32(v5))) || true) {
                if true {
                    var v29 int32 = 42
                    _ = v29
                } else {
                    if false {
                        var v30 uint16 = (v4 * 1)
                        _ = v30
                    }
                    var v31 int8 = (v2 | (v2 << ((1) & 63)))
                    v4 = (v12 - 0)
                    _ = v31
                }
                var v32 int64 = -1
                _ = v32
            }
            if (true || ((false || true) || true)) {
                v33 := GenS1{F0: 42}
                _ = v33
            }
            v4 = v4
        } else {
            fmt.Println(v4)
        }
        _ = v28
    }
    v1 = (^(v1))
    if false {
        for fi34 := 0; fi34 < 4; fi34++ {
            v35 := GenS1{F0: 1}
            _ = fi34
            _ = v35
        }
        for fi36 := 0; fi36 < 12; fi36++ {
            var v37 int64 = -1
            if true {
                var v38 int16 = 32767
                _ = v38
            } else {
                v3 = append(v3, "foo bar")
            }
            _ = fi36
            _ = v37
        }
    }
    v39 := GenS3{F0: -32768, F1: -1}
    v5 -= (v5 ^ (v5 / ((v5) | 1)))
    var v40 int16 = 0
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
