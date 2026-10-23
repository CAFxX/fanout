package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := []string{}
    v2 := make(map[uint32]bool)
    var v3 int = 127
    v4 := make(map[uint32]string)
    v5 := []int{v3, v3, v3}
    v6 := GenS1{F0: v3}
    v3 -= (v3 | (v3 | 32767))
    v3 = -619
    var v7 uint = 255
    switch 4294967295 {
    case 943:
        for fi8 := 0; fi8 < 5; fi8++ {
            if false {
                v9 := v2[1]
                var v10 string = (("\n\t" + ("\n\t" + "")) + "")
                _ = v10
                _ = v9
            } else {
                for fi11 := 0; fi11 < 19; fi11++ {
                    var v12 int8 = -1
                    var v13 int8 = (^(v12))
                    fmt.Println(fi8)
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                if (false && true) {
                    v3 = (v3 | (-(fi8)))
                }
            }
            fmt.Println(v7, v3, v3)
            v14 := GenS1{F0: 246}
            _ = fi8
            _ = v14
        }
    case -9223372036854775808:
        for fi15 := 0; fi15 < 6; fi15++ {
            if (128 < uint8(v7)) {
                v16 := GenS1{F0: fi15}
                _ = v16
            }
            _ = fi15
        }
    case 2147483647:
        if false {
            var v17 int32 = -253
            switch v3 {
            case 42:
                fmt.Println(v17)
            default:
                v18 := GenS1{F0: v3}
                _ = v18
            }
            _ = v17
        }
    }
    v3 = v3
    if true {
        for fi19 := 0; fi19 < 9; fi19++ {
            v7 += (v7 + (v7 * v7))
            _ = fi19
        }
        v7 -= 65535
        var v20 int64 = 42
        _ = v20
    }
    v3 += (v3 * v3)
    if true {
        for fi21 := 0; fi21 < 1; fi21++ {
            v3 -= 9223372036854775807
            _ = fi21
        }
    } else {
        v7 = v7
    }
    var v22 uint = (v7 * (v7 & (v7 % ((0) | 1))))
    v22 -= 2
    v3 += v3
    var v23 int = v3
    v24 := GenS1{F0: -4}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
