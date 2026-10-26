package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 string }
type GenS2 struct { F0 string; F1 int8 }

func main() {
    v1 := make(map[string]int)
    var v2 string = ("" + ("hello" + "a"))
    var v3 int32 = 42
    var v4 string = v2
    delete(v1, "\n\t")
    v3 -= -704
    v4 = v4
    var v5 float32 = (0.1 / -2.25)
    fmt.Println(v4, v3, v5)
    if true {
        v6 := GenS1{F0: 42, F1: (!((-128 < -128) && (!(!("a" <= (v2 + "\n\t")))))), F2: v2}
        for fi7 := 0; fi7 < 17; fi7++ {
            for fi8 := 0; fi8 < 3; fi8++ {
                v3 = v3
                var v9 int32 = (v3 / ((0) | 1))
                _ = fi8
                _ = v9
            }
            v10 := GenS1{F0: int64(fi7), F1: true, F2: "a"}
            v11 := GenS1{F0: -9223372036854775808, F1: ((!true) && false), F2: v4}
            _ = fi7
            _ = v10
            _ = v11
        }
        _ = v6
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := GenS1{F0: 0, F1: (fi12 < fi12), F2: "\n\t"}
        _ = fi12
        _ = v13
    }
    v14 := []int{32767, -444, 848}
    v15 := GenS2{F0: v2, F1: 42}
    v3 += (v3 | v3)
    v16 := GenS1{F0: -1, F1: (416 >= 42), F2: "hello"}
    v2 = v4
    if ((!(true || false)) && false) {
        if (((253 != 72) || false) || true) {
            var v17 uint32 = 2147483648
            var v18 uint16 = 330
            _ = v17
            _ = v18
        } else {
            fmt.Println(v3)
            v5 += (0.1 / 0.1)
        }
        fmt.Println(v2)
        for fi19 := 0; fi19 < 5; fi19++ {
            for fi20 := 0; fi20 < 15; fi20++ {
                v21 := v1[v4]
                _ = fi20
                _ = v21
            }
            v22 := GenS2{F0: "a", F1: -128}
            _ = fi19
            _ = v22
        }
    } else {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
