package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int16 }
type GenS2 struct { F0 int; F1 uint8; F2 int }
func genFn0(p0 int, p1 int8) uint64 {
    p1 -= 0
    if true {
        fmt.Println(p1, p0)
        switch p0 {
        case 68:
            var v4 int32 = 511
            _ = v4
        case 127:
            v5 := []int{(^(p0))}
            _ = v5
        default:
            if (false || (true && false)) {
                var v6 uint32 = 4294967295
                _ = v6
            }
        }
        fmt.Println(p1, p0)
    }
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[string]uint32)
    var v3 float32 = 3.14159
    for fi7 := 0; fi7 < 17; fi7++ {
        if true {
            for fi8 := 0; fi8 < 11; fi8++ {
                _ = len(v2)
                switch fi8 {
                case 2147483647:
                    var v9 uint32 = 0
                    _ = v9
                case 1:
                    v1 += genFn0(fi8, 100)
                case 361:
                    var v10 float32 = (-2.25 + ((v3 - (3.14159 * v3)) * float32(float64((float32(0.1) + (v3 + 3.14159))))))
                    _ = v10
                }
                var v11 uint8 = 2
                v3 += -2.25
                v12 := GenS1{F0: 75, F1: 803, F2: 32767}
                v11 = (^(v11))
                var v13 int32 = -2147483648
                var v14 bool = true
                _ = fi8
                _ = v11
                _ = v12
                _ = v13
                _ = v14
            }
            if true {
                v3 = ((0.0 + float32(1e10)) * v3)
                if ((true && ((0 > 0) && (false && (12 >= 2)))) && (82 != -1)) {
                    v3 = (0.1 - 1e10)
                    v15 := []int{(fi7 & fi7), fi7, fi7}
                    _ = v15
                }
                switch 1 {
                case 2147483647:
                    if false {
                        var v16 int = (fi7 / ((fi7) | 1))
                        _ = v16
                    }
                case 0:
                    var v17 uint = 82
                    _ = v17
                }
            } else {
                if true {
                    v1 -= genFn0((fi7 - 2147483647), -79)
                    v3 = (1.5 / 0.1)
                } else {
                    if false {
                        switch fi7 {
                        case 127:
                            var v18 float64 = 0.0
                            _ = v18
                        }
                    }
                }
                v1 = (v1 ^ 111)
                var v19 uint64 = (v1 >> ((42) & 63))
                _ = v19
            }
            var v20 int32 = 0
            _ = v20
        } else {
            v21 := []int{}
            if (-1 <= -805) {
                v3 = float32(float64(0.0))
                if (-2.25 >= 1e10) {
                    var v22 int = 0
                    _ = v22
                }
                for fi23 := 0; fi23 < 19; fi23++ {
                    v3 = float32((3.14159 + 3.14159))
                    var v24 uint64 = v1
                    var v25 int16 = 42
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
            switch ("a" + "hello") {
            case "hello":
                fmt.Println(fi7, v3, v3)
            }
            _ = v21
        }
        fmt.Println(v1, v3)
        _ = fi7
    }
    fmt.Println(v1, v3)
    v1 = 2
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 uint = 42
        if true {
            var v28 uint8 = 1
            _ = v28
        }
        _ = fi26
        _ = v27
    }
    switch -39 {
    case 1:
        v29 := GenS1{F0: 352, F1: 2147483647, F2: -112}
        _ = v29
    }
    v30 := v2["\n\t"]
    for fi31 := 0; fi31 < 4; fi31++ {
        var v32 float32 = float32(((((1.5 - 3.14159) + 3.14159) - float64(v3)) * float64((v3 - (((0.0 * (v3 / (v3 + (v3 / v3)))) * float32(0.0)) * v3)))))
        var v33 float32 = (0.0 * 0.1)
        _ = fi31
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 6; fi34++ {
        if false {
            v35 := GenS2{F0: fi34, F1: uint8(v1), F2: fi34}
            var v36 string = "\n\t"
            for fi37 := 0; fi37 < 10; fi37++ {
                if true {
                    fmt.Println(v36, v1, v1)
                } else {
                    var v38 uint64 = v1
                    var v39 int64 = 207
                    v3 -= float32((0.1 / 1.0))
                    _ = v38
                    _ = v39
                }
                v40 := GenS2{F0: 128, F1: 8, F2: (fi34 + (fi34 >> ((65535) & 63)))}
                _ = fi37
                _ = v40
            }
            _ = v35
            _ = v36
        } else {
            fmt.Println(v30, v1, v1)
            v1 = genFn0(fi34, -1)
        }
        v30 -= (v30 - (v30 / (((^(v30))) | 1)))
        var v41 string = "\n\t"
        _ = fi34
        _ = v41
    }
    v30 = (v30 ^ v30)
    fmt.Println(v30, v1)
    v1 = 172
    var v42 float64 = 1.5
    v43 := GenS1{F0: 0, F1: (v30 | v30), F2: 0}
    v44 := GenS1{F0: 460, F1: 42, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v42
    _ = v43
    _ = v44
}
