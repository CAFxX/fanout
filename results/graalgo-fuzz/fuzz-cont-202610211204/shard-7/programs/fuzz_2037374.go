package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 string; F2 int64 }
func genFn0(p0 uint, p1 int32, p2 uint16) int {
    var v5 int32 = p1
    var v6 uint64 = 18446744073709551615
    p1 = (p1 / (((^(v5))) | 1))
    var v7 int8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 837
}
func genFn1() string {
    var v8 float64 = float64(((1.5 / 3.14159) + 3.14159))
    var v9 int = 9223372036854775807
    _ = v8
    _ = v9
    return "foo bar"
}

func main() {
    var v1 float32 = ((0.1 + 0.1) - 0.0)
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    var v4 float64 = 1e10
    var v10 bool = true
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            v10 = v10
            _ = fi12
        }
        if v10 {
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 int64 = 0
                _ = fi13
                _ = v14
            }
            for fi15 := 0; fi15 < 7; fi15++ {
                v1 += 0.1
                _ = fi15
            }
        } else {
            v16 := []int{(fi11 / (((^(fi11))) | 1)), (fi11 * -141)}
            _ = v16
        }
        if v10 {
            var v17 uint8 = 128
            var v18 uint = 320
            v17 = (v17 << (((v18 / (((v18 << ((v18) & 63))) | 1))) & 63))
            _ = v17
            _ = v18
        } else {
            v1 += 1e10
            v19 := GenS1{F0: true, F1: 42, F2: -497}
            var v20 uint = 18446744073709551615
            _ = v19
            _ = v20
        }
        _ = fi11
    }
    if v10 {
        if (!(!((v1 + (((v1 - (float32((float64(float32(v4)) / (v4 - v4))) * v1)) + (v1 + float32(float64((float32(float64((-2.25 + (1e10 - (v1 * ((float32(((((1e10 * float64(1e10)) * v4) * v4) / 1e10)) + v1) + 1.5)))))) + 3.14159))))) * v1)) != ((v1 + v1) * 0.1)))) {
            if (2147483648 == 2147483647) {
                v21 := []int{0}
                var v22 int64 = -548
                switch 32768 {
                case -790:
                    switch "" {
                    case "a":
                        v22 -= v22
                    case "":
                        delete(v2, 127)
                    }
                case -1:
                    var v23 uint64 = 1
                    _ = v23
                default:
                    if true {
                        v24 := GenS2{F0: "a"}
                        v4 += 1.5
                        _ = v24
                    }
                }
                _ = v21
                _ = v22
            }
            v1 = v1
            var v25 int32 = 493
            _ = v25
        }
    }
    var v26 uint16 = 672
    v27 := GenS3{F0: 1000, F1: "\n\t", F2: -447}
    if true {
        var v28 int32 = 0
        v26 = v26
        _ = v28
    }
    fmt.Println(v1, v10, v1)
    for fi29 := 0; fi29 < 13; fi29++ {
        for fi30 := 0; fi30 < 18; fi30++ {
            var v31 int8 = -1
            if (false && v10) {
                _ = len(v2)
                if (!((-1 < 42) || v10)) {
                    v31 += v31
                    var v32 uint8 = 2
                    v10 = v10
                    _ = v32
                }
                switch ("" + "hello") {
                case "a":
                    var v33 uint = 1
                    _ = v33
                }
                var v34 uint16 = (^(v26))
                _ = v34
            }
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    v26 += 42
    v35 := GenS2{F0: "foo bar"}
    var v36 int16 = 42
    var v37 float32 = float32(float64(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
