package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint64 }
func genFn0() uint16 {
    v7 := []int{}
    var v8 int = 32767
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 bool = (552 != 32767)
    var v2 uint8 = 2
    v3 := make(map[string]string)
    var v4 int32 = -1
    v5 := make(map[int]string)
    v6 := make(map[uint32]int)
    fmt.Println(v2, v4)
    var v9 bool = (true && (717 >= 0))
    fmt.Println(v2, v1, v4)
    v9 = v1
    v1 = true
    if true {
        for fi10 := 0; fi10 < 6; fi10++ {
            v4 = v4
            _ = fi10
        }
        var v11 uint8 = 151
        v2 = (v11 >> ((65535) & 63))
        _ = v11
    }
    var v12 bool = (v9 || (true || v1))
    if v9 {
        fmt.Println(v2, v12)
        if (-922 < 9223372036854775807) {
            v13 := GenS1{F0: 1, F1: -1}
            _ = v13
        }
    }
    if true {
        v14 := []int{127, 9223372036854775807, 2147483647, -2147483648}
        v1 = false
        _ = v14
    } else {
        var v15 int16 = 0
        _ = v15
    }
    v16 := v5[-241]
    v5[127] = (((("hello" + (v16 + v16)) + ("hello" + v16)) + v16) + v16)
    var v17 string = ("\n\t" + (v16 + (((v16 + v16) + (v16 + (v16 + ((v16 + "foo bar") + v16)))) + v16)))
    for fi18 := 0; fi18 < 10; fi18++ {
        if v1 {
            for fi19 := 0; fi19 < 17; fi19++ {
                if v1 {
                    var v20 float32 = 1e10
                    _ = v20
                } else {
                    v21 := GenS1{F0: 18446744073709551615, F1: 584}
                    switch "hello" {
                    case "":
                        v22 := v6[2]
                        _ = v22
                    default:
                        v23 := GenS1{F0: 496, F1: -729}
                        _ = v23
                    }
                    _ = v21
                }
                fmt.Println(v16, fi18, v16)
                switch (-(fi19)) {
                case 1:
                    v24 := GenS1{F0: 18446744073709551615, F1: 1000}
                    _ = v24
                case 127:
                    switch fi18 {
                    case 32767:
                        var v25 int64 = -1
                        _ = v25
                    }
                    v1 = v12
                    v17 = "a"
                case -1:
                    switch "foo bar" {
                    case "\n\t":
                        switch 1 {
                        case 410:
                            var v26 int = (fi19 ^ 32768)
                            _ = v26
                        }
                    case "a":
                        var v27 uint32 = 104
                        _ = v27
                    case "hello":
                        if true {
                            v1 = (!v1)
                            v28 := []int{2147483647, 1}
                            v17 = v16
                            _ = v28
                        }
                    }
                    var v29 int = fi19
                    v9 = v1
                    _ = v29
                }
                _ = fi19
            }
            for fi30 := 0; fi30 < 18; fi30++ {
                var v31 uint = 2
                _ = fi30
                _ = v31
            }
        } else {
            var v32 float32 = 0.0
            if (genFn0() < uint16(v2)) {
                fmt.Println(v1)
            } else {
                v33 := []int{(^(fi18)), fi18, fi18, fi18}
                _ = len(v3)
                _ = v33
            }
            _ = v32
        }
        v4 -= (v4 | v4)
        _ = fi18
    }
    var v34 float64 = (1e10 * 0.0)
    v35 := GenS2{F0: 0, F1: 18446744073709551615}
    v16 = v17
    fmt.Println(v9)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
