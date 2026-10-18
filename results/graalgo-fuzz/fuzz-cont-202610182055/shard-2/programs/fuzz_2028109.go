package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint16 }

func main() {
    v1 := []string{(((("\n\t" + ("hello" + ("foo bar" + "a"))) + "hello") + "") + ("hello" + (("a" + ("hello" + ("hello" + ("foo bar" + ("a" + "a"))))) + ""))), ("" + ""), "foo bar"}
    var v2 int32 = 0
    v3 := make(map[int]int)
    var v4 uint32 = 4294967295
    var v5 uint64 = 18446744073709551615
    v5 = 1
    v4 = v4
    switch "foo bar" {
    case "foo bar":
        var v6 int = -2147483648
        _ = v6
    case "":
        fmt.Println(v5, v4)
    case "a":
        if (!(float32(float64(-2.25)) <= 1.5)) {
            switch -2147483648 {
            case 74:
                for fi7 := 0; fi7 < 20; fi7++ {
                    for fi8 := 0; fi8 < 6; fi8++ {
                        v9 := GenS1{F0: 881, F1: true, F2: 65535}
                        _ = fi8
                        _ = v9
                    }
                    v2 += 45
                    var v10 int = fi7
                    _ = fi7
                    _ = v10
                }
            default:
                v11 := GenS1{F0: 222, F1: (false || (false || true)), F2: 1}
                _ = v11
            }
            v12 := GenS1{F0: 18446744073709551615, F1: true, F2: 437}
            var v13 float64 = (1.5 * 1e10)
            _ = v12
            _ = v13
        } else {
            if false {
                var v14 float32 = 1e10
                _ = v14
            }
            var v15 uint32 = v4
            _ = v15
        }
    }
    v16 := GenS1{F0: 407, F1: false, F2: 255}
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 12; fi18++ {
            switch (("foo bar" + ("hello" + ("" + "hello"))) + (("" + ("foo bar" + "\n\t")) + "a")) {
            case "hello":
                var v19 uint16 = 177
                _ = v19
            case "foo bar":
                v5 = (v5 * (v5 + 946))
            default:
                v4 -= 613
            }
            delete(v3, fi17)
            _ = fi18
        }
        _ = fi17
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        var v21 string = "hello"
        var v22 float64 = -2.25
        var v23 float32 = ((0.1 - 1.5) * 1.5)
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    v2 -= v2
    if (-859 >= 32768) {
        v2 = (v2 & (v2 ^ (v2 | (-(v2)))))
    }
    for fi24 := 0; fi24 < 17; fi24++ {
        _ = len(v1)
        _ = fi24
    }
    v25 := GenS1{F0: 2, F1: true, F2: 949}
    switch "hello" {
    case "hello":
        var v26 uint16 = 1
        _ = v26
    case "\n\t":
        v5 = 412
    case "a":
        if (false || false) {
            v27 := GenS1{F0: 0, F1: ((!(false && false)) && false), F2: 32767}
            _ = v27
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
