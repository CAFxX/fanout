package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint) int64 {
    v4 := GenS1{F0: true}
    v5 := GenS1{F0: false}
    v6 := []int{26, 128, -203}
    for fi7 := 0; fi7 < 14; fi7++ {
        p0 = 515
        _ = fi7
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 int64 = 457
    var v2 uint8 = 127
    v3 := make(map[string]bool)
    if (233 > 830) {
        v2 = (v2 + v2)
    } else {
        v1 += 42
        var v8 int64 = (v1 >> ((42) & 63))
        v9 := GenS2{F0: -128}
        _ = v8
        _ = v9
    }
    var v10 int64 = 0
    v1 = 1
    var v11 uint32 = 2
    fmt.Println(v10)
    v11 = v11
    v10 += v10
    v2 = 128
    if true {
        v12 := []int{}
        _ = v12
    } else {
        fmt.Println(v1, v1)
        v13 := []int{4294967295, 4294967295, -394}
        if true {
            for fi14 := 0; fi14 < 5; fi14++ {
                switch (fi14 % ((128) | 1)) {
                case 128:
                    v1 = (-(v1))
                case -9223372036854775808:
                    for fi15 := 0; fi15 < 17; fi15++ {
                        switch "hello" {
                        case "hello":
                            if (true || false) {
                                v11 = v11
                                v10 = 125
                            } else {
                                switch fi14 {
                                case 128:
                                    fmt.Println(v1, fi14, v11)
                                case -463:
                                    v10 = genFn0(311)
                                default:
                                    var v16 uint16 = 42
                                    _ = v16
                                }
                            }
                        case "\n\t":
                            switch fi15 {
                            case 32768:
                                fmt.Println(v10, v1)
                            }
                            v17 := GenS2{F0: -32}
                            _ = len(v13)
                            _ = v17
                        case "foo bar":
                            v2 = (v2 >> ((42) & 63))
                            v18 := GenS1{F0: true}
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            _ = v18
                        }
                        _ = fi15
                    }
                }
                var v19 float64 = 3.14159
                for fi20 := 0; fi20 < 5; fi20++ {
                    v11 -= (v11 >> ((65535) & 63))
                    _ = fi20
                }
                _ = fi14
                _ = v19
            }
        }
        _ = v13
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        v22 := GenS2{F0: 42}
        v23 := GenS1{F0: true}
        v1 = (v10 * v1)
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
}
