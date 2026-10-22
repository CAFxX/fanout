package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int }
type GenS2 struct { F0 uint64; F1 int16 }

func main() {
    var v1 int32 = -1
    v2 := make(map[uint32]uint32)
    var v3 string = "a"
    var v4 string = v3
    v5 := []int{2147483647, 4294967295, 9223372036854775807}
    fmt.Println(v1)
    switch 42 {
    case -919:
        _ = len(v2)
    case 42:
        v4 = v3
    case 32767:
        var v6 int32 = -2147483648
        _ = v6
    default:
        if (-9223372036854775808 < -1) {
            if (v3 >= v4) {
                var v7 int16 = -260
                _ = v7
            }
        }
    }
    if true {
        _ = len(v2)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        var v8 int32 = (^(v1))
        switch ("a" + v4) {
        case "":
            if (!((!true) && false)) {
                var v9 int16 = 0
                _ = v9
            } else {
                var v10 uint64 = 2
                _ = v10
            }
        }
        _ = v8
    }
    if false {
        switch (v3 + v3) {
        case "a":
            fmt.Println(v1, v4)
        default:
            for fi11 := 0; fi11 < 1; fi11++ {
                switch (fi11 - 0) {
                case 127:
                    for fi12 := 0; fi12 < 16; fi12++ {
                        fmt.Println(fi12)
                        var v13 uint16 = 42
                        v5 = append(v5, (-(fi11)))
                        _ = fi12
                        _ = v13
                    }
                case -2147483648:
                    switch (-(fi11)) {
                    case 352:
                        v14 := GenS1{F0: (fi11 << ((2) & 63)), F1: 0, F2: fi11}
                        _ = v14
                    case 338:
                        v15 := GenS1{F0: (fi11 % ((-151) | 1)), F1: 1, F2: 127}
                        _ = v15
                    case 2147483647:
                        for fi16 := 0; fi16 < 16; fi16++ {
                            v17 := GenS1{F0: 0, F1: 4294967295, F2: fi16}
                            _ = fi16
                            _ = v17
                        }
                    }
                case 361:
                    for fi18 := 0; fi18 < 17; fi18++ {
                        var v19 int64 = 1
                        var v20 uint = 1
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                default:
                    var v21 int16 = 32767
                    _ = v21
                }
                for fi22 := 0; fi22 < 10; fi22++ {
                    v2[65535] = 2147483648
                    v23 := GenS2{F0: 2, F1: 1}
                    _ = fi22
                    _ = v23
                }
                _ = fi11
            }
        }
        var v24 string = ("" + v4)
        v3 = "a"
        _ = v24
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        for fi26 := 0; fi26 < 7; fi26++ {
            if true {
                v1 = -577
            } else {
                _ = len(v2)
            }
            _ = fi26
        }
        _ = fi25
    }
    v1 += 234
    if len(v5) > 0 {
        v5[0] = 32768
    }
    v27 := GenS2{F0: 42, F1: 1}
    v1 = (v1 / (((v1 - (v1 / ((2147483647) | 1)))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
