package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 uint16; F2 bool }
type GenS3 struct { F0 int64 }
func genFn0(p0 string) uint16 {
    var v5 string = p0
    v5 = (p0 + v5)
    switch (p0 + p0) {
    case "foo bar":
        if (!(!false)) {
            v6 := GenS1{F0: false}
            var v7 uint64 = 2
            v8 := GenS2{F0: -2147483648, F1: 32767, F2: ("foo bar" > (p0 + (((v5 + ("hello" + v5)) + "\n\t") + "a")))}
            _ = v6
            _ = v7
            _ = v8
        } else {
            p0 = v5
            for fi9 := 0; fi9 < 17; fi9++ {
                fmt.Println(p0)
                _ = fi9
            }
            if (true || false) {
                p0 = "foo bar"
            }
        }
    case "\n\t":
        v10 := []int{0, 304, 32768}
        _ = v10
    case "hello":
        for fi11 := 0; fi11 < 18; fi11++ {
            p0 = "hello"
            if true {
                switch 706 {
                case 1:
                    v5 = (p0 + ((v5 + "foo bar") + ""))
                case 32767:
                    var v12 int = 128
                    _ = v12
                case 997:
                    fmt.Println(v5, p0)
                }
                v5 = (v5 + ("foo bar" + p0))
                fmt.Println(fi11, fi11, fi11)
            }
            if false {
                v5 = ""
                v13 := []int{(fi11 - fi11), fi11, (^(fi11))}
                _ = v13
            } else {
                v5 = p0
                for fi14 := 0; fi14 < 9; fi14++ {
                    fmt.Println(p0)
                    var v15 float64 = (-2.25 + 1e10)
                    if (100 != 127) {
                        var v16 int64 = 574
                        _ = v16
                    }
                    _ = fi14
                    _ = v15
                }
                if true {
                    var v17 int32 = 2147483647
                    _ = v17
                }
            }
            _ = fi11
        }
    }
    var v18 int64 = -1
    _ = p0
    _ = v18
    _ = v5
    return 65535
}

func main() {
    var v1 int32 = -524
    var v2 uint = 65535
    var v3 string = "foo bar"
    var v4 uint16 = 55
    v19 := []int{-9223372036854775808, -9223372036854775808, 9223372036854775807, 9223372036854775807}
    if false {
        for fi20 := 0; fi20 < 1; fi20++ {
            _ = len(v19)
            var v21 bool = true
            _ = fi20
            _ = v21
        }
    }
    v22 := GenS1{F0: true}
    v1 = (v1 % ((-2147483648) | 1))
    var v23 int8 = -1
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    v24 := GenS1{F0: (v4 == (v4 % ((791) | 1)))}
    for fi25 := 0; fi25 < 11; fi25++ {
        for fi26 := 0; fi26 < 10; fi26++ {
            var v27 float32 = (-2.25 - float32(3.14159))
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        switch "foo bar" {
        case "\n\t":
            if ((fi28 - fi28) != int(v4)) {
                v23 += (v23 >> ((42) & 63))
                v23 = -42
            }
        case "a":
            v19 = append(v19, fi28)
        case "":
            v1 -= (v1 << (((v2 % (((v2 ^ 0)) | 1))) & 63))
        default:
            v29 := GenS3{F0: 0}
            _ = v29
        }
        var v30 uint16 = (v4 - (^(v4)))
        var v31 string = v3
        _ = fi28
        _ = v30
        _ = v31
    }
    switch 128 {
    case -587:
        _ = len(v19)
    case -9223372036854775808:
        for fi32 := 0; fi32 < 14; fi32++ {
            var v33 int = fi32
            _ = fi32
            _ = v33
        }
    }
    var v34 float32 = 0.0
    v23 -= v23
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v34
    _ = v4
}
