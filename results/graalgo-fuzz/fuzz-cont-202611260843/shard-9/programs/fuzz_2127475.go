package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint, p1 int, p2 uint) uint8 {
    p1 += p1
    if (982 >= 0) {
        if false {
            var v5 uint = (p2 ^ p0)
            var v6 int32 = 2147483647
            _ = v5
            _ = v6
        }
    }
    var v7 float32 = float32(0.0)
    var v8 uint16 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2
}
func genFn1() int16 {
    var v9 float32 = 3.14159
    if (!true) {
        var v10 float64 = -2.25
        for fi11 := 0; fi11 < 20; fi11++ {
            for fi12 := 0; fi12 < 14; fi12++ {
                v13 := []int{}
                if (-32768 >= 32767) {
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                } else {
                    switch fi12 {
                    case 1:
                        var v14 uint16 = 65535
                        _ = v14
                    case 32767:
                        v15 := GenS1{F0: 65535}
                        _ = v15
                    default:
                        var v16 int = 0
                        _ = v16
                    }
                    var v17 int64 = -1
                    v18 := GenS1{F0: 96}
                    _ = v17
                    _ = v18
                }
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
        switch "hello" {
        case "hello":
            var v19 int = 297
            _ = v19
        case "foo bar":
            v20 := []int{1}
            _ = v20
        default:
            v10 += ((((v10 / (v10 + v10)) / ((0.1 * 1e10) * v10)) + float64(v9)) + 1e10)
        }
        _ = v10
    } else {
        v9 += v9
        if false {
            v21 := []int{-981, -1, 1, -716}
            switch 9 {
            case -979:
                for fi22 := 0; fi22 < 16; fi22++ {
                    for fi23 := 0; fi23 < 2; fi23++ {
                        var v24 uint16 = 42
                        _ = fi23
                        _ = v24
                    }
                    switch fi22 {
                    case 300:
                        v21 = append(v21, fi22)
                    case -592:
                        var v25 string = ""
                        _ = v25
                    default:
                        var v26 int16 = 446
                        _ = v26
                    }
                    _ = fi22
                }
            case 127:
                switch -160 {
                case 32768:
                    switch "hello" {
                    case "a":
                        var v27 int32 = -830
                        _ = v27
                    case "":
                        var v28 int16 = 42
                        _ = v28
                    }
                case 127:
                    v29 := GenS1{F0: 2}
                    _ = v29
                }
            }
            switch 753 {
            case 42:
                for fi30 := 0; fi30 < 9; fi30++ {
                    var v31 uint16 = 255
                    var v32 uint16 = (^(v31))
                    var v33 uint = 18446744073709551615
                    _ = fi30
                    _ = v31
                    _ = v32
                    _ = v33
                }
            case 4294967295:
                if (18446744073709551615 < 33) {
                    v9 = 0.1
                    for fi34 := 0; fi34 < 6; fi34++ {
                        var v35 int16 = -521
                        _ = fi34
                        _ = v35
                    }
                }
            case 1:
                _ = len(v21)
            }
            _ = v21
        } else {
            var v36 float64 = 0.1
            fmt.Println(v36, v36)
            _ = v36
        }
        var v37 float64 = (3.14159 / (((-2.25 / 1.5) + 1.5) * (3.14159 * (float64(v9) - 3.14159))))
        _ = v37
    }
    var v38 uint = 18446744073709551615
    v9 = 3.14159
    for fi39 := 0; fi39 < 10; fi39++ {
        var v40 int64 = int64(v38)
        v40 = -306
        _ = fi39
        _ = v40
    }
    _ = v38
    _ = v9
    return -32768
}

func main() {
    var v1 int8 = 100
    var v2 bool = false
    var v3 uint = 302
    var v4 uint = v3
    v41 := []int{}
    v4 = (^(v4))
    v4 = 301
    var v42 float64 = ((((3.14159 / -2.25) - 1e10) - 1e10) - (((1e10 * 1e10) * -2.25) - 0.0))
    var v43 float64 = v42
    var v44 uint8 = 2
    for fi45 := 0; fi45 < 2; fi45++ {
        var v46 uint16 = uint16(v44)
        _ = len(v41)
        _ = fi45
        _ = v46
    }
    var v47 int64 = 1
    var v48 int = -689
    var v49 uint = (v4 - (^(v3)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v47
    _ = v48
    _ = v49
}
