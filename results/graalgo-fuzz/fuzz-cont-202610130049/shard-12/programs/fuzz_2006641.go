package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint64) int16 {
    p0 += 1
    v4 := []int{-1, 2147483647, 2147483647}
    for fi5 := 0; fi5 < 15; fi5++ {
        fmt.Println(fi5, p0)
        var v6 float32 = 3.14159
        v7 := GenS1{F0: -1}
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS2{F0: 18446744073709551615, F1: 0}
    _ = p0
    _ = v4
    _ = v8
    return 0
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    var v3 int32 = -970
    delete(v1, "foo bar")
    switch "\n\t" {
    case "":
        switch "" {
        case "foo bar":
            for fi9 := 0; fi9 < 16; fi9++ {
                v3 -= 2147483647
                v3 += (v3 ^ v3)
                v10 := GenS3{F0: 168}
                _ = fi9
                _ = v10
            }
        }
    }
    fmt.Println(v3)
    var v11 uint8 = 42
    v12 := GenS2{F0: 2, F1: 65535}
    var v13 float64 = ((3.14159 / 1.5) / (-2.25 * 1e10))
    var v14 int = 0
    switch 0 {
    case 9223372036854775807:
        if ((327 != 255) && true) {
            switch v14 {
            case 42:
                switch (v14 | (-(v14))) {
                case -2147483648:
                    switch (v14 | (v14 | (-(v14)))) {
                    case 1:
                        var v15 int32 = -2147483648
                        _ = v15
                    case -705:
                        var v16 bool = false
                        _ = v16
                    }
                }
            case 32767:
                v17 := []int{2147483647, -756, v14}
                _ = v17
            case 128:
                v14 -= v14
            }
        } else {
            v11 += v11
            v18 := v2[(v14 ^ v14)]
            _ = v18
        }
    }
    var v19 uint64 = 0
    v19 = (v19 >> ((0) & 63))
    var v20 float32 = 1e10
    for fi21 := 0; fi21 < 10; fi21++ {
        if (((v20 - v20) + float32((float64((v20 + v20)) * v13))) != ((v20 / (v20 + (0.1 + v20))) + float32(v13))) {
            var v22 int8 = 0
            for fi23 := 0; fi23 < 8; fi23++ {
                if (false || (221 < 41)) {
                    v24 := []int{fi23}
                    _ = v24
                } else {
                    v25 := v1["a"]
                    v3 = v3
                    _ = v25
                }
                _ = fi23
            }
            _ = v22
        } else {
            v26 := []int{-500, v14, 0}
            _ = v26
        }
        v11 -= (v11 ^ 145)
        _ = fi21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
