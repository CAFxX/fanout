package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 uint32 }

func main() {
    var v1 int16 = 1
    v2 := make(map[uint32]bool)
    var v3 int16 = (-(v1))
    var v4 int16 = v1
    var v5 int = 122
    var v6 uint = 0
    var v7 int8 = 0
    v3 += (v4 - v3)
    v8 := GenS1{F0: 625}
    v1 -= (v1 ^ v3)
    delete(v2, 0)
    _ = len(v2)
    var v9 uint = (^(v6))
    for fi10 := 0; fi10 < 11; fi10++ {
        v5 += (fi10 << ((v6) & 63))
        switch v5 {
        case 42:
            v7 = 1
        case 703:
            if false {
                switch int(v7) {
                case -9223372036854775808:
                    var v11 string = "foo bar"
                    _ = v11
                case 1:
                    var v12 float64 = 0.0
                    _ = v12
                }
            } else {
                v13 := []int{(fi10 * int(v3)), (fi10 - (v5 * v5)), (v5 >> ((v6) & 63)), -402}
                switch (-(fi10)) {
                case -2147483648:
                    for fi14 := 0; fi14 < 1; fi14++ {
                        if ((3.14159 - -2.25) >= ((-2.25 + 0.0) * float32(0.0))) {
                            var v15 uint16 = 65535
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            _ = v15
                        }
                        switch fi10 {
                        case 42:
                            switch "foo bar" {
                            case "foo bar":
                                var v16 float32 = (0.0 - ((0.1 + float32(3.14159)) * ((0.0 * -2.25) * (float32(-2.25) / (0.1 * (float32(3.14159) + (-2.25 - 0.1)))))))
                                _ = v16
                            }
                        case -510:
                            if (v4 == (-(v3))) {
                                var v17 float32 = -2.25
                                _ = v17
                            }
                        case -9223372036854775808:
                            v4 += (-(v1))
                        }
                        _ = fi14
                    }
                case -314:
                    for fi18 := 0; fi18 < 9; fi18++ {
                        var v19 int64 = -9223372036854775808
                        _ = fi18
                        _ = v19
                    }
                case 467:
                    var v20 int16 = v4
                    _ = v20
                }
                if (-9223372036854775808 == -1) {
                    v21 := GenS1{F0: 119}
                    _ = len(v2)
                    v22 := GenS1{F0: 1}
                    _ = v21
                    _ = v22
                } else {
                    for fi23 := 0; fi23 < 18; fi23++ {
                        var v24 float64 = 1e10
                        _ = fi23
                        _ = v24
                    }
                    fmt.Println(v5)
                }
                _ = v13
            }
        case 1:
            v25 := GenS2{F0: v9, F1: 2147483647}
            _ = v25
        default:
            v7 -= (v7 + (v7 << (((^(v9))) & 63)))
        }
        _ = fi10
    }
    fmt.Println(v6, v1)
    delete(v2, 42)
    for fi26 := 0; fi26 < 16; fi26++ {
        for fi27 := 0; fi27 < 17; fi27++ {
            v28 := GenS2{F0: v6, F1: 196}
            v29 := []int{}
            _ = fi27
            _ = v28
            _ = v29
        }
        fmt.Println(v6)
        v5 -= fi26
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
