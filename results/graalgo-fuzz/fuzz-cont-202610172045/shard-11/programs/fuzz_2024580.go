package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int; F1 int64 }
func genFn0(p0 int64, p1 int8) uint {
    var v4 uint16 = 0
    fmt.Println(v4)
    if false {
        v5 := GenS1{F0: ((("foo bar" + "") + "hello") + "foo bar")}
        _ = v5
    } else {
        p0 += p0
        var v6 int16 = 1000
        v7 := []int{2147483647, -218, -767, 32767}
        _ = v6
        _ = v7
    }
    p0 = p0
    var v8 uint = 65535
    _ = p0
    _ = p1
    _ = v4
    _ = v8
    return uint(v4)
}
func genFn1(p0 int) int16 {
    v9 := []int{}
    var v10 int16 = 0
    _ = p0
    _ = v10
    _ = v9
    return 1000
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[string]string)
    var v3 uint32 = 65535
    for fi11 := 0; fi11 < 1; fi11++ {
        v3 -= v3
        if ((v1 % (((v1 * (^(v1)))) | 1)) == v1) {
            if false {
                v1 = (v1 << ((0) & 63))
            } else {
                var v12 int16 = -1
                v2["hello"] = "\n\t"
                _ = v12
            }
            if false {
                var v13 uint64 = (^(v1))
                _ = v13
            } else {
                for fi14 := 0; fi14 < 11; fi14++ {
                    v3 -= v3
                    _ = fi14
                }
                v3 = uint32(v1)
            }
        } else {
            for fi15 := 0; fi15 < 5; fi15++ {
                delete(v2, ("hello" + ((("foo bar" + "foo bar") + "a") + "foo bar")))
                v16 := GenS1{F0: ((("a" + "\n\t") + "a") + "\n\t")}
                var v17 int = fi15
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        _ = fi11
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        v19 := GenS1{F0: ("" + "a")}
        v3 = v3
        v20 := GenS2{F0: 4294967295, F1: -906, F2: 127}
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 float32 = 1e10
    if true {
        var v22 float32 = 1.5
        var v23 string = "\n\t"
        v24 := GenS3{F0: 14, F1: 32767}
        _ = v22
        _ = v23
        _ = v24
    } else {
        v25 := GenS4{F0: v1, F1: 1}
        _ = v25
    }
    var v26 uint = genFn0(1, -128)
    if (false || false) {
        var v27 int64 = 0
        v1 += v1
        _ = v27
    }
    v28 := GenS5{F0: 849, F1: 1}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v28
    _ = v3
}
