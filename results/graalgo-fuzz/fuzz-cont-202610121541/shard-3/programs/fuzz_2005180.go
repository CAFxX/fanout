package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint64 }
func genFn0(p0 int, p1 uint64) uint32 {
    if (!(false && false)) {
        p1 = (p1 % (((p1 - p1)) | 1))
        if false {
            v5 := GenS1{F0: 0, F1: 716, F2: 42}
            switch "foo bar" {
            case "":
                p1 = 772
            case "a":
                var v6 int32 = -307
                _ = v6
            }
            switch "hello" {
            case "\n\t":
                if false {
                    if (false || (450 < 958)) {
                        fmt.Println(p0)
                        v7 := GenS1{F0: uint32(p1), F1: 581, F2: 42}
                        _ = v7
                    }
                    fmt.Println(p1)
                }
            case "foo bar":
                var v8 uint = 4294967295
                _ = v8
            case "a":
                if (3.14159 != 0.1) {
                    fmt.Println(p1)
                } else {
                    var v9 float64 = 3.14159
                    fmt.Println(p0, p0)
                    _ = v9
                }
            }
            _ = v5
        } else {
            if true {
                var v10 uint32 = 4294967295
                for fi11 := 0; fi11 < 20; fi11++ {
                    for fi12 := 0; fi12 < 11; fi12++ {
                        var v13 uint64 = p1
                        _ = fi12
                        _ = v13
                    }
                    if (178 >= 255) {
                        var v14 int32 = 794
                        switch (fi11 / ((fi11) | 1)) {
                        case -1:
                            fmt.Println(p1, fi11, fi11)
                        }
                        _ = v14
                    }
                    _ = fi11
                }
                var v15 uint8 = 128
                _ = v10
                _ = v15
            } else {
                if false {
                    v16 := []int{p0, (p0 >> ((60) & 63)), 42}
                    _ = v16
                }
                var v17 uint8 = 128
                var v18 int16 = 765
                _ = v17
                _ = v18
            }
            if true {
                v19 := GenS1{F0: 2147483647, F1: 42, F2: 255}
                var v20 int8 = -126
                fmt.Println(v20, p0)
                _ = v19
                _ = v20
            } else {
                v21 := GenS1{F0: 4294967295, F1: 32767, F2: 32767}
                v22 := GenS1{F0: 2147483648, F1: 32767, F2: 214}
                _ = v21
                _ = v22
            }
            for fi23 := 0; fi23 < 18; fi23++ {
                p1 += 18446744073709551615
                _ = fi23
            }
        }
        v24 := []int{(p0 >> ((42) & 63)), (p0 / ((127) | 1))}
        _ = v24
    } else {
        var v25 int16 = -32768
        switch "\n\t" {
        case "a":
            if true {
                switch "" {
                case "hello":
                    v25 = v25
                case "\n\t":
                    fmt.Println(p1, v25)
                }
            }
        }
        p1 = p1
        _ = v25
    }
    v26 := []int{579, 2147483647, p0}
    v27 := GenS1{F0: 1, F1: 1, F2: 0}
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    return 2
}
func genFn1(p0 int64, p1 int32) uint64 {
    v28 := []int{4294967295, 32768}
    fmt.Println(p0)
    if false {
        p1 = (^(p1))
    }
    _ = p0
    _ = p1
    _ = v28
    return 627
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    var v4 uint64 = 1
    var v29 uint = 625
    var v30 float32 = 0.1
    v4 -= v4
    var v31 uint16 = 2
    v32 := GenS1{F0: 2147483648, F1: 32767, F2: v31}
    v4 = genFn1(1, -2147483648)
    for fi33 := 0; fi33 < 1; fi33++ {
        v34 := GenS2{F0: genFn0(fi33, v4), F1: (v4 * (^(v4)))}
        for fi35 := 0; fi35 < 9; fi35++ {
            var v36 string = "foo bar"
            switch fi35 {
            case 1:
                v37 := GenS2{F0: 851, F1: v4}
                _ = v37
            default:
                v38 := []int{fi33, fi35, (fi35 / (((fi35 % ((fi35) | 1))) | 1)), (fi35 ^ int(v29))}
                _ = v38
            }
            if false {
                delete(v2, (fi35 % (((fi33 * fi33)) | 1)))
                var v39 float32 = (0.0 + v30)
                switch "foo bar" {
                case "a":
                    var v40 uint = (v29 + 293)
                    _ = v40
                case "hello":
                    fmt.Println(v31)
                case "\n\t":
                    if false {
                        var v41 int = int(v31)
                        if (-444 < 42) {
                            var v42 uint8 = 127
                            _ = v42
                        } else {
                            v43 := []int{(v41 | int(v4)), v41}
                            _ = v43
                        }
                        v44 := GenS2{F0: 616, F1: (v4 ^ genFn1(-355, -83))}
                        _ = v41
                        _ = v44
                    }
                }
                _ = v39
            }
            v45 := GenS1{F0: 42, F1: 1000, F2: v31}
            _ = fi35
            _ = v36
            _ = v45
        }
        v31 = 643
        _ = fi33
        _ = v34
    }
    v31 -= (v31 ^ v31)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
