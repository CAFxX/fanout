package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint64, p1 int16) int16 {
    v7 := []int{}
    var v8 int32 = -674
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p1
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]bool)
    v3 := []int{-1, 794}
    v4 := []int{2147483647, -2147483648, -2147483648, -2147483648}
    var v5 uint8 = 42
    var v6 float32 = 0.1
    v6 -= (v6 * v6)
    if ((!(("foo bar" == ("\n\t" + "")) && (true && false))) && true) {
        if (("foo bar" + "hello") < "foo bar") {
            var v9 uint16 = 1
            v10 := GenS1{F0: v9}
            for fi11 := 0; fi11 < 14; fi11++ {
                for fi12 := 0; fi12 < 14; fi12++ {
                    if len(v4) > 0 {
                        v4[0] = (-(fi12))
                    }
                    _ = fi12
                }
                if false {
                    var v13 int = (fi11 | fi11)
                    _ = v13
                }
                _ = fi11
            }
            _ = v10
            _ = v9
        } else {
            fmt.Println(v5, v6)
            var v14 float64 = 0.1
            v15 := GenS1{F0: 255}
            _ = v14
            _ = v15
        }
        v16 := GenS2{F0: (42 < 368)}
        v3 = append(v3, 2147483647)
        _ = v16
    } else {
        v17 := GenS2{F0: true}
        _ = v17
    }
    var v18 float32 = v6
    var v19 uint = 255
    v20 := GenS1{F0: 708}
    v2[127] = false
    for fi21 := 0; fi21 < 3; fi21++ {
        v22 := v1[fi21]
        if false {
            v23 := v1[2147483647]
            _ = v23
        } else {
            if true {
                v18 = float32(float64((1.5 + float32(-2.25))))
            }
            var v24 uint64 = 1
            _ = v24
        }
        v5 = 2
        _ = fi21
        _ = v22
    }
    _ = len(v3)
    fmt.Println(v18, v5, v18)
    v25 := GenS3{F0: 2}
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
