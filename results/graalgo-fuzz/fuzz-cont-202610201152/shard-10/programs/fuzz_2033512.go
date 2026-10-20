package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 float64 = (1e10 * (1e10 - 0.1))
    var v2 int32 = -225
    var v3 uint32 = 2
    var v4 int16 = -517
    var v5 uint32 = 673
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 string = "a"
        v8 := []int{(fi6 * 4294967295), fi6, -991, fi6}
        _ = len(v8)
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        v1 = v1
        v4 = 1000
        for fi10 := 0; fi10 < 13; fi10++ {
            if true {
                if ((uint64(v5) != 42) && true) {
                    var v11 string = "foo bar"
                    _ = v11
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        v12 := GenS1{F0: 2, F1: 4294967295, F2: -2147483648}
                        _ = v12
                    default:
                        if true {
                            if (fi9 > -386) {
                                var v13 bool = false
                                _ = v13
                            }
                            v14 := GenS2{F0: false}
                            _ = v14
                        }
                    }
                    v15 := GenS2{F0: (true && (-102 < 127))}
                    _ = v15
                }
                var v16 int = (fi9 + fi9)
                var v17 uint16 = 736
                _ = v16
                _ = v17
            }
            _ = fi10
        }
        _ = fi9
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 uint16 = 23
        if (2 >= 893) {
            v4 = v4
            v20 := GenS2{F0: true}
            _ = v20
        }
        for fi21 := 0; fi21 < 8; fi21++ {
            switch ("\n\t" + "foo bar") {
            case "hello":
                var v22 string = "a"
                _ = v22
            case "":
                var v23 uint16 = (v19 >> ((783) & 63))
                v3 -= 0
                fmt.Println(v2, v4)
                _ = v23
            default:
                var v24 int8 = 28
                _ = v24
            }
            var v25 bool = false
            _ = fi21
            _ = v25
        }
        _ = fi18
        _ = v19
    }
    v26 := []int{32767, 257}
    v2 = 2147483647
    if (!(!(!(false || true)))) {
        var v27 uint64 = 177
        _ = v27
    }
    for fi28 := 0; fi28 < 6; fi28++ {
        v5 = (v5 * 2147483647)
        _ = fi28
    }
    v2 = (v2 % (((v2 & (v2 % ((v2) | 1)))) | 1))
    switch 32768 {
    case 0:
        if len(v26) > 0 {
            v26[0] = 32768
        }
    default:
        var v29 uint = 952
        _ = v29
    }
    var v30 float64 = v1
    var v31 uint8 = 255
    switch 32768 {
    case -696:
        v32 := GenS3{F0: (^(v3))}
        _ = v32
    case -2147483648:
        v30 += (((((v30 * (-2.25 + 0.0)) / v1) + float64(3.14159)) / float64(3.14159)) - v1)
    case 483:
        v33 := GenS2{F0: true}
        _ = v33
    default:
        v3 = (v5 + 2)
    }
    var v34 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v5
}
