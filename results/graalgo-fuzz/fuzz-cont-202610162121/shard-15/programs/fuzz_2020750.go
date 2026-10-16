package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint8, p1 bool) int32 {
    v5 := []int{42}
    var v6 int = -1
    var v7 uint8 = p0
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return -1
}

func main() {
    v1 := []string{"hello", ("\n\t" + "a"), "hello", "a"}
    var v2 float32 = 3.14159
    var v3 int = 1
    var v4 int32 = -859
    var v8 int64 = -599
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "foo bar" {
    case "a":
        var v9 string = ((("foo bar" + "a") + ("" + "hello")) + (("\n\t" + "foo bar") + "hello"))
        _ = v9
    }
    v2 += float32((float64((1e10 * ((0.1 + v2) * v2))) - 3.14159))
    for fi10 := 0; fi10 < 4; fi10++ {
        for fi11 := 0; fi11 < 12; fi11++ {
            v3 -= (fi10 + fi11)
            var v12 int32 = -2147483648
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    var v13 int16 = 1000
    v13 = v13
    var v14 int8 = -119
    v14 = v14
    for fi15 := 0; fi15 < 16; fi15++ {
        v2 -= ((v2 - v2) + v2)
        _ = fi15
    }
    if true {
        v14 = 11
    }
    v16 := GenS1{F0: 0}
    var v17 int = (v3 * 962)
    v18 := GenS2{F0: 4294967295}
    switch "a" {
    case "hello":
        if (!((42 > 1) && (false || (true || (255 < 410))))) {
            for fi19 := 0; fi19 < 14; fi19++ {
                var v20 uint32 = 0
                var v21 string = (("a" + "a") + "")
                _ = fi19
                _ = v20
                _ = v21
            }
        }
    case "\n\t":
        var v22 uint64 = 823
        _ = v22
    case "a":
        switch int(v8) {
        case 32768:
            v8 += v8
        }
    }
    fmt.Println(v4, v13, v13)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
