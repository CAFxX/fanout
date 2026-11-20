package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 bool; F1 int64 }
func genFn0() bool {
    var v5 uint = 65535
    var v6 int = 42
    fmt.Println(v5)
    if (107 < 127) {
        v7 := GenS1{F0: (((("foo bar" + "a") + "") + "") + ("hello" + "foo bar")), F1: 0}
        var v8 bool = (4294967295 > 129)
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: v5, F1: 42, F2: 0}
    _ = v5
    _ = v6
    _ = v9
    return (!(false || false))
}

func main() {
    var v1 uint16 = 1
    v2 := []int{int(v1), 2147483647, -644}
    var v3 int = 1
    v4 := []int{}
    v10 := GenS1{F0: "hello", F1: 0}
    for fi11 := 0; fi11 < 10; fi11++ {
        v1 = (^(v1))
        var v12 uint64 = 1
        if true {
            for fi13 := 0; fi13 < 11; fi13++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi13
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi11
        _ = v12
    }
    v1 -= (v1 ^ (v1 * v1))
    for fi14 := 0; fi14 < 1; fi14++ {
        v3 += (fi14 * (v3 / (((fi14 >> ((4294967295) & 63))) | 1)))
        v3 = v3
        _ = fi14
    }
    if true {
        var v15 float32 = float32(3.14159)
        _ = v15
    } else {
        fmt.Println(v3, v1)
        fmt.Println(v3, v1)
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        fmt.Println(fi16)
        _ = fi16
    }
    var v17 bool = (-333 == 32767)
    v3 = (v3 + v3)
    var v18 uint = 65535
    fmt.Println(v18)
    v3 = v3
    for fi19 := 0; fi19 < 11; fi19++ {
        if v17 {
            for fi20 := 0; fi20 < 11; fi20++ {
                var v21 uint64 = 1
                if len(v2) > 0 {
                    v2[0] = 256
                }
                _ = fi20
                _ = v21
            }
            for fi22 := 0; fi22 < 16; fi22++ {
                _ = len(v4)
                var v23 string = ""
                v18 = (v18 ^ v18)
                _ = fi22
                _ = v23
            }
            v17 = (!((128 < v3) && ((64 > 100) || v17)))
        } else {
            v18 += v18
            var v24 uint = 600
            v25 := GenS3{F0: 0, F1: 42}
            _ = v24
            _ = v25
        }
        switch "\n\t" {
        case "hello":
            v3 = (-(fi19))
        default:
            v26 := GenS2{F0: (v18 / (((^(v18))) | 1)), F1: 0, F2: 2}
            _ = v26
        }
        _ = fi19
    }
    v27 := GenS4{F0: (9223372036854775807 > 355), F1: int64(v18)}
    v1 += 310
    var v28 float32 = 0.0
    for fi29 := 0; fi29 < 6; fi29++ {
        v30 := GenS1{F0: "", F1: 2147483647}
        var v31 int32 = 692
        _ = fi29
        _ = v30
        _ = v31
    }
    v32 := GenS4{F0: (0 <= v3), F1: 130}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v4
}
