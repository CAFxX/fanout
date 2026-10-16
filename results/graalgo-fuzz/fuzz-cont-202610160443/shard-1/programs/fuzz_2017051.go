package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int64 }
func genFn0(p0 int, p1 uint64) uint {
    p0 -= (p0 * 128)
    if (p0 < p0) {
        fmt.Println(p0)
        fmt.Println(p0, p0)
        fmt.Println(p0, p1)
    } else {
        p0 += (-(p0))
        switch (p0 >> ((1) & 63)) {
        case 42:
            v4 := GenS1{F0: 468, F1: (p0 - p0)}
            _ = v4
        case -900:
            var v5 int16 = 491
            _ = v5
        default:
            if false {
                for fi6 := 0; fi6 < 17; fi6++ {
                    switch fi6 {
                    case 2147483647:
                        v7 := []int{p0, fi6, 128}
                        _ = v7
                    default:
                        p0 = (fi6 | (fi6 | 0))
                    }
                    _ = fi6
                }
                p0 = p0
                var v8 float32 = 3.14159
                _ = v8
            } else {
                if true {
                    for fi9 := 0; fi9 < 13; fi9++ {
                        p0 = -989
                        if true {
                            v10 := GenS1{F0: 4294967295, F1: fi9}
                            var v11 int = fi9
                            _ = v10
                            _ = v11
                        }
                        p1 -= (p1 >> ((18446744073709551615) & 63))
                        _ = fi9
                    }
                }
                v12 := GenS1{F0: uint(p1), F1: (p0 << ((2) & 63))}
                _ = v12
            }
        }
    }
    p1 = (^(p1))
    _ = p0
    _ = p1
    return 65535
}

func main() {
    v1 := []string{"", "foo bar"}
    var v2 uint8 = 135
    var v3 uint = 733
    if true {
        var v13 bool = (!true)
        v14 := GenS1{F0: (v3 / (((^(v3))) | 1)), F1: 9223372036854775807}
        _ = v13
        _ = v14
    } else {
        for fi15 := 0; fi15 < 18; fi15++ {
            fmt.Println(v3, v2)
            v3 = genFn0(2147483647, 567)
            _ = fi15
        }
        v3 = 42
    }
    if true {
        _ = len(v1)
    }
    var v16 uint = v3
    switch "" {
    case "a":
        v2 = (v2 + 1)
    }
    switch "a" {
    case "a":
        v3 -= (v16 ^ 1)
    }
    var v17 uint = (v3 ^ (v16 << (((v3 + genFn0(-594, 822))) & 63)))
    fmt.Println(v2)
    v18 := GenS2{F0: 1}
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 int = fi19
        switch "a" {
        case "":
            v3 += (v16 & 4294967295)
        case "a":
            if false {
                v17 -= genFn0(v20, uint64(v17))
            }
        case "\n\t":
            var v21 string = (("hello" + "\n\t") + "a")
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    v3 = (v17 ^ genFn0(601, 1))
    for fi22 := 0; fi22 < 6; fi22++ {
        if true {
            if (255 != 0) {
                var v23 bool = false
                v17 = (v16 | v17)
                _ = v23
            } else {
                var v24 uint64 = 723
                switch ((((("hello" + "\n\t") + ("\n\t" + "foo bar")) + ("" + "\n\t")) + "foo bar") + "\n\t") {
                case "foo bar":
                    v16 += v16
                }
                _ = v24
            }
            v25 := GenS1{F0: 4294967295, F1: (fi22 + fi22)}
            v2 -= v2
            _ = v25
        }
        _ = fi22
    }
    if true {
        for fi26 := 0; fi26 < 17; fi26++ {
            for fi27 := 0; fi27 < 1; fi27++ {
                v1 = append(v1, ("hello" + "foo bar"))
                _ = fi27
            }
            switch "\n\t" {
            case "a":
                if len(v1) > 0 {
                    v1[0] = ""
                }
            }
            var v28 uint8 = 200
            _ = fi26
            _ = v28
        }
        var v29 uint16 = 2
        _ = v29
    }
    v30 := GenS2{F0: 1}
    v16 += v16
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
}
