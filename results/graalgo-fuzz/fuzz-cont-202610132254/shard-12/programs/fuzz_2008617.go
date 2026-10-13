package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 bool; F1 int16 }
func genFn0() bool {
    var v7 float32 = 1.5
    v7 -= 1.5
    _ = v7
    return true
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 255
    var v4 int32 = -2147483648
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]int)
    var v8 uint = 745
    var v9 bool = false
    switch 4294967295 {
    case 665:
        if false {
            switch "hello" {
            case "hello":
                v10 := GenS1{F0: 238}
                _ = v10
            case "a":
                if v1 {
                    if (!true) {
                        for fi11 := 0; fi11 < 7; fi11++ {
                            var v12 uint = v8
                            if v1 {
                                v2[109] = 2147483647
                            } else {
                                v1 = true
                                v13 := []int{(fi11 & 559), (^(fi11)), (fi11 >> ((v12) & 63))}
                                _ = v13
                            }
                            var v14 uint64 = 1
                            _ = fi11
                            _ = v12
                            _ = v14
                        }
                    }
                } else {
                    v6[""] = 969
                    for fi15 := 0; fi15 < 8; fi15++ {
                        v9 = v9
                        v16 := GenS2{F0: "a", F1: (((0.1 / 1e10) * 0.1) == (-2.25 - float64(0.1)))}
                        _ = fi15
                        _ = v16
                    }
                    var v17 bool = ((4294967295 <= 65535) && ((255 >= 128) && (4294967295 < 4294967295)))
                    _ = v17
                }
            default:
                v18 := GenS2{F0: "foo bar", F1: (4294967295 != v8)}
                _ = v18
            }
            v1 = v1
        }
    case 127:
        for fi19 := 0; fi19 < 9; fi19++ {
            v2[537] = 551
            var v20 int64 = 457
            v1 = ((32767 >= 1) || v1)
            _ = fi19
            _ = v20
        }
    default:
        if (2147483648 == 2) {
            var v21 uint32 = 42
            _ = v21
        } else {
            for fi22 := 0; fi22 < 20; fi22++ {
                var v23 int64 = 1
                _ = fi22
                _ = v23
            }
            var v24 uint8 = 127
            v9 = (v9 && (false && false))
            _ = v24
        }
    }
    var v25 int16 = 32767
    fmt.Println(v9, v8)
    switch ("\n\t" + "") {
    case "a":
        for fi26 := 0; fi26 < 12; fi26++ {
            v27 := GenS3{F0: true, F1: v25}
            _ = fi26
            _ = v27
        }
    default:
        for fi28 := 0; fi28 < 17; fi28++ {
            for fi29 := 0; fi29 < 8; fi29++ {
                fmt.Println(v8)
                v4 = v4
                v30 := GenS2{F0: (("hello" + ((("\n\t" + "a") + (("a" + ("foo bar" + "")) + "a")) + "a")) + ("foo bar" + "foo bar")), F1: ((((!v1) && (!v1)) || (4294967295 >= 742)) && v9)}
                _ = fi29
                _ = v30
            }
            _ = fi28
        }
    }
    v25 -= v25
    var v31 int32 = v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
