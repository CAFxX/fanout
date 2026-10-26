package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint64 }
type GenS2 struct { F0 uint; F1 uint; F2 int16 }
func genFn0(p0 uint64) int {
    if false {
        var v6 int16 = -1
        v7 := GenS1{F0: (!(769 >= 2)), F1: 257, F2: (p0 & p0)}
        var v8 int32 = -2147483648
        _ = v6
        _ = v7
        _ = v8
    }
    switch "\n\t" {
    case "hello":
        if false {
            for fi9 := 0; fi9 < 15; fi9++ {
                var v10 uint64 = (^(p0))
                fmt.Println(p0, fi9)
                v10 += (p0 | p0)
                _ = fi9
                _ = v10
            }
        }
    case "\n\t":
        v11 := GenS2{F0: 2, F1: 580, F2: 838}
        _ = v11
    default:
        v12 := []int{4294967295, -9223372036854775808}
        _ = v12
    }
    _ = p0
    return 35
}

func main() {
    var v1 bool = false
    v2 := make(map[string]bool)
    var v3 uint64 = 0
    var v4 float64 = 0.0
    v5 := make(map[int]int)
    for fi13 := 0; fi13 < 3; fi13++ {
        if (2147483647 == 0) {
            v4 += 3.14159
            switch ("foo bar" + "a") {
            case "foo bar":
                var v14 uint32 = 65535
                _ = v14
            case "\n\t":
                v15 := []int{fi13}
                _ = v15
            case "a":
                var v16 uint64 = 0
                _ = v16
            }
            for fi17 := 0; fi17 < 14; fi17++ {
                var v18 uint = 18446744073709551615
                _ = fi17
                _ = v18
            }
        }
        if v1 {
            if v1 {
                v4 += float64(-2.25)
            } else {
                var v19 string = "hello"
                _ = v19
            }
            var v20 int16 = -566
            _ = v20
        }
        _ = fi13
    }
    fmt.Println(v4, v1)
    if v1 {
        fmt.Println(v4, v4, v4)
        for fi21 := 0; fi21 < 10; fi21++ {
            for fi22 := 0; fi22 < 12; fi22++ {
                var v23 string = ("" + ("hello" + ("\n\t" + "hello")))
                fmt.Println(v1)
                if (!false) {
                    if (v1 || ((!false) || v1)) {
                        var v24 uint32 = 42
                        _ = v24
                    }
                    var v25 float64 = (v4 + 1e10)
                    _ = v25
                }
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
        fmt.Println(v1)
    }
    v26 := GenS2{F0: 0, F1: 862, F2: 1}
    var v27 int8 = 1
    v3 += v3
    for fi28 := 0; fi28 < 19; fi28++ {
        var v29 uint64 = (v3 << ((364) & 63))
        if (!v1) {
            var v30 uint8 = 2
            v31 := v2[""]
            _ = v30
            _ = v31
        }
        fmt.Println(v3, fi28)
        _ = fi28
        _ = v29
    }
    v1 = (false && true)
    v4 = v4
    if v1 {
        v4 = 0.0
        var v32 int64 = -1
        for fi33 := 0; fi33 < 6; fi33++ {
            v5[fi33] = fi33
            v34 := []int{fi33, 42, fi33}
            v3 += 99
            _ = fi33
            _ = v34
        }
        _ = v32
    }
    v4 = (3.14159 / 1.5)
    v35 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
