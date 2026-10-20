package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int16; F1 uint; F2 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 0
    var v3 uint8 = 1
    v4 := []uint32{2147483648, 0}
    var v5 float32 = (3.14159 + 1e10)
    v6 := GenS1{F0: 4294967295, F1: ("a" + "hello")}
    v4 = append(v4, 4294967295)
    fmt.Println(v3)
    v4 = append(v4, 0)
    var v7 int16 = 32767
    var v8 int = 128
    v9 := GenS1{F0: 727, F1: "a"}
    v2 = (v2 * v2)
    for fi10 := 0; fi10 < 1; fi10++ {
        v8 += (fi10 - (v8 % ((-9223372036854775808) | 1)))
        if true {
            v3 = 64
            var v11 uint32 = 2
            v2 = v2
            _ = v11
        } else {
            if false {
                for fi12 := 0; fi12 < 12; fi12++ {
                    fmt.Println(v2, fi10)
                    v13 := GenS1{F0: 2, F1: "foo bar"}
                    v7 += (v7 * v7)
                    _ = fi12
                    _ = v13
                }
                var v14 uint = 0
                fmt.Println(fi10, v5)
                _ = v14
            } else {
                var v15 int = fi10
                var v16 uint = 0
                _ = v15
                _ = v16
            }
            for fi17 := 0; fi17 < 10; fi17++ {
                v18 := GenS1{F0: 748, F1: ("\n\t" + ("hello" + ""))}
                var v19 uint16 = 42
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        for fi20 := 0; fi20 < 18; fi20++ {
            var v21 int32 = 0
            if (false && true) {
                v21 += v21
                var v22 uint16 = 65535
                v2 = (v21 >> ((1) & 63))
                _ = v22
            }
            fmt.Println(v2, v3)
            _ = fi20
            _ = v21
        }
        _ = fi10
    }
    v4 = append(v4, 2147483648)
    var v23 float64 = float64((((-2.25 * v5) - v5) + (v5 * (v5 - 3.14159))))
    var v24 int16 = (^(v7))
    for fi25 := 0; fi25 < 13; fi25++ {
        v23 += v23
        v26 := GenS2{F0: v7, F1: 344, F2: -905}
        var v27 bool = true
        _ = fi25
        _ = v26
        _ = v27
    }
    delete(v1, ("hello" + "foo bar"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
