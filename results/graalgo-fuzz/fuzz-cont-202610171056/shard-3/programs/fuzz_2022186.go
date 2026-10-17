package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 bool }
type GenS2 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 uint8 = 0
    var v2 int = 2147483647
    v3 := make(map[uint32]bool)
    var v4 int = v2
    fmt.Println(v4)
    if false {
        var v5 float64 = float64(1.5)
        _ = v5
    }
    for fi6 := 0; fi6 < 13; fi6++ {
        if (!true) {
            for fi7 := 0; fi7 < 15; fi7++ {
                v1 = v1
                _ = fi7
            }
            v1 -= 1
        } else {
            switch v4 {
            case -719:
                v2 += 9223372036854775807
            default:
                if ((253 == 2147483647) || true) {
                    v8 := GenS1{F0: -2147483648, F1: (^(v1)), F2: (2 != 4294967295)}
                    v9 := []int{(fi6 ^ 706), -250, (fi6 & v2), 822}
                    _ = v8
                    _ = v9
                }
            }
            var v10 int32 = 425
            delete(v3, 65535)
            _ = v10
        }
        var v11 bool = false
        _ = fi6
        _ = v11
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 1; fi13++ {
            v14 := GenS1{F0: 42, F1: (v1 * 180), F2: (false || (!(!true)))}
            switch int(v1) {
            case 128:
                v15 := GenS1{F0: 498, F1: (v1 + 1), F2: (!true)}
                _ = v15
            case 4294967295:
                v16 := GenS1{F0: -481, F1: (^(v1)), F2: (2 != 4294967295)}
                _ = v16
            case 0:
                var v17 uint8 = v1
                _ = v17
            default:
                v18 := GenS1{F0: 0, F1: 128, F2: (0 <= 530)}
                _ = v18
            }
            var v19 uint32 = 2147483648
            if ("" == (("hello" + "\n\t") + "")) {
                var v20 int32 = 1
                _ = v20
            } else {
                v2 = (v2 & v2)
                v21 := GenS1{F0: 0, F1: v1, F2: ((191 >= 653) || ("" != "a"))}
                if false {
                    v22 := []int{}
                    var v23 int64 = int64(v19)
                    _ = v22
                    _ = v23
                } else {
                    var v24 int = v2
                    _ = v24
                }
                _ = v21
            }
            _ = fi13
            _ = v14
            _ = v19
        }
        var v25 uint = 463
        _ = fi12
        _ = v25
    }
    var v26 uint16 = 255
    delete(v3, 371)
    v27 := GenS2{F0: 127, F1: 941}
    if (!false) {
        fmt.Println(v4)
        v4 += v2
        var v28 float64 = 1e10
        _ = v28
    } else {
        fmt.Println(v26)
        v26 += 1
        v29 := v3[1]
        _ = v29
    }
    if (!true) {
        if (false && (!true)) {
            if true {
                v1 += v1
                v26 -= 220
                if (v26 != v26) {
                    v4 = v4
                    var v30 uint = 2
                    _ = v30
                } else {
                    var v31 int16 = -1
                    _ = v31
                }
            }
            var v32 uint = 65535
            _ = v32
        }
        var v33 uint64 = 142
        delete(v3, 2147483647)
        _ = v33
    }
    _ = len(v3)
    for fi34 := 0; fi34 < 8; fi34++ {
        fmt.Println(v4, v26, v1)
        var v35 int = (fi34 >> ((2) & 63))
        _ = fi34
        _ = v35
    }
    v26 += v26
    if (3.14159 <= (1e10 / 1.0)) {
        v36 := []int{128, 1, v4, (v2 * (v4 * v4))}
        v1 += (v1 & 65)
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
