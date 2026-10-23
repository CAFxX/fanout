package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]string)
    var v3 uint64 = 42
    var v4 uint64 = (v3 / (((v3 % (((^(v3))) | 1))) | 1))
    var v5 int64 = 42
    v4 = (v3 * v3)
    v6 := GenS1{F0: 39, F1: 41}
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 uint16 = 1
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    if true {
        fmt.Println(v3, v4, v4)
    }
    fmt.Println(v5)
    var v10 int64 = (v5 / ((v5) | 1))
    v3 = v3
    v10 -= 42
    switch "hello" {
    case "hello":
        if true {
            if (3.14159 == -2.25) {
                for fi11 := 0; fi11 < 2; fi11++ {
                    switch (fi11 / ((fi11) | 1)) {
                    case -492:
                        if true {
                            v10 -= (v10 + v10)
                        }
                    case 127:
                        if (!false) {
                            if (2147483648 == 129) {
                                var v12 bool = ((fi11 ^ int(v3)) <= (fi11 * fi11))
                                _ = v12
                            }
                        }
                    }
                    _ = fi11
                }
            }
        } else {
            switch 42 {
            case -566:
                var v13 uint8 = 2
                _ = v13
            case 4294967295:
                if true {
                    var v14 int = 128
                    v10 += 9223372036854775807
                    var v15 int16 = 32767
                    _ = v14
                    _ = v15
                }
            case -1:
                if false {
                    v10 = (v10 & (v5 & (v10 / ((0) | 1))))
                }
            }
            fmt.Println(v4, v5, v5)
        }
    }
    fmt.Println(v5, v10)
    v5 -= (^(v5))
    var v16 uint8 = 119
    fmt.Println(v3, v4)
    var v17 int32 = -1
    var v18 int = 127
    var v19 uint64 = 973
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
