package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int16; F1 int16 }

func main() {
    var v1 uint8 = 108
    var v2 int = 127
    var v3 int32 = -2147483648
    v4 := []string{("foo bar" + ("hello" + "")), "hello"}
    var v5 uint32 = 0
    v6 := make(map[int]int)
    _ = len(v4)
    switch (v2 >> ((18446744073709551615) & 63)) {
    case 127:
        v5 += (v5 * (v5 / ((v5) | 1)))
    default:
        v7 := v6[-461]
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        if (0 >= -1) {
            var v9 int16 = -32768
            _ = v9
        }
        for fi10 := 0; fi10 < 20; fi10++ {
            if true {
                fmt.Println(v3)
                if true {
                    v2 -= (fi10 - (^(fi8)))
                    v11 := v6[(^(fi10))]
                    _ = v11
                } else {
                    fmt.Println(v1)
                    switch "\n\t" {
                    case "a":
                        v12 := GenS1{F0: int16(v5), F1: 18446744073709551615}
                        _ = v12
                    default:
                        if true {
                            v5 = (^(v5))
                            var v13 int64 = 0
                            _ = v13
                        } else {
                            v5 -= (v5 / (((v5 % ((4294967295) | 1))) | 1))
                        }
                    }
                    if true {
                        var v14 float64 = -2.25
                        if (!(!true)) {
                            var v15 string = "foo bar"
                            fmt.Println(v1)
                            var v16 float64 = (1e10 - v14)
                            _ = v15
                            _ = v16
                        }
                        fmt.Println(v1, v5)
                        _ = v14
                    } else {
                        _ = len(v4)
                    }
                }
            }
            _ = fi10
        }
        _ = fi8
    }
    var v17 float32 = float32(1.5)
    var v18 uint32 = (^(v5))
    v1 -= 1
    v19 := GenS1{F0: -517, F1: uint64(v18)}
    v6[(^(v2))] = (v2 ^ v2)
    var v20 bool = ((v2 >> ((18446744073709551615) & 63)) != 2147483647)
    v17 = 0.1
    var v21 int64 = -248
    _ = len(v6)
    if (true || (!false)) {
        v3 += (v3 & v3)
    } else {
        fmt.Println(v18, v5, v5)
    }
    var v22 int16 = 32767
    v23 := GenS2{F0: (v22 / ((1000) | 1)), F1: 42}
    v20 = (v20 && (!((v20 || (!((4294967295 <= 42) && v20))) || (!((uint16(v1) <= 65535) && true)))))
    v21 = v21
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
