package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]uint32)
    v3 := make(map[string]string)
    for fi4 := 0; fi4 < 7; fi4++ {
        if (!false) {
            v5 := []int{fi4, fi4, (^(fi4))}
            _ = v5
        } else {
            switch fi4 {
            case 9223372036854775807:
                if (-9223372036854775808 == -1) {
                    for fi6 := 0; fi6 < 4; fi6++ {
                        var v7 uint32 = 2147483648
                        var v8 float32 = 1.5
                        _ = fi6
                        _ = v7
                        _ = v8
                    }
                    for fi9 := 0; fi9 < 10; fi9++ {
                        v1[fi4] = (!true)
                        var v10 int32 = 0
                        _ = fi9
                        _ = v10
                    }
                } else {
                    for fi11 := 0; fi11 < 12; fi11++ {
                        var v12 uint16 = 32767
                        v13 := []int{-548}
                        v14 := GenS1{F0: (-(fi11)), F1: 4294967295}
                        v12 += v12
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    var v15 uint16 = 0
                    _ = v15
                }
            case 128:
                v16 := []int{-917, fi4}
                _ = v16
            default:
                v17 := GenS1{F0: (fi4 >> ((4294967295) & 63)), F1: 0}
                _ = v17
            }
        }
        var v18 int64 = 9223372036854775807
        _ = fi4
        _ = v18
    }
    var v19 bool = true
    var v20 bool = (!(v19 || v19))
    _ = len(v1)
    v21 := []int{1}
    v22 := GenS2{F0: -1, F1: 312}
    var v23 bool = true
    v24 := GenS3{F0: 0}
    v21 = append(v21, 333)
    v23 = (2 != 1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
