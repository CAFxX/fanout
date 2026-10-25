package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int64 }
func genFn0(p0 int64, p1 uint) string {
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            if (225 == 128) {
                var v8 int64 = (p0 >> (((p1 & (p1 - p1))) & 63))
                v9 := GenS1{F0: 797}
                _ = v8
                _ = v9
            }
            switch ("" + "foo bar") {
            case "a":
                if (1.5 < float64((3.14159 + 3.14159))) {
                    fmt.Println(fi7, fi6)
                }
            }
            var v10 uint64 = 0
            p0 = 0
            fmt.Println(p1, p0)
            _ = fi7
            _ = v10
        }
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 float32 = 0.1
            v13 := []int{fi11, (fi6 + int(p0)), (fi11 / (((fi11 / (((fi6 & fi11)) | 1))) | 1))}
            _ = fi11
            _ = v12
            _ = v13
        }
        v14 := GenS2{F0: -189, F1: 1}
        _ = fi6
        _ = v14
    }
    var v15 int16 = 32767
    _ = p0
    _ = p1
    _ = v15
    return "a"
}
func genFn1() uint32 {
    v16 := GenS3{F0: (708 != 531), F1: ""}
    var v17 uint64 = 42
    _ = v16
    _ = v17
    return 2
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = false
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    v5 := make(map[uint32]bool)
    v18 := []int{127, -231, 0, -9223372036854775808}
    if (((0.0 / (3.14159 * 3.14159)) != -2.25) || (464 <= 32767)) {
        var v19 int32 = 26
        for fi20 := 0; fi20 < 18; fi20++ {
            if len(v18) > 0 {
                v18[0] = fi20
            }
            _ = fi20
        }
        _ = v19
    }
    v2 = true
    fmt.Println(v2)
    switch "\n\t" {
    case "hello":
        v21 := GenS2{F0: 1, F1: 765}
        _ = v21
    default:
        v22 := GenS2{F0: -1, F1: 1}
        _ = v22
    }
    v5[859] = (4294967295 < 0)
    delete(v1, -957)
    v23 := GenS4{F0: 9223372036854775807}
    v24 := GenS1{F0: genFn1()}
    fmt.Println(v2)
    v25 := v5[2]
    var v26 uint8 = 0
    v27 := v4[4294967295]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
