package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 uint; F2 int64 }
type GenS3 struct { F0 int; F1 int16 }
func genFn0() int8 {
    switch 2147483647 {
    case 9223372036854775807:
        var v7 uint8 = 117
        _ = v7
    case -291:
        var v8 string = "foo bar"
        _ = v8
    case -9223372036854775808:
        v9 := []int{}
        _ = v9
    }
    var v10 uint16 = 425
    fmt.Println(v10)
    v10 -= (v10 & 255)
    _ = v10
    return 31
}
func genFn1(p0 int64) int {
    switch -2147483648 {
    case 128:
        v11 := GenS1{F0: false}
        _ = v11
    case -1:
        fmt.Println(p0)
    default:
        var v12 uint64 = 4294967295
        _ = v12
    }
    fmt.Println(p0)
    _ = p0
    return 128
}

func main() {
    v1 := make(map[string]string)
    var v2 float32 = float32(((1e10 + 1e10) - 0.1))
    var v3 float64 = 0.1
    var v4 uint64 = 18446744073709551615
    var v5 uint8 = 128
    v6 := []uint32{65535, 4294967295, 2147483648}
    v6 = append(v6, 292)
    var v13 uint = 42
    v5 -= 2
    for fi14 := 0; fi14 < 16; fi14++ {
        v2 = (v2 / v2)
        _ = fi14
    }
    fmt.Println(v13)
    if false {
        v15 := GenS1{F0: (!(!true))}
        var v16 int16 = 1
        v17 := GenS2{F0: 42, F1: v13, F2: 42}
        _ = v15
        _ = v16
        _ = v17
    }
    switch int(v4) {
    case -9223372036854775808:
        var v18 bool = (42 < 1)
        _ = v18
    }
    switch ("hello" + "\n\t") {
    case "a":
        var v19 uint8 = v5
        _ = v19
    default:
        if (v2 >= v2) {
            for fi20 := 0; fi20 < 15; fi20++ {
                for fi21 := 0; fi21 < 16; fi21++ {
                    var v22 float32 = 0.0
                    var v23 float64 = 0.1
                    var v24 int8 = -1
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                }
                v25 := GenS3{F0: genFn1(0), F1: 1}
                v1[("hello" + "foo bar")] = ("a" + "hello")
                _ = fi20
                _ = v25
            }
            v1["foo bar"] = "a"
            if true {
                v26 := v1[("a" + "foo bar")]
                if (0 < v13) {
                    for fi27 := 0; fi27 < 2; fi27++ {
                        v5 = (v5 ^ 252)
                        var v28 int32 = 2147483647
                        v5 += (v5 << ((v13) & 63))
                        _ = fi27
                        _ = v28
                    }
                    v5 += 2
                }
                v1["\n\t"] = "foo bar"
                _ = v26
            }
        } else {
            switch ("hello" + ("a" + ((("a" + ("hello" + "a")) + "") + "a"))) {
            case "\n\t":
                v2 = 1e10
            case "":
                v5 += 2
            case "a":
                var v29 uint16 = 1
                _ = v29
            }
            v4 -= v4
        }
    }
    v30 := GenS1{F0: true}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
