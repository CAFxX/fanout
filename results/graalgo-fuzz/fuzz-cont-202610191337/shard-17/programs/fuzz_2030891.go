package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
func genFn0(p0 int32) string {
    var v7 uint32 = 472
    fmt.Println(p0, v7)
    fmt.Println(p0)
    v8 := []int{-935, 4294967295}
    p0 -= p0
    _ = p0
    _ = v7
    _ = v8
    return "foo bar"
}
func genFn1() uint {
    var v9 string = "foo bar"
    if (true || false) {
        var v10 float32 = (0.1 * (0.0 / -2.25))
        v11 := []int{}
        var v12 uint64 = 42
        _ = v10
        _ = v11
        _ = v12
    } else {
        v13 := []int{4294967295}
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        _ = v13
    }
    var v14 int16 = -1
    v15 := []int{552, -2147483648, -545, 9223372036854775807}
    _ = v14
    _ = v15
    _ = v9
    return 255
}

func main() {
    var v1 int8 = 0
    var v2 uint = 1
    v3 := []int{}
    var v4 string = "foo bar"
    v5 := make(map[string]bool)
    var v6 float64 = 1e10
    delete(v5, "hello")
    v16 := GenS1{F0: 2}
    if true {
        var v17 int16 = 42
        fmt.Println(v2, v17, v17)
        var v18 uint16 = 255
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        fmt.Println(v4, fi19, v1)
        fmt.Println(fi19, v4, v6)
        _ = fi19
    }
    var v20 int16 = 1000
    for fi21 := 0; fi21 < 11; fi21++ {
        v3 = append(v3, fi21)
        _ = fi21
    }
    v20 = v20
    var v22 string = (v4 + v4)
    v20 = (^(v20))
    var v23 int32 = 42
    v5[v22] = (!true)
    v20 = -254
    if ((-(v23)) == v23) {
        var v24 string = (((v22 + (v22 + "\n\t")) + "foo bar") + (("\n\t" + (v4 + v4)) + (((v22 + "") + v22) + v22)))
        v25 := GenS2{F0: "foo bar"}
        var v26 float64 = (v6 + 0.1)
        _ = v24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        v6 += v6
        switch "" {
        case "foo bar":
            v1 += (-(v1))
        }
        for fi28 := 0; fi28 < 8; fi28++ {
            v29 := GenS1{F0: 0}
            var v30 uint16 = 32767
            if (217 != 957) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = fi27
    }
    fmt.Println(v22, v22)
    var v31 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
