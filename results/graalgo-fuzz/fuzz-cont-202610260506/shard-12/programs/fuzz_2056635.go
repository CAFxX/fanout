package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint16 }
func genFn0() uint64 {
    var v6 int64 = -886
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(fi7)
        v8 := []int{}
        v6 = v6
        _ = fi7
        _ = v8
    }
    if (0.0 > 0.1) {
        v9 := GenS1{F0: ("\n\t" + "")}
        _ = v9
    } else {
        v6 = v6
    }
    _ = v6
    return 29
}

func main() {
    var v1 float64 = 3.14159
    var v2 int16 = 513
    v3 := []string{}
    var v4 int = -2147483648
    var v5 string = ("a" + ("hello" + (("hello" + "a") + "a")))
    v2 -= -32768
    v4 = v4
    v3 = append(v3, v5)
    if (1 >= 0) {
        var v10 uint8 = 127
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if false {
            var v11 string = ((((v5 + "hello") + v5) + v5) + v5)
            var v12 uint64 = 35
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
            _ = v11
            _ = v12
        } else {
            var v13 uint8 = 42
            v14 := GenS2{F0: -2147483648, F1: 811}
            v15 := GenS1{F0: (("" + "\n\t") + "\n\t")}
            _ = v13
            _ = v14
            _ = v15
        }
        _ = v10
    }
    if false {
        v1 += (1.5 * ((float64(0.0) + 1e10) + 1.5))
    } else {
        v4 = (v4 & 1)
        var v16 uint16 = 981
        _ = v16
    }
    if true {
        for fi17 := 0; fi17 < 15; fi17++ {
            var v18 int = 42
            var v19 float64 = v1
            _ = fi17
            _ = v18
            _ = v19
        }
        var v20 int64 = 0
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v20
    } else {
        v4 = 127
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        if len(v3) > 0 {
            v3[0] = (v5 + v5)
        }
        _ = fi21
    }
    var v22 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
