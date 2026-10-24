package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8 }
func genFn0(p0 int64, p1 bool) int8 {
    var v4 int32 = 42
    v4 -= (v4 - (v4 * 784))
    v4 -= v4
    _ = p0
    _ = p1
    _ = v4
    return -128
}
func genFn1() uint32 {
    var v5 float32 = float32(1.5)
    var v6 int16 = 966
    _ = v5
    _ = v6
    return 65535
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[uint32]uint32)
    var v3 int16 = 0
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: 11}
        v9 := GenS2{F0: (fi7 >> ((18446744073709551615) & 63))}
        switch (fi7 + 0) {
        case 342:
            switch "hello" {
            case "\n\t":
                var v10 int64 = 1
                _ = v10
            }
        case -9223372036854775808:
            fmt.Println(fi7)
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    if (0 <= v3) {
        v3 = v3
        fmt.Println(v1)
        for fi11 := 0; fi11 < 18; fi11++ {
            v3 += (v3 + (v3 >> ((737) & 63)))
            fmt.Println(v3, v3)
            _ = fi11
        }
    }
    v12 := GenS2{F0: 32767}
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 uint16 = 255
        _ = fi13
        _ = v14
    }
    var v15 float32 = (0.1 * -2.25)
    var v16 float32 = 1.5
    for fi17 := 0; fi17 < 4; fi17++ {
        switch (fi17 % ((fi17) | 1)) {
        case -9223372036854775808:
            fmt.Println(v16, fi17)
        default:
            var v18 int = (fi17 + (fi17 & fi17))
            _ = v18
        }
        _ = len(v2)
        _ = fi17
    }
    var v19 int = -259
    for fi20 := 0; fi20 < 5; fi20++ {
        var v21 uint32 = 65535
        var v22 bool = ((!(980 <= 1)) && false)
        var v23 float64 = (1e10 * 3.14159)
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    v24 := []int{}
    v19 += (^(v19))
    fmt.Println(v3, v19)
    var v25 uint32 = v1
    for fi26 := 0; fi26 < 10; fi26++ {
        if true {
            switch "a" {
            case "hello":
                delete(v2, v25)
            default:
                if ((!(18446744073709551615 != 2)) && (45 < genFn0(874, (-126 >= 0)))) {
                    if true {
                        v19 -= (fi26 % (((-(v19))) | 1))
                        switch (-(fi26)) {
                        case 193:
                            var v27 float64 = 1e10
                            _ = v27
                        default:
                            v24 = append(v24, 127)
                        }
                        fmt.Println(v15)
                    }
                    for fi28 := 0; fi28 < 18; fi28++ {
                        v16 = v15
                        var v29 int64 = 328
                        fmt.Println(v1)
                        _ = fi28
                        _ = v29
                    }
                    var v30 bool = false
                    _ = v30
                }
            }
            fmt.Println(v3, v1, v15)
            v31 := GenS1{F0: 525}
            _ = v31
        } else {
            if true {
                var v32 int32 = -700
                var v33 uint32 = genFn1()
                _ = v32
                _ = v33
            }
        }
        fmt.Println(fi26)
        for fi34 := 0; fi34 < 9; fi34++ {
            switch (fi34 - 230) {
            case 2147483647:
                v35 := GenS3{F0: -128}
                _ = v35
            case 9223372036854775807:
                fmt.Println(v15, v16)
            }
            v3 = v3
            _ = fi34
        }
        _ = fi26
    }
    var v36 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v36
}
