package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int8 }
func genFn0(p0 uint32) uint32 {
    var v5 int64 = -59
    v6 := []int{int(p0), 32767, -214}
    switch "\n\t" {
    case "foo bar":
        v5 = v5
    case "a":
        fmt.Println(v5, v5)
    case "\n\t":
        for fi7 := 0; fi7 < 2; fi7++ {
            p0 += p0
            var v8 uint64 = 953
            if (!true) {
                fmt.Println(v5)
            } else {
                v6 = append(v6, (fi7 ^ (^(fi7))))
                v9 := GenS1{F0: "hello", F1: (((v8 == 2) && false) && (true || false)), F2: 127}
                _ = v9
            }
            _ = fi7
            _ = v8
        }
    }
    _ = p0
    _ = v5
    _ = v6
    return 2147483648
}

func main() {
    var v1 uint8 = 1
    var v2 string = ""
    var v3 int64 = 9223372036854775807
    var v4 uint = 65535
    for fi10 := 0; fi10 < 6; fi10++ {
        fmt.Println(v1, v3, v4)
        _ = fi10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        if (float32(1e10) <= -2.25) {
            v4 += v4
            if false {
                var v12 bool = true
                var v13 uint8 = v1
                if true {
                    v3 = v3
                    if v12 {
                        var v14 bool = v12
                        _ = v14
                    }
                    v2 = (v2 + v2)
                }
                _ = v12
                _ = v13
            } else {
                var v15 float32 = 1e10
                var v16 uint64 = 124
                _ = v15
                _ = v16
            }
            fmt.Println(v1, v1, v4)
        }
        var v17 bool = false
        _ = fi11
        _ = v17
    }
    var v18 uint16 = 42
    if false {
        var v19 uint16 = (v18 - (v18 ^ v18))
        _ = v19
    } else {
        v3 = v3
        if ((v4 << ((18446744073709551615) & 63)) != 58) {
            for fi20 := 0; fi20 < 12; fi20++ {
                var v21 int8 = 1
                var v22 uint16 = 32767
                _ = fi20
                _ = v21
                _ = v22
            }
            v23 := []int{9223372036854775807, 9223372036854775807}
            _ = v23
        } else {
            var v24 int64 = v3
            v3 += v24
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        v3 -= (v3 + v3)
        v1 -= 2
        _ = fi25
    }
    v26 := []int{32768}
    var v27 int64 = (v3 % ((-122) | 1))
    for fi28 := 0; fi28 < 4; fi28++ {
        fmt.Println(v1)
        v29 := GenS1{F0: (v2 + v2), F1: (-565 == 1), F2: 0}
        _ = fi28
        _ = v29
    }
    v3 -= v3
    if len(v26) > 0 {
        v26[0] = -1
    }
    var v30 int16 = 1000
    switch "hello" {
    case "hello":
        for fi31 := 0; fi31 < 19; fi31++ {
            if true {
                if len(v26) > 0 {
                    v26[0] = fi31
                }
                var v32 int = fi31
                _ = v32
            }
            var v33 uint = (v4 ^ v4)
            _ = fi31
            _ = v33
        }
    case "foo bar":
        var v34 int8 = -128
        _ = v34
    case "\n\t":
        v30 += (v30 << (((v4 & (v4 >> ((v4) & 63)))) & 63))
    default:
        v35 := GenS1{F0: (v2 + (v2 + (v2 + (v2 + "foo bar")))), F1: (!(((^(v1)) >= (^(v1))) || false)), F2: 100}
        _ = v35
    }
    fmt.Println(v3, v18, v27)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
}
