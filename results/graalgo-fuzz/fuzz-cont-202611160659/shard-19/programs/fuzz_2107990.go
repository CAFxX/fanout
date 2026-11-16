package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() int64 {
    if false {
        var v5 int = -1
        v5 = (v5 % ((173) | 1))
        _ = v5
    } else {
        if (50 >= 2) {
            var v6 int16 = 374
            v6 = v6
            _ = v6
        }
    }
    switch 32767 {
    case -9223372036854775808:
        var v7 float64 = float64(float32(((0.0 - 0.1) / float64(1.5))))
        _ = v7
    case -357:
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 float32 = 0.1
            _ = fi8
            _ = v9
        }
    case -2147483648:
        var v10 int8 = 70
        _ = v10
    }
    return -9223372036854775808
}
func genFn1(p0 uint64, p1 uint) int64 {
    p1 = 42
    var v11 int64 = -9223372036854775808
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := GenS1{F0: -671}
        var v14 string = "a"
        var v15 int8 = 1
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v11
    return (v11 / ((v11) | 1))
}

func main() {
    var v1 uint16 = 312
    var v2 uint16 = (v1 & v1)
    var v3 int16 = 42
    v4 := []uint32{65535, 65535, 674}
    v3 += v3
    v1 = (^(v2))
    v4 = append(v4, 2147483647)
    v2 += (v1 >> ((55) & 63))
    fmt.Println(v3, v2)
    var v16 int = 128
    switch v16 {
    case 53:
        v1 = (v1 % ((v1) | 1))
    case 507:
        switch 691 {
        case -329:
            v3 = 32767
        case 9223372036854775807:
            var v17 uint32 = 4294967295
            _ = v17
        }
    case 4294967295:
        switch "\n\t" {
        case "foo bar":
            for fi18 := 0; fi18 < 11; fi18++ {
                v16 = (fi18 << ((388) & 63))
                var v19 string = ("" + ((("\n\t" + "") + "a") + "foo bar"))
                v20 := GenS1{F0: 0}
                _ = fi18
                _ = v19
                _ = v20
            }
        }
    }
    _ = len(v4)
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 uint32 = 42
        var v23 int32 = 15
        switch -523 {
        case 42:
            var v24 bool = (-97 <= 0)
            _ = v24
        case 127:
            var v25 int8 = 71
            _ = v25
        case 4294967295:
            switch fi21 {
            case 42:
                v3 += (v3 + (^(v3)))
            case 51:
                if false {
                    fmt.Println(fi21, fi21)
                } else {
                    if (v3 > (v3 + v3)) {
                        for fi26 := 0; fi26 < 19; fi26++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v27 int = v16
                            _ = fi26
                            _ = v27
                        }
                        for fi28 := 0; fi28 < 1; fi28++ {
                            fmt.Println(v22, v1, v1)
                            _ = fi28
                        }
                        var v29 int32 = (-(v23))
                        _ = v29
                    } else {
                        var v30 float64 = 0.0
                        v30 += v30
                        switch int(v23) {
                        case 456:
                            v3 = (v3 - (^(v3)))
                        }
                        _ = v30
                    }
                    fmt.Println(v23)
                }
            case 32768:
                v3 = v3
            default:
                v4 = append(v4, v22)
            }
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
