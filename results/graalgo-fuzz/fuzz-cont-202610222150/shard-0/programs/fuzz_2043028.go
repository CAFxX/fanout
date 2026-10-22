package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint16; F1 uint8; F2 int8 }
func genFn0(p0 uint64) uint8 {
    var v6 string = "hello"
    var v7 string = v6
    v8 := []int{42, 0, 127}
    p0 -= 351
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return uint8(p0)
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 65535
    v3 := make(map[uint32]int)
    v4 := []uint32{2147483647, 754, 827}
    v5 := make(map[uint32]uint32)
    if false {
        switch -1 {
        case 228:
            if false {
                switch ((("hello" + "a") + "hello") + "a") {
                case "\n\t":
                    fmt.Println(v2)
                case "foo bar":
                    v9 := GenS1{F0: v2}
                    _ = v9
                case "a":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            } else {
                fmt.Println(v2)
            }
        case 2147483647:
            var v10 uint8 = 128
            _ = v10
        }
        var v11 float32 = 0.1
        v11 += v11
        _ = v11
    } else {
        v12 := v3[42]
        _ = v12
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        v2 += (v2 << ((437) & 63))
        if true {
            if false {
                v4 = append(v4, 4294967295)
                if ((true && false) || true) {
                    var v14 int16 = 0
                    var v15 uint8 = 1
                    _ = v14
                    _ = v15
                }
            }
        } else {
            var v16 float32 = -2.25
            var v17 uint32 = 792
            if (((1e10 / 0.1) / -2.25) > 0.1) {
                if len(v4) > 0 {
                    v4[0] = 443
                }
                v16 -= v16
                for fi18 := 0; fi18 < 16; fi18++ {
                    if (!false) {
                        var v19 float64 = (((0.0 - -2.25) * 1e10) * float64((((1e10 + (v16 - (v16 * v16))) * ((v16 * (v16 * v16)) + (0.1 + 0.1))) / (v16 - v16))))
                        var v20 int = fi13
                        if true {
                            var v21 uint8 = uint8(v2)
                            var v22 float32 = (v16 / (v16 * (((v16 / (v16 / v16)) + v16) + -2.25)))
                            var v23 string = ""
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        _ = v19
                        _ = v20
                    } else {
                        fmt.Println(fi13, fi18, v17)
                    }
                    var v24 int = fi18
                    _ = fi18
                    _ = v24
                }
            }
            _ = v16
            _ = v17
        }
        _ = fi13
    }
    var v25 string = ("\n\t" + "hello")
    v2 = v2
    v2 = (v2 | (v2 / ((221) | 1)))
    var v26 float32 = 0.0
    for fi27 := 0; fi27 < 19; fi27++ {
        for fi28 := 0; fi28 < 11; fi28++ {
            _ = len(v4)
            _ = fi28
        }
        v2 -= v2
        if false {
            var v29 int64 = 1
            _ = v29
        }
        _ = fi27
    }
    var v30 int64 = -654
    if true {
        if (("" > v25) || false) {
            v25 = v25
        }
        if (2 <= 2) {
            v25 = v25
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            fmt.Println(v26, v30, v2)
        }
        if true {
            switch 1 {
            case 128:
                switch 2 {
                case 127:
                    var v31 uint = 4294967295
                    _ = v31
                case 1:
                    var v32 int32 = -970
                    _ = v32
                }
            case -958:
                v33 := GenS2{F0: (^(v2)), F1: 0, F2: 159}
                _ = v33
            default:
                for fi34 := 0; fi34 < 5; fi34++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    switch "hello" {
                    case "\n\t":
                        for fi35 := 0; fi35 < 7; fi35++ {
                            v36 := GenS3{F0: v2, F1: 233, F2: 42}
                            var v37 float32 = v26
                            _ = fi35
                            _ = v36
                            _ = v37
                        }
                    case "":
                        fmt.Println(v25, v2, v26)
                    case "a":
                        var v38 float32 = float32(0.1)
                        _ = v38
                    default:
                        if false {
                            if true {
                                var v39 int8 = -1
                                _ = v39
                            } else {
                                var v40 uint32 = 338
                                _ = v40
                            }
                            v30 += v30
                        }
                    }
                    v41 := GenS1{F0: v2}
                    _ = fi34
                    _ = v41
                }
            }
        } else {
            v4 = append(v4, 42)
        }
    } else {
        var v42 int = 42
        _ = v42
    }
    for fi43 := 0; fi43 < 7; fi43++ {
        fmt.Println(v25)
        fmt.Println(v25)
        var v44 uint = 0
        _ = fi43
        _ = v44
    }
    fmt.Println(v26, v2, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
