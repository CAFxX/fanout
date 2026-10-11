package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 string }
func genFn0(p0 uint32) int8 {
    p0 += p0
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 string = (("\n\t" + ("a" + ("foo bar" + "a"))) + "")
        _ = fi6
        _ = v7
    }
    var v8 uint = 0
    if (330 != 2147483647) {
        v9 := GenS1{F0: 0, F1: 1, F2: ("hello" + ("a" + "\n\t"))}
        _ = v9
    }
    _ = p0
    _ = v8
    return 88
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 876
    var v3 bool = false
    v4 := make(map[uint32]bool)
    v5 := make(map[string]bool)
    v3 = (true || true)
    var v10 uint8 = 81
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 uint8 = v10
        fmt.Println(v3, v2, v12)
        _ = fi11
        _ = v12
    }
    var v13 uint16 = 2
    _ = len(v4)
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 uint = 0
        _ = fi14
        _ = v15
    }
    v3 = true
    v10 += v10
    for fi16 := 0; fi16 < 11; fi16++ {
        if false {
            fmt.Println(v13)
            v13 -= (v2 << ((42) & 63))
            switch "a" {
            case "hello":
                delete(v5, ("a" + ""))
            default:
                switch (fi16 >> ((255) & 63)) {
                case 0:
                    for fi17 := 0; fi17 < 18; fi17++ {
                        v18 := []int{(fi16 % ((541) | 1)), 128}
                        _ = len(v4)
                        fmt.Println(v10, fi16, fi16)
                        v3 = false
                        _ = fi17
                        _ = v18
                    }
                case -530:
                    _ = len(v5)
                default:
                    v10 = v10
                }
            }
        }
        for fi19 := 0; fi19 < 15; fi19++ {
            v20 := GenS1{F0: 0, F1: 0, F2: ""}
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 2; fi21++ {
            v2 += 595
            v13 -= (v2 | 42)
            if (72 != 42) {
                if ((v3 && false) || (false || true)) {
                    delete(v5, ((("\n\t" + "") + ("foo bar" + (("" + (("foo bar" + (("\n\t" + (((("hello" + "\n\t") + "") + ((("\n\t" + "") + (("foo bar" + "") + "hello")) + ("" + ((("a" + "") + "a") + "foo bar")))) + (("a" + "\n\t") + "a"))) + ("\n\t" + (("hello" + "a") + "\n\t")))) + ("foo bar" + ("" + "\n\t")))) + "a"))) + ((("foo bar" + "hello") + ("" + "")) + "a")))
                }
            } else {
                if v3 {
                    v22 := GenS1{F0: 42, F1: -9223372036854775808, F2: "foo bar"}
                    v3 = v3
                    v23 := GenS1{F0: 2, F1: 0, F2: "\n\t"}
                    v24 := []int{128, (^(fi16))}
                    _ = v22
                    _ = v23
                    _ = v24
                }
                if (-445 >= -9223372036854775808) {
                    var v25 int16 = 1000
                    v10 = 0
                    _ = v25
                } else {
                    v10 -= (v10 >> ((18446744073709551615) & 63))
                }
            }
            _ = fi21
        }
        _ = fi16
    }
    var v26 uint = 904
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
