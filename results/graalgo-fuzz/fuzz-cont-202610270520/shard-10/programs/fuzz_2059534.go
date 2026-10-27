package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint }
type GenS2 struct { F0 int; F1 int32; F2 uint8 }
func genFn0(p0 uint16, p1 uint, p2 uint64) int16 {
    var v7 uint32 = 65535
    fmt.Println(v7, p2, v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -32768
}
func genFn1() bool {
    if (((4294967295 != 2147483648) || (true || (675 >= 2147483647))) && (1000 != 0)) {
        v8 := GenS1{F0: 18446744073709551615, F1: "a", F2: 510}
        var v9 uint = 686
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 484, F1: "", F2: 588}
    _ = v10
    return (1 > 42)
}

func main() {
    var v1 int8 = 124
    v2 := []string{("hello" + "hello")}
    var v3 uint8 = 255
    var v4 int64 = -660
    var v5 bool = (false && true)
    var v6 uint8 = (v3 - v3)
    if ((!false) || (2 == 65535)) {
        if (0 > 427) {
            v5 = false
        }
        switch 42 {
        case -9223372036854775808:
            var v11 uint8 = (^(v3))
            _ = v11
        default:
            var v12 uint64 = 735
            _ = v12
        }
        v6 = (v6 - v3)
    } else {
        _ = len(v2)
        var v13 bool = (0 == v6)
        v14 := GenS1{F0: 610, F1: "foo bar", F2: 0}
        _ = v13
        _ = v14
    }
    if (-1 < (v4 % (((-(v4))) | 1))) {
        v6 = 128
        v6 = 255
    }
    if v5 {
        var v15 float64 = 0.1
        var v16 float32 = -2.25
        if true {
            var v17 float64 = v15
            var v18 uint16 = 65535
            _ = v17
            _ = v18
        }
        _ = v15
        _ = v16
    } else {
        var v19 uint8 = v6
        v6 = (v3 * 0)
        v6 += (v19 / (((v6 & (v19 << ((65535) & 63)))) | 1))
        _ = v19
    }
    v3 -= v3
    v20 := GenS1{F0: 1, F1: "a", F2: 976}
    if (65535 == 32767) {
        for fi21 := 0; fi21 < 13; fi21++ {
            var v22 string = "\n\t"
            switch "foo bar" {
            case "a":
                var v23 uint16 = 715
                _ = v23
            default:
                for fi24 := 0; fi24 < 8; fi24++ {
                    v25 := GenS1{F0: 114, F1: v22, F2: 255}
                    _ = fi24
                    _ = v25
                }
            }
            v26 := GenS2{F0: (^(fi21)), F1: 752, F2: 2}
            _ = fi21
            _ = v22
            _ = v26
        }
        v1 -= -23
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        v1 += v1
        v4 += v4
        _ = fi27
    }
    v28 := GenS1{F0: 391, F1: ("\n\t" + ("foo bar" + "\n\t")), F2: 808}
    switch ("foo bar" + ("\n\t" + (((("foo bar" + "hello") + "hello") + "\n\t") + ("foo bar" + "hello")))) {
    case "\n\t":
        for fi29 := 0; fi29 < 12; fi29++ {
            v2 = append(v2, "hello")
            _ = fi29
        }
    case "hello":
        switch int(v4) {
        case 4294967295:
            for fi30 := 0; fi30 < 17; fi30++ {
                v4 += (v4 >> ((65535) & 63))
                _ = fi30
            }
        case 832:
            v1 -= v1
        case 128:
            if (705 <= 2147483647) {
                for fi31 := 0; fi31 < 17; fi31++ {
                    var v32 bool = true
                    var v33 uint = uint(v3)
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                if (false && v5) {
                    for fi34 := 0; fi34 < 19; fi34++ {
                        var v35 uint32 = 2
                        _ = fi34
                        _ = v35
                    }
                    v3 = (v3 % (((v6 << ((65535) & 63))) | 1))
                }
                v2 = append(v2, "a")
            }
        default:
            fmt.Println(v5)
        }
    }
    var v36 bool = (v5 || (32767 < -2147483648))
    v6 += (^(v6))
    if v5 {
        fmt.Println(v36, v5, v3)
        _ = len(v2)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v37 int32 = 1
    var v38 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
