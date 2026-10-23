package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := []int{1}
    var v2 uint16 = 42
    v3 := []uint32{2147483648, uint32(v2), 4294967295}
    var v4 string = "\n\t"
    v2 -= 32767
    v5 := GenS1{F0: -392}
    v2 = 612
    v2 -= (v2 & 338)
    if true {
        var v6 int8 = 100
        _ = v6
    } else {
        var v7 int = 32767
        if ((194 == 654) || true) {
            if len(v1) > 0 {
                v1[0] = int(v2)
            }
            for fi8 := 0; fi8 < 7; fi8++ {
                v9 := GenS1{F0: (v7 << ((1) & 63))}
                _ = fi8
                _ = v9
            }
            fmt.Println(v4, v4)
        }
        _ = v7
    }
    var v10 uint16 = v2
    var v11 int32 = 630
    if false {
        var v12 int = 127
        if false {
            v2 = 32767
            v10 += v2
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 uint16 = 255
                for fi15 := 0; fi15 < 19; fi15++ {
                    if false {
                        var v16 int16 = 1
                        var v17 uint8 = 197
                        _ = v16
                        _ = v17
                    }
                    if true {
                        switch 107 {
                        case 388:
                            if len(v3) > 0 {
                                v3[0] = 2
                            }
                        }
                    } else {
                        switch v4 {
                        case "":
                            v4 = "hello"
                        case "\n\t":
                            var v18 uint16 = 2
                            _ = v18
                        default:
                            switch v4 {
                            case "\n\t":
                                v19 := GenS2{F0: 127}
                                _ = v19
                            default:
                                if len(v1) > 0 {
                                    v1[0] = (fi13 * fi15)
                                }
                                v3 = append(v3, 65535)
                                v2 = 668
                            }
                        }
                        v20 := GenS1{F0: (fi15 << ((937) & 63))}
                        _ = v20
                    }
                    _ = fi15
                }
                v10 += (v2 >> ((116) & 63))
                _ = fi13
                _ = v14
            }
        } else {
            v10 += v2
            for fi21 := 0; fi21 < 7; fi21++ {
                var v22 uint = 42
                switch (fi21 * (^(fi21))) {
                case 32768:
                    v11 = 1
                case 2147483647:
                    fmt.Println(v22, v2)
                }
                var v23 uint8 = 127
                _ = fi21
                _ = v22
                _ = v23
            }
            var v24 int8 = -110
            _ = v24
        }
        v2 = v2
        _ = v12
    } else {
        switch (v4 + v4) {
        case "a":
            v25 := GenS2{F0: 127}
            _ = v25
        case "\n\t":
            var v26 int = -666
            _ = v26
        case "foo bar":
            v11 -= v11
        }
    }
    var v27 uint64 = 0
    fmt.Println(v27)
    v27 = v27
    v4 = ("foo bar" + v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
