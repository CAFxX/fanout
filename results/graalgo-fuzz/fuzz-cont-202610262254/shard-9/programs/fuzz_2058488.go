package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint16 }
type GenS2 struct { F0 int64; F1 bool; F2 string }
type GenS3 struct { F0 uint16; F1 int; F2 int32 }

func main() {
    var v1 int16 = 547
    var v2 int32 = 2147483647
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    v5 := make(map[uint32]uint32)
    var v6 bool = (!(-1 > -1))
    if (42 <= -484) {
        var v7 float64 = float64((0.0 * 1.5))
        v8 := GenS1{F0: "foo bar", F1: -2147483648, F2: 2}
        _ = v7
        _ = v8
    }
    v6 = (-2147483648 > -18)
    if v6 {
        var v9 int8 = 100
        var v10 int8 = v9
        _ = v10
        _ = v9
    }
    switch "" {
    case "":
        switch "hello" {
        case "\n\t":
            v1 = (v1 >> ((42) & 63))
        case "":
            for fi11 := 0; fi11 < 18; fi11++ {
                v12 := []int{(fi11 / (((fi11 >> ((578) & 63))) | 1)), fi11}
                v13 := GenS2{F0: -488, F1: false, F2: "\n\t"}
                v6 = false
                _ = fi11
                _ = v12
                _ = v13
            }
        case "foo bar":
            for fi14 := 0; fi14 < 7; fi14++ {
                delete(v4, fi14)
                v1 = (v1 + (^(v1)))
                v6 = v6
                _ = fi14
            }
        }
    default:
        switch -9223372036854775808 {
        case -2147483648:
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 uint64 = 2
                _ = fi15
                _ = v16
            }
        case 299:
            v1 = v1
        case 32767:
            switch 1 {
            case 2147483647:
                v17 := GenS2{F0: 1, F1: ((!((v2 << ((0) & 63)) == v2)) || v6), F2: ("" + "a")}
                _ = v17
            case 32768:
                for fi18 := 0; fi18 < 17; fi18++ {
                    for fi19 := 0; fi19 < 2; fi19++ {
                        var v20 int64 = 561
                        _ = fi19
                        _ = v20
                    }
                    _ = fi18
                }
            }
        default:
            v1 -= (^(v1))
        }
    }
    var v21 int = 4294967295
    fmt.Println(v2, v1)
    var v22 int16 = v1
    var v23 float64 = 3.14159
    var v24 string = ((((((("hello" + "\n\t") + "") + ("hello" + "foo bar")) + "") + "foo bar") + "hello") + "")
    switch (v21 * 2147483647) {
    case -970:
        v25 := GenS1{F0: v24, F1: (^(v21)), F2: 0}
        _ = v25
    case 9223372036854775807:
        v26 := GenS3{F0: 32767, F1: v21, F2: v2}
        _ = v26
    case 127:
        if v6 {
            v1 = v1
            v27 := GenS2{F0: -198, F1: (((0 >= 109) && true) && (-2.25 == 1.5)), F2: ("a" + "a")}
            var v28 float32 = (3.14159 / float32(-2.25))
            _ = v27
            _ = v28
        }
    default:
        v24 = v24
    }
    var v29 int32 = (v2 - 1)
    v1 += 32767
    v30 := GenS2{F0: -1, F1: v6, F2: ("\n\t" + (v24 + "\n\t"))}
    var v31 uint64 = 4294967295
    if v6 {
        v6 = (!(644 < 2))
        v32 := []int{2147483647, 9223372036854775807}
        _ = v32
    }
    v31 += (v31 >> ((0) & 63))
    if false {
        switch "" {
        case "foo bar":
            for fi33 := 0; fi33 < 6; fi33++ {
                _ = len(v5)
                fmt.Println(v22)
                _ = fi33
            }
        case "\n\t":
            var v34 int = v21
            _ = v34
        }
    } else {
        if v6 {
            var v35 int64 = -1
            _ = v35
        }
        var v36 int8 = 0
        v37 := GenS3{F0: 1, F1: v21, F2: (v29 - 1)}
        _ = v36
        _ = v37
    }
    var v38 string = ((v24 + ((v24 + v24) + (v24 + v24))) + "a")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
