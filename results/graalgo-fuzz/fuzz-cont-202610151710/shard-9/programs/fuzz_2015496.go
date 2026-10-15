package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 uint8; F2 uint }
type GenS4 struct { F0 string; F1 uint16 }
func genFn0(p0 uint) uint64 {
    var v5 int8 = 0
    v6 := GenS1{F0: "\n\t", F1: 378}
    var v7 float32 = -2.25
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 559
}

func main() {
    var v1 uint32 = 2
    var v2 string = "\n\t"
    var v3 int32 = 309
    var v4 int = -1
    var v8 uint32 = 664
    if (1000 > 1000) {
        v9 := GenS2{F0: 42}
        _ = v9
    } else {
        v10 := []int{}
        v2 = v2
        _ = v10
    }
    v11 := GenS1{F0: v2, F1: 765}
    var v12 int16 = 32767
    for fi13 := 0; fi13 < 13; fi13++ {
        switch (v2 + "hello") {
        case "\n\t":
            v14 := GenS1{F0: (v2 + (v2 + "\n\t")), F1: 32767}
            _ = v14
        case "a":
            switch v2 {
            case "\n\t":
                var v15 float64 = 1e10
                _ = v15
            case "hello":
                var v16 bool = false
                _ = v16
            case "":
                switch 1 {
                case -952:
                    fmt.Println(v12)
                }
            default:
                for fi17 := 0; fi17 < 5; fi17++ {
                    v18 := []int{(fi13 & v4), 42, v4, v4}
                    switch "a" {
                    case "foo bar":
                        v19 := GenS3{F0: -32768, F1: 211, F2: 1}
                        _ = v19
                    case "\n\t":
                        v12 = (v12 % (((v12 + v12)) | 1))
                    default:
                        var v20 int64 = 490
                        _ = v20
                    }
                    _ = fi17
                    _ = v18
                }
            }
        case "":
            switch fi13 {
            case -1:
                v21 := GenS4{F0: ("" + v2), F1: 1}
                _ = v21
            case 32767:
                v22 := GenS3{F0: 989, F1: 128, F2: 42}
                _ = v22
            case 127:
                for fi23 := 0; fi23 < 12; fi23++ {
                    switch 32768 {
                    case 4294967295:
                        fmt.Println(v3, v1)
                        var v24 bool = true
                        _ = v24
                    }
                    switch (v2 + v2) {
                    case "\n\t":
                        v2 = v2
                    }
                    _ = fi23
                }
            }
        }
        var v25 int32 = -1
        var v26 uint = 447
        _ = fi13
        _ = v25
        _ = v26
    }
    if true {
        var v27 int64 = 0
        _ = v27
    } else {
        v1 += 0
        switch (v4 | -697) {
        case -854:
            v28 := []int{v4, -798, v4, (v4 - (v4 | -304))}
            _ = v28
        }
        switch -1 {
        case -2147483648:
            v3 = (-(v3))
        case -190:
            v4 = v4
        }
    }
    var v29 uint = 1
    var v30 int8 = -3
    var v31 uint = (v29 >> ((4294967295) & 63))
    switch "hello" {
    case "":
        for fi32 := 0; fi32 < 7; fi32++ {
            v3 += 42
            _ = fi32
        }
    case "a":
        v12 = (v12 * -32768)
    case "hello":
        v8 -= (v1 * (v1 / (((v8 >> (((v31 & 255)) & 63))) | 1)))
    default:
        switch (v4 & v4) {
        case 0:
            var v33 int8 = 1
            _ = v33
        case 127:
            for fi34 := 0; fi34 < 19; fi34++ {
                var v35 uint = (v29 ^ (v29 * v31))
                var v36 int8 = v30
                _ = fi34
                _ = v35
                _ = v36
            }
        }
    }
    if true {
        var v37 uint64 = 201
        var v38 uint32 = (v1 / ((2147483647) | 1))
        v30 -= (v30 >> ((297) & 63))
        _ = v37
        _ = v38
    }
    var v39 int8 = (^(v30))
    if true {
        v29 = v29
        var v40 int = 9223372036854775807
        v41 := GenS2{F0: 2}
        _ = v40
        _ = v41
    } else {
        fmt.Println(v2)
        v42 := []int{v4, v4}
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v39
    _ = v4
    _ = v8
}
