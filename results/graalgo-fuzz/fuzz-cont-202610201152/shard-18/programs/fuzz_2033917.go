package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 bool }
func genFn0(p0 bool) uint8 {
    var v6 int64 = 42
    if true {
        fmt.Println(p0)
    }
    v7 := GenS1{F0: -644}
    _ = p0
    _ = v6
    _ = v7
    return 255
}

func main() {
    var v1 bool = (!(2 < 18446744073709551615))
    v2 := []string{"a"}
    v3 := []int{32768}
    v4 := []string{"a"}
    var v5 string = "a"
    v8 := GenS1{F0: -1}
    fmt.Println(v5, v1)
    v5 = "foo bar"
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 int64 = -9223372036854775808
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: ((("foo bar" + (v5 + "\n\t")) + v5) + "\n\t")}
    v12 := GenS2{F0: ""}
    if (1000 < 1) {
        var v13 uint32 = 886
        var v14 uint8 = genFn0(false)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v13
        _ = v14
    } else {
        var v15 uint32 = 65535
        v16 := GenS1{F0: 657}
        var v17 int8 = 0
        _ = v15
        _ = v16
        _ = v17
    }
    var v18 bool = (42 >= 42)
    switch 42 {
    case 128:
        fmt.Println(v18, v5)
    case -398:
        for fi19 := 0; fi19 < 4; fi19++ {
            v20 := GenS1{F0: 843}
            var v21 float32 = 1.5
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    if (!((127 != 42) && true)) {
        var v22 float64 = 1.5
        var v23 string = ((v5 + v5) + "\n\t")
        _ = v22
        _ = v23
    } else {
        if v18 {
            var v24 string = (v5 + v5)
            fmt.Println(v18)
            if v18 {
                var v25 uint32 = 2147483648
                _ = v25
            }
            _ = v24
        }
        var v26 uint32 = 2
        var v27 uint8 = 247
        _ = v26
        _ = v27
    }
    v28 := GenS1{F0: 1}
    v29 := GenS3{F0: genFn0((true && true)), F1: (v18 && (9223372036854775807 < 330))}
    var v30 float64 = 1.5
    var v31 int8 = 127
    var v32 int8 = 0
    v30 = (v30 / v30)
    if false {
        if len(v2) > 0 {
            v2[0] = v5
        }
        for fi33 := 0; fi33 < 11; fi33++ {
            v2 = append(v2, v5)
            v34 := GenS1{F0: 1}
            _ = fi33
            _ = v34
        }
        var v35 int8 = v32
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v8
}
