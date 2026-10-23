package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint64; F2 int }
type GenS3 struct { F0 int8; F1 uint; F2 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64 }
func genFn0() bool {
    v6 := GenS1{F0: 1}
    var v7 float64 = ((0.1 - 3.14159) - 3.14159)
    switch -9223372036854775808 {
    case 42:
        v7 = 3.14159
    case -9223372036854775808:
        v8 := []int{32767, 722}
        _ = v8
    }
    var v9 uint32 = 0
    _ = v6
    _ = v7
    _ = v9
    return false
}

func main() {
    var v1 uint8 = 255
    var v2 uint32 = 0
    var v3 float32 = 0.0
    var v4 int64 = 9223372036854775807
    var v5 uint32 = (v2 / ((0) | 1))
    v1 = v1
    v5 = (v5 | (v2 + (^(v2))))
    var v10 uint16 = 463
    for fi11 := 0; fi11 < 18; fi11++ {
        if (1 >= 594) {
            switch -285 {
            case 32767:
                switch (^(fi11)) {
                case 128:
                    v10 -= v10
                }
            case 2147483647:
                var v12 bool = (!(!(-128 > 35)))
                _ = v12
            default:
                for fi13 := 0; fi13 < 14; fi13++ {
                    if true {
                        var v14 bool = (2 < 486)
                        _ = v14
                    } else {
                        v15 := GenS2{F0: 42, F1: 647, F2: (fi13 / (((fi11 % ((1) | 1))) | 1))}
                        _ = v15
                    }
                    v5 = 4294967295
                    _ = fi13
                }
            }
            v16 := []int{(fi11 / (((fi11 - fi11)) | 1)), fi11, fi11}
            var v17 string = "hello"
            _ = v16
            _ = v17
        }
        _ = fi11
    }
    v10 = v10
    v2 = (v5 * (v2 * 1))
    for fi18 := 0; fi18 < 15; fi18++ {
        if true {
            var v19 uint64 = 18446744073709551615
            _ = v19
        }
        if (false && (("\n\t" != "") && false)) {
            if (true || true) {
                var v20 int64 = 534
                _ = v20
            }
            fmt.Println(v1, fi18, v10)
            v21 := []int{int(v5), fi18, fi18}
            _ = v21
        } else {
            switch "foo bar" {
            case "a":
                var v22 int64 = (v4 << ((65535) & 63))
                _ = v22
            case "\n\t":
                v23 := GenS1{F0: -152}
                _ = v23
            default:
                v10 -= 551
            }
            v24 := []int{(-(fi18)), 95, 0}
            for fi25 := 0; fi25 < 6; fi25++ {
                v26 := GenS2{F0: 1, F1: 1, F2: fi18}
                fmt.Println(v5, v2)
                _ = fi25
                _ = v26
            }
            _ = v24
        }
        _ = fi18
    }
    switch (("\n\t" + "") + "\n\t") {
    case "":
        if false {
            v27 := GenS3{F0: 127, F1: 590, F2: 4294967295}
            _ = v27
        } else {
            for fi28 := 0; fi28 < 7; fi28++ {
                for fi29 := 0; fi29 < 2; fi29++ {
                    var v30 int8 = 2
                    var v31 int = 1
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                _ = fi28
            }
            fmt.Println(v2, v2)
        }
    default:
        for fi32 := 0; fi32 < 13; fi32++ {
            var v33 uint8 = (v1 << ((871) & 63))
            _ = fi32
            _ = v33
        }
    }
    var v34 int16 = 1000
    v2 = v2
    v35 := GenS1{F0: (v4 + (v4 ^ 1))}
    v10 -= v10
    v36 := GenS3{F0: 100, F1: 594, F2: 2}
    v37 := GenS4{F0: "\n\t"}
    fmt.Println(v34, v2, v4)
    switch int(v4) {
    case 32767:
        var v38 uint = 0
        _ = v38
    case 2147483647:
        if true {
            fmt.Println(v4, v5, v3)
        }
    }
    for fi39 := 0; fi39 < 2; fi39++ {
        v5 -= 186
        _ = fi39
    }
    if (!false) {
        fmt.Println(v34)
        var v40 float64 = -2.25
        _ = v40
    } else {
        v41 := GenS5{F0: 230}
        var v42 uint8 = (v1 ^ v1)
        _ = v41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
