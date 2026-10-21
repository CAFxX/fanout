package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint16 }
type GenS2 struct { F0 int8; F1 string; F2 uint16 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint32) uint32 {
    p0 += 0
    v7 := []int{}
    _ = len(v7)
    _ = p0
    _ = v7
    return (^(p0))
}
func genFn1(p0 uint64, p1 int32) uint {
    var v8 uint = 434
    switch 32767 {
    case 128:
        fmt.Println(v8, p1, p0)
    case 713:
        var v9 uint32 = 2
        _ = v9
    case 0:
        for fi10 := 0; fi10 < 1; fi10++ {
            v11 := []int{(fi10 & (fi10 / ((484) | 1))), fi10, (fi10 + 312)}
            v8 -= (v8 % ((742) | 1))
            if true {
                if false {
                    if true {
                        v12 := GenS1{F0: 0, F1: (752 >= 0), F2: 255}
                        _ = v12
                    }
                }
            } else {
                p0 = 1
            }
            _ = fi10
            _ = v11
        }
    }
    v13 := GenS1{F0: p1, F1: false, F2: 255}
    var v14 int64 = -921
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v8
    return v8
}

func main() {
    var v1 uint8 = 2
    var v2 int64 = -297
    var v3 uint = 1
    v4 := []string{}
    var v5 uint = 65535
    var v6 int8 = 42
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 uint = v5
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: 42, F1: "\n\t", F2: uint16(v1)}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) > 0 {
        v4[0] = ""
    }
    if true {
        _ = len(v4)
        if false {
            if len(v4) > 0 {
                v4[0] = "hello"
            }
            if false {
                _ = len(v4)
            }
            var v18 uint16 = 0
            _ = v18
        } else {
            switch ("a" + "") {
            case "":
                v19 := GenS2{F0: (^(v6)), F1: ("foo bar" + ("\n\t" + "hello")), F2: 255}
                _ = v19
            case "foo bar":
                var v20 bool = true
                _ = v20
            case "a":
                for fi21 := 0; fi21 < 17; fi21++ {
                    if (v5 == (v3 & v5)) {
                        if ((false && true) && (false || ((v2 >> ((161) & 63)) == v2))) {
                            v3 = v5
                        }
                        var v22 string = "hello"
                        if (fi21 <= (fi21 ^ fi21)) {
                            v3 = (v5 | (v3 << ((genFn1(18446744073709551615, 0)) & 63)))
                        }
                        _ = v22
                    }
                    switch (fi21 ^ fi21) {
                    case 0:
                        v23 := GenS3{F0: (v1 + 64), F1: (v1 / ((2) | 1))}
                        _ = v23
                    case 769:
                        if (!true) {
                            v24 := GenS4{F0: v3}
                            v5 = v3
                            _ = v24
                        }
                    default:
                        v25 := GenS4{F0: genFn1(869, -1)}
                        _ = v25
                    }
                    var v26 uint64 = 42
                    _ = fi21
                    _ = v26
                }
            }
            v27 := GenS3{F0: 127, F1: v1}
            _ = v27
        }
    }
    var v28 string = "a"
    for fi29 := 0; fi29 < 17; fi29++ {
        for fi30 := 0; fi30 < 9; fi30++ {
            v3 -= (v5 << ((v5) & 63))
            v1 = (v1 % ((v1) | 1))
            _ = fi30
        }
        var v31 int = (-(fi29))
        _ = fi29
        _ = v31
    }
    fmt.Println(v1, v28, v2)
    switch -1 {
    case -2147483648:
        _ = len(v4)
    }
    v1 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
