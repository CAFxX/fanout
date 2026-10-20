package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint64 }
func genFn0(p0 bool) int16 {
    fmt.Println(p0)
    fmt.Println(p0)
    v6 := GenS1{F0: 513, F1: 2, F2: 2}
    p0 = (!p0)
    _ = p0
    _ = v6
    return 1000
}

func main() {
    var v1 uint32 = 65535
    v2 := []string{"\n\t", ("a" + (("hello" + "a") + "\n\t")), ("a" + "foo bar"), ("" + "")}
    var v3 int = 2147483647
    var v4 bool = false
    var v5 uint16 = 0
    for fi7 := 0; fi7 < 10; fi7++ {
        if v4 {
            for fi8 := 0; fi8 < 1; fi8++ {
                if (!((1000 > 42) || false)) {
                    var v9 int16 = 1
                    var v10 int32 = -825
                    _ = v10
                    _ = v9
                }
                if (!(18446744073709551615 < 0)) {
                    v1 += (v1 ^ (v1 << ((18446744073709551615) & 63)))
                    var v11 float64 = (3.14159 * 1.5)
                    _ = v11
                }
                _ = fi8
            }
            var v12 float64 = 3.14159
            _ = v12
        }
        _ = fi7
    }
    if (!(0 >= 72)) {
        var v13 float64 = (1e10 / -2.25)
        v13 = (v13 * (0.0 * (-2.25 - (v13 / ((3.14159 * 0.1) + v13)))))
        switch (("a" + "hello") + "hello") {
        case "foo bar":
            if len(v2) > 0 {
                v2[0] = "hello"
            }
        case "":
            _ = len(v2)
        case "a":
            v14 := GenS1{F0: 2, F1: 0, F2: 18446744073709551615}
            _ = v14
        default:
            _ = len(v2)
        }
        _ = v13
    } else {
        v15 := GenS1{F0: 666, F1: 242, F2: 42}
        _ = v15
    }
    fmt.Println(v5)
    var v16 uint16 = (v5 << ((0) & 63))
    var v17 string = ("\n\t" + "hello")
    for fi18 := 0; fi18 < 11; fi18++ {
        v4 = (v17 == (v17 + v17))
        for fi19 := 0; fi19 < 3; fi19++ {
            fmt.Println(v1, v1)
            var v20 int8 = -109
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    v17 = v17
    v16 += (v5 & 1)
    var v21 uint16 = 42
    if (v17 == "a") {
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 int16 = 42
            _ = fi22
            _ = v23
        }
        var v24 float64 = (0.0 - 3.14159)
        _ = v24
    } else {
        var v25 float32 = 3.14159
        var v26 uint64 = 481
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
