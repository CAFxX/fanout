package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 string }

func main() {
    var v1 uint16 = 457
    var v2 int8 = 1
    v3 := make(map[uint32]bool)
    v4 := GenS1{F0: -798, F1: ("hello" + "hello"), F2: ("" + ("foo bar" + ("hello" + "foo bar")))}
    fmt.Println(v1, v2)
    v5 := GenS1{F0: 128, F1: "a", F2: "hello"}
    var v6 bool = (42 <= 1000)
    switch int(v2) {
    case 9223372036854775807:
        var v7 bool = (92 == 115)
        _ = v7
    case 32767:
        v8 := []int{-2147483648}
        _ = v8
    case 42:
        if (891 <= 2) {
            v6 = (18446744073709551615 >= 2)
            v9 := []int{112, 42, -936}
            if (((v1 != 42) && ((v1 << ((225) & 63)) >= (v1 + (^(v1))))) || (!v6)) {
                fmt.Println(v2, v2, v1)
                v1 -= (^(v1))
            } else {
                switch 2147483647 {
                case 32768:
                    for fi10 := 0; fi10 < 1; fi10++ {
                        v1 = (v1 & (^(v1)))
                        _ = fi10
                    }
                case -998:
                    var v11 uint32 = 803
                    _ = v11
                case 127:
                    var v12 uint16 = v1
                    _ = v12
                }
                var v13 uint32 = 296
                var v14 int64 = 42
                _ = v13
                _ = v14
            }
            _ = v9
        } else {
            var v15 uint16 = (^(v1))
            _ = v15
        }
    default:
        if v6 {
            var v16 int32 = -1
            fmt.Println(v2)
            _ = v16
        }
    }
    v17 := GenS1{F0: 32768, F1: "\n\t", F2: "\n\t"}
    v1 = (v1 * v1)
    switch 0 {
    case 42:
        if ((!(true || (!v6))) || (v6 && (0 != 1))) {
            var v18 uint64 = 0
            if true {
                _ = len(v3)
            } else {
                var v19 uint32 = 84
                var v20 int32 = 0
                v21 := []int{}
                _ = v19
                _ = v20
                _ = v21
            }
            _ = v18
        }
    case 0:
        var v22 bool = false
        _ = v22
    }
    v23 := GenS1{F0: 32767, F1: ((("" + "") + ("a" + ("hello" + ("foo bar" + "foo bar")))) + (("foo bar" + ("hello" + ("a" + (("a" + "hello") + ("foo bar" + "hello"))))) + "foo bar")), F2: "foo bar"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
