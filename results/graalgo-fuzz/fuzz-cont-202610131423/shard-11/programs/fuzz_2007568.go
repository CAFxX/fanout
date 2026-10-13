package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int16 }
func genFn0(p0 uint16) string {
    var v4 float64 = -2.25
    if false {
        switch -1 {
        case 4294967295:
            switch "" {
            case "hello":
                for fi5 := 0; fi5 < 18; fi5++ {
                    v6 := []int{}
                    _ = fi5
                    _ = v6
                }
            case "foo bar":
                var v7 uint64 = 18446744073709551615
                _ = v7
            case "":
                v4 = float64(-2.25)
            default:
                v4 = v4
            }
        case -1:
            if ((float64(0.0) * v4) <= (0.1 * 3.14159)) {
                for fi8 := 0; fi8 < 3; fi8++ {
                    var v9 float64 = v4
                    var v10 int8 = -128
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                var v11 int = -2147483648
                _ = v11
            }
        }
        for fi12 := 0; fi12 < 5; fi12++ {
            v4 -= (1.5 + v4)
            _ = fi12
        }
    }
    _ = p0
    _ = v4
    return "a"
}
func genFn1(p0 int64, p1 uint8, p2 bool) int32 {
    for fi13 := 0; fi13 < 5; fi13++ {
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 float32 = 0.0
            v16 := GenS1{F0: "hello"}
            _ = fi14
            _ = v15
            _ = v16
        }
        fmt.Println(fi13, p0)
        _ = fi13
    }
    v17 := []int{-191, 128, 410, 4294967295}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return -953
}

func main() {
    var v1 int16 = -1
    var v2 bool = (313 != 2147483648)
    var v3 bool = (v2 && v2)
    v18 := GenS1{F0: "a"}
    if (499 == -284) {
        v19 := []int{15, 1, -35}
        if v2 {
            v20 := GenS2{F0: 1}
            v21 := GenS3{F0: 0, F1: -1}
            _ = v20
            _ = v21
        } else {
            var v22 int16 = (v1 ^ 32767)
            if false {
                var v23 int64 = 1
                fmt.Println(v22)
                v2 = (!v2)
                _ = v23
            } else {
                switch 4294967295 {
                case 32767:
                    var v24 uint32 = 2
                    _ = v24
                case 127:
                    for fi25 := 0; fi25 < 3; fi25++ {
                        var v26 uint = 18446744073709551615
                        _ = fi25
                        _ = v26
                    }
                default:
                    switch 513 {
                    case 127:
                        v22 = (v22 & -1)
                    case 4294967295:
                        if v2 {
                            v27 := GenS1{F0: "\n\t"}
                            var v28 uint32 = 0
                            v29 := GenS1{F0: "hello"}
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                    default:
                        var v30 int32 = 42
                        _ = v30
                    }
                }
                switch -416 {
                case 4294967295:
                    v31 := GenS3{F0: 599, F1: -1}
                    _ = v31
                case -995:
                    v3 = v3
                case 127:
                    v2 = v2
                }
                if v3 {
                    var v32 uint32 = 4294967295
                    _ = v32
                }
            }
            var v33 uint16 = 32767
            _ = v22
            _ = v33
        }
        for fi34 := 0; fi34 < 17; fi34++ {
            v35 := GenS2{F0: 2147483647}
            var v36 int8 = 0
            var v37 string = "hello"
            _ = fi34
            _ = v35
            _ = v36
            _ = v37
        }
        _ = v19
    }
    var v38 int16 = 82
    v1 = (v38 / ((v1) | 1))
    switch "" {
    case "hello":
        v39 := []int{}
        _ = v39
    }
    var v40 string = "\n\t"
    var v41 uint32 = 396
    v42 := []int{}
    for fi43 := 0; fi43 < 2; fi43++ {
        var v44 int16 = 1
        _ = fi43
        _ = v44
    }
    v2 = (1 >= 2)
    _ = len(v42)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v42:", len(v42), v42)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v38
    _ = v40
    _ = v41
    _ = v42
}
