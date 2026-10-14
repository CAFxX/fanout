package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 uint32 }
func genFn0(p0 int, p1 string) int16 {
    fmt.Println(p1)
    v5 := []int{(p0 / ((p0) | 1)), 32767}
    if len(v5) > 0 {
        v5[0] = p0
    }
    _ = p0
    _ = p1
    _ = v5
    return 42
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[uint32]int)
    v3 := make(map[string]bool)
    var v4 int = 42
    if true {
        fmt.Println(v1)
        if false {
            var v6 uint32 = 2147483647
            v7 := []int{v4}
            _ = v6
            _ = v7
        } else {
            for fi8 := 0; fi8 < 14; fi8++ {
                v9 := v3["\n\t"]
                _ = fi8
                _ = v9
            }
            v4 -= v4
            v10 := []int{(v4 - -858), v4, v4, (v4 - -958)}
            _ = v10
        }
        fmt.Println(v4, v1)
    } else {
        var v11 float32 = (0.0 + (3.14159 / ((-2.25 / (3.14159 / 0.1)) * float32(1e10))))
        _ = v11
    }
    var v12 int64 = -1
    var v13 uint8 = 2
    if (!(true || false)) {
        for fi14 := 0; fi14 < 3; fi14++ {
            for fi15 := 0; fi15 < 13; fi15++ {
                v16 := GenS1{F0: (v13 - (v1 - (v13 % ((v1) | 1))))}
                delete(v3, "a")
                var v17 uint32 = 65535
                v17 = v17
                _ = fi15
                _ = v16
                _ = v17
            }
            if true {
                var v18 int8 = 42
                v19 := GenS2{F0: 2}
                _ = v18
                _ = v19
            }
            _ = fi14
        }
        for fi20 := 0; fi20 < 17; fi20++ {
            for fi21 := 0; fi21 < 7; fi21++ {
                switch 32768 {
                case -9223372036854775808:
                    v4 = fi20
                }
                if true {
                    v12 -= 497
                    var v22 bool = (v1 >= 2)
                    _ = v22
                }
                switch "foo bar" {
                case "hello":
                    v13 = v13
                case "foo bar":
                    if true {
                        fmt.Println(v13)
                    }
                }
                _ = fi21
            }
            switch (v4 * (v4 / (((v4 % ((v4) | 1))) | 1))) {
            case 42:
                var v23 int64 = int64(v13)
                _ = v23
            }
            _ = len(v3)
            _ = fi20
        }
        v24 := []int{v4, v4, v4, (v4 * (v4 << ((2) & 63)))}
        _ = v24
    }
    var v25 uint32 = 0
    v26 := GenS2{F0: 352}
    var v27 uint32 = 4294967295
    if (!((v1 != 127) && true)) {
        fmt.Println(v1)
        v12 = v12
    }
    v28 := GenS3{F0: -2147483648, F1: 2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
