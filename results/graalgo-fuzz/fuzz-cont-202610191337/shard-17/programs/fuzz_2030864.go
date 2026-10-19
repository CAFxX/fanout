package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }

func main() {
    var v1 uint = 326
    var v2 uint8 = 128
    v3 := make(map[string]string)
    v4 := GenS1{F0: (v1 & (v1 | v1)), F1: ""}
    for fi5 := 0; fi5 < 15; fi5++ {
        if (-28 != -128) {
            if (!((!true) && (!true))) {
                switch 68 {
                case -9223372036854775808:
                    v6 := v3["a"]
                    _ = v6
                default:
                    var v7 bool = true
                    _ = v7
                }
                for fi8 := 0; fi8 < 12; fi8++ {
                    var v9 int16 = 548
                    _ = fi8
                    _ = v9
                }
            }
            _ = len(v3)
        } else {
            var v10 string = "foo bar"
            switch 2147483647 {
            case 32767:
                v11 := GenS1{F0: (v1 - (v1 / ((4294967295) | 1))), F1: v10}
                _ = v11
            }
            v12 := GenS1{F0: v1, F1: v10}
            _ = v10
            _ = v12
        }
        v2 -= 88
        v2 += v2
        _ = fi5
    }
    v1 -= 877
    for fi13 := 0; fi13 < 19; fi13++ {
        v2 -= (v2 + v2)
        _ = fi13
    }
    var v14 int32 = 42
    v1 -= (v1 * 244)
    for fi15 := 0; fi15 < 2; fi15++ {
        for fi16 := 0; fi16 < 18; fi16++ {
            v3["hello"] = "hello"
            v2 -= v2
            v17 := GenS1{F0: v1, F1: "a"}
            _ = fi16
            _ = v17
        }
        v18 := GenS1{F0: (v1 << ((v1) & 63)), F1: "\n\t"}
        v19 := []int{(fi15 >> ((1) & 63)), (^(fi15)), fi15, (fi15 << (((v1 + (v1 | 2))) & 63))}
        _ = fi15
        _ = v18
        _ = v19
    }
    v20 := GenS1{F0: 65535, F1: "foo bar"}
    if true {
        v14 += (v14 - int32(v1))
        v1 = (v1 ^ (v1 * v1))
        v3["foo bar"] = "a"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
