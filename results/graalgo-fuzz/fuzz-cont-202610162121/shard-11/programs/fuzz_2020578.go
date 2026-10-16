package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int64 }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{0, -270, 127}
    v3 := make(map[uint32]int)
    var v4 uint16 = 648
    v5 := make(map[uint32]int)
    if true {
        if (!false) {
            if true {
                var v6 uint64 = 0
                for fi7 := 0; fi7 < 15; fi7++ {
                    var v8 float64 = float64(1.5)
                    _ = fi7
                    _ = v8
                }
                v6 = (v6 & (v6 / ((4294967295) | 1)))
                _ = v6
            }
            fmt.Println(v4)
            v4 -= (v4 / ((v4) | 1))
        }
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 uint64 = 40
            v11 := GenS1{F0: 42, F1: (v4 + (v4 ^ v4)), F2: 389}
            var v12 int16 = 1000
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    }
    if (false && false) {
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 uint8 = 128
            for fi15 := 0; fi15 < 10; fi15++ {
                var v16 int16 = -32768
                switch "foo bar" {
                case "hello":
                    var v17 float64 = 0.0
                    _ = v17
                }
                if (!(!(true && (4294967295 <= 0)))) {
                    switch "foo bar" {
                    case "\n\t":
                        var v18 float64 = 0.1
                        _ = v18
                    case "a":
                        var v19 uint8 = v14
                        _ = v19
                    }
                    var v20 uint64 = 978
                    _ = v20
                } else {
                    var v21 uint64 = 0
                    _ = v21
                }
                _ = fi15
                _ = v16
            }
            _ = fi13
            _ = v14
        }
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        switch ("foo bar" + ("\n\t" + "hello")) {
        case "hello":
            if len(v2) > 0 {
                v2[0] = fi22
            }
        case "\n\t":
            v4 = v4
        }
        switch ("\n\t" + "") {
        case "\n\t":
            switch "" {
            case "\n\t":
                v4 = (v4 / ((2) | 1))
            case "a":
                v23 := GenS1{F0: 1, F1: v4, F2: 0}
                _ = v23
            default:
                fmt.Println(v4, v4)
            }
        case "":
            var v24 int8 = 42
            _ = v24
        }
        _ = fi22
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        v4 = 32767
        _ = fi25
    }
    v4 = v4
    switch -1 {
    case 42:
        var v26 bool = (((127 == int8(v4)) || false) || (!true))
        _ = v26
    case 2147483647:
        if len(v2) > 0 {
            v2[0] = 841
        }
    default:
        delete(v1, 599)
    }
    v4 = v4
    v4 = (v4 | 631)
    var v27 int = 111
    v4 = (v4 + v4)
    if (!(!true)) {
        v27 += (v27 >> ((18446744073709551615) & 63))
        v28 := GenS1{F0: 42, F1: 821, F2: -1}
        _ = v28
    } else {
        v4 -= 853
        switch (v27 * v27) {
        case 128:
            v27 = 127
        case -9223372036854775808:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 32768:
            var v29 int8 = 42
            _ = v29
        }
        v30 := GenS1{F0: 0, F1: (v4 ^ (v4 / ((v4) | 1))), F2: 9223372036854775807}
        _ = v30
    }
    if false {
        v4 = v4
        var v31 int32 = -926
        v27 += v27
        _ = v31
    } else {
        fmt.Println(v4, v4)
        if len(v2) > 0 {
            v2[0] = (v27 % ((v27) | 1))
        }
        var v32 int16 = 1
        _ = v32
    }
    switch ("\n\t" + ("hello" + (("\n\t" + "\n\t") + ("a" + "")))) {
    case "a":
        fmt.Println(v4)
    }
    v33 := GenS1{F0: 0, F1: (v4 / ((v4) | 1)), F2: 806}
    _ = len(v2)
    if true {
        if true {
            v27 = 1
            v4 = v4
        }
        for fi34 := 0; fi34 < 16; fi34++ {
            v4 = v4
            fmt.Println(fi34, v4)
            _ = fi34
        }
    }
    v27 += v27
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
