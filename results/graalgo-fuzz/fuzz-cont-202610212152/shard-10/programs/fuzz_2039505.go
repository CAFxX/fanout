package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int16; F1 int; F2 int16 }
type GenS3 struct { F0 bool }

func main() {
    var v1 bool = false
    v2 := make(map[string]int)
    var v3 uint64 = 18446744073709551615
    v4 := make(map[int]bool)
    var v5 uint16 = 675
    fmt.Println(v1, v3, v5)
    v6 := GenS1{F0: 18446744073709551615, F1: 0}
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 float32 = (3.14159 - -2.25)
            v10 := GenS1{F0: (v3 - 0), F1: -130}
            v11 := GenS1{F0: 0, F1: 841}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        _ = fi7
    }
    var v12 bool = v1
    var v13 uint = 538
    v5 += (^(v5))
    for fi14 := 0; fi14 < 1; fi14++ {
        v15 := []int{32767}
        _ = fi14
        _ = v15
    }
    delete(v4, -9223372036854775808)
    if ("a" >= "hello") {
        if v1 {
            switch "hello" {
            case "foo bar":
                for fi16 := 0; fi16 < 9; fi16++ {
                    v17 := GenS1{F0: 734, F1: -920}
                    _ = fi16
                    _ = v17
                }
            case "hello":
                var v18 float64 = 0.0
                _ = v18
            default:
                v19 := []int{int(v13), -623}
                _ = v19
            }
        } else {
            for fi20 := 0; fi20 < 2; fi20++ {
                if true {
                    var v21 uint64 = v3
                    fmt.Println(v5)
                    v22 := GenS2{F0: 0, F1: (fi20 + fi20), F2: -32768}
                    _ = v21
                    _ = v22
                } else {
                    for fi23 := 0; fi23 < 19; fi23++ {
                        if true {
                            v24 := GenS1{F0: 990, F1: 0}
                            _ = v24
                        }
                        _ = len(v2)
                        v13 = 843
                        v3 -= (v3 + v3)
                        _ = fi23
                    }
                }
                v3 = uint64(v5)
                v25 := GenS3{F0: (!(2147483648 < 642))}
                _ = fi20
                _ = v25
            }
        }
        v13 = (v13 ^ 0)
        if (!((2 <= 1) || (245 == 12))) {
            var v26 uint8 = 0
            fmt.Println(v3, v12, v1)
            var v27 int8 = 0
            _ = v26
            _ = v27
        }
    } else {
        if v12 {
            var v28 uint8 = 128
            v29 := GenS3{F0: v12}
            _ = v28
            _ = v29
        } else {
            if v12 {
                for fi30 := 0; fi30 < 19; fi30++ {
                    v13 -= 736
                    v31 := GenS1{F0: 60, F1: -230}
                    var v32 int16 = 32767
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                switch "foo bar" {
                case "\n\t":
                    var v33 uint64 = v3
                    _ = v33
                }
            }
        }
    }
    var v34 uint = (v13 & (^(v13)))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
