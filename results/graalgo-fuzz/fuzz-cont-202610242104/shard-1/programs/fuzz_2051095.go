package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int32) int16 {
    if true {
        var v6 string = "\n\t"
        p0 = (p0 - p0)
        _ = v6
    } else {
        p0 = p0
    }
    v7 := []int{9223372036854775807, -332, 32767}
    var v8 int = 32768
    _ = p0
    _ = v7
    _ = v8
    return 0
}
func genFn1() string {
    var v9 bool = (!(127 < -1))
    v10 := []int{-9223372036854775808, 127, -703}
    v11 := GenS1{F0: ("hello" + ("foo bar" + "a")), F1: (!v9)}
    _ = v10
    _ = v11
    _ = v9
    return (("a" + "foo bar") + "foo bar")
}

func main() {
    v1 := []string{"foo bar", "", (("\n\t" + "foo bar") + ("" + ("foo bar" + (("hello" + (("a" + "a") + "")) + "a"))))}
    var v2 int = 0
    var v3 uint16 = 210
    var v4 int16 = 42
    var v5 int64 = -973
    var v12 uint64 = 4294967295
    fmt.Println(v5)
    if ((!false) && false) {
        for fi13 := 0; fi13 < 16; fi13++ {
            v14 := GenS1{F0: "", F1: (v3 <= (v3 << ((845) & 63)))}
            fmt.Println(v4)
            _ = fi13
            _ = v14
        }
        var v15 int = v2
        v16 := GenS2{F0: 1, F1: "hello"}
        _ = v15
        _ = v16
    }
    v2 = (v2 ^ -2147483648)
    var v17 int = 319
    var v18 string = "a"
    v19 := GenS1{F0: v18, F1: (!((false && false) || false))}
    if (363 < 405) {
        var v20 float32 = 0.1
        _ = v20
    } else {
        var v21 int16 = (v4 << ((2) & 63))
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 uint64 = v12
            switch "" {
            case "\n\t":
                var v24 uint = uint(v3)
                _ = v24
            default:
                v25 := GenS3{F0: 2}
                _ = v25
            }
            v26 := GenS3{F0: 841}
            _ = fi22
            _ = v23
            _ = v26
        }
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
