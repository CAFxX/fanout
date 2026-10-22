package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 uint16; F2 uint32 }
func genFn0(p0 int16, p1 int32) int {
    if (-2.25 < (1e10 * 0.1)) {
        p0 += p0
        var v6 bool = false
        _ = v6
    }
    v7 := []int{4294967295, -867}
    _ = p0
    _ = p1
    _ = v7
    return -2147483648
}

func main() {
    var v1 float32 = 1e10
    var v2 string = "hello"
    var v3 int16 = 0
    var v4 string = ("hello" + ((v2 + v2) + v2))
    v5 := make(map[string]string)
    switch -525 {
    case -325:
        if ((65535 > 566) || true) {
            v4 = v4
            switch -1 {
            case 0:
                if ((-9223372036854775808 > -411) || (!(42 <= -1))) {
                    v8 := GenS1{F0: 1, F1: 42}
                    var v9 int8 = 0
                    v10 := []int{-612, 42, 342, genFn0(int16(v9), -515)}
                    _ = v10
                    _ = v8
                    _ = v9
                } else {
                    var v11 int32 = -2147483648
                    _ = v11
                }
            case 127:
                v12 := GenS2{F0: -2147483648, F1: v4}
                _ = v12
            case 894:
                v13 := GenS1{F0: 0, F1: 0}
                _ = v13
            default:
                var v14 uint32 = 2147483648
                _ = v14
            }
        }
    case -1:
        v2 = v4
    case 32767:
        v15 := []int{241, 128}
        _ = v15
    }
    v16 := GenS3{F0: 865}
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 float32 = v1
        _ = fi17
        _ = v18
    }
    var v19 string = "a"
    v4 = (v19 + "foo bar")
    var v20 int8 = -1
    switch 32767 {
    case 128:
        var v21 string = (v19 + "hello")
        _ = v21
    case 0:
        if true {
            var v22 int8 = (v20 * (v20 + 100))
            if true {
                v4 = v19
            }
            _ = v22
        }
    }
    v2 = v2
    switch -706 {
    case -958:
        for fi23 := 0; fi23 < 12; fi23++ {
            var v24 bool = (false && false)
            _ = fi23
            _ = v24
        }
    }
    if true {
        for fi25 := 0; fi25 < 3; fi25++ {
            v26 := []int{}
            var v27 float64 = (3.14159 - (-2.25 + 0.0))
            _ = fi25
            _ = v26
            _ = v27
        }
        v19 = "foo bar"
        for fi28 := 0; fi28 < 18; fi28++ {
            var v29 int64 = 494
            _ = fi28
            _ = v29
        }
    }
    v30 := GenS4{F0: 1, F1: 65535, F2: 317}
    var v31 string = v19
    var v32 uint64 = 23
    v20 = v20
    v4 = (v4 + v31)
    var v33 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
