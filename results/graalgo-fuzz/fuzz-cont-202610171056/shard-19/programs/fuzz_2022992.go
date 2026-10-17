package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint16 }
func genFn0(p0 uint, p1 int32) int {
    for fi4 := 0; fi4 < 18; fi4++ {
        fmt.Println(p0, p0, p1)
        p1 += 2147483647
        p1 -= (p1 | p1)
        _ = fi4
    }
    var v5 uint32 = 42
    v6 := GenS1{F0: 42, F1: 42, F2: 308}
    v5 += (v5 / ((v5) | 1))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 523
}

func main() {
    var v1 uint = 42
    var v2 bool = false
    var v3 uint16 = 65535
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 uint8 = 0
            switch ("hello" + (("hello" + ("\n\t" + "")) + (("foo bar" + "a") + ("a" + "")))) {
            case "":
                switch (fi8 & (^(fi8))) {
                case -741:
                    if (-473 < -32768) {
                        v10 := []int{fi7}
                        switch -9223372036854775808 {
                        case -1:
                            var v11 bool = v2
                            _ = v11
                        default:
                            v9 = 127
                        }
                        _ = v10
                    }
                }
            case "hello":
                v12 := GenS1{F0: -32768, F1: v3, F2: v3}
                _ = v12
            case "a":
                var v13 uint64 = 1
                var v14 int8 = 37
                v15 := GenS1{F0: -780, F1: v3, F2: v3}
                _ = v13
                _ = v14
                _ = v15
            }
            _ = fi8
            _ = v9
        }
        var v16 int32 = 2147483647
        if (v2 || v2) {
            fmt.Println(fi7, v1, fi7)
            v17 := GenS1{F0: 144, F1: 65535, F2: v3}
            v3 -= (^(v3))
            _ = v17
        }
        _ = fi7
        _ = v16
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        v19 := GenS1{F0: 835, F1: 1, F2: v3}
        v3 -= 32767
        _ = fi18
        _ = v19
    }
    var v20 uint8 = 127
    v1 = (v1 | (v1 & v1))
    v2 = v2
    switch "foo bar" {
    case "\n\t":
        v20 = (v20 & (v20 ^ (v20 / (((^(v20))) | 1))))
    }
    switch ("a" + "") {
    case "\n\t":
        v2 = ("a" > "")
    case "a":
        for fi21 := 0; fi21 < 11; fi21++ {
            v22 := []int{}
            fmt.Println(v20, v2)
            _ = fi21
            _ = v22
        }
    }
    var v23 int8 = 74
    var v24 float64 = 0.1
    if v2 {
        fmt.Println(v3, v20)
        if v2 {
            v1 += (v1 | v1)
            for fi25 := 0; fi25 < 11; fi25++ {
                switch fi25 {
                case 127:
                    v23 += v23
                case 2147483647:
                    v20 = (v20 + (v20 * (^(v20))))
                case 9223372036854775807:
                    var v26 uint = (v1 - (v1 << ((4294967295) & 63)))
                    _ = v26
                default:
                    v27 := []int{-1, (fi25 % ((fi25) | 1)), genFn0((^(v1)), 1)}
                    _ = v27
                }
                var v28 uint = v1
                fmt.Println(v24, v3, v20)
                _ = fi25
                _ = v28
            }
        } else {
            for fi29 := 0; fi29 < 9; fi29++ {
                fmt.Println(v1)
                var v30 int16 = 323
                v31 := GenS1{F0: v30, F1: v3, F2: (v3 - 42)}
                _ = fi29
                _ = v30
                _ = v31
            }
            var v32 int8 = v23
            var v33 bool = ((v20 * uint8(v3)) >= v20)
            _ = v32
            _ = v33
        }
    } else {
        v3 = v3
        v34 := []int{}
        _ = v34
    }
    var v35 int8 = v23
    var v36 uint32 = uint32(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v35
    _ = v36
}
