package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 uint32; F1 int64 }
func genFn0() uint {
    var v7 uint8 = 255
    var v8 uint16 = uint16(v7)
    v9 := []int{-1, 2147483647, -195}
    _ = v7
    _ = v8
    _ = v9
    return 501
}

func main() {
    v1 := make(map[int]int)
    var v2 int64 = -661
    var v3 string = (((("" + "") + ("\n\t" + "\n\t")) + "\n\t") + "\n\t")
    v4 := make(map[uint32]bool)
    var v5 uint32 = 2
    var v6 uint = 0
    v6 -= v6
    var v10 uint16 = 832
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 int16 = -1
        _ = fi11
        _ = v12
    }
    delete(v4, (v5 & (v5 * (v5 >> ((genFn0()) & 63)))))
    for fi13 := 0; fi13 < 17; fi13++ {
        v14 := []int{(fi13 % ((fi13) | 1))}
        var v15 string = "foo bar"
        v16 := GenS1{F0: v10, F1: (^(v6))}
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    switch 42 {
    case 968:
        v17 := GenS1{F0: 1, F1: v6}
        _ = v17
    case -9223372036854775808:
        for fi18 := 0; fi18 < 7; fi18++ {
            v6 = genFn0()
            fmt.Println(v6)
            v19 := GenS1{F0: (v10 ^ v10), F1: (v6 | (v6 / ((642) | 1)))}
            _ = fi18
            _ = v19
        }
    default:
        v20 := GenS2{F0: (v10 & v10), F1: 32767}
        _ = v20
    }
    var v21 uint8 = uint8(v5)
    var v22 bool = true
    var v23 uint64 = 2
    var v24 float32 = (0.0 * 3.14159)
    if v22 {
        for fi25 := 0; fi25 < 8; fi25++ {
            var v26 int64 = v2
            switch fi25 {
            case 9223372036854775807:
                switch (fi25 << (((^(v6))) & 63)) {
                case -2147483648:
                    v4[2147483647] = ((int8(fi25) < 100) && (fi25 <= fi25))
                }
            case 2147483647:
                for fi27 := 0; fi27 < 2; fi27++ {
                    v28 := v1[(fi25 & fi27)]
                    var v29 uint8 = 87
                    fmt.Println(v24, v22)
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            }
            _ = fi25
            _ = v26
        }
        switch 9223372036854775807 {
        case 32767:
            fmt.Println(v5, v23)
        case 1:
            v30 := GenS3{F0: 4294967295, F1: v2}
            _ = v30
        case 2147483647:
            if ((v2 & v2) != v2) {
                if (false || true) {
                    var v31 int8 = 127
                    _ = v31
                }
                switch v3 {
                case "a":
                    var v32 int32 = -2147483648
                    _ = v32
                case "hello":
                    v3 = ""
                default:
                    v33 := v4[4294967295]
                    _ = v33
                }
                if v22 {
                    var v34 uint16 = v10
                    _ = v34
                } else {
                    if true {
                        var v35 uint8 = v21
                        _ = v35
                    } else {
                        var v36 int64 = -192
                        var v37 uint32 = uint32(v23)
                        _ = v36
                        _ = v37
                    }
                }
            } else {
                v38 := []int{792, 4294967295, 127, -2147483648}
                v39 := GenS2{F0: uint16(v5), F1: -836}
                var v40 float32 = v24
                _ = v38
                _ = v39
                _ = v40
            }
        default:
            for fi41 := 0; fi41 < 11; fi41++ {
                v21 = uint8(v6)
                _ = fi41
            }
        }
    }
    v42 := GenS3{F0: 4294967295, F1: v2}
    var v43 uint8 = (v21 ^ 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
