package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64 }

func main() {
    var v1 uint32 = 731
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]bool)
    v4 := make(map[int]int)
    fmt.Println(v1)
    if true {
        fmt.Println(v1)
    } else {
        var v5 int16 = 1
        v6 := []int{-9223372036854775808, 806, 37, -260}
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 string = ("hello" + "foo bar")
            var v9 string = v8
            switch (fi7 + 42) {
            case -457:
                var v10 float64 = 1.5
                _ = v10
            default:
                v11 := GenS1{F0: 42}
                _ = v11
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = v5
        _ = v6
    }
    if (1 == 2) {
        var v12 int8 = 0
        v12 -= v12
        if true {
            for fi13 := 0; fi13 < 18; fi13++ {
                var v14 int32 = 0
                v15 := GenS2{F0: 0, F1: (v1 & v1)}
                _ = fi13
                _ = v14
                _ = v15
            }
            switch 32767 {
            case 32768:
                v16 := GenS3{F0: 691}
                _ = v16
            case 689:
                for fi17 := 0; fi17 < 17; fi17++ {
                    switch (("a" + "\n\t") + ("foo bar" + ("\n\t" + ""))) {
                    case "\n\t":
                        v12 = v12
                    }
                    var v18 string = "foo bar"
                    v19 := GenS4{F0: 2}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            case -2147483648:
                v20 := GenS4{F0: 255}
                _ = v20
            default:
                var v21 uint = 2
                _ = v21
            }
        } else {
            if (!true) {
                var v22 int64 = -793
                if false {
                    v22 = v22
                    switch 804 {
                    case 1:
                        var v23 int16 = 138
                        _ = v23
                    }
                }
                _ = v22
            }
            for fi24 := 0; fi24 < 17; fi24++ {
                var v25 uint64 = 42
                switch fi24 {
                case 1:
                    v12 += v12
                default:
                    if (!(809 != 685)) {
                        v25 = (v25 & (v25 + (v25 >> ((65535) & 63))))
                    }
                }
                _ = fi24
                _ = v25
            }
            v12 -= 127
        }
        _ = v12
    }
    v1 = v1
    v26 := v2["hello"]
    var v27 float64 = ((0.0 / 1e10) / 0.1)
    var v28 uint16 = 306
    if (18446744073709551615 != 42) {
        for fi29 := 0; fi29 < 15; fi29++ {
            fmt.Println(v27, v27, fi29)
            v27 += 0.1
            _ = fi29
        }
        for fi30 := 0; fi30 < 18; fi30++ {
            var v31 int = (fi30 << ((852) & 63))
            _ = fi30
            _ = v31
        }
        v28 = (v28 << ((84) & 63))
    } else {
        var v32 uint64 = 336
        _ = v32
    }
    fmt.Println(v26, v27)
    v33 := GenS5{F0: 104}
    var v34 bool = (!(uint(v28) <= 255))
    var v35 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
