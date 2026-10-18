package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int64, p1 int) uint32 {
    p1 -= (^(p1))
    for fi6 := 0; fi6 < 11; fi6++ {
        if false {
            v7 := GenS1{F0: 2}
            v8 := GenS1{F0: 65535}
            _ = v7
            _ = v8
        } else {
            switch (fi6 / ((32767) | 1)) {
            case -278:
                for fi9 := 0; fi9 < 6; fi9++ {
                    v10 := []int{}
                    _ = len(v10)
                    p0 = p0
                    _ = fi9
                    _ = v10
                }
            default:
                var v11 int32 = -259
                _ = v11
            }
            p0 = (-(p0))
            var v12 int16 = -1
            _ = v12
        }
        _ = fi6
    }
    v13 := GenS2{F0: 1, F1: 4294967295, F2: p0}
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = v13
    return 2147483648
}

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 887
    v3 := make(map[string]uint32)
    var v4 int8 = 42
    v5 := make(map[int]uint32)
    v2 = 9223372036854775807
    var v14 bool = (0 >= 2147483647)
    var v15 int32 = 0
    v2 = 42
    for fi16 := 0; fi16 < 11; fi16++ {
        switch 845 {
        case -20:
            var v17 int = (fi16 / (((-(fi16))) | 1))
            _ = v17
        case -985:
            for fi18 := 0; fi18 < 2; fi18++ {
                v19 := GenS2{F0: 127, F1: 1, F2: v2}
                v20 := GenS3{F0: 734, F1: (-(v2))}
                _ = fi18
                _ = v19
                _ = v20
            }
        case 4294967295:
            v2 = v2
        }
        v21 := GenS2{F0: 42, F1: genFn0(0, fi16), F2: (v2 / (((v2 >> ((333) & 63))) | 1))}
        for fi22 := 0; fi22 < 6; fi22++ {
            var v23 int32 = (v15 * (v15 | v15))
            _ = fi22
            _ = v23
        }
        _ = fi16
        _ = v21
    }
    v14 = v14
    var v24 string = "hello"
    if (880 < 339) {
        var v25 bool = (false && v14)
        _ = v25
    }
    switch 128 {
    case -9223372036854775808:
        _ = len(v3)
    case 32768:
        var v26 uint16 = 2
        _ = v26
    default:
        for fi27 := 0; fi27 < 9; fi27++ {
            v14 = ((!true) && (!(v4 < -128)))
            _ = fi27
        }
    }
    v2 = (^(v2))
    v1["hello"] = 2147483647
    var v28 string = ("\n\t" + "\n\t")
    _ = len(v3)
    for fi29 := 0; fi29 < 17; fi29++ {
        for fi30 := 0; fi30 < 2; fi30++ {
            v31 := GenS3{F0: 4294967295, F1: -9223372036854775808}
            if true {
                var v32 int8 = (v4 & v4)
                var v33 uint16 = 255
                _ = v32
                _ = v33
            } else {
                v34 := GenS4{F0: 1, F1: 59}
                v15 -= (v15 - v15)
                _ = v34
            }
            v14 = ((true && (-32768 > 122)) && ((!(true && ((0 >= 781) || v14))) || (!v14)))
            v24 = "hello"
            _ = len(v5)
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
