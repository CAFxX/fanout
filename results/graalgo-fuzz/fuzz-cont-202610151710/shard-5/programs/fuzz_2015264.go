package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint16, p1 int, p2 int16) uint8 {
    for fi4 := 0; fi4 < 5; fi4++ {
        fmt.Println(p2, p2, fi4)
        _ = fi4
    }
    if false {
        if true {
            if (!false) {
                v5 := []int{294, p1}
                for fi6 := 0; fi6 < 5; fi6++ {
                    v7 := GenS1{F0: -32768}
                    for fi8 := 0; fi8 < 6; fi8++ {
                        v9 := GenS1{F0: 1000}
                        switch (("a" + "\n\t") + "\n\t") {
                        case "\n\t":
                            if ((p0 >> ((203) & 63)) < (p0 * 32767)) {
                                fmt.Println(p2)
                                var v10 int64 = -9223372036854775808
                                _ = v10
                            }
                        case "":
                            fmt.Println(p2, fi6)
                        }
                        fmt.Println(fi6, p0)
                        _ = fi8
                        _ = v9
                    }
                    _ = fi6
                    _ = v7
                }
                _ = v5
            } else {
                var v11 uint16 = (p0 / (((p0 ^ p0)) | 1))
                _ = v11
            }
            var v12 uint64 = 374
            _ = v12
        }
        var v13 int = -89
        switch "hello" {
        case "":
            v14 := []int{p1, p1, v13, (v13 << ((255) & 63))}
            _ = v14
        case "\n\t":
            var v15 bool = (float64(3.14159) >= -2.25)
            _ = v15
        default:
            for fi16 := 0; fi16 < 18; fi16++ {
                v17 := []int{p1, p1, v13}
                _ = fi16
                _ = v17
            }
        }
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return 206
}
func genFn1(p0 int8, p1 uint64) int32 {
    p1 -= (^(p1))
    v18 := []int{743}
    switch -2147483648 {
    case -2147483648:
        fmt.Println(p1, p1)
    default:
        if (!true) {
            var v19 int64 = 42
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            _ = v19
        } else {
            for fi20 := 0; fi20 < 16; fi20++ {
                p0 += p0
                _ = fi20
            }
            for fi21 := 0; fi21 < 19; fi21++ {
                fmt.Println(fi21, p1)
                p0 -= 0
                _ = fi21
            }
        }
    }
    var v22 uint = 663
    _ = p0
    _ = p1
    _ = v18
    _ = v22
    return 0
}

func main() {
    var v1 uint16 = 65535
    var v2 int16 = 32767
    var v3 uint64 = 556
    if (v3 < 0) {
        v3 -= (v3 * (v3 << ((35) & 63)))
        var v23 int16 = (v2 + (v2 << ((334) & 63)))
        v3 -= (v3 & (v3 ^ v3))
        _ = v23
    } else {
        switch ("hello" + "") {
        case "\n\t":
            if true {
                var v24 bool = true
                var v25 string = ("" + ("" + ("a" + "a")))
                _ = v24
                _ = v25
            } else {
                v26 := GenS2{F0: v1}
                _ = v26
            }
        case "hello":
            v3 = (v3 * 2)
        default:
            var v27 uint = 111
            _ = v27
        }
        fmt.Println(v2, v2)
    }
    v28 := GenS2{F0: v1}
    fmt.Println(v2, v3)
    v1 += (v1 * (v1 >> ((1) & 63)))
    if (true && false) {
        v29 := []int{2147483647, -954, 9223372036854775807, -2147483648}
        switch ("a" + "") {
        case "hello":
            var v30 uint8 = genFn0(v1, 32768, (v2 ^ -599))
            _ = v30
        default:
            v31 := GenS2{F0: v1}
            _ = v31
        }
        _ = v29
    }
    v32 := GenS1{F0: 1}
    var v33 bool = (4294967295 > 2147483647)
    var v34 int64 = 1
    v33 = v33
    var v35 int32 = genFn1(0, (v3 << ((4294967295) & 63)))
    var v36 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
}
