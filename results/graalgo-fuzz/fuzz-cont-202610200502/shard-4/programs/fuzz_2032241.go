package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int8 }
func genFn0(p0 int16, p1 string) int {
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 bool = true
        _ = fi7
        _ = v8
    }
    var v9 uint64 = 4294967295
    switch 925 {
    case 2147483647:
        p1 = p1
    default:
        for fi10 := 0; fi10 < 9; fi10++ {
            fmt.Println(fi10, v9, p0)
            _ = fi10
        }
    }
    v11 := []int{42, 4294967295}
    _ = p0
    _ = p1
    _ = v11
    _ = v9
    return 32768
}

func main() {
    var v1 bool = (32767 <= 42)
    var v2 bool = false
    var v3 uint32 = 0
    var v4 float32 = -2.25
    v5 := make(map[int]uint32)
    var v6 string = (("foo bar" + "hello") + "\n\t")
    v12 := GenS1{F0: 904, F1: (v6 + v6), F2: 1}
    for fi13 := 0; fi13 < 7; fi13++ {
        v5[(fi13 | (-(fi13)))] = (v3 >> ((4294967295) & 63))
        var v14 int16 = -32768
        if (v1 && v1) {
            v3 += (v3 & (v3 % (((v3 << ((467) & 63))) | 1)))
            var v15 float64 = 0.1
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(v4, v3, v3)
    for fi16 := 0; fi16 < 15; fi16++ {
        for fi17 := 0; fi17 < 7; fi17++ {
            if v2 {
                fmt.Println(v1, v4, v3)
                fmt.Println(v1, v4, v1)
            } else {
                v18 := GenS1{F0: (fi17 >> ((0) & 63)), F1: v6, F2: 1}
                _ = v18
            }
            fmt.Println(v2)
            v3 -= v3
            v1 = v2
            _ = fi17
        }
        v4 = ((v4 / (v4 / v4)) / (3.14159 / (v4 - v4)))
        _ = fi16
    }
    fmt.Println(v3, v4)
    var v19 uint64 = 675
    switch ("hello" + "a") {
    case "\n\t":
        v19 = (v19 | (v19 | 4294967295))
    case "hello":
        var v20 uint8 = 79
        _ = v20
    default:
        v1 = false
    }
    var v21 float64 = -2.25
    v19 += (^(v19))
    var v22 int32 = -2147483648
    var v23 uint64 = 129
    v1 = (v1 && ((v3 + 465) != (v3 % ((1) | 1))))
    if v1 {
        if v1 {
            for fi24 := 0; fi24 < 5; fi24++ {
                v25 := []int{fi24, (fi24 + genFn0(-45, v6)), (fi24 ^ (fi24 >> ((931) & 63))), (fi24 / ((int(v3)) | 1))}
                if v1 {
                    v22 += 2147483647
                    v2 = v2
                } else {
                    if v1 {
                        v5[fi24] = (^(v3))
                        var v26 uint = 255
                        _ = v26
                    }
                    var v27 int32 = -677
                    _ = v27
                }
                _ = fi24
                _ = v25
            }
            switch 866 {
            case 0:
                v21 += float64(v4)
            default:
                for fi28 := 0; fi28 < 17; fi28++ {
                    var v29 int32 = (v22 ^ 0)
                    _ = fi28
                    _ = v29
                }
            }
            for fi30 := 0; fi30 < 20; fi30++ {
                var v31 float32 = (v4 / float32(0.1))
                _ = fi30
                _ = v31
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
