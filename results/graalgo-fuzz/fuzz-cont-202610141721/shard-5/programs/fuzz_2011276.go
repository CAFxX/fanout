package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint, p1 uint8) uint64 {
    p1 = 128
    var v7 float32 = (1e10 + 1.5)
    v7 = ((v7 / v7) * float32(1e10))
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 0
    v3 := make(map[uint32]int)
    var v4 uint32 = 42
    v5 := make(map[int]string)
    v6 := make(map[int]int)
    v8 := v5[-2147483648]
    for fi9 := 0; fi9 < 12; fi9++ {
        if true {
            var v10 uint = 1
            _ = v10
        }
        _ = fi9
    }
    if (false || (2147483647 <= -133)) {
        _ = len(v3)
        var v11 int16 = -953
        _ = v11
    } else {
        _ = len(v6)
    }
    if (0 > 856) {
        var v12 int = 32768
        var v13 string = v8
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: v2}
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 int32 = 0
        v17 := GenS1{F0: (^(v2))}
        var v18 uint16 = 32767
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    switch 127 {
    case 127:
        delete(v3, v4)
    case -1:
        for fi19 := 0; fi19 < 10; fi19++ {
            for fi20 := 0; fi20 < 9; fi20++ {
                var v21 int64 = 9223372036854775807
                var v22 int16 = 300
                _ = fi20
                _ = v21
                _ = v22
            }
            if false {
                v23 := GenS1{F0: v2}
                if false {
                    var v24 uint = 741
                    switch 42 {
                    case -212:
                        v8 = v8
                    case -405:
                        var v25 uint64 = v2
                        _ = v25
                    case -440:
                        var v26 int32 = 1
                        _ = v26
                    default:
                        var v27 float32 = -2.25
                        _ = v27
                    }
                    _ = v24
                } else {
                    for fi28 := 0; fi28 < 2; fi28++ {
                        v29 := []int{(fi19 >> ((65535) & 63))}
                        var v30 uint8 = 57
                        v2 += (v2 & (v2 ^ v2))
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    var v31 uint32 = (v4 % ((65535) | 1))
                    _ = v31
                }
                var v32 uint = 991
                _ = v23
                _ = v32
            } else {
                v4 = 803
                if (42 > 2) {
                    var v33 int64 = 1
                    for fi34 := 0; fi34 < 10; fi34++ {
                        if true {
                            switch 9223372036854775807 {
                            case 2147483647:
                                v4 = 1
                            }
                        }
                        var v35 float64 = -2.25
                        _ = fi34
                        _ = v35
                    }
                    var v36 float64 = 1e10
                    _ = v33
                    _ = v36
                }
                fmt.Println(v2, v4, v8)
            }
            var v37 int8 = 88
            _ = fi19
            _ = v37
        }
    case 4294967295:
        var v38 string = "hello"
        _ = v38
    }
    fmt.Println(v4, v4, v8)
    var v39 int16 = 0
    v2 = v2
    v40 := GenS2{F0: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v8
}
