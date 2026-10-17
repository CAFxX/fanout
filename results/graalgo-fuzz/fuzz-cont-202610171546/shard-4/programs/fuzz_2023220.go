package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int; F1 int8; F2 int8 }
type GenS5 struct { F0 int64 }
func genFn0(p0 int64, p1 int64) uint16 {
    v7 := GenS1{F0: true, F1: 780}
    p0 -= -1
    v8 := GenS1{F0: false, F1: 1}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 255
}

func main() {
    v1 := []string{"hello"}
    var v2 int16 = 1
    var v3 uint64 = 0
    var v4 float64 = ((-2.25 + -2.25) + 3.14159)
    var v5 uint64 = (v3 % ((979) | 1))
    var v6 uint64 = v5
    v2 -= 1000
    switch ("a" + "") {
    case "hello":
        v9 := GenS1{F0: (!false), F1: 443}
        _ = v9
    case "\n\t":
        for fi10 := 0; fi10 < 18; fi10++ {
            v1 = append(v1, "hello")
            _ = fi10
        }
    default:
        v11 := GenS2{F0: 1, F1: "foo bar"}
        _ = v11
    }
    var v12 uint32 = 2
    for fi13 := 0; fi13 < 2; fi13++ {
        if true {
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 float64 = ((1e10 + 0.0) / float64(1e10))
                _ = fi14
                _ = v15
            }
            v2 += 0
        } else {
            fmt.Println(fi13, v6, v6)
        }
        _ = fi13
    }
    fmt.Println(v4, v3, v12)
    for fi16 := 0; fi16 < 8; fi16++ {
        v5 = v3
        v6 = 18446744073709551615
        if false {
            fmt.Println(fi16)
            for fi17 := 0; fi17 < 20; fi17++ {
                if true {
                    var v18 float32 = (float32(v4) - ((((3.14159 - (1e10 - float32(v4))) * -2.25) + 1.5) / float32((v4 * v4))))
                    var v19 uint8 = 2
                    var v20 string = "foo bar"
                    _ = v18
                    _ = v19
                    _ = v20
                } else {
                    switch (fi17 + 127) {
                    case -2147483648:
                        var v21 uint8 = 127
                        _ = v21
                    case 796:
                        var v22 uint8 = 2
                        _ = v22
                    }
                    v23 := GenS3{F0: 65535, F1: uint8(v3), F2: v6}
                    v1 = append(v1, "a")
                    fmt.Println(v5, fi17, v2)
                    var v24 uint16 = 42
                    var v25 float32 = ((1e10 * ((1e10 / (1e10 - float32(v4))) * -2.25)) * (-2.25 + 1.5))
                    _ = v23
                    _ = v24
                    _ = v25
                }
                _ = fi17
            }
            for fi26 := 0; fi26 < 9; fi26++ {
                if true {
                    switch (fi16 % (((-(fi16))) | 1)) {
                    case 0:
                        v5 = (v5 + 18446744073709551615)
                    }
                    if false {
                        v2 = v2
                    }
                    v4 = v4
                }
                v27 := GenS4{F0: fi16, F1: 1, F2: 0}
                var v28 bool = (!(!false))
                v1 = append(v1, "")
                if ((3.14159 / v4) < (((v4 + -2.25) - 0.0) / ((float64(0.0) / v4) * 1.5))) {
                    v1 = append(v1, "hello")
                    var v29 int = (fi26 >> ((62) & 63))
                    var v30 int = (fi26 % ((42) | 1))
                    _ = v29
                    _ = v30
                } else {
                    switch "a" {
                    case "a":
                        var v31 uint32 = v12
                        _ = v31
                    default:
                        fmt.Println(v2, v3)
                    }
                }
                if false {
                    fmt.Println(fi16, v2, v3)
                    v5 += (v6 / (((v5 ^ (v3 << ((191) & 63)))) | 1))
                }
                var v32 int64 = 0
                fmt.Println(v3)
                _ = fi26
                _ = v27
                _ = v28
                _ = v32
            }
        } else {
            v33 := GenS5{F0: -288}
            _ = v33
        }
        _ = fi16
    }
    if true {
        var v34 int = -385
        _ = v34
    }
    v12 -= v12
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
