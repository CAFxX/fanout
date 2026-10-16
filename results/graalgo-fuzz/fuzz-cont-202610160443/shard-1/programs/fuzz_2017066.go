package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int16; F1 int; F2 uint16 }
type GenS4 struct { F0 uint64; F1 int16 }
func genFn0(p0 int32, p1 int, p2 uint32) uint32 {
    for fi6 := 0; fi6 < 15; fi6++ {
        p1 = (fi6 & (fi6 << ((65535) & 63)))
        v7 := GenS1{F0: 215}
        if (1 <= -1) {
            p1 = fi6
            fmt.Println(p0)
        }
        _ = fi6
        _ = v7
    }
    p1 = p1
    if (42 < 1) {
        var v8 uint64 = 2
        var v9 uint16 = 2
        var v10 uint32 = (p2 >> ((895) & 63))
        _ = v10
        _ = v8
        _ = v9
    } else {
        for fi11 := 0; fi11 < 7; fi11++ {
            v12 := GenS1{F0: uint8(p2)}
            _ = fi11
            _ = v12
        }
        v13 := GenS1{F0: 1}
        if false {
            fmt.Println(p0)
        }
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}
func genFn1() int64 {
    var v14 string = ("" + "foo bar")
    fmt.Println(v14)
    _ = v14
    return 1
}

func main() {
    var v1 int32 = 1
    var v2 int16 = 247
    var v3 int32 = (-(v1))
    var v4 int16 = (v2 + v2)
    v5 := make(map[int]uint32)
    if (false && (4294967295 == -191)) {
        v15 := []int{4294967295, -2147483648}
        var v16 int8 = 100
        var v17 int64 = -221
        _ = v15
        _ = v16
        _ = v17
    }
    fmt.Println(v4, v1)
    switch -128 {
    case 0:
        if (true && false) {
            if true {
                v18 := GenS1{F0: 1}
                if true {
                    var v19 int8 = -1
                    _ = v19
                } else {
                    var v20 float32 = 1e10
                    _ = v20
                }
                _ = v18
            }
        } else {
            var v21 uint32 = 978
            _ = v21
        }
    }
    fmt.Println(v1)
    v22 := GenS2{F0: v1, F1: 103}
    v2 -= -595
    for fi23 := 0; fi23 < 10; fi23++ {
        switch 0 {
        case 1:
            var v24 bool = true
            _ = v24
        case 32767:
            for fi25 := 0; fi25 < 8; fi25++ {
                v26 := GenS2{F0: (v1 * v1), F1: 100}
                _ = fi25
                _ = v26
            }
        }
        switch (fi23 - fi23) {
        case -289:
            var v27 int64 = 9223372036854775807
            _ = v27
        }
        _ = fi23
    }
    v28 := GenS2{F0: v1, F1: 127}
    v1 = -1
    var v29 int32 = 42
    delete(v5, 32768)
    if false {
        v30 := GenS3{F0: -516, F1: 2147483647, F2: 65535}
        v31 := GenS2{F0: v29, F1: 127}
        _ = v30
        _ = v31
    } else {
        if true {
            for fi32 := 0; fi32 < 18; fi32++ {
                for fi33 := 0; fi33 < 17; fi33++ {
                    var v34 float32 = (1.5 + (3.14159 - float32(0.1)))
                    _ = fi33
                    _ = v34
                }
                fmt.Println(fi32, v29)
                v35 := GenS4{F0: 135, F1: v4}
                _ = fi32
                _ = v35
            }
            v1 = (v1 ^ (v3 + 42))
            v2 += (v2 - -32768)
        }
        var v36 int16 = (v2 + 1)
        var v37 float64 = 0.0
        _ = v36
        _ = v37
    }
    v38 := v5[2147483647]
    var v39 uint32 = (v38 | v38)
    fmt.Println(v1, v2)
    v40 := GenS2{F0: v1, F1: 100}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
