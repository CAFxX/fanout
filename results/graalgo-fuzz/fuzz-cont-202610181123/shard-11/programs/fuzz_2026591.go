package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
func genFn0() int64 {
    var v5 float32 = 1e10
    v6 := []int{-356, 1, 9223372036854775807, 42}
    _ = v5
    _ = v6
    return -672
}

func main() {
    var v1 uint16 = 756
    v2 := make(map[int]uint32)
    var v3 int16 = 32767
    v4 := []string{"foo bar", "hello"}
    if ((2 < 4294967295) || true) {
        var v7 int16 = 42
        _ = len(v2)
        _ = v7
    }
    v3 = v3
    v3 += (v3 >> ((18446744073709551615) & 63))
    var v8 float32 = -2.25
    if true {
        var v9 uint16 = v1
        if (!true) {
            if (-417 != -289) {
                fmt.Println(v8, v1, v8)
            } else {
                switch ((((("" + ("hello" + ("hello" + "a"))) + "a") + ((("\n\t" + "hello") + "\n\t") + ("foo bar" + ("hello" + "foo bar")))) + "") + "foo bar") {
                case "foo bar":
                    v9 = v1
                }
                if len(v4) > 0 {
                    v4[0] = "foo bar"
                }
                var v10 int16 = (v3 + (v3 | 1))
                _ = v10
            }
            v1 = v9
            for fi11 := 0; fi11 < 11; fi11++ {
                v12 := v2[fi11]
                if true {
                    for fi13 := 0; fi13 < 13; fi13++ {
                        fmt.Println(v8, v1)
                        v14 := v2[fi13]
                        fmt.Println(v14, v9, v12)
                        _ = fi13
                        _ = v14
                    }
                    v8 = 0.0
                } else {
                    fmt.Println(v1, v3, v8)
                    v1 = (^(v9))
                }
                if len(v4) > 0 {
                    v4[0] = "foo bar"
                }
                _ = fi11
                _ = v12
            }
        }
        _ = v9
    } else {
        v8 += v8
    }
    if ("hello" < ("" + ("a" + "a"))) {
        var v15 int = 128
        _ = v15
    } else {
        v1 = (^(v1))
        var v16 uint = 65535
        v17 := GenS1{F0: 0}
        _ = v16
        _ = v17
    }
    var v18 int16 = v3
    for fi19 := 0; fi19 < 13; fi19++ {
        v20 := GenS2{F0: ("hello" + ("foo bar" + "foo bar"))}
        v1 -= (v1 << ((42) & 63))
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        v4 = append(v4, "a")
        if true {
            var v22 uint = 65535
            _ = v22
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v23 int16 = v3
            _ = v23
        }
        fmt.Println(v3)
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
