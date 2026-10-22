package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int16) int8 {
    if true {
        p0 = p0
    } else {
        var v5 int8 = -35
        v6 := GenS1{F0: 9223372036854775807}
        v5 += v5
        _ = v5
        _ = v6
    }
    fmt.Println(p0)
    var v7 int64 = 0
    _ = p0
    _ = v7
    return 42
}

func main() {
    var v1 uint8 = 128
    var v2 int8 = -85
    var v3 uint32 = 2147483648
    var v4 uint16 = 1
    var v8 int8 = (v2 | (v2 * genFn0(-138)))
    var v9 uint = 4294967295
    var v10 int8 = (v2 ^ v8)
    var v11 int16 = 561
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 int32 = 2147483647
        fmt.Println(v1, v13)
        switch -9223372036854775808 {
        case 32767:
            var v14 int32 = -856
            _ = v14
        case 4294967295:
            fmt.Println(v1, v8)
        case 42:
            if false {
                var v15 uint = (v9 ^ 2)
                v2 = (v2 ^ genFn0((^(v11))))
                fmt.Println(v4)
                _ = v15
            } else {
                if (uint32(v9) > (v3 % (((^(v3))) | 1))) {
                    if false {
                        for fi16 := 0; fi16 < 4; fi16++ {
                            v17 := GenS1{F0: -413}
                            v18 := []int{(fi16 * fi12), fi12}
                            fmt.Println(v10)
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    }
                }
            }
        default:
            v11 = int16(v13)
        }
        _ = fi12
        _ = v13
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        fmt.Println(v9, v3, v2)
        var v20 int64 = 1
        v20 = v20
        _ = fi19
        _ = v20
    }
    v8 -= genFn0((v11 + (v11 % (((v11 << (((v9 >> ((v9) & 63))) & 63))) | 1))))
    v21 := []int{660, 9223372036854775807, 127, 1}
    for fi22 := 0; fi22 < 8; fi22++ {
        var v23 int = -480
        _ = fi22
        _ = v23
    }
    var v24 uint64 = 194
    if false {
        fmt.Println(v1, v4, v11)
        if false {
            v25 := GenS1{F0: -9223372036854775808}
            _ = v25
        }
    } else {
        for fi26 := 0; fi26 < 17; fi26++ {
            for fi27 := 0; fi27 < 9; fi27++ {
                v28 := GenS2{F0: fi27}
                var v29 float32 = -2.25
                v4 = 42
                var v30 uint = v9
                var v31 uint8 = (v1 >> ((490) & 63))
                if (42 != -502) {
                    v32 := GenS3{F0: 2}
                    _ = v32
                }
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
                _ = v31
            }
            var v33 int16 = v11
            _ = fi26
            _ = v33
        }
        var v34 int64 = 446
        _ = v34
    }
    for fi35 := 0; fi35 < 14; fi35++ {
        fmt.Println(v2, v10, v9)
        v36 := GenS3{F0: v4}
        _ = fi35
        _ = v36
    }
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
