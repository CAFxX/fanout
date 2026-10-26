package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 2147483647
    var v3 uint = 56
    v4 := []uint32{268, 2, 412, 42}
    for fi5 := 0; fi5 < 3; fi5++ {
        switch (-(fi5)) {
        case -959:
            for fi6 := 0; fi6 < 11; fi6++ {
                v2 = v2
                switch 42 {
                case 42:
                    var v7 float64 = 0.0
                    _ = v7
                case -2147483648:
                    v8 := v1[0]
                    _ = v8
                }
                v2 = v2
                _ = fi6
            }
        case 969:
            v9 := GenS1{F0: ("\n\t" + "hello"), F1: 1, F2: (1 != 2147483648)}
            _ = v9
        }
        var v10 uint16 = 914
        _ = fi5
        _ = v10
    }
    v3 = v3
    v2 -= (v2 ^ v2)
    v11 := GenS1{F0: "hello", F1: 1, F2: true}
    v2 -= v2
    var v12 int64 = -9223372036854775808
    fmt.Println(v12, v3)
    switch 42 {
    case -1:
        if (false || (((306 != 4294967295) || (-390 <= 385)) && false)) {
            fmt.Println(v12, v3, v2)
            for fi13 := 0; fi13 < 16; fi13++ {
                if len(v4) > 0 {
                    v4[0] = 755
                }
                var v14 bool = (100 >= 54)
                delete(v1, fi13)
                _ = fi13
                _ = v14
            }
        }
    case 632:
        var v15 int16 = 42
        _ = v15
    case 2147483647:
        switch 32767 {
        case 4294967295:
            delete(v1, 42)
        case -1:
            if ((true || ((v2 << ((v3) & 63)) > (v2 | (v2 / ((1) | 1))))) || false) {
                for fi16 := 0; fi16 < 5; fi16++ {
                    v2 -= v2
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi16
                }
                fmt.Println(v12)
            }
        case -9223372036854775808:
            v17 := v1[128]
            _ = v17
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
