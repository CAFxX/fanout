package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 string; F1 int8 }
func genFn0() uint8 {
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := []int{(fi4 ^ 128), fi4, -9223372036854775808}
        for fi6 := 0; fi6 < 9; fi6++ {
            _ = len(v5)
            if false {
                var v7 float32 = 0.1
                _ = v7
            } else {
                var v8 int32 = 63
                fmt.Println(fi4, fi6)
                _ = v8
            }
            if len(v5) > 0 {
                v5[0] = (-(fi6))
            }
            _ = fi6
        }
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 uint32 = 850
            if false {
                v10 = (v10 / (((v10 & v10)) | 1))
            }
            var v11 uint64 = 67
            if ((-2.25 >= (0.0 - 0.1)) || true) {
                var v12 int32 = 42
                var v13 int8 = 60
                var v14 int16 = -183
                v15 := GenS1{F0: 0, F1: (v12 ^ (v12 / ((v12) | 1)))}
                _ = v12
                _ = v13
                _ = v14
                _ = v15
            } else {
                var v16 int32 = 42
                fmt.Println(v10, fi4)
                fmt.Println(v10)
                _ = v16
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi4
        _ = v5
    }
    v17 := GenS2{F0: 671, F1: 42}
    _ = v17
    return 128
}

func main() {
    var v1 int64 = -915
    v2 := make(map[string]uint32)
    v3 := []string{}
    var v18 int32 = 42
    for fi19 := 0; fi19 < 9; fi19++ {
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        _ = fi19
    }
    v20 := GenS3{F0: true, F1: -1}
    v21 := GenS2{F0: 0, F1: 65535}
    var v22 uint = 176
    if false {
        v1 = (v1 / ((v1) | 1))
        v23 := GenS4{F0: "", F1: 127}
        _ = v23
    }
    v18 = 1
    v24 := GenS1{F0: 748, F1: -2147483648}
    for fi25 := 0; fi25 < 20; fi25++ {
        var v26 uint32 = 1
        delete(v2, "foo bar")
        var v27 uint = v22
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
}
