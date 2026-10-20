package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint8, p1 int64, p2 bool) uint {
    var v4 int32 = -1
    v5 := GenS1{F0: (v4 - v4)}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 0
}

func main() {
    var v1 uint32 = 195
    var v2 int64 = -402
    var v3 int32 = -1
    var v6 int32 = (v3 * 42)
    v7 := GenS1{F0: 28}
    v6 = (v6 / ((v3) | 1))
    var v8 uint = 0
    if (874 == -423) {
        fmt.Println(v8)
    } else {
        if (255 <= uint16(v1)) {
            v3 = v3
            switch ("" + "foo bar") {
            case "foo bar":
                v9 := GenS1{F0: (v6 + (^(v3)))}
                _ = v9
            case "a":
                v6 += v3
            default:
                for fi10 := 0; fi10 < 4; fi10++ {
                    var v11 int = 127
                    v12 := GenS2{F0: 680}
                    var v13 uint32 = v1
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
            }
        }
        v2 = (v2 << (((v8 | v8)) & 63))
        var v14 uint8 = 33
        _ = v14
    }
    v2 -= v2
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 int8 = -86
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 string = "\n\t"
            _ = fi17
            _ = v18
        }
        v1 -= v1
        _ = fi15
        _ = v16
    }
    var v19 bool = (2 >= 255)
    var v20 int = 32767
    v2 = (v2 * (v2 ^ (v2 ^ 0)))
    v1 = v1
    if (!(26 >= 232)) {
        v21 := []int{v20}
        if v19 {
            switch (v20 | -9223372036854775808) {
            case 757:
                fmt.Println(v19, v6)
            case 458:
                fmt.Println(v2, v3, v2)
            case 2147483647:
                fmt.Println(v8, v19)
            default:
                fmt.Println(v20, v1)
            }
            v3 -= v3
            if true {
                var v22 int32 = v3
                var v23 int8 = -19
                _ = v22
                _ = v23
            }
        } else {
            if (v1 != (v1 - v1)) {
                var v24 float64 = 1e10
                var v25 float64 = (v24 + -2.25)
                _ = v24
                _ = v25
            }
            var v26 uint32 = (v1 ^ (v1 + v1))
            _ = v26
        }
        _ = v21
    } else {
        fmt.Println(v2, v20, v3)
        v27 := GenS1{F0: (v3 / ((v3) | 1))}
        fmt.Println(v6)
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
