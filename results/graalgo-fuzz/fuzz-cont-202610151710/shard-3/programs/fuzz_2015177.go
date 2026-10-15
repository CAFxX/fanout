package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int16, p1 uint, p2 uint16) uint16 {
    var v7 uint32 = 490
    switch (("" + "a") + (("" + "a") + "")) {
    case "a":
        switch "a" {
        case "foo bar":
            v8 := []int{-794, 32768}
            _ = v8
        case "\n\t":
            fmt.Println(p0, p1, p2)
        case "a":
            v9 := []int{764, 32768}
            _ = v9
        }
    case "\n\t":
        var v10 uint64 = 485
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p2 - p2)
}

func main() {
    var v1 int32 = -2147483648
    v2 := []int{9223372036854775807, 9223372036854775807, 490}
    var v3 int16 = -1
    var v4 string = "a"
    v5 := []int{127, -1}
    v6 := make(map[string]string)
    v4 = ("\n\t" + v4)
    for fi11 := 0; fi11 < 8; fi11++ {
        fmt.Println(fi11)
        v12 := GenS1{F0: 42}
        var v13 uint16 = 247
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := v6[(("hello" + ("\n\t" + "")) + v4)]
    switch v4 {
    case "\n\t":
        if false {
            v5 = append(v5, -833)
        }
    case "foo bar":
        for fi15 := 0; fi15 < 3; fi15++ {
            for fi16 := 0; fi16 < 8; fi16++ {
                v3 += (v3 >> ((18446744073709551615) & 63))
                if false {
                    v14 = ""
                    v3 += v3
                } else {
                    v4 = (v4 + v4)
                }
                _ = fi16
            }
            fmt.Println(fi15)
            v17 := GenS1{F0: 596}
            _ = fi15
            _ = v17
        }
    case "a":
        if (0.0 >= (-2.25 / 1.0)) {
            var v18 int8 = 127
            v19 := GenS1{F0: 0}
            v20 := v6[v14]
            _ = v18
            _ = v19
            _ = v20
        } else {
            for fi21 := 0; fi21 < 4; fi21++ {
                v1 -= -989
                var v22 bool = (v14 != v14)
                _ = fi21
                _ = v22
            }
            fmt.Println(v3, v4)
            _ = len(v2)
        }
    }
    v1 -= 0
    var v23 int32 = v1
    fmt.Println(v23, v14, v14)
    _ = len(v6)
    if (!(0 < 42)) {
        var v24 string = v4
        fmt.Println(v14, v24)
        v25 := v6[(v4 + "\n\t")]
        _ = v24
        _ = v25
    }
    var v26 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
