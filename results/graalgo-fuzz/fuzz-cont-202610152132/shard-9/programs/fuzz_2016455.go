package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint32, p1 uint8, p2 string) uint {
    var v4 uint32 = 1
    var v5 float32 = -2.25
    if false {
        v6 := []int{4294967295, -9223372036854775808}
        _ = v6
    } else {
        p0 = p0
        var v7 int16 = 32767
        _ = v7
    }
    fmt.Println(p1, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 65535
}
func genFn1(p0 int64, p1 uint64, p2 uint8) int32 {
    v8 := []int{-619, 945, 1, 32767}
    p0 += (p0 * (p0 * p0))
    p2 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 557
}

func main() {
    var v1 bool = true
    var v2 uint32 = 2147483648
    v3 := make(map[string]uint32)
    fmt.Println(v1)
    switch -219 {
    case 32768:
        var v9 uint64 = 15
        _ = v9
    case 0:
        var v10 uint32 = (^(v2))
        _ = v10
    case 32767:
        var v11 uint = 65535
        _ = v11
    default:
        if v1 {
            v12 := GenS1{F0: v1, F1: 42}
            _ = v12
        } else {
            for fi13 := 0; fi13 < 2; fi13++ {
                for fi14 := 0; fi14 < 3; fi14++ {
                    delete(v3, "a")
                    _ = fi14
                }
                _ = fi13
            }
            if false {
                fmt.Println(v1, v1)
                var v15 uint8 = 2
                _ = v15
            } else {
                var v16 string = "a"
                for fi17 := 0; fi17 < 9; fi17++ {
                    for fi18 := 0; fi18 < 13; fi18++ {
                        v19 := GenS1{F0: (0 <= 188), F1: 1}
                        if (32767 <= 33) {
                            fmt.Println(v16, v2, v2)
                            v1 = false
                        }
                        _ = fi18
                        _ = v19
                    }
                    var v20 int16 = -938
                    for fi21 := 0; fi21 < 5; fi21++ {
                        v22 := GenS2{F0: v2, F1: (v2 - (v2 << ((genFn0(2, 92, v16)) & 63))), F2: v20}
                        switch 0 {
                        case 4294967295:
                            fmt.Println(v2, v2)
                        case 528:
                            v20 += v20
                        }
                        _ = fi21
                        _ = v22
                    }
                    _ = fi17
                    _ = v20
                }
                _ = v16
            }
            for fi23 := 0; fi23 < 4; fi23++ {
                fmt.Println(fi23)
                v24 := []int{fi23, fi23}
                v25 := GenS3{F0: 517, F1: -1}
                _ = fi23
                _ = v24
                _ = v25
            }
        }
    }
    fmt.Println(v1)
    var v26 int16 = 1
    v27 := GenS3{F0: 32767, F1: v26}
    v28 := []int{}
    fmt.Println(v2, v2, v2)
    v26 -= 561
    if true {
        v26 = (v26 * 32767)
        v29 := GenS2{F0: (v2 ^ (^(v2))), F1: 0, F2: v26}
        _ = v29
    } else {
        switch 42 {
        case 32767:
            if (!true) {
                var v30 float32 = float32(3.14159)
                for fi31 := 0; fi31 < 16; fi31++ {
                    v26 += 0
                    _ = len(v28)
                    _ = fi31
                }
                v32 := GenS4{F0: 226}
                _ = v30
                _ = v32
            }
        case -2147483648:
            var v33 int64 = -1
            _ = v33
        }
    }
    fmt.Println(v26, v1)
    var v34 uint16 = 255
    fmt.Println(v2)
    var v35 uint8 = 128
    _ = len(v3)
    var v36 bool = (3.14159 == (1.5 - ((1e10 * -2.25) + 3.14159)))
    var v37 int8 = 42
    v35 += v35
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}
