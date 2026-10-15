package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint, p1 int16) uint32 {
    var v7 int64 = 0
    v8 := []int{32767, 9223372036854775807, int(p1)}
    p1 = 32767
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 182
}

func main() {
    v1 := []uint32{2, 55, 2, 65535}
    var v2 uint32 = 2147483648
    v3 := []int{}
    var v4 uint8 = 0
    var v5 uint = 18446744073709551615
    var v6 int = 128
    var v9 bool = true
    if ((3.14159 + 1e10) <= 1.5) {
        if (!v9) {
            fmt.Println(v6, v5)
        } else {
            v10 := GenS1{F0: "\n\t"}
            _ = v10
        }
        for fi11 := 0; fi11 < 19; fi11++ {
            v2 = v2
            _ = fi11
        }
    }
    var v12 uint16 = 32767
    for fi13 := 0; fi13 < 18; fi13++ {
        for fi14 := 0; fi14 < 9; fi14++ {
            v6 += (fi14 % ((v6) | 1))
            _ = fi14
        }
        if (!(v9 || v9)) {
            for fi15 := 0; fi15 < 7; fi15++ {
                switch v6 {
                case -9223372036854775808:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case -963:
                    v16 := GenS1{F0: ("a" + "hello")}
                    _ = v16
                default:
                    v4 -= 194
                }
                _ = fi15
            }
            for fi17 := 0; fi17 < 6; fi17++ {
                if v9 {
                    var v18 int64 = -9223372036854775808
                    _ = v18
                }
                var v19 uint64 = 609
                if len(v3) > 0 {
                    v3[0] = fi17
                }
                _ = fi17
                _ = v19
            }
        } else {
            var v20 uint = 149
            _ = v20
        }
        _ = fi13
    }
    v21 := GenS2{F0: -171}
    for fi22 := 0; fi22 < 11; fi22++ {
        v23 := GenS3{F0: v2}
        _ = fi22
        _ = v23
    }
    var v24 uint = v5
    v2 += (v2 - genFn0((^(v5)), 0))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
