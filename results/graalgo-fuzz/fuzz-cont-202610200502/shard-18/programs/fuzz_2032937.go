package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint8 }
type GenS2 struct { F0 string; F1 int64 }
func genFn0() int {
    for fi4 := 0; fi4 < 2; fi4++ {
        switch 2147483647 {
        case 230:
            if false {
                var v5 int64 = 237
                _ = v5
            } else {
                var v6 uint16 = 946
                _ = v6
            }
        }
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 int64 = 228
            _ = fi7
            _ = v8
        }
        _ = fi4
    }
    var v9 string = "a"
    v10 := []int{2147483647, 32768, -2147483648}
    _ = v10
    _ = v9
    return 2147483647
}
func genFn1(p0 int64, p1 int8, p2 string) uint8 {
    p0 = 9223372036854775807
    if false {
        switch ((p2 + ("foo bar" + (p2 + "hello"))) + p2) {
        case "\n\t":
            v11 := []int{127}
            _ = v11
        case "foo bar":
            v12 := GenS1{F0: (1 >= 65535), F1: p2, F2: 116}
            _ = v12
        }
    } else {
        if true {
            if true {
                if (3.14159 < 1e10) {
                    p0 = (p0 << ((4294967295) & 63))
                } else {
                    if false {
                        var v13 uint = 255
                        _ = v13
                    } else {
                        p0 += (p0 / (((p0 % ((p0) | 1))) | 1))
                        v14 := []int{}
                        if len(v14) > 0 {
                            v14[0] = 127
                        }
                        _ = v14
                    }
                }
            }
        }
    }
    p2 = ("\n\t" + "foo bar")
    p1 = -128
    var v15 uint8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return (v15 << ((18446744073709551615) & 63))
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = -1
    var v3 int64 = 42
    var v16 int8 = (-(v2))
    v17 := GenS1{F0: true, F1: ("" + "foo bar"), F2: 128}
    v1[2147483647] = true
    fmt.Println(v2, v3)
    delete(v1, 2)
    switch "hello" {
    case "\n\t":
        var v18 uint64 = 141
        _ = v18
    case "hello":
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 float32 = (float32(3.14159) * ((1.5 / ((0.1 - 1.5) / float32(3.14159))) * float32((0.1 + 0.0))))
            _ = fi19
            _ = v20
        }
    default:
        if false {
            if true {
                for fi21 := 0; fi21 < 11; fi21++ {
                    v16 -= (v16 << ((2) & 63))
                    v22 := GenS1{F0: ("\n\t" <= "hello"), F1: ("a" + ("" + "foo bar")), F2: 127}
                    _ = fi21
                    _ = v22
                }
                v2 = v2
            } else {
                v2 += (v2 + v16)
            }
        } else {
            v23 := []int{}
            var v24 uint = 1
            var v25 int = -9223372036854775808
            _ = v23
            _ = v24
            _ = v25
        }
    }
    var v26 uint16 = 65535
    if (!true) {
        if false {
            v27 := v1[523]
            _ = v27
        } else {
            _ = len(v1)
            v28 := []int{}
            _ = v28
        }
    }
    v29 := GenS1{F0: (-1 < 42), F1: "hello", F2: genFn1((^(v3)), -1, "a")}
    switch genFn0() {
    case 208:
        v30 := GenS2{F0: "", F1: v3}
        _ = v30
    case -9223372036854775808:
        var v31 int16 = -807
        _ = v31
    }
    v2 = (v16 % ((v16) | 1))
    _ = len(v1)
    for fi32 := 0; fi32 < 17; fi32++ {
        if false {
            var v33 float32 = -2.25
            for fi34 := 0; fi34 < 16; fi34++ {
                var v35 uint64 = 805
                fmt.Println(fi32, v33, fi32)
                var v36 uint16 = v26
                _ = fi34
                _ = v35
                _ = v36
            }
            _ = v33
        }
        _ = fi32
    }
    v2 -= -128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v29
    _ = v3
}
