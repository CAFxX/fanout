package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint8) int8 {
    for fi4 := 0; fi4 < 14; fi4++ {
        v5 := GenS1{F0: ("hello" + "a"), F1: 2}
        _ = fi4
        _ = v5
    }
    var v6 int = -2147483648
    for fi7 := 0; fi7 < 19; fi7++ {
        switch fi7 {
        case -261:
            v6 = 42
        }
        v8 := GenS1{F0: (("foo bar" + ("foo bar" + ("hello" + "a"))) + ("" + "a")), F1: 1}
        v9 := GenS1{F0: "a", F1: 42}
        _ = fi7
        _ = v8
        _ = v9
    }
    switch 9223372036854775807 {
    case -1:
        for fi10 := 0; fi10 < 4; fi10++ {
            v6 -= v6
            switch "a" {
            case "foo bar":
                var v11 int16 = 1
                _ = v11
            case "hello":
                for fi12 := 0; fi12 < 5; fi12++ {
                    fmt.Println(fi12, fi12, fi12)
                    v6 = -2147483648
                    _ = fi12
                }
            case "":
                p0 -= 0
            }
            v6 = fi10
            _ = fi10
        }
    case 747:
        v13 := []int{42}
        _ = v13
    }
    _ = p0
    _ = v6
    return 9
}
func genFn1() uint64 {
    v14 := []int{955, -220}
    var v15 bool = (!((!true) && true))
    var v16 uint16 = 1
    _ = v14
    _ = v15
    _ = v16
    return 42
}

func main() {
    var v1 string = "hello"
    var v2 int64 = 9223372036854775807
    v3 := []int{127, 128, 9223372036854775807, 128}
    v2 = v2
    v1 = (v1 + v1)
    v2 -= (v2 | (v2 >> ((2) & 63)))
    for fi17 := 0; fi17 < 15; fi17++ {
        v18 := GenS1{F0: ("foo bar" + "\n\t"), F1: 65535}
        var v19 string = v1
        var v20 string = ("a" + v1)
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    if len(v3) > 0 {
        v3[0] = -2147483648
    }
    _ = len(v3)
    _ = len(v3)
    v2 -= (v2 << ((0) & 63))
    fmt.Println(v2, v2)
    v1 = ("a" + "\n\t")
    var v21 uint64 = 42
    _ = len(v3)
    var v22 string = v1
    if ((v21 / ((v21) | 1)) >= v21) {
        fmt.Println(v22, v21)
        if false {
            if false {
                v2 += 0
                if true {
                    var v23 uint64 = v21
                    var v24 string = (v22 + "hello")
                    var v25 uint16 = 490
                    _ = v23
                    _ = v24
                    _ = v25
                } else {
                    if (!false) {
                        v22 = "foo bar"
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        switch v1 {
                        case "hello":
                            fmt.Println(v21, v2)
                        default:
                            v1 = v1
                        }
                    }
                    v21 = 18446744073709551615
                }
            } else {
                v2 += 1
                var v26 int64 = v2
                _ = v26
            }
        }
        fmt.Println(v2, v2, v22)
    } else {
        for fi27 := 0; fi27 < 18; fi27++ {
            var v28 int64 = 42
            _ = fi27
            _ = v28
        }
    }
    var v29 int32 = 569
    for fi30 := 0; fi30 < 16; fi30++ {
        var v31 bool = false
        _ = fi30
        _ = v31
    }
    v32 := GenS2{F0: 9223372036854775807}
    if true {
        v2 -= (-(v2))
        v21 = v21
        var v33 bool = (!(58 < -127))
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v32
}
