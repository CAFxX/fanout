package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 bool; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 int16, p1 uint, p2 int64) uint32 {
    switch 4294967295 {
    case 127:
        p1 -= p1
    case 224:
        var v4 int64 = p2
        _ = v4
    }
    fmt.Println(p2)
    v5 := GenS1{F0: (p1 / ((2) | 1))}
    var v6 float32 = 0.1
    var v7 uint8 = 84
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = (42 <= 127)
    v3 := []int{32768, 32767, 32767}
    var v8 float64 = 1e10
    switch 32767 {
    case -1:
        if false {
            v9 := GenS2{F0: 1, F1: -1}
            v10 := GenS2{F0: 90, F1: -9223372036854775808}
            var v11 bool = v2
            _ = v10
            _ = v11
            _ = v9
        } else {
            v12 := GenS3{F0: ((v2 && (!(4294967295 <= genFn0(-263, 0, 0)))) && v2), F1: (3.14159 >= float64(-2.25)), F2: 239}
            for fi13 := 0; fi13 < 8; fi13++ {
                v14 := GenS4{F0: 872, F1: 0}
                var v15 uint64 = 745
                v16 := GenS3{F0: v2, F1: v2, F2: 42}
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            v8 = 3.14159
            _ = v12
        }
    }
    var v17 int = 0
    v18 := GenS3{F0: ((3.14159 / 1e10) < 1e10), F1: (((-(v17)) == 2147483647) || ((42 == -942) && (!((^(v17)) < (v17 + -481))))), F2: 128}
    v19 := GenS2{F0: -128, F1: 274}
    v2 = v2
    switch "foo bar" {
    case "foo bar":
        for fi20 := 0; fi20 < 19; fi20++ {
            switch (^(v17)) {
            case 2147483647:
                delete(v1, 845)
            case 1:
                for fi21 := 0; fi21 < 5; fi21++ {
                    var v22 uint64 = 732
                    fmt.Println(v8, fi21, fi21)
                    var v23 int64 = -654
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            }
            _ = fi20
        }
    case "\n\t":
        v24 := GenS1{F0: 2}
        _ = v24
    }
    if v2 {
        v17 = (v17 % ((v17) | 1))
        v8 -= 3.14159
        switch "foo bar" {
        case "hello":
            v17 = (v17 >> ((0) & 63))
        }
    }
    if (!false) {
        var v25 int32 = 2147483647
        var v26 uint = 154
        _ = v25
        _ = v26
    } else {
        v17 += (^(v17))
        for fi27 := 0; fi27 < 17; fi27++ {
            var v28 uint = 380
            _ = fi27
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        if v2 {
            v30 := v1[42]
            _ = v30
        }
        fmt.Println(v8, v17, v8)
        _ = fi29
    }
    v8 = (3.14159 + v8)
    v31 := GenS2{F0: 127, F1: 42}
    var v32 uint64 = 0
    v1[genFn0(1, 127, 42)] = v2
    fmt.Println(v2, v8)
    switch v17 {
    case -1:
        v2 = false
    case -36:
        if true {
            v17 = 32767
        } else {
            v33 := GenS5{F0: 907}
            if v2 {
                if v2 {
                    var v34 uint = 670
                    _ = v34
                } else {
                    var v35 int = (v17 & (v17 & v17))
                    v36 := GenS4{F0: 42, F1: 127}
                    _ = v35
                    _ = v36
                }
            }
            _ = v33
        }
    default:
        var v37 int64 = 0
        _ = v37
    }
    fmt.Println(v2, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v8
}
