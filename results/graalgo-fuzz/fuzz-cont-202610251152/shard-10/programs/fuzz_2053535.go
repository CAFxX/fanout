package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := make(map[int]string)
    var v2 uint = 213
    v3 := []string{"foo bar", "\n\t"}
    v4 := []uint32{4294967295, 4294967295}
    v5 := make(map[uint32]bool)
    var v6 string = ("foo bar" + ("\n\t" + "a"))
    if true {
        if (311 < 32767) {
            for fi7 := 0; fi7 < 17; fi7++ {
                var v8 string = "a"
                var v9 int64 = -871
                _ = fi7
                _ = v8
                _ = v9
            }
            for fi10 := 0; fi10 < 19; fi10++ {
                var v11 uint16 = 600
                switch -364 {
                case -9223372036854775808:
                    v12 := GenS1{F0: (-(fi10))}
                    _ = v12
                case 1:
                    if true {
                        for fi13 := 0; fi13 < 10; fi13++ {
                            v14 := GenS1{F0: 32767}
                            v15 := GenS2{F0: 2}
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    } else {
                        var v16 float32 = float32(float64(float32((((1e10 / 1.0) / float64(1.5)) * 1.5))))
                        _ = v16
                    }
                default:
                    v11 -= v11
                }
                _ = fi10
                _ = v11
            }
        }
    } else {
        v17 := GenS3{F0: 0}
        _ = v17
    }
    var v18 uint16 = 1
    var v19 string = ((v6 + "hello") + v6)
    _ = len(v1)
    switch v6 {
    case "\n\t":
        var v20 uint64 = 2
        _ = v20
    default:
        var v21 int = -2147483648
        _ = v21
    }
    fmt.Println(v6)
    v22 := v1[0]
    v2 += (v2 | v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
