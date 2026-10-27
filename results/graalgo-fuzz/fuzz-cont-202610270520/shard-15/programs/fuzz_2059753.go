package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int; F2 uint16 }

func main() {
    var v1 uint = 255
    var v2 uint8 = 2
    var v3 uint32 = 0
    if (("" + ("\n\t" + ("" + ("foo bar" + "\n\t")))) > "hello") {
        var v4 int16 = 42
        if true {
            v3 += 42
            v2 += v2
        } else {
            if ((2 <= 264) && (true && false)) {
                v5 := []int{-160, int(v1), 0}
                _ = v5
            } else {
                v6 := []int{0, 2147483647, -47}
                fmt.Println(v1, v3)
                var v7 uint = v1
                _ = v6
                _ = v7
            }
            if false {
                v8 := GenS1{F0: (^(v1))}
                v2 -= v2
                fmt.Println(v3, v2)
                _ = v8
            } else {
                v1 = (v1 ^ v1)
                v9 := []int{-535, -422}
                _ = v9
            }
            var v10 uint16 = 306
            _ = v10
        }
        _ = v4
    }
    v11 := []int{451, -56, 1}
    if len(v11) > 0 {
        v11[0] = -858
    }
    _ = len(v11)
    var v12 int8 = 1
    if (!true) {
        v1 -= 65535
    }
    v13 := GenS2{F0: 1, F1: -605, F2: 2}
    var v14 uint32 = (v3 | (v3 | (v3 & v3)))
    v15 := GenS2{F0: 32767, F1: 4294967295, F2: 65535}
    if (!(182 >= 1)) {
        var v16 uint64 = 837
        if true {
            var v17 bool = (false || true)
            _ = v17
        } else {
            v16 = (v16 * (v16 + 543))
            fmt.Println(v14, v1, v1)
            for fi18 := 0; fi18 < 4; fi18++ {
                for fi19 := 0; fi19 < 18; fi19++ {
                    v11 = append(v11, (fi19 ^ 341))
                    v20 := GenS1{F0: (v1 ^ v1)}
                    switch "foo bar" {
                    case "foo bar":
                        v16 = (^(v16))
                    case "\n\t":
                        v3 = v14
                        v3 -= v14
                    default:
                        fmt.Println(fi19, fi19)
                    }
                    _ = fi19
                    _ = v20
                }
                for fi21 := 0; fi21 < 18; fi21++ {
                    v14 = 69
                    switch "a" {
                    case "\n\t":
                        fmt.Println(v3, v2)
                    case "a":
                        fmt.Println(v2, v16, v14)
                    }
                    var v22 uint = (v1 + 291)
                    _ = fi21
                    _ = v22
                }
                var v23 uint8 = 1
                _ = fi18
                _ = v23
            }
        }
        v1 = v1
        _ = v16
    }
    switch 4294967295 {
    case 32768:
        var v24 int8 = v12
        _ = v24
    }
    v2 = v2
    var v25 int64 = 42
    v2 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v3
}
