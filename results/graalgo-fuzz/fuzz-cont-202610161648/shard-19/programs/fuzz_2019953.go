package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16 }
func genFn0(p0 string) uint8 {
    v4 := GenS1{F0: -32768, F1: 42, F2: 4294967295}
    p0 = (p0 + p0)
    p0 = p0
    fmt.Println(p0)
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 string = (p0 + ("hello" + ("foo bar" + ("" + (p0 + "foo bar")))))
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = v4
    return 42
}

func main() {
    var v1 int8 = -128
    var v2 float32 = (0.1 / 1.0)
    var v3 int8 = v1
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := GenS2{F0: (2 == 716)}
            if false {
                v2 -= (float32(0.0) + ((v2 * 0.0) / (v2 * (1e10 + 1.5))))
                v10 := GenS3{F0: -1}
                _ = v10
            } else {
                if true {
                    v11 := GenS4{F0: (2147483648 <= 4294967295)}
                    _ = v11
                } else {
                    var v12 uint64 = 42
                    v1 -= v1
                    v12 = (^(v12))
                    var v13 uint16 = 170
                    fmt.Println(v1, v2)
                    v14 := GenS1{F0: -26, F1: (v12 & 550), F2: 257}
                    var v15 bool = true
                    _ = v12
                    _ = v13
                    _ = v14
                    _ = v15
                }
                v2 = 1.5
                v16 := GenS3{F0: 858}
                _ = v16
            }
            var v17 string = ""
            _ = fi8
            _ = v17
            _ = v9
        }
        v2 -= 0.0
        fmt.Println(v3, v1, fi7)
        _ = fi7
    }
    fmt.Println(v3, v1)
    v3 -= 127
    var v18 int32 = 1
    var v19 uint32 = 1
    v2 = -2.25
    var v20 float32 = v2
    for fi21 := 0; fi21 < 2; fi21++ {
        for fi22 := 0; fi22 < 18; fi22++ {
            var v23 uint32 = 768
            v1 -= (v1 << ((255) & 63))
            if (v19 == (v19 - 475)) {
                if (!(137 >= 17)) {
                    v23 = v23
                    var v24 uint = 255
                    v25 := []int{32768}
                    _ = v24
                    _ = v25
                } else {
                    v26 := GenS2{F0: false}
                    fmt.Println(fi22, v1)
                    _ = v26
                }
                var v27 float32 = -2.25
                switch (fi21 & fi21) {
                case -295:
                    v28 := GenS1{F0: -149, F1: 4294967295, F2: 2}
                    _ = v28
                case -406:
                    v29 := GenS5{F0: 32767}
                    _ = v29
                default:
                    v30 := []int{fi22, -2147483648}
                    _ = v30
                }
                _ = v27
            }
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    v1 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
