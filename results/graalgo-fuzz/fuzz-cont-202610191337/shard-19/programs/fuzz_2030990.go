package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 string; F1 bool; F2 int32 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 bool; F2 uint32 }
func genFn0(p0 int16, p1 uint32) int64 {
    p0 = p0
    var v7 uint32 = (p1 % (((p1 * p1)) | 1))
    var v8 int = 4294967295
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 uint32, p1 int16) int32 {
    p1 = -1
    v9 := GenS1{F0: (p1 * -770), F1: 110}
    _ = p0
    _ = p1
    _ = v9
    return -2147483648
}

func main() {
    var v1 uint32 = 4294967295
    v2 := []uint32{2147483647, 42, (v1 << ((602) & 63))}
    var v3 int64 = -9223372036854775808
    var v4 float64 = (3.14159 / (1.5 * -2.25))
    v5 := []uint32{}
    v6 := []string{"a", "a", "\n\t", "a"}
    for fi10 := 0; fi10 < 20; fi10++ {
        if true {
            v6 = append(v6, "a")
        } else {
            for fi11 := 0; fi11 < 3; fi11++ {
                var v12 int8 = -69
                _ = fi11
                _ = v12
            }
            v4 = v4
        }
        var v13 uint32 = (v1 >> ((4294967295) & 63))
        _ = fi10
        _ = v13
    }
    v4 += v4
    var v14 uint16 = 1
    if len(v6) > 0 {
        v6[0] = (("foo bar" + "a") + "a")
    }
    if len(v5) > 0 {
        v5[0] = 0
    }
    var v15 int = 32768
    v1 += (v1 | 467)
    switch v15 {
    case -266:
        v16 := GenS1{F0: -1, F1: 255}
        _ = v16
    case 482:
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 bool = true
            var v19 uint = 255
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    switch (v15 % (((v15 >> ((0) & 63))) | 1)) {
    case 1:
        fmt.Println(v4)
    default:
        v20 := GenS2{F0: ("\n\t" + (("foo bar" + "a") + ((("" + "\n\t") + ("\n\t" + "")) + "\n\t"))), F1: false, F2: 0}
        _ = v20
    }
    fmt.Println(v3)
    v21 := GenS3{F0: 4294967295, F1: 0}
    for fi22 := 0; fi22 < 4; fi22++ {
        fmt.Println(v14, v4, v14)
        for fi23 := 0; fi23 < 14; fi23++ {
            v24 := GenS4{F0: v3}
            fmt.Println(v1, v4)
            if (380 < 1) {
                v5 = append(v5, v1)
                if true {
                    v25 := GenS5{F0: -106, F1: (1 >= (v1 + 2147483648)), F2: 987}
                    _ = v25
                } else {
                    fmt.Println(fi22, v1, fi22)
                }
            }
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    var v26 uint = 42
    var v27 uint8 = 1
    if len(v2) > 0 {
        v2[0] = 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
