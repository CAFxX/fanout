package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int16; F1 int16 }
func genFn0(p0 int16) uint {
    v6 := GenS1{F0: 128}
    var v7 uint = 170
    _ = p0
    _ = v6
    _ = v7
    return (v7 - v7)
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 4294967295
    v3 := []int{128, -360}
    var v4 bool = false
    v5 := make(map[string]uint32)
    if v4 {
        for fi8 := 0; fi8 < 4; fi8++ {
            v4 = false
            var v9 uint64 = (v2 << ((374) & 63))
            _ = fi8
            _ = v9
        }
        v4 = (!true)
        v2 = v2
    } else {
        v10 := GenS2{F0: 0, F1: 18446744073709551615}
        fmt.Println(v4, v4)
        v2 += v2
        _ = v10
    }
    v2 = v2
    v2 += (^(v2))
    v11 := GenS3{F0: -592, F1: 0}
    v12 := v5["hello"]
    v2 += (v2 - (v2 * v2))
    if (v4 && true) {
        var v13 int16 = 325
        var v14 string = (("" + "a") + "\n\t")
        if v4 {
            if (-128 != 0) {
                v2 += (v2 % ((v2) | 1))
            } else {
                v15 := GenS1{F0: 42}
                fmt.Println(v12)
                if (!(62 > 127)) {
                    if v4 {
                        var v16 float64 = float64(3.14159)
                        var v17 float64 = v16
                        _ = v16
                        _ = v17
                    } else {
                        var v18 int = -2147483648
                        v13 = v13
                        switch v14 {
                        case "a":
                            v3 = append(v3, (v18 | (v18 & -834)))
                        case "":
                            var v19 string = v14
                            _ = v19
                        }
                        _ = v18
                    }
                }
                _ = v15
            }
            v1[(^(v12))] = (1 != 65535)
        } else {
            v14 = "a"
            var v20 int32 = 2147483647
            _ = len(v1)
            _ = v20
        }
        _ = v13
        _ = v14
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            if true {
                v23 := GenS1{F0: uint8(v12)}
                var v24 uint64 = v2
                fmt.Println(fi21, v4, v12)
                _ = v23
                _ = v24
            }
            var v25 int16 = 32767
            v4 = (((!(v4 || (v12 <= (^(v12))))) || v4) || v4)
            v3 = append(v3, (fi22 >> ((101) & 63)))
            _ = fi22
            _ = v25
        }
        _ = fi21
    }
    fmt.Println(v2, v4)
    var v26 string = "hello"
    if false {
        v12 = 0
        var v27 int = -9223372036854775808
        _ = v27
    }
    if (47 >= 14) {
        v26 = "foo bar"
        v2 = v2
    } else {
        for fi28 := 0; fi28 < 19; fi28++ {
            for fi29 := 0; fi29 < 18; fi29++ {
                v4 = false
                _ = fi29
            }
            _ = fi28
        }
        for fi30 := 0; fi30 < 18; fi30++ {
            var v31 int8 = 127
            _ = fi30
            _ = v31
        }
    }
    v32 := v1[v12]
    for fi33 := 0; fi33 < 12; fi33++ {
        var v34 float32 = 1.5
        v4 = v32
        for fi35 := 0; fi35 < 16; fi35++ {
            v36 := v5[(v26 + v26)]
            var v37 int32 = 42
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = fi33
        _ = v34
    }
    var v38 float32 = ((float32((-2.25 * 0.0)) + (float32(((3.14159 + 1.5) + (3.14159 - 1e10))) * (1.5 - -2.25))) / (1.5 / 1.5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v38
    _ = v4
    _ = v5
}
