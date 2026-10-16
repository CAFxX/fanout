package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{4294967295, 611, 0}
    var v3 uint = 503
    v4 := GenS1{F0: 894}
    var v5 uint = 4294967295
    switch 32768 {
    case 83:
        switch -2147483648 {
        case 2147483647:
            v5 += 586
        case -324:
            fmt.Println(v5, v5)
        case -9223372036854775808:
            if (65535 <= 2147483648) {
                var v6 uint64 = 18446744073709551615
                if len(v2) > 0 {
                    v2[0] = 58
                }
                v7 := GenS1{F0: 32767}
                _ = v6
                _ = v7
            } else {
                if true {
                    v1[4294967295] = uint32(v3)
                    switch 128 {
                    case -1:
                        if (127 <= 18) {
                            var v8 uint16 = 8
                            _ = v8
                        }
                    case -737:
                        switch ("a" + ("\n\t" + "a")) {
                        case "a":
                            v3 -= (v3 | v3)
                        case "\n\t":
                            v9 := GenS1{F0: 42}
                            _ = v9
                        case "foo bar":
                            v3 += v3
                        default:
                            v5 = v5
                        }
                    case -2147483648:
                        _ = len(v1)
                    }
                }
                v3 = v5
                var v10 int16 = 1000
                _ = v10
            }
        default:
            var v11 int32 = 0
            _ = v11
        }
    case -218:
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 uint = v3
            var v14 float32 = 0.0
            switch fi12 {
            case 9223372036854775807:
                for fi15 := 0; fi15 < 14; fi15++ {
                    v16 := GenS1{F0: -401}
                    if ((-1 == -1) && true) {
                        var v17 int = fi12
                        var v18 float32 = 1e10
                        _ = v17
                        _ = v18
                    } else {
                        var v19 float32 = ((-2.25 / v14) * float32(1e10))
                        _ = v19
                    }
                    _ = fi15
                    _ = v16
                }
            }
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    v3 = (v5 % ((v5) | 1))
    v5 += v3
    var v20 uint = v3
    v21 := GenS1{F0: -248}
    if true {
        fmt.Println(v5)
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 int16 = 32767
        v23 = (v23 & v23)
        _ = fi22
        _ = v23
    }
    if len(v2) > 0 {
        v2[0] = 65535
    }
    v24 := v1[644]
    if (2147483647 < -1) {
        for fi25 := 0; fi25 < 12; fi25++ {
            var v26 uint8 = 0
            _ = fi25
            _ = v26
        }
        var v27 float32 = 1.5
        _ = v27
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        for fi29 := 0; fi29 < 7; fi29++ {
            switch (fi29 << ((2) & 63)) {
            case 4294967295:
                delete(v1, 2147483648)
            case -916:
                var v30 int64 = -1
                _ = v30
            default:
                var v31 float32 = 0.0
                _ = v31
            }
            v5 += 255
            fmt.Println(fi28, v20)
            _ = fi29
        }
        _ = fi28
    }
    v2 = append(v2, (v24 * v24))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
