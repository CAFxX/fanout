package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 bool }
func genFn0(p0 uint16) int8 {
    var v7 bool = false
    p0 += p0
    for fi8 := 0; fi8 < 12; fi8++ {
        switch "\n\t" {
        case "":
            for fi9 := 0; fi9 < 5; fi9++ {
                v10 := []int{-546}
                fmt.Println(fi9, p0)
                var v11 float64 = 0.1
                _ = fi9
                _ = v10
                _ = v11
            }
        case "foo bar":
            var v12 uint64 = 539
            _ = v12
        case "\n\t":
            var v13 uint32 = 4294967295
            _ = v13
        }
        _ = fi8
    }
    _ = p0
    _ = v7
    return 0
}
func genFn1() uint64 {
    var v14 uint = 124
    v15 := []int{}
    var v16 int16 = 465
    v14 -= (v14 << (((^(v14))) & 63))
    _ = v14
    _ = v15
    _ = v16
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 4294967295
    v3 := make(map[string]string)
    v4 := make(map[string]bool)
    v5 := make(map[uint32]int)
    var v6 float32 = float32(float64(1e10))
    for fi17 := 0; fi17 < 6; fi17++ {
        switch "\n\t" {
        case "foo bar":
            var v18 bool = false
            _ = v18
        case "hello":
            for fi19 := 0; fi19 < 17; fi19++ {
                var v20 int16 = 701
                var v21 int32 = 1
                if true {
                    v21 = v21
                }
                _ = fi19
                _ = v20
                _ = v21
            }
        case "a":
            v22 := GenS1{F0: -32768, F1: 1, F2: false}
            _ = v22
        }
        _ = fi17
    }
    var v23 uint64 = 751
    for fi24 := 0; fi24 < 6; fi24++ {
        if true {
            var v25 uint = 42
            switch "a" {
            case "hello":
                for fi26 := 0; fi26 < 8; fi26++ {
                    v27 := GenS1{F0: 371, F1: 70, F2: false}
                    var v28 uint16 = 465
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
            _ = v25
        } else {
            var v29 int32 = 2147483647
            _ = v29
        }
        _ = fi24
    }
    delete(v1, 128)
    v2 -= v23
    fmt.Println(v23)
    fmt.Println(v2)
    var v30 int16 = -847
    switch 32767 {
    case 0:
        v2 += genFn1()
    case -344:
        switch 9223372036854775807 {
        case 32768:
            for fi31 := 0; fi31 < 15; fi31++ {
                v23 += (^(v23))
                _ = fi31
            }
        }
    }
    var v32 int64 = 0
    v2 += v23
    var v33 int16 = (v30 / ((v30) | 1))
    fmt.Println(v32, v33)
    v5[4294967295] = 284
    if true {
        if (!((-128 <= 1) || true)) {
            var v34 int32 = 1
            _ = v34
        }
        var v35 uint = 83
        _ = v35
    }
    v33 -= (v33 | 0)
    v36 := []int{127, -907, -19}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
