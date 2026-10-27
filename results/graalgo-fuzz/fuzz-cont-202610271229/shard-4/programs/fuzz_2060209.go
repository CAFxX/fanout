package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 bool; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint64; F1 int64; F2 string }

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 183
    var v3 int32 = -1
    v4 := GenS1{F0: true, F1: (v2 >> ((65535) & 63))}
    v2 -= (v2 >> ((595) & 63))
    for fi5 := 0; fi5 < 15; fi5++ {
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 uint16 = 255
            _ = fi6
            _ = v7
        }
        var v8 uint = 4294967295
        _ = fi5
        _ = v8
    }
    var v9 int32 = 2147483647
    if true {
        v10 := v1["\n\t"]
        delete(v1, "a")
        var v11 int64 = 0
        _ = v10
        _ = v11
    } else {
        switch ("foo bar" + "") {
        case "a":
            v3 = (v9 * (v9 >> ((973) & 63)))
        case "foo bar":
            for fi12 := 0; fi12 < 19; fi12++ {
                v13 := GenS2{F0: 302, F1: 255}
                if (-1 == 32767) {
                    v14 := []int{(^(fi12))}
                    _ = v14
                } else {
                    v15 := GenS3{F0: (-60 > 42), F1: 1, F2: 1}
                    v16 := []int{fi12, fi12, fi12}
                    _ = v15
                    _ = v16
                }
                v17 := []int{fi12, (fi12 | (fi12 >> ((0) & 63))), fi12, fi12}
                _ = fi12
                _ = v13
                _ = v17
            }
        }
        v18 := []int{}
        v18 = append(v18, 0)
        _ = v18
    }
    fmt.Println(v9)
    fmt.Println(v9, v2)
    v2 += v2
    v2 -= 128
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 int = fi19
        fmt.Println(v3, v3)
        _ = fi19
        _ = v20
    }
    if true {
        for fi21 := 0; fi21 < 17; fi21++ {
            for fi22 := 0; fi22 < 2; fi22++ {
                switch "\n\t" {
                case "hello":
                    v23 := GenS1{F0: true, F1: 207}
                    _ = v23
                case "":
                    v9 += v3
                default:
                    v2 = 0
                }
                if (false && (true && true)) {
                    var v24 string = "hello"
                    v25 := GenS3{F0: (!false), F1: 42, F2: 0}
                    var v26 int64 = 0
                    _ = v24
                    _ = v25
                    _ = v26
                }
                v3 -= (-(v9))
                _ = fi22
            }
            v27 := GenS4{F0: 18446744073709551615, F1: 0, F2: ""}
            _ = fi21
            _ = v27
        }
        var v28 float64 = float64(0.1)
        v2 = (^(v2))
        _ = v28
    }
    var v29 int8 = 1
    fmt.Println(v3, v29)
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v9
}
