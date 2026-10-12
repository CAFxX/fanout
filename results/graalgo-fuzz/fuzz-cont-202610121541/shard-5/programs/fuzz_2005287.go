package main

import (
    "fmt"
)

func genFn0(p0 uint16, p1 uint32) uint32 {
    switch "a" {
    case "a":
        var v4 string = "foo bar"
        _ = v4
    default:
        p0 = uint16(p1)
    }
    fmt.Println(p1)
    var v5 int32 = -1
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 uint32 = 2147483648
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := []int{fi8, fi8, (fi8 << ((0) & 63))}
        var v10 int64 = 0
        for fi11 := 0; fi11 < 3; fi11++ {
            fmt.Println(p1)
            if false {
                if len(v9) > 0 {
                    v9[0] = fi11
                }
            } else {
                var v12 uint16 = p0
                if ((^(v12)) <= v12) {
                    fmt.Println(v5, fi8, p1)
                    v12 = 255
                } else {
                    p1 += p1
                    if len(v9) > 0 {
                        v9[0] = (fi8 - fi11)
                    }
                }
                v12 = (p0 << ((1) & 63))
                _ = v12
            }
            if false {
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            } else {
                var v13 string = "foo bar"
                _ = v13
            }
            _ = fi11
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    return 42
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 1
    var v3 float32 = (-2.25 * 1.5)
    var v14 uint16 = 1
    fmt.Println(v14, v2, v3)
    fmt.Println(v3, v2)
    v2 = v2
    if (false && (false || false)) {
        v2 = v2
        fmt.Println(v2)
    }
    var v15 int8 = -45
    var v16 string = ("" + "a")
    var v17 bool = (!(!false))
    fmt.Println(v3, v15, v14)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
