package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
func genFn0(p0 bool, p1 int16) int16 {
    var v6 int = -1
    if (!(255 == 128)) {
        p1 = p1
    }
    v7 := []int{}
    var v8 uint = 604
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 1000
}
func genFn1(p0 int64, p1 int64, p2 uint8) int8 {
    p1 = (p0 ^ 731)
    v9 := GenS1{F0: (p0 | (p0 << ((4294967295) & 63))), F1: 32768}
    for fi10 := 0; fi10 < 15; fi10++ {
        p1 = 9223372036854775807
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 42
}

func main() {
    var v1 int64 = 1
    v2 := []int{4294967295, 661}
    var v3 int64 = -149
    var v4 uint16 = 255
    v5 := []string{"hello", "hello", ("\n\t" + "foo bar"), ""}
    for fi11 := 0; fi11 < 9; fi11++ {
        v4 = (v4 | 32767)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if (float32(0.0) >= (1.5 + 0.0)) {
            v12 := GenS1{F0: v3, F1: 0}
            switch fi11 {
            case 9223372036854775807:
                var v13 uint64 = 18446744073709551615
                _ = v13
            default:
                var v14 int32 = 42
                _ = v14
            }
            _ = v12
        }
        _ = fi11
    }
    v15 := GenS1{F0: (v1 ^ -1), F1: int(v4)}
    switch 1 {
    case 4294967295:
        v1 += (v3 + v3)
    case -249:
        var v16 bool = false
        _ = v16
    }
    if (!(3.14159 > 0.0)) {
        if false {
            for fi17 := 0; fi17 < 18; fi17++ {
                v4 -= (v4 + v4)
                v3 = (v3 & v3)
                var v18 int64 = 301
                _ = fi17
                _ = v18
            }
        }
        v19 := GenS1{F0: v3, F1: 42}
        _ = v19
    } else {
        var v20 int8 = genFn1((v1 % (((v3 - 9223372036854775807)) | 1)), (v3 << ((704) & 63)), 128)
        _ = v20
    }
    switch 2147483647 {
    case 356:
        _ = len(v5)
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        _ = len(v2)
        if false {
            v5 = append(v5, "a")
            var v22 uint = 4294967295
            v23 := GenS1{F0: 616, F1: (fi21 ^ fi21)}
            _ = v22
            _ = v23
        }
        _ = fi21
    }
    var v24 uint16 = (v4 << ((764) & 63))
    v25 := GenS1{F0: (-(v1)), F1: -668}
    fmt.Println(v24, v3)
    var v26 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
