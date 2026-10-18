package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int }
type GenS2 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 int; F1 uint8 }

func main() {
    var v1 string = ("foo bar" + "a")
    var v2 uint64 = 18446744073709551615
    v3 := make(map[uint32]string)
    var v4 uint32 = 0
    var v5 uint8 = 212
    var v6 int8 = -128
    switch v1 {
    case "foo bar":
        switch 4294967295 {
        case -982:
            _ = len(v3)
        case -9223372036854775808:
            var v7 uint16 = 0
            _ = v7
        case 0:
            v8 := []int{42}
            _ = v8
        default:
            var v9 uint64 = 4294967295
            _ = v9
        }
    case "":
        var v10 int = 128
        _ = v10
    case "hello":
        v6 = 127
    default:
        switch -1 {
        case 32767:
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 bool = true
                for fi13 := 0; fi13 < 2; fi13++ {
                    var v14 int16 = 837
                    var v15 int32 = 42
                    v16 := GenS1{F0: (fi11 % ((fi13) | 1)), F1: (v1 + v1), F2: 9223372036854775807}
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
                var v17 uint8 = v5
                _ = fi11
                _ = v12
                _ = v17
            }
        case -2147483648:
            v3[4294967295] = "foo bar"
        }
    }
    v18 := GenS2{F0: (!(147 > (v4 << ((42) & 63)))), F1: v2, F2: 2}
    v19 := GenS3{F0: (32767 > 0), F1: 0}
    var v20 uint = 2
    switch "" {
    case "":
        v21 := GenS1{F0: 875, F1: v1, F2: 128}
        _ = v21
    default:
        if false {
            v6 = 64
            fmt.Println(v4)
            v22 := GenS4{F0: int(v6), F1: 0}
            _ = v22
        } else {
            v2 += v2
        }
    }
    v4 = v4
    v23 := v3[v4]
    if false {
        for fi24 := 0; fi24 < 1; fi24++ {
            for fi25 := 0; fi25 < 4; fi25++ {
                switch ("" + ("\n\t" + v23)) {
                case "a":
                    var v26 int64 = -669
                    _ = v26
                case "foo bar":
                    var v27 string = v23
                    _ = v27
                case "\n\t":
                    v3[0] = v23
                default:
                    v28 := []int{(fi24 / (((fi25 >> (((v20 % ((v20) | 1))) & 63))) | 1)), -569, (fi24 + (fi24 & 32768)), -2147483648}
                    _ = v28
                }
                delete(v3, 42)
                if (9223372036854775807 < -129) {
                    var v29 int32 = 42
                    _ = v29
                } else {
                    switch ("a" + "hello") {
                    case "":
                        v23 = v23
                    case "foo bar":
                        v2 += v2
                    }
                    var v30 int64 = -878
                    if (301 == 75) {
                        var v31 int16 = 376
                        _ = v31
                    } else {
                        fmt.Println(v23, v4, fi24)
                        var v32 int16 = 0
                        v33 := GenS3{F0: (2 == 42), F1: (v30 + (^(v30)))}
                        _ = v32
                        _ = v33
                    }
                    var v34 int64 = (v30 ^ v30)
                    _ = v30
                    _ = v34
                }
                v35 := []int{}
                _ = fi25
                _ = v35
            }
            for fi36 := 0; fi36 < 1; fi36++ {
                fmt.Println(fi36, fi36)
                _ = fi36
            }
            _ = fi24
        }
        v37 := GenS3{F0: (false || (false && true)), F1: -1}
        _ = v37
    }
    var v38 uint = 198
    var v39 float64 = 1e10
    v40 := GenS3{F0: (v23 < (v1 + v1)), F1: int64(v38)}
    for fi41 := 0; fi41 < 13; fi41++ {
        if ((1 == 42) || (1000 <= -1)) {
            v42 := GenS3{F0: (!true), F1: 1}
            v23 = "\n\t"
            _ = v42
        }
        _ = fi41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
