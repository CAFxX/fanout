package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int }

func main() {
    var v1 int16 = 32767
    var v2 float32 = 3.14159
    var v3 int32 = 0
    if (false && (-952 >= 32768)) {
        if (false && (32768 > 713)) {
            v1 += -1
            switch 32768 {
            case 0:
                if (!false) {
                    fmt.Println(v2, v3)
                } else {
                    for fi4 := 0; fi4 < 6; fi4++ {
                        var v5 bool = false
                        _ = fi4
                        _ = v5
                    }
                    v1 = 0
                    if true {
                        v6 := GenS1{F0: 428, F1: 65535}
                        var v7 string = (("a" + "") + "foo bar")
                        var v8 uint16 = 42
                        _ = v6
                        _ = v7
                        _ = v8
                    }
                }
            case 2147483647:
                v2 = (v2 + v2)
            case 4294967295:
                switch 175 {
                case 4294967295:
                    var v9 uint32 = 2
                    _ = v9
                case 128:
                    v2 = 3.14159
                default:
                    if (!((98 > 42) || false)) {
                        var v10 bool = ((!true) || false)
                        for fi11 := 0; fi11 < 3; fi11++ {
                            if v10 {
                                v12 := GenS2{F0: 32767}
                                _ = v12
                            } else {
                                switch -580 {
                                case 0:
                                    var v13 int64 = 363
                                    _ = v13
                                default:
                                    var v14 uint32 = 243
                                    _ = v14
                                }
                            }
                            _ = fi11
                        }
                        v2 = (0.0 / v2)
                        _ = v10
                    } else {
                        var v15 uint16 = 876
                        _ = v15
                    }
                }
            }
        } else {
            var v16 int8 = 100
            var v17 string = "hello"
            v2 -= (v2 * 0.1)
            _ = v16
            _ = v17
        }
    }
    fmt.Println(v2)
    v18 := []int{-1}
    fmt.Println(v2)
    v18 = append(v18, 128)
    fmt.Println(v3, v1)
    v2 -= 0.1
    if len(v18) > 0 {
        v18[0] = 206
    }
    var v19 uint = 644
    fmt.Println(v19, v19, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
