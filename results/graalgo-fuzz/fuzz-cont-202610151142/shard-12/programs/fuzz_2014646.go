package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint; F1 string; F2 int64 }
type GenS3 struct { F0 int8 }
func genFn0() bool {
    var v4 uint8 = 2
    v4 += v4
    var v5 uint16 = 1
    v6 := GenS1{F0: v4, F1: 18446744073709551615, F2: (^(v5))}
    v4 -= v4
    _ = v4
    _ = v5
    _ = v6
    return false
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    for fi7 := 0; fi7 < 19; fi7++ {
        if ((false && true) && (true && (((^(fi7)) <= fi7) && false))) {
            switch (fi7 + (^(fi7))) {
            case 32767:
                var v8 int32 = 2147483647
                _ = v8
            }
        } else {
            v1 -= 108
            fmt.Println(fi7)
        }
        _ = fi7
    }
    var v9 int64 = 0
    if false {
        var v10 int32 = 1
        if (!true) {
            for fi11 := 0; fi11 < 7; fi11++ {
                switch "a" {
                case "\n\t":
                    v12 := []int{9223372036854775807}
                    _ = v12
                }
                for fi13 := 0; fi13 < 10; fi13++ {
                    var v14 float32 = (1.5 - float32(-2.25))
                    v15 := GenS2{F0: 18446744073709551615, F1: "a", F2: v9}
                    fmt.Println(v9, v14)
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = fi11
            }
            if (65535 != 586) {
                var v16 bool = (true || false)
                v9 = -1
                _ = v16
            } else {
                var v17 uint16 = (v1 + (v1 | 2))
                var v18 float32 = (1.5 / 1e10)
                for fi19 := 0; fi19 < 3; fi19++ {
                    var v20 uint16 = (v17 ^ 404)
                    _ = fi19
                    _ = v20
                }
                _ = v17
                _ = v18
            }
        } else {
            if true {
                var v21 uint32 = 0
                _ = v21
            }
            switch ("foo bar" + "") {
            case "foo bar":
                for fi22 := 0; fi22 < 16; fi22++ {
                    var v23 uint32 = 4294967295
                    var v24 string = ("foo bar" + "a")
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            case "\n\t":
                v9 -= (v9 << ((2) & 63))
            }
        }
        switch "foo bar" {
        case "a":
            v9 = (v9 << ((255) & 63))
        case "":
            var v25 int32 = (v10 % (((^(v10))) | 1))
            _ = v25
        default:
            var v26 uint8 = 42
            _ = v26
        }
        _ = v10
    } else {
        if true {
            if false {
                v1 = v1
            }
            fmt.Println(v9, v9)
            var v27 bool = (v1 <= 1)
            _ = v27
        }
        v28 := []int{}
        _ = v28
    }
    if false {
        v1 = v1
        for fi29 := 0; fi29 < 20; fi29++ {
            for fi30 := 0; fi30 < 9; fi30++ {
                fmt.Println(v1)
                var v31 int64 = v9
                _ = fi30
                _ = v31
            }
            for fi32 := 0; fi32 < 16; fi32++ {
                v33 := GenS2{F0: 65535, F1: (("a" + "\n\t") + (("a" + "hello") + "a")), F2: (v9 | v9)}
                _ = fi32
                _ = v33
            }
            switch (fi29 >> ((229) & 63)) {
            case 186:
                var v34 bool = true
                _ = v34
            case 2147483647:
                v9 = (v9 >> ((331) & 63))
            default:
                fmt.Println(v1, v1, v9)
            }
            _ = fi29
        }
    } else {
        switch -2147483648 {
        case 643:
            v1 += v1
        case 9223372036854775807:
            var v35 int64 = (-(v9))
            _ = v35
        case 42:
            v36 := GenS1{F0: 127, F1: 763, F2: v1}
            _ = v36
        default:
            for fi37 := 0; fi37 < 10; fi37++ {
                if (3.14159 < ((0.0 - (0.1 + 3.14159)) - (1e10 + 0.1))) {
                    for fi38 := 0; fi38 < 20; fi38++ {
                        v3["foo bar"] = 1
                        fmt.Println(v9)
                        _ = fi38
                    }
                }
                _ = fi37
            }
        }
        switch "" {
        case "":
            var v39 uint64 = 18446744073709551615
            _ = v39
        default:
            var v40 uint64 = 0
            _ = v40
        }
        v41 := GenS3{F0: 100}
        _ = v41
    }
    v42 := GenS1{F0: 0, F1: 18446744073709551615, F2: (v1 * 61)}
    v1 += v1
    var v43 float32 = 3.14159
    v44 := GenS2{F0: 2, F1: "", F2: (v9 - (-(v9)))}
    var v45 uint32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v9
}
