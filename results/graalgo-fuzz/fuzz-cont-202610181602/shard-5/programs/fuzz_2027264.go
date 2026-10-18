package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }

func main() {
    var v1 uint16 = 1
    var v2 int64 = -9223372036854775808
    var v3 int64 = (^(v2))
    var v4 int8 = 17
    var v5 uint16 = 65535
    var v6 int16 = 0
    switch "\n\t" {
    case "foo bar":
        switch "foo bar" {
        case "":
            v7 := GenS1{F0: 2147483647, F1: 1}
            _ = v7
        default:
            v5 = v1
        }
    case "\n\t":
        if true {
            fmt.Println(v5)
            var v8 float64 = (float64(1.5) / ((3.14159 * 1.5) - -2.25))
            v4 -= v4
            _ = v8
        } else {
            v1 = v1
            var v9 uint32 = 159
            v10 := []int{2147483647, 375}
            _ = v10
            _ = v9
        }
    case "":
        if true {
            var v11 float64 = 3.14159
            _ = v11
        }
    default:
        v6 = v6
    }
    fmt.Println(v3, v3)
    v5 = v5
    for fi12 := 0; fi12 < 4; fi12++ {
        if (852 > fi12) {
            v3 = v2
            fmt.Println(v2)
        } else {
            switch 0 {
            case -978:
                v13 := []int{fi12, fi12, fi12, fi12}
                _ = v13
            case 191:
                v4 += -128
            }
            var v14 int64 = 9223372036854775807
            v4 = (v4 << ((65535) & 63))
            _ = v14
        }
        switch "\n\t" {
        case "":
            v15 := []int{(fi12 & 32767)}
            _ = v15
        case "\n\t":
            v16 := []int{-1, (fi12 * fi12), fi12, fi12}
            _ = v16
        }
        _ = fi12
    }
    fmt.Println(v4)
    var v17 string = "a"
    v6 = 0
    v17 = v17
    v1 += 2
    fmt.Println(v3, v17)
    switch 1 {
    case 692:
        v18 := GenS1{F0: 42, F1: 1}
        _ = v18
    case 9223372036854775807:
        for fi19 := 0; fi19 < 9; fi19++ {
            var v20 string = v17
            v2 += v2
            var v21 float32 = 1e10
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    if ((!(1 == 0)) || (true && false)) {
        var v22 string = "a"
        _ = v22
    }
    v23 := []int{0, 0, 127, 469}
    v4 += v4
    for fi24 := 0; fi24 < 15; fi24++ {
        v25 := GenS1{F0: 599, F1: 42}
        _ = fi24
        _ = v25
    }
    fmt.Println(v2)
    for fi26 := 0; fi26 < 4; fi26++ {
        for fi27 := 0; fi27 < 5; fi27++ {
            var v28 int16 = (v6 * (v6 - (^(v6))))
            v4 = v4
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        for fi30 := 0; fi30 < 1; fi30++ {
            if ((fi30 + fi29) == (fi30 >> ((4294967295) & 63))) {
                var v31 uint16 = v5
                fmt.Println(v1, fi29, fi30)
                v4 = 0
                _ = v31
            } else {
                if (!false) {
                    if true {
                        v32 := GenS1{F0: -2147483648, F1: 4294967295}
                        _ = v32
                    }
                    switch ("a" + v17) {
                    case "":
                        fmt.Println(v2, v2, v3)
                    default:
                        switch "hello" {
                        case "\n\t":
                            var v33 uint8 = 172
                            var v34 uint64 = 780
                            _ = v33
                            _ = v34
                        case "a":
                            _ = len(v23)
                        default:
                            v4 = (^(v4))
                        }
                    }
                    var v35 int32 = 2147483647
                    var v36 bool = ((12 != 939) || true)
                    _ = v35
                    _ = v36
                }
                fmt.Println(v6, v4, v3)
                fmt.Println(v2, v6, v3)
            }
            _ = fi30
        }
        v37 := GenS1{F0: 1, F1: 0}
        v17 = "foo bar"
        _ = fi29
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
