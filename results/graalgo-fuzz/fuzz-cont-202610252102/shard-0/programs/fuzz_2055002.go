package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint16; F1 int64; F2 int64 }
func genFn0(p0 string, p1 string) string {
    p0 = "hello"
    v6 := []int{128}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v6
    return p0
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = 127
    v3 := []uint32{43, 995, 88, 2}
    var v4 float32 = (0.0 - -2.25)
    v5 := make(map[uint32]uint32)
    var v7 bool = true
    fmt.Println(v4, v2)
    v8 := v1[42]
    switch v2 {
    case 32768:
        fmt.Println(v4, v4, v2)
    }
    var v9 uint = 18446744073709551615
    var v10 int32 = -1
    if v7 {
        v11 := GenS1{F0: v10, F1: -45}
        var v12 uint8 = 216
        v2 = (v2 / (((v2 - v2)) | 1))
        _ = v11
        _ = v12
    }
    var v13 uint32 = 2147483648
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 int64 = 0
        v3 = append(v3, (^(v13)))
        v7 = (v8 || (386 == v15))
        _ = fi14
        _ = v15
    }
    if ((v13 | v13) == 495) {
        v16 := v5[v13]
        _ = v16
    } else {
        v7 = true
    }
    var v17 uint16 = 580
    var v18 float64 = 3.14159
    if len(v3) > 0 {
        v3[0] = (v13 & 1)
    }
    if (v7 && true) {
        if v7 {
            var v19 uint16 = v17
            _ = v19
        }
        if ("\n\t" < ("\n\t" + "a")) {
            for fi20 := 0; fi20 < 2; fi20++ {
                switch (v2 - 4294967295) {
                case 32767:
                    v4 -= float32(v18)
                case -9223372036854775808:
                    var v21 uint16 = v17
                    _ = v21
                }
                var v22 string = "\n\t"
                _ = fi20
                _ = v22
            }
        } else {
            if (v8 && true) {
                for fi23 := 0; fi23 < 1; fi23++ {
                    v24 := GenS2{F0: (v13 / ((v13) | 1)), F1: 65535, F2: 1}
                    v25 := GenS3{F0: 32767, F1: 1, F2: 42}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            } else {
                v26 := GenS2{F0: (v13 & 65535), F1: v13, F2: 42}
                v8 = (v8 && (255 != (v9 >> (((^(v9))) & 63))))
                _ = v26
            }
            var v27 int32 = int32(v9)
            var v28 bool = ((v2 / ((-1) | 1)) == 42)
            _ = v27
            _ = v28
        }
        var v29 int = (v2 / ((128) | 1))
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
