package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int16; F1 string; F2 uint8 }
func genFn0(p0 int64, p1 uint32) int16 {
    switch "hello" {
    case "\n\t":
        var v5 float32 = 1e10
        _ = v5
    case "a":
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 uint16 = 1
            _ = fi6
            _ = v7
        }
    default:
        if (false && true) {
            v8 := GenS1{F0: 0}
            var v9 uint32 = (p1 - 2147483648)
            if (42 < 0) {
                var v10 int64 = p0
                _ = v10
            }
            _ = v8
            _ = v9
        } else {
            v11 := GenS1{F0: 862}
            for fi12 := 0; fi12 < 10; fi12++ {
                if true {
                    var v13 bool = false
                    var v14 uint8 = 227
                    _ = v13
                    _ = v14
                } else {
                    v15 := GenS2{F0: 255, F1: -60}
                    var v16 int32 = 42
                    var v17 bool = (true && (false || ((!(!(255 != 2))) && (2 >= 0))))
                    _ = v15
                    _ = v16
                    _ = v17
                }
                v18 := GenS2{F0: 0, F1: 0}
                _ = fi12
                _ = v18
            }
            for fi19 := 0; fi19 < 7; fi19++ {
                fmt.Println(fi19)
                v20 := GenS3{F0: 42, F1: ("a" + "hello"), F2: 127}
                _ = fi19
                _ = v20
            }
            _ = v11
        }
    }
    fmt.Println(p1, p1)
    var v21 string = ("a" + "")
    _ = p0
    _ = p1
    _ = v21
    return 939
}

func main() {
    var v1 int16 = 1000
    var v2 float64 = 0.0
    var v3 uint16 = 437
    v4 := make(map[string]string)
    var v22 uint32 = 4294967295
    v3 = (v3 >> ((uint(v22)) & 63))
    v1 = (^(v1))
    fmt.Println(v2)
    var v23 float32 = 0.1
    if true {
        var v24 int = 32767
        _ = v24
    } else {
        if (-173 != 42) {
            v1 = v1
        } else {
            var v25 uint32 = (v22 + v22)
            v3 = 194
            _ = v25
        }
    }
    v3 += 32767
    for fi26 := 0; fi26 < 18; fi26++ {
        var v27 int = fi26
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
