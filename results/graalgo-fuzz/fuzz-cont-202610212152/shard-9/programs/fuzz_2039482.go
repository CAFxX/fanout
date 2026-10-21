package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int, p1 uint32, p2 int64) int16 {
    var v5 int32 = 52
    if (!true) {
        v6 := GenS1{F0: 42}
        _ = v6
    }
    v7 := []int{p0, p0, p0, p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return 1000
}

func main() {
    var v1 string = (("\n\t" + (("foo bar" + "a") + "hello")) + "")
    var v2 uint8 = 127
    var v3 int8 = 1
    var v4 int16 = 75
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 14; fi9++ {
            if true {
                v3 -= (v3 & 36)
                if false {
                    var v10 bool = (false && false)
                    var v11 uint16 = 255
                    _ = v10
                    _ = v11
                } else {
                    if false {
                        if false {
                            v12 := []int{}
                            v13 := GenS1{F0: 2}
                            _ = v12
                            _ = v13
                        }
                        var v14 int16 = (v4 % ((genFn0(fi8, 104, 1)) | 1))
                        switch fi8 {
                        case -40:
                            var v15 uint16 = 65535
                            _ = v15
                        case -776:
                            v14 = v4
                        case 430:
                            var v16 float32 = float32((0.1 - 0.0))
                            _ = v16
                        default:
                            var v17 int32 = 141
                            _ = v17
                        }
                        _ = v14
                    }
                    if (2 < 151) {
                        var v18 int64 = 42
                        if (v3 != v3) {
                            v3 = (^(v3))
                        }
                        _ = v18
                    } else {
                        switch 482 {
                        case -438:
                            var v19 float64 = 1e10
                            _ = v19
                        case 2147483647:
                            switch v1 {
                            case "a":
                                v3 -= -56
                            }
                        }
                        fmt.Println(fi9, v4, v3)
                        var v20 float64 = 1e10
                        _ = v20
                    }
                    v3 = (v3 / ((0) | 1))
                }
            }
            _ = fi9
        }
        if false {
            v2 = 0
            v4 = v4
        }
        v21 := []int{(fi8 >> ((42) & 63)), (fi8 << ((0) & 63)), fi8, (fi8 % ((365) | 1))}
        _ = fi8
        _ = v21
    }
    v22 := GenS1{F0: 0}
    v23 := GenS1{F0: 483}
    v24 := GenS1{F0: 0}
    v25 := GenS1{F0: 42}
    for fi26 := 0; fi26 < 18; fi26++ {
        switch (-(fi26)) {
        case 42:
            switch v1 {
            case "foo bar":
                fmt.Println(v3, fi26, v1)
            default:
                v3 -= (v3 * v3)
            }
        case 32767:
            var v27 float64 = float64(-2.25)
            _ = v27
        case 0:
            var v28 int8 = (v3 >> ((4294967295) & 63))
            _ = v28
        }
        v3 -= (v3 & v3)
        v1 = v1
        _ = fi26
    }
    var v29 uint16 = 408
    v30 := []int{2147483647, -9223372036854775808, 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
