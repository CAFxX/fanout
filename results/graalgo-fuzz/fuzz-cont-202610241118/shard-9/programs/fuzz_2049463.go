package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint }
func genFn0() bool {
    for fi4 := 0; fi4 < 1; fi4++ {
        v5 := GenS1{F0: 42}
        var v6 float32 = 1.5
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 uint32 = 984
    v8 := GenS1{F0: -1}
    _ = v7
    _ = v8
    return false
}

func main() {
    var v1 int8 = 127
    var v2 int = 127
    v3 := []string{("\n\t" + ("hello" + "hello")), "hello", "hello", ("" + "hello")}
    v9 := GenS1{F0: 621}
    v1 = (^(v1))
    for fi10 := 0; fi10 < 4; fi10++ {
        fmt.Println(v1)
        _ = fi10
    }
    if (32767 < 2) {
        if (3.14159 <= -2.25) {
            var v11 uint16 = 217
            _ = v11
        }
    }
    var v12 int64 = 42
    v13 := GenS1{F0: (v12 * v12)}
    v1 += v1
    fmt.Println(v2, v2, v1)
    switch "\n\t" {
    case "a":
        for fi14 := 0; fi14 < 17; fi14++ {
            if ((!true) && (796 != 1)) {
                switch (fi14 | v2) {
                case 32768:
                    fmt.Println(v2)
                case 9223372036854775807:
                    var v15 float32 = float32(((0.0 - float64(0.0)) - (0.0 / 1.0)))
                    _ = v15
                case 4294967295:
                    var v16 uint16 = 0
                    _ = v16
                default:
                    if false {
                        for fi17 := 0; fi17 < 14; fi17++ {
                            v2 = 4294967295
                            fmt.Println(fi17, fi17)
                            _ = fi17
                        }
                        v3 = append(v3, (("" + "a") + ("a" + ("a" + "\n\t"))))
                        fmt.Println(fi14, v2)
                    }
                }
            }
            v12 += v12
            _ = fi14
        }
    case "\n\t":
        _ = len(v3)
    case "":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        if false {
            v12 -= v12
            switch ("a" + ("" + "foo bar")) {
            case "":
                for fi18 := 0; fi18 < 5; fi18++ {
                    v12 = 42
                    _ = fi18
                }
            case "\n\t":
                if (1 != 1) {
                    switch (v2 * v2) {
                    case 127:
                        v2 = (v2 & v2)
                    }
                    for fi19 := 0; fi19 < 4; fi19++ {
                        v1 -= 100
                        v2 = (^(v2))
                        _ = fi19
                    }
                    v1 -= v1
                }
            case "hello":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                v12 = v12
            }
        } else {
            v20 := GenS2{F0: 42, F1: 335}
            _ = v20
        }
    }
    if false {
        v12 += v12
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
