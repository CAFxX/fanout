package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 string; F1 bool; F2 bool }
func genFn0() uint8 {
    v4 := []int{}
    var v5 int64 = 167
    var v6 string = ("a" + "\n\t")
    fmt.Println(v5, v5)
    _ = v4
    _ = v5
    _ = v6
    return 93
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]string)
    var v3 int16 = 32767
    switch 1 {
    case 4294967295:
        switch "\n\t" {
        case "a":
            switch 4294967295 {
            case 4294967295:
                if true {
                    v1[0] = 2147483648
                    v7 := GenS1{F0: 794}
                    _ = v7
                }
            default:
                for fi8 := 0; fi8 < 19; fi8++ {
                    delete(v1, 327)
                    _ = fi8
                }
            }
        case "foo bar":
            var v9 uint64 = 18446744073709551615
            _ = v9
        case "\n\t":
            for fi10 := 0; fi10 < 11; fi10++ {
                v11 := []int{fi10, fi10, (^(fi10))}
                for fi12 := 0; fi12 < 4; fi12++ {
                    v13 := GenS1{F0: 2}
                    if (((((-2.25 + -2.25) * -2.25) - (-2.25 / -2.25)) * 0.1) == (0.0 / (1.5 + 0.0))) {
                        _ = len(v2)
                    }
                    _ = fi12
                    _ = v13
                }
                v14 := GenS2{F0: (fi10 - 579), F1: 58, F2: -32768}
                _ = fi10
                _ = v11
                _ = v14
            }
        }
    case -684:
        v3 = (v3 + (v3 % (((-(v3))) | 1)))
    case 557:
        if (true && (!(!(127 > 127)))) {
            v3 = v3
        }
    default:
        var v15 uint32 = 42
        _ = v15
    }
    v16 := GenS1{F0: 0}
    switch (("\n\t" + "\n\t") + "a") {
    case "":
        var v17 float64 = 1.5
        _ = v17
    case "a":
        if false {
            var v18 string = "hello"
            _ = v18
        }
    case "hello":
        for fi19 := 0; fi19 < 10; fi19++ {
            switch "hello" {
            case "a":
                for fi20 := 0; fi20 < 5; fi20++ {
                    if true {
                        v21 := GenS3{F0: 42, F1: (v3 * v3)}
                        switch ((("\n\t" + ("a" + "foo bar")) + ((("a" + "") + "a") + "foo bar")) + "\n\t") {
                        case "hello":
                            v3 += (v3 | v3)
                        }
                        v22 := GenS4{F0: ("foo bar" + "\n\t"), F1: (0.0 <= 1.5), F2: (2 != 0)}
                        _ = v21
                        _ = v22
                    } else {
                        v23 := GenS1{F0: 18446744073709551615}
                        _ = v23
                    }
                    _ = fi20
                }
            case "hello":
                v3 = v3
            default:
                _ = len(v1)
            }
            var v24 bool = ((3.14159 + float64(0.1)) >= 1.5)
            _ = fi19
            _ = v24
        }
    }
    v25 := v2[(("hello" + "a") + "")]
    var v26 float32 = 3.14159
    v27 := GenS2{F0: -197, F1: 2147483647, F2: (-(v3))}
    v26 = v26
    for fi28 := 0; fi28 < 6; fi28++ {
        _ = len(v1)
        v3 = (v3 - 42)
        _ = fi28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
