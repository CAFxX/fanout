package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 string }
func genFn0(p0 uint64, p1 int32, p2 int) uint16 {
    switch (p2 * p2) {
    case 321:
        p0 += p0
    }
    v5 := []int{p2, p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2
}
func genFn1() int {
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := []int{fi6, (fi6 - 1), 127}
        _ = fi6
        _ = v7
    }
    switch 1 {
    case -153:
        var v8 uint = 4294967295
        _ = v8
    case -73:
        var v9 float64 = 0.0
        _ = v9
    }
    var v10 bool = (((0.1 - 1.5) <= (3.14159 * (((((1e10 + float64((((1.5 * 1.5) + float32(3.14159)) / float32(1e10)))) * 1e10) * 0.0) * 0.0) * 0.1))) && false)
    _ = v10
    return 0
}

func main() {
    var v1 int16 = 965
    v2 := make(map[int]int)
    var v3 int16 = v1
    v4 := []string{((("foo bar" + "hello") + (("a" + "") + "a")) + "hello")}
    v3 = (-(v3))
    for fi11 := 0; fi11 < 14; fi11++ {
        fmt.Println(fi11)
        _ = fi11
    }
    var v12 int64 = 0
    var v13 uint64 = 18446744073709551615
    v14 := GenS1{F0: 0, F1: 4294967295}
    for fi15 := 0; fi15 < 20; fi15++ {
        fmt.Println(v3)
        _ = fi15
    }
    switch "a" {
    case "hello":
        for fi16 := 0; fi16 < 2; fi16++ {
            var v17 bool = ("hello" <= "foo bar")
            fmt.Println(v3)
            _ = fi16
            _ = v17
        }
    case "":
        for fi18 := 0; fi18 < 2; fi18++ {
            for fi19 := 0; fi19 < 20; fi19++ {
                v2[(-(fi19))] = (-(fi18))
                if len(v4) > 0 {
                    v4[0] = (((("foo bar" + "hello") + (("\n\t" + "") + ("\n\t" + ""))) + "\n\t") + ("hello" + ((("" + ("hello" + "foo bar")) + ((("a" + "\n\t") + "hello") + "\n\t")) + "a")))
                }
                var v20 float64 = -2.25
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
    case "a":
        switch "foo bar" {
        case "\n\t":
            var v21 int64 = 1
            _ = v21
        }
    }
    v22 := GenS1{F0: 2, F1: (v13 | (v13 << ((18446744073709551615) & 63)))}
    var v23 int = -763
    var v24 uint8 = 13
    fmt.Println(v13, v23)
    if len(v4) > 0 {
        v4[0] = "a"
    }
    var v25 int16 = (v3 % ((v3) | 1))
    if (v23 >= (v23 - (v23 - -9223372036854775808))) {
        if (-36 > 42) {
            v26 := GenS2{F0: "\n\t"}
            switch (v23 >> ((2) & 63)) {
            case -244:
                v12 -= 9223372036854775807
            case 42:
                for fi27 := 0; fi27 < 12; fi27++ {
                    v23 = 42
                    if false {
                        for fi28 := 0; fi28 < 14; fi28++ {
                            v29 := GenS1{F0: v13, F1: v13}
                            _ = fi28
                            _ = v29
                        }
                        v30 := GenS1{F0: uint64(v24), F1: 18446744073709551615}
                        var v31 int64 = (v12 >> ((0) & 63))
                        _ = v30
                        _ = v31
                    }
                    _ = fi27
                }
            case -2147483648:
                delete(v2, (v23 / ((v23) | 1)))
            }
            _ = len(v4)
            _ = v26
        } else {
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
            v32 := GenS2{F0: ("hello" + "")}
            _ = v32
        }
    }
    v13 += v13
    v33 := GenS1{F0: v13, F1: (v13 + v13)}
    if (-9223372036854775808 < v12) {
        var v34 string = ("" + "hello")
        _ = v34
    } else {
        var v35 uint = 4294967295
        var v36 uint8 = (v24 + v24)
        _ = v35
        _ = v36
    }
    if (!(42 <= 1)) {
        switch v23 {
        case 727:
            var v37 int = (v23 * (v23 << ((0) & 63)))
            _ = v37
        case -227:
            for fi38 := 0; fi38 < 8; fi38++ {
                var v39 uint64 = uint64(v24)
                _ = fi38
                _ = v39
            }
        default:
            for fi40 := 0; fi40 < 19; fi40++ {
                var v41 uint32 = 600
                _ = len(v4)
                v2[613] = v23
                _ = fi40
                _ = v41
            }
        }
        v12 -= (v12 - (-(v12)))
    } else {
        var v42 float32 = float32((0.1 / 3.14159))
        v24 -= 114
        _ = len(v2)
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v4
}
