package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int64, p1 int16, p2 int64) int {
    fmt.Println(p0)
    switch ("a" + (("foo bar" + "a") + "a")) {
    case "hello":
        if true {
            for fi6 := 0; fi6 < 13; fi6++ {
                for fi7 := 0; fi7 < 11; fi7++ {
                    p0 = 315
                    var v8 bool = (0 != 0)
                    _ = fi7
                    _ = v8
                }
                _ = fi6
            }
            v9 := []int{}
            _ = v9
        }
    case "foo bar":
        p0 = p0
    }
    switch ("hello" + "foo bar") {
    case "a":
        p0 = 1
    }
    v10 := GenS1{F0: ("a" + "foo bar"), F1: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 2147483647
}
func genFn1(p0 string, p1 int32) bool {
    v11 := []int{127, 382, 9223372036854775807, 1}
    p1 = (p1 - 0)
    _ = p0
    _ = p1
    _ = v11
    return (true || (255 < 586))
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]uint32)
    var v3 int = 128
    var v4 bool = (v3 < -2147483648)
    v5 := make(map[int]int)
    v12 := GenS2{F0: v3, F1: 4294967295}
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 uint = 0
        v15 := GenS1{F0: "foo bar", F1: 1}
        fmt.Println(v4, v4, v4)
        _ = fi13
        _ = v14
        _ = v15
    }
    if (255 <= 216) {
        delete(v5, (-(v3)))
        delete(v2, 0)
        v16 := []int{}
        _ = v16
    } else {
        v17 := GenS2{F0: v3, F1: 0}
        v18 := GenS3{F0: 65535}
        if (1.5 <= (-2.25 - 0.0)) {
            v19 := GenS4{F0: 780}
            _ = v19
        } else {
            var v20 int16 = 1000
            _ = v20
        }
        _ = v17
        _ = v18
    }
    switch (^(v3)) {
    case 9223372036854775807:
        var v21 float64 = ((1.5 * float64(3.14159)) + (3.14159 + 0.0))
        _ = v21
    case 0:
        v2[4294967295] = 1
    }
    var v22 string = "a"
    v23 := []int{2147483647, genFn0(1, 1000, 9223372036854775807), (v3 + (v3 / ((genFn0(9223372036854775807, 1000, -833)) | 1))), (v3 % ((2147483647) | 1))}
    v22 = "a"
    var v24 uint16 = 1
    v25 := GenS3{F0: v24}
    v24 -= (v24 >> ((179) & 63))
    v5[(^(v3))] = (^(v3))
    for fi26 := 0; fi26 < 11; fi26++ {
        v24 = (v24 & 65535)
        v22 = "a"
        _ = fi26
    }
    if (42 != -32768) {
        fmt.Println(v3, v24)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
