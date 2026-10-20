package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint }
func genFn0() string {
    switch -9223372036854775808 {
    case 32768:
        var v6 int16 = 1000
        _ = v6
    }
    var v7 uint8 = 2
    v7 = (v7 + 42)
    fmt.Println(v7)
    _ = v7
    return "\n\t"
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 float32 = float32((3.14159 / 1.0))
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]string)
    var v5 float32 = float32(((3.14159 * (((0.0 * 1.5) / 0.1) / 1.0)) * 1.5))
    switch "\n\t" {
    case "hello":
        var v8 int8 = 1
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        if (!(false || false)) {
            v1 = v1
        } else {
            for fi10 := 0; fi10 < 4; fi10++ {
                switch (fi10 & 9223372036854775807) {
                case 880:
                    if (!(2 != 8)) {
                        var v11 int16 = 0
                        _ = v11
                    } else {
                        fmt.Println(v5, fi10)
                        var v12 float32 = v5
                        v1 = 0
                        _ = v12
                    }
                }
                if false {
                    v13 := GenS1{F0: -53, F1: 35}
                    _ = v13
                }
                v5 = ((v5 + v5) * (0.0 * -2.25))
                v14 := []int{}
                v5 -= v2
                v5 -= 0.1
                _ = fi10
                _ = v14
            }
            switch (("a" + "a") + (("foo bar" + "a") + (("foo bar" + "foo bar") + "a"))) {
            case "hello":
                v5 += v2
            }
        }
        v15 := GenS1{F0: -112, F1: 0}
        _ = fi9
        _ = v15
    }
    if true {
        v16 := GenS2{F0: 778}
        fmt.Println(v1, v5, v1)
        _ = v16
    }
    if (-32768 != 1) {
        v3[2] = 2147483647
        var v17 int64 = 0
        _ = v17
    } else {
        v1 = v1
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 bool = (65535 >= 42)
        v20 := GenS3{F0: 945}
        var v21 int64 = 578
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    v5 = (v5 - v2)
    fmt.Println(v1, v5)
    v1 = v1
    if false {
        var v22 string = "foo bar"
        v3[4294967295] = 169
        v2 -= v5
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
