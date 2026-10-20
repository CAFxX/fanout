package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }

func main() {
    var v1 string = ""
    v2 := []uint32{65535, 443, 42}
    v3 := []uint32{2147483647, 65535}
    var v4 uint = 4294967295
    var v5 float64 = 1.5
    fmt.Println(v4, v1)
    if len(v2) > 0 {
        v2[0] = 667
    }
    v1 = "hello"
    var v6 int16 = 0
    v6 = (v6 + v6)
    if true {
        _ = len(v2)
        v7 := GenS1{F0: v4}
        if false {
            v5 -= 0.0
            var v8 bool = (v4 >= (^(v4)))
            _ = v8
        } else {
            var v9 int32 = 269
            if false {
                if false {
                    fmt.Println(v1, v6)
                    var v10 int = 1
                    v10 = 2147483647
                    _ = v10
                } else {
                    switch 1 {
                    case 32767:
                        if (false && false) {
                            _ = len(v2)
                            var v11 string = v1
                            _ = v11
                        }
                    default:
                        var v12 bool = false
                        _ = v12
                    }
                    var v13 int8 = -1
                    _ = v13
                }
                for fi14 := 0; fi14 < 15; fi14++ {
                    switch "a" {
                    case "\n\t":
                        for fi15 := 0; fi15 < 20; fi15++ {
                            v16 := GenS2{F0: false}
                            v4 = v4
                            fmt.Println(fi15, v1, v4)
                            _ = fi15
                            _ = v16
                        }
                    case "a":
                        v1 = v1
                    default:
                        v17 := GenS1{F0: 42}
                        _ = v17
                    }
                    var v18 uint16 = 2
                    v5 -= (3.14159 + v5)
                    _ = fi14
                    _ = v18
                }
            } else {
                var v19 float64 = v5
                _ = v19
            }
            for fi20 := 0; fi20 < 17; fi20++ {
                for fi21 := 0; fi21 < 2; fi21++ {
                    if ((fi20 % ((fi20) | 1)) <= fi20) {
                        _ = len(v2)
                    } else {
                        v9 = -682
                        var v22 string = ((v1 + (((((v1 + (v1 + v1)) + (((v1 + v1) + v1) + v1)) + (v1 + "a")) + ("a" + "\n\t")) + (v1 + v1))) + (v1 + v1))
                        _ = v22
                    }
                    v23 := GenS1{F0: v4}
                    _ = fi21
                    _ = v23
                }
                v1 = "\n\t"
                _ = fi20
            }
            _ = v9
        }
        _ = v7
    }
    var v24 int8 = 127
    fmt.Println(v5, v24, v6)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
