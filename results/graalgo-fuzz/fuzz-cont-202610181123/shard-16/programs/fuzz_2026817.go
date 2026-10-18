package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint32; F2 int }
type GenS3 struct { F0 int32; F1 int32; F2 bool }
func genFn0(p0 string, p1 uint32) bool {
    for fi4 := 0; fi4 < 17; fi4++ {
        for fi5 := 0; fi5 < 20; fi5++ {
            var v6 uint32 = p1
            var v7 bool = true
            var v8 float32 = (0.1 * 0.1)
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
        var v9 bool = (p1 != (p1 >> ((798) & 63)))
        _ = fi4
        _ = v9
    }
    p0 = (p0 + p0)
    _ = p0
    _ = p1
    return true
}

func main() {
    var v1 bool = (-536 != 9223372036854775807)
    v2 := make(map[string]uint32)
    v3 := make(map[string]bool)
    var v10 int16 = 614
    v11 := GenS1{F0: ("\n\t" + "foo bar")}
    v1 = (!((!v1) || false))
    var v12 float64 = 3.14159
    if (!v1) {
        v13 := v2[((("a" + "") + ("\n\t" + "hello")) + "\n\t")]
        _ = v13
    }
    switch 32767 {
    case 32768:
        if (0 > 127) {
            var v14 uint = 0
            v10 += (v10 << ((v14) & 63))
            _ = v14
        } else {
            switch 449 {
            case 491:
                if true {
                    delete(v2, "hello")
                    v15 := GenS2{F0: (v10 / ((v10) | 1)), F1: 4294967295, F2: 9223372036854775807}
                    for fi16 := 0; fi16 < 19; fi16++ {
                        switch fi16 {
                        case 2147483647:
                            if v1 {
                                fmt.Println(fi16, v10)
                            } else {
                                var v17 uint8 = 127
                                _ = v17
                            }
                        default:
                            v1 = v1
                        }
                        var v18 int8 = 100
                        _ = fi16
                        _ = v18
                    }
                    _ = v15
                } else {
                    v12 = v12
                }
            default:
                v12 -= (v12 - 0.1)
            }
            v19 := GenS3{F0: 1, F1: 2147483647, F2: (127 >= 255)}
            _ = v19
        }
    case -9223372036854775808:
        var v20 bool = (!(1 > 68))
        _ = v20
    }
    if v1 {
        v12 += (v12 - ((v12 / (v12 - 0.0)) / 3.14159))
    } else {
        v21 := GenS3{F0: -1, F1: 2147483647, F2: v1}
        v12 -= 3.14159
        fmt.Println(v10)
        _ = v21
    }
    switch "\n\t" {
    case "":
        v22 := []int{681, 0, 228}
        _ = v22
    case "hello":
        fmt.Println(v10, v12, v12)
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
}
