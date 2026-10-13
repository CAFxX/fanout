package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 int; F2 uint64 }
type GenS4 struct { F0 int; F1 bool; F2 uint8 }
type GenS5 struct { F0 string }

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 65535
    var v3 int16 = 1
    var v4 uint8 = 0
    var v5 uint16 = 65535
    if (v2 > v2) {
        var v6 int16 = (^(v3))
        v7 := v1["a"]
        var v8 int = -1
        _ = v6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        if (0 != -21) {
            fmt.Println(v3, v5)
            v2 = (v2 ^ 190)
            if true {
                for fi10 := 0; fi10 < 19; fi10++ {
                    delete(v1, "a")
                    if (false && false) {
                        v4 = 41
                        var v11 uint8 = (^(v4))
                        if false {
                            v12 := v1["a"]
                            _ = v12
                        } else {
                            v4 = (v11 - (v11 % (((v4 >> ((1) & 63))) | 1)))
                        }
                        _ = v11
                    }
                    v13 := GenS1{F0: 65535, F1: 96}
                    _ = fi10
                    _ = v13
                }
                v2 -= 660
            }
        } else {
            var v14 int8 = 100
            var v15 int32 = 552
            if true {
                var v16 uint64 = 978
                if false {
                    if (!true) {
                        for fi17 := 0; fi17 < 19; fi17++ {
                            var v18 uint32 = (v2 & (v2 << ((566) & 63)))
                            var v19 int16 = v3
                            v16 += v16
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        v20 := GenS1{F0: 0, F1: int64(v15)}
                        _ = v20
                    }
                    var v21 int16 = (v3 & int16(fi9))
                    var v22 string = ""
                    _ = v21
                    _ = v22
                }
                _ = v16
            }
            _ = v14
            _ = v15
        }
        if true {
            switch "a" {
            case "\n\t":
                v23 := GenS1{F0: 65535, F1: -9223372036854775808}
                _ = v23
            case "foo bar":
                switch "hello" {
                case "foo bar":
                    v3 = 42
                }
            default:
                for fi24 := 0; fi24 < 10; fi24++ {
                    v25 := GenS1{F0: 1, F1: 42}
                    var v26 uint8 = (v4 / ((1) | 1))
                    _ = len(v1)
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
        }
        v4 += (v4 & 128)
        _ = fi9
    }
    if false {
        v3 = 566
    }
    v3 -= v3
    fmt.Println(v2, v3)
    switch "foo bar" {
    case "foo bar":
        if (!false) {
            for fi27 := 0; fi27 < 6; fi27++ {
                var v28 string = ("foo bar" + (("\n\t" + "\n\t") + ""))
                fmt.Println(v2)
                _ = fi27
                _ = v28
            }
            switch -814 {
            case 127:
                var v29 float64 = 0.0
                _ = v29
            case 42:
                switch "hello" {
                case "hello":
                    delete(v1, "hello")
                case "foo bar":
                    v30 := GenS2{F0: 1}
                    _ = v30
                case "a":
                    v31 := GenS3{F0: 88, F1: 402, F2: 921}
                    _ = v31
                default:
                    v4 = 0
                }
            }
            var v32 int8 = 127
            _ = v32
        }
    }
    fmt.Println(v5, v4, v5)
    switch 1 {
    case -846:
        v3 = (-(v3))
    case -1:
        var v33 uint8 = v4
        _ = v33
    case -613:
        fmt.Println(v2)
    }
    v34 := GenS1{F0: 0, F1: 9223372036854775807}
    v35 := GenS4{F0: 2147483647, F1: (593 > uint64(v2)), F2: 42}
    if false {
        if (!false) {
            v36 := []int{675, 2147483647, 32768}
            _ = v36
        }
    }
    if false {
        for fi37 := 0; fi37 < 4; fi37++ {
            v38 := GenS5{F0: "\n\t"}
            v5 = v5
            fmt.Println(v4, fi37, v2)
            _ = fi37
            _ = v38
        }
    }
    var v39 float32 = 0.0
    for fi40 := 0; fi40 < 2; fi40++ {
        v39 = v39
        v41 := []int{fi40, (fi40 << ((574) & 63)), 128, (fi40 % ((fi40) | 1))}
        _ = fi40
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v39
    _ = v4
    _ = v5
}
