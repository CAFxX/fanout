package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = false
    v3 := []uint32{0, 0, 65535}
    v2 = false
    _ = len(v1)
    fmt.Println(v2)
    v4 := GenS1{F0: 1}
    switch "a" {
    case "foo bar":
        v2 = false
    case "a":
        var v5 uint64 = 147
        _ = v5
    }
    var v6 string = "foo bar"
    v6 = ("" + (v6 + v6))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v7 int8 = -1
    if ((v6 + ("" + v6)) == (("\n\t" + "\n\t") + (v6 + v6))) {
        v8 := GenS1{F0: 597}
        var v9 uint16 = 0
        v9 += (^(v9))
        _ = v8
        _ = v9
    }
    var v10 uint8 = 127
    v10 += (^(v10))
    v10 -= (^(v10))
    if (-1 == 723) {
        if v2 {
            delete(v1, 9223372036854775807)
        }
    } else {
        for fi11 := 0; fi11 < 1; fi11++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v12 int64 = 1
            if v2 {
                for fi13 := 0; fi13 < 4; fi13++ {
                    var v14 uint16 = 805
                    v7 = (v7 | v7)
                    if len(v3) > 0 {
                        v3[0] = 479
                    }
                    _ = fi13
                    _ = v14
                }
                var v15 float64 = (float64((0.1 + 1e10)) + (((0.0 * (1e10 + 1.5)) + 0.0) - 0.1))
                if ((164 >= 42) || v2) {
                    var v16 bool = true
                    if v16 {
                        fmt.Println(v7, v12, v7)
                        v17 := GenS1{F0: 2}
                        _ = v17
                    } else {
                        fmt.Println(v12, v2, v16)
                        var v18 uint = 42
                        if true {
                            _ = len(v1)
                            v19 := GenS2{F0: fi11}
                            var v20 int8 = v7
                            _ = v19
                            _ = v20
                        }
                        _ = v18
                    }
                    v15 += v15
                    _ = v16
                }
                _ = v15
            } else {
                var v21 float32 = float32(((float64((-2.25 * 3.14159)) / -2.25) - (((float64(1.5) - 1.5) * (1e10 + 3.14159)) * 0.0)))
                v7 = (v7 * (^(v7)))
                _ = v21
            }
            _ = fi11
            _ = v12
        }
    }
    switch int(v10) {
    case 9223372036854775807:
        for fi22 := 0; fi22 < 2; fi22++ {
            var v23 uint8 = v10
            _ = fi22
            _ = v23
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
