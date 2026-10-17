package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 bool; F1 uint8; F2 uint64 }
func genFn0(p0 uint32, p1 uint8, p2 int) int64 {
    fmt.Println(p1, p0)
    p0 = p0
    p1 = (p1 % (((p1 >> ((uint(p1)) & 63))) | 1))
    _ = p0
    _ = p1
    _ = p2
    return -1
}
func genFn1(p0 int64, p1 string) int64 {
    p0 += (p0 | (p0 / ((p0) | 1)))
    var v5 uint16 = 256
    _ = p0
    _ = p1
    _ = v5
    return (p0 << ((42) & 63))
}

func main() {
    var v1 float32 = 1e10
    v2 := []int{128, 42, 626}
    var v3 int = -1
    var v4 float32 = (v1 / (float32(1.5) + float32(0.1)))
    switch (v3 + 127) {
    case -290:
        if true {
            for fi6 := 0; fi6 < 6; fi6++ {
                if len(v2) > 0 {
                    v2[0] = v3
                }
                _ = fi6
            }
            var v7 int32 = 1
            _ = v7
        } else {
            v1 += 1.5
        }
    }
    fmt.Println(v3, v1)
    v8 := GenS1{F0: false}
    _ = len(v2)
    var v9 int32 = 1
    for fi10 := 0; fi10 < 8; fi10++ {
        v1 -= 1e10
        if true {
            for fi11 := 0; fi11 < 13; fi11++ {
                var v12 uint8 = 137
                var v13 uint = 750
                if ((!false) && (!true)) {
                    v4 += v4
                }
                switch fi11 {
                case 128:
                    v14 := GenS1{F0: (v12 == v12)}
                    _ = v14
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                switch (-(v3)) {
                case 32768:
                    if ((!true) || (v12 <= 169)) {
                        v2 = append(v2, (fi11 & fi10))
                        switch (fi10 / ((-800) | 1)) {
                        case 842:
                            v15 := GenS1{F0: true}
                            _ = v15
                        }
                        if true {
                            _ = len(v2)
                        }
                    }
                case 127:
                    var v16 int16 = -1
                    _ = v16
                case -836:
                    v17 := GenS2{F0: v13, F1: 1}
                    _ = v17
                default:
                    fmt.Println(v9)
                }
                var v18 int64 = 42
                _ = fi11
                _ = v12
                _ = v13
                _ = v18
            }
        }
        switch (v3 % (((fi10 - 312)) | 1)) {
        case 42:
            if (true || (18446744073709551615 > 255)) {
                v19 := GenS2{F0: 255, F1: -9223372036854775808}
                _ = v19
            }
        case 128:
            if true {
                v1 = (v1 - (3.14159 - 0.1))
                var v20 float32 = float32(1.5)
                _ = v20
            }
        default:
            v9 -= v9
        }
        _ = fi10
    }
    v21 := GenS1{F0: (42 < 4294967295)}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 = v4
    var v22 uint8 = 1
    if true {
        switch "hello" {
        case "hello":
            var v23 uint8 = v22
            _ = v23
        case "foo bar":
            var v24 int32 = (v9 ^ (v9 & v9))
            _ = v24
        case "a":
            v25 := GenS3{F0: ((v9 | 942) >= -15), F1: v22, F2: 0}
            _ = v25
        }
        v9 = (-(v9))
    }
    v26 := GenS2{F0: 806, F1: 744}
    for fi27 := 0; fi27 < 7; fi27++ {
        if false {
            v9 = v9
            v4 -= ((((1.5 + v4) * v4) / -2.25) - v4)
            switch "foo bar" {
            case "foo bar":
                for fi28 := 0; fi28 < 17; fi28++ {
                    v3 += fi27
                    var v29 float64 = -2.25
                    _ = fi28
                    _ = v29
                }
            case "hello":
                v4 = (v4 / v1)
            case "\n\t":
                var v30 uint32 = 945
                _ = v30
            default:
                var v31 bool = (701 == 1)
                _ = v31
            }
        }
        if len(v2) > 0 {
            v2[0] = fi27
        }
        _ = fi27
    }
    fmt.Println(v3)
    v9 = (v9 & v9)
    v3 = (v3 | (v3 >> ((874) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
