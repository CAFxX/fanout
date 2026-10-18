package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 uint }

func main() {
    var v1 int32 = -588
    var v2 int16 = 1000
    var v3 uint32 = 42
    var v4 uint16 = 42
    var v5 int16 = v2
    v6 := GenS1{F0: 9223372036854775807, F1: ""}
    switch ("\n\t" + ("\n\t" + "\n\t")) {
    case "":
        switch 0 {
        case 42:
            if (true || true) {
                for fi7 := 0; fi7 < 17; fi7++ {
                    v4 -= (^(v4))
                    _ = fi7
                }
            } else {
                var v8 uint64 = 1
                _ = v8
            }
        case 32767:
            var v9 int8 = 127
            _ = v9
        }
    }
    if (329 >= 1) {
        if false {
            var v10 int16 = (v5 | v2)
            _ = v10
        }
        fmt.Println(v4)
    } else {
        var v11 uint8 = uint8(v4)
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 4; fi13++ {
            v2 = (-(v2))
            _ = fi13
        }
        v14 := []int{(fi12 - (fi12 << ((18446744073709551615) & 63))), -2147483648, (fi12 >> ((4294967295) & 63))}
        _ = fi12
        _ = v14
    }
    v15 := []int{-2147483648, 128}
    var v16 int64 = 884
    switch 42 {
    case -143:
        v3 = v3
    }
    v16 += int64(v2)
    v4 = v4
    v17 := GenS2{F0: 2147483647, F1: v4, F2: v4}
    for fi18 := 0; fi18 < 14; fi18++ {
        v3 -= 65535
        _ = fi18
    }
    if false {
        if len(v15) > 0 {
            v15[0] = -2147483648
        }
        v19 := GenS3{F0: 4294967295}
        v2 += (^(v2))
        _ = v19
    } else {
        if (42 > -2147483648) {
            v1 = (v1 >> ((0) & 63))
            var v20 uint = 399
            _ = v20
        }
        if true {
            v2 = v2
        } else {
            if false {
                v21 := GenS1{F0: -9223372036854775808, F1: ("" + "")}
                if len(v15) > 0 {
                    v15[0] = -9223372036854775808
                }
                if false {
                    v2 = v2
                    fmt.Println(v16, v2, v3)
                    v3 += v3
                }
                _ = v21
            }
        }
    }
    v15 = append(v15, 128)
    for fi22 := 0; fi22 < 20; fi22++ {
        v23 := GenS4{F0: 18446744073709551615, F1: 18446744073709551615}
        v24 := GenS3{F0: v3}
        fmt.Println(v3, fi22)
        _ = fi22
        _ = v23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
