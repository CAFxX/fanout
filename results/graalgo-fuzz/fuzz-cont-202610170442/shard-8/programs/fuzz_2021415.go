package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := []string{"hello"}
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    if ((-2.25 < 0.0) || true) {
        var v4 float32 = 0.1
        for fi5 := 0; fi5 < 14; fi5++ {
            if true {
                v6 := GenS1{F0: (0 != 42)}
                var v7 string = ("a" + ("a" + (("" + "hello") + (("hello" + "foo bar") + ("\n\t" + "hello")))))
                _ = v6
                _ = v7
            }
            v4 = ((1.5 * v4) / (float32(float64(v4)) * (1.5 / float32(-2.25))))
            _ = fi5
        }
        if (1.5 <= 1.5) {
            var v8 int32 = 0
            if true {
                var v9 int8 = 58
                fmt.Println(v4)
                var v10 bool = (true || ((18446744073709551615 > 922) || false))
                _ = v10
                _ = v9
            } else {
                v11 := GenS1{F0: true}
                _ = v11
            }
            v4 = v4
            _ = v8
        } else {
            v12 := v3[4294967295]
            var v13 int8 = 0
            _ = v12
            _ = v13
        }
        _ = v4
    } else {
        delete(v2, 2147483647)
    }
    var v14 float32 = 0.0
    var v15 int16 = -703
    v1 = append(v1, ("a" + ("\n\t" + "")))
    var v16 string = "foo bar"
    v16 = v16
    var v17 int8 = -128
    v17 = -128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
