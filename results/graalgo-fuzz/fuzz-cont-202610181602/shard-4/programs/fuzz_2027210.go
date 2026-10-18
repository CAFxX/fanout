package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 string }
type GenS2 struct { F0 uint32; F1 uint32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]int)
    var v3 uint8 = 0
    var v4 uint8 = v3
    var v5 bool = false
    v6 := GenS1{F0: 295, F1: 758, F2: "\n\t"}
    for fi7 := 0; fi7 < 14; fi7++ {
        v5 = ((^(fi7)) >= 128)
        _ = fi7
    }
    v5 = (v5 && true)
    var v8 float32 = 1e10
    v9 := GenS2{F0: 65535, F1: 4294967295}
    fmt.Println(v3)
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            switch ("a" + "foo bar") {
            case "foo bar":
                var v12 int = 32768
                _ = v12
            case "hello":
                if v5 {
                    var v13 int8 = -128
                    _ = v13
                } else {
                    if v5 {
                        v4 += v3
                    } else {
                        v14 := GenS1{F0: 582, F1: 2, F2: "\n\t"}
                        var v15 bool = v5
                        v16 := GenS2{F0: 2147483647, F1: 65535}
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                }
            default:
                switch ("foo bar" + "\n\t") {
                case "":
                    v17 := []int{(fi11 << ((1) & 63))}
                    _ = v17
                case "hello":
                    switch (fi10 << ((255) & 63)) {
                    case -511:
                        var v18 uint64 = 42
                        _ = v18
                    case 9223372036854775807:
                        if false {
                            var v19 float32 = v8
                            var v20 int = (-(fi10))
                            _ = v19
                            _ = v20
                        } else {
                            fmt.Println(v8, v4)
                        }
                    }
                case "\n\t":
                    var v21 float32 = (float32(1.5) - (v8 * float32(0.1)))
                    _ = v21
                }
                var v22 bool = (!(2147483647 > 2))
                fmt.Println(v22, fi11)
                v4 = (v4 + 35)
                if false {
                    v23 := GenS1{F0: 4294967295, F1: 975, F2: "a"}
                    var v24 int = fi10
                    var v25 string = "\n\t"
                    _ = v23
                    _ = v24
                    _ = v25
                }
                _ = v22
            }
            v8 += ((1.5 - (v8 / (v8 * 1.5))) * 3.14159)
            _ = fi11
        }
        _ = fi10
    }
    var v26 float32 = 1.5
    var v27 string = ("\n\t" + "\n\t")
    fmt.Println(v4, v4)
    if ((v4 << ((42) & 63)) != v3) {
        v27 = ("" + "\n\t")
    }
    v28 := GenS1{F0: 57, F1: 920, F2: "foo bar"}
    v26 = (v8 + v26)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
