package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 int32; F2 uint8 }

func main() {
    var v1 float64 = (1.5 / (1e10 + ((0.0 + 0.0) / 1e10)))
    var v2 float64 = (v1 - 0.1)
    var v3 uint8 = 1
    var v4 int8 = -128
    var v5 float64 = v2
    if false {
        fmt.Println(v3, v3)
        var v6 int = int(v4)
        _ = v6
    }
    v7 := []int{}
    if ((42 != 1) || (!true)) {
        switch 128 {
        case -672:
            _ = len(v7)
        case 127:
            switch 4294967295 {
            case 32767:
                if len(v7) > 0 {
                    v7[0] = 32767
                }
            case -2147483648:
                v8 := GenS1{F0: 2, F1: v4, F2: "hello"}
                _ = v8
            case 1:
                if (true || true) {
                    var v9 float32 = (0.0 - (((3.14159 * 1e10) + 1e10) + (1e10 + ((0.0 + 1e10) + 1.5))))
                    v10 := GenS2{F0: 2147483648}
                    _ = v10
                    _ = v9
                } else {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                }
            }
        }
        v2 -= v1
    } else {
        if true {
            switch 32767 {
            case 433:
                v1 -= v2
            default:
                var v11 int64 = -1
                _ = v11
            }
            fmt.Println(v3)
            var v12 int32 = -417
            _ = v12
        }
    }
    if (!false) {
        if true {
            var v13 uint = 358
            fmt.Println(v3, v1, v1)
            var v14 uint8 = v3
            _ = v13
            _ = v14
        } else {
            v2 -= float64(-2.25)
            var v15 uint64 = 676
            var v16 string = (("foo bar" + ("a" + (("a" + ("hello" + "a")) + "foo bar"))) + ("a" + "hello"))
            _ = v15
            _ = v16
        }
        v17 := GenS3{F0: (v2 > (v2 / ((v1 - v2) - v5))), F1: 42, F2: 2}
        _ = v17
    }
    v1 = ((v1 + v2) / 0.1)
    _ = len(v7)
    if false {
        switch -364 {
        case -9223372036854775808:
            v7 = append(v7, 32768)
        case -2:
            var v18 float64 = 0.0
            _ = v18
        case -12:
            for fi19 := 0; fi19 < 11; fi19++ {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                var v20 uint8 = (^(v3))
                v21 := GenS1{F0: 99, F1: -1, F2: "hello"}
                _ = fi19
                _ = v20
                _ = v21
            }
        }
    } else {
        var v22 uint64 = 209
        if len(v7) > 0 {
            v7[0] = 128
        }
        var v23 uint = 599
        _ = v22
        _ = v23
    }
    v4 = (^(v4))
    fmt.Println(v2, v3)
    if false {
        fmt.Println(v1, v3)
    } else {
        for fi24 := 0; fi24 < 4; fi24++ {
            var v25 int64 = 1
            _ = fi24
            _ = v25
        }
        if (false && (3.14159 > (1.5 / 1e10))) {
            var v26 string = "\n\t"
            if (42 == 255) {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
            _ = v26
        } else {
            _ = len(v7)
            var v27 int16 = 1
            v4 -= 100
            _ = v27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
