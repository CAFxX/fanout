package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int16 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int = -769
    var v3 int8 = 42
    var v4 int32 = 42
    var v5 int64 = -9223372036854775808
    fmt.Println(v2)
    v1 -= 42
    fmt.Println(v5, v2)
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 uint32 = 4294967295
        v4 += 2147483647
        for fi8 := 0; fi8 < 12; fi8++ {
            v9 := GenS1{F0: (-(v5)), F1: 124, F2: 1}
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v10 int8 = v3
    v11 := []int{v2}
    if true {
        if true {
            var v12 uint = 480
            v13 := GenS1{F0: v5, F1: v10, F2: 1000}
            for fi14 := 0; fi14 < 1; fi14++ {
                if (1000 >= 0) {
                    for fi15 := 0; fi15 < 18; fi15++ {
                        if ((v5 % ((v5) | 1)) != (v5 + v5)) {
                            v5 -= (v5 ^ (v5 << (((v12 ^ v12)) & 63)))
                        }
                        v16 := GenS1{F0: (v5 | (v5 >> (((^(v12))) & 63))), F1: -1, F2: -497}
                        _ = fi15
                        _ = v16
                    }
                }
                _ = fi14
            }
            _ = v12
            _ = v13
        }
        for fi17 := 0; fi17 < 13; fi17++ {
            var v18 float32 = 1e10
            var v19 uint16 = 32767
            var v20 int32 = v4
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        v21 := GenS1{F0: (^(v5)), F1: 42, F2: 992}
        _ = v21
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        fmt.Println(v2, fi22, v3)
        var v23 uint16 = 65535
        _ = fi22
        _ = v23
    }
    v24 := GenS1{F0: v5, F1: (v10 + -70), F2: 0}
    for fi25 := 0; fi25 < 18; fi25++ {
        if false {
            if (false || (132 > 4294967295)) {
                if (false || (v10 > v10)) {
                    v2 = 128
                } else {
                    switch fi25 {
                    case 2147483647:
                        fmt.Println(fi25, v2, v10)
                    case 1:
                        var v26 int64 = 1
                        _ = v26
                    }
                    v27 := GenS1{F0: v5, F1: 0, F2: 1000}
                    fmt.Println(v10, v1, v1)
                    _ = v27
                }
            }
            v1 = v1
        }
        if false {
            v1 = 545
            var v28 string = ("\n\t" + "foo bar")
            _ = v28
        } else {
            v29 := GenS1{F0: v5, F1: 127, F2: -1}
            v10 = (v10 - (v10 | v10))
            _ = v29
        }
        for fi30 := 0; fi30 < 11; fi30++ {
            switch (v2 * (-(v2))) {
            case 575:
                v31 := GenS1{F0: v5, F1: 1, F2: 0}
                _ = v31
            case -1:
                _ = len(v11)
            case 0:
                var v32 int64 = v5
                _ = v32
            }
            if false {
                fmt.Println(fi25)
                var v33 uint = 42
                fmt.Println(v1, v2)
                _ = v33
            }
            _ = len(v11)
            _ = fi30
        }
        _ = fi25
    }
    v1 = 1
    for fi34 := 0; fi34 < 14; fi34++ {
        for fi35 := 0; fi35 < 7; fi35++ {
            var v36 uint32 = 2147483648
            _ = len(v11)
            _ = fi35
            _ = v36
        }
        var v37 int8 = (v10 >> ((0) & 63))
        _ = fi34
        _ = v37
    }
    if len(v11) > 0 {
        v11[0] = v2
    }
    v5 += v5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
