package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
func genFn0(p0 int8, p1 uint32) uint16 {
    p0 = p0
    for fi7 := 0; fi7 < 9; fi7++ {
        p0 -= 1
        v8 := GenS1{F0: 18446744073709551615}
        fmt.Println(p0, p1)
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    return 255
}

func main() {
    var v1 bool = ((("" == ("a" + "foo bar")) || false) || (-2.25 >= float32(-2.25)))
    var v2 float32 = float32(0.0)
    v3 := make(map[int]string)
    var v4 uint = 2
    v5 := make(map[int]bool)
    var v6 int = 760
    v5[v6] = v1
    fmt.Println(v2)
    if (!("" <= "")) {
        v4 = v4
        var v9 string = "\n\t"
        _ = v9
    }
    var v10 bool = v1
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := GenS2{F0: ("\n\t" + "hello")}
        _ = fi11
        _ = v12
    }
    if v1 {
        var v13 bool = v1
        if false {
            switch "" {
            case "a":
                _ = len(v5)
            case "\n\t":
                var v14 uint32 = 65535
                _ = v14
            }
            v6 += v6
        } else {
            if false {
                var v15 int8 = -38
                _ = v15
            } else {
                var v16 bool = (994 > (v4 / ((1) | 1)))
                v4 += (v4 & 653)
                _ = v16
            }
            v10 = (false && false)
        }
        _ = v13
    }
    fmt.Println(v2, v4)
    var v17 int16 = 335
    if (!v1) {
        v10 = (18446744073709551615 >= 666)
    }
    v18 := []int{-725, (v6 + -2147483648)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
