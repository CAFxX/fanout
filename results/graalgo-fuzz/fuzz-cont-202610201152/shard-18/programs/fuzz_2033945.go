package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int64 }
func genFn0(p0 bool, p1 int32, p2 bool) int64 {
    for fi7 := 0; fi7 < 14; fi7++ {
        fmt.Println(p2, p0, p1)
        _ = fi7
    }
    fmt.Println(p1)
    v8 := GenS1{F0: -2147483648, F1: 42}
    var v9 float32 = (float32(1e10) * 3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := []uint32{2147483647, 2147483648, 4294967295, 0}
    var v2 uint64 = 78
    var v3 int16 = -311
    var v4 int8 = 116
    var v5 uint32 = 42
    var v6 int32 = 2147483647
    v2 -= 904
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        fmt.Println(v5, v2, v5)
    case "":
        fmt.Println(v4, v2, v4)
    }
    v10 := GenS1{F0: 9223372036854775807, F1: v2}
    var v11 uint16 = uint16(v5)
    switch 4294967295 {
    case 127:
        _ = len(v1)
    case 32767:
        v6 = -207
    case 128:
        v5 -= 1
    }
    v2 += (v2 + (^(v2)))
    if ("\n\t" != "") {
        var v12 uint8 = 0
        _ = v12
    } else {
        v5 = 2147483647
        if true {
            var v13 float32 = 3.14159
            if (2 == 5) {
                v14 := GenS2{F0: (v6 ^ v6), F1: genFn0(true, v6, false)}
                var v15 int = -1
                _ = v14
                _ = v15
            } else {
                v13 -= 3.14159
            }
            _ = v13
        } else {
            var v16 uint = 1
            var v17 uint16 = v11
            _ = v16
            _ = v17
        }
    }
    switch 530 {
    case 32767:
        for fi18 := 0; fi18 < 6; fi18++ {
            v19 := GenS1{F0: 0, F1: 0}
            for fi20 := 0; fi20 < 2; fi20++ {
                if false {
                    fmt.Println(v5, fi18)
                    v21 := GenS3{F0: 1}
                    _ = v21
                }
                _ = len(v1)
                var v22 int16 = (v3 / (((v3 * v3)) | 1))
                if len(v1) > 0 {
                    v1[0] = (v5 & (v5 + (v5 - 0)))
                }
                _ = len(v1)
                _ = fi20
                _ = v22
            }
            _ = fi18
            _ = v19
        }
    default:
        for fi23 := 0; fi23 < 5; fi23++ {
            fmt.Println(fi23)
            _ = fi23
        }
    }
    if true {
        switch 32768 {
        case 9223372036854775807:
            v24 := GenS2{F0: v6, F1: 1}
            _ = v24
        case 0:
            if false {
                v5 -= v5
                fmt.Println(v11, v2, v2)
                var v25 float32 = 1e10
                _ = v25
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        default:
            var v26 uint = 18446744073709551615
            _ = v26
        }
        fmt.Println(v3, v4, v2)
        switch 32768 {
        case -26:
            for fi27 := 0; fi27 < 11; fi27++ {
                v5 += v5
                for fi28 := 0; fi28 < 1; fi28++ {
                    v29 := GenS2{F0: v6, F1: genFn0(false, (v6 + v6), (!(false && (!((3.14159 > 0.0) || (4294967295 >= uint(v5)))))))}
                    var v30 int64 = 1
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi27
            }
        case 42:
            fmt.Println(v2)
        case -1:
            var v31 float32 = float32(0.1)
            _ = v31
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi32 := 0; fi32 < 2; fi32++ {
            var v33 uint = 515
            switch int(v33) {
            case 4294967295:
                v4 -= 127
            case -1:
                if len(v1) > 0 {
                    v1[0] = (v5 & v5)
                }
            }
            if (v11 >= (v11 & 65535)) {
                var v34 int32 = int32(v5)
                _ = v34
            }
            _ = fi32
            _ = v33
        }
    }
    var v35 uint64 = uint64(v11)
    var v36 int64 = 0
    v11 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
