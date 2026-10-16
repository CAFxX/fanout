package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int16 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 int; F1 int32; F2 int16 }
type GenS5 struct { F0 bool; F1 int32; F2 int8 }
type GenS6 struct { F0 string; F1 int64 }

func main() {
    var v1 uint = 255
    var v2 string = "foo bar"
    var v3 int8 = -66
    var v4 uint = (v1 & v1)
    var v5 float32 = -2.25
    switch v2 {
    case "\n\t":
        v6 := []int{127, -377, -2147483648, 9223372036854775807}
        _ = v6
    case "foo bar":
        if (!false) {
            v7 := GenS1{F0: (4294967295 > 1), F1: -554, F2: -1}
            var v8 bool = (false && true)
            for fi9 := 0; fi9 < 7; fi9++ {
                var v10 uint64 = 0
                for fi11 := 0; fi11 < 2; fi11++ {
                    v12 := GenS2{F0: 0, F1: 0}
                    v4 += v1
                    switch 32768 {
                    case 125:
                        v2 = v2
                    default:
                        v13 := []int{fi11, -549, fi11, 2147483647}
                        v4 = (v1 & (^(v1)))
                        v14 := GenS3{F0: v8, F1: true}
                        var v15 uint32 = 42
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    v16 := []int{fi11, (fi9 & fi11), 32768}
                    _ = fi11
                    _ = v12
                    _ = v16
                }
                v17 := GenS4{F0: fi9, F1: 42, F2: 1000}
                _ = fi9
                _ = v10
                _ = v17
            }
            _ = v7
            _ = v8
        }
    case "hello":
        fmt.Println(v3, v2, v1)
    default:
        v3 += v3
    }
    if (1.5 <= -2.25) {
        v18 := GenS5{F0: true, F1: 2147483647, F2: (-(v3))}
        v19 := []int{734, 0, 42}
        _ = v18
        _ = v19
    }
    v20 := GenS6{F0: (((v2 + ("a" + "foo bar")) + v2) + "\n\t"), F1: -867}
    var v21 uint8 = 1
    if false {
        var v22 int32 = -891
        var v23 int16 = -1
        fmt.Println(v1)
        _ = v22
        _ = v23
    }
    v2 = "a"
    fmt.Println(v2, v5)
    switch -9223372036854775808 {
    case 127:
        v24 := GenS1{F0: false, F1: 42, F2: 75}
        _ = v24
    case 392:
        var v25 uint16 = 338
        _ = v25
    }
    fmt.Println(v1, v1)
    v21 -= (v21 / ((148) | 1))
    v26 := []int{int(v1), -546}
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    for fi27 := 0; fi27 < 10; fi27++ {
        var v28 bool = false
        v28 = v28
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
