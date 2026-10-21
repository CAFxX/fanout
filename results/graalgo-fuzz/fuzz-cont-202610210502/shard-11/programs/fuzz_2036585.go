package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int8, p1 int, p2 uint64) int32 {
    v5 := []int{}
    v6 := GenS1{F0: 128}
    if (42 == 835) {
        var v7 uint16 = 948
        var v8 uint = 963
        var v9 int8 = (-(p0))
        _ = v7
        _ = v8
        _ = v9
    }
    v5 = append(v5, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint64) int16 {
    fmt.Println(p0)
    p0 = 1
    for fi10 := 0; fi10 < 19; fi10++ {
        if ((fi10 | fi10) > (fi10 & 1)) {
            switch (fi10 << ((153) & 63)) {
            case -704:
                v11 := []int{(fi10 ^ 32767), -2147483648, fi10, (-(fi10))}
                _ = v11
            case -609:
                var v12 uint16 = 42
                _ = v12
            case 32768:
                for fi13 := 0; fi13 < 16; fi13++ {
                    var v14 int64 = 9223372036854775807
                    _ = fi13
                    _ = v14
                }
            default:
                switch "\n\t" {
                case "\n\t":
                    v15 := []int{42, fi10, fi10}
                    _ = v15
                default:
                    var v16 uint8 = uint8(p0)
                    _ = v16
                }
            }
        }
        for fi17 := 0; fi17 < 2; fi17++ {
            switch (fi17 & (fi10 - -120)) {
            case 127:
                p0 += (p0 - 18446744073709551615)
            case -2147483648:
                fmt.Println(fi17)
            case -476:
                fmt.Println(p0, fi10, fi10)
            }
            var v18 uint32 = 0
            switch fi10 {
            case 127:
                v19 := []int{127, (-(fi17)), fi10}
                _ = v19
            case -1:
                p0 -= (p0 << ((42) & 63))
            case -2147483648:
                if (p0 >= p0) {
                    fmt.Println(v18, p0)
                }
            default:
                var v20 float64 = -2.25
                _ = v20
            }
            _ = fi17
            _ = v18
        }
        _ = fi10
    }
    var v21 uint8 = 1
    v21 -= 127
    _ = p0
    _ = v21
    return 0
}

func main() {
    v1 := []uint32{2147483647, 628, 837}
    var v2 uint64 = 0
    var v3 string = "hello"
    var v4 string = v3
    var v22 int = 145
    var v23 int64 = -9223372036854775808
    v24 := GenS1{F0: 42}
    var v25 float64 = 1.5
    var v26 uint8 = 1
    var v27 uint = 1
    v1 = append(v1, 2)
    v28 := GenS2{F0: 42}
    v2 += (v2 ^ v2)
    fmt.Println(v3, v25)
    v29 := GenS2{F0: 777}
    var v30 int = v22
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
