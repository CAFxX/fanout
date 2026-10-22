package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 string; F1 uint64; F2 int64 }
type GenS3 struct { F0 int16; F1 int8; F2 uint32 }

func main() {
    var v1 uint64 = 2
    v2 := make(map[string]int)
    var v3 string = "hello"
    var v4 uint16 = uint16(v1)
    v5 := make(map[int]uint32)
    switch 127 {
    case -1:
        var v6 uint32 = 736
        _ = v6
    default:
        for fi7 := 0; fi7 < 4; fi7++ {
            for fi8 := 0; fi8 < 15; fi8++ {
                fmt.Println(v3, v3)
                _ = fi8
            }
            _ = fi7
        }
    }
    var v9 int = 1
    for fi10 := 0; fi10 < 1; fi10++ {
        fmt.Println(v1)
        if (-1 >= 1) {
            v1 -= v1
            v2[v3] = v9
        }
        v9 += (v9 % (((v9 | v9)) | 1))
        _ = fi10
    }
    switch 9223372036854775807 {
    case -9223372036854775808:
        for fi11 := 0; fi11 < 17; fi11++ {
            for fi12 := 0; fi12 < 19; fi12++ {
                v4 -= 1
                _ = fi12
            }
            _ = fi11
        }
    }
    v13 := GenS1{F0: ((("" + v3) + v3) + ("foo bar" + "a")), F1: 1}
    _ = len(v2)
    v14 := v5[(v9 | v9)]
    switch v9 {
    case 2147483647:
        var v15 int64 = 9223372036854775807
        _ = v15
    case 177:
        if (true || (433 < 1)) {
            v3 = (v3 + (("foo bar" + "") + ((v3 + "\n\t") + v3)))
            fmt.Println(v14, v9, v4)
            v16 := GenS1{F0: ("" + v3), F1: 100}
            _ = v16
        } else {
            v17 := GenS1{F0: "a", F1: -111}
            fmt.Println(v1, v14)
            v4 = v4
            _ = v17
        }
    case -607:
        switch (v9 ^ (v9 * (v9 & v9))) {
        case 2147483647:
            if ((255 < 2) && false) {
                for fi18 := 0; fi18 < 17; fi18++ {
                    fmt.Println(fi18)
                    v19 := GenS1{F0: v3, F1: 69}
                    _ = fi18
                    _ = v19
                }
                v20 := GenS1{F0: "a", F1: 100}
                v9 = (v9 % ((v9) | 1))
                _ = v20
            }
        case 1:
            var v21 int8 = int8(v9)
            _ = v21
        default:
            switch -824 {
            case -900:
                switch (v9 << ((18446744073709551615) & 63)) {
                case -9223372036854775808:
                    delete(v5, v9)
                }
            case 32768:
                v4 = (^(v4))
            case 42:
                v22 := []int{-158, v9, (v9 << ((65535) & 63)), (v9 % (((^(v9))) | 1))}
                _ = v22
            }
        }
    default:
        fmt.Println(v4, v9, v3)
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        fmt.Println(fi23, fi23, v9)
        if false {
            var v24 int64 = 42
            _ = v24
        } else {
            v25 := GenS1{F0: v3, F1: -128}
            fmt.Println(v3, v3)
            switch v3 {
            case "\n\t":
                v26 := GenS2{F0: v3, F1: v1, F2: 42}
                _ = v26
            case "":
                var v27 string = v3
                _ = v27
            default:
                v28 := []int{}
                _ = v28
            }
            _ = v25
        }
        v29 := v5[(fi23 + int(v4))]
        _ = fi23
        _ = v29
    }
    v30 := v2["hello"]
    v31 := GenS3{F0: 0, F1: 0, F2: (v14 * 356)}
    v9 = (v9 ^ int(v14))
    v9 += v30
    _ = len(v2)
    fmt.Println(v1, v3)
    v9 -= (v9 % ((v30) | 1))
    v9 -= 1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v9
}
