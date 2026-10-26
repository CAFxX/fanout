package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 int64; F2 int64 }
func genFn0() int {
    if (2 != 0) {
        v7 := []int{-157, -2147483648, -469}
        var v8 int = -394
        switch (v8 | 32767) {
        case 127:
            var v9 int32 = -621
            _ = v9
        case 42:
            if true {
                var v10 int32 = 2147483647
                _ = v10
            } else {
                v11 := GenS1{F0: -128}
                if len(v7) > 0 {
                    v7[0] = v8
                }
                _ = v11
            }
        case 128:
            v12 := GenS1{F0: 1}
            _ = v12
        default:
            switch (v8 * (v8 - v8)) {
            case -2147483648:
                var v13 int32 = 42
                _ = v13
            case 4294967295:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            default:
                var v14 uint16 = 42
                _ = v14
            }
        }
        _ = v7
        _ = v8
    }
    var v15 uint32 = 65535
    _ = v15
    return 944
}
func genFn1(p0 int16, p1 uint16) int64 {
    switch 32768 {
    case -2147483648:
        v16 := GenS1{F0: 100}
        _ = v16
    case 32767:
        for fi17 := 0; fi17 < 10; fi17++ {
            for fi18 := 0; fi18 < 11; fi18++ {
                v19 := []int{}
                if len(v19) > 0 {
                    v19[0] = (fi17 % (((^(fi18))) | 1))
                }
                var v20 uint64 = 18446744073709551615
                _ = fi18
                _ = v19
                _ = v20
            }
            for fi21 := 0; fi21 < 10; fi21++ {
                var v22 float32 = 1e10
                v23 := GenS1{F0: 100}
                var v24 int16 = p0
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            var v25 uint = 2
            _ = fi17
            _ = v25
        }
    }
    v26 := []int{-9223372036854775808, 32768}
    _ = len(v26)
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    _ = p0
    _ = p1
    _ = v26
    return -358
}

func main() {
    var v1 int16 = 0
    var v2 float32 = -2.25
    var v3 bool = false
    v4 := make(map[string]uint32)
    v5 := []int{2147483647, 128}
    var v6 uint64 = 982
    v27 := GenS2{F0: -733}
    var v28 uint16 = 42
    var v29 uint16 = 603
    if true {
        var v30 int8 = 0
        _ = v30
    }
    var v31 int16 = v1
    var v32 int64 = genFn1(-627, (^(v28)))
    var v33 uint32 = 65535
    v4[""] = 59
    v34 := GenS2{F0: int(v6)}
    if (!(8 > 127)) {
        for fi35 := 0; fi35 < 4; fi35++ {
            v36 := GenS3{F0: (fi35 / ((int(v29)) | 1)), F1: -9223372036854775808, F2: 42}
            _ = fi35
            _ = v36
        }
        v37 := GenS1{F0: 100}
        var v38 uint64 = (v6 - (^(v6)))
        _ = v37
        _ = v38
    }
    if (false || false) {
        v6 += (v6 ^ 18446744073709551615)
    } else {
        v3 = true
        var v39 uint32 = v33
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
