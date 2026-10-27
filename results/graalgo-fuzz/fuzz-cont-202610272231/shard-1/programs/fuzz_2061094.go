package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int16) uint8 {
    var v6 string = ("" + "foo bar")
    fmt.Println(p0, v6)
    v7 := GenS1{F0: v6}
    p0 += (p0 & 32767)
    fmt.Println(p0, p0)
    _ = p0
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[string]int)
    var v3 bool = true
    v4 := make(map[uint32]int)
    var v5 int32 = 992
    var v8 int = 17
    if (true && v3) {
        v9 := GenS1{F0: "hello"}
        var v10 uint16 = v1
        for fi11 := 0; fi11 < 17; fi11++ {
            fmt.Println(v3)
            _ = fi11
        }
        _ = v10
        _ = v9
    }
    v3 = (v3 && (((1e10 + 1e10) * (0.1 / 0.1)) > ((0.0 / float64(1e10)) / 1.0)))
    v12 := []int{(v8 * -9223372036854775808), 569}
    fmt.Println(v1, v8)
    v3 = (("\n\t" + ("\n\t" + (("a" + ("hello" + (("a" + "hello") + "foo bar"))) + "hello"))) == "hello")
    for fi13 := 0; fi13 < 15; fi13++ {
        v3 = false
        _ = fi13
    }
    if (!(371 != 0)) {
        fmt.Println(v3, v8, v5)
    } else {
        for fi14 := 0; fi14 < 17; fi14++ {
            v15 := GenS1{F0: "hello"}
            if v3 {
                if (!(false || (v3 || (!false)))) {
                    fmt.Println(v3, v1)
                    v8 -= 0
                }
                v8 = fi14
            }
            _ = fi14
            _ = v15
        }
        if (0 >= 174) {
            v8 -= v8
            v16 := v2["a"]
            _ = v16
        }
        fmt.Println(v1, v1)
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        fmt.Println(v5)
        var v18 uint = 12
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 uint = (v18 & (v18 + v18))
            _ = len(v4)
            _ = fi19
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    var v21 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
