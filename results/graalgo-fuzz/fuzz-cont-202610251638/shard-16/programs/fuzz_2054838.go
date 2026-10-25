package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := []uint32{2147483647, 243}
    var v2 uint = 65535
    var v3 string = ("" + ("" + "\n\t"))
    v4 := []string{(v3 + v3), ((v3 + v3) + v3), v3}
    v5 := make(map[string]string)
    v6 := GenS1{F0: 255}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        v5[("hello" + v3)] = (v3 + "hello")
        var v8 uint8 = 1
        _ = fi7
        _ = v8
    }
    var v9 int32 = -1
    v9 = v9
    v9 = (v9 >> ((1) & 63))
    v9 = -520
    v9 += (^(v9))
    if (!true) {
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := v5[(v3 + ((v3 + v3) + v3))]
            _ = fi10
            _ = v11
        }
        if false {
            if false {
                v2 = v2
                if false {
                    v3 = v3
                    v1 = append(v1, 634)
                } else {
                    var v12 uint = (v2 & 2)
                    if (0.0 == (1.5 + 0.1)) {
                        var v13 int64 = -958
                        switch v3 {
                        case "\n\t":
                            var v14 int16 = 666
                            _ = v14
                        case "a":
                            for fi15 := 0; fi15 < 11; fi15++ {
                                var v16 bool = (3 > 32767)
                                _ = fi15
                                _ = v16
                            }
                        default:
                            var v17 float64 = (((1e10 + (3.14159 * (0.0 / (0.0 + ((float64(0.0) * (-2.25 / float64(((float32(3.14159) - float32((-2.25 - 0.0))) / ((3.14159 + float32(1.5)) / ((float32(1.5) + float32(-2.25)) + 0.0)))))) + (1e10 + 0.0)))))) + (3.14159 / 1e10)) / 1e10)
                            _ = v17
                        }
                        _ = v13
                    }
                    for fi18 := 0; fi18 < 5; fi18++ {
                        v19 := GenS2{F0: -1}
                        fmt.Println(v3)
                        _ = fi18
                        _ = v19
                    }
                    _ = v12
                }
            } else {
                if (!true) {
                    fmt.Println(v9, v9)
                    fmt.Println(v2, v2)
                } else {
                    var v20 int32 = (v9 >> ((990) & 63))
                    _ = v20
                }
                fmt.Println(v3, v2)
            }
            v21 := GenS1{F0: 255}
            _ = v21
        }
        if len(v4) > 0 {
            v4[0] = v3
        }
    }
    var v22 int32 = (v9 / ((19) | 1))
    if false {
        v4 = append(v4, ("a" + ""))
        switch (v3 + v3) {
        case "foo bar":
            for fi23 := 0; fi23 < 20; fi23++ {
                if (!false) {
                    for fi24 := 0; fi24 < 10; fi24++ {
                        v22 = (v9 >> ((v2) & 63))
                        _ = fi24
                    }
                    var v25 int32 = v9
                    if false {
                        var v26 float64 = 0.0
                        _ = v26
                    }
                    _ = v25
                }
                switch (fi23 ^ (fi23 & fi23)) {
                case -515:
                    for fi27 := 0; fi27 < 7; fi27++ {
                        v22 = (v9 << ((122) & 63))
                        v28 := v5[v3]
                        _ = fi27
                        _ = v28
                    }
                case 0:
                    var v29 uint8 = 128
                    _ = v29
                case 868:
                    v9 += v22
                }
                fmt.Println(v3, v2)
                _ = fi23
            }
        case "hello":
            v9 -= 0
        default:
            var v30 string = "\n\t"
            _ = v30
        }
    }
    v31 := GenS2{F0: 1}
    var v32 float32 = (0.0 * float32((3.14159 - (((0.0 + (1.5 * 3.14159)) + -2.25) / 1.0))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
