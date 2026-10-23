package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
func genFn0() bool {
    var v7 string = (("a" + "\n\t") + "\n\t")
    var v8 int16 = -1
    var v9 int16 = v8
    var v10 uint = 2
    var v11 float32 = ((1.5 + 0.1) - -2.25)
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return (1 <= 128)
}

func main() {
    var v1 uint8 = 151
    var v2 float64 = float64(float32(((3.14159 / (-2.25 / (float64(-2.25) * -2.25))) / ((1e10 + ((1e10 - ((0.0 * 0.1) * 3.14159)) / 1.5)) / 1e10))))
    var v3 bool = true
    v4 := make(map[uint32]uint32)
    var v5 int16 = -32768
    v6 := make(map[string]int)
    var v12 uint64 = 42
    var v13 int8 = -42
    switch "a" {
    case "hello":
        v14 := GenS1{F0: 2}
        _ = v14
    case "":
        var v15 uint = 506
        _ = v15
    case "a":
        var v16 int8 = (v13 % ((v13) | 1))
        _ = v16
    default:
        switch ("\n\t" + "hello") {
        case "hello":
            for fi17 := 0; fi17 < 16; fi17++ {
                for fi18 := 0; fi18 < 17; fi18++ {
                    v19 := v6["a"]
                    v19 = (fi18 % ((fi18) | 1))
                    v13 = (v13 << ((950) & 63))
                    _ = fi18
                    _ = v19
                }
                var v20 uint = 632
                _ = fi17
                _ = v20
            }
        case "\n\t":
            fmt.Println(v3, v2, v3)
        case "a":
            fmt.Println(v3, v2)
        }
    }
    if (!(!(4294967295 <= 42))) {
        v13 = -91
    } else {
        switch "\n\t" {
        case "hello":
            v3 = (true && v3)
        case "foo bar":
            v5 = v5
        case "a":
            fmt.Println(v13, v2)
        }
    }
    v1 = (v1 | v1)
    if ((!((-2.25 / 1e10) != float32(v2))) || (!v3)) {
        for fi21 := 0; fi21 < 3; fi21++ {
            v2 = v2
            var v22 uint8 = v1
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 12; fi23++ {
            v24 := GenS1{F0: 974}
            var v25 float64 = 1e10
            for fi26 := 0; fi26 < 13; fi26++ {
                v27 := GenS2{F0: (v13 << ((254) & 63))}
                switch ("hello" + ("foo bar" + "a")) {
                case "hello":
                    fmt.Println(v25)
                    _ = len(v6)
                }
                _ = fi26
                _ = v27
            }
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    v12 -= (^(v12))
    v12 -= 18446744073709551615
    switch ("" + ("hello" + "hello")) {
    case "foo bar":
        var v28 int32 = -2147483648
        _ = v28
    case "\n\t":
        switch ("a" + "hello") {
        case "":
            for fi29 := 0; fi29 < 6; fi29++ {
                v2 -= v2
                _ = fi29
            }
        }
    case "":
        v1 = uint8(v12)
    }
    v5 -= (v5 | int16(v13))
    if (false || (!(1 > -2147483648))) {
        if (int32(v12) >= -773) {
            _ = len(v6)
        } else {
            switch 127 {
            case 589:
                var v30 float32 = 1e10
                _ = v30
            case 890:
                var v31 int64 = 648
                _ = v31
            case -438:
                var v32 bool = (1e10 < (-2.25 - -2.25))
                _ = v32
            default:
                v13 += v13
            }
        }
        var v33 int16 = v5
        _ = v33
    } else {
        if (v3 && ((1 > 0) && false)) {
            var v34 uint32 = 2147483648
            for fi35 := 0; fi35 < 18; fi35++ {
                if ("a" > "") {
                    v2 = v2
                    var v36 uint = 18446744073709551615
                    if v3 {
                        var v37 uint64 = (v12 % (((v12 >> ((v36) & 63))) | 1))
                        var v38 int64 = int64(v13)
                        _ = v37
                        _ = v38
                    } else {
                        for fi39 := 0; fi39 < 15; fi39++ {
                            v40 := GenS2{F0: 1}
                            _ = fi39
                            _ = v40
                        }
                    }
                    _ = v36
                }
                switch "hello" {
                case "hello":
                    if v3 {
                        v41 := GenS1{F0: 65535}
                        v42 := GenS2{F0: (v13 % ((-44) | 1))}
                        _ = v41
                        _ = v42
                    }
                case "foo bar":
                    v34 = (v34 & (v34 * v34))
                case "\n\t":
                    v4[(v34 - (v34 + v34))] = (v34 & (v34 << ((0) & 63)))
                }
                _ = fi35
            }
            _ = v34
        } else {
            var v43 uint64 = (^(v12))
            var v44 string = "hello"
            _ = v43
            _ = v44
        }
    }
    var v45 float32 = ((((0.0 / 3.14159) - ((3.14159 / (((1e10 / (float32(v2) - 3.14159)) * 1e10) - ((0.1 + 1.5) * 1e10))) * (1.5 + (1e10 - (0.0 - (((-2.25 / -2.25) / 3.14159) - (float32((v2 * (v2 - v2))) - 1.5))))))) - 3.14159) - 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
