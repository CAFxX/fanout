package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 bool; F2 uint64 }
func genFn0(p0 int16, p1 string) string {
    var v5 uint8 = 245
    var v6 float64 = 0.0
    var v7 int32 = 1
    var v8 bool = true
    p1 = p1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return ((p1 + p1) + ("a" + ""))
}
func genFn1(p0 uint) int8 {
    v9 := []int{229}
    var v10 uint8 = 42
    v9 = append(v9, 9223372036854775807)
    _ = len(v9)
    _ = p0
    _ = v10
    _ = v9
    return -1
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    var v4 uint8 = 255
    v4 = 0
    var v11 bool = true
    switch -2147483648 {
    case 32768:
        var v12 int8 = genFn1(2)
        _ = v12
    }
    var v13 int64 = -497
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := []int{(fi14 << ((255) & 63))}
        if len(v15) > 0 {
            v15[0] = fi14
        }
        _ = fi14
        _ = v15
    }
    fmt.Println(v13)
    for fi16 := 0; fi16 < 9; fi16++ {
        v13 = (v13 * (v13 >> ((255) & 63)))
        fmt.Println(v11)
        _ = fi16
    }
    var v17 string = ("hello" + ("hello" + "foo bar"))
    var v18 float32 = 3.14159
    _ = len(v3)
    var v19 uint8 = 42
    switch -408 {
    case -131:
        for fi20 := 0; fi20 < 13; fi20++ {
            v21 := GenS1{F0: -89}
            var v22 uint64 = 469
            _ = fi20
            _ = v21
            _ = v22
        }
    case 9223372036854775807:
        for fi23 := 0; fi23 < 6; fi23++ {
            var v24 bool = (!false)
            for fi25 := 0; fi25 < 1; fi25++ {
                var v26 int32 = -2147483648
                v27 := GenS2{F0: (v26 * -1), F1: ("hello" < v17), F2: 2}
                _ = fi25
                _ = v26
                _ = v27
            }
            for fi28 := 0; fi28 < 7; fi28++ {
                v29 := v2[-1]
                _ = fi28
                _ = v29
            }
            _ = fi23
            _ = v24
        }
    default:
        fmt.Println(v17)
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        if false {
            v11 = (0 > 18446744073709551615)
            var v31 uint8 = v19
            _ = v31
        }
        _ = fi30
    }
    v19 = v19
    var v32 int64 = (v13 << ((869) & 63))
    v33 := v2[-2147483648]
    for fi34 := 0; fi34 < 12; fi34++ {
        v35 := []int{}
        for fi36 := 0; fi36 < 5; fi36++ {
            switch int(v19) {
            case 0:
                var v37 int8 = -1
                _ = v37
            case -209:
                var v38 float32 = v18
                _ = v38
            default:
                if ((v4 << ((4294967295) & 63)) < v4) {
                    if v11 {
                        var v39 string = v17
                        if v11 {
                            v35 = append(v35, fi34)
                        }
                        _ = v39
                    }
                } else {
                    _ = len(v35)
                    v13 += v13
                    if len(v35) >= 2 {
                        v35 = v35[1:len(v35)]
                    }
                }
            }
            var v40 uint = 65535
            var v41 uint = (v40 * v40)
            _ = fi36
            _ = v40
            _ = v41
        }
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
