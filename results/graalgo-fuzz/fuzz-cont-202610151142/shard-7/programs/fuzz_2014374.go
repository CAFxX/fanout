package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint32; F1 uint8; F2 int16 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]string)
    var v3 uint = 42
    var v4 uint = (v3 ^ (v3 >> ((1) & 63)))
    v5 := []uint32{42, 543, 2, 65535}
    fmt.Println(v4, v3)
    v6 := GenS1{F0: uint32(v3), F1: (v3 & (v4 & (v4 + v4)))}
    switch "foo bar" {
    case "foo bar":
        for fi7 := 0; fi7 < 16; fi7++ {
            v3 -= (v4 >> ((v4) & 63))
            var v8 float32 = 0.0
            if true {
                if len(v5) > 0 {
                    v5[0] = 33
                }
            }
            _ = fi7
            _ = v8
        }
    default:
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 uint32 = 0
            for fi11 := 0; fi11 < 20; fi11++ {
                if false {
                    var v12 int32 = -1
                    v5 = append(v5, v10)
                    if true {
                        var v13 uint64 = 2
                        _ = v13
                    }
                    v1[(v10 & v10)] = (4294967295 >= 1)
                    if false {
                        var v14 int16 = 32767
                        if true {
                            var v15 uint32 = 65535
                            switch 127 {
                            case 0:
                                var v16 int64 = 484
                                _ = v16
                            case 42:
                                var v17 uint = v4
                                _ = v17
                            case 595:
                                var v18 float32 = 1.5
                                _ = v18
                            }
                            _ = v15
                        } else {
                            var v19 int32 = -695
                            _ = v19
                        }
                        _ = v14
                    }
                    _ = len(v1)
                    v10 = 0
                    _ = v12
                } else {
                    var v20 bool = true
                    v21 := GenS2{F0: (v10 ^ v10), F1: 42, F2: -1}
                    if true {
                        var v22 int64 = 1
                        switch (-(fi11)) {
                        case 247:
                            v22 = v22
                        case 42:
                            if v20 {
                                v4 += (v3 * 4294967295)
                                v23 := GenS2{F0: (v10 | 2147483648), F1: 255, F2: -59}
                                _ = v23
                            }
                        }
                        switch "hello" {
                        case "":
                            var v24 uint64 = 42
                            _ = v24
                        case "foo bar":
                            var v25 int32 = 1
                            _ = v25
                        default:
                            fmt.Println(v3, v20)
                        }
                        _ = v22
                    }
                    _ = v20
                    _ = v21
                }
                v3 = 65535
                v1[490] = (-766 >= 32767)
                _ = fi11
            }
            _ = fi9
            _ = v10
        }
    }
    var v26 uint8 = 103
    if (!true) {
        for fi27 := 0; fi27 < 5; fi27++ {
            delete(v1, 2)
            var v28 string = "foo bar"
            v28 = v28
            _ = fi27
            _ = v28
        }
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v29 uint16 = 65535
    var v30 uint32 = 0
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
