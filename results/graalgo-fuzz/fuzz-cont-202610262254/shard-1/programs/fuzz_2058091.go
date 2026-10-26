package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint16, p1 int32, p2 int64) int64 {
    var v5 int32 = (p1 | p1)
    v6 := []int{-940, -9223372036854775808}
    if (float32(0.0) >= 1e10) {
        v7 := GenS1{F0: 242}
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 float32 = float32(0.1)
            if len(v6) > 0 {
                v6[0] = fi8
            }
            switch "" {
            case "":
                switch fi8 {
                case 289:
                    p2 = p2
                case 1:
                    switch ((("\n\t" + "hello") + "\n\t") + "foo bar") {
                    case "hello":
                        var v10 int64 = (p2 | p2)
                        _ = v10
                    case "foo bar":
                        v11 := GenS1{F0: 52}
                        _ = v11
                    case "\n\t":
                        v12 := GenS2{F0: -2147483648}
                        _ = v12
                    }
                case 6:
                    p1 = p1
                default:
                    v13 := GenS2{F0: v5}
                    _ = v13
                }
            case "hello":
                v5 = 0
            case "a":
                var v14 uint = 65535
                _ = v14
            }
            _ = fi8
            _ = v9
        }
        _ = v7
    } else {
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 string = "a"
            var v17 string = "hello"
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p2
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    var v3 string = ("\n\t" + (("\n\t" + "a") + "a"))
    var v4 bool = true
    v18 := []int{634, 32767}
    var v19 int64 = 1
    fmt.Println(v3, v4)
    v20 := v1[2]
    var v21 uint = 548
    var v22 int16 = 0
    v19 -= (^(v19))
    v23 := GenS3{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
