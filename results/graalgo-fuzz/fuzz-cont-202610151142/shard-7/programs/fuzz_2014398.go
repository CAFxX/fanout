package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 bool }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint8, p1 uint16) int16 {
    v5 := GenS1{F0: 1, F1: 82, F2: true}
    var v6 float64 = (float64(1e10) + -2.25)
    fmt.Println(p0)
    v7 := []int{-1, 32768, 0}
    var v8 float64 = ((((float64(-2.25) - v6) / ((v6 / v6) / v6)) / -2.25) / 1.0)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint8, p1 bool, p2 int32) uint16 {
    p0 = 88
    for fi9 := 0; fi9 < 17; fi9++ {
        if (!(917 < -9223372036854775808)) {
            fmt.Println(p2, p0)
            var v10 int32 = p2
            for fi11 := 0; fi11 < 9; fi11++ {
                v12 := GenS2{F0: 2147483647}
                switch "\n\t" {
                case "a":
                    v13 := []int{(fi11 << ((1) & 63)), (-(fi11))}
                    _ = v13
                }
                _ = fi11
                _ = v12
            }
            _ = v10
        }
        _ = fi9
    }
    p0 = p0
    var v14 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 32767
}

func main() {
    var v1 uint64 = 2
    var v2 int32 = 42
    var v3 float64 = 0.1
    v4 := make(map[string]int)
    fmt.Println(v1, v1)
    delete(v4, "hello")
    var v15 int = 32768
    if (!true) {
        v2 -= (v2 - (v2 % ((-1) | 1)))
        var v16 int16 = 0
        _ = v16
    } else {
        if (!true) {
            v1 += 42
        }
    }
    if (!false) {
        for fi17 := 0; fi17 < 13; fi17++ {
            fmt.Println(fi17, v15, v1)
            v15 += (^(fi17))
            if false {
                fmt.Println(v15, v3, v15)
                var v18 int16 = -239
                _ = v18
            } else {
                delete(v4, ("foo bar" + "a"))
                for fi19 := 0; fi19 < 5; fi19++ {
                    var v20 uint = 0
                    var v21 int = fi19
                    v3 = (v3 * v3)
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
            _ = fi17
        }
        v1 = (v1 << ((4294967295) & 63))
        switch "a" {
        case "":
            for fi22 := 0; fi22 < 10; fi22++ {
                var v23 float32 = 1.5
                v3 += v3
                _ = fi22
                _ = v23
            }
        case "foo bar":
            var v24 int64 = 1
            _ = v24
        }
    }
    fmt.Println(v1, v2)
    var v25 bool = (false && false)
    if v25 {
        v15 += v15
        v3 = 1e10
        if (!(945 != 0)) {
            var v26 uint64 = 52
            v15 += 959
            if true {
                fmt.Println(v3, v15)
                v4["hello"] = (v15 - v15)
            } else {
                var v27 uint16 = 2
                _ = v27
            }
            _ = v26
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
