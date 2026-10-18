package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int32 }
func genFn0(p0 int, p1 uint64, p2 int32) bool {
    var v4 int8 = 42
    v5 := GenS1{F0: v4}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return true
}

func main() {
    var v1 string = "foo bar"
    var v2 float32 = (0.1 + (0.1 + 0.1))
    var v3 uint64 = 967
    var v6 int64 = 9223372036854775807
    switch 2147483647 {
    case 42:
        var v7 int64 = 42
        _ = v7
    }
    var v8 string = (((((("\n\t" + "") + v1) + ("" + v1)) + v1) + v1) + "\n\t")
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 uint = 294
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        if true {
            fmt.Println(v1, v6, v8)
        } else {
            switch (fi11 / ((fi11) | 1)) {
            case 32768:
                for fi12 := 0; fi12 < 6; fi12++ {
                    var v13 uint8 = 255
                    if (fi11 <= (^(fi12))) {
                        v3 += v3
                        var v14 uint16 = 255
                        v15 := []int{(fi11 * fi12)}
                        _ = v14
                        _ = v15
                    }
                    var v16 int8 = -5
                    _ = fi12
                    _ = v13
                    _ = v16
                }
            case 812:
                for fi17 := 0; fi17 < 13; fi17++ {
                    var v18 bool = (-375 <= 32767)
                    var v19 int8 = 127
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            case -2147483648:
                if (!((!(-3 != 1)) || (0.1 >= 0.1))) {
                    v20 := []int{}
                    for fi21 := 0; fi21 < 3; fi21++ {
                        v3 = 56
                        _ = fi21
                    }
                    fmt.Println(v6)
                    _ = v20
                }
            }
            v22 := GenS1{F0: -1}
            _ = v22
        }
        fmt.Println(v3)
        _ = fi11
    }
    var v23 int16 = 0
    for fi24 := 0; fi24 < 20; fi24++ {
        v25 := GenS1{F0: 0}
        var v26 uint8 = 1
        _ = fi24
        _ = v25
        _ = v26
    }
    switch v8 {
    case "hello":
        fmt.Println(v6)
    case "":
        var v27 float64 = float64((3.14159 + v2))
        _ = v27
    case "a":
        var v28 int8 = 127
        _ = v28
    default:
        fmt.Println(v8, v23)
    }
    v1 = "\n\t"
    var v29 int16 = (v23 * v23)
    for fi30 := 0; fi30 < 7; fi30++ {
        if false {
            v31 := []int{(^(fi30)), 42, 127}
            v23 += 0
            _ = v31
        } else {
            var v32 string = "\n\t"
            if false {
                v33 := []int{fi30, fi30, 362, (^(fi30))}
                _ = v33
            } else {
                var v34 uint32 = 594
                _ = v34
            }
            _ = v32
        }
        _ = fi30
    }
    v35 := GenS2{F0: -2147483648, F1: -899}
    v36 := []int{}
    v8 = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v6
    _ = v8
}
