package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 int; F2 int32 }
func genFn0(p0 uint64) uint64 {
    fmt.Println(p0)
    p0 = (p0 - (^(p0)))
    var v7 int32 = -1
    _ = p0
    _ = v7
    return 144
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 42
    var v3 bool = (-617 >= -711)
    var v4 int = -2147483648
    var v5 int16 = -32768
    var v6 float32 = (3.14159 * (1.5 + 0.0))
    if false {
        v8 := GenS1{F0: -905, F1: 0, F2: v5}
        _ = v8
    } else {
        if (42 != 0) {
            v9 := GenS2{F0: 333}
            v10 := GenS3{F0: v5}
            _ = len(v1)
            _ = v10
            _ = v9
        } else {
            v11 := []int{2147483647, 42}
            var v12 string = "hello"
            _ = v11
            _ = v12
        }
        var v13 int8 = 1
        _ = v13
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        v1["hello"] = ""
        _ = fi14
    }
    if true {
        if false {
            v1["foo bar"] = ("foo bar" + "foo bar")
        }
        if (!((v3 || (-128 <= 100)) && (1 != 145))) {
            var v15 int8 = 0
            for fi16 := 0; fi16 < 19; fi16++ {
                v17 := GenS4{F0: 0, F1: v4, F2: 0}
                _ = fi16
                _ = v17
            }
            switch "hello" {
            case "hello":
                v18 := []int{20, (v4 << ((42) & 63)), 1, (-(v4))}
                _ = v18
            }
            _ = v15
        } else {
            var v19 int8 = -1
            _ = v19
        }
    } else {
        v2 = (v2 << ((661) & 63))
    }
    var v20 uint64 = 1
    for fi21 := 0; fi21 < 11; fi21++ {
        var v22 uint16 = 42
        if ((true && false) || v3) {
            var v23 int32 = -2147483648
            v3 = v3
            for fi24 := 0; fi24 < 9; fi24++ {
                if v3 {
                    v2 = 65535
                    var v25 int16 = 154
                    _ = v25
                }
                var v26 float32 = (float32(1e10) / v6)
                _ = fi24
                _ = v26
            }
            _ = v23
        }
        v3 = false
        _ = fi21
        _ = v22
    }
    switch ("\n\t" + ("" + (("" + ("a" + "a")) + "\n\t"))) {
    case "hello":
        v2 -= (v2 + v2)
    }
    v3 = v3
    fmt.Println(v6, v4, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
