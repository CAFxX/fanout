package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint32; F1 uint16; F2 string }
func genFn0(p0 uint64, p1 int64, p2 uint8) uint8 {
    if true {
        for fi5 := 0; fi5 < 13; fi5++ {
            if ((32767 <= 1000) && true) {
                p2 = (p2 + p2)
                v6 := GenS1{F0: uint8(p0)}
                _ = v6
            } else {
                for fi7 := 0; fi7 < 17; fi7++ {
                    var v8 float32 = -2.25
                    _ = fi7
                    _ = v8
                }
            }
            if (!true) {
                p0 += uint64(p2)
            } else {
                v9 := []int{}
                for fi10 := 0; fi10 < 8; fi10++ {
                    switch (fi10 ^ (fi5 | fi5)) {
                    case 42:
                        v11 := GenS1{F0: 2}
                        _ = v11
                    case 2147483647:
                        if ((p0 >= (^(p0))) || true) {
                            v12 := GenS1{F0: 0}
                            var v13 bool = (-1 == 42)
                            v14 := GenS1{F0: (p2 / ((p2) | 1))}
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                        var v15 uint32 = 680
                        _ = v15
                    case 329:
                        p2 = p2
                    default:
                        var v16 uint64 = 1
                        _ = v16
                    }
                    if len(v9) > 0 {
                        v9[0] = (fi5 ^ (^(fi5)))
                    }
                    p0 += (p0 - p0)
                    _ = fi10
                }
                fmt.Println(p0, p2, p1)
                _ = v9
            }
            v17 := []int{(fi5 - fi5), -88, fi5}
            _ = fi5
            _ = v17
        }
    } else {
        switch -2147483648 {
        case 1:
            p1 = (p1 / ((42) | 1))
        case -102:
            var v18 string = "hello"
            _ = v18
        case 206:
            var v19 int32 = -179
            _ = v19
        default:
            p0 = 351
        }
        if false {
            v20 := GenS1{F0: p2}
            if true {
                p0 += p0
                for fi21 := 0; fi21 < 9; fi21++ {
                    p0 = (^(p0))
                    var v22 int16 = 0
                    _ = fi21
                    _ = v22
                }
            }
            _ = v20
        } else {
            v23 := GenS1{F0: (p2 % ((p2) | 1))}
            var v24 int64 = -9223372036854775808
            _ = v23
            _ = v24
        }
        var v25 int32 = 103
        _ = v25
    }
    p0 -= 661
    _ = p0
    _ = p1
    _ = p2
    return (p2 | (p2 & p2))
}

func main() {
    v1 := []string{("" + ""), "\n\t", ("\n\t" + ("foo bar" + "foo bar"))}
    var v2 uint = 65535
    var v3 int32 = 883
    v4 := []uint32{2147483648, 0, 565, 42}
    fmt.Println(v3, v3)
    var v26 int64 = -1
    fmt.Println(v2, v2)
    for fi27 := 0; fi27 < 16; fi27++ {
        v28 := GenS2{F0: (fi27 >> (((v2 / ((v2) | 1))) & 63)), F1: v3, F2: 0}
        var v29 int8 = 42
        _ = fi27
        _ = v28
        _ = v29
    }
    switch -951 {
    case 1:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 396:
        for fi30 := 0; fi30 < 12; fi30++ {
            var v31 uint64 = 710
            _ = fi30
            _ = v31
        }
    case 9223372036854775807:
        v26 = v26
    }
    var v32 uint16 = 65535
    v3 = v3
    v33 := GenS3{F0: 65535, F1: 0, F2: ""}
    switch "\n\t" {
    case "hello":
        v3 = v3
    case "a":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "foo bar":
        var v34 int8 = 127
        _ = v34
    }
    var v35 int64 = (v26 << ((v2) & 63))
    var v36 uint16 = v32
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v35
    _ = v36
    _ = v4
}
