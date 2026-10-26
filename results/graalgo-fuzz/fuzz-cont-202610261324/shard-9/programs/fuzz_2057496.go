package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int16, p1 string) bool {
    p1 = ((p1 + (p1 + p1)) + (p1 + (p1 + p1)))
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return true
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[int]int)
    v3 := []uint32{586}
    v4 := []int{}
    v5 := make(map[uint32]int)
    var v6 int64 = -773
    switch "foo bar" {
    case "\n\t":
        for fi7 := 0; fi7 < 13; fi7++ {
            for fi8 := 0; fi8 < 1; fi8++ {
                switch (("hello" + "") + "") {
                case "":
                    fmt.Println(v6, v6)
                default:
                    switch "" {
                    case "hello":
                        v6 = (v6 << ((350) & 63))
                        v1 -= (v1 << ((2) & 63))
                    default:
                        v9 := GenS1{F0: 0}
                        var v10 int16 = -32768
                        _ = v10
                        _ = v9
                    }
                }
                var v11 uint = 849
                _ = fi8
                _ = v11
            }
            _ = fi7
        }
    case "hello":
        if false {
            var v12 uint64 = 551
            _ = v12
        }
    default:
        for fi13 := 0; fi13 < 20; fi13++ {
            switch (fi13 * fi13) {
            case 532:
                for fi14 := 0; fi14 < 19; fi14++ {
                    switch (fi14 / (((fi13 & fi13)) | 1)) {
                    case -2147483648:
                        switch ("\n\t" + "") {
                        case "hello":
                            if (!false) {
                                var v15 string = "foo bar"
                                var v16 string = v15
                                _ = v15
                                _ = v16
                            } else {
                                fmt.Println(fi13, v1)
                                var v17 bool = false
                                fmt.Println(fi13)
                                _ = v17
                            }
                        }
                    case 32768:
                        v18 := GenS1{F0: -1}
                        _ = v18
                    case -188:
                        v19 := GenS1{F0: 362}
                        _ = v19
                    default:
                        var v20 float32 = (0.0 * 0.0)
                        _ = v20
                    }
                    if ((4 != 42) && true) {
                        var v21 uint16 = 760
                        var v22 float64 = 1e10
                        var v23 uint16 = (v21 + v21)
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        v6 -= -9223372036854775808
                        v24 := GenS1{F0: -1}
                        _ = v24
                    }
                    v6 -= (v6 * (^(v6)))
                    _ = fi14
                }
            case 1:
                v25 := GenS1{F0: 2147483647}
                _ = v25
            case 938:
                var v26 uint32 = 0
                _ = v26
            }
            var v27 uint32 = 4294967295
            v1 = (v1 >> ((979) & 63))
            _ = fi13
            _ = v27
        }
    }
    v6 -= -1
    var v28 string = ""
    v28 = ""
    v1 -= v1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v6, v28, v6)
    v29 := GenS1{F0: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
