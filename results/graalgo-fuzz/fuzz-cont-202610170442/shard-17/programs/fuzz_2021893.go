package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 uint; F2 uint }

func main() {
    var v1 int = 9223372036854775807
    var v2 int8 = 0
    v3 := make(map[int]string)
    v4 := make(map[string]string)
    var v5 int64 = 0
    v6 := GenS1{F0: "foo bar"}
    switch "hello" {
    case "\n\t":
        v7 := v3[(-(v1))]
        _ = v7
    case "":
        fmt.Println(v5, v1)
    default:
        v8 := GenS1{F0: ("\n\t" + "")}
        _ = v8
    }
    var v9 float32 = 1e10
    var v10 uint64 = 18446744073709551615
    v9 = v9
    var v11 uint = 65535
    if false {
        var v12 uint8 = 2
        fmt.Println(v2, v11, v2)
        _ = v12
    } else {
        var v13 uint16 = 255
        var v14 uint32 = 2147483647
        switch (("\n\t" + "foo bar") + "\n\t") {
        case "hello":
            switch "hello" {
            case "hello":
                v14 = (v14 << ((1) & 63))
            case "\n\t":
                for fi15 := 0; fi15 < 20; fi15++ {
                    var v16 int64 = (v5 - (-(v5)))
                    if false {
                        var v17 int16 = -32768
                        for fi18 := 0; fi18 < 13; fi18++ {
                            var v19 string = (("\n\t" + "a") + "foo bar")
                            switch (v1 / ((fi15) | 1)) {
                            case 0:
                                fmt.Println(v10)
                            }
                            switch v1 {
                            case 32767:
                                v17 -= v17
                            default:
                                fmt.Println(fi15)
                            }
                            v20 := []int{749, (v1 + 9223372036854775807), (fi15 % ((fi18) | 1))}
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        _ = v17
                    }
                    fmt.Println(v11)
                    _ = fi15
                    _ = v16
                }
            default:
                v21 := GenS1{F0: "hello"}
                _ = v21
            }
        case "":
            v22 := GenS2{F0: (v5 & (v5 % ((v5) | 1))), F1: (v11 << (((v11 / ((v11) | 1))) & 63)), F2: (v11 * (v11 * (v11 ^ v11)))}
            _ = v22
        default:
            v9 = (v9 - v9)
        }
        _ = v13
        _ = v14
    }
    if true {
        v23 := GenS1{F0: ((("\n\t" + "foo bar") + "foo bar") + "")}
        v2 -= int8(v5)
        _ = v23
    } else {
        v9 = 0.0
    }
    for fi24 := 0; fi24 < 2; fi24++ {
        for fi25 := 0; fi25 < 2; fi25++ {
            var v26 float64 = 0.1
            if (1000 >= -572) {
                fmt.Println(v2, v10, v2)
                fmt.Println(v26, fi25, v26)
            }
            v2 = v2
            _ = fi25
            _ = v26
        }
        v27 := []int{}
        _ = fi24
        _ = v27
    }
    v9 = (v9 + float32(-2.25))
    for fi28 := 0; fi28 < 11; fi28++ {
        for fi29 := 0; fi29 < 2; fi29++ {
            switch (fi29 / (((fi28 ^ (fi28 / (((fi28 & (-(v1)))) | 1)))) | 1)) {
            case 703:
                v5 = v5
            }
            var v30 int32 = 891
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    v9 = v9
    var v31 uint64 = 42
    var v32 uint8 = uint8(v31)
    fmt.Println(v31)
    if false {
        delete(v3, v1)
        v33 := []int{v1, 127}
        for fi34 := 0; fi34 < 17; fi34++ {
            if (!((v31 >> (((^(v11))) & 63)) != v31)) {
                _ = len(v4)
                v32 = (v32 & v32)
            }
            if ((0.1 * -2.25) >= 1.5) {
                v11 -= v11
                v1 += 1
                v32 -= (^(v32))
            }
            v4[""] = "hello"
            _ = fi34
        }
        _ = v33
    }
    switch (v1 << ((v11) & 63)) {
    case 2147483647:
        switch (v1 >> (((^(v11))) & 63)) {
        case -2147483648:
            v5 -= v5
        case 128:
            if false {
                switch v1 {
                case 1:
                    v35 := GenS2{F0: (-(v5)), F1: 501, F2: v11}
                    _ = v35
                case 0:
                    v36 := []int{}
                    _ = v36
                }
                v37 := GenS1{F0: "hello"}
                _ = v37
            }
        case 2147483647:
            v38 := GenS2{F0: v5, F1: (v11 % ((v11) | 1)), F2: 42}
            _ = v38
        }
    default:
        v39 := []int{(v1 % (((v1 - (-(v1)))) | 1)), v1, 579}
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
