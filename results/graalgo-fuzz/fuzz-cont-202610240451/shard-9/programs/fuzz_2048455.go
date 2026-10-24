package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int64) uint {
    fmt.Println(p0)
    p0 = 621
    if true {
        v4 := []int{954, 523}
        _ = v4
    }
    _ = p0
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 1
    v3 := make(map[uint32]uint32)
    v5 := GenS1{F0: 300}
    v2 -= 65535
    if true {
        v2 += 2
    } else {
        var v6 int8 = 100
        v1[2] = (0.0 == -2.25)
        _ = v6
    }
    v7 := GenS1{F0: 0}
    for fi8 := 0; fi8 < 11; fi8++ {
        if (!false) {
            v2 = (v2 / (((v2 * v2)) | 1))
            v9 := GenS1{F0: 0}
            v10 := GenS1{F0: -890}
            _ = v10
            _ = v9
        } else {
            v2 += v2
        }
        v2 += (^(v2))
        _ = fi8
    }
    v11 := GenS1{F0: -1}
    var v12 float32 = (((float32(1e10) * 1.5) / 1e10) * -2.25)
    switch "" {
    case "":
        if false {
            switch 32768 {
            case 42:
                for fi13 := 0; fi13 < 13; fi13++ {
                    var v14 bool = false
                    fmt.Println(v12, v2)
                    _ = fi13
                    _ = v14
                }
            case 1:
                switch (("\n\t" + ("foo bar" + ("\n\t" + ""))) + (("\n\t" + "a") + ("" + "a"))) {
                case "":
                    var v15 string = "a"
                    _ = v15
                case "a":
                    v2 = v2
                }
            default:
                var v16 float32 = ((v12 - v12) / 1.0)
                _ = v16
            }
            var v17 float32 = 1.5
            if true {
                v17 = v12
                v18 := []int{160, 2147483647, 32767, -1}
                _ = v18
            }
            _ = v17
        } else {
            v19 := GenS1{F0: 42}
            _ = v19
        }
    default:
        fmt.Println(v12, v2)
    }
    v2 -= v2
    v2 += (^(v2))
    var v20 int8 = 0
    if (9223372036854775807 == 9223372036854775807) {
        switch "a" {
        case "hello":
            v21 := GenS2{F0: 128}
            _ = v21
        default:
            v22 := v3[v2]
            _ = v22
        }
        v2 = 60
    } else {
        var v23 uint8 = 128
        v20 = v20
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 uint16 = 1
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v5
    _ = v7
}
