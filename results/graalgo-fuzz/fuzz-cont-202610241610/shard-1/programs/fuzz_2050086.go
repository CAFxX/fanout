package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int64; F1 uint64 }

func main() {
    v1 := []int{127}
    var v2 int = 622
    var v3 uint8 = 42
    v4 := make(map[int]bool)
    var v5 uint16 = 32767
    var v6 string = "hello"
    var v7 float64 = 3.14159
    v8 := GenS1{F0: 1}
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 int = 370
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := GenS1{F0: 1}
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 14; fi14++ {
            switch v6 {
            case "foo bar":
                var v15 uint = 18446744073709551615
                _ = v15
            }
            _ = fi14
        }
        _ = fi9
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        var v17 bool = false
        v3 += 127
        for fi18 := 0; fi18 < 12; fi18++ {
            switch v6 {
            case "foo bar":
                var v19 int8 = -55
                _ = v19
            default:
                var v20 string = "foo bar"
                _ = v20
            }
            v21 := v4[fi18]
            _ = fi18
            _ = v21
        }
        _ = fi16
        _ = v17
    }
    var v22 uint16 = v5
    v2 = (v2 << ((542) & 63))
    for fi23 := 0; fi23 < 2; fi23++ {
        v24 := GenS2{F0: 530, F1: 255}
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        for fi26 := 0; fi26 < 2; fi26++ {
            if len(v1) > 0 {
                v1[0] = v2
            }
            v6 = v6
            if true {
                delete(v4, (fi26 & v2))
                var v27 int32 = 2147483647
                switch fi25 {
                case 128:
                    if true {
                        v28 := v4[fi25]
                        switch (v2 << ((0) & 63)) {
                        case 193:
                            var v29 uint = 255
                            _ = v29
                        case 824:
                            var v30 bool = (!(!false))
                            _ = v30
                        }
                        _ = v28
                    } else {
                        _ = len(v1)
                    }
                case 865:
                    var v31 int32 = -980
                    _ = v31
                default:
                    v32 := GenS3{F0: -9223372036854775808, F1: 18446744073709551615}
                    _ = v32
                }
                _ = v27
            } else {
                v33 := GenS2{F0: 42, F1: 988}
                var v34 float32 = (0.1 * 1e10)
                _ = v33
                _ = v34
            }
            _ = fi26
        }
        _ = fi25
    }
    if (!((v3 ^ v3) < v3)) {
        for fi35 := 0; fi35 < 13; fi35++ {
            if false {
                v36 := v4[(fi35 / (((-(fi35))) | 1))]
                delete(v4, int(v3))
                _ = v36
            }
            v22 = 65535
            _ = fi35
        }
    }
    var v37 uint32 = 2
    var v38 uint = 65535
    v39 := GenS1{F0: -984}
    fmt.Println(v37)
    var v40 int = 2147483647
    v7 = ((v7 - ((v7 - v7) + (float64(0.0) * 0.1))) * float64(3.14159))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
