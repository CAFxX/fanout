package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint; F1 int64; F2 uint }
type GenS3 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = "\n\t"
    var v3 uint = 255
    var v4 int64 = -9223372036854775808
    v1[2147483648] = false
    delete(v1, 2)
    v2 = v2
    v3 += (v3 << ((0) & 63))
    if false {
        if ((v3 | v3) > 18446744073709551615) {
            v3 -= (^(v3))
        } else {
            var v5 uint64 = 18446744073709551615
            _ = v5
        }
        switch 42 {
        case 32768:
            v4 += (v4 / ((v4) | 1))
        }
        switch ((v2 + v2) + "a") {
        case "a":
            var v6 uint64 = 42
            _ = v6
        default:
            switch v2 {
            case "hello":
                for fi7 := 0; fi7 < 3; fi7++ {
                    v3 = (^(v3))
                    _ = fi7
                }
            case "\n\t":
                for fi8 := 0; fi8 < 20; fi8++ {
                    var v9 uint = v3
                    _ = fi8
                    _ = v9
                }
            case "":
                var v10 int32 = 249
                _ = v10
            }
        }
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 uint32 = 2147483648
        var v13 string = "foo bar"
        _ = fi11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        v15 := GenS1{F0: v4, F1: 42, F2: uint32(v3)}
        v4 -= v4
        v1[12] = (-2147483648 < -760)
        _ = fi14
        _ = v15
    }
    if true {
        v16 := v1[2]
        v16 = (!(-697 == -2147483648))
        _ = v16
    }
    switch ((v2 + ("hello" + v2)) + v2) {
    case "a":
        for fi17 := 0; fi17 < 20; fi17++ {
            _ = len(v1)
            if ((-128 >= 0) && true) {
                fmt.Println(fi17, v4)
            }
            v2 = "a"
            _ = fi17
        }
    case "":
        var v18 uint64 = 2
        _ = v18
    case "foo bar":
        var v19 int16 = -498
        _ = v19
    }
    switch -9223372036854775808 {
    case 1:
        v3 = (v3 ^ 488)
    case 4294967295:
        v2 = (((v2 + "foo bar") + v2) + "\n\t")
    case 2147483647:
        switch "\n\t" {
        case "a":
            fmt.Println(v3, v2)
        case "foo bar":
            fmt.Println(v2)
        default:
            v20 := GenS2{F0: v3, F1: (v4 << ((65535) & 63)), F2: v3}
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        v1[0] = true
        var v22 string = ""
        _ = fi21
        _ = v22
    }
    if false {
        for fi23 := 0; fi23 < 9; fi23++ {
            for fi24 := 0; fi24 < 6; fi24++ {
                v3 -= (v3 % ((v3) | 1))
                _ = fi24
            }
            var v25 int = fi23
            _ = fi23
            _ = v25
        }
        v26 := GenS3{F0: 100}
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
