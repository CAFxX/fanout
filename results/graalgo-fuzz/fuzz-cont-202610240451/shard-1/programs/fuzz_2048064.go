package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 int; F1 uint32; F2 uint64 }
func genFn0(p0 int8, p1 uint64) int8 {
    var v6 int32 = 568
    for fi7 := 0; fi7 < 2; fi7++ {
        if true {
            if true {
                var v8 bool = (0 >= 32767)
                _ = v8
            }
        }
        _ = fi7
    }
    var v9 int32 = v6
    v10 := []int{-1}
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    _ = v9
    return (p0 | (^(p0)))
}
func genFn1() int64 {
    v11 := GenS1{F0: false, F1: 1, F2: 862}
    var v12 int32 = 884
    fmt.Println(v12)
    var v13 string = ""
    _ = v11
    _ = v12
    _ = v13
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 42
    v3 := make(map[string]uint32)
    v4 := make(map[int]uint32)
    var v5 int = -916
    var v14 uint64 = 698
    var v15 uint64 = (v14 + 113)
    var v16 float32 = (1e10 * 1.5)
    v17 := v4[(-(v5))]
    fmt.Println(v14, v17, v5)
    for fi18 := 0; fi18 < 3; fi18++ {
        if ((!((v16 * 1e10) != v16)) || (!false)) {
            v3[("hello" + "\n\t")] = v17
            v14 -= 1
        } else {
            v17 += (v17 / ((v17) | 1))
            v19 := []int{fi18}
            var v20 uint8 = 128
            _ = v19
            _ = v20
        }
        if true {
            v16 = 1e10
        } else {
            var v21 uint16 = 241
            v22 := []int{fi18, int(v17), 1}
            _ = v21
            _ = v22
        }
        _ = fi18
    }
    switch (v5 >> ((v2) & 63)) {
    case 32768:
        var v23 uint = (v2 / (((v2 | 255)) | 1))
        _ = v23
    case 2147483647:
        var v24 int = (^(v5))
        _ = v24
    case -2147483648:
        v14 = v14
    default:
        v25 := GenS2{F0: v15, F1: 1}
        _ = v25
    }
    v5 -= v5
    if false {
        if true {
            v17 -= 2147483648
        } else {
            if (uint16(v15) < 211) {
                var v26 uint = 255
                _ = v26
            }
        }
        if (4294967295 < v2) {
            v16 = 0.1
            if false {
                if (true && (!false)) {
                    switch (v5 & (v5 - 32767)) {
                    case 9223372036854775807:
                        fmt.Println(v16, v17)
                    case -9223372036854775808:
                        switch "hello" {
                        case "\n\t":
                            var v27 float64 = (0.0 / float64(v16))
                            _ = v27
                        case "a":
                            v17 += 2
                        }
                    }
                    v14 = v15
                    v4[(^(v5))] = v17
                }
                var v28 float32 = v16
                switch "hello" {
                case "foo bar":
                    var v29 int32 = 239
                    _ = v29
                case "\n\t":
                    fmt.Println(v28, v5)
                }
                _ = v28
            } else {
                v30 := GenS1{F0: (42 != 128), F1: 581, F2: v14}
                v31 := GenS1{F0: true, F1: 0, F2: v15}
                if true {
                    v14 += (v15 % ((v15) | 1))
                    v2 -= v2
                    var v32 int = v5
                    _ = v32
                } else {
                    if (42 != 1) {
                        switch -748 {
                        case 127:
                            v2 = v2
                        default:
                            v14 = 18446744073709551615
                        }
                        v33 := GenS3{F0: -603, F1: (^(v17)), F2: v15}
                        _ = v33
                    }
                }
                _ = v30
                _ = v31
            }
            v34 := []int{9223372036854775807}
            _ = v34
        } else {
            v5 = -2147483648
        }
    }
    var v35 string = "foo bar"
    for fi36 := 0; fi36 < 11; fi36++ {
        v16 = (-2.25 - v16)
        v2 = v2
        delete(v1, fi36)
        _ = fi36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
