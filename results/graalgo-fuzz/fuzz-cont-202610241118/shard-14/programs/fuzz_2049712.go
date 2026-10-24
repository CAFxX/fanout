package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int64 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint64, p1 uint16) uint8 {
    v6 := []int{2147483647, 32768, 127, 9223372036854775807}
    switch 821 {
    case 0:
        var v7 int = -481
        _ = v7
    default:
        v8 := GenS1{F0: 0, F1: (p0 - (^(p0))), F2: 42}
        _ = v8
    }
    v9 := GenS2{F0: 753}
    var v10 bool = ((!((p1 + (p1 << ((671) & 63))) <= 255)) && true)
    for fi11 := 0; fi11 < 20; fi11++ {
        v12 := GenS1{F0: 0, F1: p0, F2: -9223372036854775808}
        p0 = p0
        if (42 < 1) {
            switch ("hello" + (("" + "a") + (("a" + "a") + ("a" + ("hello" + ("" + ("" + "\n\t"))))))) {
            case "":
                switch fi11 {
                case 32767:
                    var v13 int = (fi11 >> ((1) & 63))
                    _ = v13
                default:
                    if (!(!v10)) {
                        v14 := GenS2{F0: -32768}
                        v15 := GenS2{F0: 1}
                        if (!true) {
                            fmt.Println(fi11, p1, v10)
                            if len(v6) > 0 {
                                v6[0] = 307
                            }
                        } else {
                            v6 = append(v6, 0)
                            var v16 uint64 = 2
                            v10 = ((-2.25 + (1e10 + 0.1)) > ((float64(float32(0.1)) - 0.0) / 3.14159))
                            _ = v16
                        }
                        _ = v14
                        _ = v15
                    }
                }
            }
        } else {
            p1 = p1
            var v17 uint16 = 340
            _ = v17
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    _ = v9
    return 2
}
func genFn1() int {
    var v18 uint32 = 511
    v18 -= v18
    for fi19 := 0; fi19 < 10; fi19++ {
        switch fi19 {
        case -2147483648:
            switch "\n\t" {
            case "\n\t":
                for fi20 := 0; fi20 < 8; fi20++ {
                    var v21 uint = 42
                    var v22 string = "foo bar"
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            default:
                fmt.Println(fi19, v18)
            }
        }
        switch "a" {
        case "foo bar":
            v18 += v18
        }
        _ = fi19
    }
    switch 32767 {
    case 42:
        if (false && false) {
            var v23 uint16 = 253
            v24 := []int{271, 2147483647}
            _ = v23
            _ = v24
        }
    case -830:
        if false {
            fmt.Println(v18)
            for fi25 := 0; fi25 < 18; fi25++ {
                var v26 float32 = 1.5
                _ = fi25
                _ = v26
            }
        }
    }
    for fi27 := 0; fi27 < 8; fi27++ {
        switch fi27 {
        case -324:
            if false {
                v28 := []int{fi27, fi27, fi27, fi27}
                _ = v28
            }
        }
        v29 := GenS1{F0: -2147483648, F1: 2, F2: 0}
        v18 = (^(v18))
        _ = fi27
        _ = v29
    }
    _ = v18
    return 423
}

func main() {
    var v1 int16 = 208
    var v2 int64 = 9223372036854775807
    var v3 float64 = float64((0.1 * 3.14159))
    v4 := []int{}
    var v5 int = -2147483648
    v3 = (v3 + 0.1)
    v5 += (v5 | (v5 * 680))
    var v30 float32 = 0.0
    v2 -= 1
    if (v30 > (0.0 - v30)) {
        v1 = v1
        var v31 uint32 = 2147483647
        _ = v31
    }
    switch "\n\t" {
    case "foo bar":
        v32 := GenS1{F0: 42, F1: 2, F2: -1}
        _ = v32
    case "\n\t":
        switch ("\n\t" + "a") {
        case "foo bar":
            if (!false) {
                fmt.Println(v5, v3)
                v30 = 1e10
            }
        }
    case "a":
        var v33 uint16 = 1
        _ = v33
    }
    v30 += float32(v3)
    var v34 string = (("" + ("hello" + "hello")) + "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
}
