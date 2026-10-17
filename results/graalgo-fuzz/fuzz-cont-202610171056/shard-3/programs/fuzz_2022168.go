package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint16 }

func main() {
    var v1 int32 = 692
    var v2 float32 = 0.0
    var v3 uint32 = 42
    var v4 int = -681
    v5 := GenS1{F0: 9223372036854775807, F1: 1, F2: 1}
    fmt.Println(v1, v2, v1)
    if ((0 >= 0) || (!false)) {
        v4 = (v4 >> ((42) & 63))
    }
    fmt.Println(v3)
    var v6 int64 = 9223372036854775807
    v2 = v2
    v4 -= (v4 + (v4 & int(v6)))
    var v7 uint16 = 0
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 uint8 = 1
        var v10 float64 = -2.25
        for fi11 := 0; fi11 < 20; fi11++ {
            v12 := []int{}
            _ = fi11
            _ = v12
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    var v13 int8 = 0
    switch (v4 % ((349) | 1)) {
    case -889:
        switch v4 {
        case -292:
            v6 = int64(v13)
        case 2147483647:
            for fi14 := 0; fi14 < 18; fi14++ {
                if (v7 > 42) {
                    v15 := []int{v4, 0, -832}
                    _ = v15
                }
                _ = fi14
            }
        }
    case 586:
        switch (("a" + "hello") + ("hello" + ("\n\t" + "hello"))) {
        case "hello":
            if (!false) {
                for fi16 := 0; fi16 < 18; fi16++ {
                    if (9223372036854775807 <= 1) {
                        var v17 int32 = (^(v1))
                        var v18 uint8 = 222
                        if (true || ((v1 | (v1 + v17)) >= v1)) {
                            var v19 uint = 322
                            v20 := []int{}
                            _ = v19
                            _ = v20
                        }
                        _ = v17
                        _ = v18
                    }
                    var v21 int64 = v6
                    _ = fi16
                    _ = v21
                }
                var v22 int64 = 0
                _ = v22
            }
        default:
            var v23 uint32 = v3
            _ = v23
        }
    case 32768:
        for fi24 := 0; fi24 < 12; fi24++ {
            v4 = v4
            var v25 uint8 = 42
            _ = fi24
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
