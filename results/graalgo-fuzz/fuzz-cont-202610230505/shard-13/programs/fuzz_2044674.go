package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int32; F1 string; F2 uint32 }
func genFn0() uint64 {
    var v4 uint = 1
    if ((("" + ("a" + ("" + (("" + (("\n\t" + "") + "foo bar")) + "\n\t")))) + ("foo bar" + "hello")) == (("foo bar" + (("hello" + "hello") + "\n\t")) + "\n\t")) {
        var v5 string = "hello"
        _ = v5
    }
    _ = v4
    return 458
}

func main() {
    var v1 uint64 = 2
    var v2 bool = true
    v3 := make(map[int]bool)
    switch 2147483647 {
    case 145:
        v6 := []int{}
        _ = v6
    default:
        var v7 uint = 2
        _ = v7
    }
    v1 -= genFn0()
    v1 -= 4294967295
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 uint32 = 2147483647
        var v10 bool = (v2 || false)
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        if (32767 > 42) {
            v12 := GenS1{F0: fi11, F1: -888}
            switch (fi11 >> ((42) & 63)) {
            case -2147483648:
                var v13 float32 = ((1e10 - 1e10) * 0.0)
                _ = v13
            case 130:
                fmt.Println(v1, fi11)
            default:
                if (456 >= 213) {
                    for fi14 := 0; fi14 < 13; fi14++ {
                        v15 := []int{}
                        _ = fi14
                        _ = v15
                    }
                    v16 := GenS2{F0: -2147483648, F1: "a", F2: 2147483647}
                    _ = v16
                }
            }
            _ = v12
        } else {
            var v17 uint64 = 0
            for fi18 := 0; fi18 < 1; fi18++ {
                if (1 <= -1) {
                    var v19 int16 = -102
                    v20 := GenS1{F0: (fi18 << ((65535) & 63)), F1: 1}
                    if (!v2) {
                        v21 := []int{361, (fi18 / (((fi18 >> ((586) & 63))) | 1))}
                        _ = v21
                    } else {
                        var v22 int8 = 42
                        _ = v22
                    }
                    _ = v19
                    _ = v20
                } else {
                    v23 := []int{970, fi11}
                    if len(v23) > 0 {
                        v23[0] = (^(fi18))
                    }
                    v2 = v2
                    _ = v23
                }
                _ = fi18
            }
            _ = v17
        }
        _ = fi11
    }
    v24 := v3[-9223372036854775808]
    var v25 uint = 0
    _ = len(v3)
    fmt.Println(v2)
    if (255 <= 2) {
        var v26 float32 = (1.5 + 0.0)
        v27 := GenS2{F0: 2147483647, F1: "hello", F2: 333}
        _ = v26
        _ = v27
    } else {
        switch "hello" {
        case "hello":
            switch 2147483647 {
            case 128:
                v24 = false
            }
        case "\n\t":
            for fi28 := 0; fi28 < 18; fi28++ {
                if false {
                    if v24 {
                        v29 := GenS2{F0: 42, F1: "a", F2: 65535}
                        fmt.Println(v1, v24)
                        _ = v29
                    }
                } else {
                    var v30 int32 = -775
                    _ = v30
                }
                var v31 int = (fi28 % ((fi28) | 1))
                _ = fi28
                _ = v31
            }
        case "":
            if (225 >= 65535) {
                for fi32 := 0; fi32 < 17; fi32++ {
                    v1 = (^(v1))
                    _ = fi32
                }
                for fi33 := 0; fi33 < 3; fi33++ {
                    if v2 {
                        v25 = 639
                    } else {
                        var v34 int16 = 408
                        v24 = (fi33 != (fi33 % ((135) | 1)))
                        _ = v34
                    }
                    if (!((1e10 - 1.5) != 1e10)) {
                        var v35 uint32 = 65535
                        var v36 string = ""
                        _ = v35
                        _ = v36
                    }
                    _ = fi33
                }
                var v37 string = "hello"
                _ = v37
            }
        default:
            var v38 uint8 = 2
            _ = v38
        }
        var v39 float32 = 3.14159
        _ = v39
    }
    v1 -= 18446744073709551615
    switch "" {
    case "\n\t":
        v40 := GenS1{F0: -375, F1: -2147483648}
        _ = v40
    case "foo bar":
        v41 := GenS1{F0: 4294967295, F1: 2147483647}
        _ = v41
    default:
        if v2 {
            v24 = v24
            var v42 uint = (v25 + (v25 >> (((v25 & 255)) & 63)))
            var v43 int = 645
            _ = v42
            _ = v43
        } else {
            v44 := GenS2{F0: -1, F1: "foo bar", F2: 1}
            fmt.Println(v2)
            v25 = (v25 & v25)
            _ = v44
        }
    }
    v45 := GenS2{F0: -990, F1: ("foo bar" + ("" + ((("a" + "") + ("a" + "hello")) + "\n\t"))), F2: 883}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v45
}
