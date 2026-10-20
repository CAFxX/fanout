package main

import (
    "fmt"
)

func genFn0(p0 string) uint32 {
    p0 = (p0 + p0)
    v5 := []int{0}
    _ = p0
    _ = v5
    return 470
}
func genFn1(p0 int64, p1 int, p2 int16) string {
    if ((!true) && false) {
        p0 += -444
        p1 = (-(p1))
    }
    fmt.Println(p0)
    var v6 uint = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "foo bar"
}

func main() {
    var v1 uint32 = 2147483648
    var v2 float64 = 3.14159
    var v3 bool = (0 != 1000)
    v4 := make(map[uint32]string)
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 bool = (42 > 4294967295)
        var v9 int64 = 2
        v3 = (!(-599 != 507))
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        v2 = 1e10
        v1 -= v1
        _ = fi10
    }
    v3 = false
    for fi11 := 0; fi11 < 3; fi11++ {
        v1 = 2147483648
        switch (fi11 << ((779) & 63)) {
        case 128:
            var v12 string = ""
            _ = v12
        }
        _ = fi11
    }
    v4[v1] = "\n\t"
    if false {
        var v13 int16 = -1
        fmt.Println(v3, v3, v1)
        _ = v13
    }
    v14 := v4[(v1 - (v1 / ((42) | 1)))]
    var v15 int32 = -2147483648
    fmt.Println(v14)
    if true {
        _ = len(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
