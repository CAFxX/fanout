package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int8; F2 uint64 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint16, p1 int32) bool {
    for fi4 := 0; fi4 < 14; fi4++ {
        if true {
            var v5 uint = 255
            _ = v5
        }
        _ = fi4
    }
    p0 -= (p0 | 255)
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return false
}

func main() {
    var v1 int32 = -1
    v2 := make(map[string]uint32)
    var v3 uint8 = 42
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := GenS1{F0: 130, F1: 255, F2: ("hello" + "\n\t")}
        _ = fi6
        _ = v7
    }
    delete(v2, ("foo bar" + "hello"))
    var v8 uint8 = (v3 % ((2) | 1))
    switch 669 {
    case 127:
        v9 := GenS2{F0: 0}
        _ = v9
    case -183:
        var v10 int16 = 1
        _ = v10
    case 32767:
        v11 := GenS3{F0: -1, F1: 1, F2: 2}
        _ = v11
    }
    v8 += 128
    v12 := []int{}
    if ((!(v1 >= (v1 + 2147483647))) && (1000 <= -790)) {
        if false {
            if (!(!true)) {
                for fi13 := 0; fi13 < 19; fi13++ {
                    var v14 int32 = int32(v8)
                    _ = fi13
                    _ = v14
                }
            } else {
                var v15 int16 = 707
                _ = len(v12)
                _ = v15
            }
            for fi16 := 0; fi16 < 9; fi16++ {
                v17 := GenS4{F0: false, F1: (-128 == 91)}
                fmt.Println(fi16, fi16, fi16)
                v3 = 0
                _ = fi16
                _ = v17
            }
        } else {
            if len(v12) > 0 {
                v12[0] = 767
            }
            var v18 int = -631
            v2[""] = 1
            _ = v18
        }
        for fi19 := 0; fi19 < 10; fi19++ {
            if false {
                var v20 bool = ((v1 != 2147483647) || false)
                for fi21 := 0; fi21 < 18; fi21++ {
                    _ = len(v2)
                    _ = fi21
                }
                for fi22 := 0; fi22 < 8; fi22++ {
                    v23 := GenS3{F0: (fi19 * fi22), F1: 0, F2: 1}
                    _ = fi22
                    _ = v23
                }
                _ = v20
            }
            _ = fi19
        }
    } else {
        v24 := GenS2{F0: -74}
        fmt.Println(v8, v3)
        var v25 uint64 = 1
        _ = v24
        _ = v25
    }
    var v26 uint8 = 202
    for fi27 := 0; fi27 < 15; fi27++ {
        if false {
            v12 = append(v12, (fi27 + (-(fi27))))
        }
        _ = fi27
    }
    var v28 uint8 = 1
    var v29 uint32 = 42
    for fi30 := 0; fi30 < 9; fi30++ {
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        var v31 int8 = -128
        _ = fi30
        _ = v31
    }
    if true {
        var v32 string = ("foo bar" + ("\n\t" + "hello"))
        _ = len(v12)
        _ = v32
    } else {
        v33 := GenS5{F0: 2147483647}
        _ = v33
    }
    fmt.Println(v8, v26)
    var v34 float32 = (1.5 * 1e10)
    if (0 != 32767) {
        var v35 float64 = 1e10
        var v36 int16 = 1000
        v26 = v3
        _ = v35
        _ = v36
    }
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    for fi37 := 0; fi37 < 7; fi37++ {
        v26 = (v28 / ((0) | 1))
        v3 = (^(v26))
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v8
}
