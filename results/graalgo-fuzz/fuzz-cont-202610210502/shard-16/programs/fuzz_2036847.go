package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int }
func genFn0(p0 uint8, p1 int) int32 {
    v7 := []int{(p1 * p1), (p1 ^ -1), (p1 >> ((18446744073709551615) & 63)), -851}
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 float32 = 0.1
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        if (false && (int(p0) >= (p1 + 126))) {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            switch (fi10 / ((p1) | 1)) {
            case -2147483648:
                fmt.Println(p0)
            case 32768:
                fmt.Println(fi10)
            }
        }
        for fi11 := 0; fi11 < 19; fi11++ {
            p1 -= 32768
            var v12 uint64 = 522
            _ = fi11
            _ = v12
        }
        var v13 uint64 = 290
        _ = fi10
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v7
    return 0
}
func genFn1(p0 int, p1 uint16) int16 {
    var v14 float32 = float32((0.1 * ((0.0 / float64(float32(3.14159))) + float64((3.14159 * ((0.1 + 1e10) * 1.5))))))
    v15 := GenS1{F0: -1, F1: p0}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return -1
}

func main() {
    var v1 uint8 = 202
    var v2 uint8 = 2
    v3 := make(map[string]string)
    var v4 int64 = 42
    var v5 uint = 566
    var v6 int8 = 42
    var v16 int32 = int32(v2)
    var v17 int64 = -9223372036854775808
    if ((v17 | (-(v17))) > (v4 << (((^(v5))) & 63))) {
        if true {
            var v18 uint = 687
            if (1e10 < 3.14159) {
                v19 := GenS2{F0: 128, F1: 100}
                _ = v19
            } else {
                _ = len(v3)
                fmt.Println(v17, v4)
            }
            fmt.Println(v4)
            _ = v18
        } else {
            v20 := []int{32767}
            for fi21 := 0; fi21 < 7; fi21++ {
                if true {
                    v6 = (-(v6))
                    v16 += genFn0((v1 & v1), fi21)
                }
                if false {
                    switch "" {
                    case "a":
                        v4 = int64(v5)
                    default:
                        delete(v3, "hello")
                    }
                } else {
                    v16 -= (v16 << ((v5) & 63))
                    v6 = v6
                    if ((v2 != (^(v1))) && true) {
                        _ = len(v20)
                        fmt.Println(v6, v17, v5)
                        if false {
                            v22 := GenS2{F0: fi21, F1: v6}
                            _ = v22
                        } else {
                            delete(v3, "\n\t")
                            v23 := GenS2{F0: fi21, F1: v6}
                            _ = v23
                        }
                    }
                }
                _ = fi21
            }
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            _ = v20
        }
    } else {
        v24 := GenS3{F0: 0, F1: uint(v1)}
        v25 := GenS2{F0: -2147483648, F1: v6}
        var v26 uint = 255
        _ = v24
        _ = v25
        _ = v26
    }
    v27 := GenS4{F0: v2}
    var v28 int8 = (-(v6))
    v16 = genFn0((v2 & (v1 ^ (v2 + v2))), -530)
    var v29 int64 = int64(v5)
    for fi30 := 0; fi30 < 20; fi30++ {
        v4 = v4
        _ = fi30
    }
    var v31 bool = (v4 > (v29 >> ((607) & 63)))
    var v32 bool = ((v1 & (^(v2))) > v1)
    if ((!(167 == 2)) || true) {
        if false {
            var v33 string = "\n\t"
            for fi34 := 0; fi34 < 15; fi34++ {
                _ = len(v3)
                v4 = int64(v28)
                _ = fi34
            }
            v28 = v28
            _ = v33
        }
        fmt.Println(v4, v28, v28)
        v35 := GenS5{F0: 9223372036854775807}
        _ = v35
    }
    var v36 float32 = -2.25
    var v37 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
