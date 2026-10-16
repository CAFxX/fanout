package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 uint8; F2 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = -724
    v3 := make(map[uint32]bool)
    v4 := make(map[int]string)
    var v5 float64 = float64(1.5)
    var v6 uint64 = 0
    switch -9223372036854775808 {
    case 4294967295:
        v6 = v6
    case -774:
        delete(v1, 65535)
    case 872:
        v2 = (v2 & 42)
    }
    switch (("foo bar" + ("a" + ((("" + "a") + ("hello" + "foo bar")) + ("hello" + "\n\t")))) + "\n\t") {
    case "hello":
        v2 = (v2 * (v2 - (v2 ^ (-(v2)))))
    case "":
        var v7 uint = 954
        _ = v7
    default:
        var v8 string = "hello"
        _ = v8
    }
    v9 := GenS1{F0: int16(v6)}
    v10 := []int{-655, 4294967295}
    fmt.Println(v6)
    switch "a" {
    case "hello":
        for fi11 := 0; fi11 < 11; fi11++ {
            switch "hello" {
            case "hello":
                v2 -= -777
            case "":
                v6 -= v6
            case "a":
                v6 += v6
            default:
                fmt.Println(v2, v6)
            }
            fmt.Println(v2, fi11, v2)
            _ = fi11
        }
    case "\n\t":
        if (false && (396 <= 1)) {
            if true {
                for fi12 := 0; fi12 < 20; fi12++ {
                    v2 -= v2
                    _ = fi12
                }
            } else {
                var v13 float64 = v5
                _ = v13
            }
            var v14 int8 = 100
            v2 = (v2 << ((799) & 63))
            _ = v14
        } else {
            v2 = 9223372036854775807
            var v15 int32 = -788
            _ = v15
        }
    case "":
        v16 := GenS2{F0: 942}
        _ = v16
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        switch ("hello" + "a") {
        case "hello":
            var v18 int8 = -128
            _ = v18
        }
        var v19 uint32 = 2147483647
        v2 -= 1
        _ = fi17
        _ = v19
    }
    v20 := GenS1{F0: -1}
    for fi21 := 0; fi21 < 4; fi21++ {
        var v22 bool = (!true)
        var v23 int64 = 525
        fmt.Println(fi21, v5)
        _ = fi21
        _ = v22
        _ = v23
    }
    v2 = (v2 >> ((4294967295) & 63))
    fmt.Println(v6, v2, v2)
    switch -853 {
    case -1:
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 uint64 = (v6 << ((1) & 63))
            _ = fi24
            _ = v25
        }
    case -2147483648:
        v26 := GenS1{F0: 944}
        _ = v26
    case -9223372036854775808:
        v2 -= (-(v2))
    default:
        v5 -= (float64(0.0) + v5)
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        v28 := GenS3{F0: false, F1: 1, F2: 42}
        _ = fi27
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
