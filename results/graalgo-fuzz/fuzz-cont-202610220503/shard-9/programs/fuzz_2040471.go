package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint; F1 int32 }

func main() {
    var v1 uint = 65535
    var v2 bool = (true && (90 <= 210))
    var v3 int = 127
    var v4 string = ""
    var v5 int16 = -32768
    for fi6 := 0; fi6 < 7; fi6++ {
        fmt.Println(v4, v2)
        _ = fi6
    }
    fmt.Println(v2)
    switch v3 {
    case 128:
        v4 = v4
    default:
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 int = v3
            _ = fi7
            _ = v8
        }
    }
    if false {
        var v9 float64 = 0.1
        if v2 {
            v3 -= 0
            v1 += (^(v1))
        } else {
            switch 978 {
            case 128:
                v4 = v4
            case 2147483647:
                switch (v3 / ((2147483647) | 1)) {
                case 0:
                    var v10 int16 = (v5 % ((v5) | 1))
                    _ = v10
                case 128:
                    fmt.Println(v9, v2)
                }
            case -541:
                var v11 uint8 = 0
                _ = v11
            }
            var v12 uint32 = 2
            v1 -= (v1 - v1)
            _ = v12
        }
        for fi13 := 0; fi13 < 9; fi13++ {
            if (154 > 18446744073709551615) {
                switch "foo bar" {
                case "a":
                    for fi14 := 0; fi14 < 3; fi14++ {
                        v9 = (0.1 - v9)
                        var v15 int8 = -1
                        v16 := GenS1{F0: 0}
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                default:
                    v17 := GenS1{F0: 855}
                    _ = v17
                }
                v18 := GenS1{F0: 2}
                switch v4 {
                case "":
                    if false {
                        for fi19 := 0; fi19 < 15; fi19++ {
                            var v20 float32 = (0.1 * 3.14159)
                            v21 := GenS1{F0: 1}
                            v22 := GenS2{F0: v4, F1: -9223372036854775808, F2: 42}
                            _ = fi19
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                        var v23 uint32 = 494
                        var v24 uint64 = 2
                        _ = v23
                        _ = v24
                    } else {
                        v25 := GenS3{F0: 578, F1: -548}
                        v5 += v5
                        v2 = false
                        _ = v25
                    }
                case "hello":
                    switch "" {
                    case "\n\t":
                        fmt.Println(v5)
                    default:
                        v26 := []int{fi13, fi13, -843, v3}
                        _ = v26
                    }
                default:
                    v27 := GenS3{F0: (^(v1)), F1: 2147483647}
                    _ = v27
                }
                _ = v18
            }
            v5 += 32767
            for fi28 := 0; fi28 < 10; fi28++ {
                v29 := GenS2{F0: "hello", F1: 564, F2: 2}
                _ = fi28
                _ = v29
            }
            _ = fi13
        }
        _ = v9
    }
    var v30 bool = (42 < 4294967295)
    var v31 int16 = v5
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
