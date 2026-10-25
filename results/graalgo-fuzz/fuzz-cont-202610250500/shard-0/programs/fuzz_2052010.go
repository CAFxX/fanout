package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 bool; F2 string }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 string; F1 int16 }

func main() {
    v1 := []string{}
    var v2 float64 = 0.1
    v3 := make(map[int]int)
    var v4 uint8 = 127
    v5 := make(map[string]uint32)
    v6 := []uint32{2147483648, 935, 2, 475}
    v7 := GenS1{F0: "hello", F1: 14, F2: 127}
    _ = len(v1)
    v8 := GenS2{F0: -1}
    if (-1 < 780) {
        var v9 float64 = float64((0.1 * 1.5))
        _ = len(v6)
        _ = v9
    }
    var v10 uint64 = 1
    _ = len(v1)
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 int16 = -1
        if true {
            var v13 int32 = 0
            var v14 uint8 = (v4 << ((uint(v4)) & 63))
            var v15 int = fi11
            _ = v13
            _ = v14
            _ = v15
        }
        var v16 int64 = 1
        _ = fi11
        _ = v12
        _ = v16
    }
    var v17 int64 = -1
    v18 := GenS3{F0: 1}
    v19 := GenS4{F0: v17, F1: (true && (515 > 967)), F2: "hello"}
    switch -156 {
    case 4294967295:
        v10 = (v10 >> ((18446744073709551615) & 63))
    }
    v20 := GenS3{F0: -52}
    switch -655 {
    case 32768:
        _ = len(v1)
    default:
        var v21 uint16 = 32767
        _ = v21
    }
    if (2 != (v10 & v10)) {
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 float32 = -2.25
            var v24 string = "a"
            _ = len(v6)
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = len(v3)
    } else {
        if (100 >= 0) {
            var v25 bool = (4294967295 > 65535)
            var v26 string = ("hello" + "hello")
            for fi27 := 0; fi27 < 8; fi27++ {
                var v28 string = v26
                _ = fi27
                _ = v28
            }
            _ = v25
            _ = v26
        }
    }
    var v29 string = "hello"
    if ((2147483648 == 42) || false) {
        if true {
            v17 = -1
            v2 += ((float64(((float32((-2.25 + (float64(-2.25) + ((v2 / v2) / -2.25)))) * 0.0) + (0.0 + (-2.25 * (0.1 + 0.0))))) / v2) / float64((0.0 - 0.1)))
        }
    }
    if true {
        v30 := GenS4{F0: v17, F1: false, F2: (v29 + (("hello" + "") + ("\n\t" + "a")))}
        var v31 uint32 = 211
        _ = v30
        _ = v31
    } else {
        if true {
            var v32 float32 = -2.25
            v10 = 393
            fmt.Println(v29, v29)
            _ = v32
        } else {
            fmt.Println(v2)
        }
        v33 := GenS5{F0: uint64(v4)}
        if false {
            for fi34 := 0; fi34 < 19; fi34++ {
                for fi35 := 0; fi35 < 20; fi35++ {
                    fmt.Println(v4)
                    delete(v5, (v29 + v29))
                    if false {
                        v17 = (-(v17))
                        v36 := GenS6{F0: "a", F1: int16(v4)}
                        var v37 int16 = 1000
                        _ = v36
                        _ = v37
                    }
                    _ = fi35
                }
                _ = fi34
            }
        }
        _ = v33
    }
    switch 938 {
    case 141:
        var v38 uint8 = 205
        _ = v38
    case 2147483647:
        if false {
            v4 = 176
            var v39 uint = 2
            var v40 float32 = (float32(-2.25) + 1e10)
            _ = v39
            _ = v40
        } else {
            var v41 int64 = (-(v17))
            v2 = (v2 * v2)
            _ = v41
        }
    case 42:
        delete(v3, 853)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
