package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint64, p1 int16) int64 {
    p0 = (p0 & 1)
    p1 -= (p1 | p1)
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 int32 = 42
    var v2 uint32 = 800
    var v3 bool = (float64((3.14159 / (-2.25 * (0.0 - (-2.25 / 3.14159))))) > (3.14159 - 1e10))
    var v4 bool = false
    var v5 uint32 = (^(v2))
    v6 := make(map[string]string)
    if v4 {
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 int = (-(fi7))
            v9 := []int{fi7, (v8 - fi7), 128}
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 bool = v3
        for fi11 := 0; fi11 < 5; fi11++ {
            if v10 {
                if (!(1e10 != 0.1)) {
                    fmt.Println(fi11, v1)
                    v12 := []int{fi11, 1}
                    _ = v12
                }
            }
            _ = fi11
        }
        _ = v10
    } else {
        for fi13 := 0; fi13 < 18; fi13++ {
            v14 := []int{(fi13 | fi13), 42, (fi13 + (-(fi13)))}
            _ = fi13
            _ = v14
        }
    }
    var v15 int64 = genFn0(1, -445)
    fmt.Println(v1)
    v16 := []int{58, -237, -903, 127}
    if true {
        _ = len(v16)
        switch ((("foo bar" + "foo bar") + "foo bar") + "a") {
        case "":
            var v17 string = "foo bar"
            _ = v17
        case "\n\t":
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        }
    } else {
        v2 = (v5 - v2)
        v4 = v3
    }
    v5 += (v2 | v5)
    if ((127 == 105) && ((4294967295 < 2) && ((-568 <= 505) && v3))) {
        fmt.Println(v15, v5)
    }
    v1 -= v1
    v15 = -945
    if (79 > 2) {
        fmt.Println(v15)
        v6["hello"] = "hello"
        fmt.Println(v5, v5, v3)
    } else {
        v3 = (v15 > (v15 >> ((839) & 63)))
        v18 := GenS1{F0: 94}
        _ = v18
    }
    if true {
        v3 = (83 < 4294967295)
    } else {
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    }
    delete(v6, "hello")
    v1 += v1
    for fi19 := 0; fi19 < 19; fi19++ {
        var v20 string = ""
        v21 := GenS1{F0: 18446744073709551615}
        var v22 int64 = -9223372036854775808
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    v23 := GenS1{F0: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
