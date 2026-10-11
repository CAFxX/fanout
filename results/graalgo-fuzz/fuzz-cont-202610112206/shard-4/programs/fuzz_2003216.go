package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 uint8 }
func genFn0() int {
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := GenS1{F0: -2147483648}
        for fi8 := 0; fi8 < 16; fi8++ {
            if true {
                fmt.Println(fi6, fi6)
                var v9 float32 = 1.5
                _ = v9
            }
            _ = fi8
        }
        var v10 float32 = ((1.5 * 1e10) / -2.25)
        _ = fi6
        _ = v10
        _ = v7
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 bool = true
        _ = fi11
        _ = v12
    }
    var v13 int16 = 32767
    _ = v13
    return -917
}

func main() {
    var v1 uint = 655
    var v2 uint8 = 1
    var v3 uint16 = 32767
    v4 := make(map[string]bool)
    var v5 uint8 = (v2 / ((73) | 1))
    if true {
        if (-1 == 42) {
            v14 := GenS1{F0: -26}
            v15 := []int{32767, -597}
            _ = v14
            _ = v15
        } else {
            fmt.Println(v3, v5)
        }
        if false {
            v16 := v4["hello"]
            for fi17 := 0; fi17 < 7; fi17++ {
                var v18 uint16 = (v3 | 2)
                _ = fi17
                _ = v18
            }
            v3 += 643
            _ = v16
        }
    }
    _ = len(v4)
    if false {
        v2 = 128
        var v19 uint32 = 0
        _ = v19
    } else {
        var v20 uint32 = 0
        for fi21 := 0; fi21 < 6; fi21++ {
            var v22 float32 = 3.14159
            switch ("\n\t" + "") {
            case "\n\t":
                var v23 string = "\n\t"
                _ = v23
            case "a":
                v24 := []int{(fi21 * -539), 32768}
                _ = v24
            default:
                var v25 uint16 = 2
                _ = v25
            }
            var v26 float64 = 0.1
            _ = fi21
            _ = v22
            _ = v26
        }
        var v27 string = ("hello" + "\n\t")
        _ = v20
        _ = v27
    }
    if true {
        var v28 int64 = 9223372036854775807
        v2 = v2
        _ = v28
    }
    switch 32768 {
    case 0:
        for fi29 := 0; fi29 < 10; fi29++ {
            for fi30 := 0; fi30 < 12; fi30++ {
                fmt.Println(v5, v2)
                switch "" {
                case "hello":
                    v31 := []int{(^(fi29)), (fi29 << (((v1 >> (((v1 ^ 545)) & 63))) & 63)), fi30}
                    _ = v31
                case "a":
                    var v32 uint64 = 621
                    _ = v32
                default:
                    v3 += v3
                }
                _ = fi30
            }
            var v33 uint = v1
            _ = fi29
            _ = v33
        }
    case 42:
        fmt.Println(v1)
    }
    fmt.Println(v1, v1)
    v1 = uint(v2)
    var v34 int16 = 425
    for fi35 := 0; fi35 < 11; fi35++ {
        v2 = (^(v5))
        var v36 int = fi35
        for fi37 := 0; fi37 < 13; fi37++ {
            if (-246 < 9223372036854775807) {
                var v38 int = v36
                _ = v38
            }
            v39 := GenS2{F0: (v2 % ((159) | 1)), F1: 0}
            _ = fi37
            _ = v39
        }
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
