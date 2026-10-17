package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }
func genFn0(p0 uint16, p1 bool, p2 uint8) uint8 {
    if (0 != 4294967295) {
        v7 := []int{-9223372036854775808}
        _ = v7
    }
    v8 := []int{}
    fmt.Println(p1, p2)
    var v9 int = 9223372036854775807
    if (!((p0 << ((4294967295) & 63)) != (p0 | p0))) {
        for fi10 := 0; fi10 < 1; fi10++ {
            for fi11 := 0; fi11 < 10; fi11++ {
                _ = len(v8)
                if (!p1) {
                    var v12 float32 = 0.1
                    if (!(!p1)) {
                        v12 = (v12 / v12)
                        p1 = (p1 && true)
                        var v13 uint = 42
                        _ = v13
                    }
                    _ = v12
                }
                _ = fi11
            }
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi10
        }
        p2 -= 0
        fmt.Println(v9, p1, p2)
    } else {
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 uint64 = 0
            _ = fi14
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return p2
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]int)
    var v3 uint16 = 927
    var v4 string = "\n\t"
    var v5 uint64 = 1
    var v6 string = (v4 + v4)
    v3 = (v3 - uint16(v5))
    v16 := GenS1{F0: 936, F1: 1, F2: v4}
    fmt.Println(v4, v3, v6)
    var v17 uint = 65535
    for fi18 := 0; fi18 < 13; fi18++ {
        v19 := GenS1{F0: 2147483648, F1: (fi18 & fi18), F2: ""}
        v20 := v1[(fi18 * (fi18 % ((fi18) | 1)))]
        v21 := GenS1{F0: 2147483648, F1: (v20 * (v20 + 42)), F2: ""}
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    switch -24 {
    case 1:
        v22 := GenS1{F0: 2147483647, F1: 42, F2: (v6 + ((((((("a" + (v6 + ("" + "foo bar"))) + v6) + ("" + v6)) + "\n\t") + (("\n\t" + (((((("a" + (v6 + ((("foo bar" + ((((v4 + v6) + v6) + v6) + "\n\t")) + v4) + "\n\t"))) + "foo bar") + (v6 + v6)) + "a") + "foo bar") + "\n\t")) + "\n\t")) + v6) + ((v4 + v4) + "hello")))}
        _ = v22
    case -9223372036854775808:
        if true {
            fmt.Println(v3, v4, v17)
        } else {
            v23 := GenS2{F0: (v3 * (v3 + (v3 | (v3 % (((^(v3))) | 1)))))}
            if (false || (false || (789 > uint32(v17)))) {
                for fi24 := 0; fi24 < 3; fi24++ {
                    v25 := GenS3{F0: v17}
                    _ = fi24
                    _ = v25
                }
            } else {
                fmt.Println(v3, v5, v4)
            }
            switch -1 {
            case 2147483647:
                v5 = v5
            case -328:
                fmt.Println(v5, v3)
            default:
                if false {
                    v17 -= 255
                } else {
                    var v26 float32 = 0.0
                    _ = v26
                }
            }
            _ = v23
        }
    default:
        v27 := GenS2{F0: 32767}
        _ = v27
    }
    var v28 int32 = 0
    fmt.Println(v6)
    v6 = v6
    switch v4 {
    case "hello":
        var v29 int16 = -262
        _ = v29
    case "\n\t":
        var v30 int16 = 0
        _ = v30
    case "":
        v31 := []int{-1}
        _ = v31
    default:
        fmt.Println(v17, v5)
    }
    v5 += v5
    delete(v2, -1)
    v6 = ""
    if false {
        v28 += v28
        var v32 uint16 = 338
        for fi33 := 0; fi33 < 1; fi33++ {
            v34 := []int{}
            var v35 uint = v17
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
