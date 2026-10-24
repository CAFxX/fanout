package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int, p1 uint) uint32 {
    var v5 int64 = 1
    fmt.Println(p1, v5)
    _ = p0
    _ = p1
    _ = v5
    return 0
}
func genFn1(p0 int32, p1 int8) uint16 {
    v6 := []int{128, 9223372036854775807, -1}
    var v7 uint8 = 232
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := []uint32{4294967295}
    var v2 uint = 255
    var v3 string = "\n\t"
    v4 := []uint32{496, 2}
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 int = fi8
        v1 = append(v1, 42)
        _ = fi8
        _ = v9
    }
    if (!(false && false)) {
        if len(v4) > 0 {
            v4[0] = 2
        }
        if true {
            for fi10 := 0; fi10 < 11; fi10++ {
                v3 = ""
                v3 = "\n\t"
                _ = fi10
            }
            var v11 float32 = -2.25
            _ = v11
        }
    }
    v2 = (v2 % (((^(v2))) | 1))
    switch v3 {
    case "foo bar":
        v12 := GenS1{F0: -9223372036854775808}
        _ = v12
    case "":
        var v13 int64 = -1
        _ = v13
    case "hello":
        var v14 uint16 = 65535
        _ = v14
    default:
        var v15 uint8 = 255
        _ = v15
    }
    var v16 float64 = 0.1
    switch -2147483648 {
    case 127:
        v17 := GenS1{F0: 388}
        _ = v17
    case 1:
        v1 = append(v1, 717)
    case 475:
        var v18 uint = 65535
        _ = v18
    default:
        for fi19 := 0; fi19 < 1; fi19++ {
            fmt.Println(v3)
            v2 -= 4294967295
            switch fi19 {
            case -2147483648:
                if (0 >= 222) {
                    v20 := GenS1{F0: 9223372036854775807}
                    if false {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    if (genFn1(-1, -128) == 813) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v21 int32 = 2147483647
                        v4 = append(v4, genFn0(-995, (v2 << (((v2 - 2)) & 63))))
                        _ = v21
                    } else {
                        if true {
                            v3 = "foo bar"
                        } else {
                            for fi22 := 0; fi22 < 17; fi22++ {
                                v1 = append(v1, genFn0((fi22 * 128), 0))
                                _ = fi22
                            }
                        }
                    }
                    _ = v20
                } else {
                    if len(v1) > 0 {
                        v1[0] = 65535
                    }
                }
            }
            _ = fi19
        }
    }
    v23 := GenS1{F0: 42}
    var v24 float32 = 1.5
    v1 = append(v1, 513)
    v2 = (v2 - (v2 >> ((v2) & 63)))
    v16 = v16
    v16 -= 0.0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
