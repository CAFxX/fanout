package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 string; F2 int64 }
type GenS3 struct { F0 string; F1 int }
func genFn0(p0 uint32, p1 int32) int64 {
    v7 := GenS1{F0: 42}
    for fi8 := 0; fi8 < 12; fi8++ {
        p0 = (p0 + 992)
        _ = fi8
    }
    var v9 int = 127
    if (false || false) {
        v9 -= v9
        var v10 float64 = (3.14159 * float64(-2.25))
        fmt.Println(v9, v9, v10)
        _ = v10
    } else {
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 uint64 = 4294967295
            v13 := GenS1{F0: 104}
            _ = fi11
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 15; fi14++ {
            if false {
                p1 = (-(p1))
                var v15 bool = true
                var v16 int16 = 1000
                _ = v15
                _ = v16
            } else {
                for fi17 := 0; fi17 < 12; fi17++ {
                    var v18 int64 = 948
                    var v19 uint16 = 736
                    if ((1 > 9223372036854775807) && (1 > 128)) {
                        var v20 int16 = -1
                        _ = v20
                    }
                    v9 += fi14
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
            v9 += (^(v9))
            _ = fi14
        }
        v21 := []int{(v9 / ((v9) | 1)), (-(v9)), v9}
        _ = v21
    }
    p0 += p0
    _ = p0
    _ = p1
    _ = v7
    _ = v9
    return -59
}
func genFn1(p0 uint64) uint {
    for fi22 := 0; fi22 < 17; fi22++ {
        p0 += (p0 << ((2) & 63))
        if false {
            if false {
                var v23 uint = 2
                var v24 int = (fi22 % ((805) | 1))
                var v25 int64 = 0
                _ = v23
                _ = v24
                _ = v25
            }
        }
        p0 = (^(p0))
        _ = fi22
    }
    var v26 uint16 = 255
    _ = p0
    _ = v26
    return 255
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]bool)
    var v3 string = ("a" + (("\n\t" + "foo bar") + "hello"))
    v4 := make(map[uint32]string)
    var v5 int16 = -737
    var v6 int16 = v5
    v27 := []int{0, 1}
    v5 = (v6 | (v5 >> ((18446744073709551615) & 63)))
    v5 += (v5 >> ((996) & 63))
    fmt.Println(v5, v6)
    fmt.Println(v5, v3, v3)
    var v28 uint16 = 255
    switch -324 {
    case -9223372036854775808:
        v3 = "hello"
    default:
        var v29 int32 = -1
        _ = v29
    }
    v30 := GenS1{F0: 65535}
    var v31 int8 = -90
    var v32 uint = 4294967295
    var v33 uint8 = 0
    v34 := GenS2{F0: 2147483647, F1: (v3 + (("hello" + "a") + v3)), F2: 42}
    if (-1 == genFn0(2147483647, 1)) {
        v6 -= v6
        v32 -= v32
        var v35 uint16 = v28
        _ = v35
    }
    var v36 bool = (0 != v32)
    if v36 {
        if false {
            for fi37 := 0; fi37 < 15; fi37++ {
                v6 = v5
                _ = fi37
            }
            fmt.Println(v3)
        } else {
            var v38 float32 = -2.25
            _ = v38
        }
    } else {
        var v39 int32 = 186
        v36 = (v32 >= (v32 | v32))
        v40 := GenS3{F0: "a", F1: 127}
        _ = v39
        _ = v40
    }
    var v41 uint64 = 18446744073709551615
    for fi42 := 0; fi42 < 14; fi42++ {
        switch v3 {
        case "a":
            v33 = v33
        case "":
            for fi43 := 0; fi43 < 6; fi43++ {
                fmt.Println(v33, v5)
                var v44 int8 = (-(v31))
                switch "a" {
                case "a":
                    var v45 uint = v32
                    _ = v45
                case "\n\t":
                    switch ((v3 + v3) + "a") {
                    case "a":
                        var v46 uint32 = uint32(v32)
                        _ = v46
                    case "\n\t":
                        var v47 int16 = (v6 | 42)
                        _ = v47
                    case "":
                        var v48 uint8 = v33
                        _ = v48
                    }
                }
                var v49 uint16 = v28
                v50 := GenS1{F0: (v32 >> (((v32 + genFn1(2))) & 63))}
                _ = fi43
                _ = v44
                _ = v49
                _ = v50
            }
        default:
            v32 -= v32
        }
        var v51 float64 = ((-2.25 - 0.1) - ((1.5 - (0.0 + float64(1e10))) / 0.1))
        v41 = v41
        _ = fi42
        _ = v51
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v36
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
