package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int }
func genFn0(p0 uint16, p1 int32) uint32 {
    v7 := []int{55, -555, -1}
    var v8 int16 = 0
    _ = len(v7)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 float32 = (-2.25 - 0.1)
    var v2 int8 = 127
    v3 := make(map[uint32]string)
    v4 := make(map[string]int)
    v5 := []string{"a"}
    var v6 uint = 65535
    if true {
        fmt.Println(v1)
    } else {
        _ = len(v4)
        if (!(2147483647 >= -1)) {
            var v9 int32 = -636
            _ = v9
        } else {
            v1 -= 1e10
        }
        v10 := GenS1{F0: "foo bar"}
        _ = v10
    }
    var v11 int = 0
    var v12 uint32 = 407
    if ((!true) && true) {
        for fi13 := 0; fi13 < 19; fi13++ {
            v11 = (v11 % (((v11 | (v11 ^ (-(fi13))))) | 1))
            _ = fi13
        }
        _ = len(v5)
        var v14 int8 = int8(v6)
        _ = v14
    }
    v1 = (v1 * (1e10 * v1))
    v2 = v2
    for fi15 := 0; fi15 < 7; fi15++ {
        if (!(368 <= (^(v12)))) {
            var v16 int = (v11 - fi15)
            v12 -= (^(v12))
            _ = v16
        } else {
            var v17 float32 = v1
            if len(v5) > 0 {
                v5[0] = "\n\t"
            }
            _ = v17
        }
        if true {
            switch ("a" + ("\n\t" + (("foo bar" + ("hello" + "")) + ""))) {
            case "":
                v6 = (v6 * (v6 << (((v6 ^ (^(v6)))) & 63)))
            case "\n\t":
                for fi18 := 0; fi18 < 12; fi18++ {
                    var v19 float64 = (0.0 + (0.1 + 0.0))
                    _ = fi18
                    _ = v19
                }
            }
            v6 = v6
        } else {
            v5 = append(v5, ("foo bar" + "hello"))
        }
        _ = fi15
    }
    if (0 == 1) {
        if (!true) {
            var v20 int = (v11 << ((v6) & 63))
            _ = v20
        }
        v6 = v6
        var v21 uint32 = 87
        _ = v21
    } else {
        var v22 int64 = -9223372036854775808
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v22
    }
    _ = len(v4)
    for fi23 := 0; fi23 < 14; fi23++ {
        if false {
            fmt.Println(v12, v11)
            v12 = v12
        } else {
            switch (((("a" + "a") + ("" + "\n\t")) + (("\n\t" + "") + "hello")) + (((("foo bar" + "foo bar") + "a") + ("foo bar" + "")) + ((((("foo bar" + "\n\t") + ("foo bar" + "\n\t")) + (("\n\t" + (("hello" + "") + "foo bar")) + "hello")) + "hello") + (("" + ("" + "a")) + "hello")))) {
            case "a":
                fmt.Println(v12)
            default:
                fmt.Println(v2)
            }
        }
        _ = len(v4)
        if true {
            if false {
                v24 := GenS1{F0: "a"}
                fmt.Println(v11)
                _ = v24
            } else {
                v12 = (^(v12))
            }
            fmt.Println(fi23, v1)
            v25 := GenS1{F0: "foo bar"}
            _ = v25
        } else {
            fmt.Println(v6, v12, v1)
            v26 := GenS2{F0: 2, F1: fi23}
            _ = v26
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
