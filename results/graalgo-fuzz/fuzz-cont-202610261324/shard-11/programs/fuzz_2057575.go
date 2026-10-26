package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 bool }
type GenS2 struct { F0 string }
func genFn0(p0 int8, p1 uint16) int8 {
    var v7 int32 = -291
    if ((!(18446744073709551615 <= 4294967295)) && (p1 != p1)) {
        v8 := []int{2147483647, 128, 876, -372}
        _ = v8
    }
    v7 -= 25
    var v9 uint8 = 18
    _ = p0
    _ = p1
    _ = v7
    _ = v9
    return 1
}

func main() {
    var v1 uint32 = 752
    var v2 float64 = (1e10 * (0.1 * 3.14159))
    var v3 int32 = 0
    var v4 uint16 = 32767
    v5 := []int{42, -604, -747, -933}
    v6 := []string{"foo bar", "hello", ("" + ("" + "hello"))}
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    v2 -= 3.14159
    var v10 uint16 = (v4 + (v4 / (((v4 + v4)) | 1)))
    v11 := GenS1{F0: -2147483648, F1: v3, F2: false}
    v6 = append(v6, "foo bar")
    for fi12 := 0; fi12 < 14; fi12++ {
        if true {
            v13 := GenS1{F0: -9223372036854775808, F1: (v3 * (-(v3))), F2: (!(!false))}
            _ = v13
        }
        var v14 int8 = -128
        if len(v6) > 0 {
            v6[0] = ("foo bar" + (("foo bar" + ("\n\t" + "\n\t")) + "a"))
        }
        _ = fi12
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        for fi16 := 0; fi16 < 17; fi16++ {
            fmt.Println(v10, v4, v1)
            _ = fi16
        }
        var v17 float64 = v2
        _ = fi15
        _ = v17
    }
    v10 = 691
    switch ("a" + ("\n\t" + "")) {
    case "hello":
        var v18 int8 = -40
        _ = v18
    case "a":
        var v19 float32 = 1.5
        _ = v19
    case "":
        var v20 int8 = 127
        _ = v20
    default:
        if false {
            var v21 uint = uint(v10)
            v6 = append(v6, (("foo bar" + (("" + "\n\t") + ("hello" + ""))) + "a"))
            _ = v21
        } else {
            v3 += v3
            fmt.Println(v1)
        }
    }
    var v22 uint16 = (v10 >> ((13) & 63))
    switch "hello" {
    case "hello":
        v23 := GenS1{F0: -9223372036854775808, F1: 42, F2: false}
        _ = v23
    case "a":
        for fi24 := 0; fi24 < 18; fi24++ {
            v25 := GenS2{F0: ((("foo bar" + ("\n\t" + "")) + "") + ("" + (("foo bar" + (("\n\t" + "a") + (("" + ("foo bar" + (("" + ("" + "")) + "a"))) + "foo bar"))) + (("" + "\n\t") + (("a" + "a") + ("hello" + "a"))))))}
            _ = fi24
            _ = v25
        }
    }
    v2 += 1e10
    v3 -= v3
    var v26 bool = (((v1 == uint32(v22)) && false) && true)
    for fi27 := 0; fi27 < 15; fi27++ {
        var v28 float32 = 0.0
        for fi29 := 0; fi29 < 11; fi29++ {
            var v30 bool = v26
            v1 -= (v1 >> ((65535) & 63))
            v26 = (!v26)
            _ = fi29
            _ = v30
        }
        var v31 int32 = v3
        _ = fi27
        _ = v28
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
