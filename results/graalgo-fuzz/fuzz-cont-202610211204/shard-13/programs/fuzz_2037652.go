package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 string; F1 uint32; F2 int16 }
func genFn0() uint8 {
    var v7 uint16 = 652
    var v8 float64 = 0.1
    _ = v7
    _ = v8
    return 255
}

func main() {
    v1 := []string{"foo bar"}
    var v2 uint16 = 869
    v3 := make(map[uint32]string)
    var v4 float64 = 0.1
    var v5 float32 = 1e10
    var v6 int = -416
    var v9 uint64 = 673
    if (v5 == (0.1 + v5)) {
        v10 := GenS1{F0: 42, F1: 35}
        if false {
            for fi11 := 0; fi11 < 15; fi11++ {
                if false {
                    v6 += (fi11 * fi11)
                }
                if (-2147483648 >= -1) {
                    v1 = append(v1, "")
                    var v12 uint32 = 65535
                    _ = v12
                }
                _ = fi11
            }
        } else {
            switch -1 {
            case -9223372036854775808:
                var v13 float32 = float32(1.5)
                _ = v13
            case 128:
                var v14 int64 = 706
                _ = v14
            case 2147483647:
                var v15 uint = 255
                _ = v15
            default:
                fmt.Println(v4)
            }
            var v16 int64 = 9223372036854775807
            _ = v16
        }
        fmt.Println(v5)
        _ = v10
    }
    var v17 int = 42
    v1 = append(v1, ("" + ((("" + "") + "\n\t") + (("a" + "hello") + ("a" + "foo bar")))))
    v18 := GenS2{F0: ("\n\t" + ""), F1: 1, F2: -479}
    fmt.Println(v2, v6, v17)
    var v19 uint16 = 32767
    var v20 float32 = 1.5
    _ = len(v1)
    for fi21 := 0; fi21 < 1; fi21++ {
        for fi22 := 0; fi22 < 6; fi22++ {
            v17 = (v6 / ((128) | 1))
            var v23 string = ("" + "")
            _ = fi22
            _ = v23
        }
        var v24 uint = 901
        _ = fi21
        _ = v24
    }
    _ = len(v3)
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        if (-32768 != 32767) {
            if true {
                var v26 int = (v17 << ((18446744073709551615) & 63))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = len(v3)
                _ = v26
            } else {
                for fi27 := 0; fi27 < 10; fi27++ {
                    var v28 float64 = v4
                    _ = fi27
                    _ = v28
                }
                for fi29 := 0; fi29 < 16; fi29++ {
                    v17 -= (v6 / ((v17) | 1))
                    _ = len(v3)
                    v5 -= 0.0
                    _ = fi29
                }
                var v30 bool = true
                _ = v30
            }
        }
        _ = fi25
    }
    var v31 int8 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
