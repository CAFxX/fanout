package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint32) int16 {
    var v5 int16 = -1
    switch 127 {
    case -802:
        fmt.Println(v5)
    case -275:
        var v6 uint8 = 128
        _ = v6
    default:
        v7 := []int{785, 32768}
        _ = v7
    }
    fmt.Println(p0)
    _ = p0
    _ = v5
    return 40
}
func genFn1(p0 int16, p1 int32) string {
    v8 := []int{2147483647, 2147483647, 127}
    var v9 uint8 = 0
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return "a"
}

func main() {
    var v1 string = "hello"
    var v2 uint = 994
    var v3 bool = true
    v4 := make(map[int]string)
    if false {
        for fi10 := 0; fi10 < 1; fi10++ {
            if v3 {
                if (-128 != -1) {
                    for fi11 := 0; fi11 < 15; fi11++ {
                        var v12 int8 = 42
                        var v13 uint64 = 240
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    v14 := GenS1{F0: 4294967295}
                    _ = v14
                } else {
                    if (3.14159 != 0.0) {
                        v15 := GenS2{F0: int32(fi10)}
                        var v16 int16 = -1
                        _ = v15
                        _ = v16
                    } else {
                        _ = len(v4)
                        var v17 bool = (v3 && v3)
                        _ = v17
                    }
                    v18 := GenS2{F0: 495}
                    v2 += (v2 % (((^(v2))) | 1))
                    _ = v18
                }
                for fi19 := 0; fi19 < 7; fi19++ {
                    v20 := GenS2{F0: -412}
                    _ = fi19
                    _ = v20
                }
            } else {
                fmt.Println(v1, fi10)
            }
            v3 = (((825 == 47) && v3) || v3)
            if (109 == 1) {
                for fi21 := 0; fi21 < 5; fi21++ {
                    v22 := []int{(fi21 ^ 0), (fi21 / (((fi21 + (fi21 ^ int(v2)))) | 1)), (fi21 ^ fi10), fi21}
                    var v23 bool = ((fi21 / (((fi10 >> ((42) & 63))) | 1)) <= (fi21 | int(v2)))
                    v24 := v4[522]
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                }
            }
            _ = fi10
        }
    } else {
        var v25 int = 1
        var v26 uint16 = 890
        _ = v25
        _ = v26
    }
    var v27 uint64 = 4294967295
    var v28 int8 = 100
    var v29 uint8 = 254
    if v3 {
        delete(v4, 127)
    }
    if v3 {
        v2 -= v2
    }
    var v30 bool = true
    _ = len(v4)
    switch 894 {
    case 127:
        v30 = (-2.25 == 3.14159)
    case -2147483648:
        for fi31 := 0; fi31 < 10; fi31++ {
            if true {
                v27 = (v27 ^ v27)
                fmt.Println(fi31, v30)
            }
            _ = fi31
        }
    }
    var v32 int8 = v28
    v29 = v29
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
}
