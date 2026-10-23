package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }

func main() {
    var v1 int = -262
    var v2 uint32 = 428
    v3 := make(map[string]int)
    var v4 uint64 = 2
    v3[(("" + "foo bar") + ("foo bar" + "foo bar"))] = 32767
    v5 := []int{(v1 << ((0) & 63)), (v1 + v1), v1, (v1 + (v1 + -9223372036854775808))}
    v4 = v4
    if false {
        if (false && true) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v5 = append(v5, int(v2))
            for fi6 := 0; fi6 < 10; fi6++ {
                _ = len(v5)
                _ = fi6
            }
        } else {
            v7 := GenS1{F0: 42}
            if true {
                switch "\n\t" {
                case "":
                    v1 = -9223372036854775808
                }
            }
            _ = v7
        }
        for fi8 := 0; fi8 < 10; fi8++ {
            if (255 <= 2) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            fmt.Println(v1, v1)
            _ = fi8
        }
    } else {
        fmt.Println(v4)
        v1 -= v1
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        v4 = (v4 + v4)
        var v10 uint64 = v4
        _ = fi9
        _ = v10
    }
    var v11 uint16 = 0
    v12 := GenS1{F0: 42}
    v13 := GenS1{F0: 128}
    v11 += v11
    v11 = v11
    if (false || ((0 == v4) && (v11 < 894))) {
        v14 := GenS1{F0: 216}
        for fi15 := 0; fi15 < 6; fi15++ {
            if false {
                if len(v5) > 0 {
                    v5[0] = v1
                }
                if (!true) {
                    v3[(("\n\t" + "foo bar") + "hello")] = v1
                }
            } else {
                var v16 uint8 = 128
                var v17 int64 = -822
                _ = v16
                _ = v17
            }
            if (true || true) {
                if true {
                    v11 -= 42
                    if false {
                        var v18 bool = (1 <= -1)
                        _ = v18
                    } else {
                        var v19 float32 = 1.5
                        for fi20 := 0; fi20 < 1; fi20++ {
                            var v21 float32 = 0.1
                            var v22 bool = (v19 < (v21 / ((((1e10 - v19) / 1e10) / (-2.25 - 0.1)) * ((v19 / float32(3.14159)) / float32((1.5 + (float64(v19) - (0.1 / 0.1))))))))
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                        _ = v19
                    }
                    fmt.Println(v2, v11)
                }
            } else {
                if len(v5) > 0 {
                    v5[0] = fi15
                }
                v4 -= (v4 * (v4 / (((v4 + uint64(v11))) | 1)))
            }
            _ = fi15
        }
        _ = v14
    } else {
        v23 := GenS2{F0: v1}
        v11 = (v11 % (((v11 << ((1) & 63))) | 1))
        for fi24 := 0; fi24 < 11; fi24++ {
            v11 = v11
            delete(v3, ("foo bar" + "foo bar"))
            if true {
                v3[""] = (fi24 - -746)
            }
            _ = fi24
        }
        _ = v23
    }
    v4 += (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
