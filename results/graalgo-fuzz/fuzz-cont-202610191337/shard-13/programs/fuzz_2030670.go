package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint16) uint8 {
    fmt.Println(p0)
    var v4 uint16 = (p0 % ((65535) | 1))
    _ = p0
    _ = v4
    return 2
}
func genFn1(p0 int8, p1 bool) string {
    v5 := GenS1{F0: 42}
    fmt.Println(p0, p0)
    v6 := []int{}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return ("hello" + ("\n\t" + "a"))
}

func main() {
    var v1 int32 = 854
    var v2 uint = 0
    v3 := []int{127, -2147483648, 1, 9223372036854775807}
    if len(v3) > 0 {
        v3[0] = 521
    }
    v1 = v1
    var v7 float64 = (1e10 + -2.25)
    if false {
        switch ("\n\t" + (("" + "hello") + "")) {
        case "\n\t":
            if (2 != 298) {
                if (true && true) {
                    v7 = v7
                    var v8 uint = v2
                    _ = v8
                }
                for fi9 := 0; fi9 < 18; fi9++ {
                    var v10 float32 = 1.5
                    _ = fi9
                    _ = v10
                }
            }
        case "":
            v11 := GenS1{F0: -1}
            _ = v11
        }
        v12 := GenS2{F0: 18446744073709551615}
        v1 -= 793
        _ = v12
    }
    var v13 bool = true
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    var v14 int = -2147483648
    if (4294967295 <= 1) {
        for fi15 := 0; fi15 < 3; fi15++ {
            v13 = (!v13)
            v16 := GenS1{F0: int8(fi15)}
            _ = fi15
            _ = v16
        }
        fmt.Println(v14, v13, v1)
        var v17 uint8 = 1
        _ = v17
    }
    var v18 uint8 = 217
    if ((v13 || ((true || (42 < 0)) || v13)) || (!v13)) {
        if true {
            var v19 int8 = 49
            _ = v19
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v7
}
