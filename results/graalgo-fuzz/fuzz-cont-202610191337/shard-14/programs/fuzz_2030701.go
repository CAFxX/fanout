package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 string }
func genFn0(p0 int64) string {
    var v7 uint64 = 42
    switch ("\n\t" + ("hello" + "hello")) {
    case "\n\t":
        v7 = 4294967295
    case "foo bar":
        v8 := []int{-9223372036854775808, 32767, 9223372036854775807}
        _ = v8
    }
    v9 := []int{42}
    _ = p0
    _ = v7
    _ = v9
    return "foo bar"
}
func genFn1(p0 uint64, p1 uint) string {
    var v10 int64 = 9223372036854775807
    if true {
        v11 := GenS1{F0: 1000, F1: (-36 >= -60), F2: ("" + "a")}
        fmt.Println(p1, p1)
        _ = v11
    } else {
        for fi12 := 0; fi12 < 10; fi12++ {
            v13 := []int{4294967295}
            _ = fi12
            _ = v13
        }
    }
    v14 := GenS1{F0: 139, F1: (1e10 <= 1.5), F2: "\n\t"}
    _ = p0
    _ = p1
    _ = v10
    _ = v14
    return (("" + "foo bar") + "\n\t")
}

func main() {
    var v1 float32 = 0.0
    v2 := []uint32{4294967295, 2147483647, 0, 1}
    var v3 float32 = (1e10 / float32(-2.25))
    v4 := []uint32{42, 42, 1}
    var v5 int = 127
    v6 := make(map[string]int)
    v15 := GenS1{F0: 1000, F1: (0.0 >= float64(v3)), F2: "\n\t"}
    if (v5 > (-(v5))) {
        fmt.Println(v3)
    } else {
        for fi16 := 0; fi16 < 3; fi16++ {
            v1 -= v3
            v3 = v1
            if (false && false) {
                v3 = v3
                if (fi16 >= v5) {
                    for fi17 := 0; fi17 < 7; fi17++ {
                        v4 = append(v4, 42)
                        var v18 int64 = -353
                        var v19 bool = ((!(!true)) && false)
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                }
                _ = len(v6)
            }
            _ = fi16
        }
    }
    var v20 uint16 = 0
    v3 += (-2.25 / 0.1)
    var v21 uint = 42
    var v22 int8 = -69
    var v23 string = ""
    var v24 uint64 = 4294967295
    v5 = int(v21)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
