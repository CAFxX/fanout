package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]string)
    var v3 int64 = 400
    v4 := make(map[uint32]string)
    var v5 int = 4294967295
    delete(v2, 0)
    delete(v2, 0)
    var v6 uint8 = 128
    var v7 uint16 = 282
    for fi8 := 0; fi8 < 10; fi8++ {
        if false {
            v5 += fi8
            if ((fi8 & (v5 + v5)) <= v5) {
                v7 = 32767
                v9 := GenS1{F0: 42}
                _ = v9
            } else {
                v5 -= fi8
                switch (("a" + "hello") + ("" + "hello")) {
                case "foo bar":
                    for fi10 := 0; fi10 < 12; fi10++ {
                        if true {
                            v3 = v3
                        }
                        var v11 int32 = -1
                        if true {
                            v5 = 1
                        }
                        _ = fi10
                        _ = v11
                    }
                }
            }
        } else {
            switch (v5 | v5) {
            case 127:
                for fi12 := 0; fi12 < 14; fi12++ {
                    v3 += (v3 & v3)
                    switch "hello" {
                    case "\n\t":
                        var v13 string = "\n\t"
                        _ = v13
                    case "foo bar":
                        v14 := GenS1{F0: 2147483647}
                        _ = v14
                    case "":
                        if false {
                            if true {
                                v3 -= 771
                            }
                            v15 := GenS1{F0: 2}
                            _ = v15
                        } else {
                            var v16 int64 = 9223372036854775807
                            var v17 int64 = (v3 | -9223372036854775808)
                            _ = v16
                            _ = v17
                        }
                    default:
                        v18 := GenS2{F0: (v6 | 42)}
                        _ = v18
                    }
                    v6 = (v6 >> ((65535) & 63))
                    _ = fi12
                }
            default:
                var v19 uint = uint(v7)
                _ = v19
            }
            for fi20 := 0; fi20 < 6; fi20++ {
                var v21 float64 = (1.5 * 0.1)
                switch 586 {
                case -191:
                    fmt.Println(v5, v5, v3)
                case -74:
                    v22 := []int{fi8, (-(fi20)), 127}
                    _ = v22
                case 63:
                    var v23 uint = 97
                    _ = v23
                }
                v24 := []int{(v5 & v5), (fi20 | v5), (fi8 + (fi20 % ((fi8) | 1)))}
                _ = fi20
                _ = v21
                _ = v24
            }
        }
        if false {
            for fi25 := 0; fi25 < 15; fi25++ {
                var v26 uint16 = v7
                v27 := []int{-1, v5}
                _ = fi25
                _ = v26
                _ = v27
            }
        }
        for fi28 := 0; fi28 < 2; fi28++ {
            switch "foo bar" {
            case "\n\t":
                v6 = 127
            default:
                v7 += (v7 << ((65535) & 63))
            }
            v7 = (^(v7))
            v7 += (v7 / ((v7) | 1))
            _ = fi28
        }
        _ = fi8
    }
    if false {
        v6 -= (v6 << ((18446744073709551615) & 63))
    } else {
        var v29 int8 = 127
        _ = v29
    }
    v7 -= v7
    var v30 uint8 = (v6 | 0)
    fmt.Println(v6)
    if (0 == 100) {
        for fi31 := 0; fi31 < 16; fi31++ {
            var v32 uint16 = (v7 & (^(v7)))
            v3 = int64(v32)
            _ = fi31
            _ = v32
        }
    }
    _ = len(v1)
    var v33 float64 = (0.1 * 0.1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
