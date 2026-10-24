package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int32 }
func genFn0() uint {
    var v5 float64 = -2.25
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS1{F0: ("a" + "hello")}
        _ = fi6
        _ = v7
    }
    _ = v5
    return 709
}

func main() {
    var v1 int8 = -3
    v2 := []string{"foo bar", "a"}
    v3 := make(map[int]uint32)
    v4 := []int{127, 0, 42}
    var v8 int8 = (v1 % (((v1 - v1)) | 1))
    switch 128 {
    case -9223372036854775808:
        v2 = append(v2, ("a" + ""))
    case 32768:
        for fi9 := 0; fi9 < 19; fi9++ {
            if (false || false) {
                for fi10 := 0; fi10 < 2; fi10++ {
                    v11 := GenS1{F0: ""}
                    v1 -= -14
                    v1 = -1
                    _ = fi10
                    _ = v11
                }
            } else {
                v12 := GenS2{F0: (-(v8)), F1: fi9}
                _ = v12
            }
            _ = fi9
        }
    default:
        if true {
            if true {
                var v13 int8 = (v8 >> ((4294967295) & 63))
                v14 := GenS2{F0: v13, F1: 124}
                v13 = v1
                _ = v13
                _ = v14
            }
        } else {
            _ = len(v3)
        }
    }
    var v15 int16 = 1
    var v16 bool = true
    for fi17 := 0; fi17 < 2; fi17++ {
        v18 := GenS3{F0: 0}
        v1 = (v8 - -1)
        _ = fi17
        _ = v18
    }
    delete(v3, 2147483647)
    var v19 int8 = (^(v1))
    var v20 string = "foo bar"
    if (689 > 65535) {
        v21 := GenS2{F0: (v19 & v19), F1: int(v15)}
        _ = v21
    }
    var v22 uint8 = 127
    v15 += v15
    var v23 int8 = (v19 & -1)
    v1 += 100
    v24 := GenS1{F0: ("" + v20)}
    var v25 uint = 18446744073709551615
    fmt.Println(v23, v19)
    fmt.Println(v23, v19)
    delete(v3, -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v8
}
