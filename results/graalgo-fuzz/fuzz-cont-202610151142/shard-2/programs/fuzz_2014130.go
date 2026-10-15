package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint32; F1 uint; F2 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 uint64; F1 int }
func genFn0(p0 uint, p1 int) uint32 {
    for fi4 := 0; fi4 < 7; fi4++ {
        if false {
            fmt.Println(p0)
            if false {
                v5 := GenS1{F0: -515, F1: 1, F2: 2}
                if ((p1 / ((2147483647) | 1)) != (p1 & (^(p1)))) {
                    v6 := GenS1{F0: 766, F1: 32767, F2: 255}
                    if true {
                        var v7 bool = true
                        _ = v7
                    } else {
                        v8 := GenS1{F0: 0, F1: -992, F2: 2}
                        p0 += p0
                        _ = v8
                    }
                    v9 := []int{}
                    _ = v6
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 18; fi10++ {
                        v11 := GenS1{F0: 32767, F1: 1, F2: 1}
                        _ = fi10
                        _ = v11
                    }
                }
                v12 := GenS2{F0: 2147483647, F1: 0}
                _ = v12
                _ = v5
            } else {
                v13 := []int{32768}
                var v14 int16 = -781
                _ = v13
                _ = v14
            }
        } else {
            var v15 int32 = -1
            for fi16 := 0; fi16 < 2; fi16++ {
                fmt.Println(fi16)
                var v17 float64 = (-2.25 + ((3.14159 + 0.0) * (0.1 - float64(1e10))))
                v18 := GenS3{F0: 4294967295, F1: p0, F2: false}
                _ = fi16
                _ = v17
                _ = v18
            }
            p1 -= (p1 * (p1 << (((p0 >> ((p0) & 63))) & 63)))
            _ = v15
        }
        _ = fi4
    }
    p0 -= (p0 + (p0 << ((p0) & 63)))
    v19 := []int{(^(p1)), p1, p1, 2147483647}
    v20 := GenS4{F0: 32767}
    v21 := GenS5{F0: p0, F1: 141}
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    _ = v21
    return 785
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = 0.0
    v3 := make(map[uint32]uint32)
    var v22 uint16 = 0
    v2 += (v2 + (float64((float32(3.14159) * 0.0)) + v2))
    if false {
        var v23 float32 = (0.0 / 1.0)
        v24 := GenS6{F0: 205, F1: -525}
        v3[1] = 42
        _ = v23
        _ = v24
    }
    var v25 float32 = 1.5
    if true {
        var v26 uint16 = (^(v22))
        for fi27 := 0; fi27 < 1; fi27++ {
            for fi28 := 0; fi28 < 12; fi28++ {
                v29 := []int{fi27, fi27, -597}
                var v30 uint8 = 0
                var v31 string = ""
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
            var v32 float64 = (v2 - ((v2 + 1e10) + -2.25))
            _ = fi27
            _ = v32
        }
        v33 := GenS6{F0: 1, F1: 9223372036854775807}
        _ = v26
        _ = v33
    } else {
        if (true || false) {
            for fi34 := 0; fi34 < 3; fi34++ {
                for fi35 := 0; fi35 < 8; fi35++ {
                    v22 = 1
                    _ = fi35
                }
                var v36 string = "hello"
                v1[(((v36 + v36) + v36) + (("" + "a") + v36))] = v36
                _ = fi34
                _ = v36
            }
            var v37 int32 = -2147483648
            for fi38 := 0; fi38 < 6; fi38++ {
                var v39 int8 = 53
                var v40 int8 = v39
                _ = fi38
                _ = v39
                _ = v40
            }
            _ = v37
        } else {
            var v41 int64 = 0
            _ = v41
        }
        var v42 uint8 = 1
        v43 := GenS3{F0: 863, F1: uint(v22), F2: (42 != 4294967295)}
        _ = v42
        _ = v43
    }
    if false {
        v22 += (v22 % ((v22) | 1))
    }
    var v44 int64 = 1
    v45 := v1["hello"]
    v2 = (0.1 - (v2 + -2.25))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v44
    _ = v45
}
