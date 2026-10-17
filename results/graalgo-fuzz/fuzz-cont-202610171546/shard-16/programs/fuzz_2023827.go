package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint }
func genFn0(p0 string, p1 uint16, p2 int8) int16 {
    if (p0 > "hello") {
        v7 := GenS1{F0: 1}
        _ = v7
    } else {
        var v8 bool = false
        var v9 uint8 = 34
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: -299}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 889
}
func genFn1(p0 int, p1 uint16) uint32 {
    if (660 >= 0) {
        var v11 uint = 870
        _ = v11
    }
    if (-9223372036854775808 >= 42) {
        for fi12 := 0; fi12 < 20; fi12++ {
            p1 = 32767
            _ = fi12
        }
    } else {
        if true {
            var v13 int32 = 42
            v13 += 0
            fmt.Println(v13)
            _ = v13
        } else {
            v14 := []int{p0}
            _ = v14
        }
        p1 = (p1 - (p1 - (p1 * (p1 - p1))))
        fmt.Println(p1, p0)
    }
    if (471 > 4294967295) {
        v15 := []int{}
        p1 = 1
        var v16 string = ("foo bar" + "a")
        _ = v15
        _ = v16
    } else {
        switch p0 {
        case 32768:
            p1 -= 32767
        case -1:
            for fi17 := 0; fi17 < 6; fi17++ {
                p0 += (fi17 << ((255) & 63))
                v18 := []int{}
                for fi19 := 0; fi19 < 6; fi19++ {
                    p0 += (p0 ^ (fi17 | (^(fi17))))
                    v20 := GenS1{F0: -1}
                    var v21 string = ""
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                _ = fi17
                _ = v18
            }
        case -420:
            if false {
                for fi22 := 0; fi22 < 3; fi22++ {
                    for fi23 := 0; fi23 < 17; fi23++ {
                        v24 := []int{fi23, 42, 0, (p0 / ((p0) | 1))}
                        _ = fi23
                        _ = v24
                    }
                    _ = fi22
                }
                var v25 int64 = 1
                v26 := GenS2{F0: 127}
                _ = v25
                _ = v26
            }
        }
        var v27 uint = 18446744073709551615
        _ = v27
    }
    var v28 uint = 0
    _ = p0
    _ = p1
    _ = v28
    return 71
}

func main() {
    var v1 int8 = -27
    v2 := []int{}
    v3 := []int{42, 0}
    var v4 float32 = ((0.1 + (0.0 / 0.1)) - 3.14159)
    var v5 int16 = 32767
    v6 := make(map[uint32]uint32)
    _ = len(v6)
    v29 := GenS3{F0: 65535}
    for fi30 := 0; fi30 < 17; fi30++ {
        v31 := GenS4{F0: fi30}
        _ = fi30
        _ = v31
    }
    fmt.Println(v5, v5, v4)
    var v32 uint64 = 42
    v32 += 23
    fmt.Println(v5, v5, v5)
    if (0 >= -2147483648) {
        for fi33 := 0; fi33 < 19; fi33++ {
            v5 = 1000
            _ = fi33
        }
        for fi34 := 0; fi34 < 2; fi34++ {
            fmt.Println(v4)
            _ = fi34
        }
        var v35 float32 = ((0.1 + 3.14159) / (v4 + ((float32(3.14159) * 3.14159) * v4)))
        _ = v35
    } else {
        var v36 int32 = -2147483648
        v36 = 0
        _ = v36
    }
    v37 := GenS4{F0: -969}
    v4 = float32(1.5)
    v38 := GenS5{F0: 2}
    switch int(v5) {
    case 128:
        v1 -= (v1 & 100)
    case 32767:
        v39 := v6[914]
        _ = v39
    case 9223372036854775807:
        v32 -= (v32 % ((v32) | 1))
    default:
        switch 128 {
        case -2147483648:
            v3 = append(v3, -957)
        default:
            fmt.Println(v4, v1, v4)
        }
    }
    var v40 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v37
    _ = v38
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
