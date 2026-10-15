package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
func genFn0() uint32 {
    v7 := GenS1{F0: 0}
    v8 := GenS2{F0: 42}
    var v9 string = "a"
    if false {
        v9 = v9
    } else {
        var v10 uint16 = 2
        _ = v10
    }
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int32) uint {
    var v11 int64 = -1
    fmt.Println(p0, p0)
    switch ("" + "foo bar") {
    case "a":
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 string = "a"
            _ = fi12
            _ = v13
        }
    }
    _ = p0
    _ = v11
    return 1
}

func main() {
    var v1 uint16 = 42
    var v2 string = (((("foo bar" + ((((("hello" + "a") + "a") + "a") + "a") + ("" + "a"))) + ("foo bar" + "")) + "foo bar") + ("foo bar" + "foo bar"))
    var v3 int = 1
    var v4 int16 = -612
    v5 := make(map[string]string)
    var v6 float64 = -2.25
    delete(v5, "hello")
    if (genFn0() > 1) {
        v2 = v2
        var v14 uint32 = genFn0()
        _ = v14
    } else {
        var v15 uint64 = 42
        var v16 uint32 = 65535
        _ = v15
        _ = v16
    }
    fmt.Println(v4, v4)
    if ((!false) && (false || false)) {
        fmt.Println(v3, v1)
        v1 -= v1
    }
    v17 := []int{127, (^(v3)), v3}
    var v18 float32 = ((-2.25 * 0.0) - -2.25)
    var v19 int64 = -9223372036854775808
    switch "foo bar" {
    case "":
        fmt.Println(v4, v3)
    }
    v18 -= v18
    v17 = append(v17, (v3 | v3))
    fmt.Println(v19)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
