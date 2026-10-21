package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
func genFn0(p0 uint32, p1 uint16, p2 int64) string {
    v7 := GenS1{F0: 0, F1: p1}
    var v8 uint64 = 555
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return "\n\t"
}
func genFn1() int {
    var v9 uint64 = 42
    v9 = v9
    var v10 int16 = 32767
    v11 := []int{4294967295}
    _ = v10
    _ = v11
    _ = v9
    return -602
}

func main() {
    v1 := make(map[int]string)
    v2 := []int{-2147483648}
    v3 := make(map[uint32]int)
    var v4 int = 0
    var v5 int8 = -128
    v6 := make(map[uint32]int)
    for fi12 := 0; fi12 < 13; fi12++ {
        v6[213] = fi12
        for fi13 := 0; fi13 < 7; fi13++ {
            v4 += fi12
            _ = fi13
        }
        fmt.Println(v5, fi12)
        _ = fi12
    }
    var v14 int8 = v5
    for fi15 := 0; fi15 < 18; fi15++ {
        _ = len(v3)
        if len(v2) > 0 {
            v2[0] = (v4 * (fi15 << ((490) & 63)))
        }
        fmt.Println(v5)
        _ = fi15
    }
    v14 = v5
    fmt.Println(v5, v4, v5)
    var v16 string = ("" + "\n\t")
    switch 128 {
    case 2147483647:
        v16 = (("\n\t" + v16) + "")
    }
    var v17 int32 = -2147483648
    _ = len(v2)
    if true {
        v5 += (v5 + -79)
        var v18 string = v16
        for fi19 := 0; fi19 < 2; fi19++ {
            fmt.Println(v17)
            _ = fi19
        }
        _ = v18
    } else {
        for fi20 := 0; fi20 < 10; fi20++ {
            if (!true) {
                if (true && (-32768 != -1)) {
                    switch v16 {
                    case "foo bar":
                        v1[(fi20 & (fi20 & genFn1()))] = "foo bar"
                    case "":
                        v16 = v16
                    }
                    for fi21 := 0; fi21 < 16; fi21++ {
                        v16 = v16
                        _ = fi21
                    }
                    v3[9] = (fi20 >> ((65535) & 63))
                }
                var v22 int16 = 1
                delete(v3, 685)
                _ = v22
            }
            _ = fi20
        }
    }
    v23 := GenS1{F0: 255, F1: 189}
    v14 -= v14
    if (v17 > -1) {
        for fi24 := 0; fi24 < 5; fi24++ {
            if false {
                v4 += (fi24 / ((v4) | 1))
                v5 += v14
            } else {
                switch 1 {
                case 222:
                    v17 = (v17 % (((-(v17))) | 1))
                default:
                    delete(v1, (-(fi24)))
                }
                fmt.Println(v17)
            }
            _ = fi24
        }
        var v25 int16 = 1
        v14 += (v14 * (-(v14)))
        _ = v25
    } else {
        var v26 bool = ((-107 >= v14) || (!true))
        var v27 int16 = 0
        _ = v26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
