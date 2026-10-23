package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int16 }
type GenS2 struct { F0 int; F1 uint; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint16; F2 int32 }
func genFn0(p0 int) int8 {
    var v5 int8 = -35
    fmt.Println(p0)
    if false {
        var v6 int8 = 1
        v5 = v6
        switch int(v5) {
        case -476:
            fmt.Println(v5)
        default:
            switch "hello" {
            case "a":
                v7 := []int{(p0 & (^(p0))), (p0 % ((-1) | 1)), (p0 % (((^(p0))) | 1))}
                _ = v7
            case "":
                fmt.Println(p0, v5)
            case "hello":
                v8 := GenS1{F0: "hello", F1: 4294967295, F2: -32768}
                _ = v8
            }
        }
        _ = v6
    } else {
        var v9 int8 = int8(p0)
        for fi10 := 0; fi10 < 18; fi10++ {
            p0 = (p0 ^ fi10)
            v11 := []int{}
            var v12 uint32 = 249
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = v9
    }
    v13 := GenS1{F0: (((("\n\t" + "\n\t") + ("a" + "\n\t")) + "\n\t") + ""), F1: 4294967295, F2: 42}
    _ = p0
    _ = v13
    _ = v5
    return v5
}
func genFn1(p0 uint, p1 int64) uint64 {
    if (false || true) {
        var v14 bool = (false || (1 > 100))
        if ((42 > 0) || (v14 || v14)) {
            fmt.Println(p1, v14, p0)
            if (0 > -523) {
                var v15 int64 = -268
                var v16 string = "a"
                _ = v15
                _ = v16
            }
        } else {
            v17 := GenS2{F0: 962, F1: p0, F2: 664}
            _ = v17
        }
        _ = v14
    }
    var v18 float32 = 1e10
    switch 4294967295 {
    case 9223372036854775807:
        v19 := []int{-947, 32767, 42, 32768}
        _ = v19
    case 32767:
        v20 := GenS3{F0: -1, F1: 0, F2: -122}
        _ = v20
    default:
        fmt.Println(v18, p0, p1)
    }
    _ = p0
    _ = p1
    _ = v18
    return 1
}

func main() {
    var v1 float64 = 0.1
    var v2 int64 = -1
    var v3 bool = true
    var v4 uint8 = 127
    for fi21 := 0; fi21 < 6; fi21++ {
        v22 := []int{}
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
        var v23 float32 = 0.1
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS1{F0: "hello", F1: 294, F2: 1000}
    switch -9223372036854775808 {
    case 791:
        if true {
            v3 = false
            var v25 float32 = 1e10
            _ = v25
        }
    case 128:
        var v26 uint = 223
        _ = v26
    case 2147483647:
        var v27 int32 = 851
        _ = v27
    default:
        for fi28 := 0; fi28 < 16; fi28++ {
            v29 := []int{(-(fi28)), fi28, 4294967295, fi28}
            _ = fi28
            _ = v29
        }
    }
    fmt.Println(v2, v3, v3)
    if ((^(v2)) <= v2) {
        for fi30 := 0; fi30 < 2; fi30++ {
            var v31 uint = 65535
            v32 := GenS1{F0: ("hello" + "a"), F1: 90, F2: 1}
            var v33 float32 = 1.5
            _ = fi30
            _ = v31
            _ = v32
            _ = v33
        }
    } else {
        var v34 string = "a"
        fmt.Println(v34)
        _ = v34
    }
    if false {
        v3 = v3
        if (!(1.5 < 1.5)) {
            v35 := GenS3{F0: 1, F1: 2, F2: 42}
            _ = v35
        }
    } else {
        var v36 uint16 = 1
        v2 = v2
        _ = v36
    }
    v3 = true
    v1 -= 1e10
    for fi37 := 0; fi37 < 17; fi37++ {
        v4 = 171
        v38 := []int{}
        var v39 uint = 950
        _ = fi37
        _ = v38
        _ = v39
    }
    var v40 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v40
}
