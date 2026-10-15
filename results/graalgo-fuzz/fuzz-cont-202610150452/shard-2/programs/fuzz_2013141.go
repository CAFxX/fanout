package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int32; F2 string }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 uint32 = 896
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]int)
    var v5 float32 = ((float32(0.0) + 0.1) * ((0.1 - (1.5 - 1e10)) / (1e10 / (((((-2.25 / 3.14159) * 1.5) * ((3.14159 + 0.0) - -2.25)) / 3.14159) + 3.14159))))
    if true {
        var v6 int32 = 859
        v1 += (v1 ^ v1)
        var v7 uint32 = (v1 >> ((0) & 63))
        _ = v6
        _ = v7
    }
    if false {
        switch -9223372036854775808 {
        case 127:
            if (1 > 124) {
                v8 := GenS1{F0: 1}
                v9 := GenS2{F0: 127, F1: -292, F2: ""}
                _ = v8
                _ = v9
            }
        case 222:
            for fi10 := 0; fi10 < 5; fi10++ {
                v11 := []int{fi10, 273, (fi10 % ((fi10) | 1)), -9223372036854775808}
                v1 = (v1 ^ v1)
                var v12 uint16 = 633
                _ = fi10
                _ = v11
                _ = v12
            }
        case 2147483647:
            if (false && (true && (true || (!true)))) {
                v1 = v1
                v5 -= 3.14159
                var v13 int16 = 1
                _ = v13
            } else {
                var v14 string = "foo bar"
                var v15 bool = true
                _ = len(v3)
                _ = v14
                _ = v15
            }
        }
        var v16 bool = (4294967295 == v1)
        for fi17 := 0; fi17 < 2; fi17++ {
            v5 = v5
            _ = fi17
        }
        _ = v16
    }
    v18 := GenS2{F0: 1, F1: -1, F2: "hello"}
    v19 := []int{42, 4294967295, -1, -680}
    v1 = v1
    var v20 uint16 = 0
    v19 = append(v19, -9223372036854775808)
    for fi21 := 0; fi21 < 6; fi21++ {
        v22 := v2[(v1 / ((607) | 1))]
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        for fi24 := 0; fi24 < 15; fi24++ {
            fmt.Println(fi23, v5)
            switch "a" {
            case "\n\t":
                fmt.Println(v20)
            }
            var v25 int64 = -1
            switch (fi24 << ((1) & 63)) {
            case -291:
                _ = len(v4)
            default:
                v26 := GenS3{F0: uint64(v20)}
                _ = v26
            }
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    if len(v19) > 0 {
        v19[0] = 9223372036854775807
    }
    var v27 float64 = -2.25
    var v28 int8 = int8(v1)
    v5 += v5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
