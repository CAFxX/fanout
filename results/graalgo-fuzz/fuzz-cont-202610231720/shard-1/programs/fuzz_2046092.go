package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint16; F2 int16 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint32 }
func genFn0() uint64 {
    v5 := GenS1{F0: 0}
    var v6 string = "foo bar"
    if (((900 == 32767) && false) || (1 != 550)) {
        fmt.Println(v6)
        v7 := []int{}
        v6 = ("hello" + (v6 + "a"))
        _ = v7
    } else {
        v6 = v6
        v8 := []int{-590}
        _ = v8
    }
    fmt.Println(v6)
    _ = v5
    _ = v6
    return 680
}
func genFn1(p0 uint32, p1 int64) int8 {
    var v9 int16 = 0
    fmt.Println(v9, p1)
    v9 = (v9 / ((-1) | 1))
    v9 += (-(v9))
    fmt.Println(v9, p1)
    _ = p0
    _ = p1
    _ = v9
    return -128
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{}
    v3 := make(map[int]int)
    var v4 string = ("foo bar" + "foo bar")
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := GenS2{F0: 127, F1: 42, F2: 32767}
        fmt.Println(v4)
        _ = fi10
        _ = v11
    }
    var v12 uint16 = 32767
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v12 = 637
    var v13 uint = 2
    if (9223372036854775807 < 42) {
        v14 := GenS2{F0: 489, F1: v12, F2: -185}
        _ = v14
    } else {
        var v15 int64 = 9223372036854775807
        if ((v13 & 42) > v13) {
            var v16 float64 = 0.1
            var v17 string = (v4 + "\n\t")
            var v18 string = v17
            _ = v16
            _ = v17
            _ = v18
        } else {
            for fi19 := 0; fi19 < 3; fi19++ {
                var v20 int = (fi19 * (fi19 * (fi19 << ((371) & 63))))
                v20 += int(v12)
                if true {
                    v12 = v12
                    var v21 uint8 = 2
                    _ = v21
                }
                _ = fi19
                _ = v20
            }
            var v22 float32 = (0.1 - (float32((0.1 + float64(-2.25))) / 0.1))
            _ = v22
        }
        _ = v15
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        v24 := GenS2{F0: (fi23 + fi23), F1: 32767, F2: -238}
        _ = fi23
        _ = v24
    }
    if false {
        v12 = 146
    } else {
        delete(v3, 640)
        if (0.0 == 0.1) {
            var v25 uint32 = 65535
            if len(v2) > 0 {
                v2[0] = 0
            }
            if true {
                delete(v3, -863)
                for fi26 := 0; fi26 < 11; fi26++ {
                    v4 = ("foo bar" + v4)
                    fmt.Println(v4, fi26)
                    v27 := GenS3{F0: 1, F1: 42}
                    _ = fi26
                    _ = v27
                }
                var v28 int = -1
                _ = v28
            }
            _ = v25
        }
        v3[9223372036854775807] = -241
    }
    v29 := GenS3{F0: 4294967295, F1: 42}
    if false {
        var v30 string = ("foo bar" + "")
        v4 = (v30 + "foo bar")
        _ = v30
    } else {
        v2 = append(v2, -2147483648)
    }
    fmt.Println(v4, v13, v13)
    v3[4294967295] = 802
    if (true || false) {
        v31 := GenS2{F0: -388, F1: 42, F2: 163}
        var v32 bool = (1e10 >= 0.0)
        v32 = (!(11 >= 255))
        _ = v31
        _ = v32
    } else {
        if true {
            var v33 float64 = 1e10
            _ = v33
        } else {
            v4 = (("foo bar" + "foo bar") + v4)
            v34 := GenS3{F0: 677, F1: 42}
            _ = v34
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    for fi35 := 0; fi35 < 16; fi35++ {
        fmt.Println(v13, v4)
        var v36 int64 = int64(v12)
        _ = fi35
        _ = v36
    }
    var v37 uint64 = 788
    v12 -= v12
    v38 := GenS4{F0: 939}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
