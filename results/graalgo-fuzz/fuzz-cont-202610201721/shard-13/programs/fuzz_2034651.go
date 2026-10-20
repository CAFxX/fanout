package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 float32 = -2.25
    var v2 int64 = 1
    var v3 int = -536
    v4 := make(map[string]uint32)
    var v5 int8 = 0
    if true {
        v6 := GenS1{F0: -9223372036854775808}
        fmt.Println(v2, v5, v1)
        _ = v6
    }
    fmt.Println(v3, v2, v1)
    var v7 uint8 = 132
    if ("foo bar" > "") {
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 uint64 = 0
            var v10 float32 = 3.14159
            _ = fi8
            _ = v10
            _ = v9
        }
    } else {
        v3 = (-(v3))
        var v11 uint32 = 2147483647
        _ = v11
    }
    if (false || true) {
        fmt.Println(v2)
        var v12 uint = 2
        for fi13 := 0; fi13 < 7; fi13++ {
            for fi14 := 0; fi14 < 6; fi14++ {
                if true {
                    v1 = (1.5 * v1)
                    var v15 uint32 = 474
                    fmt.Println(v15)
                    _ = v15
                }
                var v16 uint16 = 42
                _ = fi14
                _ = v16
            }
            v3 -= (fi13 * 127)
            v17 := GenS2{F0: 1}
            _ = fi13
            _ = v17
        }
        _ = v12
    } else {
        v18 := []int{v3, (-(v3)), 127}
        _ = v18
    }
    v19 := []int{}
    for fi20 := 0; fi20 < 20; fi20++ {
        v3 -= (v3 & (fi20 & v3))
        v5 += v5
        _ = fi20
    }
    if true {
        v21 := v4["a"]
        fmt.Println(v3, v1)
        for fi22 := 0; fi22 < 4; fi22++ {
            if (v1 >= float32(1.5)) {
                v19 = append(v19, v3)
                v5 = (v5 | int8(v2))
            } else {
                v23 := GenS1{F0: v2}
                fmt.Println(fi22)
                _ = v23
            }
            v24 := GenS2{F0: (v5 | 127)}
            v25 := GenS1{F0: 370}
            _ = fi22
            _ = v24
            _ = v25
        }
        _ = v21
    } else {
        var v26 int8 = -128
        v5 = -38
        _ = v26
    }
    v27 := v4[("a" + "")]
    for fi28 := 0; fi28 < 17; fi28++ {
        fmt.Println(v2, fi28)
        for fi29 := 0; fi29 < 19; fi29++ {
            v30 := GenS2{F0: v5}
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    v31 := GenS2{F0: -66}
    fmt.Println(v27, v3)
    for fi32 := 0; fi32 < 15; fi32++ {
        var v33 uint64 = 540
        _ = len(v19)
        _ = fi32
        _ = v33
    }
    var v34 uint64 = uint64(v7)
    switch -473 {
    case 2147483647:
        for fi35 := 0; fi35 < 11; fi35++ {
            for fi36 := 0; fi36 < 10; fi36++ {
                v3 = int(v5)
                var v37 int8 = 100
                var v38 int32 = int32(v2)
                _ = fi36
                _ = v37
                _ = v38
            }
            v7 = (^(v7))
            _ = fi35
        }
    case -9223372036854775808:
        v39 := GenS3{F0: 755}
        _ = v39
    }
    var v40 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v40
    _ = v5
    _ = v7
}
