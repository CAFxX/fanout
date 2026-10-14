package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 bool; F2 bool }
func genFn0(p0 int64, p1 int) uint64 {
    p0 = (-(p0))
    v6 := GenS1{F0: 0, F1: 128}
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 float32 = 1e10
        p1 = 128
        _ = fi7
        _ = v8
    }
    var v9 string = ("hello" + "\n\t")
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v6
    _ = v9
    return 314
}

func main() {
    var v1 string = ("\n\t" + "")
    v2 := make(map[int]bool)
    var v3 int64 = -1
    var v4 string = v1
    var v5 uint8 = 128
    if (2 == 486) {
        if false {
            v1 = "foo bar"
        }
        for fi10 := 0; fi10 < 17; fi10++ {
            delete(v2, (fi10 & 32767))
            _ = fi10
        }
        v4 = v4
    }
    var v11 uint16 = 42
    if ((-2147483648 < -2147483648) || false) {
        fmt.Println(v1, v1, v5)
        if true {
            v3 -= v3
        }
        switch 739 {
        case -955:
            var v12 uint64 = 781
            _ = v12
        case 128:
            v1 = ("\n\t" + v1)
        }
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        for fi14 := 0; fi14 < 20; fi14++ {
            v4 = v4
            _ = fi14
        }
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 uint = 0
            _ = fi15
            _ = v16
        }
        var v17 int16 = 0
        _ = fi13
        _ = v17
    }
    delete(v2, -1)
    var v18 uint16 = v11
    var v19 uint8 = v5
    if (784 == 100) {
        fmt.Println(v18, v1)
        fmt.Println(v3, v4)
    } else {
        v20 := GenS2{F0: 96, F1: v3}
        _ = v20
    }
    v11 = 0
    var v21 uint64 = 4294967295
    v22 := GenS3{F0: 365}
    fmt.Println(v21, v3)
    var v23 uint8 = v19
    v24 := GenS3{F0: 2}
    switch -904 {
    case -9223372036854775808:
        v25 := GenS1{F0: uint64(v11), F1: (v23 + (v5 - (v5 << ((42) & 63))))}
        _ = v25
    case 9223372036854775807:
        v26 := GenS4{F0: (v3 % (((v3 - 42)) | 1)), F1: true, F2: false}
        _ = v26
    default:
        if (1 > -1) {
            var v27 int16 = 1000
            switch 128 {
            case 9223372036854775807:
                var v28 bool = true
                _ = v28
            case 475:
                var v29 uint = 655
                _ = v29
            case 253:
                switch 0 {
                case 127:
                    v19 -= v19
                default:
                    v30 := GenS1{F0: 1, F1: v23}
                    _ = v30
                }
            }
            var v31 uint = 255
            _ = v27
            _ = v31
        }
    }
    if false {
        v2[4294967295] = (!false)
    } else {
        v3 += v3
        var v32 uint64 = v21
        _ = v32
    }
    delete(v2, 0)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
