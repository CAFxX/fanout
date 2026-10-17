package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 string }
func genFn0() uint32 {
    v4 := []int{-336, -2147483648, -585}
    if false {
        var v5 bool = (false && false)
        var v6 int = 128
        _ = v5
        _ = v6
    }
    if false {
        v7 := GenS1{F0: 255, F1: 4294967295}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 float32 = float32(-2.25)
            v9 = v9
            _ = fi8
            _ = v9
        }
    }
    var v10 int8 = -128
    var v11 uint = 2
    _ = v10
    _ = v11
    _ = v4
    return uint32(v11)
}
func genFn1(p0 uint8, p1 uint16, p2 string) int8 {
    switch 2147483647 {
    case 102:
        if true {
            if false {
                p0 += p0
                for fi12 := 0; fi12 < 19; fi12++ {
                    var v13 string = p2
                    v14 := []int{fi12, -909, fi12}
                    v13 = v13
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            } else {
                for fi15 := 0; fi15 < 20; fi15++ {
                    for fi16 := 0; fi16 < 4; fi16++ {
                        var v17 int = fi15
                        _ = fi16
                        _ = v17
                    }
                    _ = fi15
                }
                var v18 uint8 = (^(p0))
                _ = v18
            }
        }
    case 0:
        var v19 int32 = 337
        _ = v19
    }
    fmt.Println(p2, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    var v1 int64 = -1
    v2 := make(map[int]string)
    var v3 uint16 = 65535
    fmt.Println(v3)
    if false {
        v3 = (v3 >> ((65535) & 63))
    } else {
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 int64 = v1
            _ = fi20
            _ = v21
        }
    }
    fmt.Println(v3, v3)
    v22 := GenS2{F0: ("foo bar" + "\n\t")}
    for fi23 := 0; fi23 < 10; fi23++ {
        fmt.Println(v3, v3)
        fmt.Println(v1, fi23, v3)
        for fi24 := 0; fi24 < 14; fi24++ {
            if false {
                var v25 bool = ((-2.25 + 0.0) == 3.14159)
                v26 := v2[(-(fi24))]
                fmt.Println(v3)
                _ = v25
                _ = v26
            } else {
                v3 -= 255
                var v27 bool = false
                _ = v27
            }
            var v28 uint8 = 42
            var v29 uint32 = 42
            fmt.Println(v28)
            fmt.Println(v1, fi23, v3)
            v3 = (v3 - 953)
            _ = fi24
            _ = v28
            _ = v29
        }
        _ = fi23
    }
    var v30 float32 = -2.25
    _ = len(v2)
    var v31 uint16 = v3
    switch "hello" {
    case "":
        v2[-506] = "a"
    case "hello":
        v32 := []int{-1}
        _ = v32
    case "a":
        var v33 float32 = v30
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v31
}
