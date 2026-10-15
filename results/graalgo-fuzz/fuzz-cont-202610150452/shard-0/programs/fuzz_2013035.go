package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint64; F1 int16 }
func genFn0() uint64 {
    v4 := []int{1, 9223372036854775807, 32768, -1}
    v5 := GenS1{F0: -1, F1: 320}
    var v6 int64 = -1
    _ = v4
    _ = v5
    _ = v6
    return 301
}

func main() {
    var v1 uint = 0
    v2 := make(map[string]string)
    var v3 int64 = 1
    var v7 int8 = 42
    switch 832 {
    case -782:
        delete(v2, ("hello" + ("\n\t" + (("foo bar" + "\n\t") + "hello"))))
    case 128:
        var v8 uint16 = 0
        _ = v8
    }
    fmt.Println(v3, v3)
    var v9 int8 = -65
    if false {
        for fi10 := 0; fi10 < 9; fi10++ {
            for fi11 := 0; fi11 < 11; fi11++ {
                fmt.Println(v1, v9, v9)
                v3 -= (v3 + -9223372036854775808)
                _ = fi11
            }
            for fi12 := 0; fi12 < 14; fi12++ {
                switch ("hello" + ("foo bar" + "")) {
                case "":
                    v9 -= (v9 & 1)
                }
                v1 = 0
                if true {
                    v3 = v3
                    v13 := GenS2{F0: (v1 + (v1 % ((v1) | 1))), F1: -1}
                    v7 += v7
                    var v14 float64 = 1.5
                    _ = v13
                    _ = v14
                } else {
                    var v15 int = (-(fi12))
                    fmt.Println(fi10)
                    _ = v15
                }
                var v16 int16 = -469
                var v17 int32 = 1
                _ = fi12
                _ = v16
                _ = v17
            }
            _ = fi10
        }
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 int8 = (v7 * v7)
            v20 := []int{(fi18 - fi18)}
            var v21 float64 = (3.14159 - (-2.25 - 0.1))
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
    }
    v22 := []int{}
    fmt.Println(v9, v7)
    v9 = v9
    switch -574 {
    case -9223372036854775808:
        for fi23 := 0; fi23 < 14; fi23++ {
            v24 := GenS3{F0: 1, F1: 42}
            delete(v2, "hello")
            var v25 uint16 = 65535
            _ = fi23
            _ = v24
            _ = v25
        }
    case 128:
        for fi26 := 0; fi26 < 18; fi26++ {
            for fi27 := 0; fi27 < 13; fi27++ {
                var v28 int8 = v7
                v29 := GenS3{F0: 1, F1: -1}
                fmt.Println(fi26)
                _ = fi27
                _ = v28
                _ = v29
            }
            v30 := v2[("\n\t" + "hello")]
            fmt.Println(v7)
            _ = fi26
            _ = v30
        }
    }
    if (true && ((42 != 18446744073709551615) && false)) {
        var v31 uint64 = 2
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v7
    _ = v9
}
