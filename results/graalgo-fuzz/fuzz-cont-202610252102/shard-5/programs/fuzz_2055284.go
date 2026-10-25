package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() int32 {
    if true {
        var v4 int = 4294967295
        if (0 < 950) {
            v4 = v4
            var v5 uint8 = 176
            for fi6 := 0; fi6 < 12; fi6++ {
                v4 -= fi6
                v7 := GenS1{F0: v5}
                _ = fi6
                _ = v7
            }
            _ = v5
        }
        var v8 uint64 = 18446744073709551615
        _ = v4
        _ = v8
    }
    var v9 int32 = 47
    switch -526 {
    case -7:
        if false {
            v10 := []int{506, 4294967295, 32767, 9223372036854775807}
            _ = v10
        }
    default:
        if true {
            for fi11 := 0; fi11 < 19; fi11++ {
                fmt.Println(fi11)
                var v12 float32 = (1.5 + float32((0.0 * 0.0)))
                _ = fi11
                _ = v12
            }
        } else {
            v9 = (v9 / ((-2147483648) | 1))
        }
    }
    var v13 uint32 = 42
    _ = v13
    _ = v9
    return (v9 ^ (^(v9)))
}

func main() {
    var v1 uint64 = 289
    v2 := make(map[uint32]int)
    v3 := make(map[string]int)
    v14 := GenS1{F0: 255}
    if (1e10 <= (float32((-2.25 - 1e10)) + 3.14159)) {
        fmt.Println(v1)
    }
    var v15 int64 = 42
    v15 = (v15 | v15)
    var v16 int = -196
    v15 = (v15 + (v15 & v15))
    for fi17 := 0; fi17 < 20; fi17++ {
        switch ("" + "hello") {
        case "hello":
            switch "hello" {
            case "a":
                switch "foo bar" {
                case "a":
                    switch v16 {
                    case -9223372036854775808:
                        var v18 uint64 = (v1 | 42)
                        _ = v18
                    case 373:
                        if ((639 <= 1) || (0 <= 4294967295)) {
                            v2[888] = v16
                        }
                    default:
                        delete(v2, 4294967295)
                    }
                case "\n\t":
                    v16 -= (fi17 & fi17)
                default:
                    v19 := []int{fi17}
                    _ = v19
                }
            case "foo bar":
                v1 = 4294967295
            }
        }
        _ = fi17
    }
    v20 := []int{(v16 / ((-147) | 1)), v16, v16, -821}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
}
