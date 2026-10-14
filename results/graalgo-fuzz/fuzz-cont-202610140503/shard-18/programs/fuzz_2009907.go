package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint16, p1 int) bool {
    v4 := GenS1{F0: 32767, F1: 2147483647}
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := []int{}
        var v7 string = ("" + "")
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 int = p1
            v9 = fi8
            var v10 int16 = 42
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    switch p1 {
    case 42:
        if true {
            v11 := []int{}
            p0 = 255
            _ = v11
        } else {
            if ((!(!(false || (410 != 264)))) || true) {
                for fi12 := 0; fi12 < 3; fi12++ {
                    if true {
                        v13 := GenS2{F0: -920, F1: 0, F2: 65535}
                        _ = v13
                    } else {
                        v14 := []int{}
                        _ = v14
                    }
                    _ = fi12
                }
                v15 := []int{p1, (^(p1))}
                _ = v15
            }
            fmt.Println(p0)
            p0 -= p0
        }
    case 170:
        var v16 int8 = -1
        _ = v16
    case 32767:
        var v17 uint8 = 0
        _ = v17
    }
    p1 += 4294967295
    _ = p0
    _ = p1
    _ = v4
    return true
}

func main() {
    var v1 int = 2147483647
    var v2 int32 = 42
    var v3 int = (v1 % (((-(v1))) | 1))
    var v18 uint = 255
    if (true && ((((1.5 / 0.1) / 1.0) - (1.5 / 1.5)) >= 3.14159)) {
        for fi19 := 0; fi19 < 6; fi19++ {
            v18 = 4294967295
            switch "a" {
            case "":
                var v20 bool = (v18 <= v18)
                _ = v20
            case "\n\t":
                var v21 bool = false
                _ = v21
            case "hello":
                v3 = (fi19 + (fi19 - fi19))
            default:
                var v22 int = (^(v3))
                _ = v22
            }
            var v23 int = fi19
            _ = fi19
            _ = v23
        }
    }
    v18 += (v18 >> ((v18) & 63))
    fmt.Println(v3)
    if true {
        fmt.Println(v3, v1)
        v2 = (v2 / (((v2 & v2)) | 1))
        var v24 uint32 = 2147483648
        _ = v24
    }
    if true {
        var v25 float64 = (0.0 + 0.0)
        var v26 float32 = 0.0
        var v27 int32 = (v2 % ((v2) | 1))
        _ = v25
        _ = v26
        _ = v27
    } else {
        fmt.Println(v18, v3, v18)
    }
    for fi28 := 0; fi28 < 5; fi28++ {
        v29 := []int{fi28, (v3 / ((128) | 1)), v1}
        var v30 int = v3
        for fi31 := 0; fi31 < 20; fi31++ {
            var v32 float64 = 1.5
            v18 += 0
            v33 := GenS3{F0: 64}
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi28
        _ = v29
        _ = v30
    }
    if true {
        var v34 int = v3
        var v35 uint32 = 42
        _ = v34
        _ = v35
    }
    v36 := []int{v3, (v1 / (((v3 << ((v18) & 63))) | 1)), 0}
    if len(v36) > 0 {
        v36[0] = v1
    }
    if (!true) {
        fmt.Println(v2, v1, v1)
        v36 = append(v36, (v1 & v1))
        if false {
            v2 += (-(v2))
        } else {
            fmt.Println(v18, v3)
            if false {
                var v37 string = "\n\t"
                v1 = v3
                var v38 uint16 = 2
                _ = v37
                _ = v38
            } else {
                v1 -= int(v2)
                v2 = v2
                if len(v36) > 0 {
                    v36[0] = -9223372036854775808
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v36
}
