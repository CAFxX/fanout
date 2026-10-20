package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 uint; F2 bool }
func genFn0(p0 uint, p1 uint8) string {
    v6 := GenS1{F0: p0}
    p0 = (p0 + p0)
    p1 += (p1 << ((p0) & 63))
    _ = p0
    _ = p1
    _ = v6
    return ("hello" + "hello")
}

func main() {
    var v1 bool = (-1 >= 395)
    var v2 bool = ((float64((0.0 + 0.1)) != (0.1 * 3.14159)) || v1)
    v3 := make(map[string]bool)
    var v4 int32 = 352
    var v5 string = (("\n\t" + "\n\t") + "hello")
    v7 := GenS1{F0: 1}
    if (973 != 222) {
        var v8 int64 = 436
        if (true && ((float32(-2.25) + ((3.14159 - 1.5) - (0.1 - (((0.0 * 1e10) - (float32(0.0) / (-2.25 * 1e10))) / float32(1.5))))) > 1e10)) {
            for fi9 := 0; fi9 < 13; fi9++ {
                if (v2 || v2) {
                    v10 := GenS2{F0: v4, F1: 0}
                    fmt.Println(v5, v8, v1)
                    _ = v10
                }
                switch ((v5 + (v5 + (v5 + v5))) + ((v5 + v5) + v5)) {
                case "hello":
                    for fi11 := 0; fi11 < 6; fi11++ {
                        if ((v4 / ((v4) | 1)) == 2147483647) {
                            v12 := GenS1{F0: 255}
                            _ = v12
                        }
                        v13 := []int{(fi11 / ((-9223372036854775808) | 1)), fi11}
                        var v14 int16 = 0
                        _ = fi11
                        _ = v13
                        _ = v14
                    }
                case "a":
                    switch fi9 {
                    case -547:
                        switch fi9 {
                        case 32767:
                            var v15 int = (fi9 ^ (fi9 | 42))
                            _ = v15
                        case 0:
                            var v16 int32 = (v4 << ((2) & 63))
                            _ = v16
                        case 205:
                            v17 := []int{fi9, (^(fi9)), -1, fi9}
                            _ = v17
                        }
                    }
                default:
                    if (v2 || v2) {
                        var v18 uint16 = 32767
                        v4 += v4
                        if (!true) {
                            var v19 uint16 = 65535
                            _ = len(v3)
                            _ = v19
                        } else {
                            v20 := []int{-605, fi9}
                            _ = v20
                        }
                        _ = v18
                    }
                }
                v4 = (v4 * v4)
                _ = fi9
            }
            v21 := GenS3{F0: 2}
            _ = v21
        }
        _ = v8
    } else {
        var v22 uint32 = 42
        var v23 uint8 = 42
        if ((!v2) && (952 == 0)) {
            v2 = v2
            v24 := GenS2{F0: (-(v4)), F1: 0}
            var v25 int32 = v4
            _ = v24
            _ = v25
        }
        _ = v22
        _ = v23
    }
    v26 := GenS1{F0: 255}
    fmt.Println(v1, v4, v4)
    for fi27 := 0; fi27 < 6; fi27++ {
        var v28 uint32 = 2147483647
        var v29 float32 = 1.5
        _ = fi27
        _ = v28
        _ = v29
    }
    v30 := v3[((v5 + ("" + "a")) + v5)]
    v31 := []int{}
    for fi32 := 0; fi32 < 12; fi32++ {
        var v33 int = fi32
        switch ("\n\t" + ("" + "hello")) {
        case "hello":
            if ((^(v4)) != (v4 - v4)) {
                v33 -= v33
            } else {
                switch v33 {
                case 32768:
                    if ((-1 != 1) || ((1 < 127) && (!(1 >= 42)))) {
                        for fi34 := 0; fi34 < 16; fi34++ {
                            var v35 int = 42
                            var v36 int64 = 29
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                        var v37 float32 = -2.25
                        for fi38 := 0; fi38 < 20; fi38++ {
                            v5 = v5
                            _ = fi38
                        }
                        _ = v37
                    } else {
                        if (true || (-1 < 42)) {
                            var v39 uint64 = 1
                            v30 = true
                            v40 := GenS4{F0: 255, F1: (v39 ^ v39)}
                            _ = v39
                            _ = v40
                        } else {
                            v41 := GenS3{F0: 255}
                            var v42 float32 = 1e10
                            _ = v41
                            _ = v42
                        }
                        var v43 int64 = 1
                        _ = v43
                    }
                case 9223372036854775807:
                    v44 := GenS5{F0: -606}
                    _ = v44
                }
            }
        case "":
            var v45 int64 = -1
            _ = v45
        case "\n\t":
            fmt.Println(v30, fi32)
        default:
            if v30 {
                if v1 {
                    if (9223372036854775807 == 421) {
                        v2 = v2
                        v46 := GenS6{F0: ((v33 | fi32) < v33), F1: 0, F2: false}
                        _ = v46
                    }
                    for fi47 := 0; fi47 < 16; fi47++ {
                        if len(v31) > 0 {
                            v31[0] = 127
                        }
                        fmt.Println(v4, v33)
                        var v48 float64 = (float64(((1e10 / 1e10) * 1.5)) * 1e10)
                        _ = fi47
                        _ = v48
                    }
                }
                if (!(v1 || true)) {
                    v4 = v4
                    if len(v31) > 0 {
                        v31[0] = (-(v33))
                    }
                    var v49 int64 = 42
                    _ = v49
                }
            }
        }
        v3[(v5 + v5)] = v2
        _ = fi32
        _ = v33
    }
    var v50 int64 = -1
    for fi51 := 0; fi51 < 13; fi51++ {
        for fi52 := 0; fi52 < 15; fi52++ {
            var v53 uint32 = 561
            v2 = v1
            _ = fi52
            _ = v53
        }
        switch v5 {
        case "foo bar":
            v31 = append(v31, fi51)
        }
        _ = fi51
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v50
    _ = v7
}
