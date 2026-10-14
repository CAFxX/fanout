package main

import (
    "fmt"
)

func genFn0(p0 int32, p1 uint) int32 {
    v5 := []int{-1, -1, 127}
    for fi6 := 0; fi6 < 15; fi6++ {
        p1 = p1
        _ = fi6
    }
    p1 = (^(p1))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = p0
    _ = p1
    _ = v5
    return p0
}
func genFn1(p0 uint) int8 {
    p0 -= p0
    fmt.Println(p0)
    _ = p0
    return 92
}

func main() {
    v1 := []string{"\n\t", "a", "a"}
    v2 := []int{-602}
    var v3 string = ("" + "foo bar")
    var v4 uint8 = 255
    var v7 uint = 65535
    var v8 uint8 = v4
    fmt.Println(v7)
    _ = len(v1)
    fmt.Println(v4, v3)
    if len(v2) > 0 {
        v2[0] = -9223372036854775808
    }
    fmt.Println(v3, v4, v4)
    var v9 uint = v7
    switch v3 {
    case "a":
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 uint32 = 2147483648
            _ = fi10
            _ = v11
        }
    default:
        if (!((v4 | 128) == v4)) {
            for fi12 := 0; fi12 < 17; fi12++ {
                if false {
                    var v13 float64 = 1.5
                    _ = v13
                }
                if ((v3 + (v3 + "\n\t")) != v3) {
                    for fi14 := 0; fi14 < 2; fi14++ {
                        var v15 string = v3
                        var v16 int = (fi12 << (((v9 * v9)) & 63))
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                } else {
                    for fi17 := 0; fi17 < 17; fi17++ {
                        fmt.Println(v7, fi12)
                        var v18 int8 = -1
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v19 int16 = 42
                        if true {
                            v9 -= (^(v9))
                        }
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                }
                _ = fi12
            }
            var v20 uint64 = 4294967295
            var v21 uint = (v9 << ((285) & 63))
            _ = v20
            _ = v21
        } else {
            var v22 uint8 = v8
            _ = v22
        }
    }
    v9 += uint(v4)
    _ = len(v1)
    var v23 uint64 = 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
