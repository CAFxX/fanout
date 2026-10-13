package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 int64; F2 int32 }
func genFn0(p0 uint16) int16 {
    p0 = (p0 & (p0 / ((32767) | 1)))
    if (true && (!false)) {
        v4 := []int{1, 0}
        p0 += (p0 * (p0 & p0))
        _ = len(v4)
        _ = v4
    } else {
        if true {
            var v5 uint32 = 349
            switch ((("foo bar" + (("hello" + "foo bar") + "")) + "hello") + ("\n\t" + "foo bar")) {
            case "\n\t":
                v6 := []int{0}
                _ = v6
            case "hello":
                p0 = (p0 ^ 42)
            }
            _ = v5
        }
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 int32 = -1
            v9 := []int{fi7, fi7, fi7, (fi7 - (fi7 ^ fi7))}
            _ = fi7
            _ = v8
            _ = v9
        }
        switch 4294967295 {
        case -2147483648:
            for fi10 := 0; fi10 < 2; fi10++ {
                switch fi10 {
                case 128:
                    p0 = p0
                }
                _ = fi10
            }
        case -9223372036854775808:
            if false {
                p0 -= p0
            }
        }
    }
    fmt.Println(p0)
    _ = p0
    return 140
}

func main() {
    var v1 float64 = 1.5
    var v2 int32 = -252
    var v3 bool = ((-2.25 + 0.0) != (v1 + 1.5))
    v11 := []int{}
    v11 = append(v11, -1)
    v2 += v2
    v12 := GenS1{F0: 76, F1: 0, F2: (v2 ^ (v2 << ((4294967295) & 63)))}
    switch 32767 {
    case -436:
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    default:
        for fi13 := 0; fi13 < 6; fi13++ {
            fmt.Println(v1)
            _ = fi13
        }
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        fmt.Println(v2)
        _ = fi14
    }
    var v15 uint16 = 985
    v16 := GenS1{F0: 673, F1: 2, F2: v2}
    switch -823 {
    case 127:
        var v17 uint16 = (^(v15))
        _ = v17
    case 264:
        v18 := GenS2{F0: 2}
        _ = v18
    case -40:
        v2 += -2147483648
    default:
        var v19 uint32 = 94
        _ = v19
    }
    v2 -= (v2 - -869)
    if v3 {
        if v3 {
            var v20 uint32 = 1
            for fi21 := 0; fi21 < 17; fi21++ {
                v3 = v3
                v22 := GenS1{F0: 65535, F1: 18446744073709551615, F2: -878}
                var v23 float64 = (v1 / 1.5)
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = v20
        } else {
            var v24 int = 0
            for fi25 := 0; fi25 < 9; fi25++ {
                v24 = fi25
                switch 665 {
                case 0:
                    var v26 string = ""
                    _ = v26
                case 9223372036854775807:
                    for fi27 := 0; fi27 < 14; fi27++ {
                        var v28 int = (-(v24))
                        var v29 int16 = genFn0(612)
                        v29 += (v29 ^ (v29 - v29))
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                }
                switch (("" + "\n\t") + "hello") {
                case "hello":
                    v30 := GenS1{F0: 893, F1: 18446744073709551615, F2: 42}
                    _ = v30
                case "\n\t":
                    v15 += 76
                case "foo bar":
                    v11 = append(v11, fi25)
                default:
                    var v31 uint16 = (v15 ^ 65535)
                    _ = v31
                }
                _ = fi25
            }
            switch v24 {
            case 0:
                var v32 int64 = -1
                _ = v32
            case 32767:
                for fi33 := 0; fi33 < 18; fi33++ {
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    var v34 uint8 = 1
                    _ = fi33
                    _ = v34
                }
            case 32768:
                v35 := GenS3{F0: v24, F1: 0, F2: v2}
                _ = v35
            }
            _ = v24
        }
        var v36 int8 = 127
        _ = v36
    } else {
        for fi37 := 0; fi37 < 19; fi37++ {
            switch "a" {
            case "a":
                var v38 uint64 = 2
                _ = v38
            case "hello":
                v11 = append(v11, fi37)
            default:
                v39 := GenS1{F0: 4294967295, F1: 4294967295, F2: 0}
                _ = v39
            }
            v40 := GenS1{F0: 1, F1: 1, F2: (v2 >> ((174) & 63))}
            _ = fi37
            _ = v40
        }
        var v41 float64 = float64(0.0)
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
