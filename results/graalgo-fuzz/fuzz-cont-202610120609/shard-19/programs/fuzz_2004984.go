package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int }
func genFn0() int64 {
    if true {
        var v4 float32 = 1e10
        var v5 int16 = -32768
        _ = v4
        _ = v5
    }
    if false {
        var v6 float64 = 0.1
        _ = v6
    } else {
        if true {
            var v7 int32 = 257
            _ = v7
        } else {
            if true {
                var v8 uint = 42
                switch ("\n\t" + "\n\t") {
                case "":
                    var v9 uint = (v8 << ((v8) & 63))
                    _ = v9
                default:
                    switch 1 {
                    case 199:
                        v10 := []int{2147483647}
                        _ = v10
                    case 6:
                        var v11 uint16 = 65535
                        _ = v11
                    }
                }
                _ = v8
            } else {
                switch 32767 {
                case 645:
                    var v12 uint8 = 128
                    _ = v12
                case 32768:
                    var v13 bool = (-9223372036854775808 >= -1)
                    _ = v13
                }
                v14 := GenS1{F0: 792, F1: 2}
                var v15 int8 = -128
                _ = v14
                _ = v15
            }
            if (-128 != 39) {
                v16 := GenS2{F0: 1}
                var v17 float32 = 0.1
                _ = v16
                _ = v17
            }
            var v18 uint32 = 922
            _ = v18
        }
    }
    if (1 != -333) {
        var v19 uint8 = 128
        if true {
            v20 := []int{1, 4294967295, 714}
            _ = v20
        }
        var v21 float32 = 0.0
        _ = v19
        _ = v21
    } else {
        v22 := GenS3{F0: 128}
        _ = v22
    }
    if true {
        v23 := GenS3{F0: 2147483647}
        _ = v23
    } else {
        v24 := []int{}
        for fi25 := 0; fi25 < 18; fi25++ {
            var v26 int = -2147483648
            _ = fi25
            _ = v26
        }
        for fi27 := 0; fi27 < 8; fi27++ {
            v24 = append(v24, 9223372036854775807)
            var v28 int32 = -1
            _ = fi27
            _ = v28
        }
        _ = v24
    }
    v29 := GenS1{F0: 42, F1: 1}
    _ = v29
    return 42
}
func genFn1() uint {
    for fi30 := 0; fi30 < 8; fi30++ {
        fmt.Println(fi30)
        _ = fi30
    }
    v31 := []int{147, -447, 128, 9223372036854775807}
    var v32 int8 = 1
    v33 := GenS2{F0: 0}
    _ = v31
    _ = v32
    _ = v33
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 971
    var v3 int8 = 1
    for fi34 := 0; fi34 < 18; fi34++ {
        v2 = (v2 << ((genFn1()) & 63))
        switch -9223372036854775808 {
        case -493:
            fmt.Println(fi34, fi34, v3)
        }
        v35 := GenS3{F0: (fi34 >> ((543) & 63))}
        _ = fi34
        _ = v35
    }
    var v36 int = -398
    v37 := GenS3{F0: v36}
    v38 := []int{}
    v2 += (v2 - v2)
    delete(v1, ("a" + "\n\t"))
    if false {
        v39 := GenS3{F0: 128}
        v2 += (v2 | 274)
        _ = v39
    } else {
        var v40 uint64 = 1
        _ = v40
    }
    var v41 string = "\n\t"
    v3 -= 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v41
}
