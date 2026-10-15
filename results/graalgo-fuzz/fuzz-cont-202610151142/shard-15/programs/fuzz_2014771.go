package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32 }
func genFn0(p0 uint16, p1 int64) int64 {
    switch "\n\t" {
    case "foo bar":
        v4 := GenS1{F0: true}
        _ = v4
    case "a":
        p1 = (^(p1))
    case "":
        switch "a" {
        case "a":
            switch int(p1) {
            case -626:
                var v5 int64 = (p1 % ((p1) | 1))
                _ = v5
            }
        case "\n\t":
            v6 := []int{32767, -2147483648, 127, -515}
            _ = v6
        default:
            for fi7 := 0; fi7 < 20; fi7++ {
                v8 := GenS1{F0: false}
                _ = fi7
                _ = v8
            }
        }
    default:
        p1 += p1
    }
    p0 = p0
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 int8, p1 bool, p2 uint) int16 {
    for fi9 := 0; fi9 < 1; fi9++ {
        switch ("" + ("foo bar" + (("" + (("foo bar" + "") + ("foo bar" + "\n\t"))) + ""))) {
        case "":
            for fi10 := 0; fi10 < 14; fi10++ {
                var v11 int = 42
                v12 := GenS1{F0: p1}
                _ = fi10
                _ = v11
                _ = v12
            }
        default:
            var v13 int32 = 2147483647
            _ = v13
        }
        var v14 uint8 = 255
        _ = fi9
        _ = v14
    }
    p1 = false
    _ = p0
    _ = p1
    _ = p2
    return 1000
}

func main() {
    var v1 string = "a"
    var v2 uint32 = 42
    var v3 uint8 = 127
    v1 = v1
    fmt.Println(v2)
    v15 := GenS2{F0: 0}
    if false {
        v2 = (^(v2))
        v16 := []int{}
        v17 := GenS3{F0: 2, F1: 61}
        _ = v16
        _ = v17
    } else {
        var v18 int16 = genFn1(-30, ((uint8(v2) >= uint8(v2)) && (true || (127 == -1))), 65535)
        var v19 int32 = -1
        _ = v18
        _ = v19
    }
    if true {
        var v20 int64 = 700
        var v21 float64 = 0.0
        _ = v20
        _ = v21
    }
    v3 += 2
    v3 = (v3 << ((uint(v3)) & 63))
    for fi22 := 0; fi22 < 3; fi22++ {
        v23 := GenS4{F0: -339}
        v24 := []int{(fi22 - fi22), 127, (-(fi22)), (fi22 << ((224) & 63))}
        v25 := GenS2{F0: 4294967295}
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        v27 := GenS5{F0: -2147483648}
        _ = fi26
        _ = v27
    }
    v3 = (v3 - v3)
    for fi28 := 0; fi28 < 11; fi28++ {
        v3 += (v3 / (((v3 | 42)) | 1))
        v1 = (v1 + v1)
        var v29 string = "foo bar"
        _ = fi28
        _ = v29
    }
    var v30 string = v1
    var v31 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
}
