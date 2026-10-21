package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 bool) int8 {
    v7 := []int{}
    fmt.Println(p0)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = v7
    return 127
}
func genFn1(p0 int, p1 uint32, p2 int32) int32 {
    v8 := []int{0}
    v9 := GenS1{F0: 548}
    v10 := GenS2{F0: (p1 / ((42) | 1))}
    v8 = append(v8, 9223372036854775807)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return p2
}

func main() {
    var v1 int16 = -1
    var v2 int8 = -38
    v3 := make(map[int]string)
    var v4 uint16 = 677
    var v5 int32 = 0
    var v6 int32 = (^(v5))
    v11 := []int{-439, 127}
    var v12 float64 = -2.25
    switch "" {
    case "":
        fmt.Println(v6)
    case "hello":
        v6 = v6
    }
    switch -266 {
    case 0:
        fmt.Println(v5)
    case 9223372036854775807:
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 int16 = v1
            if len(v11) > 0 {
                v11[0] = fi13
            }
            _ = fi13
            _ = v14
        }
    case 128:
        fmt.Println(v12)
    default:
        v15 := GenS3{F0: 65535}
        _ = v15
    }
    switch "a" {
    case "hello":
        v1 = v1
    case "":
        v2 += v2
    case "\n\t":
        v16 := v3[1]
        _ = v16
    }
    var v17 bool = ((!(v12 <= v12)) && (269 < 42))
    var v18 uint32 = 7
    v19 := GenS1{F0: v1}
    var v20 string = ("" + ("a" + ("\n\t" + ("\n\t" + "a"))))
    var v21 float64 = float64(3.14159)
    if (!v17) {
        if ((v5 >= (v6 - v5)) || (false && ((v6 ^ (v6 - -688)) > -1))) {
            if len(v11) > 0 {
                v11[0] = 850
            }
            if true {
                if len(v11) > 0 {
                    v11[0] = -475
                }
            }
            if (v2 <= -119) {
                v22 := GenS2{F0: (^(v18))}
                v12 = (-2.25 / (v21 / v21))
                _ = v22
            } else {
                var v23 int32 = 350
                var v24 int = 4294967295
                _ = v23
                _ = v24
            }
        } else {
            for fi25 := 0; fi25 < 17; fi25++ {
                v18 = (v18 + v18)
                _ = fi25
            }
            var v26 bool = (((255 != 42) && false) && v17)
            _ = v26
        }
        var v27 uint = 18446744073709551615
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
