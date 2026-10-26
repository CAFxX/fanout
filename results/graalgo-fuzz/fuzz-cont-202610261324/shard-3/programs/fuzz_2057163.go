package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS5 struct { F0 int8; F1 bool }

func main() {
    var v1 uint8 = 0
    v2 := []int{127, -385, 42}
    v3 := make(map[int]int)
    var v4 int64 = 541
    v5 := []string{"", "a", "", ("" + "\n\t")}
    v6 := GenS1{F0: "", F1: 65535}
    fmt.Println(v4, v1)
    if true {
        v4 += v4
        switch 9223372036854775807 {
        case -2147483648:
            if (128 > 147) {
                var v7 uint = uint(v1)
                _ = v7
            }
        case 127:
            for fi8 := 0; fi8 < 16; fi8++ {
                v3[fi8] = 0
                _ = fi8
            }
        default:
            var v9 string = "hello"
            _ = v9
        }
        fmt.Println(v1)
    } else {
        var v10 int8 = 1
        _ = len(v5)
        if true {
            var v11 string = "a"
            v12 := GenS2{F0: 32767, F1: 127}
            v13 := GenS3{F0: 2}
            _ = v11
            _ = v12
            _ = v13
        } else {
            var v14 uint64 = 0
            for fi15 := 0; fi15 < 18; fi15++ {
                v1 = v1
                v16 := GenS1{F0: "", F1: 59}
                if (fi15 != (fi15 % (((fi15 / ((fi15) | 1))) | 1))) {
                    v17 := GenS4{F0: (^(v14)), F1: 673, F2: 42}
                    v18 := GenS4{F0: 934, F1: 1000, F2: 474}
                    _ = v17
                    _ = v18
                }
                _ = fi15
                _ = v16
            }
            _ = v14
        }
        _ = v10
    }
    fmt.Println(v1, v4)
    for fi19 := 0; fi19 < 14; fi19++ {
        v20 := GenS2{F0: 32767, F1: 42}
        var v21 uint = 255
        _ = fi19
        _ = v20
        _ = v21
    }
    if true {
        delete(v3, 128)
        var v22 uint8 = (^(v1))
        var v23 uint32 = 1
        _ = v22
        _ = v23
    } else {
        var v24 int = 32767
        var v25 bool = true
        _ = v24
        _ = v25
    }
    if (0 >= 0) {
        v26 := GenS5{F0: 42, F1: (!(("hello" + "") != "hello"))}
        if ((^(v1)) < v1) {
            for fi27 := 0; fi27 < 13; fi27++ {
                v28 := v3[(fi27 >> ((315) & 63))]
                _ = fi27
                _ = v28
            }
        } else {
            v4 = v4
            v1 -= v1
        }
        _ = v26
    } else {
        var v29 uint32 = 0
        _ = v29
    }
    _ = len(v5)
    if false {
        v30 := GenS1{F0: "hello", F1: 2}
        _ = v30
    }
    if false {
        for fi31 := 0; fi31 < 10; fi31++ {
            var v32 uint = 2
            _ = fi31
            _ = v32
        }
        v4 = v4
        var v33 float32 = 3.14159
        _ = v33
    } else {
        var v34 uint = 1
        var v35 uint32 = 785
        var v36 int32 = -2147483648
        _ = v34
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 20; fi37++ {
        var v38 int = (fi37 + fi37)
        _ = fi37
        _ = v38
    }
    v39 := GenS2{F0: 1000, F1: 32768}
    v1 = (v1 << ((4294967295) & 63))
    fmt.Println(v1, v1)
    var v40 int8 = -40
    v5 = append(v5, ("hello" + "\n\t"))
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
