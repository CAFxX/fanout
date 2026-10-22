package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int32, p1 int8) int32 {
    var v7 int8 = p1
    v8 := []int{832, -239}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -2147483648
}

func main() {
    var v1 float32 = float32((float64(-2.25) / 1.5))
    v2 := []string{"hello", ""}
    var v3 uint = 4294967295
    v4 := make(map[string]uint32)
    var v5 int32 = -173
    var v6 uint16 = 255
    v2 = append(v2, (("a" + "hello") + ("hello" + ("a" + "foo bar"))))
    _ = len(v2)
    v5 += (v5 * genFn0(genFn0(0, 100), 43))
    switch 264 {
    case 1:
        for fi9 := 0; fi9 < 2; fi9++ {
            if true {
                var v10 uint64 = 4294967295
                switch (fi9 & -501) {
                case 2147483647:
                    var v11 bool = true
                    _ = v11
                case 1:
                    if len(v2) > 0 {
                        v2[0] = ("\n\t" + ("\n\t" + ""))
                    }
                case 127:
                    _ = len(v2)
                default:
                    if (genFn0(-2147483648, 42) != (^(v5))) {
                        for fi12 := 0; fi12 < 14; fi12++ {
                            fmt.Println(fi9, v10)
                            _ = fi12
                        }
                    }
                }
                for fi13 := 0; fi13 < 13; fi13++ {
                    fmt.Println(v3)
                    if ((v1 - v1) <= (1e10 + 3.14159)) {
                        fmt.Println(fi13, v10)
                        var v14 uint64 = (v10 << (((^(v3))) & 63))
                        _ = v14
                    }
                    _ = fi13
                }
                _ = v10
            } else {
                v5 = (v5 ^ v5)
                var v15 float32 = ((v1 - v1) / (-2.25 - v1))
                _ = v15
            }
            _ = len(v2)
            fmt.Println(v1)
            _ = fi9
        }
    case 9223372036854775807:
        var v16 uint32 = 462
        _ = v16
    default:
        if false {
            for fi17 := 0; fi17 < 10; fi17++ {
                var v18 float64 = 3.14159
                if (42 != -376) {
                    fmt.Println(v3)
                    fmt.Println(v5)
                    _ = len(v4)
                }
                v19 := GenS1{F0: 9223372036854775807}
                _ = fi17
                _ = v18
                _ = v19
            }
            var v20 int32 = int32(v6)
            _ = v20
        }
    }
    v21 := GenS1{F0: 58}
    var v22 uint = (v3 ^ (v3 - (^(v3))))
    fmt.Println(v1, v5)
    for fi23 := 0; fi23 < 12; fi23++ {
        switch 615 {
        case -1:
            var v24 int = 222
            _ = v24
        }
        _ = fi23
    }
    v3 -= (^(v3))
    v1 = (v1 - ((0.1 - (v1 - ((((v1 - float32(0.1)) / (v1 * ((v1 - v1) + (v1 - 1.5)))) * (v1 / ((v1 * v1) - 1.5))) + ((v1 / v1) * (v1 + v1))))) + (v1 / 1.0)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
