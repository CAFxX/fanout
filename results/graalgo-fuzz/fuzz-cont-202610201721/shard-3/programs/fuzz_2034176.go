package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
func genFn0(p0 string, p1 uint16, p2 uint) int16 {
    var v5 int16 = -466
    fmt.Println(v5, p0, p2)
    switch p0 {
    case "a":
        v6 := []int{-1, -342, 32768}
        _ = v6
    case "foo bar":
        switch "\n\t" {
        case "a":
            v7 := GenS1{F0: 255}
            _ = v7
        case "":
            var v8 uint8 = 128
            _ = v8
        case "foo bar":
            fmt.Println(v5)
        default:
            for fi9 := 0; fi9 < 11; fi9++ {
                if (true && true) {
                    p1 = (p1 << ((p2) & 63))
                }
                for fi10 := 0; fi10 < 16; fi10++ {
                    p1 += 111
                    if false {
                        v11 := []int{}
                        v11 = append(v11, (^(fi10)))
                        _ = v11
                    }
                    _ = fi10
                }
                for fi12 := 0; fi12 < 10; fi12++ {
                    p0 = p0
                    _ = fi12
                }
                _ = fi9
            }
        }
    default:
        v13 := []int{}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 32767
}

func main() {
    v1 := []uint32{1, 886, 42}
    v2 := make(map[int]string)
    v3 := []uint32{2147483647, 65535}
    var v4 float64 = (3.14159 / ((((1.5 + float64(3.14159)) * -2.25) - 3.14159) - 0.0))
    var v14 string = ("\n\t" + "")
    fmt.Println(v4, v4)
    v1 = append(v1, 42)
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 int64 = -1
        v16 -= int64(fi15)
        _ = fi15
        _ = v16
    }
    fmt.Println(v14)
    var v17 int16 = 1000
    v18 := GenS1{F0: 2}
    if false {
        v4 -= (float64((-2.25 - ((0.0 - 1.5) * ((float32(float64(0.0)) / 0.1) * float32(v4))))) * v4)
        fmt.Println(v14, v14)
        fmt.Println(v17, v14)
    } else {
        v14 = "a"
    }
    var v19 int64 = -691
    switch 32767 {
    case 127:
        var v20 int64 = v19
        _ = v20
    default:
        _ = len(v3)
    }
    if (1 > 100) {
        v4 = (float64(1e10) / float64(0.1))
    }
    if (22 <= 42) {
        var v21 uint8 = 176
        _ = v21
    } else {
        if true {
            v14 = ("a" + v14)
            var v22 int32 = -2147483648
            var v23 int16 = (^(v17))
            _ = v22
            _ = v23
        } else {
            v2[-484] = v14
        }
        v14 = "foo bar"
        delete(v2, 546)
    }
    if len(v3) > 0 {
        v3[0] = 1
    }
    v19 -= v19
    v4 = (((v4 - v4) / 3.14159) * v4)
    v24 := GenS1{F0: 518}
    v25 := v2[32767]
    v26 := GenS2{F0: 4294967295}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
