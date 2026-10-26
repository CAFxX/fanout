package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int64 }
func genFn0() uint32 {
    v6 := []int{}
    var v7 bool = true
    v7 = true
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 uint = 1
    v2 := []int{2147483647, -461, 42}
    var v3 int32 = -2147483648
    var v4 string = (("" + ("\n\t" + (("" + "foo bar") + "hello"))) + "\n\t")
    var v5 int16 = 32767
    for fi8 := 0; fi8 < 5; fi8++ {
        v1 += v1
        var v9 int32 = v3
        _ = fi8
        _ = v9
    }
    if true {
        v5 += (v5 >> ((v1) & 63))
        v5 = v5
    } else {
        fmt.Println(v4)
        switch v4 {
        case "":
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
        default:
            var v10 float32 = 1.5
            _ = v10
        }
    }
    var v11 int16 = (v5 << (((v1 ^ (^(v1)))) & 63))
    fmt.Println(v3)
    fmt.Println(v5)
    v1 = v1
    for fi12 := 0; fi12 < 16; fi12++ {
        v1 = (v1 | (v1 + v1))
        _ = fi12
    }
    if false {
        fmt.Println(v11, v5, v4)
        if true {
            v4 = v4
        } else {
            var v13 int16 = (v5 | v5)
            _ = v13
        }
    } else {
        v14 := GenS1{F0: 128, F1: 97, F2: 42}
        _ = v14
    }
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    fmt.Println(v11, v5, v5)
    for fi15 := 0; fi15 < 9; fi15++ {
        for fi16 := 0; fi16 < 4; fi16++ {
            if false {
                v17 := GenS2{F0: genFn0()}
                if (false || (true && true)) {
                    v18 := GenS2{F0: 894}
                    var v19 uint32 = 2
                    var v20 uint32 = v19
                    _ = v18
                    _ = v19
                    _ = v20
                } else {
                    v11 = v11
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                switch (fi15 + -301) {
                case 32768:
                    var v21 uint64 = uint64(v1)
                    _ = v21
                case 42:
                    if (("foo bar" + v4) == ("a" + v4)) {
                        v1 -= v1
                        v4 = (v4 + "\n\t")
                    }
                    v22 := GenS3{F0: 4294967295, F1: 501}
                    _ = v22
                default:
                    v23 := GenS3{F0: 42, F1: 973}
                    _ = v23
                }
                _ = v17
            }
            switch 982 {
            case -877:
                v5 = 0
            default:
                v11 = v5
            }
            v11 -= v11
            _ = fi16
        }
        fmt.Println(v11, v4, v1)
        _ = fi15
    }
    var v24 int16 = v5
    var v25 uint32 = 198
    var v26 int = int(v25)
    var v27 uint8 = 255
    var v28 bool = (1e10 != 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
