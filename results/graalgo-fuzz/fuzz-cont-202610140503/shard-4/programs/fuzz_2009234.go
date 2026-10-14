package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int32 }
func genFn0(p0 string, p1 uint32, p2 int64) uint8 {
    var v6 int64 = (p2 >> ((1) & 63))
    v7 := GenS1{F0: 1, F1: ("foo bar" + ""), F2: (("\n\t" + "\n\t") + "a")}
    fmt.Println(p2, p1)
    for fi8 := 0; fi8 < 4; fi8++ {
        p1 = p1
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 128
}

func main() {
    var v1 int64 = 42
    var v2 int8 = -1
    var v3 string = "foo bar"
    v4 := make(map[int]uint32)
    var v5 uint16 = 0
    var v9 int32 = 915
    v9 = (v9 >> ((42) & 63))
    var v10 int64 = 42
    if false {
        v11 := GenS1{F0: 2, F1: "", F2: v3}
        v12 := GenS2{F0: -331}
        _ = v11
        _ = v12
    }
    fmt.Println(v9, v3, v1)
    var v13 uint32 = 1
    var v14 uint8 = 128
    fmt.Println(v2, v5, v1)
    v15 := GenS2{F0: 4294967295}
    var v16 int = 42
    var v17 string = (v3 + (("a" + v3) + ((v3 + "hello") + v3)))
    var v18 uint8 = 1
    v19 := GenS3{F0: (v14 << ((341) & 63)), F1: (v9 / ((-1) | 1))}
    var v20 uint32 = v13
    var v21 int8 = (v2 | -1)
    if false {
        if (892 > 255) {
            v14 = v14
            if false {
                for fi22 := 0; fi22 < 5; fi22++ {
                    var v23 float32 = (1e10 + 3.14159)
                    var v24 int16 = 34
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                v13 -= 62
            }
        }
        v25 := GenS3{F0: (v18 / ((0) | 1)), F1: -1}
        v9 = v9
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
