package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 bool) int64 {
    var v5 int16 = -607
    fmt.Println(v5, p0)
    var v6 uint32 = 2
    _ = p0
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 string = (("" + "foo bar") + "a")
    v2 := make(map[string]string)
    var v3 uint64 = 4294967295
    var v4 int = -312
    if (((^(v4)) != v4) || (!true)) {
        var v7 uint64 = v3
        if ((v4 != (v4 & v4)) || (53 > -335)) {
            v7 = (v7 * (v7 - v7))
            v7 += v3
            v8 := []int{v4, 9223372036854775807, v4, 637}
            _ = v8
        }
        var v9 int32 = 42
        _ = v7
        _ = v9
    }
    v10 := []int{79, v4}
    v1 = ("hello" + ((v1 + v1) + "foo bar"))
    v11 := GenS1{F0: ((((v1 + "\n\t") + (("hello" + v1) + (((("hello" + ((v1 + v1) + (v1 + "\n\t"))) + (v1 + (v1 + v1))) + v1) + "a"))) + v1) + v1)}
    fmt.Println(v3)
    v4 = (v4 << ((65535) & 63))
    var v12 uint64 = v3
    var v13 int16 = -1
    v2["\n\t"] = "foo bar"
    var v14 int = (v4 - (v4 / ((32768) | 1)))
    var v15 int32 = 42
    var v16 float64 = 1e10
    for fi17 := 0; fi17 < 10; fi17++ {
        for fi18 := 0; fi18 < 16; fi18++ {
            if true {
                var v19 uint64 = (v3 + v12)
                v16 = (0.1 / v16)
                v13 = v13
                _ = v19
            }
            fmt.Println(v3, v12, v14)
            if (true || true) {
                switch v1 {
                case "a":
                    v1 = v1
                case "hello":
                    if true {
                        var v20 uint = 526
                        var v21 int8 = -1
                        _ = v20
                        _ = v21
                    } else {
                        if true {
                            var v22 string = v1
                            var v23 uint = 255
                            v12 -= (v12 & v12)
                            _ = v22
                            _ = v23
                        } else {
                            _ = len(v2)
                            var v24 float64 = float64((-2.25 * (1.5 - (((((3.14159 - 0.1) + 3.14159) / 3.14159) * (1.5 - float32(0.1))) * float32((v16 / v16))))))
                            v4 -= (fi18 - fi17)
                            var v25 uint64 = v12
                            _ = v24
                            _ = v25
                        }
                        v3 = (v12 >> ((2) & 63))
                    }
                }
                v2["a"] = ((v1 + "") + "foo bar")
                var v26 int8 = -128
                _ = v26
            }
            _ = fi18
        }
        for fi27 := 0; fi27 < 8; fi27++ {
            var v28 int16 = v13
            _ = fi27
            _ = v28
        }
        for fi29 := 0; fi29 < 15; fi29++ {
            delete(v2, ("\n\t" + "hello"))
            _ = fi29
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
