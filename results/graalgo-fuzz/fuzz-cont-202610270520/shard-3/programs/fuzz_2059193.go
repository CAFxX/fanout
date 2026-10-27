package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS2 struct { F0 int16; F1 int32; F2 int32 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint32, p1 int16) int16 {
    for fi6 := 0; fi6 < 7; fi6++ {
        switch (fi6 ^ fi6) {
        case 42:
            p1 = p1
        }
        _ = fi6
    }
    v7 := GenS1{F0: 758, F1: 0, F2: 339}
    if (2147483647 < 32768) {
        if (368 == 35) {
            if false {
                v8 := GenS1{F0: 796, F1: 127, F2: 2}
                for fi9 := 0; fi9 < 4; fi9++ {
                    switch (fi9 % (((fi9 ^ (^(fi9)))) | 1)) {
                    case 1:
                        switch ("\n\t" + ("" + (("" + ("\n\t" + "a")) + "a"))) {
                        case "\n\t":
                            v10 := GenS2{F0: (p1 | (-(p1))), F1: -2147483648, F2: 2147483647}
                            _ = v10
                        default:
                            v11 := []int{fi9, (fi9 ^ fi9), (fi9 >> ((255) & 63))}
                            _ = v11
                        }
                    case -1:
                        for fi12 := 0; fi12 < 15; fi12++ {
                            if (false || (float64(1.5) > (float64(1.5) / float64((3.14159 * float32((3.14159 - 1.5))))))) {
                                v13 := GenS1{F0: 20, F1: 42, F2: 0}
                                var v14 uint = 65535
                                v15 := GenS1{F0: 65535, F1: -100, F2: 968}
                                _ = v13
                                _ = v14
                                _ = v15
                            }
                            _ = fi12
                        }
                    }
                    p0 += (p0 | (p0 << ((191) & 63)))
                    _ = fi9
                }
                p1 = (p1 % (((p1 - -549)) | 1))
                _ = v8
            } else {
                p1 -= p1
            }
            var v16 bool = (p0 < 1)
            _ = v16
        } else {
            fmt.Println(p0)
        }
    }
    switch "foo bar" {
    case "":
        fmt.Println(p1)
    case "a":
        switch "foo bar" {
        case "a":
            var v17 int64 = 634
            _ = v17
        }
    default:
        v18 := GenS1{F0: 65535, F1: 121, F2: 181}
        _ = v18
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = v7
    return 0
}

func main() {
    var v1 uint32 = 744
    v2 := make(map[uint32]string)
    v3 := []uint32{(v1 & (v1 * v1)), 4294967295, 147, v1}
    v4 := []int{2147483647}
    var v5 uint64 = 4294967295
    var v19 int64 = 42
    var v20 uint8 = 108
    fmt.Println(v5, v5, v1)
    v21 := GenS2{F0: -707, F1: 0, F2: -2147483648}
    var v22 int = 2147483647
    fmt.Println(v20, v1, v19)
    v23 := GenS3{F0: v1}
    switch v22 {
    case 149:
        fmt.Println(v5, v5, v20)
    case 54:
        switch "a" {
        case "foo bar":
            v20 = v20
        case "hello":
            for fi24 := 0; fi24 < 15; fi24++ {
                var v25 uint16 = 255
                _ = fi24
                _ = v25
            }
        case "\n\t":
            var v26 int64 = v19
            _ = v26
        }
    case -255:
        if len(v3) > 0 {
            v3[0] = v1
        }
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        v5 += v5
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
