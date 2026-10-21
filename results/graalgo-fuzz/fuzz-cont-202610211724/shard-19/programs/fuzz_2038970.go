package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 bool }
func genFn0(p0 int64, p1 int8, p2 uint64) string {
    p2 += (p2 << ((18446744073709551615) & 63))
    var v7 uint8 = 64
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "foo bar"
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 2
    var v3 int8 = 100
    var v4 uint = 0
    var v5 float64 = ((0.0 / 1e10) - -2.25)
    v6 := make(map[int]string)
    var v8 uint64 = 1
    fmt.Println(v8, v3)
    fmt.Println(v5)
    if true {
        var v9 int = 127
        v10 := GenS1{F0: 1, F1: 926, F2: false}
        switch "" {
        case "hello":
            v11 := []int{(v9 >> (((^(v4))) & 63)), (v9 >> ((0) & 63)), (v9 & 524)}
            _ = v11
        case "":
            delete(v1, (v9 - (v9 / (((v9 & v9)) | 1))))
        case "foo bar":
            v5 += v5
        }
        _ = v10
        _ = v9
    }
    v8 = v8
    v8 -= 18446744073709551615
    var v12 uint16 = 2
    for fi13 := 0; fi13 < 4; fi13++ {
        fmt.Println(fi13, v4, v2)
        _ = fi13
    }
    v4 += uint(v12)
    if false {
        if (!true) {
            v8 += (v8 ^ v8)
            var v14 string = (("" + "foo bar") + ("\n\t" + ("a" + "\n\t")))
            for fi15 := 0; fi15 < 14; fi15++ {
                var v16 float64 = v5
                _ = fi15
                _ = v16
            }
            _ = v14
        } else {
            var v17 uint8 = 127
            var v18 int8 = v3
            v12 = 2
            _ = v17
            _ = v18
        }
    } else {
        var v19 uint32 = v2
        switch 494 {
        case 9223372036854775807:
            var v20 bool = (!false)
            _ = v20
        case 305:
            var v21 uint8 = 128
            _ = v21
        case 529:
            if (true || true) {
                var v22 int8 = v3
                var v23 uint64 = (^(v8))
                _ = v22
                _ = v23
            } else {
                switch "" {
                case "hello":
                    v12 += 255
                default:
                    v24 := []int{}
                    _ = v24
                }
                var v25 uint64 = v8
                for fi26 := 0; fi26 < 6; fi26++ {
                    var v27 int16 = int16(fi26)
                    _ = fi26
                    _ = v27
                }
                _ = v25
            }
        default:
            for fi28 := 0; fi28 < 13; fi28++ {
                var v29 bool = (-2147483648 <= 42)
                for fi30 := 0; fi30 < 20; fi30++ {
                    v31 := v1[401]
                    v32 := GenS1{F0: -658, F1: v8, F2: (v31 && v31)}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                switch "" {
                case "":
                    if (v5 >= (0.0 / 1e10)) {
                        var v33 uint = 2
                        _ = v33
                    }
                case "hello":
                    var v34 bool = (0 >= v3)
                    _ = v34
                }
                _ = fi28
                _ = v29
            }
        }
        _ = v19
    }
    v35 := GenS1{F0: 0, F1: (^(v8)), F2: (true && (v12 == (v12 + 509)))}
    v8 = (v8 * uint64(v12))
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
