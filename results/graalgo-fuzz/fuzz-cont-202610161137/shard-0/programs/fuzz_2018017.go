package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 bool = false
    v2 := make(map[string]string)
    v3 := make(map[int]bool)
    var v4 uint = 18446744073709551615
    var v5 string = ("hello" + "")
    fmt.Println(v5, v1)
    for fi6 := 0; fi6 < 9; fi6++ {
        v4 += 18446744073709551615
        _ = fi6
    }
    var v7 float64 = float64(float32((float64((-2.25 / 3.14159)) + ((0.0 / (1.5 / ((float64(1.5) + (0.1 * 3.14159)) * (3.14159 / (1.5 / (1e10 * 1.5)))))) / 1e10))))
    switch 760 {
    case -828:
        var v8 uint16 = 2
        _ = v8
    case 127:
        v9 := []int{-940, 128, 2147483647}
        _ = v9
    case 4294967295:
        v10 := GenS1{F0: 157}
        _ = v10
    }
    v7 = v7
    v11 := GenS1{F0: 52}
    v12 := GenS1{F0: 2}
    v7 -= 1.5
    for fi13 := 0; fi13 < 10; fi13++ {
        if (42 == 1) {
            delete(v3, fi13)
            var v14 int64 = -9223372036854775808
            v15 := []int{}
            _ = v14
            _ = v15
        } else {
            if true {
                for fi16 := 0; fi16 < 16; fi16++ {
                    var v17 int = fi13
                    _ = fi16
                    _ = v17
                }
            } else {
                fmt.Println(fi13, v5, v7)
            }
            switch (v5 + "a") {
            case "a":
                v5 = "hello"
            case "":
                v18 := []int{fi13, 128}
                _ = v18
            default:
                v1 = ((18446744073709551615 > 456) || (!(32767 < -32768)))
            }
            if true {
                v5 = "hello"
                if (!v1) {
                    var v19 int32 = -2147483648
                    _ = v19
                }
            } else {
                v20 := []int{fi13, (^(fi13)), 2147483647, (fi13 + (fi13 >> (((v4 / ((v4) | 1))) & 63)))}
                switch fi13 {
                case -373:
                    v21 := v2[(v5 + v5)]
                    _ = v21
                }
                var v22 string = v5
                _ = v20
                _ = v22
            }
        }
        _ = fi13
    }
    switch 127 {
    case 241:
        v7 += float64((0.1 - (-2.25 * 0.1)))
    case 0:
        switch v5 {
        case "hello":
            v5 = ""
        case "\n\t":
            if (v4 > v4) {
                if ((v4 & v4) != v4) {
                    if (v1 && (((-2.25 - (3.14159 / 1e10)) - 3.14159) >= 3.14159)) {
                        var v23 int8 = -12
                        fmt.Println(v7, v7)
                        for fi24 := 0; fi24 < 2; fi24++ {
                            var v25 string = v5
                            _ = fi24
                            _ = v25
                        }
                        _ = v23
                    } else {
                        var v26 uint8 = 239
                        _ = v26
                    }
                }
            }
        default:
            var v27 float32 = 0.1
            _ = v27
        }
    case -1:
        for fi28 := 0; fi28 < 14; fi28++ {
            fmt.Println(v4, v1, v4)
            _ = fi28
        }
    default:
        delete(v2, "foo bar")
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
