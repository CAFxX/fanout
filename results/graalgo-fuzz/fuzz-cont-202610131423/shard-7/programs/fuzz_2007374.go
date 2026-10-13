package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int32; F1 uint }
func genFn0(p0 int64) int64 {
    fmt.Println(p0)
    v6 := []int{}
    v7 := GenS1{F0: p0}
    if len(v6) > 0 {
        v6[0] = 128
    }
    var v8 float32 = (float32(-2.25) / 1.0)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return p0
}
func genFn1(p0 int16) int8 {
    var v9 uint16 = 402
    if false {
        var v10 uint8 = 1
        _ = v10
    } else {
        v11 := GenS2{F0: 2, F1: 437}
        v9 += v9
        _ = v11
    }
    v12 := []int{-9223372036854775808}
    _ = p0
    _ = v12
    _ = v9
    return 1
}

func main() {
    var v1 uint32 = 2
    var v2 uint64 = 4294967295
    v3 := make(map[int]string)
    v4 := make(map[string]string)
    var v5 uint32 = (v1 + (v1 << ((255) & 63)))
    delete(v4, "")
    fmt.Println(v1, v2)
    if false {
        v13 := GenS2{F0: v2, F1: (v2 >> ((4294967295) & 63))}
        if false {
            var v14 uint32 = v5
            if (!true) {
                v5 += (v5 + v1)
            } else {
                v15 := GenS3{F0: 255, F1: 698}
                _ = v15
            }
            _ = v14
        } else {
            for fi16 := 0; fi16 < 20; fi16++ {
                fmt.Println(v2)
                v1 = (v1 >> ((65535) & 63))
                v17 := GenS4{F0: (fi16 / ((fi16) | 1))}
                _ = fi16
                _ = v17
            }
            v2 = (^(v2))
            if true {
                if true {
                    if ((42 > 65535) || true) {
                        if (9223372036854775807 <= 32768) {
                            v5 = v5
                        }
                        fmt.Println(v1)
                    } else {
                        v18 := GenS5{F0: -119}
                        for fi19 := 0; fi19 < 18; fi19++ {
                            v20 := GenS5{F0: 1}
                            v5 = v1
                            v21 := []int{}
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                        _ = v18
                    }
                    if (-32768 <= 1000) {
                        for fi22 := 0; fi22 < 15; fi22++ {
                            v23 := GenS3{F0: 589, F1: 255}
                            _ = fi22
                            _ = v23
                        }
                    }
                }
            }
        }
        _ = v13
    }
    var v24 string = "\n\t"
    var v25 int64 = 1
    v26 := GenS1{F0: v25}
    var v27 int32 = -843
    v28 := GenS5{F0: v27}
    for fi29 := 0; fi29 < 6; fi29++ {
        v30 := []int{}
        v1 -= (^(v5))
        for fi31 := 0; fi31 < 15; fi31++ {
            if false {
                v32 := v4[v24]
                var v33 uint32 = (v5 + (v5 * 1))
                v5 = 1
                fmt.Println(v1, fi31)
                _ = v32
                _ = v33
            } else {
                v2 += v2
            }
            var v34 uint = 0
            _ = fi31
            _ = v34
        }
        _ = fi29
        _ = v30
    }
    v35 := v4[("" + ((("hello" + v24) + (v24 + ((((((v24 + v24) + "\n\t") + v24) + "hello") + ("" + (("" + v24) + v24))) + v24))) + "\n\t"))]
    v36 := GenS6{F0: (^(v27)), F1: 1}
    var v37 int16 = -32768
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
