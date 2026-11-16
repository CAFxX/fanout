package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint64; F2 bool }
type GenS3 struct { F0 int32; F1 uint32; F2 bool }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 int; F1 int }

func main() {
    var v1 float64 = (3.14159 / -2.25)
    v2 := []string{}
    v3 := []int{-1}
    var v4 int32 = 567
    v5 := make(map[int]int)
    var v6 bool = (!(true || (32767 > 65535)))
    var v7 uint64 = 18446744073709551615
    fmt.Println(v7, v1, v4)
    var v8 int8 = -1
    for fi9 := 0; fi9 < 4; fi9++ {
        for fi10 := 0; fi10 < 20; fi10++ {
            fmt.Println(v6, v4, v1)
            _ = fi10
        }
        _ = fi9
    }
    if v6 {
        var v11 string = ""
        switch -2147483648 {
        case 4294967295:
            for fi12 := 0; fi12 < 9; fi12++ {
                _ = len(v2)
                _ = fi12
            }
        case 1:
            if true {
                for fi13 := 0; fi13 < 12; fi13++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    fmt.Println(v7, v6, v7)
                    _ = fi13
                }
                var v14 string = v11
                fmt.Println(v6, v4, v8)
                _ = v14
            } else {
                v15 := GenS1{F0: "hello"}
                _ = len(v5)
                _ = v15
            }
        case 2147483647:
            var v16 int64 = -766
            _ = v16
        }
        var v17 bool = v6
        _ = v11
        _ = v17
    }
    if v6 {
        for fi18 := 0; fi18 < 7; fi18++ {
            for fi19 := 0; fi19 < 8; fi19++ {
                fmt.Println(v7, fi18)
                var v20 string = "a"
                v21 := GenS1{F0: v20}
                _ = fi19
                _ = v20
                _ = v21
            }
            var v22 uint8 = 1
            _ = fi18
            _ = v22
        }
        var v23 int64 = 1
        v24 := GenS1{F0: "\n\t"}
        _ = v23
        _ = v24
    } else {
        var v25 int32 = 205
        v5[42] = 32767
        if (((1.5 != (0.0 * (3.14159 - 3.14159))) || (!(360 == 130))) && (!v6)) {
            v26 := GenS2{F0: (v8 + 127), F1: v7, F2: false}
            _ = v26
        } else {
            v1 = -2.25
            var v27 uint16 = 32767
            _ = v27
        }
        _ = v25
    }
    v28 := GenS3{F0: (v4 * (v4 + (v4 + -1))), F1: 42, F2: (((-(v8)) != (v8 / (((v8 - (v8 ^ v8))) | 1))) || v6)}
    v29 := GenS1{F0: ("a" + "")}
    if ((1000 < -1) || (9223372036854775807 < -872)) {
        v4 += 2147483647
        _ = len(v5)
    } else {
        v4 += v4
    }
    for fi30 := 0; fi30 < 16; fi30++ {
        for fi31 := 0; fi31 < 12; fi31++ {
            v7 = v7
            v32 := GenS4{F0: 1000, F1: -1}
            _ = fi31
            _ = v32
        }
        v1 = (1e10 - (v1 + v1))
        if (v6 && (15 != 42)) {
            v33 := GenS2{F0: v8, F1: v7, F2: true}
            _ = v33
        } else {
            switch int(v7) {
            case 901:
                v34 := GenS4{F0: -1, F1: 1}
                _ = v34
            }
            v5[fi30] = fi30
            v8 = (v8 % (((v8 << ((18446744073709551615) & 63))) | 1))
        }
        _ = fi30
    }
    if v6 {
        v35 := GenS5{F0: 4294967295, F1: 127}
        fmt.Println(v8, v1)
        _ = v35
    } else {
        var v36 uint64 = (v7 - v7)
        switch -1 {
        case -1:
            var v37 bool = ((v7 ^ (v7 - 18446744073709551615)) > v36)
            _ = v37
        default:
            fmt.Println(v8)
        }
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
