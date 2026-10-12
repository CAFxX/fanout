package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int64) uint64 {
    if (true || (0.0 >= 3.14159)) {
        var v7 bool = ((((((0.1 + float32(0.0)) + ((3.14159 * -2.25) + ((0.0 + 1e10) / 1.5))) / 3.14159) + 0.1) <= (0.1 / 1.0)) && ((65535 >= 1) && false))
        _ = v7
    } else {
        var v8 uint16 = 32767
        _ = v8
    }
    if false {
        p0 += (^(p0))
    }
    switch "foo bar" {
    case "":
        v9 := GenS1{F0: 255}
        _ = v9
    case "a":
        if (42 >= -1) {
            if (1 == 255) {
                v10 := []int{9223372036854775807}
                v11 := GenS1{F0: 42}
                if false {
                    for fi12 := 0; fi12 < 2; fi12++ {
                        p0 = (-(p0))
                        _ = fi12
                    }
                    for fi13 := 0; fi13 < 8; fi13++ {
                        p0 = -121
                        for fi14 := 0; fi14 < 5; fi14++ {
                            p0 = p0
                            _ = fi14
                        }
                        p0 = p0
                        _ = fi13
                    }
                    for fi15 := 0; fi15 < 18; fi15++ {
                        p0 = p0
                        var v16 string = "hello"
                        _ = fi15
                        _ = v16
                    }
                }
                _ = v10
                _ = v11
            } else {
                if false {
                    var v17 int = 32768
                    v18 := []int{v17, (v17 + v17), 127, v17}
                    _ = v17
                    _ = v18
                } else {
                    var v19 float32 = 1.5
                    _ = v19
                }
            }
            p0 += (p0 << ((255) & 63))
        }
    }
    _ = p0
    return 2
}

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[int]int)
    v3 := []string{}
    var v4 uint = 255
    v5 := make(map[int]bool)
    var v6 int8 = 2
    switch "\n\t" {
    case "foo bar":
        var v20 int64 = 1
        _ = v20
    case "hello":
        v6 -= (-(v6))
    case "a":
        if false {
            switch "a" {
            case "hello":
                var v21 int8 = (v6 + 127)
                _ = v21
            case "":
                for fi22 := 0; fi22 < 17; fi22++ {
                    var v23 uint = v4
                    _ = fi22
                    _ = v23
                }
            }
            v1 -= v1
        } else {
            var v24 int16 = 542
            v25 := v5[-2147483648]
            _ = v24
            _ = v25
        }
    }
    v3 = append(v3, "a")
    if true {
        v4 -= v4
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v1 = (v1 * 4294967295)
    var v26 uint16 = 1
    var v27 uint8 = 128
    v28 := GenS1{F0: v26}
    var v29 int8 = v6
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
