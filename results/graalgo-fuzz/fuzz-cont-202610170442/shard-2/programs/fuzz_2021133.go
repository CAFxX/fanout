package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint8; F1 int8; F2 int32 }
type GenS3 struct { F0 int32; F1 uint; F2 uint64 }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string; F1 string; F2 uint32 }
type GenS7 struct { F0 int16; F1 int16 }
func genFn0() uint8 {
    switch 345 {
    case 154:
        v4 := GenS1{F0: 50, F1: -309}
        _ = v4
    case -740:
        v5 := GenS2{F0: 0, F1: 100, F2: 1}
        _ = v5
    case 0:
        switch "foo bar" {
        case "\n\t":
            if (65535 < 42) {
                var v6 bool = true
                _ = v6
            }
        }
    }
    switch "a" {
    case "\n\t":
        v7 := []int{32767}
        _ = v7
    default:
        v8 := GenS3{F0: 42, F1: 4294967295, F2: 0}
        _ = v8
    }
    v9 := GenS4{F0: 42, F1: 251}
    var v10 uint64 = 18446744073709551615
    switch (("" + ((("hello" + "\n\t") + ("\n\t" + ("a" + ("foo bar" + ("hello" + "\n\t"))))) + "a")) + "foo bar") {
    case "hello":
        v10 = 4294967295
    case "a":
        var v11 bool = true
        _ = v11
    default:
        v12 := []int{32767}
        _ = v12
    }
    _ = v10
    _ = v9
    return 166
}
func genFn1() int32 {
    var v13 string = ("a" + ("\n\t" + "foo bar"))
    if (!(0.1 < (1e10 - float64(3.14159)))) {
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 int16 = 42
            for fi16 := 0; fi16 < 18; fi16++ {
                fmt.Println(v15, v13)
                _ = fi16
            }
            _ = fi14
            _ = v15
        }
        v17 := GenS2{F0: 0, F1: -1, F2: -1}
        _ = v17
    }
    v18 := []int{32767, -2147483648}
    _ = v13
    _ = v18
    return 42
}

func main() {
    v1 := []int{32767}
    v2 := make(map[uint32]string)
    var v3 float32 = float32(-2.25)
    var v19 int16 = 741
    switch 32767 {
    case 0:
        var v20 uint8 = 1
        _ = v20
    case 2147483647:
        var v21 int8 = -96
        _ = v21
    case -2147483648:
        var v22 int = 127
        _ = v22
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        v24 := GenS5{F0: 846}
        v19 -= (v19 % (((v19 ^ (v19 + (v19 * 0)))) | 1))
        v25 := GenS6{F0: ("hello" + ((("" + ("foo bar" + "foo bar")) + "") + ("\n\t" + "hello"))), F1: ("a" + "a"), F2: 42}
        _ = fi23
        _ = v24
        _ = v25
    }
    v26 := v2[322]
    for fi27 := 0; fi27 < 1; fi27++ {
        for fi28 := 0; fi28 < 16; fi28++ {
            v19 -= (v19 | (-(v19)))
            v26 = (v26 + (v26 + v26))
            _ = fi28
        }
        _ = fi27
    }
    v29 := GenS3{F0: 0, F1: 4294967295, F2: 1}
    fmt.Println(v3, v3, v26)
    switch 4294967295 {
    case -1:
        for fi30 := 0; fi30 < 3; fi30++ {
            for fi31 := 0; fi31 < 8; fi31++ {
                var v32 bool = (((1.5 / 1.5) * 1e10) <= -2.25)
                _ = fi31
                _ = v32
            }
            var v33 uint32 = 2147483648
            _ = fi30
            _ = v33
        }
    case 2147483647:
        switch v26 {
        case "foo bar":
            for fi34 := 0; fi34 < 16; fi34++ {
                v3 = v3
                _ = fi34
            }
        case "a":
            fmt.Println(v26)
        }
    default:
        v19 = v19
    }
    v35 := GenS4{F0: 2, F1: 18446744073709551615}
    if (!((0.0 + 0.0) > v3)) {
        switch 4294967295 {
        case 32767:
            delete(v2, 42)
        case 2147483647:
            fmt.Println(v3, v19, v19)
        }
        var v36 uint16 = 1
        _ = v36
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        if (1 == genFn1()) {
            for fi38 := 0; fi38 < 3; fi38++ {
                if true {
                    var v39 int64 = -9223372036854775808
                    delete(v2, 399)
                    var v40 uint8 = 0
                    _ = v39
                    _ = v40
                } else {
                    var v41 int8 = 0
                    v42 := GenS4{F0: 42, F1: 869}
                    if false {
                        var v43 string = v26
                        _ = v43
                    }
                    _ = v41
                    _ = v42
                }
                v1 = append(v1, 261)
                v44 := GenS7{F0: (v19 ^ v19), F1: v19}
                _ = fi38
                _ = v44
            }
        }
        v19 = (^(v19))
        _ = fi37
    }
    var v45 string = ("foo bar" + v26)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v19
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v35
    _ = v45
}
