package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS2 struct { F0 int64; F1 uint16; F2 int32 }
func genFn0(p0 bool) int8 {
    v4 := []int{-190, 32767, 127}
    p0 = p0
    _ = len(v4)
    fmt.Println(p0)
    _ = p0
    _ = v4
    return -128
}

func main() {
    var v1 uint32 = 65535
    var v2 bool = false
    var v3 string = ""
    var v5 uint16 = 42
    if (!true) {
        v6 := GenS1{F0: 2147483647, F1: v5, F2: -315}
        switch -462 {
        case 1:
            v7 := []int{2147483647}
            _ = v7
        case 32767:
            v2 = v2
        case 4294967295:
            v8 := []int{32768, 679, 965}
            _ = v8
        }
        _ = v6
    }
    if (false || v2) {
        v5 = 0
    }
    fmt.Println(v5)
    v2 = ((^(v1)) < (v1 / (((v1 - v1)) | 1)))
    if false {
        var v9 float64 = 3.14159
        var v10 int16 = -555
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 int32 = int32(fi11)
            for fi13 := 0; fi13 < 7; fi13++ {
                v14 := GenS1{F0: v1, F1: 255, F2: -1}
                v12 -= v12
                v15 := GenS1{F0: (^(v1)), F1: v5, F2: 32767}
                _ = fi13
                _ = v14
                _ = v15
            }
            v16 := []int{fi11, fi11, fi11, fi11}
            _ = fi11
            _ = v12
            _ = v16
        }
        v3 = (v3 + ((v3 + v3) + v3))
        if true {
            v17 := GenS2{F0: -9223372036854775808, F1: (v5 >> ((0) & 63)), F2: 0}
            _ = v17
        }
    }
    if false {
        for fi18 := 0; fi18 < 5; fi18++ {
            if ((v3 + "hello") <= v3) {
                v19 := []int{}
                _ = v19
            } else {
                for fi20 := 0; fi20 < 4; fi20++ {
                    v21 := GenS1{F0: 4294967295, F1: (v5 + (v5 | (v5 * 65535))), F2: -1}
                    _ = fi20
                    _ = v21
                }
            }
            _ = fi18
        }
    } else {
        for fi22 := 0; fi22 < 8; fi22++ {
            v3 = v3
            var v23 uint16 = (v5 & v5)
            _ = fi22
            _ = v23
        }
        if (v1 > 42) {
            var v24 uint = 848
            if (v2 || v2) {
                v25 := []int{-511}
                _ = v25
            } else {
                v24 -= 2
            }
            _ = v24
        } else {
            var v26 int32 = 2147483647
            v2 = (!false)
            _ = v26
        }
        v27 := []int{725, 914, -1}
        _ = v27
    }
    var v28 uint64 = 739
    fmt.Println(v3, v1, v28)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v5
}
