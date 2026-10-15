package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint32 }
func genFn0() uint {
    var v4 string = "foo bar"
    v4 = "hello"
    _ = v4
    return 255
}

func main() {
    var v1 float64 = (3.14159 / (0.1 - (3.14159 + -2.25)))
    v2 := make(map[string]uint32)
    var v3 uint32 = 2147483648
    var v5 uint32 = v3
    v1 = v1
    var v6 string = ""
    switch 32768 {
    case -793:
        for fi7 := 0; fi7 < 10; fi7++ {
            for fi8 := 0; fi8 < 18; fi8++ {
                v5 += v3
                switch (fi8 % ((-939) | 1)) {
                case 32767:
                    var v9 uint32 = (v3 ^ v3)
                    _ = v9
                }
                _ = fi8
            }
            v10 := GenS1{F0: (((v6 + (("foo bar" + "\n\t") + ("foo bar" + v6))) > ((v6 + v6) + v6)) || true), F1: ""}
            if (65535 == 42) {
                v5 = 42
                v6 = "foo bar"
                for fi11 := 0; fi11 < 19; fi11++ {
                    v12 := []int{}
                    v13 := GenS1{F0: true, F1: ""}
                    var v14 uint8 = 128
                    if (0 <= 1) {
                        if (v3 > (v3 / ((v3) | 1))) {
                            var v15 uint16 = 42
                            v16 := GenS2{F0: 100}
                            _ = v15
                            _ = v16
                        }
                    } else {
                        v17 := GenS3{F0: fi7, F1: v3}
                        v6 = (("a" + "\n\t") + (("" + "\n\t") + v6))
                        _ = v17
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
            }
            _ = fi7
            _ = v10
        }
    }
    fmt.Println(v5, v1)
    v6 = ("\n\t" + v6)
    v6 = (v6 + ((v6 + "a") + "a"))
    v1 += v1
    v18 := GenS3{F0: -2147483648, F1: 65535}
    fmt.Println(v6, v1, v1)
    fmt.Println(v3, v3, v1)
    v19 := GenS4{F0: 515}
    for fi20 := 0; fi20 < 12; fi20++ {
        v1 += v1
        _ = fi20
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        fmt.Println(v6, v6, v6)
        for fi22 := 0; fi22 < 1; fi22++ {
            v23 := GenS3{F0: (fi22 - fi21), F1: (v5 >> ((1) & 63))}
            fmt.Println(v3, fi22)
            fmt.Println(v6, v1)
            _ = fi22
            _ = v23
        }
        v24 := GenS1{F0: false, F1: "\n\t"}
        _ = fi21
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
