package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 int32 }
func genFn0(p0 uint16, p1 int8) string {
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS1{F0: p1, F1: 65535, F2: 9223372036854775807}
        v8 := []int{-991, fi6, fi6, (-(fi6))}
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS2{F0: (-(fi9))}
        if true {
            switch fi9 {
            case 1:
                v11 := []int{(fi9 | -660), -797}
                _ = v11
            case 4294967295:
                for fi12 := 0; fi12 < 17; fi12++ {
                    var v13 string = ("a" + "\n\t")
                    _ = fi12
                    _ = v13
                }
            case -17:
                if (false || true) {
                    for fi14 := 0; fi14 < 4; fi14++ {
                        fmt.Println(p0, fi9, fi14)
                        _ = fi14
                    }
                }
            }
            v15 := GenS2{F0: 1}
            _ = v15
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return "\n\t"
}
func genFn1() int16 {
    var v16 float64 = 1.5
    v17 := GenS3{F0: -2147483648}
    _ = v16
    _ = v17
    return -982
}

func main() {
    var v1 uint16 = 2
    v2 := []int{0, 4294967295}
    var v3 string = "a"
    var v4 float64 = 1e10
    v5 := []int{489, 32767, 0, 128}
    var v18 int32 = 2147483647
    if len(v5) > 0 {
        v5[0] = 127
    }
    var v19 int64 = 0
    for fi20 := 0; fi20 < 9; fi20++ {
        v3 = v3
        v19 -= -1
        if false {
            var v21 uint16 = v1
            _ = v21
        }
        _ = fi20
    }
    if (2147483647 == uint32(v1)) {
        v22 := GenS4{F0: 65535, F1: v18}
        fmt.Println(v4)
        switch "a" {
        case "":
            var v23 float64 = float64(-2.25)
            _ = v23
        case "hello":
            v19 += 1
        }
        _ = v22
    }
    v24 := GenS1{F0: 2, F1: 274, F2: v19}
    v25 := GenS1{F0: 1, F1: 877, F2: v19}
    v3 = ""
    for fi26 := 0; fi26 < 15; fi26++ {
        for fi27 := 0; fi27 < 6; fi27++ {
            switch (fi27 + (fi27 >> ((42) & 63))) {
            case 127:
                v19 += v19
            }
            _ = fi27
        }
        _ = fi26
    }
    var v28 uint8 = 1
    _ = len(v2)
    var v29 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
