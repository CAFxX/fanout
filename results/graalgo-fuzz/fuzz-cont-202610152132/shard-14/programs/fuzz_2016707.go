package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int64 }

func main() {
    v1 := []string{"", ("\n\t" + "")}
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    var v4 int8 = 100
    v5 := make(map[int]uint32)
    var v6 int = 128
    fmt.Println(v4)
    if true {
        if (!true) {
            v4 += v4
            if (951 >= 567) {
                if false {
                    var v7 int = (v6 & v6)
                    _ = v7
                }
                fmt.Println(v6, v4)
            }
            if (false && true) {
                fmt.Println(v4, v6)
            } else {
                switch v6 {
                case 128:
                    delete(v5, (v6 - v6))
                case -2147483648:
                    v3[591] = "a"
                }
                for fi8 := 0; fi8 < 16; fi8++ {
                    fmt.Println(fi8, v6)
                    _ = fi8
                }
                v2[(("" + "hello") + "\n\t")] = -9223372036854775808
            }
        } else {
            var v9 uint16 = 65535
            for fi10 := 0; fi10 < 15; fi10++ {
                v9 = v9
                if (0.0 <= float32(0.1)) {
                    fmt.Println(v4)
                } else {
                    v11 := GenS1{F0: 4294967295, F1: 219, F2: -559}
                    _ = v11
                }
                delete(v3, 65535)
                _ = fi10
            }
            _ = v9
        }
        var v12 uint64 = 142
        v6 = (v6 | v6)
        _ = v12
    }
    fmt.Println(v4)
    if (9223372036854775807 == -1) {
        fmt.Println(v4, v6)
    }
    v13 := GenS1{F0: 284, F1: 44, F2: -1}
    v4 = v4
    var v14 uint32 = 65535
    if ((v14 % ((v14) | 1)) <= v14) {
        var v15 int = (v6 >> ((4294967295) & 63))
        var v16 string = (("hello" + "\n\t") + "\n\t")
        _ = v15
        _ = v16
    } else {
        var v17 string = (("hello" + "a") + "")
        if (true || (!false)) {
            fmt.Println(v17)
            for fi18 := 0; fi18 < 13; fi18++ {
                v4 += (v4 ^ v4)
                if true {
                    v1 = append(v1, "")
                }
                if true {
                    for fi19 := 0; fi19 < 17; fi19++ {
                        if (548 < 18446744073709551615) {
                            switch -1 {
                            case 801:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        }
                        v2["\n\t"] = v6
                        v14 = (v14 / ((v14) | 1))
                        if (!false) {
                            v6 += (v6 - v6)
                        }
                        _ = fi19
                    }
                    var v20 string = (v17 + "\n\t")
                    var v21 int8 = (v4 * (v4 - (v4 * v4)))
                    _ = v20
                    _ = v21
                } else {
                    v5[v6] = v14
                }
                _ = fi18
            }
            var v22 string = v17
            _ = v22
        }
        _ = v17
    }
    var v23 uint8 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
