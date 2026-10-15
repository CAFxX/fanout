package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS3 struct { F0 int }
func genFn0(p0 int64, p1 string) int32 {
    var v7 uint64 = 698
    fmt.Println(v7, v7)
    var v8 uint8 = 0
    fmt.Println(p0, v8)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    v1 := []string{"foo bar"}
    var v2 float32 = (3.14159 - 0.1)
    var v3 float32 = (((v2 + (v2 - (v2 * v2))) / 0.1) * (float32(1.5) - 1.5))
    var v4 int8 = -128
    v5 := make(map[uint32]uint32)
    var v6 uint32 = 65535
    var v9 uint = 4294967295
    for fi10 := 0; fi10 < 2; fi10++ {
        for fi11 := 0; fi11 < 4; fi11++ {
            fmt.Println(fi11, fi10, fi11)
            var v12 string = "foo bar"
            _ = fi11
            _ = v12
        }
        var v13 uint32 = (v6 / (((v6 | 4294967295)) | 1))
        _ = fi10
        _ = v13
    }
    var v14 uint16 = 1
    if false {
        v1 = append(v1, "foo bar")
        if (904 != 1) {
            v1 = append(v1, ("" + "\n\t"))
            delete(v5, uint32(v14))
        } else {
            v9 -= v9
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 bool = true
                fmt.Println(fi15)
                _ = fi15
                _ = v16
            }
        }
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        v9 -= v9
        for fi18 := 0; fi18 < 12; fi18++ {
            v19 := GenS1{F0: (-(v4)), F1: 2}
            _ = fi18
            _ = v19
        }
        var v20 int16 = 1
        _ = fi17
        _ = v20
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        for fi22 := 0; fi22 < 20; fi22++ {
            v23 := v5[450]
            v24 := GenS2{F0: 2, F1: 2, F2: 1000}
            _ = fi22
            _ = v23
            _ = v24
        }
        for fi25 := 0; fi25 < 12; fi25++ {
            v9 -= uint(v14)
            if (0 <= genFn0(1, "")) {
                var v26 float64 = (3.14159 + 1e10)
                _ = v26
            }
            _ = fi25
        }
        if (v4 <= v4) {
            for fi27 := 0; fi27 < 15; fi27++ {
                var v28 uint = 763
                var v29 uint64 = uint64(v9)
                v30 := GenS1{F0: (v4 - v4), F1: v29}
                var v31 uint8 = uint8(v14)
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
                _ = v31
            }
        }
        _ = fi21
    }
    if (!false) {
        v32 := GenS3{F0: -9223372036854775808}
        var v33 float64 = 3.14159
        _ = v32
        _ = v33
    }
    v34 := v5[v6]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
