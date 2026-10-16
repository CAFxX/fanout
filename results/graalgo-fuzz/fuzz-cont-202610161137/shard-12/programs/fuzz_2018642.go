package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int32 }
type GenS2 struct { F0 bool; F1 int64 }

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = 546
    v3 := make(map[int]bool)
    fmt.Println(v2)
    if ("foo bar" >= "a") {
        fmt.Println(v2)
    }
    v4 := GenS1{F0: 675, F1: 2, F2: 0}
    var v5 int64 = 42
    for fi6 := 0; fi6 < 8; fi6++ {
        if false {
            var v7 uint16 = 641
            v5 = (v5 - v5)
            for fi8 := 0; fi8 < 13; fi8++ {
                v7 = 2
                _ = fi8
            }
            _ = v7
        }
        v9 := GenS2{F0: (617 < 746), F1: (v5 * (v5 * (^(v5))))}
        var v10 float32 = 1.5
        _ = fi6
        _ = v10
        _ = v9
    }
    v11 := v1[-1]
    var v12 uint16 = 32767
    fmt.Println(v11, v12, v11)
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 string = "foo bar"
        _ = len(v1)
        v15 := v1[fi13]
        _ = fi13
        _ = v14
        _ = v15
    }
    delete(v3, -449)
    var v16 int32 = -933
    if (v11 && (201 > 0)) {
        if (65535 == 2) {
            for fi17 := 0; fi17 < 14; fi17++ {
                for fi18 := 0; fi18 < 7; fi18++ {
                    fmt.Println(v11, fi18, fi18)
                    _ = fi18
                }
                if v11 {
                    var v19 int16 = v2
                    _ = v19
                }
                _ = fi17
            }
        } else {
            if (!(8 != 0)) {
                for fi20 := 0; fi20 < 17; fi20++ {
                    var v21 bool = v11
                    var v22 uint32 = 626
                    v23 := GenS1{F0: 18446744073709551615, F1: 1, F2: (v16 * v16)}
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
                for fi24 := 0; fi24 < 11; fi24++ {
                    v11 = false
                    var v25 string = "a"
                    _ = fi24
                    _ = v25
                }
                fmt.Println(v12, v12, v11)
            }
            switch "hello" {
            case "a":
                switch 32768 {
                case 127:
                    var v26 float64 = 0.0
                    _ = v26
                default:
                    v27 := GenS1{F0: 42, F1: 255, F2: v16}
                    _ = v27
                }
            }
        }
    } else {
        fmt.Println(v16, v5)
    }
    var v28 int8 = 1
    if (!(v11 && v11)) {
        v12 -= (v12 - 867)
        for fi29 := 0; fi29 < 13; fi29++ {
            v5 -= -1
            v30 := GenS1{F0: 0, F1: 2, F2: int32(v5)}
            var v31 float64 = 1.5
            _ = fi29
            _ = v30
            _ = v31
        }
    } else {
        v32 := GenS2{F0: (524 >= 677), F1: v5}
        if (18446744073709551615 == 18446744073709551615) {
            var v33 uint8 = 255
            fmt.Println(v16, v2)
            for fi34 := 0; fi34 < 16; fi34++ {
                var v35 uint8 = v33
                _ = fi34
                _ = v35
            }
            _ = v33
        }
        _ = v32
    }
    var v36 float32 = 1e10
    var v37 string = ""
    v12 = v12
    var v38 int16 = (v2 - (v2 & int16(v28)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
