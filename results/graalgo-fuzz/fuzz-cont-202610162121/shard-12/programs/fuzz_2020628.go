package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int16) int32 {
    p0 = -1
    v4 := []int{42, 42, 346}
    var v5 int16 = (p0 & 1)
    switch 100 {
    case -354:
        for fi6 := 0; fi6 < 17; fi6++ {
            fmt.Println(p0)
            _ = fi6
        }
    case 42:
        p0 += (p0 | (-(v5)))
    }
    v5 = 0
    _ = p0
    _ = v4
    _ = v5
    return -2147483648
}

func main() {
    v1 := []string{"foo bar", ("a" + ""), "foo bar", "hello"}
    v2 := make(map[string]uint32)
    var v3 uint16 = 2
    switch "a" {
    case "":
        fmt.Println(v3)
    case "hello":
        for fi7 := 0; fi7 < 12; fi7++ {
            switch (("foo bar" + "") + "foo bar") {
            case "a":
                if false {
                    v3 = (v3 - v3)
                }
            }
            v8 := GenS1{F0: (v3 - 2)}
            v3 += (v3 << ((758) & 63))
            _ = fi7
            _ = v8
        }
    case "a":
        v3 += (^(v3))
    }
    fmt.Println(v3)
    for fi9 := 0; fi9 < 4; fi9++ {
        v3 = v3
        _ = fi9
    }
    var v10 int32 = 2147483647
    if false {
        var v11 uint64 = 0
        _ = v11
    }
    v10 += (-(v10))
    for fi12 := 0; fi12 < 4; fi12++ {
        v13 := GenS1{F0: 0}
        var v14 uint64 = 1
        for fi15 := 0; fi15 < 11; fi15++ {
            if true {
                if (false || (!false)) {
                    if true {
                        var v16 string = ("\n\t" + (("" + "\n\t") + ("" + "")))
                        var v17 int64 = 9223372036854775807
                        v18 := GenS1{F0: (v3 - (v3 * v3))}
                        _ = v16
                        _ = v17
                        _ = v18
                    } else {
                        fmt.Println(v3, v3, v10)
                        v10 += v10
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                fmt.Println(v10)
                fmt.Println(v14)
            }
            _ = len(v1)
            _ = fi15
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = len(v1)
    v2[("" + "\n\t")] = 1
    v19 := GenS1{F0: v3}
    switch "foo bar" {
    case "":
        v10 += (v10 << ((2) & 63))
    case "a":
        v20 := v2["foo bar"]
        _ = v20
    case "\n\t":
        v3 -= 197
    }
    fmt.Println(v3, v10)
    if ((!false) || (!false)) {
        var v21 int16 = -988
        var v22 int16 = -32768
        var v23 float64 = 3.14159
        _ = v21
        _ = v22
        _ = v23
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println(v10, v10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
}
