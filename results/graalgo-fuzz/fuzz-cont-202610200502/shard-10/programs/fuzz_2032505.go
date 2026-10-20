package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 uint16 }
func genFn0(p0 bool, p1 string, p2 uint) uint32 {
    var v6 int8 = -128
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 0, F1: 115, F2: (1 <= 470)}
        _ = fi7
        _ = v8
    }
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}

func main() {
    var v1 uint16 = 686
    var v2 uint = 65535
    var v3 float32 = 3.14159
    var v4 int8 = 82
    var v5 uint8 = 162
    var v9 uint = (^(v2))
    if false {
        switch "a" {
        case "foo bar":
            for fi10 := 0; fi10 < 5; fi10++ {
                for fi11 := 0; fi11 < 4; fi11++ {
                    var v12 int8 = (v4 - -1)
                    switch (("hello" + ((("\n\t" + (("hello" + "foo bar") + (("\n\t" + ("a" + "")) + ("hello" + "hello")))) + (("" + "\n\t") + ("foo bar" + "foo bar"))) + ("a" + "hello"))) + "") {
                    case "\n\t":
                        v13 := []int{}
                        _ = v13
                    case "foo bar":
                        if false {
                            v4 = v4
                            var v14 int = fi11
                            _ = v14
                        }
                    case "a":
                        fmt.Println(v5)
                    }
                    v2 = (v9 | 18446744073709551615)
                    _ = fi11
                    _ = v12
                }
                _ = fi10
            }
        case "":
            var v15 float32 = (1e10 * -2.25)
            _ = v15
        default:
            v16 := GenS2{F0: v1}
            _ = v16
        }
    } else {
        if true {
            switch 32767 {
            case -810:
                for fi17 := 0; fi17 < 10; fi17++ {
                    var v18 uint8 = uint8(v2)
                    var v19 uint16 = (v1 & 255)
                    var v20 int = fi17
                    _ = fi17
                    _ = v18
                    _ = v19
                    _ = v20
                }
            case 571:
                v21 := []int{9223372036854775807, 42, -1, 196}
                _ = v21
            }
        }
    }
    v1 += 0
    var v22 int32 = -2147483648
    v23 := GenS2{F0: v1}
    for fi24 := 0; fi24 < 6; fi24++ {
        for fi25 := 0; fi25 < 4; fi25++ {
            var v26 uint8 = (v5 % ((v5) | 1))
            _ = fi25
            _ = v26
        }
        v27 := GenS3{F0: 1, F1: 2}
        fmt.Println(v22, v9, v2)
        _ = fi24
        _ = v27
    }
    v28 := []int{74, int(v1), -807, -360}
    v2 += (v2 / ((v2) | 1))
    for fi29 := 0; fi29 < 14; fi29++ {
        v28 = append(v28, int(v9))
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
