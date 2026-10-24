package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint64) uint16 {
    var v4 int = -235
    if (2147483647 == 1) {
        var v5 float32 = float32(3.14159)
        _ = v5
    }
    var v6 bool = (!false)
    _ = p0
    _ = v4
    _ = v6
    return 2
}

func main() {
    var v1 string = ("a" + ("" + ""))
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    v1 = "hello"
    fmt.Println(v1)
    v1 = v1
    fmt.Println(v1)
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint = 2
        _ = fi7
        _ = v8
    }
    if false {
        var v9 float32 = (1.5 - 0.1)
        for fi10 := 0; fi10 < 10; fi10++ {
            v11 := GenS1{F0: 993}
            v12 := []int{fi10, fi10, fi10, (fi10 * fi10)}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = v9
    } else {
        for fi13 := 0; fi13 < 1; fi13++ {
            v1 = v1
            if (!(3.14159 <= 1.5)) {
                v14 := GenS1{F0: -9223372036854775808}
                var v15 string = v1
                _ = v14
                _ = v15
            }
            v16 := GenS2{F0: 356, F1: fi13}
            _ = fi13
            _ = v16
        }
        v17 := v3[780]
        switch 32767 {
        case -293:
            fmt.Println(v1, v1)
        case 127:
            if true {
                v1 = ""
                v17 -= 2
            } else {
                var v18 float64 = 1e10
                if (32767 == 0) {
                    switch 1 {
                    case 1:
                        fmt.Println(v17, v1, v18)
                    default:
                        if ("hello" > v1) {
                            var v19 int16 = 1000
                            var v20 int = 42
                            _ = v19
                            _ = v20
                        } else {
                            fmt.Println(v1, v18, v17)
                        }
                    }
                    for fi21 := 0; fi21 < 10; fi21++ {
                        v18 = (0.1 - v18)
                        if (v1 == "") {
                            var v22 uint32 = (v17 << ((1) & 63))
                            var v23 float64 = v18
                            _ = v22
                            _ = v23
                        }
                        _ = fi21
                    }
                    var v24 float64 = v18
                    _ = v24
                } else {
                    v18 = v18
                    if true {
                        for fi25 := 0; fi25 < 2; fi25++ {
                            for fi26 := 0; fi26 < 13; fi26++ {
                                v27 := []int{fi25}
                                v28 := GenS3{F0: 32767}
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                            _ = fi25
                        }
                        v2[-9223372036854775808] = 9223372036854775807
                    } else {
                        var v29 int64 = 9223372036854775807
                        v1 = v1
                        _ = v29
                    }
                    v30 := []int{}
                    _ = v30
                }
                var v31 int32 = 1
                _ = v18
                _ = v31
            }
        case 1:
            for fi32 := 0; fi32 < 14; fi32++ {
                for fi33 := 0; fi33 < 6; fi33++ {
                    v34 := v2[fi32]
                    _ = fi33
                    _ = v34
                }
                _ = fi32
            }
        default:
            fmt.Println(v1, v1)
        }
        _ = v17
    }
    v35 := GenS1{F0: 9223372036854775807}
    v1 = "hello"
    var v36 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
}
