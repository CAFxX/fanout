package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 uint16 = 2
    var v2 float64 = ((0.0 + -2.25) - (float64(1e10) / 0.1))
    v3 := []uint32{792}
    v4 := make(map[uint32]bool)
    var v5 uint32 = 957
    v6 := []int{}
    if ((107 >= -128) && (!false)) {
        if false {
            if (!(!(999 < v5))) {
                fmt.Println(v2, v2)
            }
            fmt.Println(v5, v1)
        } else {
            var v7 string = ""
            for fi8 := 0; fi8 < 3; fi8++ {
                if len(v6) > 0 {
                    v6[0] = 127
                }
                var v9 uint = 0
                var v10 string = "hello"
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = v7
        }
        var v11 uint16 = (^(v1))
        _ = v11
    } else {
        for fi12 := 0; fi12 < 1; fi12++ {
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 string = ("hello" + "\n\t")
                if (1 != 2147483647) {
                    var v15 int = 261
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v6 = append(v6, fi12)
                    _ = v15
                } else {
                    switch v14 {
                    case "":
                        v5 += 2147483647
                    default:
                        v14 = v14
                    }
                }
                var v16 int = fi12
                _ = fi13
                _ = v14
                _ = v16
            }
            var v17 uint8 = 127
            _ = fi12
            _ = v17
        }
        fmt.Println(v5, v1)
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 int16 = 1
        v19 += v19
        _ = fi18
        _ = v19
    }
    if ((9223372036854775807 != 9223372036854775807) || false) {
        for fi20 := 0; fi20 < 6; fi20++ {
            v5 += v5
            _ = fi20
        }
        switch 268 {
        case 127:
            for fi21 := 0; fi21 < 14; fi21++ {
                _ = len(v4)
                if true {
                    var v22 string = "a"
                    _ = v22
                }
                _ = fi21
            }
        case 17:
            var v23 float32 = (1.5 * float32(v2))
            _ = v23
        default:
            v5 = (v5 * v5)
        }
    }
    var v24 int = 32767
    v25 := v4[(v5 << ((455) & 63))]
    var v26 int64 = 0
    v6 = append(v6, (v24 << ((65535) & 63)))
    var v27 uint64 = 903
    for fi28 := 0; fi28 < 18; fi28++ {
        var v29 int32 = -1
        for fi30 := 0; fi30 < 8; fi30++ {
            v2 -= float64(float32(0.1))
            if false {
                v27 -= v27
            }
            if v25 {
                v6 = append(v6, (fi28 / (((^(fi28))) | 1)))
                var v31 float32 = (-2.25 * ((0.1 + 3.14159) * (3.14159 / 0.1)))
                _ = v31
            }
            _ = fi30
        }
        _ = fi28
        _ = v29
    }
    var v32 uint16 = v1
    v6 = append(v6, -229)
    for fi33 := 0; fi33 < 19; fi33++ {
        v27 -= v27
        _ = fi33
    }
    if v25 {
        v4[0] = v25
        v27 = uint64(v5)
        switch int(v27) {
        case 673:
            v34 := GenS1{F0: -492}
            _ = v34
        case -2147483648:
            if (255 <= 65535) {
                var v35 int16 = -633
                v25 = true
                _ = v35
            }
        }
    }
    var v36 int16 = 1
    var v37 float64 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
