package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint32 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := []string{}
    v2 := make(map[int]string)
    var v3 uint64 = 775
    var v4 int = 127
    var v5 int64 = 475
    for fi6 := 0; fi6 < 18; fi6++ {
        v4 -= (v4 & (fi6 * -1))
        _ = fi6
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: fi7, F1: (fi7 << ((1) & 63)), F2: 653}
        var v9 uint32 = 65535
        var v10 bool = (int8(v9) != 11)
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    v11 := GenS1{F0: (v4 | (v4 >> ((2) & 63))), F1: 2147483647, F2: 0}
    v12 := GenS1{F0: 2147483647, F1: v4, F2: 902}
    if false {
        for fi13 := 0; fi13 < 6; fi13++ {
            v14 := GenS2{F0: 720}
            v4 += 32768
            switch (fi13 & -9223372036854775808) {
            case 128:
                v4 = -9223372036854775808
            }
            _ = fi13
            _ = v14
        }
        fmt.Println(v5, v5, v4)
        for fi15 := 0; fi15 < 19; fi15++ {
            if true {
                var v16 float32 = float32(3.14159)
                for fi17 := 0; fi17 < 4; fi17++ {
                    fmt.Println(v5, v4)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v18 int16 = 42
                    _ = fi17
                    _ = v18
                }
                v2[v4] = "a"
                _ = v16
            }
            _ = fi15
        }
    } else {
        for fi19 := 0; fi19 < 2; fi19++ {
            if (!false) {
                if (!false) {
                    var v20 uint16 = 320
                    switch v4 {
                    case 32767:
                        var v21 uint16 = (v20 | 102)
                        _ = v21
                    }
                    _ = v20
                }
            } else {
                var v22 int16 = 42
                _ = v22
            }
            _ = fi19
        }
    }
    v3 = 42
    var v23 uint32 = 189
    switch (-(v4)) {
    case -801:
        if false {
            fmt.Println(v23)
            fmt.Println(v3)
        }
    case 2147483647:
        for fi24 := 0; fi24 < 16; fi24++ {
            for fi25 := 0; fi25 < 9; fi25++ {
                fmt.Println(v5)
                _ = fi25
            }
            _ = fi24
        }
    default:
        for fi26 := 0; fi26 < 11; fi26++ {
            fmt.Println(v5, fi26)
            _ = fi26
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
