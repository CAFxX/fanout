package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() int16 {
    v5 := GenS1{F0: -602}
    var v6 uint32 = 927
    if false {
        switch "" {
        case "foo bar":
            fmt.Println(v6)
        case "hello":
            v7 := []int{32767, 4294967295, 42}
            _ = v7
        case "\n\t":
            v8 := []int{9223372036854775807}
            _ = v8
        default:
            var v9 int8 = 42
            _ = v9
        }
    } else {
        v6 -= 1
        if (-2.25 == (((-2.25 - 1.5) - -2.25) / float32(float64((1.5 / ((3.14159 - ((0.1 - -2.25) - 3.14159)) * 3.14159)))))) {
            v6 += 0
        }
    }
    _ = v5
    _ = v6
    return -187
}

func main() {
    v1 := []string{"", "foo bar"}
    var v2 int8 = 100
    var v3 uint = 255
    var v4 int = -2147483648
    var v10 int = (^(v4))
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 uint = (v3 | v3)
        if (v12 >= 2) {
            for fi13 := 0; fi13 < 4; fi13++ {
                v4 -= (fi11 - v4)
                if true {
                    v12 -= 65535
                    var v14 bool = (float64(0.0) > 3.14159)
                    if v14 {
                        v4 = v10
                    }
                    _ = v14
                } else {
                    var v15 int64 = 146
                    v1 = append(v1, "")
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                    _ = v15
                }
                _ = fi13
            }
            if ((genFn0() > 804) || (!true)) {
                var v16 float64 = -2.25
                v10 += -9223372036854775808
                var v17 uint16 = 1
                _ = v16
                _ = v17
            }
            v1 = append(v1, "\n\t")
        }
        v10 += v4
        _ = fi11
        _ = v12
    }
    var v18 uint = v3
    v2 -= v2
    fmt.Println(v4, v2, v10)
    var v19 int8 = 94
    v1 = append(v1, (("hello" + (("" + (("a" + "") + "a")) + ("a" + ""))) + "\n\t"))
    v19 += (v2 + 42)
    var v20 uint64 = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
