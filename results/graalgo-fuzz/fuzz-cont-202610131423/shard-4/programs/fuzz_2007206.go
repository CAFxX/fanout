package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 int16; F1 int32 }
func genFn0(p0 int32) uint64 {
    var v5 int64 = 1
    var v6 int8 = -128
    v7 := []int{}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 2
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = -32768
    var v3 uint16 = 65535
    v4 := make(map[uint32]bool)
    var v8 float64 = ((float64(float32(0.0)) + float64(0.0)) / 3.14159)
    v9 := GenS1{F0: 363, F1: 0}
    v10 := []int{863, 1}
    if (v8 != v8) {
        if (4294967295 == 2) {
            fmt.Println(v2, v2)
            var v11 int64 = 1
            _ = v11
        } else {
            v12 := GenS2{F0: 255}
            switch 32767 {
            case 128:
                v3 -= v3
            case -148:
                if true {
                    if ((-827 < 42) && false) {
                        v13 := GenS3{F0: -2147483648, F1: -902}
                        v8 = (v8 * 1e10)
                        if false {
                            v2 = 659
                            var v14 float32 = 1e10
                            _ = v14
                        }
                        _ = v13
                    } else {
                        if false {
                            for fi15 := 0; fi15 < 16; fi15++ {
                                v8 = ((1e10 * v8) * ((1.5 / (v8 / (v8 / (0.1 * v8)))) + ((float64(0.0) - float64(0.1)) - (v8 * 1e10))))
                                _ = fi15
                            }
                            delete(v1, 65535)
                        } else {
                            fmt.Println(v3)
                            v2 = v2
                        }
                    }
                    v16 := v1[2147483648]
                    v17 := GenS1{F0: 269, F1: -9223372036854775808}
                    _ = v16
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 2; fi18++ {
                        switch "\n\t" {
                        case "foo bar":
                            for fi19 := 0; fi19 < 11; fi19++ {
                                v20 := GenS3{F0: 42, F1: 42}
                                _ = fi19
                                _ = v20
                            }
                        case "hello":
                            v2 -= 1
                        default:
                            for fi21 := 0; fi21 < 15; fi21++ {
                                switch (("hello" + "\n\t") + "\n\t") {
                                case "a":
                                    v3 += v3
                                    var v22 uint = 42
                                    _ = v22
                                }
                                v3 += v3
                                _ = fi21
                            }
                        }
                        for fi23 := 0; fi23 < 3; fi23++ {
                            fmt.Println(v2, v2)
                            _ = fi23
                        }
                        if (true || ((fi18 ^ fi18) < (^(fi18)))) {
                            v3 = (^(v3))
                        } else {
                            v24 := v4[890]
                            _ = v24
                        }
                        _ = fi18
                    }
                }
            }
            fmt.Println(v2, v8)
            _ = v12
        }
    }
    switch "\n\t" {
    case "\n\t":
        var v25 uint = 42
        _ = v25
    case "":
        switch -919 {
        case 32767:
            _ = len(v4)
        case 1:
            var v26 uint = 0
            _ = v26
        default:
            if false {
                var v27 float64 = v8
                var v28 uint16 = (v3 >> ((2) & 63))
                v2 = v2
                _ = v27
                _ = v28
            }
        }
    }
    v29 := v1[0]
    var v30 uint8 = 214
    for fi31 := 0; fi31 < 19; fi31++ {
        var v32 uint64 = 42
        v33 := GenS4{F0: (^(v2)), F1: 0}
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 float64 = 1.5
    var v35 int = -805
    var v36 bool = false
    var v37 bool = v36
    if ((v2 - v2) <= v2) {
        for fi38 := 0; fi38 < 20; fi38++ {
            fmt.Println(v3, v34)
            v34 -= (v34 * 1e10)
            for fi39 := 0; fi39 < 14; fi39++ {
                fmt.Println(v34, v30, v30)
                _ = fi39
            }
            _ = fi38
        }
        var v40 uint32 = 0
        _ = v40
    } else {
        for fi41 := 0; fi41 < 1; fi41++ {
            if false {
                var v42 int16 = (v2 + 459)
                var v43 uint = uint(v30)
                switch fi41 {
                case -1:
                    switch (fi41 % ((v29) | 1)) {
                    case 9223372036854775807:
                        for fi44 := 0; fi44 < 4; fi44++ {
                            var v45 int64 = 1
                            v8 = 1e10
                            _ = fi44
                            _ = v45
                        }
                    case -305:
                        v30 = uint8(v43)
                    case 2147483647:
                        v3 = v3
                    }
                default:
                    v35 -= v29
                }
                _ = v42
                _ = v43
            }
            v30 += 127
            _ = fi41
        }
        for fi46 := 0; fi46 < 16; fi46++ {
            delete(v1, 4294967295)
            for fi47 := 0; fi47 < 18; fi47++ {
                v48 := GenS2{F0: 741}
                v8 = v8
                _ = fi47
                _ = v48
            }
            _ = fi46
        }
    }
    v30 += (v30 + (v30 - v30))
    if len(v10) > 0 {
        v10[0] = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v8
    _ = v9
}
