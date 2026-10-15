package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int32, p1 string) uint8 {
    switch 699 {
    case 4294967295:
        for fi4 := 0; fi4 < 19; fi4++ {
            var v5 int8 = 1
            switch 0 {
            case 0:
                v6 := []int{fi4, -1, fi4, fi4}
                _ = v6
            case 32767:
                for fi7 := 0; fi7 < 3; fi7++ {
                    var v8 int16 = 153
                    var v9 float32 = ((float32(1e10) - -2.25) * float32(0.0))
                    var v10 uint64 = 0
                    _ = fi7
                    _ = v10
                    _ = v8
                    _ = v9
                }
            case -621:
                var v11 string = (p1 + p1)
                _ = v11
            }
            v12 := []int{int(p0), 32767, -9223372036854775808}
            _ = fi4
            _ = v12
            _ = v5
        }
    case -756:
        p1 = (p1 + (p1 + p1))
    case -721:
        if false {
            p0 = -284
            p1 = p1
        }
    default:
        v13 := GenS1{F0: 111, F1: 127}
        _ = v13
    }
    switch "hello" {
    case "":
        for fi14 := 0; fi14 < 7; fi14++ {
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 int32 = (p0 + -1)
                v16 += 0
                var v17 int8 = 127
                _ = fi15
                _ = v16
                _ = v17
            }
            var v18 float64 = 1.5
            _ = fi14
            _ = v18
        }
    }
    var v19 int8 = 42
    var v20 string = (p1 + p1)
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return 2
}

func main() {
    v1 := []uint32{0, 18, 16, 94}
    v2 := []int{32767, 1, 1}
    var v3 uint64 = 4294967295
    v2 = append(v2, 128)
    v21 := GenS1{F0: 32767, F1: 0}
    if ((-2147483648 >= 664) || (!false)) {
        v22 := GenS1{F0: 1, F1: 171}
        _ = v22
    }
    v3 = v3
    for fi23 := 0; fi23 < 9; fi23++ {
        var v24 uint = 42
        _ = fi23
        _ = v24
    }
    _ = len(v1)
    v1 = append(v1, 42)
    if true {
        var v25 uint32 = 42
        if len(v2) > 0 {
            v2[0] = -1
        }
        v26 := GenS1{F0: 32767, F1: 141}
        _ = v25
        _ = v26
    }
    v3 = (v3 ^ v3)
    for fi27 := 0; fi27 < 15; fi27++ {
        var v28 int8 = -120
        _ = fi27
        _ = v28
    }
    switch "a" {
    case "foo bar":
        for fi29 := 0; fi29 < 7; fi29++ {
            v3 += (^(v3))
            v30 := GenS2{F0: 2, F1: 175}
            _ = fi29
            _ = v30
        }
    case "":
        switch ("" + "foo bar") {
        case "\n\t":
            v31 := GenS3{F0: 127}
            _ = v31
        }
    default:
        v3 = 4294967295
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
}
