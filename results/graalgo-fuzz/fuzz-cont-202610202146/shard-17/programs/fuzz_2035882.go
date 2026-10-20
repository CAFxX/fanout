package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int64 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = -1
    var v3 string = ("\n\t" + "a")
    v4 := make(map[string]int)
    v5 := make(map[int]uint32)
    v6 := make(map[int]int)
    var v7 float64 = 0.1
    v8 := []int{14}
    var v9 uint32 = 42
    for fi10 := 0; fi10 < 7; fi10++ {
        switch (v3 + v3) {
        case "foo bar":
            var v11 uint64 = 18446744073709551615
            _ = v11
        default:
            var v12 uint64 = 537
            _ = v12
        }
        fmt.Println(fi10, fi10)
        if len(v8) > 0 {
            v8[0] = fi10
        }
        _ = fi10
    }
    v13 := GenS1{F0: 32767, F1: (v3 + v3), F2: 42}
    if true {
        v9 = 2147483647
        v14 := GenS1{F0: 789, F1: v3, F2: 9223372036854775807}
        _ = v14
    }
    var v15 uint64 = 2
    for fi16 := 0; fi16 < 10; fi16++ {
        switch 9223372036854775807 {
        case 128:
            for fi17 := 0; fi17 < 10; fi17++ {
                var v18 float64 = 1e10
                if true {
                    v19 := GenS1{F0: 0, F1: v3, F2: -1}
                    v20 := GenS2{F0: v15, F1: 598}
                    fmt.Println(v18, fi17)
                    _ = v19
                    _ = v20
                } else {
                    v21 := GenS2{F0: 42, F1: 0}
                    _ = v21
                }
                fmt.Println(v7, v7, v2)
                _ = fi17
                _ = v18
            }
        default:
            for fi22 := 0; fi22 < 11; fi22++ {
                v7 = (v7 / 0.1)
                var v23 uint64 = (^(v15))
                v8 = append(v8, (fi22 - fi16))
                _ = fi22
                _ = v23
            }
        }
        if true {
            switch (fi16 ^ 0) {
            case 128:
                var v24 bool = (true || (true && false))
                _ = v24
            case 4294967295:
                v25 := GenS1{F0: 65535, F1: v3, F2: 222}
                _ = v25
            case 216:
                fmt.Println(v3)
            }
            if true {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                for fi26 := 0; fi26 < 18; fi26++ {
                    var v27 uint = 65535
                    _ = fi26
                    _ = v27
                }
                for fi28 := 0; fi28 < 1; fi28++ {
                    v2 += (^(v2))
                    _ = fi28
                }
            } else {
                v15 = v15
                v29 := GenS3{F0: -2147483648}
                _ = v29
            }
        } else {
            v30 := GenS1{F0: 65535, F1: (v3 + "foo bar"), F2: 308}
            var v31 float64 = (0.0 - ((1.5 - v7) / v7))
            fmt.Println(v31)
            _ = v30
            _ = v31
        }
        fmt.Println(v15, v2, v15)
        _ = fi16
    }
    var v32 uint64 = 0
    var v33 uint8 = 42
    v32 += 42
    switch (v2 + 127) {
    case 0:
        if true {
            fmt.Println(v7)
        } else {
            v9 = (^(v9))
            if false {
                var v34 int8 = 39
                _ = v34
            }
        }
    case 32768:
        _ = len(v5)
    case -724:
        var v35 bool = (true && true)
        _ = v35
    }
    v36 := GenS4{F0: (!((^(v33)) >= (v33 - (v33 << ((255) & 63)))))}
    v7 += 0.1
    fmt.Println(v7, v2, v32)
    v9 = 618
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
