package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint; F1 bool }
func genFn0(p0 int16, p1 string, p2 int64) int {
    if ((3.14159 / float32(1e10)) <= 1e10) {
        var v5 float64 = 0.0
        fmt.Println(p1)
        _ = v5
    } else {
        v6 := GenS1{F0: (((1e10 * 0.1) - 1e10) >= 3.14159), F1: 0, F2: p2}
        p1 = p1
        if true {
            if (!(42 != 1)) {
                if true {
                    v7 := GenS2{F0: 2, F1: 42}
                    _ = v7
                }
                var v8 string = p1
                _ = v8
            } else {
                p1 = p1
                fmt.Println(p0, p0)
                var v9 uint16 = 530
                _ = v9
            }
            v10 := GenS1{F0: true, F1: 42, F2: p2}
            _ = v10
        }
        _ = v6
    }
    var v11 uint64 = 18446744073709551615
    fmt.Println(v11, v11, v11)
    v12 := []int{-432}
    p1 = (p1 + p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 1
}
func genFn1(p0 uint32, p1 int32, p2 int16) bool {
    p1 = p1
    for fi13 := 0; fi13 < 15; fi13++ {
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 uint8 = 127
            _ = fi14
            _ = v15
        }
        if false {
            fmt.Println(p1)
        } else {
            var v16 uint64 = 42
            var v17 uint16 = 761
            _ = v16
            _ = v17
        }
        v18 := []int{fi13}
        _ = fi13
        _ = v18
    }
    p2 += (p2 * -1)
    for fi19 := 0; fi19 < 13; fi19++ {
        for fi20 := 0; fi20 < 13; fi20++ {
            switch fi19 {
            case 189:
                p1 += p1
            case 175:
                switch fi19 {
                case 4294967295:
                    var v21 bool = (((!(2147483648 >= p0)) || (1 != fi19)) || (false || (((fi20 % ((617) | 1)) == 2147483647) || (true || (true || true)))))
                    _ = v21
                default:
                    fmt.Println(fi19, fi19)
                }
            default:
                if true {
                    p0 -= (p0 >> ((1) & 63))
                    fmt.Println(p0, p0)
                }
            }
            var v22 int64 = 100
            _ = fi20
            _ = v22
        }
        _ = fi19
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        for fi24 := 0; fi24 < 19; fi24++ {
            var v25 uint64 = 911
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    _ = p0
    _ = p1
    _ = p2
    return false
}

func main() {
    var v1 int16 = 42
    v2 := make(map[string]int)
    var v3 uint32 = 65535
    var v4 int32 = 0
    var v26 int32 = (v4 >> ((2) & 63))
    for fi27 := 0; fi27 < 12; fi27++ {
        for fi28 := 0; fi28 < 13; fi28++ {
            var v29 int64 = 0
            _ = fi28
            _ = v29
        }
        var v30 float64 = (3.14159 - 0.0)
        _ = fi27
        _ = v30
    }
    var v31 uint = 175
    var v32 int32 = (v26 * v26)
    var v33 uint8 = 2
    v32 += (^(v4))
    fmt.Println(v32, v26, v33)
    v33 = v33
    switch -1 {
    case -1:
        switch -719 {
        case 32768:
            v34 := []int{2147483647, 128, -134, genFn0((v1 % ((-1) | 1)), "\n\t", 308)}
            _ = v34
        }
    case 32767:
        var v35 bool = ((v32 - int32(v1)) != 0)
        _ = v35
    case 280:
        fmt.Println(v26, v4)
    default:
        v36 := GenS3{F0: (v31 ^ v31), F1: (!(!(9223372036854775807 >= -1)))}
        _ = v36
    }
    var v37 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v37
    _ = v4
}
