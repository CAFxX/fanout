package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint; F1 uint; F2 int64 }

func main() {
    v1 := make(map[int]bool)
    var v2 string = "a"
    var v3 bool = (42 > 207)
    var v4 uint64 = 4294967295
    var v5 float32 = 1e10
    var v6 int64 = 466
    var v7 uint32 = 2147483648
    v8 := GenS1{F0: (v4 | 2), F1: v6}
    var v9 float64 = (float64(v5) / 1.0)
    var v10 uint32 = v7
    for fi11 := 0; fi11 < 1; fi11++ {
        switch fi11 {
        case 42:
            if ((105 <= 100) || (!false)) {
                for fi12 := 0; fi12 < 1; fi12++ {
                    delete(v1, (fi11 ^ fi12))
                    var v13 string = v2
                    v14 := []int{(fi12 + (^(fi12))), fi11, (-(fi12))}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                v4 -= 241
            } else {
                if false {
                    fmt.Println(v9)
                    v7 += v10
                    var v15 uint64 = (v4 >> ((0) & 63))
                    _ = v15
                } else {
                    var v16 int = 192
                    if v3 {
                        v1[(fi11 >> ((1) & 63))] = (v9 > float64(float32((v9 / v9))))
                        fmt.Println(v3, v16, v2)
                        v3 = (true && (-32768 != -32768))
                    }
                    _ = v16
                }
            }
        }
        var v17 bool = (!((v6 ^ v6) > (v6 | v6)))
        switch ("a" + "a") {
        case "hello":
            var v18 uint64 = v4
            _ = v18
        case "":
            for fi19 := 0; fi19 < 7; fi19++ {
                if v17 {
                    var v20 int = -1
                    _ = v20
                }
                var v21 float32 = float32(float64((v5 / (v5 / 1.0))))
                _ = fi19
                _ = v21
            }
        case "\n\t":
            v22 := GenS1{F0: 2, F1: v6}
            _ = v22
        default:
            var v23 uint = 0
            _ = v23
        }
        _ = fi11
        _ = v17
    }
    v1[2147483647] = (v2 == ("hello" + v2))
    v24 := GenS1{F0: v4, F1: v6}
    v25 := []int{127, 16}
    var v26 uint16 = 32767
    v27 := GenS2{F0: 195, F1: 270, F2: (v6 ^ (v6 % ((v6) | 1)))}
    for fi28 := 0; fi28 < 8; fi28++ {
        var v29 int16 = 328
        var v30 string = "hello"
        v7 += v7
        _ = fi28
        _ = v29
        _ = v30
    }
    v26 += v26
    if len(v25) > 0 {
        v25[0] = 0
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
