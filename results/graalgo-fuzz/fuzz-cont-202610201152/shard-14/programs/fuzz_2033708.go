package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int }
type GenS4 struct { F0 uint16 }
func genFn0(p0 string, p1 int8) int32 {
    fmt.Println(p1)
    switch -963 {
    case 127:
        for fi6 := 0; fi6 < 14; fi6++ {
            v7 := GenS1{F0: 255}
            fmt.Println(p1, fi6, p1)
            _ = fi6
            _ = v7
        }
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 478
}
func genFn1(p0 uint8, p1 int16, p2 uint32) uint64 {
    p0 += (p0 * (^(p0)))
    v8 := GenS1{F0: 123}
    var v9 uint = 214
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 774
}

func main() {
    var v1 float64 = ((1.5 + 0.0) + 1e10)
    var v2 int16 = 32767
    var v3 string = ("hello" + "foo bar")
    var v4 uint64 = 4294967295
    var v5 float32 = 0.0
    var v10 int32 = -839
    v11 := []int{361, -1, -9223372036854775808, 32768}
    v3 = v3
    for fi12 := 0; fi12 < 9; fi12++ {
        v10 = v10
        _ = fi12
    }
    v13 := GenS1{F0: 127}
    v14 := GenS2{F0: v10, F1: (v2 << ((833) & 63))}
    var v15 int16 = (v2 ^ (v2 & (v2 + (v2 * -1))))
    switch ("\n\t" + v3) {
    case "hello":
        for fi16 := 0; fi16 < 3; fi16++ {
            fmt.Println(v1, v10, v4)
            v4 = 834
            v2 = (v2 | (v15 & (^(v2))))
            _ = fi16
        }
    case "\n\t":
        var v17 float32 = v5
        _ = v17
    case "foo bar":
        _ = len(v11)
    default:
        var v18 int32 = v10
        _ = v18
    }
    fmt.Println(v10)
    var v19 float64 = 1.5
    v3 = v3
    for fi20 := 0; fi20 < 15; fi20++ {
        v21 := GenS2{F0: (v10 * v10), F1: (v2 ^ (-(v2)))}
        _ = fi20
        _ = v21
    }
    v10 = (v10 / ((2147483647) | 1))
    var v22 string = ("hello" + (v3 + (v3 + v3)))
    v23 := GenS3{F0: 153, F1: 1, F2: 971}
    for fi24 := 0; fi24 < 3; fi24++ {
        if false {
            for fi25 := 0; fi25 < 20; fi25++ {
                v10 += v10
                fmt.Println(v15, v19)
                fmt.Println(v1, v19, v5)
                _ = fi25
            }
            if (2 < 104) {
                v22 = ("hello" + v3)
            }
        } else {
            if (397 < 121) {
                if ((1.5 + 0.1) < float64(((0.0 + v5) + 0.0))) {
                    for fi26 := 0; fi26 < 13; fi26++ {
                        v2 -= (^(v15))
                        var v27 int = 42
                        _ = fi26
                        _ = v27
                    }
                    if len(v11) > 0 {
                        v11[0] = fi24
                    }
                    switch fi24 {
                    case 127:
                        for fi28 := 0; fi28 < 20; fi28++ {
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            _ = fi28
                        }
                    case 32767:
                        if (true || (-2.25 < v19)) {
                            if false {
                                for fi29 := 0; fi29 < 12; fi29++ {
                                    var v30 uint32 = 1
                                    v31 := GenS4{F0: 0}
                                    v15 += 32767
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                }
                                v4 = v4
                                switch (v3 + (v3 + v22)) {
                                case "":
                                    var v32 uint64 = genFn1(0, v15, 42)
                                    _ = v32
                                default:
                                    if len(v11) > 0 {
                                        v11[0] = (-(fi24))
                                    }
                                }
                            } else {
                                var v33 float32 = v5
                                _ = v33
                            }
                        } else {
                            v5 -= (((((float32((((((((v19 - v1) + 3.14159) + 0.1) - (v1 + v1)) / v19) - float64(float32(v1))) / float64(((3.14159 - (((v5 / (1e10 * v5)) + ((3.14159 + 0.1) * ((float32((((1e10 / float64(((float32((v1 * (-2.25 - 0.0))) + float32(v19)) * v5))) / (((-2.25 - v1) - (1.5 * float64(1e10))) - (v19 - (0.0 - 1e10)))) / v19)) / v5) + (1.5 + 0.1)))) + v5)) / float32(v1))))) + float32((((v19 + (v19 / ((float64(v5) + 1e10) - float64(v5)))) - 1.5) - v1))) * 0.1) - float32(float64((0.1 / v5)))) * v5) + v5)
                            if true {
                                v22 = (v22 + ((v3 + v3) + ""))
                                for fi34 := 0; fi34 < 11; fi34++ {
                                    v11 = append(v11, fi24)
                                    v10 += (v10 << ((834) & 63))
                                    v35 := GenS3{F0: 128, F1: 255, F2: (fi34 + int(v10))}
                                    _ = fi34
                                    _ = v35
                                }
                            }
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                        }
                    case 42:
                        v15 = int16(v10)
                    }
                } else {
                    if true {
                        var v36 bool = (uint32(v4) > 2147483647)
                        if ((-113 <= 55) || v36) {
                            var v37 int16 = v2
                            for fi38 := 0; fi38 < 17; fi38++ {
                                if true {
                                    v2 = (v37 - (v2 << ((155) & 63)))
                                    var v39 int64 = -1
                                    var v40 int32 = (v10 - v10)
                                    _ = v39
                                    _ = v40
                                } else {
                                    fmt.Println(fi38)
                                }
                                v41 := GenS3{F0: 81, F1: 509, F2: -9223372036854775808}
                                v1 = (0.0 / v19)
                                _ = fi38
                                _ = v41
                            }
                            _ = v37
                        }
                        _ = v36
                    } else {
                        _ = len(v11)
                        var v42 bool = (!(!(((v1 / (0.0 + (v1 - v19))) - (v1 * float64(v5))) > float64((0.1 * float32(v19))))))
                        _ = v42
                    }
                    switch fi24 {
                    case -881:
                        var v43 string = v3
                        _ = v43
                    case -409:
                        v5 = 3.14159
                    case -9223372036854775808:
                        if len(v11) > 0 {
                            v11[0] = (fi24 - (fi24 ^ fi24))
                        }
                    }
                }
                var v44 uint32 = 0
                v45 := GenS1{F0: 0}
                _ = v44
                _ = v45
            }
            fmt.Println(v19)
        }
        for fi46 := 0; fi46 < 8; fi46++ {
            fmt.Println(fi24, fi24, fi24)
            _ = fi46
        }
        v22 = ((v3 + "foo bar") + "foo bar")
        _ = fi24
    }
    v47 := GenS2{F0: v10, F1: 0}
    if true {
        var v48 int8 = 100
        _ = v48
    } else {
        var v49 uint64 = (v4 & v4)
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
